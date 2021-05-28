package logs

import (
	"fmt"
	"strings"
	"testing"
)

func TestPassed(module string, function string, descr string) {
	fmt.Println("\033[32m[PASSED]  >>>  ",
		"[MODULE]:", module, strings.Repeat(" ", 12-len(module)),
		"[FUNCTION]:", function, strings.Repeat(" ", 12-len(function)),
		"[DESCRIPTION]:", descr,
		"\033[0m",
	)

}

func TestFailed(module string, function string, descr string, t *testing.T) {
	fmt.Println("\033[31m[FAILED]  >>>  ",
		"[MODULE]:", module, strings.Repeat(" ", 12-len(module)),
		"[FUNCTION]:", function, strings.Repeat(" ", 12-len(function)),
		"[DESCRIPTION]:", descr,
		"\033[0m",
	)
	t.FailNow()
}
