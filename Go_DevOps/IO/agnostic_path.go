package main

import (
	"fmt"
	_ "runtime" // Decided to use an anonymous import, where a package is needed to be loaded but not using its functionality directly. Without _, we would receive a compile error for not using the imported package.
	"os"
	"path/filepath"
)

// Using the runtime package, you can detect the OS and platform you are running on:
// func main() {
// 	fmt.Println(runtime.GOOS) // linux, darwin, ...
// 	fmt.Println(runtime.GOARCH) // amd64, arm64, ...
// }

// Say we want to access a configuration file - config.json, thats stored in the config/ directory, which is in the same directory as our binary.

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting work directory:", err)
		os.Exit(1)
	}
	content, err := os.ReadFile(filepath.Join(wd, "config", "config.json"))
	fmt.Printf("%s", content)
}

// filepath.Join() allows us to join the components of our path into a single path. This fills in the OS-specific directory separators for you and uses the native pathing rules.

