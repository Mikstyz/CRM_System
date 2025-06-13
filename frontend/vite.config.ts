import { defineConfig } from "vite";
// @ts-ignore
import react from "@vitejs/plugin-react";
import svgr from "vite-plugin-svgr";
import tsconfigPaths from "vite-tsconfig-paths";
// @ts-ignore
import tailwindcss from "@tailwindcss/vite";
import * as path from "path";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
    tsconfigPaths(),
    svgr({
      svgrOptions: {
        exportType: "named",
        ref: true,
        svgo: false,
        titleProp: true,
      },
      include: "**/*.svg?react",
    }),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@wails": path.resolve(__dirname, "./wailsjs"),
    },
  },
  root: ".", // корень — frontend
  base: "./", // относительные пути в сборке
  build: {
    outDir: "./dist", // собираем в ../build
    emptyOutDir: true,
  },
});
