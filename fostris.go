package fostris

import (
	"strings"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const ostrisHigh = "OSTRIS"
const ostrisLow = "ostris"

func Encode(vs string) string {
	v := []byte(vs)
	out := ""
	for i := 0; i < len(v); i++ {
		if v[i] == ' ' {
			out += " "
			continue
		}
		pos := strings.IndexByte(encodeStd, v[i])
		outRow := []byte(ostrisLow)
		for j := 5; j >= 0; j-- {
			if pos&(1<<uint(j)) != 0 {
				outRow[j] = ostrisHigh[j]
			} else {
				outRow[j] = ostrisLow[j]
			}
		}
		out = out + "f" + string(outRow)
	}
	return out
}

func Decode(vs string) string {
	v := []byte(vs)
	out := []byte{}
	for i := 0; i < len(v); i += 7 {
		if v[i] == ' ' {
			out = append(out, ' ')
			i -= 6
			continue
		}
		pos := 0
		for j := 0; j < 6; j++ {
			if v[i+j+1] == ostrisHigh[j] {
				pos |= 1 << uint(j)
			}
		}
		out = append(out, encodeStd[pos])
	}
	return string(out)
}
