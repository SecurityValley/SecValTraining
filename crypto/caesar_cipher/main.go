package main

import (
	"fmt"

	"github.com/SecurityValley/SecValTraining/crypto/caesar_cipher/app"
	"github.com/SecurityValley/SecValTraining/crypto/caesar_cipher/utils"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	mode  = kingpin.Arg("mode", "name of mode. encrypt, decrypt or crack").Required().String()
	input = kingpin.Arg("input", "input to process select mode").Required().String()
	rot   = kingpin.Flag("rot", "number of rotations").Short('r').Default("3").Int()
)

func main() {
	kingpin.Parse()
	switch *mode {
	case "encrypt":
		res := app.Encrypt(utils.Clean(*input), (*rot % 26))
		fmt.Println(res)
	case "decrypt":
		res := app.Decrypt(utils.Clean(*input), (*rot % 26))
		fmt.Println(res)
	case "crack":
		res := app.Crack(utils.Clean(*input))
		for i, v := range res {
			fmt.Printf("res for rot: %d\n%s\n\n", 26-(i+1), v)
		}

	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println(`
	wrong program parameters
	************************
	please run this program like follows:
	
	./caesar_cipher encrypt <input> --rot=<number of rotation rounds>
	./caesar_cipher decrypt <input> --rot=<number of rotation rounds>
	./caesar_cipher crack <input>
	`)
}
