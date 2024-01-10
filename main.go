package main

import (
	"strings"

	"github.com/Aldoihm/module-examples/slices"
)

func main() {
	list := []string{"EDteam", "gophers", "golang"}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})

}
