package cmd

import (
	"log"
	"os"

	// migration "pab-admin/cmd/database/migration"
	// seeder "pab-admin/cmd/database/seeder"
	doctor "riskmanagement/cmd/doctor"
	http "riskmanagement/cmd/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Eform Service",
		Short: "Eform - Backend Service",
	}
)

func Execute() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Register Command
	rootCmd.AddCommand(http.ServeHTTPCmd())

	http.ServeHTTPCmd().Flags().StringP("config", "c", "config/file", "Config URL dir i.e. config/file")
	http.ServeHTTPCmd().Flags().StringP("env", "e", "", "Config env file")

	rootCmd.AddCommand(doctor.DoctorCmd())

	doctor.DoctorCmd().Flags().StringP("config", "c", "config/file", "Config URL dir i.e. config/file")
	doctor.DoctorCmd().Flags().StringP("env", "e", "", "Config env file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(-1)
	}
}
