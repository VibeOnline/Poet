package utils

import "time"

func Sleep(seconds float32) {
	time.Sleep(time.Millisecond * time.Duration(seconds*1000))
}
