$ENV:GOARCH = 386
$ENV:CGO_ENABLED = 1
go build -o armago.dll -buildmode=c-shared .