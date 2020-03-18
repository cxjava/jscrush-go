package main

import (
	"bytes"
	"errors"
	"strings"
)

const LowWord = 2
const HiWord = 50

type Script struct {
	Buf []byte
}

// Creates a dictionary of the given word length and counts ocurrences.
func dict(buf []byte, wlen int) map[string]int {

	l := len(buf)

	dict := make(map[string]int, l)

	for i := 0; i+wlen < l; i++ {
		s := string(buf[i : i+wlen])
		if _, ok := dict[s]; ok == false {
			dict[s] = 0
		} else {
			dict[s] = dict[s] + 1
		}
	}

	return dict
}

// Returns the first unused byte, we need it to store a chunk.
func getUnusedByte(buf []byte) (byte, error) {
	var c byte
	for i := byte(126); i > byte(0); i-- {
		switch i {
		case
			// Non-printable bytes.
			byte(10),
			byte(13),
			byte(34),
			byte(39),
			byte(92):
		default:
			c = i
			if bytes.Contains(buf, []byte{c}) == false {
				return c, nil
			}
		}
	}
	return byte(0), errors.New("No more unused bytes\n")
}

// Takes a buffer and its keys and reconstructs the original string.
func Uncompress(buf []byte, keys []byte) ([]byte, error) {
	for i, n := 0, len(keys); i < n; i++ {

		slices := bytes.Split(buf, []byte{keys[i]})
		l := len(slices)

		buf = bytes.Join(slices[0:l-1], slices[l-1])
	}

	return buf, nil
}

// compresses a buffer into a chunk and chunk keys.
func compress(buf []byte) ([]byte, []byte, error) {
	keys := make([]byte, 0, int(len(buf)*8/10))

	mv := -1

	for true {

		t, err := getUnusedByte(buf)

		if err != nil {
			break
		}

		mk := ""

		for i := LowWord; i < HiWord; i++ {

			d := dict(buf, i)

			for k, v := range d {
				if v > 1 {
					// Calculating lenght.
					tlen := len(buf) - v*len(k) + v

					slen := tlen + len(k) + len(keys) + 2

					if slen < mv || mv < 0 {
						if tlen < len(buf) {
							mk = k
							mv = slen
						}
					}
				}
			}

		}

		if mk != "" {

			buf = bytes.Replace(buf, []byte(mk), []byte{t}, -1)
			buf = append(buf, t)
			buf = append(buf, []byte(mk)...)

			keys = append([]byte{t}, keys...)

		} else {
			break
		}
	}

	return buf, keys, nil
}

func SmartQuotes(str string) string {
	a := len(strings.Split(str, `"`))
	b := len(strings.Split(str, `'`))

	if a >= b {
		return `'` + strings.Replace(str, `'`, `\'`, -1) + `'`
	} else {
		return `"` + strings.Replace(str, `"`, `\"`, -1) + `"`
	}
}

// Uses buf and key to create the javascript decompressor.
func pack(buf []byte, keys []byte) string {
	return `_=` + SmartQuotes(string(buf)) + `;for(Y in $="` + string(keys) + `")with(_.split($[Y]))_=join(pop());eval(_)`
}

func JsCrush(buf []byte) string {
	t, k, err := compress(buf)
	if err == nil {
		return pack(t, k)
	}
	return ""
}
