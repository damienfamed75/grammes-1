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

package model

import (
	"encoding/json"
	"fmt"

	"github.com/northwesternmutual/grammes/internal/common"
	"github.com/northwesternmutual/grammes/internal/janus"
	"github.com/northwesternmutual/grammes/internal/model"
)

// UnmarshalVertexList returns a slice of Vertices.
func UnmarshalVertexList(db common.DatabaseType, data []byte) ([]model.Vertex, error) {
	switch db {
	case common.JanusGraph:
		var l janus.VertexList

		if err := json.Unmarshal(data, &l); err != nil {
			return nil, fmt.Errorf("unmarshal error: %w", err)
		}

		return l.Interface(), nil
	case common.Cosmos:
		//TODO
	}

	return nil, fmt.Errorf("no database found: %s", db)
}

// UnmarshalEdgeList returns a slice of Edges.
func UnmarshalEdgeList(db common.DatabaseType, data []byte) ([]model.Edge, error) {
	switch db {
	case common.JanusGraph:
		var l janus.EdgeList

		if err := json.Unmarshal(data, &l); err != nil {
			return nil, fmt.Errorf("unmarshal error: %w", err)
		}

		return l.Interface(), nil
	case common.Cosmos:
		//TODO
	}

	return nil, fmt.Errorf("no database found: %s", db)
}

// UnmarshalPropertyList returns a slice of Edges.
func UnmarshalPropertyList(db common.DatabaseType, data []byte) ([]model.Property, error) {
	switch db {
	case common.JanusGraph:
		var l janus.PropertyList

		if err := json.Unmarshal(data, &l); err != nil {
			return nil, fmt.Errorf("unmarshal error: %w", err)
		}

		return l.Interface(), nil
	case common.Cosmos:
		//TODO
	}

	return nil, fmt.Errorf("no database found: %s", db)
}

// UnmarshalIDList returns a slice of Edges.
func UnmarshalIDList(db common.DatabaseType, data []byte) ([]model.ID, error) {
	switch db {
	case common.JanusGraph:
		var l janus.IDList

		if err := json.Unmarshal(data, &l); err != nil {
			return nil, fmt.Errorf("unmarshal error: %w", err)
		}

		return l.Interface(), nil
	case common.Cosmos:
		//TODO
	}

	return nil, fmt.Errorf("no database found: %s", db)
}
