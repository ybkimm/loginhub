$(shell mkdir -p internal/assets)

POSTCSS := npx postcss-cli

style: internal/assets/style.css

internal/assets/%.css: styles/%.css
	${POSTCSS} \
		--use postcss-nesting \
		--use postcss-import \
		--use tailwindcss \
		--use autoprefixer --autoprefixer.browsers \
			"cover 99.5% or IE 11 and not IE < 11" \
		--use cssnano  \
		-o $@ $<
