module.exports = {
  purge: [
    './internal/tpls/html/*.html'
  ],
  darkMode: 'media',
  theme: {
    extend: {
      lineHeight: (() => {
        /**
         * @type {Record<number, string>}
         */
        const map = {}
        for (let i = 11; i <= 30; i++) {
          map[i] = `${i * 0.25}rem`
        }
        return map
      })(),
      colors: {
        primary: {
          light: '#b2f1ad',
          DEFAULT: '#71e46d',
          dark: '#1eb83f'
        },
        secondary: {
          light: '#f2c3f3',
          DEFAULT: '#e06de4',
          dark: '#a716c6'
        }
      },
      spacing: {
        login: '2.33rem'
      },
      fontSize: {
        login: ['1rem', '2.33rem']
      }
    }
  },
  variants: {
    extend: {
      margin: ['last']
    }
  },
  plugins: [
    require('@tailwindcss/forms')
  ]
}
