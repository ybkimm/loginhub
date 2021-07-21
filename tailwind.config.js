const plugin = require("tailwindcss/plugin")

const groupFocusWithinPlugin = plugin(({ addVariant }) => {
  addVariant('group-focus-within', ({ container }) => {
    container.walkRules((rule) => {
      rule.selector = `.group:focus-within .group-focus-within\\:${rule.selector.slice(1)}`
    })
  })
})

const inputNotEmptyPlugin = plugin(function ({ addVariant }) {
  addVariant('input-not-empty', ({ container }) => {
    container.walkRules((rule) => {
      rule.selector = `input:not(:placeholder-shown) + .input-not-empty\\:${rule.selector.slice(1)}`
    })
  })
})

module.exports = {
  purge: [
    './internal/tpls/html/*.html'
  ],
  darkMode: 'media',
  theme: {
    extend: {
			width: {
				'xs': '20rem',
				'sm': '24rem',
				'md': '28rem',
				'lg': '32rem',
				'xl': '36rem',
				'2xl': '42rem',
				'3xl': '48rem',
				'4xl': '56rem',
				'5xl': '64rem',
				'6xl': '72rem',
				'7xl': '80rem'
			},
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
      margin: ['last'],
      translate: ['group-focus-within', 'input-not-empty'],
      scale: ['group-focus-within', 'input-not-empty'],
      opacity: ['group-focus-within', 'input-not-empty']
    }
  },
  plugins: [
    require('@tailwindcss/forms'),
    groupFocusWithinPlugin,
    inputNotEmptyPlugin
  ]
}
