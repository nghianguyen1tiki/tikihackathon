package feedfx

import "go.uber.org/fx"

var Invoke = fx.Invoke(
	invokeFeedHandler,
)

var Provide = fx.Provide(
	provideRepo,
)
