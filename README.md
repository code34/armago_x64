# armago

Arma Golang Extension

The minimal requirement to build a good & nice GOLANG .dll or .so extension with ARMA3 :)

## Requirements & Build

First at all, install the 64bits TDD gcc compiler on your machine
https://sourceforge.net/projects/tdm-gcc/

Second, use armago as your entrie point template for your extension :)

Third, build your extension with this command line :
go build -o armago_x64.dll -buildmode=c-shared armago_x64.go