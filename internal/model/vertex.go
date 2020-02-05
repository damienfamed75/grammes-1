package model

import (
	"github.com/northwesternmutual/grammes/internal/common"
	"github.com/northwesternmutual/grammes/query/traversal"
)

type Vertex interface {
	PropertyValue(key interface{}, index int) interface{}
	PropertyMap() PropertyMap
	Value() VertexValue
	ID() ID
	Label() string
	Traversal() traversal.String

	// Functions that require a queryClient to perform.
	QueryRefresh(c common.QueryClient) error
	QueryBothEdges(c common.QueryClient, labels ...string) ([]Edge, error)
	QueryOutEdges(c common.QueryClient, labels ...string) ([]Edge, error)
	QueryInEdges(c common.QueryClient, labels ...string) ([]Edge, error)
	QueryAddEdge(c common.QueryClient, label string, outVID interface{}, properties ...interface{}) (Edge, error)
	QueryDrop(c common.QueryClient) error
	QueryDropProperties(c common.QueryClient, properties ...interface{}) error
	QueryAddProperty(c common.QueryClient, key, value interface{}) error
}

type VertexValue interface {
	ID() ID
	Label() string
	Properties() PropertyMap
}

// type PropertyMap interface {
// 	Get(key interface{}) []Property
// 	Len() int
// }
type PropertyMap map[interface{}][]Property

type Property interface {
	Label() string
	Value() interface{}
	// ID() PropertyID
	ID() ID // consolidate to the single ID interface
}

// Idea: ID is just a []interface{} and returns that for V() and can be parsed
// down to int64 if you need them.
// This should solve single value int64 and multi-value IDs
type ID interface {
	Value() []interface{}
	ValueInt() int64
}
