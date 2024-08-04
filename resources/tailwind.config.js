import defaultTheme from 'tailwindcss/defaultTheme'

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './ts/**/*.{ts,vue}',
    './views/index.html',
    './index.html',
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['Noto Sans', ...defaultTheme.fontFamily.sans],
      },
      colors: {
        obscure: {
          500: '#353535',
          700: '#171717',
          800: '#121212',
        },
      },
    },
  },
  plugins: [],
}

