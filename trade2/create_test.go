package trade2

import (
	"testing"
)

func TestCreatePool(t *testing.T) {
	newPool := CreatePool()
	if newPool == nil || newPool.Buys == nil || newPool.Sells == nil {
		t.Error("pool is nil")
	}
}
