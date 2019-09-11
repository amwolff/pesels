package pesels

import (
	"reflect"
	"testing"
)

func Test_btoi(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			"1",
			args{'1'},
			1,
			true,
		},
		{
			"2",
			args{'2'},
			2,
			true,
		},
		{
			"3",
			args{'3'},
			3,
			true,
		},
		{
			"4",
			args{'4'},
			4,
			true,
		},
		{
			"5",
			args{'5'},
			5,
			true,
		},
		{
			"6",
			args{'6'},
			6,
			true,
		},
		{
			"7",
			args{'7'},
			7,
			true,
		},
		{
			"8",
			args{'8'},
			8,
			true,
		},
		{
			"9",
			args{'9'},
			9,
			true,
		},
		{
			"0",
			args{'0'},
			0,
			true,
		},
		{
			"-",
			args{'-'},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := btoi(tt.args.b)
			if got != tt.want {
				t.Errorf("btoi() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("btoi() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestValid(t *testing.T) {
	type args struct {
		pesel string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"invalid length",
			args{"520810485320"},
			false,
		},
		{
			"contains non-digit characters",
			args{"520810485.2"},
			false,
		},
		{
			"contains non-digit characters",
			args{"5208104853."},
			false,
		},
		{
			"wrong checksum",
			args{"44051401358"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.pesel); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		pesel string
	}
	tests := []struct {
		name    string
		args    args
		want    PESEL
		wantErr bool
	}{
		{
			"invalid input",
			args{"44051401358"},
			PESEL{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.pesel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
