/** @type {import('tailwindcss').Config} */
export default {
content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      animation: {
        'appear': 'appear 0.25s'
      },
      keyframes: {
        appear: {
          '100%': {opacity: '1'}
        }
      }
    },
  },
  plugins: [],
}

