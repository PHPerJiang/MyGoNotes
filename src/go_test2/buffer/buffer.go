package buffer

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func GetStdout() {
	var B bytes.Buffer
	B.Write([]byte("Hello"))
	_,err := fmt.Fprintf(&B, " Word！")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(os.Stdout,&B)
	if err != nil  {
		panic(err)
	}
}
