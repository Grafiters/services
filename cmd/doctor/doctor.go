package cmd

import (
	"riskmanagement/lib"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	doctorCmd = &cobra.Command{
		Use:              "doctor",
		Short:            "doctor check dependencies",
		Long:             "doctor tool for check dependencies",
		PersistentPreRun: doctorPreRun,
		RunE:             runDoctor,
	}
)

func DoctorCmd() *cobra.Command {
	return doctorCmd
}

func doctorPreRun(cmd *cobra.Command, args []string) {
	// fmt.Println("doctor Pre RUN")
}

func runDoctor(cmd *cobra.Command, args []string) error {
	lib.LogInfo("----------------------------------------")
	lib.LogInfo("----------- Go Doctor ------------------")
	lib.LogInfo("----------------------------------------")
	_ = godotenv.Load()

	lib.LogInfo("----------------------------------------")
	return nil
}
