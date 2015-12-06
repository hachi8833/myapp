# myapp
Temporal projects for questioning. 

## Environment
- Go 1.5. 

## go get to build

- `go get "github.com/elazarl/go-bindata-assetfs"`
- `go get "github.com/unrolled/render"`
- `go get "github.com/zenazn/goji"`
- `go get "github.com/zenazn/goji/web/middleware"`

## To reproduce the issue

1. `git clone https://github.com/hachi8833/myapp.git` to obtain this.
2. `cd myapp`
3. `go run main.go bindata.go`
4. Open 'localhost:8000' and view the source.
5. Check that 'styles/main.17e92ed5.css' and 'scripts/main.8bd37191.js' on the source HTML are empty when opened. It is weird that HTTP status are all '200'.

>    2015/12/06 21:35:14.003325 [MacBook-Pro.local/vr0NUvmPPI-000001] Started GET "/" from [::1]:55755
>   2015/12/06 21:35:14.004796 [MacBook-Pro.local/vr0NUvmPPI-000001] Returning 200 in 1.418064ms
>   2015/12/06 21:35:14.166211 [MacBook-Pro.local/vr0NUvmPPI-000002] Started GET "/styles/main.17e92ed5.css" from [::1]:55755
>   2015/12/06 21:35:14.166247 [MacBook-Pro.local/vr0NUvmPPI-000002] Returning 200 in 9.849µs
>   2015/12/06 21:35:14.166345 [MacBook-Pro.local/vr0NUvmPPI-000003] Started GET "/scripts/main.8bd37191.js" from [::1]:55756
>   2015/12/06 21:35:14.166370 [MacBook-Pro.local/vr0NUvmPPI-000003] Returning 200 in 5.058µs
>   2015/12/06 21:35:19.865308 [MacBook-Pro.local/vr0NUvmPPI-000004] Started GET "/styles/main.17e92ed5.css" from [::1]:55756
>   2015/12/06 21:35:19.865339 [MacBook-Pro.local/vr0NUvmPPI-000004] Returning 200 in 7.313µs
>   2015/12/06 21:35:22.319075 [MacBook-Pro.local/vr0NUvmPPI-000005] Started GET "/scripts/main.8bd37191.js" from [::1]:55756
>   2015/12/06 21:35:22.319103 [MacBook-Pro.local/vr0NUvmPPI-000005] Returning 200 in 6.581µs

## To recompile the bindata

1. `mv renamed_02_dist 02_dist`
2. `go generate`, which actually runs `go-bindata 02_dist/...`
3. `mv 02_dist renamed_02_dist`

Renaming '02_dist' is just to avoid to be reffered instad of bindata.go generated.
