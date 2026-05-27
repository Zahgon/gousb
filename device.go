// Copyright 2013 Google Inc.  All rights reserved.
// Copyright 2016 the gousb Authors.  All rights reserved.
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

// DeviceDesc is a representation of a USB device descriptor.
type DeviceDesc struct {
	// The bus on which the device was detected
	Bus int
	// The address of the device on the bus
	Address int
	// The negotiated operating speed for the device
	Speed Speed
	// The pyhsical port on the parent hub on which the device is connected.
	// Ports are numbered from 1, excepting root hub devices which are always 0.
	Port int
	// Physical path of connected parent ports, starting at the root hub device.
	// A path length of 0 represents a root hub device,
	// a path length of 1 represents a device directly connected to a root hub,
	// a path length of 2 or more are connected to intermediate hub devices.
	// e.g. [1,2,3] represents a device connected to port 3 of a hub connected
	// to port 2 of a hub connected to port 1 of a root hub.
	Path []int

	// Version information
	Spec   BCD // USB Specification Release Number
	Device BCD // The device version

	// Product information
	Vendor  ID // The Vendor identifer
	Product ID // The Product identifier

	// Protocol information
	Class                Class    // The class of this device
	SubClass             Class    // The sub-class (within the class) of this device
	Protocol             Protocol // The protocol (within the sub-class) of this device
	MaxControlPacketSize int      // Maximum size of the control transfer

	// Configuration information
	Configs map[int]ConfigDesc

	iManufacturer int // The Manufacturer descriptor index
	iProduct      int // The Product descriptor index
	iSerialNumber int // The SerialNumber descriptor index
}

// String returns a human-readable version of the device descriptor.
func (d *DeviceDesc) String() string { _ = "STUB: not implemented"; return "" }

func (d *DeviceDesc) sortedConfigIds() []int { _ = "STUB: not implemented"; return nil }

func (d *DeviceDesc) cfgDesc(cfgNum int) (*ConfigDesc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Device represents an opened USB device.
// Device allows sending USB control commands through the Command() method.
// For data transfers select a device configuration through a call to
// Config().
// A Device must be Close()d after use.
type Device struct {
	handle *libusbDevHandle
	ctx    *Context

	// Embed the device information for easy access
	Desc *DeviceDesc
	// Timeout for control commands
	ControlTimeout time.Duration

	// Claimed config
	mu      sync.Mutex
	claimed *Config

	// Handle AutoDetach in this library
	autodetach bool
}

// String represents a human readable representation of the device.
func (d *Device) String() string { _ = "STUB: not implemented"; return "" }

// Reset performs a USB port reset to reinitialize a device.
func (d *Device) Reset() error { _ = "STUB: not implemented"; return nil }

// ActiveConfigNum returns the config id of the active configuration.
// The value corresponds to the ConfigInfo.Config field of one of the
// ConfigInfos of this Device.
func (d *Device) ActiveConfigNum() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Config returns a USB device set to use a particular config.
// The cfgNum provided is the config id (not the index) of the configuration to
// set, which corresponds to the ConfigInfo.Config field.
// USB supports only one active config per device at a time. Config claims the
// device before setting the desired config and keeps it locked until Close is
// called.
// A claimed config needs to be Close()d after use.
func (d *Device) Config(cfgNum int) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// DefaultInterface opens interface #0 with alternate setting #0 of the currently active
// config. It's intended as a shortcut for devices that have the simplest
// interface of a single config, interface and alternate setting.
// The done func should be called to release the claimed interface and config.
func (d *Device) DefaultInterface() (intf *Interface, done func(), err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Control sends a control request to the device.
func (d *Device) Control(rType, request uint8, val, idx uint16, data []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close closes the device.
func (d *Device) Close() error { _ = "STUB: not implemented"; return nil }

// GetStringDescriptor returns a device string descriptor with the given index
// number. The first supported language is always used and the returned
// descriptor string is converted to ASCII (non-ASCII characters are replaced
// with "?").
func (d *Device) GetStringDescriptor(descIndex int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// string descriptor index value of 0 indicates no string descriptor.

// Manufacturer returns the device's manufacturer name.
// GetStringDescriptor's string conversion rules apply.
func (d *Device) Manufacturer() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Product returns the device's product name.
// GetStringDescriptor's string conversion rules apply.
func (d *Device) Product() (string, error) { _ = "STUB: not implemented"; return "", nil }

// SerialNumber returns the device's serial number.
// GetStringDescriptor's string conversion rules apply.
func (d *Device) SerialNumber() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ConfigDescription returns the description of the selected device
// configuration. GetStringDescriptor's string conversion rules apply.
func (d *Device) ConfigDescription(cfg int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// InterfaceDescription returns the description of the selected interface and
// its alternate setting in a selected configuration. GetStringDescriptor's
// string conversion rules apply.
func (d *Device) InterfaceDescription(cfgNum, intfNum, altNum int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetAutoDetach enables/disables automatic kernel driver detachment.
// When autodetach is enabled gousb will automatically detach the kernel driver
// on the interface and reattach it when releasing the interface.
// Automatic kernel driver detachment is disabled on newly opened device handles by default.
func (d *Device) SetAutoDetach(autodetach bool) error { _ = "STUB: not implemented"; return nil }
