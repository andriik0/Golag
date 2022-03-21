package main

import (
	"testing"
)

func Test_arithmeticSigns(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Negative test",
			args{r: 'A'},
			false,
		},
		{
			"Positive test",
			args{r: '+'},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arithmeticSigns(tt.args.r); got != tt.want {
				t.Errorf("arithmeticSigns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		numbers string
	}
	tests := []struct {
		name       string
		args       args
		want       int
		wantResult bool
	}{
		{
			"Negative result",
			args{"2+5"},
			5,
			false,
		},

		{
			"Positive result",
			args{"9+5"},
			14,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := calculate(tt.args.numbers); err == nil && got != tt.want && tt.wantResult {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSign(t *testing.T) {
	type args struct {
		operation string
		operators *[]string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"Valid result",
			args{"2*8",
				&[]string{"2", "8"},
			},
			'*',
		},
		{"Valid result two digits",
			args{"24+84",
				&[]string{"24", "84"},
			},
			'+',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSign(tt.args.operation, tt.args.operators); got != tt.want {
				t.Errorf("getSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
