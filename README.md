# myapp
Temporal projects for questioning. 

## Environment
- Go 1.5. 

## go get to build

`go get "github.com/elazarl/go-bindata-assetfs"`
`go get "github.com/unrolled/render"`
`go get "github.com/zenazn/goji"`
`go get "github.com/zenazn/goji/web/middleware"`

## To reproduce the issue

1. `git clone https://github.com/hachi8833/myapp.git` to obtain this.
2. `cd myapp`
3. `go run main.go bindata.go`
4. Open 'localhost:8000' and view the source.
5. Check that 'styles/main.17e92ed5.css' and 'scripts/main.8bd37191.js' on the source HTML are empty when opened. It is weird that HTTP status are all '200'.

## To recompile the bindata

1. `mv renamed_02_dist 02_dist`
2. `go generate`, which actually runs `go-bindata 02_dist/...`
3. `mv 02_dist renamed_02_dist`

Renaming '02_dist' is just to avoid to be reffered instad of bindata.go generated.
