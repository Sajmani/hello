//go:guide This file, hello.go, contains the Go source code for a
// command line tool that prints "hello, world", with flags to customize
// the message. The comment below, directly above the "package main"
// declaration, is a "doc comment" that will be rendered in the generated
// documentation for this package. You can read more about doc comments
// [here](https://go.dev/doc/comment).

// Hello prints "hello, world". The -greeting flag customizes the "hello" part
// of the message, and the -whom flag customizes the "world" part.
package main

//go:guide The import block makes other Go packages available to the
// source code in this file. The hello command imports the "flag" package
// to set up command line flags, and it imports "fmt" to format and print
// strings. Most editors can be set up to populate the "import" block
// automatically from the available packages.
import (
	"flag"
	"fmt"
)

//go:guide This var block registers the -greeting and -whom command line flags.
var (
	greeting = flag.String("greeting", "hello", `the "hello" part of "hello, world"`)
	whom     = flag.String("whom", "world", `the "world" part of "hello, world"`)
)

//go:guide Go programs start by calling the function named "main" in package main.
// This main function parses the command line flags then prints the "hello, world"
// message, followed by a newline. The expression *greeting evaluates to the string
// value of the -greeting command line flag, and likewise for *whom.

func main() {
	flag.Parse()
	fmt.Printf("%s, %s\n", *greeting, *whom)
}
