module github.com/mztlive/searchengine

go 1.19

replace (
	github.com/mztlive/foundation => ../foundation
	github.com/mztlive/logger => ../logger
	github.com/mztlive/repository => ../repository
	github.com/mztlive/utils => ../utils
)

require (
	github.com/meilisearch/meilisearch-go v0.20.1
	github.com/morikuni/failure v1.0.0
	github.com/mztlive/foundation v0.0.0-00010101000000-000000000000
	github.com/mztlive/logger v0.0.0-00010101000000-000000000000
	github.com/mztlive/repository v0.0.0-00010101000000-000000000000
	github.com/mztlive/utils v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.23.0
)

require (
	github.com/Masterminds/squirrel v1.5.3 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/gookit/config/v2 v2.1.6 // indirect
	github.com/gookit/goutil v0.5.12 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.37.1-0.20220607072126-8a320890c08d // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
	golang.org/x/text v0.3.7 // indirect
)
