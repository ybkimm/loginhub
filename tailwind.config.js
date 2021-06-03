module.exports = {
  purge: [
    './internal/tpls/html/*.html'
  ],
  darkMode: 'media',
  theme: {
    extend: {
      lineHeight: (() => {
        let map = {}
        for (let i = 11; i <= 30; i++) {
          map[i] = `${i*0.25}rem`
        }
        return map
      })()
    },
  },
  variants: {},
  plugins: []
}
