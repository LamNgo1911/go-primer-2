package myutf8

import (
	"testing"
)

func TestEncodeRune(t *testing.T) {
	for _, input := range byteInputs {
		t.Run(input.name, func(t *testing.T) {
			got, gotOK := EncodeRune(input.p, input.r)

			if gotOK != input.ok {
				t.Errorf("EncodeRune(% X, %c) ok = %t, want %t",
					input.p, input.r, gotOK, input.ok)
			}

			if got != input.output {
				t.Errorf("EncodeRune(% X, %c) = %d, want %d",
					input.p, input.r, got, input.output)
			}
		})
	}
}
