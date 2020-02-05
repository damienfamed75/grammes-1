package grammes

import (
	"github.com/northwesternmutual/grammes/internal/common"
)

type DatabaseType = common.DatabaseType

const (
	JanusGraph = common.JanusGraph
	Cosmos     = common.Cosmos
	// Neptune = common.Neptune
	// DSEGraph = common.DSEGraph
)
