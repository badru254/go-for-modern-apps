package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	Msg io.Writer
}

func (c *Capper) Write(data []byte) (int, error) {
	//Option A
	diff := byte('a' - 'A')
	out := make([]byte, len(data))

	for i, c := range data {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}

		out[i] = c

	}

	//Option A output
	return c.Msg.Write(out)

	// //Option B
	// str := string(data)

	// strUpper := strings.ToUpper(str)

	// //Option B output
	// return c.Msg.Write([]byte(strUpper))

}

func main() {

	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")

}
