package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		out("[NOTE] No arguments given, using example data.\n[NOTE]To read stdin pass '--' as the argument\n")
		args = append(args, "Nord̩n Näs Åmål Västra Mêléé - 가구")
	}

	for _, a := range args {
		if a == "--" {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				outDecodeAll(scanner.Text())
			}
		} else {
			outDecodeAll(a)
		}
	}
}

func halt(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, "[HALT] "+format, a...)
	os.Exit(1)
}

func out(format string, a ...interface{}) {
	if _, err := fmt.Fprintf(os.Stdout, format, a...); err != nil {
		halt(err.Error())
	}
}

func outDecode(input string) {
	i, u, x := decodeUTF8([]byte(input))
	out("%s\ni:%s\nu:%s\nx:%s\n", input, i, u, x)
}

func outDecodeAll(input string) {
	out("Raw Input: ")
	outDecode(input)
	out("\n")

	out("NFC: ")
	outDecode(norm.NFC.String(input))
	out("\n")

	out("NFKC: ")
	outDecode(norm.NFKC.String(input))
	out("\n")

	out("NFD: ")
	outDecode(norm.NFD.String(input))
	out("\n")

	out("NFKD: ")
	outDecode(norm.NFKD.String(input))
	out("\n")
}

// Create formatted []byte of input (in) with unicode code points (cp) and the hex encoding (hex).
//in:     N    o    r    d    ̩    n
//cp:  004E 006F 0072 0064 0329 006E
//hex:   4E   6F   72   64 CCA9   6E
func decodeUTF8(input []byte) (in, cp, hex []byte) {
	for len(input) > 0 {
		if len(in) > 0 {
			in = append(in, ' ')
			cp = append(cp, ' ')
			hex = append(hex, ' ')
		}

		r, size := utf8.DecodeRune(input)
		cur := input[:size]
		input = input[size:]

		width := fmt.Sprintf("%d", size<<1)
		in = append(in, []byte(fmt.Sprintf("%"+width+"s", cur))...)
		cp = append(cp, []byte(fmt.Sprintf("%0"+width+"X", r))...)
		hex = append(hex, []byte(fmt.Sprintf("%"+width+"X", cur))...)
	}
	return
}
