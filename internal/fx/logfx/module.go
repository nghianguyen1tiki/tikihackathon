package logfx

import "go.uber.org/fx"

var Invoke = fx.Invoke(
	invokeLogger,
)
