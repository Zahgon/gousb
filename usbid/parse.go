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

package usbid

import (
	"io"

	"github.com/google/gousb"
)

// A Vendor contains the name of the vendor and mappings corresponding to all
// known products by their ID.
type Vendor struct {
	Name    string
	Product map[gousb.ID]*Product
}

// String returns the name of the vendor.
func (v Vendor) String() string {
	_ = "STUB: not implemented"

	// A Product contains the name of the product (from a particular vendor) and
	// the names of any interfaces that were specified.
	return ""
}

type Product struct {
	Name      string
	Interface map[gousb.ID]string
}

// String returns the name of the product.
func (p Product) String() string {
	_ = "STUB: not implemented"

	// A Class contains the name of the class and mappings for each subclass.
	return ""
}

type Class struct {
	Name     string
	SubClass map[gousb.Class]*SubClass
}

// String returns the name of the class.
func (c Class) String() string {
	_ = "STUB: not implemented"

	// A SubClass contains the name of the subclass and any associated protocols.
	return ""
}

type SubClass struct {
	Name     string
	Protocol map[gousb.Protocol]string
}

// String returns the name of the SubClass.
func (s SubClass) String() string {
	_ = "STUB: not implemented"

	// ParseIDs parses and returns mappings from the given reader.  In general, this
	// should not be necessary, as a set of mappings is already embedded in the library.
	// If a new or specialized file is obtained, this can be used to retrieve the mappings,
	// which can be stored in the global Vendors and Classes map.
	return ""
}

func ParseIDs(r io.Reader) (map[gousb.ID]*Vendor, map[gousb.Class]*Class, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// TODO(kevlar): count

// Save the name

// Parse out the level

// Parse the first piece to see if it has a kind

// Parse the ID

// Hold the interim values

// Hold the interim values

// TODO(kevlar): Parse class information, etc
//var class *Class
//var subclass *SubClass
