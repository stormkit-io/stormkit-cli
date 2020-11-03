package app

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/stormkit-io/stormkit-cli/api"
	"github.com/stormkit-io/stormkit-cli/stormkit"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Chose an application via app id",
	Long: `Choose an application by providing the application id. The chosen application will be cached and will be used as the default application for subsequent commands.

Example: 

$ stormkit-cli app use <app_id>

Please note that this command will have no effect in case it is executed from a directory that contains a "stormkit.config.yml" file. The config file takes precedence over the "app use" command.`,
	RunE: runAppUse,
}

func init() {
	appCmd.AddCommand(useCmd)
}

func runAppUse(cmd *cobra.Command, args []string) error {
	// Fetch Apps from  API
	apps, err := api.GetApps()
	if err != nil {
		return err
	}

	// check if arguments are present
	if len(args) == 0 {
		return errors.New("not enought arguments")
	}

	// check if app exists
	for _, a := range apps.Apps {
		if a.ID == args[0] {
			stormkit.SetEngineAppID(a.ID)
			return nil
		}
	}

	return errors.New("app not found")
}
