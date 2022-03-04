package crypt

import (
	"testing"
)

func getTestSign() string {
	return `nHmOEHJC+jXdtGdNeEhPhzjtDtGBo/SJ3SSEa/PT45+Dpasmz3K4Hp9HHF2l0RSriALn8EtegV83wBvn0LKRjg==`
}

func TestSign(t *testing.T) {
	key := getTestKey()
	realsign := getTestSign()
	mes := "hello"
	Setup(key)
	outsign, err := Sign(mes)
	if err != nil {
		t.Error(err)
	}
	if outsign != realsign {
		t.Error(`real sign is not matching with actual`)
	}
}
