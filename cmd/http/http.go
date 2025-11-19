package cmd

import (
	"fmt"

	"riskmanagement/bootstrap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:              "serve-http",
		Short:            "Eform Service HTTP",
		Long:             "Service Eform through HTTP",
		PersistentPreRun: httpPreRun,
		RunE:             runHTTP,
	}
)

func ServeHTTPCmd() *cobra.Command {
	return serveHTTPCmd
}

func httpPreRun(cmd *cobra.Command, args []string) {
	// log.LogInit()

	// if env.Get() == env.EnvDevelopment {
	// 	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	// 	log.Level("debug")
	// }
	fmt.Println("root Pre RUN")
}

func runHTTP(cmd *cobra.Command, args []string) error {
	fmt.Println("runHTTP is being called")
	_ = godotenv.Load()
	fx.New(bootstrap.Module).Run()
	return nil
}
