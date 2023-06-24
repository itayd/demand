package main

import (
	"log"
	"os"
)

var logFlags = struct {
	verbose bool
	quiet   bool
}{}

func init() {
	log.SetOutput(os.Stderr)
	log.SetFlags(0)
}

func print(f string, vs ...any) { log.Printf(f+"\n", vs...) }

func debug(f string, vs ...any) {
	if logFlags.verbose {
		print("[debug] "+f, vs)
	}
}

func info(f string, vs ...any) {
	if !logFlags.quiet {
		print(f, vs...)
	}
}

func hiss(f string, vs ...any) {
	print("error: "+f, vs...)
}
