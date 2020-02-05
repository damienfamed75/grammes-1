// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package janus

import (
	"github.com/northwesternmutual/grammes/internal/model"
)

var (
	_ model.Property = &Property{}
	_ model.ID       = PropertyID{}
)

// Tinkerpop:
// http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/structure/Property.html

// PropertyMap is the map used to hold
// the properties itself. the string key is equivalent
// to the Gremlin key and the []Property is the value.
// Properties can have multiple values; this is why we must
// have it as a slice of Property.
type PropertyMap = model.PropertyMap

// type PropertyMap map[interface{}][]model.Property

// Property holds the type and
// value of the property. It's extra
// information used by PropertyDetail.
type Property struct {
	// Type is the Gremlin-type. For this particular
	// structure it typically will be "g:VertexProperty"
	Type string `json:"@type"`
	// Value stores the actual data of the Property.
	// This would be its Key, Value, and ID.
	Val PropertyValue `json:"@value"`
}

// NewProperty will just shorten the struggle of filling
// a property struct. This is meant to be used when creating a Vertex struct.
func NewProperty(label string, value interface{}) Property {
	return Property{
		Val: PropertyValue{
			Label: label,
			Val: ValueWrapper{
				PropertyDetailedValue: PropertyDetailedValue{
					Val: value,
				},
			},
		},
	}
}

// Value is a shortcut for taking the raw interface{}
// from the property itself without redundancy.
func (p Property) Value() interface{} {
	return p.Val.Val.Val
}

// Label will return the key that is used to find
// this particular property and its value.
func (p Property) Label() string {
	return p.Val.Label
}

// ID returns the ID object.
func (p Property) ID() model.ID {
	return p.Val.ID
}

// PropertyValue contains the ID,
// value, and label of this property's value.
type PropertyValue struct {
	ID    PropertyID   `json:"id"`
	Val   ValueWrapper `json:"value"`
	Label string       `json:"label"`
}

// PropertyID holds the ID that is used
// for the property itself.
type PropertyID struct {
	Type string          `json:"@type"`
	Val  PropertyIDValue `json:"@value"`
}

func (p PropertyID) Value() []interface{} {
	return []interface{}{p.Val.RelationID}
}

func (p PropertyID) ValueInt() int64 {
	return -1
}

// PropertyIDValue holds the value
// of the PropertyID.
type PropertyIDValue struct {
	RelationID string `json:"relationId"`
}
