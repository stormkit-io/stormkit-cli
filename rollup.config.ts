import { defineConfig } from "rollup";
import { nodeResolve } from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";
import typescript from "@rollup/plugin-typescript";
import jsonResolve from "@rollup/plugin-json";

export default defineConfig({
  input: { index: "src/main.ts" },
  output: {
    dir: "dist",
    format: "esm",
    exports: "auto",
    banner: "#!/usr/bin/env node",
  },
  plugins: [
    typescript(),
    commonjs({
      // Required to use `require` instead of `commonJsRequire` for dynamic imports.
      ignoreDynamicRequires: true,
    }),
    nodeResolve(),
    jsonResolve(),
  ],
});
