package common

import "strconv"

// go:generate type=DatabaseType

type DatabaseType uint8

const (
	JanusGraph DatabaseType = iota + 1
	Cosmos
	// Neptune
	// DSEGraph
)

func (d DatabaseType) String() string {
	switch d {
	case JanusGraph:
		return "janusgraph"
	case Cosmos:
		return "cosmos"
	}

	return strconv.Itoa(int(d))
}
