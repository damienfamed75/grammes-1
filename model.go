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

package grammes

import (
	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/model"
)

var (
	// NewWebSocketDialer returns websocket with established connection.
	NewWebSocketDialer = gremconnect.NewWebSocketDialer
	// NewVertex returns a vertex struct meant for adding it.
	// NewVertex = model.NewVertex
	// NewProperty returns a property struct meant for adding it to a vertex.
	// NewProperty = model.NewProperty

	// Unmarshal functions.

	// UnmarshalEdgeList is a utility to unmarshal a list
	// or array of edges properly.
	UnmarshalEdgeList = model.UnmarshalEdgeList
	// UnmarshalIDList is a utility to unmarshal a list
	// or array of IDs properly.
	UnmarshalIDList = model.UnmarshalIDList
	// UnmarshalVertexList is a utility to unmarshal a list
	// or array of vertices properly.
	UnmarshalVertexList = model.UnmarshalVertexList
	// UnmarshalPropertyList is a utility to unmarshal a list
	// or array of IDs properly.
	UnmarshalPropertyList = model.UnmarshalPropertyList
)

// Vertex is used to get quick access
// to the model.Vertex without having to
// import it everywhere in the grammes package.
//
// TinkerPop: http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/structure/Vertex.html
//
//  ---inEdges---> vertex ---outEdges--->.
//
// Vertex maintains pointers to both a set
// of incoming and outgoing Edge objects. The
// outgoing edges are the edges for which the
// Vertex is a tail. The incoming edges are those
// edges for which the Vertex is the head.
type Vertex = model.Vertex

// VertexValue is used to get quick access
// to the model.VertexValue without having to
// import it everywhere in the grammes package.
//
// VertexValue contains the 'value' data
// from the Vertex object.
type VertexValue = model.VertexValue

// Edge is used to get quick access
// to the model.Edge without having to
// import it everywhere in the grammes package.
//
// Tinkerpop: http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/structure/Edge.html
//
//  outVertex ---label---> inVertex.
//
// Edge is the object that builds a
// connection between two or more vertices.
type Edge = model.Edge

// Property is used to get quick access
// to the model.Property without having to
// import it everywhere in the grammes package.
//
// Property holds the type and
// value of the property. It's extra
// information used by PropertyDetail.
type Property = model.Property

// PropertyMap is used to get quick access
// to the model.PropertyMap without having to
// import it everywhere in the grammes package.
//
// Tinkerpop: http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/structure/Property.html
//
// PropertyMap is the map used to hold
// the properties itself. the string key is equivalent
// to the Gremlin key and the []Property is the value.
// Properties can have multiple values; this is why we must
// have it as a slice of Property.
type PropertyMap = model.PropertyMap

// ID is used to get quick access
// to the model.ID without having to
// import it everywhere in the grammes package.
//
// ID contains the data stores in the
// 'ID' data including the type and Value
type ID = model.ID
