# armago

Arma Golang Extension

The minimal requirement to build a good & nice GOLANG .dll or .so extension with ARMA3 :)

## Requirements & Build

1- install the 64bits TDD gcc compiler on your machine
https://sourceforge.net/projects/tdm-gcc/

2- use armago_x64 as your entrie point template for your extension :)

3- build your extension with this command line :
go build -o armago_x64.dll -buildmode=c-shared armago_x64.go