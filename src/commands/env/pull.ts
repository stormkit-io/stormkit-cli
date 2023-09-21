import path from "node:path";
import fs from "node:fs";
import api from "../../api";

export default async (options: { debug?: boolean }) => {
  if (!process.env.SK_API_KEY) {
    throw new Error(
      "No `SK_API_KEY` environment variable created. Create one from your environment page and make it available for this command."
    );
  }

  try {
    const skEnvVars = await api.get<Record<string, string>>("/env/pull");
    const dotEnv = path.resolve(process.cwd(), ".env");
    const envVars: string[] = [];

    if (fs.existsSync(dotEnv)) {
      console.log(".env found");
      console.log(".env is going to patched with variables from Stormkit");

      const existingVars = fs
        .readFileSync(dotEnv, "utf-8")
        .split("\n")
        .reduce((obj, line) => {
          const indexOfEqual = line.indexOf("=");

          obj[line.slice(0, indexOfEqual).trim()] = line
            .slice(indexOfEqual + 1)
            .trim();

          return obj;
        }, {} as Record<string, string>);

      const mergedVars = { ...existingVars, ...skEnvVars };

      Object.keys(mergedVars).forEach((key) => {
        envVars.push(`${key}=${mergedVars[key]}`);
      });
    } else {
      console.log(".env not found");

      Object.keys(skEnvVars).forEach((key) => {
        envVars.push(`${key}=${skEnvVars}`);
      });
    }

    fs.writeFileSync(dotEnv, envVars.join("\n"));

    console.log(`.env is now ready to be consumed\n`);

    if (options?.debug) {
      console.log(`file location: ${dotEnv}`);
    }
  } catch {
    console.error(
      "Cannot pull environment variables. Make sure your API key is correct or that the environment exists."
    );

    process.exit(1);
  }
};
