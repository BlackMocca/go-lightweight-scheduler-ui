/** @type {import('tailwindcss').Config} */
module.exports = {
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
      kanit: "Kanit-Regular",
      kanitLight: "Kanit-Light",
      kanitBold: "Kanit-Bold",
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
    },
    extend: {
      fontFamily: {
        sans: ['Kanit-Regular', 'Kanit-Light', 'Kanit-Bold']
      },
    }
  },
  plugins: [],
  safelist: [],
}

