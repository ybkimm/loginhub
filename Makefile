CSS_SRC := assets/styles
CSS_DST := build/styles

$(shell mkdir -p ${CSS_DST})

POSTCSS := npx postcss-cli

.PHONY: tpldemo
tpldemo:
	go run github.com/ybkimm/loginhub/cmd/tpldemo

.PHONY: style
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

.PHONY: clean
clean:
	rm ${CSS_DST}/*.css
