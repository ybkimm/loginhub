CSS_SRC := assets/styles
CSS_DST := build/styles
$(shell mkdir -p ${CSS_DST})

TS_SRC := assets/scripts
TS_DST := build/scripts
$(shell mkdir -p ${TS_DST})

KEY_DST := internal/secrets
KEY_FILES := $(addprefix ${KEY_DST}/,token.key)

SQL_SRC_FILES := $(shell find ./sql -name '*.sql' -type f)
SQL_DST := internal/db
SQL_DST_FILE := ${SQL_DST}/db.go
$(shell mkdir -p ${SQL_DST})

POSTCSS := npx postcss-cli
ESBUILD := npx esbuild --bundle
KEYGEN := ./scripts/keygen.sh
SQLC := sqlc

.PHONY: tpldemo
tpldemo:
	go run github.com/ybkimm/loginhub/cmd/tpldemo

.PHONY: run
run: ${CSS_DST}/style.css ${TS_DST}/main.bundle.js
	go run github.com/ybkimm/loginhub/cmd/loginhub

.PHONY: build
build: ESBUILD += --minify

.PHONY: style
style: ${CSS_DST}/style.css

${CSS_DST}/%.css: ${CSS_SRC}/%.css tailwind.config.js
	${POSTCSS} \
		--no-map \
		--use postcss-nesting \
		--use postcss-import \
		--use tailwindcss \
		--use autoprefixer --autoprefixer.browsers \
			"cover 99.5% or IE 11 and not IE < 11" \
		--use cssnano  \
		-o $@ $(firstword $<)

.PHONY: script
script: ${TS_DST}/main.bundle.js

${TS_DST}/%.bundle.js: ${TS_SRC}/%.ts tsconfig.json $(shell find assets/scripts \( -name '*.ts' -o -name '*.tsx' -o -name '*.js' -o -name '*.jsx' \))
	${ESBUILD} --outfile=$@ $(firstword $<)

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
