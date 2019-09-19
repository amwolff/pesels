// Package pesels provides utilities for working with Polish PESEL numbers.
package pesels

import (
	"errors"
	"time"
	"unicode/utf8"
)

var loc *time.Location

func mustLoadLocation() {
	var err error
	if loc, err = time.LoadLocation("Europe/Warsaw"); err != nil {
		panic(err)
	}
}

func init() {
	mustLoadLocation()
}

func btoi(b byte) (int, bool) {
	d := int(b - '0')
	if d >= 0 && d <= 9 {
		return d, true
	}
	return 0, false
}

func valid(pesel []byte) ([]int, bool) {
	if !(utf8.Valid(pesel) && len(pesel) == 11) {
		return nil, false // Invalid length.
	}

	var (
		checksum  int
		rewritten []int
	)
	mfs := []int{9, 7, 3, 1}
	for i, b := range pesel[:10] {
		d, ok := btoi(b)
		if !ok {
			return nil, false // Contains non-digit characters.
		}
		checksum += mfs[i%4] * d
		rewritten = append(rewritten, d)
	}

	ld, ok := btoi(pesel[10])
	if !ok {
		return nil, false // Contains non-digit characters.
	}

	return append(rewritten, ld), checksum%10 == ld
}

// Valid reports wheter `pesel` represents a valid PESEL number (e.g. has a
// valid checksum).
func Valid(pesel string) bool {
	_, ok := valid([]byte(pesel))
	return ok
}

type Sex int

const (
	Female Sex = iota
	Male
)

func (s Sex) String() string {
	if s == Female {
		return "Female"
	}
	return "Male"
}

type PESEL struct {
	DateOfBirth   time.Time
	OrdinalNumber int
	Sex           Sex
	CheckDigit    int
}

var ErrInvalid = errors.New("invalid input string")

// Decode decodes valid `pesel` into a `PESEL`. It returns `ErrInvalid` if
// `pesel` does not represent a valid PESEL number. Date of birth is parsed in
// the Europe/Warsaw time zone.
func Decode(pesel string) (PESEL, error) {
	ret := PESEL{}

	digits, ok := valid([]byte(pesel))
	if !ok {
		return ret, ErrInvalid
	}

	var y int
	var m time.Month
	yShift, mShift := 10*digits[0]+digits[1], time.Month(digits[3])
	switch digits[2] {
	case 8:
		y = 1800 + yShift
		m = mShift
	case 9:
		y = 1800 + yShift
		m = 10 + mShift
	case 0:
		y = 1900 + yShift
		m = mShift
	case 1:
		y = 1900 + yShift
		m = 10 + mShift
	case 2:
		y = 2000 + yShift
		m = mShift
	case 3:
		y = 2000 + yShift
		m = 10 + mShift
	case 4:
		y = 2100 + yShift
		m = mShift
	case 5:
		y = 2100 + yShift
		m = 10 + mShift
	case 6:
		y = 2200 + yShift
		m = mShift
	case 7:
		y = 2200 + yShift
		m = 10 + mShift
	}

	ret.DateOfBirth = time.Date(y, m, 10*digits[4]+digits[5], 0, 0, 0, 0, loc)

	ret.OrdinalNumber = 1000*digits[6] + 100*digits[7] + 10*digits[8] + digits[9]

	if digits[9]%2 == 0 {
		ret.Sex = Female
	} else {
		ret.Sex = Male
	}

	ret.CheckDigit = digits[10]

	return ret, nil
}
