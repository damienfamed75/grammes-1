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
	"encoding/json"

	"github.com/northwesternmutual/grammes/internal/model"
)

// List is used in Gremlin v3.0+ instead of arrays in some cases.
type List struct {
	Type  string        `json:"@type"`
	Value []interface{} `json:"@value"`
}

// IDList is used for unmarshalling after querying.
// We use this instead of []ID for Gremlin v3.0 compatibility.
type IDList struct {
	listOfIDs List
	IDs       []ID
}

func (l IDList) Interface() []model.ID {
	var ll []model.ID

	for _, i := range l.IDs {
		ll = append(ll, i)
	}

	return ll
}

// UnmarshalJSON overrides to assure a proper unmarshal.
func (l *IDList) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &l.listOfIDs); err == nil {
		if data, err = json.Marshal(l.listOfIDs.Value); err != nil {
			return err
		}
	}

	return json.Unmarshal(data, &l.IDs)
}

// EdgeList is used for unmarshalling after querying.
// We use this instead of []Edge for Gremlin v3.0 compatibility.
type EdgeList struct {
	listOfEdges List
	Edges       []Edge
}

func (l EdgeList) Interface() []model.Edge {
	var ll []model.Edge

	for _, i := range l.Edges {
		ll = append(ll, &i)
	}

	return ll
}

// UnmarshalJSON overrides to assure a proper unmarshal.
func (l *EdgeList) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &l.listOfEdges); err == nil {
		if data, err = json.Marshal(l.listOfEdges.Value); err != nil {
			return err
		}
	}

	return json.Unmarshal(data, &l.Edges)
}

// VertexList is used for unmarshalling after querying.
// We use this instead of []Vertex for Gremlin v3.0 compatibility.
type VertexList struct {
	listOfVertices List
	Vertices       []Vertex
}

func (l VertexList) Interface() []model.Vertex {
	var ll []model.Vertex

	for _, i := range l.Vertices {
		ll = append(ll, &i)
	}

	return ll
}

// UnmarshalJSON overrides to assure a proper unmarshal.
func (l *VertexList) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &l.listOfVertices); err == nil {
		if data, err = json.Marshal(l.listOfVertices.Value); err != nil {
			return err
		}
	}

	return json.Unmarshal(data, &l.Vertices)
}

// PropertyList is used for unmarshalling after querying.
// We use this instead of []Edge for Gremlin v3.0 compatibility.
type PropertyList struct {
	listOfProperties List
	Properties       []Property
}

func (l PropertyList) Interface() []model.Property {
	var ll []model.Property

	for _, i := range l.Properties {
		ll = append(ll, &i)
	}

	return ll
}

// UnmarshalJSON overrides to assure a proper unmarshal.
func (l *PropertyList) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &l.listOfProperties); err == nil {
		if data, err = json.Marshal(l.listOfProperties.Value); err != nil {
			return err
		}
	}

	return json.Unmarshal(data, &l.Properties)
}
