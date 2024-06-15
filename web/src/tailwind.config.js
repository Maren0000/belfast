/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../views/**/*.html", "tw.css"],
  safelist: [
    "badge-neutral",
    "badge-primary",
    "badge-accent",
    "badge-success",
    "text-primary",
    "text-secondary",
  ],
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["autumn", "nord"],
    styled: true,
    base: true,
    utils: true,
    darkTheme: "nord",
  },
};
