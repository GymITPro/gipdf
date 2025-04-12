package gipdf

import (
	"math/rand"
)

func randomString() string {
	return randStringBytes(20)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func configRunner(pdf *Document, x, y, width, height float64, runnerFunc func(pdf *Document, x, y, width, height float64), configs ...ConfigFunc) {
	var final = func() {
		runnerFunc(pdf, x, y, width, height)
	}
	if len(configs) == 0 {
		final()
		return
	}

	for i := len(configs) - 1; i >= 0; i-- {
		var prevFinal = final
		var f = configs[i]
		final = func() {
			var finalCalled bool
			pdf.SetX(x)
			pdf.SetY(y)
			f(pdf, x, y, width, height, func() {
				if finalCalled {
					return
				}
				finalCalled = true
				prevFinal()
			})
			if !finalCalled {
				prevFinal()
			}
		}
	}

	final()
	return
}
