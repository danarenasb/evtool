package evtool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeErr(t *testing.T) {
	tests := []struct {
		in string
	}{

		// Valid.
		{in: "email@yahoo.com"},
		{in: "e.m.a.i.l.ema.il@hotmail.com"},
		{in: "email+extra@lycos.com"},
	}

	for _, tt := range tests {
		_, err := NormalizeGmail(tt.in)
		if err != nil {
			assert.EqualErrorf(t, err, "Not a Gmail address", "Not a Gmail address")
		}
	}
}

func TestNormalizeSuccess(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{

		// Valid.
		{in: "email@gmail.com", out: "email@gmail.com"},
		{in: "e.m.a.i.l.ema.il@gmail.com", out: "emailemail@gmail.com"},
		{in: "email+extra@gmail.com", out: "email@gmail.com"},
	}

	for _, tt := range tests {
		s, err := NormalizeGmail(tt.in)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.out, s)
	}
}

func TestValidateSuccess(t *testing.T) {
	tests := []struct {
		in string
	}{

		// Valid.
		{in: "email@gmail.com"},
		{in: "e.m.a.i.l.ema.il@yahoo.com"},
		{in: "email+extra@hotmail.com"},
	}

	for _, tt := range tests {
		valid := Validate(tt.in)
		assert.Equal(t, valid, true)
	}
}

func TestValidateFailure(t *testing.T) {
	tests := []struct {
		in     string
		result bool
	}{

		// Valid.
		{in: "e.m.a.i.l.ema.il@notaval$%iddomain.com", result: false},
	}

	for _, tt := range tests {
		valid := Validate(tt.in)
		assert.Equal(t, valid, tt.result)
	}
}
