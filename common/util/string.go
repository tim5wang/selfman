package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

const encodeURL = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

func Print(os ...interface{}) {
	for _, o := range os {
		printJSONIndent(o)
	}
}

func PrintJSON(o interface{}) {
	s, _ := json.Marshal(o)
	fmt.Println(string(s))
}

func printJSONIndent(o interface{}) {
	s, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(s))
}
func ToBase64(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}
func Base64(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}
func Base64ToString(data string) string {
	b, _ := base64.URLEncoding.DecodeString(data)
	return string(b)
}
func ShortBase64(in uint64) string {
	bys := Uint64ToBytes(in)
	ss := strings.TrimLeft(string(bys), "A")
	return Base64(ss)
}

func ShortBase64ToUint64(in string) uint64 {
	ss := Base64ToString(in)
	_ = ss
	return 0
}

func encode(charset string, in []byte) string {
	c := len(charset)
	res := make([]byte, 0)
	batch := int(math.Floor(math.Log2(float64(c))))
	var left, right = 0, batch
	total := 8 * len(in)
	for left < total {
		if right > total {
			right = total
		}
		lb := left / 8
		li := left % 8
		rb := right / 8
		ri := right % 8
		if li == 0 && lb > 0 {
			lb--
		}
		if ri == 0 && rb > 0 {
			rb--
			ri = 8
		}
		var index byte
		if lb != rb {
			l := in[lb]
			r := in[rb]
			l <<= li           // 左边byte里不要的移出去
			l >>= byte(8 - ri) // 往右移动到正确的位置
			index |= l         // 赋值给index
			r >>= byte(8 - ri) // 左边byte里不要的移出去
			index |= r
		} else {
			b := in[lb]
			b <<= li              // 向左移除不要的位
			b >>= byte(8 - batch) // 向右移除不要的位
			index = b
		}
		left += batch
		right += batch
		res = append(res, charset[index])
	}
	return string(res)
}

func decode(charset string, in string) []byte {
	res := make([]byte, 0)
	c := len(charset)
	batch := byte(math.Floor(math.Log2(float64(c))))
	var vb, index byte
	for _, v := range in {
		var b byte
		for i := range charset {
			if charset[i] == byte(v) {
				b = byte(i)
				break
			}
		}
		if index+batch < 8 {
			index += batch
			vb <<= batch
			vb |= b
		} else {
			ri := (index + batch) - 8
			li := batch - ri
			index = (index + batch) % 8
			vb <<= li
			lb := b >> (batch - li)
			vb |= lb
			res = append(res, vb)
			vb = 0
		}
	}
	return nil
}
