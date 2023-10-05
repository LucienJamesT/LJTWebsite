/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './templates/**/*.html',  // Looks for .html files in the templates directory and its subdirectories
  ],
  theme: {
    extend: {
      backgroundColor: {
        'brey' : '#182e3d',
      }
    },
  },
  plugins: [],
}
