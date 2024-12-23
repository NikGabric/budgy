import daisyui from 'daisyui';

/** @type {import('tailwindcss').Config} */
export default {
  content: ['./intext.html', './src/**/*.{js,ts,jsx,tsx,vue}'],
  plugins: [daisyui],
  daisyui: {
    themes: ['halloween', 'cupcake'],
  },
};
