import { defineConfig } from "rollup";
import { nodeResolve } from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";
import typescript from "@rollup/plugin-typescript";
import jsonResolve from "@rollup/plugin-json";

export default defineConfig({
  input: { index: "src/main.ts" },
  output: {
    dir: "dist",
    format: "cjs",
    exports: "auto",
    banner: "#!/usr/bin/env node",
  },

  plugins: [typescript(), commonjs(), nodeResolve(), jsonResolve()],
});
