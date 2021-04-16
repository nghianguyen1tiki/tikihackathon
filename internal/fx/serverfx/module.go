package serverfx

import "go.uber.org/fx"

var Provide = fx.Provide(
	provideServer,
)
