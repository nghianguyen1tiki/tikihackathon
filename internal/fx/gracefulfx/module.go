package gracefulfx

import "go.uber.org/fx"

var Invoke = fx.Invoke(
	invokeGraceful,
)
