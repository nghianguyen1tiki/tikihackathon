package recipefx

import "go.uber.org/fx"

var Provide = fx.Provide(
	provideRepo,
)

var Invoke = fx.Invoke(
	invokeRecipeHandler,
)
