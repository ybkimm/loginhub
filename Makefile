$(shell mkdir -p internal/assets)

POSTCSS := npx postcss-cli

style: internal/assets/style.css

internal/assets/%.css: styles/%.css
	${POSTCSS} \
		--use postcss-nesting \
		--use postcss-import \
		--use tailwindcss \
		--use autoprefixer --autoprefixer.browsers "> 5%" \
		--use cssnano  \
		-o $@ $<
