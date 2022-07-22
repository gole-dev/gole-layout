package cache

import (
	"github.com/gole-dev/gole/pkg/redis"
	"github.com/google/wire"
)

// ProviderSet is cache providers.
var ProviderSet = wire.NewSet(redis.Init)
