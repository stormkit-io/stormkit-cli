package log

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stormkit-io/stormkit-cli/api"
	"github.com/stormkit-io/stormkit-cli/cmd"
	"github.com/stormkit-io/stormkit-cli/stormkit"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log <deploy_id>",
	Short: "Print deployment logs",
	Long:  `Print the logs of the deployment through his ID and the ID of the application, the id of the application from file ~/.stomkit-cli.yml or ./stormkit.config.yml`,
	RunE:  runLog,
	Args:  cobra.ExactArgs(1),
}

func init() {
	cmd.GetRootCmd().AddCommand(logCmd)
}

func runLog(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("not enought arguments")
	}

	appID := stormkit.GetEngineAppID()
	id := args[0]

	deploy, err := api.DeployByID(appID, id)
	if err != nil {
		return err
	}

	for _, l := range deploy.Deploy.Logs {
		fmt.Println(l.Dump())
	}

	return nil
}
