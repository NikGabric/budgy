import daisyui from "daisyui";

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
    "./error.vue",
  ],
  theme: {
    extend: {},
  },
  plugins: [daisyui],
  daisyui: {
    themes: [
      {
        budgy: {
          primary: "#4A90E2",
          secondary: "#50E3C2",
          accent: "#F5A623",
          neutral: "#9B9B9B",
          "base-100": "#FFFFFF",
        },
        "budgy-dark": {
          primary: "#3A7BD5",
          secondary: "#4FD1B4",
          accent: "#F5A623",
          neutral: "#A6A6A6",
          "base-100": "#1C1C1C",
        },
      },
    ],
  },
};
