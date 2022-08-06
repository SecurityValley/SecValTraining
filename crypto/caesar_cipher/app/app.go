package app

import (
	"bytes"

	"github.com/SecurityValley/SecValTraining/crypto/caesar_cipher/utils"
)

func Encrypt(input string, rot int) string {

	out := bytes.Buffer{}
	da := utils.GetDefaultAlphabet()

	for _, v := range input {
		if v >= 97 && v <= 122 {
			for i, l := range da {
				if v == l {
					out.WriteRune(da[((i + rot) % 26)])
				}
			}
		} else {
			out.WriteRune(v)
		}
	}

	return out.String()
}

func Decrypt(input string, rot int) string {
	out := bytes.Buffer{}
	da := utils.GetDefaultAlphabet()

	for _, v := range input {
		if v >= 97 && v <= 122 {
			for i, l := range da {
				if v == l {
					index := ((i-rot)%26 + 26) % 26
					out.WriteRune(da[index])
				}
			}
		} else {
			out.WriteRune(v)
		}
	}

	return out.String()
}

func Crack(input string) []string {

	results := make([]string, 0)

	for i := 1; i < 26; i++ {
		out := bytes.Buffer{}
		da := utils.GetDefaultAlphabet()

		for _, v := range input {
			if v >= 97 && v <= 122 {
				for a, l := range da {
					if v == l {
						out.WriteRune(da[((a + i) % 26)])
					}
				}
			} else {
				out.WriteRune(v)
			}
		}

		results = append(results, out.String())
	}

	return results
}
