package main

import (
	"github.com/stormkit-io/stormkit-cli/cmd"
	_ "github.com/stormkit-io/stormkit-cli/cmd/app"
	_ "github.com/stormkit-io/stormkit-cli/cmd/deploy"
	_ "github.com/stormkit-io/stormkit-cli/cmd/log"
)

func main() {
	cmd.Execute()
}
