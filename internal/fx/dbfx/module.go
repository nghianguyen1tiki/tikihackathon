package dbfx

import (
	"go.uber.org/fx"
)

var Provide = fx.Provide(
	provideDB,
)
