package repository

import (
	"github.com/gole-dev/gole-layout/internal/model"
	"github.com/google/wire"
)

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(model.Init)
