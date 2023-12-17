/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "../domain/**.go",
    "../pages/**.go",
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
      'xs': '0.75rem',   
      'sm': '0.875rem', 
      'base': '1rem',    
      'lg': '1.125rem',  
      'xl': '1.25rem',   
      '2xl': '1.5rem',  
    },
    colors: {
      primary: {
        base: "#264653"
      },
      secondary: {
        base: "#f0fdfa"
      }
    },
    extend: {
      fontFamily: {
        sans: ['Kanit-Regular', 'Kanit-Light', 'Kanit-Bold']
      }
    }
  },
  plugins: [],
  safelist: [],
}

