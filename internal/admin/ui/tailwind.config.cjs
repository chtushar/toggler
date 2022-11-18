const colors = require('tailwindcss/colors');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      white: colors.white,
      neutral: colors.neutral,
      accent: colors.indigo,
    },
    extend: {
      fontFamily: {
        sans: ["Inter var", "sans-serif"],
      }
    },
  },
  plugins: [],
}
