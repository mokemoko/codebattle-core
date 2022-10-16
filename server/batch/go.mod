module batch

go 1.18

require github.com/volatiletech/sqlboiler/v4 v4.13.0

require (
	github.com/friendsofgo/errors v0.9.2 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/null/v8 v8.1.2 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	github.com/volatiletech/strmangle v0.0.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

require (
	github.com/google/uuid v1.3.0
	github.com/mattn/go-sqlite3 v1.14.10
	github.com/mokemoko/codebattle-core/server/models v0.1.0
	gonum.org/v1/gonum v0.12.0
)

replace github.com/mokemoko/codebattle-core/server/models => ../models
