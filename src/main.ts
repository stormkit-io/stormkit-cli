import yargs from "yargs";
import * as cmds from "./index";

const argv = yargs
  .scriptName("stormkit")
  .usage("$0 <cmd> [args]")
  .command("api ", "Starts an API development server", (yargs) => {
    return yargs
      .option("p", {
        alias: "port",
        default: "9090",
        description: "Specify the port on which the API server should listen.",
      })
      .option("dir", {
        alias: "d",
        default: "api",
        description:
          "Specify the directory in which the API routes are located. This path is relative to project root.",
      });
  })
  .help()
  .parseSync();

if (argv.api) {
  new cmds.apiServer({
    port: argv.port ? Number(argv.port) : undefined,
    dir: argv.dir,
  }).listen();
}
