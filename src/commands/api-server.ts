import path from "node:path";
import fs from "node:fs";
import express from "express";
import { green, blue } from "colorette";
import { createServer as createViteServer } from "vite";
import apiMiddleware from "@stormkit/serverless/middlewares";

interface DevServerConfig {
  // The port to listen
  port?: number;
  // The host to listen
  host?: string;
  // If provided, the directory will be used as a file-system based routing root.
  dir?: string;
}

const defaultConfig: DevServerConfig = {
  dir: process.env.SERVERLESS_DIR || "api",
  host: process.env.SERVERLESS_HOST || "localhost",
  port: Number(process.env.SERVERLESS_PORT) || 3000,
};

const getRootFolder = (apiDir: string = "api") => {
  const cwd = process.cwd();

  if (fs.existsSync(path.join(cwd, apiDir))) {
    return path.join(cwd, apiDir);
  }

  let dir = require?.main?.filename || cwd;

  if (dir.indexOf("node_modules") > -1) {
    return /^(.*?)node_modules/.exec(dir)?.[1] || dir;
  }

  while (dir !== path.sep) {
    if (fs.existsSync(path.join(dir, "package.json"))) {
      return dir;
    }

    dir = path.dirname(dir);
  }

  return path.join(dir, apiDir);
};

class DevServer {
  config: DevServerConfig;

  constructor(config: DevServerConfig) {
    Object.keys(defaultConfig).forEach((k) => {
      const key = k as keyof DevServerConfig;

      if (typeof config[key] === "undefined") {
        // @ts-ignore
        config[key] = defaultConfig[key];
      }
    });

    this.config = config;
  }

  async listen(): Promise<void> {
    const app = express();

    const vite = await createViteServer({
      server: { middlewareMode: true },
      appType: "custom",
      optimizeDeps: {
        disabled: true,
      },
    });

    // use vite's connect instance as middleware
    // if you use your own express router (express.Router()), you should use router.use
    app.use(vite.middlewares);

    app.all(
      "*",
      apiMiddleware({
        middleware: "express",
        apiDir: getRootFolder(this.config.dir),
        moduleLoader: vite.ssrLoadModule,
      })
    );

    app.listen(this.config.port!, this.config.host!, () => {
      console.log(
        `Server running at ${green(
          `http://${this.config.host}:${this.config.port}/`
        )}`
      );

      console.log(`Listening changes on directory: ${blue(this.config.dir!)}`);
    });
  }
}

export default DevServer;
