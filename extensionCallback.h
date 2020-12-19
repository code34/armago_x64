#pragma once

#include <stdlib.h>

typedef int (*extensionCallback)(char const *name, char const *function, char const *data);

/*
** * To avoid the multiple definations error in go files "static inline" is used
** * https://golang.org/cmd/cgo/#hdr-C_references_to_Go
*/
static inline int runExtensionCallback(extensionCallback fnc,
	char const *name, char const *function, char const *data)
{
	return fnc(name, function, data);
}
