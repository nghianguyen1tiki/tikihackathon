package crawlfx

import "go.uber.org/fx"

var Provide = fx.Provide(
	provideCrawler,
)
