/** @type {import('tailwindcss').Config} */
export default {
  content: ["./views/**/*.templ", "./assets/**/*.js"],
  theme: {
    extend: {},
  },
  plugins: [
        require('daisyui'),
  ],
}

