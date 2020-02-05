package model

import "github.com/northwesternmutual/grammes/internal/common"

// Edge is the object that builds a
// connection between two or more vertices.
//
// Tinkerpop: http://tinkerpop.apache.org/javadocs/3.2.1/core/org/apache/tinkerpop/gremlin/structure/Edge.html
//
//  outVertex ---label---> inVertex.
type Edge interface {
	Value() EdgeValue

	EdgeValue
}

type EdgeValue interface {
	ID() ID
	Type() string
	Label() string
	Property(key interface{}) interface{}
	// Out Vertex
	OutVertexID() ID
	OutVertexLabel() string
	// In Vertex
	InVertexID() ID
	InVertexLabel() string

	QueryOutVertex(common.QueryClient) (Vertex, error)
	QueryInVertex(common.QueryClient) (Vertex, error)
}
