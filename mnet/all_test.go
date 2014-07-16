package mnet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestB(t *test.T) {

}

func TestA(t *testing.T) {
	t.Skip("...")
	data := "hello"
	pkgLength := len([]byte(data))

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, uint16(pkgLength))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	err = binary.Write(buf, binary.LittleEndian, []byte(data))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	fmt.Println("haaa")
	fmt.Printf("% x", buf.Bytes())

	buf2 := bytes.NewBuffer(buf)
	binary.Read(buf2, binary.LittleEndian, data)

	fmt.Println("")
}
