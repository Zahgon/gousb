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

// InterfaceDesc contains information about a USB interface, extracted from
// the descriptor.
type InterfaceDesc struct {
	// Number is the number of this interface.
	Number int
	// AltSettings is a list of alternate settings supported by the interface.
	AltSettings []InterfaceSetting
}

func (i *InterfaceDesc) altSetting(alt int) (*InterfaceSetting, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// String returns a human-readable description of the interface descriptor and
// its alternate settings.
func (i InterfaceDesc) String() string { _ = "STUB: not implemented"; return "" }

// InterfaceSetting contains information about a USB interface with a particular
// alternate setting, extracted from the descriptor.
type InterfaceSetting struct {
	// Number is the number of this interface, the same as in InterfaceDesc.
	Number int
	// Alternate is the number of this alternate setting.
	Alternate int
	// Class is the USB-IF (Implementers Forum) class code, as defined by the USB spec.
	Class Class
	// SubClass is the USB-IF (Implementers Forum) subclass code, as defined by the USB spec.
	SubClass Class
	// Protocol is USB protocol code, as defined by the USB spe.c
	Protocol Protocol
	// Endpoints enumerates the endpoints available on this interface with
	// this alternate setting.
	Endpoints map[EndpointAddress]EndpointDesc

	iInterface int // index of a string descriptor describing this interface.
}

func (a InterfaceSetting) sortedEndpointIds() []string { _ = "STUB: not implemented"; return nil }

// String returns a human-readable description of the particular
// alternate setting of an interface.
func (a InterfaceSetting) String() string { _ = "STUB: not implemented"; return "" }

// Interface is a representation of a claimed interface with a particular setting.
// To access device endpoints use InEndpoint() and OutEndpoint() methods.
// The interface should be Close()d after use.
type Interface struct {
	Setting InterfaceSetting

	config *Config
}

func (i *Interface) String() string { _ = "STUB: not implemented"; return "" }

// Close releases the interface.
func (i *Interface) Close() { _ = "STUB: not implemented"; return }

func (i *Interface) openEndpoint(epAddr EndpointAddress) (*endpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InEndpoint prepares an IN endpoint for transfer.
func (i *Interface) InEndpoint(epNum int) (*InEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OutEndpoint prepares an OUT endpoint for transfer.
func (i *Interface) OutEndpoint(epNum int) (*OutEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
