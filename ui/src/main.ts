
import '@skeletonlabs/skeleton/styles/skeleton.css';
import "./app.postcss";
import './styles/theme.postcss';
// @ts-nocheck
import { createInertiaApp } from "@inertiajs/svelte";
// import "./main.css";
import "./styles/main.scss";
createInertiaApp({
  resolve: (name: string) => {
    const pages = import.meta.glob("./pages/**/*.svelte", { eager: true });
    console.log(pages[`./pages/${name}.svelte`]);
    return pages[`./pages/${name}.svelte`] || import("./pages/error.svelte");
  },
  setup({ el, App, props }) {
    new App({ target: el, props });
  },
  progress: {
    color: "#d00",
  },
});
