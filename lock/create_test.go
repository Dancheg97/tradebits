package lock

import "testing"

func TestGenerateLockers(t *testing.T) {
	for _, mapOrPointer := range blockers {
		if mapOrPointer == nil {
			t.Error("map should be initialized, and should not be nil pointer")
		}
	}
}
