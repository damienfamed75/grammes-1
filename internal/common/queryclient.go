package common

import (
	"github.com/northwesternmutual/grammes/query"
)

type QueryClient interface {
	ExecuteQuery(query.Query) ([][]byte, error)
	ExecuteStringQuery(string) ([][]byte, error)
}
