package tikingonfx

import "go.uber.org/fx"

var Invoke = fx.Invoke(
	invokeTikiNgonHandler,
)
