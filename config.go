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
)

// ConfigDesc contains the information about a USB device configuration,
// extracted from the device descriptor.
type ConfigDesc struct {
	// Number is the configuration number.
	Number int
	// SelfPowered is true if the device is powered externally, i.e. not
	// drawing power from the USB bus.
	SelfPowered bool
	// RemoteWakeup is true if the device supports remote wakeup, i.e.
	// an external signal that will wake up a suspended USB device. An example
	// might be a keyboard that can wake up through a keypress after
	// the host put it in suspend mode. Note that gousb does not support
	// device power management, RemoteWakeup only refers to the reported device
	// capability.
	RemoteWakeup bool
	// MaxPower is the maximum current the device draws from the USB bus
	// in this configuration.
	MaxPower Milliamperes
	// Interfaces has a list of USB interfaces available in this configuration.
	Interfaces []InterfaceDesc

	iConfiguration int // index of a string descriptor describing this configuration
}

// String returns the human-readable description of the configuration descriptor.
func (c ConfigDesc) String() string { _ = "STUB: not implemented"; return "" }

func (c ConfigDesc) intfDesc(num int) (*InterfaceDesc, error) {
	_ = "STUB: not implemented"
	// In an ideal world, interfaces in the descriptor would be numbered
	// contiguously starting from 0, as required by the specification. In the
	// real world however the specification is sometimes ignored:
	// https://github.com/google/gousb/issues/65
	return nil, nil
}

// Config represents a USB device set to use a particular configuration.
// Only one Config of a particular device can be used at any one time.
// To access device endpoints, claim an interface and it's alternate
// setting number through a call to Interface().
type Config struct {
	Desc ConfigDesc

	dev *Device

	// Claimed interfaces
	mu      sync.Mutex
	claimed map[int]bool
}

// Close releases the underlying device, allowing the caller to switch the device to a different configuration.
func (c *Config) Close() error { _ = "STUB: not implemented"; return nil }

// String returns the human-readable description of the configuration.
func (c *Config) String() string { _ = "STUB: not implemented"; return "" }

// Interface claims and returns an interface on a USB device.
// num specifies the number of an interface to claim, and alt specifies the
// alternate setting number for that interface.
func (c *Config) Interface(num, alt int) (*Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Claim the interface

// Select an alternate setting if needed (device has multiple alternate settings).
