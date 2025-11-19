package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	installHTTPCmd = &cobra.Command{
		Use:              "install",
		Short:            "Installation Program",
		Long:             "Installation Program",
		PersistentPreRun: installPreRun,
		RunE:             runInstall,
	}
)

func InstallHTTPCmd() *cobra.Command {
	return installHTTPCmd
}

func installPreRun(cmd *cobra.Command, args []string) {
	// create systemd service
	// create nginx service
	fmt.Println("root Pre RUN")
}

func runInstall(cmd *cobra.Command, args []string) error {

	return nil
}
