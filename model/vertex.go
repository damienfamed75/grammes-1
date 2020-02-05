package model

import "github.com/northwesternmutual/grammes/internal/model"

func IsVertexNil(v Vertex) bool {
	if len(v.ID().Value()) > 0 {
		return false
	}

	return true
}

type (
	Vertex      = model.Vertex
	VertexValue = model.VertexValue
	PropertyMap = model.PropertyMap
	Property    = model.Property
	ID          = model.ID
)
