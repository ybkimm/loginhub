CSS_SRC := assets/public
CSS_DST := internal/assets/public

$(shell mkdir -p ${CSS_DST})

POSTCSS := npx postcss-cli

tpldemo:
	go run github.com/ybkimm/loginhub/cmd/tpldemo

style: ${CSS_DST}/style.css

${CSS_DST}/%.css: ${CSS_SRC}/%.css tailwind.config.js
	${POSTCSS} \
		--use postcss-nesting \
		--use postcss-import \
		--use tailwindcss \
		--use autoprefixer --autoprefixer.browsers \
			"cover 99.5% or IE 11 and not IE < 11" \
		--use cssnano  \
		-o $@ $(firstword $<)
