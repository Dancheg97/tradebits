package logs

import "testing"

func TestLogs(t *testing.T) {
	Init()
	Info("some info", "with some vars", 1, 2, []byte{123, 21, 23, 56})
	Error("some info", "with some vars", 1, 2, []byte{123, 21, 23, 56})
	Critical("some info", "with some vars", 1, 2, []byte{123, 21, 23, 56})
	Warning("some info", "with some vars", 1, 2, []byte{123, 21, 23, 56})
}
