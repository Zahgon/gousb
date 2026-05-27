// Copyright 2017 the gousb Authors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gousb

import (
	"sync"
	"time"
)

/*
#cgo pkg-config: libusb-1.0
#include <libusb.h>

int gousb_compact_iso_data(struct libusb_transfer *xfer, unsigned char *status);
struct libusb_transfer *gousb_alloc_transfer_and_buffer(int bufLen, int numIsoPackets);
void gousb_free_transfer_and_buffer(struct libusb_transfer *xfer);
int submit(struct libusb_transfer *xfer);
void gousb_set_debug(libusb_context *ctx, int lvl);
*/
import "C"

type libusbContext C.libusb_context
type libusbDevice C.libusb_device
type libusbDevHandle C.libusb_device_handle
type libusbTransfer C.struct_libusb_transfer
type libusbEndpoint C.struct_libusb_endpoint_descriptor

func (ep libusbEndpoint) endpointDesc(dev *DeviceDesc) EndpointDesc {
	_ = "STUB: not implemented"
	return *new(EndpointDesc)
}

// bits 0-10 identify the packet size, bits 11-12 are the number of additional transactions per microframe.
// Don't use libusb_get_max_iso_packet_size, as it has a bug where it returns the same value
// regardless of alternative setting used, where different alternative settings might define different
// max packet sizes.
// See http://libusb.org/ticket/77 for more background.

// If the device conforms to USB1.x:
//   Interval for polling endpoint for data transfers. Expressed in
//   milliseconds.
//   This field is ignored for bulk and control endpoints. For
//   isochronous endpoints this field must be set to 1. For interrupt
//   endpoints, this field may range from 1 to 255.
// Note: in low-speed mode, isochronous transfers are not supported.

// If the device conforms to USB[23].x and the device is in low or full
// speed mode:
//   Interval for polling endpoint for data transfers.  Expressed in
//   frames (1ms)
//   For full-speed isochronous endpoints, the value of this field should
//   be 1.
//   For full-/low-speed interrupt endpoints, the value of this field may
//   be from 1 to 255.
// Note: in low-speed mode, isochronous transfers are not supported.

// If the device conforms to USB[23].x and the device is in high speed
// mode:
//   Interval is expressed in microframe units (125 µs).
//   For high-speed bulk/control OUT endpoints, the bInterval must
//   specify the maximum NAK rate of the endpoint. A value of 0 indicates
//   the endpoint never NAKs. Other values indicate at most 1 NAK each
//   bInterval number of microframes. This value must be in the range
//   from 0 to 255.

// If the device conforms to USB[23].x and the device is in high speed
// mode:
//   For high-speed isochronous endpoints, this value must be in
//   the range from 1 to 16. The bInterval value is used as the exponent
//   for a 2bInterval-1 value; e.g., a bInterval of 4 means a period
//   of 8 (2^(4-1)).
//   For high-speed interrupt endpoints, the bInterval value is used as
//   the exponent for a 2bInterval-1 value; e.g., a bInterval of 4 means
//   a period of 8 (2^(4-1)). This value must be from 1 to 16.
// If the device conforms to USB3.x and the device is in SuperSpeed mode:
//   Interval for servicing the endpoint for data transfers. Expressed in
//   125-µs units.
//   For Enhanced SuperSpeed isochronous and interrupt endpoints, this
//   value shall be in the range from 1 to 16. However, the valid ranges
//   are 8 to 16 for Notification type Interrupt endpoints. The bInterval
//   value is used as the exponent for a 2(^bInterval-1) value; e.g., a
//   bInterval of 4 means a period of 8 (2^(4-1) → 2^3 → 8).
//   This field is reserved and shall not be used for Enhanced SuperSpeed
//   bulk or control endpoints.

// libusbIntf is a set of trivial idiomatic Go wrappers around libusb C functions.
// The underlying code is generally not testable or difficult to test,
// since libusb interacts directly with the host USB stack.
//
// All functions here should operate on types defined on C.libusb* data types,
// and occasionally on convenience data types (like TransferType or DeviceDesc).
type libusbIntf interface {
	// context
	init() (*libusbContext, error)
	handleEvents(*libusbContext, <-chan struct{})
	getDevices(*libusbContext) ([]*libusbDevice, error)
	exit(*libusbContext) error
	setDebug(*libusbContext, int)

	// device
	dereference(*libusbDevice)
	getDeviceDesc(*libusbDevice) (*DeviceDesc, error)
	open(*libusbDevice) (*libusbDevHandle, error)
	wrapSysDevice(*libusbContext, uintptr) (*libusbDevHandle, error)

	close(*libusbDevHandle)
	reset(*libusbDevHandle) error
	control(*libusbDevHandle, time.Duration, uint8, uint8, uint16, uint16, []byte) (int, error)
	getConfig(*libusbDevHandle) (uint8, error)
	setConfig(*libusbDevHandle, uint8) error
	getStringDesc(*libusbDevHandle, int) (string, error)
	setAutoDetach(*libusbDevHandle, int) error
	detachKernelDriver(*libusbDevHandle, uint8) error
	getDevice(*libusbDevHandle) *libusbDevice

	// interface
	claim(*libusbDevHandle, uint8) error
	release(*libusbDevHandle, uint8)
	setAlt(*libusbDevHandle, uint8, uint8) error

	// transfer
	alloc(*libusbDevHandle, *EndpointDesc, int, int, chan struct{}) (*libusbTransfer, error)
	cancel(*libusbTransfer) error
	submit(*libusbTransfer) error
	buffer(*libusbTransfer) []byte
	data(*libusbTransfer) (int, TransferStatus)
	free(*libusbTransfer)
	setIsoPacketLengths(*libusbTransfer, uint32)
}

