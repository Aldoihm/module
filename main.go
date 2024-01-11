package main

import (
	"strings"

	"github.com/Aldoihm/module/slices"
	"rsc.io/quote"
)

func main() {
	list := []string{"EDteam", "gophers", "golang", "sotec", quote.Hello()}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})

}
