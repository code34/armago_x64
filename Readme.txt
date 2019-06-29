$ENV:GOARCH = 386
$ENV:CGO_ENABLED = 1
gcc -c RVExtension.c -m32
ar cru libRVExtension.a RVExtension.o
go build -o armago.dll -buildmode=c-shared armago.go