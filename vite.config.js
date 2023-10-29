import { resolve } from "path";
import { defineConfig } from "vite";

export default defineConfig({
  build: {
    lib: {
      entry: resolve(__dirname, "src/htmx.js"),
      name: "[name]",
      formats: ["es"],
      fileName: "[name]",
    },
    outDir: "static",
    emptyOutDir: false,
  },
});
