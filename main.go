// CLI utility for sorting strings from file
// Usage: sort [-nru] [-k START] [FILE]...
package main

import (
	"log"

	"github.com/lastlife77/sort/cmd"
)

func main() {
	log.SetFlags(0)

	cmd.Execute()
}
