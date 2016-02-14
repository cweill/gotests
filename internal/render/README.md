bindata.go must be generated with https://github.com/jteeuwen/go-bindata.

From the gotests root run `go generate ./...`.

Or from this directory run `go-bindata -pkg=bindata -o "./bindata/bindata.go" templates`.

During development run `go-bindata -pkg=bindata -o "./bindata/bindata.go" -debug templates` instead.