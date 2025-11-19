package lib

import (
	env "riskmanagement/lib/env"
	jobs "riskmanagement/lib/jobs"

	minioEnv "gitlab.com/golang-package-library/env"
	logger "gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"
	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(env.NewEnv),
	fx.Provide(logger.NewLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewDatabases),
	fx.Provide(minioEnv.NewEnv),
	fx.Provide(minio.NewMinio),
	fx.Provide(jobs.NewCronJob),
)
