package calc

import "testing"

func TestGetKeys(t *testing.T) {
	keys := Gen()
	lenSum := len(keys.PersPriv) + len(keys.PersPub) + len(keys.MesPriv) + len(keys.MesPub)
	if lenSum < 4330 || lenSum > 4340 {
		t.Error("failed to generate correct keys")
	}
}
