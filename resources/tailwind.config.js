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
          800: '#121212',
        },
      },
    },
  },
  plugins: [],
}

