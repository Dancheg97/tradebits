package redis

// import (
// 	"testing"
// 	"time"
// )

// func TestLockID(t *testing.T) {
// 	Setup("")
// 	timeout := time.After(3 * time.Second)
// 	done := make(chan bool)
// 	go func() {
// 		lockBytes := make([]byte, 64)
// 		lockBytes[0] = 65
// 		lockBytes[1] = 69
// 		Lock(lockBytes)
// 		Unlock(lockBytes)
// 		done <- true
// 	}()
// 	select {
// 	case <-timeout:
// 		t.Fatal("Test didn't finish in time")
// 	case <-done:
// 	}
// }
