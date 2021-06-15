//+build !windows

package main

import (
	"github.com/zalando/go-keyring"
)

func getChromeMacOS() (string, error) {
	return keyring.Get("Chrome Safe Storage", "Chrome")
}

//DPApi ...
func DPApi(data []byte) ([]byte, error) {
	return nil, nil
}
