package pesels

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func Test_btoi(t *testing.T) {
	var b byte
	for ; ; b++ {
		d, ok := btoi(b)
		switch b {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			if !(d >= 0 && d <= 9 && ok) {
				t.Errorf("btoi() = %v, want %c", d, b)
			}
		default:
			if !(d == 0 && !ok) {
				t.Errorf("btoi() = %v, want %v", d, 0)
			}
		}
		if b == math.MaxUint8 { // A typical loop would overflow b.
			break
		}
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
			args{"520810485#2"},
			false,
		},
		{
			"contains non-digit characters",
			args{"5208104853@"},
			false,
		},
		{
			"wrong checksum",
			args{"44051401358"},
			false,
		},
		{
			"OK",
			args{"70010198266"},
			true,
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
		{
			"1800-01-01/0186/F/7",
			args{"00810101867"},
			PESEL{
				time.Date(1800, time.January, 1, 0, 0, 0, 0, loc),
				186,
				Female,
				7,
			},
			false,
		},
		{
			"1800-10-01/1307/M/9",
			args{"00900113079"},
			PESEL{
				time.Date(1800, time.October, 1, 0, 0, 0, 0, loc),
				1307,
				Male,
				9,
			},
			false,
		},
		{
			"1900-01-01/0970/F/0",
			args{"00010109700"},
			PESEL{
				time.Date(1900, time.January, 1, 0, 0, 0, 0, loc),
				970,
				Female,
				0,
			},
			false,
		},
		{
			"1900-10-01/1327/M/3",
			args{"00100113273"},
			PESEL{
				time.Date(1900, time.October, 1, 0, 0, 0, 0, loc),
				1327,
				Male,
				3,
			},
			false,
		},
		{
			"2000-01-01/1780/F/6",
			args{"00210117806"},
			PESEL{
				time.Date(2000, time.January, 1, 0, 0, 0, 0, loc),
				1780,
				Female,
				6,
			},
			false,
		},
		{
			"2000-10-01/1511/M/0",
			args{"00300115110"},
			PESEL{
				time.Date(2000, time.October, 1, 0, 0, 0, 0, loc),
				1511,
				Male,
				0,
			},
			false,
		},
		{
			"2100-01-01/1072/F/0",
			args{"00410110720"},
			PESEL{
				time.Date(2100, time.January, 1, 0, 0, 0, 0, loc),
				1072,
				Female,
				0,
			},
			false,
		},
		{
			"2100-10-01/0295/M/0",
			args{"00500102950"},
			PESEL{
				time.Date(2100, time.October, 1, 0, 0, 0, 0, loc),
				295,
				Male,
				0,
			},
			false,
		},
		{
			"2200-01-01/0220/F/6",
			args{"00610102206"},
			PESEL{
				time.Date(2200, time.January, 1, 0, 0, 0, 0, loc),
				220,
				Female,
				6,
			},
			false,
		},
		{
			"2200-10-01/0909/M/0",
			args{"00700109090"},
			PESEL{
				time.Date(2200, time.October, 1, 0, 0, 0, 0, loc),
				909,
				Male,
				0,
			},
			false,
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
