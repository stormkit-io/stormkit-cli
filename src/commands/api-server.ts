import type { RequestEvent, ServerlessResponse } from "@stormkit/serverless";
import http from "node:http";
import path from "node:path";
import fs from "node:fs";
import express from "express";
import { green, blue } from "colorette";
import sk from "@stormkit/serverless";

const handler = (event: any, root: string): Promise<ServerlessResponse> => {
  var tsNode = require("ts-node");

  tsNode.createEsmHooks(
    tsNode.register({
      // We need to ovewrite the `"type": "module"` specified in
      // package.json.
      moduleTypes: {
        [`${root}/**/*`]: "cjs",
      },
      transpileOnly: true,
    })
  );

  // @ts-ignore
  return sk(undefined, "stormkit:api")(event, root);
};

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

const rootDir = ((): string => {
  let dir = require?.main?.filename || process.cwd();

  if (dir.indexOf("node_modules") > -1) {
    return /^(.*?)node_modules/.exec(dir)?.[1] || dir;
  }

  while (dir !== "/") {
    if (fs.existsSync(path.join(dir, "package.json"))) {
      return dir;
    }

    dir = path.dirname(dir);
  }

  return dir;
})();

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

  async _readBody(req: http.IncomingMessage): Promise<string> {
    const body: string[] = [];

    return new Promise((resolve, _) => {
      if (req.method?.toLowerCase() === "get") {
        return resolve("");
      }

      req.on("data", (chunks) => {
        body.push(chunks.toString("utf-8"));
      });

      req.on("end", () => {
        resolve(body.join(""));
      });
    });
  }

  async _transformToRequestEvent(
    req: http.IncomingMessage
  ): Promise<RequestEvent> {
    const headers: Record<string, string> = {};
    const body = await this._readBody(req);

    Object.keys(req.headers).forEach((key) => {
      const headerVal = req.headers[key];
      const headerKey = key.toLowerCase();

      if (Array.isArray(headerVal)) {
        headers[headerKey] = headerVal.join(",");
      } else if (headerVal) {
        headers[headerKey] = headerVal;
      }
    });

    const request: RequestEvent = {
      method: req.method || "get",
      url: req.url || "/",
      path: req.url?.split("?")?.[0] || "/",
      body,
      headers,
    };

    return request;
  }

  listen(): void {
    const app = express();

    app.all("*", async (req, res) => {
      const request = await this._transformToRequestEvent(req);
      const root = path.join(rootDir, this.config.dir || "");

      try {
        const data = await handler(request, root.replace(/\\/g, "/"));

        Object.keys(data.headers || {}).forEach((key) => {
          res.set(key, data.headers[key]);
        });

        res.status(data.status);
        res.send(Buffer.from(data.buffer || "", "base64").toString("utf-8"));
      } catch (err) {
        console.log("execute ts-node error err:", err);
      }
    });

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
