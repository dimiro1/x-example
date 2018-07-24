package app

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			// Routes
			Index,
			About,
			// Health Checks
			GoogleHealthCheck,
			HelloFreshHealthCheck,
			CustomHealthCheck,

			// Template filters
			Uppercase,
			Slugify,
		),
	)
}
