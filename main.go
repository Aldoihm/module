package main

import (
	"strings"

	"github.com/Aldoihm/module/slices"
)

func main() {
	list := []string{"EDteam", "gophers", "golang", "sotec"}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "s")
	})

}
