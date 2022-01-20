package server

import (
	"github.com/google/wire"
)

// ProviderSet1 is server providers.
var ProviderSet1 = wire.NewSet(NewHTTPServer, NewGRPCServer)
