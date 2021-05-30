$(shell mkdir -p internal/styles)

POSTCSS := npx postcss-cli

tpldemo:
	go run github.com/ybkimm/loginhub/cmd/tpldemo

style: internal/styles/style.css

internal/styles/%.css: styles/%.css
	${POSTCSS} \
		--use postcss-nesting \
		--use postcss-import \
		--use tailwindcss \
		--use autoprefixer --autoprefixer.browsers \
			"cover 99.5% or IE 11 and not IE < 11" \
		--use cssnano  \
		-o $@ $<
