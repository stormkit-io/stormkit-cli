#!/usr/bin/env node
import yargs from "yargs";
import * as cmds from "./index";
import dotenv from "dotenv";

dotenv.config();

yargs()
  .scriptName("stormkit")
  .usage("$0 <cmd> [args]")
  .command({
    command: "api",
    describe: "Starts an API development server",
    builder: {
      port: {
        alias: "p",
        describe: "Specify the port on which the API server should listen.",
        default: "9090",
      },
      dir: {
        alias: "d",
        default: "api",
        description:
          "Specify the directory in which the API routes are located. This path is relative to project root.",
      },
    },
    handler(argv) {
      new cmds.apiServer({
        port: argv.port ? Number(argv.port) : undefined,
        dir: typeof argv.dir === "string" ? argv.dir : undefined,
      }).listen();
    },
  })
  .command({
    command: "env pull",
    describe: "Pull environment variables from the environment.",
    builder: {
      debug: {
        alias: "d",
        describe: "Prints additional logs.",
      },
    },
    handler(argv) {
      cmds.envPull({ debug: argv.debug as boolean });
    },
  })
  .parse(process.argv.slice(2));