// libusbImpl is an implementation of libusbIntf using real CGo-wrapped libusb.
type libusbImpl struct {
	discovery DeviceDiscovery
}

func (impl libusbImpl) init() (*libusbContext, error) {
	var ctx *C.libusb_context

	var libusbOpts [4]C.struct_libusb_init_option // fixed to 4 - there are maximum 4 options
	nOpts := 0
	if impl.discovery == DisableDeviceDiscovery {
		libusbOpts[nOpts].option = C.LIBUSB_OPTION_NO_DEVICE_DISCOVERY
		nOpts++
	}

	if err := fromErrNo(C.libusb_init_context(&ctx, &(libusbOpts[0]), C.int(nOpts))); err != nil {
		return nil, err
	}

	return (*libusbContext)(ctx), nil
}

func (libusbImpl) handleEvents(c *libusbContext, done <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// handler can be interrupted by a signal and this doesn't indicate an error, we'll retry on the next loop iteration

func (libusbImpl) getDevices(ctx *libusbContext) ([]*libusbDevice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// devices must be dereferenced by the caller to prevent memory leaks.

func (libusbImpl) wrapSysDevice(ctx *libusbContext, fd uintptr) (*libusbDevHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (libusbImpl) getDevice(d *libusbDevHandle) *libusbDevice {
	_ = "STUB: not implemented"
	return nil
}

func (libusbImpl) exit(c *libusbContext) error { _ = "STUB: not implemented"; return nil }

func (libusbImpl) setDebug(c *libusbContext, lvl int) { _ = "STUB: not implemented"; return }

func (libusbImpl) getDeviceDesc(d *libusbDevice) (*DeviceDesc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Defaults to port = 0, path = [] for root device

// Enumerate configurations

// at GenX speeds MaxPower is expressed in units of 8mA, not 2mA.

// a map of interface numbers to a set of alternate settings numbers

func (libusbImpl) dereference(d *libusbDevice) { _ = "STUB: not implemented"; return }

func (libusbImpl) open(d *libusbDevice) (*libusbDevHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (libusbImpl) close(d *libusbDevHandle) { _ = "STUB: not implemented"; return }

func (libusbImpl) reset(d *libusbDevHandle) error { _ = "STUB: not implemented"; return nil }

func (libusbImpl) control(d *libusbDevHandle, timeout time.Duration, rType, request uint8, val, idx uint16, data []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (libusbImpl) getConfig(d *libusbDevHandle) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (libusbImpl) setConfig(d *libusbDevHandle, cfg uint8) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(sebek): device string descriptors are natively in UTF16 and support
// multiple languages. get_string_descriptor_ascii uses always the first
// language and discards non-ascii bytes. We could do better if needed.
func (libusbImpl) getStringDesc(d *libusbDevHandle, index int) (string, error) {
	_ = "STUB: not implemented"
	// allocate 200-byte array limited the length of string descriptor
	return "", nil
}

// get string descriptor from libusb. if errno < 0 then there are any errors.
// if errno >= 0; it is a length of result string descriptor

func (libusbImpl) setAutoDetach(d *libusbDevHandle, val int) error {
	_ = "STUB: not implemented"
	return nil
}

func (libusbImpl) detachKernelDriver(d *libusbDevHandle, iface uint8) error {
	_ = "STUB: not implemented"
	return nil
}

// ErrorNotSupported is returned in non linux systems
// ErrorNotFound is returned if libusb's driver is already attached to the device

func (libusbImpl) claim(d *libusbDevHandle, iface uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (libusbImpl) release(d *libusbDevHandle, iface uint8) { _ = "STUB: not implemented"; return }

func (libusbImpl) setAlt(d *libusbDevHandle, iface, setup uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (libusbImpl) alloc(d *libusbDevHandle, ep *EndpointDesc, isoPackets int, bufLen int, done chan struct{}) (*libusbTransfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (libusbImpl) cancel(t *libusbTransfer) error { _ = "STUB: not implemented"; return nil }

func (libusbImpl) submit(t *libusbTransfer) error { _ = "STUB: not implemented"; return nil }

func (libusbImpl) buffer(t *libusbTransfer) []byte {
	_ = "STUB: not implemented"
	// TODO(go1.10?): replace with more user-friendly construct once
	// one exists. https://github.com/golang/go/issues/13656
	return nil
}

func (libusbImpl) data(t *libusbTransfer) (int, TransferStatus) {
	_ = "STUB: not implemented"
	return 0, *new(TransferStatus)
}

func (libusbImpl) free(t *libusbTransfer) { _ = "STUB: not implemented"; return }

func (libusbImpl) setIsoPacketLengths(t *libusbTransfer, length uint32) {
	_ = "STUB: not implemented"
	return
}

// xferDoneMap keeps a map of done callback channels for all allocated transfers.
var xferDoneMap = struct {
	m map[*libusbTransfer]chan struct{}
	sync.RWMutex
}{
	m: make(map[*libusbTransfer]chan struct{}),
}

//export xferCallback
func xferCallback(xfer *C.struct_libusb_transfer) { _ = "STUB: not implemented"; return }

// for benchmarking of method on implementation vs vanilla function.
func libusbSetDebug(c *libusbContext, lvl int) { _ = "STUB: not implemented"; return }

// for obtaining unique CGo pointers.
func newDevicePointer() *libusbDevice { _ = "STUB: not implemented"; return nil }

func newFakeTransferPointer() *libusbTransfer { _ = "STUB: not implemented"; return nil }

func newContextPointer() *libusbContext { _ = "STUB: not implemented"; return nil }

func newDevHandlePointer() *libusbDevHandle { _ = "STUB: not implemented"; return nil }
