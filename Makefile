CSS_DIST := internal/assets/public

$(shell mkdir -p ${CSS_DIST})

POSTCSS := npx postcss-cli

tpldemo:
	go run github.com/ybkimm/loginhub/cmd/tpldemo

style: ${CSS_DIST}/style.css

${CSS_DIST}/%.css: styles/%.css
	${POSTCSS} \
		--use postcss-nesting \
		--use postcss-import \
		--use tailwindcss \
		--use autoprefixer --autoprefixer.browsers \
			"cover 99.5% or IE 11 and not IE < 11" \
		--use cssnano  \
		-o $@ $<
