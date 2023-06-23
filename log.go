package main

import (
	"fmt"
	"os"
)

var logFlags = struct {
	verbose bool
	quiet   bool
}{}

func print(f string, vs []any) { fmt.Fprintf(os.Stderr, f+"\n", vs...) }

func debug(f string, vs ...any) {
	if logFlags.verbose {
		print("[debug] "+f, vs)
	}
}

func info(f string, vs ...any) {
	if !logFlags.quiet {
		print(f, vs)
	}
}

func hiss(f string, vs ...any) {
	print("error: "+f, vs)
}
