package crawlfx

import "go.uber.org/fx"

var Invoke = fx.Invoke(
	invokeCrawler,
)
