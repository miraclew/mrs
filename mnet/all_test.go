package mnet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
)

func TestB(t *testing.T) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	handleErr(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	handleErr(err)

	payload := Payload{cmd: 1, body: []byte("hello")}

	b, err := payload.Encode()
	handleErr(err)

	fmt.Printf("Send: % x\n", b)
	_, err = conn.Write(b)
	handleErr(err)
	result, err := ioutil.ReadAll(conn)
	handleErr(err)
	fmt.Println(string(result))
	// conn.Close()
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("FATAL error: %s", err.Error())
		os.Exit(1)
	}
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

	buf2 := bytes.NewBuffer(buf.Bytes())
	binary.Read(buf2, binary.LittleEndian, data)

	fmt.Println("")
}
