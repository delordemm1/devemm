import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import laravel from "laravel-vite-plugin";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    laravel({
      input: ["src/main.ts"],
      buildDirectory: "build",
      publicDirectory: "../static",
      // ssr: "src/ssr.ts",
      refresh: true,
    }),
    svelte(),
  ],
  publicDir: 'static',
  build:{
    emptyOutDir: true,
  }
});
