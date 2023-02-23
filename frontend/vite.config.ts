import suidPlugin from "@suid/vite-plugin";
import { defineConfig } from "vite";
import solidPlugin from "vite-plugin-solid";

export default defineConfig({
  plugins: [solidPlugin(), suidPlugin()],
  server: {
    port: 3000,
    proxy: {
      "/api": "http://localhost:4030",
    },
  },
  build: {
    target: "esnext",
    outDir: "../server/frontend",
    emptyOutDir: true,
  },
});
