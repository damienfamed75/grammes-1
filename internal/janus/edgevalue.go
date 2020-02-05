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
	"errors"

	"github.com/northwesternmutual/grammes/internal/common"
	"github.com/northwesternmutual/grammes/internal/model"
	"github.com/northwesternmutual/grammes/query/traversal"
)

// EdgeValue contains the 'value' data
// from the Edge object.
type EdgeValue struct {
	ValueID    EdgeID         `json:"id"`
	ValueLabel string         `json:"label"`
	InVLabel   string         `json:"inVLabel,omitempty"`
	OutVLabel  string         `json:"outVLabel,omitempty"`
	InV        ID             `json:"inV,omitempty"`
	OutV       ID             `json:"outV,omitempty"`
	Props      EdgeProperties `json:"properties,omitempty"`
}

// Property will retrieve the Property for you.
func (e *EdgeValue) Property(key interface{}) interface{} {
	return e.Props[key].Value.Value.PropertyDetailedValue.Val
}

func (e *EdgeValue) Type() string {
	return e.Type()
}

// ID will retrieve the Edge ID for you.
func (e *EdgeValue) ID() model.ID {
	return e.ValueID
}

func (e *EdgeValue) Label() string {
	return e.ValueLabel
}

// OutVertexID will retrieve the id for the
// vertex that the edge goes out of.
func (e *EdgeValue) OutVertexID() model.ID {
	return e.OutV
}

// InVertexID will retrieve the id for the
// vertex that the edge goes into.
func (e *EdgeValue) InVertexID() model.ID {
	return e.InV
}

// OutVertexLabel will retrieve the label
// for the vertex the edge goes out of.
func (e *EdgeValue) OutVertexLabel() string {
	return e.OutVLabel
}

// InVertexLabel will retrieve the label
// for the vertex the edge goes into.
func (e *EdgeValue) InVertexLabel() string {
	return e.InVLabel
}

// QueryOutVertex will retrieve the vertex that
// the edge comes out of.
func (e *EdgeValue) QueryOutVertex(client common.QueryClient) (model.Vertex, error) {
	if client == nil {
		return nil, errors.New("QueryOutVertex: nil client given to Edge")
	}

	responses, err := client.ExecuteQuery(traversal.NewTraversal().
		V().HasID(e.OutVertexID()))
	if err != nil {
		return nil, err
	}

	vertList, err := UnmarshalVertexList(responses)
	if err != nil {
		return nil, err
	}

	return &vertList[0], nil
}

// QueryInVertex will retrieve the vertex that
// the edge comes out of.
func (e *EdgeValue) QueryInVertex(client common.QueryClient) (model.Vertex, error) {
	if client == nil {
		return nil, errors.New("QueryInVertex: nil client given to Edge")
	}

	responses, err := client.ExecuteQuery(traversal.NewTraversal().
		V().HasID(e.InVertexID()))
	if err != nil {
		return nil, nil
	}

	vertList, err := UnmarshalVertexList(responses)
	if err != nil {
		return nil, err
	}

	return &vertList[0], nil
}

// EdgeVertex only contains the type
// and ID of the vertex.
// type EdgeVertex struct {
// 	Type  string `json:"@type"`
// 	Value int64  `json:"@value"`
// }
