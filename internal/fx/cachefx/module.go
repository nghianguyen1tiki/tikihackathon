package cachefx

import "go.uber.org/fx"

var Provide = fx.Provide(
	provideCache,
)
