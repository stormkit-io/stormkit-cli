package app

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/stormkit-io/stormkit-cli/api"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runAppLs,
}

func init() {
	appCmd.AddCommand(lsCmd)

	lsCmd.Flags().BoolP("details", "d", false, "Show details of the apps")
}

func runAppLs(cmd *cobra.Command, args []string) error {
	apps, err := api.GetApps()
	if err != nil {
		return err
	}

	details, err := cmd.Flags().GetBool("details")
	if err != nil {
		return err
	}

	idMaxLength := 0
	for _, a := range apps.Apps {
		if len(a.ID) > idMaxLength {
			idMaxLength = len(a.ID)
		}
	}
	lenf := strconv.Itoa(idMaxLength)
	printf := "%" + lenf + "v  %s\n"
	tabf := fmt.Sprintf("%"+lenf+"v", "")

	if !details {
		fmt.Printf("ID%sRepository\n", tabf)
	}

	for _, a := range apps.Apps {
		if details {
			fmt.Print(api.DumpApp(a))
		} else {
			fmt.Printf(printf, a.ID, a.Repo)
		}
	}

	return nil
}
