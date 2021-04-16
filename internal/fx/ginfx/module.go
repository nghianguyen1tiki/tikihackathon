package ginfx

import "go.uber.org/fx"

var Provide = fx.Provide(
	provideGin,
)
