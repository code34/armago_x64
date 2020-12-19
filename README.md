# Armago

Arma Golang Extension 32/64 bits

The minimal requirement to build a good & nice GOLANG .dll or .so extension with ARMA3 :)

This template extension complete those jobs:
- receive ARMA input
- by default, return an asynchrone result through "ExtensionCallback" of arma handler
- or return a synchrone result throught callExtension arma function (if extensionCallbackFnc is unset)

## Requirements & Build

1- install the 32/64bits TDD gcc compiler on your machine

https://sourceforge.net/projects/tdm-gcc/

2- use armago as your entrie point template to develop your own extension :)

3- build your extension with this command line :

32 bits version
```
$ENV:GOARCH = 386
$ENV:CGO_ENABLED = 1
go build -o armago.dll -buildmode=c-shared .
```

64bits version
```
$ENV:GOARCH = "amd64"
$ENV:CGO_ENABLED = 1
go build -o armago_x64.dll -buildmode=c-shared .
```