CSS_SRC := assets/styles
CSS_DST := build/styles
$(shell mkdir -p ${CSS_DST})

JS_SRC := assets/scripts
JS_DST := build/scripts
$(shell mkdir -p ${JS_DST})

KEY_DST := internal/secrets
KEY_FILES := $(addprefix ${KEY_DST}/,token.key)

SQL_SRC_FILES := $(shell find ./sql -name '*.sql' -type f)
SQL_DST := internal/db
SQL_DST_FILE := ${SQL_DST}/db.go
$(shell mkdir -p ${SQL_DST})

POSTCSS := npx postcss-cli
DENO_BUNDLE := deno bundle -c ./tsconfig.json
KEYGEN := ./scripts/keygen.sh
SQLC := sqlc

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

.PHONY: script
script: ${JS_DST}/main.bundle.js

${JS_DST}/%.bundle.js: ${JS_SRC}/%.ts tsconfig.json
	${DENO_BUNDLE} $(firstword $<) $@

.PHONY: secrets
secrets: ${KEY_FILES}

${KEY_DST}/%.key:
	-${KEYGEN} > $@

.PHONY: sql
sql: ${SQL_DST_FILE}

${SQL_DST_FILE}: ${SQL_SRC_FILES}
	${SQLC} generate

.PHONY: clean
clean:
	rm ${CSS_DST}/*.css \
	   ${JS_DST}/*.js
