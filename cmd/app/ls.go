package app

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
	Run: runAppLs,
}

func init() {
	appCmd.AddCommand(lsCmd)

	lsCmd.Flags().BoolP("details", "d", false, "Show details of the apps")
}

func runAppLs(cmd *cobra.Command, args []string) {
	apps, err := api.GetApps()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	details, err := cmd.Flags().GetBool("details")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
			fmt.Printf("Repo: %s\n", a.Repo)
			fmt.Printf("  ID: %s\n", a.ID)
			fmt.Printf("  Status: %t\n", a.Status)
			fmt.Printf("  AutoDeploy: %s\n", a.AutoDeploy)
			fmt.Printf("  DefaultEnv: %s\n", a.DefaultEnv)
			fmt.Printf("  Endpoint: %s\n", a.Endpoint)
			fmt.Printf("  DisplayName: %s\n", a.DisplayName)
			fmt.Printf("  CreatedAt: %s\n", time.Unix(int64(a.CreatedAt), 0))
			fmt.Printf("  DeployedAt: %s\n", time.Unix(int64(a.DeployedAt), 0))
			fmt.Println()
		} else {
			fmt.Printf(printf, a.ID, a.Repo)
		}
	}
}
