/** @type {import('tailwindcss').Config} */
var defaultColorConfig = require('./tailwind.config-color.js')

module.exports = {
  mode: 'aot',
  content: [
    "../domain/**.go",
    "../domain/**/*.go",
    "../pages/**.go",
    "../pages/**/*.go",
  ],
  theme: {
    screens: {
      iphone: { 
        min: "280px", max: "639px" 
      },
      tablet: "640px",
      laptop: "1280px",
    },
    fontFamily: {
      kanit: "Kanit-Light",
      kanitBold: "Kanit-Regular",
    },
    fontSize: {
      sm: '0.8rem',
      base: '1rem',
      xl: '1.25rem',
      '2xl': '1.563rem',
      '3xl': '1.953rem',
      '4xl': '2.441rem',
      '5xl': '3.052rem',
    },
    colors: {
      primary: {
        base: "#264653"
      },
      secondary: {
        base: "#FFFAFA"
      },
      red: "red",
    },
    extend: {
      fontFamily: {
        sans: ['Kanit-Light', 'Kanit-Regular']
      },
      borderWidth: {
        '0.5': '0.5px',
      },
      colors:{
        ...defaultColorConfig.theme.colors
      },
    }
  },
  plugins: [],
  safelist: [],
}

