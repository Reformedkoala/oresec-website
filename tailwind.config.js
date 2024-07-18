/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./static/templates/*.html'], // This is where your HTML templates / JSX files are located
  safelist: [
    'text-2xl',
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Iosevka Aile Iaso", "sans-serif"],
        mono: ["Iosevka Curly Iaso", "monospace"],
        serif: ["Iosevka Etoile Iaso", "serif"],
      },
    },
  },
  plugins: [],
};

