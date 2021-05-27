package __tests

import (
	"fmt"
	"strings"
)

func Passed(module string, function string, description string) {

	fmt.Println("\033[32m[PASSED]  >>>  ",
		"[MODULE]:", module, strings.Repeat(" ", 8-len(module)),
		"[FUNCTION]:", function, strings.Repeat(" ", 8-len(function)),
		"[DESCRIPTION]:", description,
		"\033[0m",
	)
}

func Failed(module string, function string, description string) {
	fmt.Println("\033[31m[FAILED]  >>>  ",
		"[MODULE]:", module, strings.Repeat(" ", 8-len(module)),
		"[FUNCTION]:", function, strings.Repeat(" ", 8-len(function)),
		"[DESCRIPTION]:", description,
		"\033[0m",
	)
}
