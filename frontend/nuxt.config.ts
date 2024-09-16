// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",

  app: {
    pageTransition: { name: "page", mode: "out-in", duration: 100 },
  },

  devtools: { enabled: true },
  css: ["~/assets/css/main.css"],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },

  modules: ["@pinia/nuxt"],
  imports: {
    dirs: ["stores"],
  },
});
