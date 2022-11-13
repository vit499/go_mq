package utils

//import "log"

// '3' -> 0x03, 'A' -> 0x0A
func char2Bcd(c uint8) uint8 {
	var b uint8
	if c > 0x39 {
		b = c - 0x37
	} else {
		b = c - 0x30
	}
	return b
}
// [ 0x32, 0x33 ] -> 0x23
func symHex2Hex(s []uint8, offset int) uint8 {
	var r uint8
	r = char2Bcd(s[offset]) << 4
	r += char2Bcd(s[offset + 1])
	return r
}
// "23" -> 0x23
func sym2Hex(s string, offset int) uint8 {
	b := []uint8(s)
	r := symHex2Hex(b, offset)
	return r
}

func Str2Hex(s string) []uint8 {
	len_dst := len(s) / 2
	dst := make([]uint8, len_dst)
	for i := 0; i < len_dst; i++ {
		dst[i] = sym2Hex(s, i * 2)
	}
	//log.Printf("src= %s, dst= % X ", s, dst)
	return dst
}
//------------------------------

// 0xC1 -> [1,0,0,0, 0,0,1,1]
func hex2Bits(c uint8) []uint8 {
	b := make([]uint8, 8)
	for i := 0; i < 8; i++ {
		if (c & (1 << i) != 0) {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
	return b
}
func hexBuf2Bits(src []uint8) []uint8 {
	len_src := len(src)
	dst := make([]uint8, len_src * 8)
	for i := 0; i < len_src; i++ {
		a := hex2Bits(src[i])
		for j := 0; j < 8; j++ {
			dst[i * 8 + j] = a[j]
		}
	}
	return dst
}
// "0482" -> [0,0,1,0,0,0,0,0, 0,1,0,0,0,0,0,1]
func Str2Bits(s string) []uint8 {
	b := Str2Hex(s)
	dst := hexBuf2Bits(b)
	//log.Printf("src= %s, dst= % X ", s, dst)
	return dst
}
//====================================

func bcd2CharH(b uint8) uint8 {
	var c uint8
	c = b >> 4 & 0x0f
	if c > 9 {
		c = c + 0x37
	} else {
		c = c + 0x30
	}
	return c
}
func bcd2CharL(b uint8) uint8 {
	var c uint8
	c = b & 0x0f
	if c > 9 {
		c = c + 0x37
	} else {
		c = c + 0x30
	}
	return c
}
// 0x25 0x31 -> "2531"
func Hex2Str(buf []uint8, len_src int) string {
	len_dst := len_src * 2
	strBuf := make([]uint8, len_dst)
	for i := 0; i < len_src; i++ {
		strBuf[i * 2] = bcd2CharH(buf[i])
		strBuf[i * 2 + 1] = bcd2CharL(buf[i])
	}
	s := string(strBuf)
	return s
}