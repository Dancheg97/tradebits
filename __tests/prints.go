package __tests

import (
	"fmt"
)

func Passed(module string, function string, description string) {
	fmt.Println("\033[32m[PASSED] ===>",
		"[MODULE]:", module,
		"[FUNCTION]:", function,
		"[DESCRIPTION]:", description,
		"\033[0m",
	)
}

func Failed(module string, function string, description string) {
	fmt.Println("\033[31m[FAILED] ===>",
		"[MODULE]:", module,
		"[FUNCTION]:", function,
		"[DESCRIPTION]:", description,
		"\033[0m",
	)
}
