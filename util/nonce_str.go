package util

import "github.com/fxrobot/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
