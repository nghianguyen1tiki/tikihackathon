package httpfx

import "go.uber.org/fx"

var Provide = fx.Provide(
	provideHTTPClient,
)
