package mnet

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"fmt"
	"github.com/miraclew/mrs/pb"
	"hash/crc32"
	// "log"
)

const (
	PKG_HEAD_BYTES = 8
)

type Packet struct {
	Body interface{}
}

type Message struct {
	Code pb.Code
	MSG  proto.Message
}

func (m *Message) String() string {
	return fmt.Sprintf("Message: Code=%#v MSG: %#v", m.Code, m.MSG.String())
}

type Payload struct {
	Code   uint16
	Length uint16 // body length
	Crc32  uint32
	Body   []byte
}

func (p *Payload) Encode() (data []byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, p.Code)
	if err != nil {
		return
	}

	p.Length = uint16(len(p.Body))
	err = binary.Write(buf, binary.LittleEndian, p.Length)
	if err != nil {
		return
	}

	p.Crc32 = crc32.ChecksumIEEE(p.Body)
	err = binary.Write(buf, binary.LittleEndian, p.Crc32)
	if err != nil {
		return
	}

	err = binary.Write(buf, binary.LittleEndian, p.Body)
	if err != nil {
		return
	}

	return buf.Bytes(), nil
}

func (p *Payload) Decode(data []byte) (err error, more bool, left []byte) {
	// log.Printf("Decode: % x\n", data)
	more = false
	err = nil
	if len(data) < PKG_HEAD_BYTES {
		more = true
		return
	}

	buf := bytes.NewBuffer(data[0:PKG_HEAD_BYTES])
	err = binary.Read(buf, binary.LittleEndian, &p.Code)
	if err != nil {
		return
	}

	err = binary.Read(buf, binary.LittleEndian, &p.Length)
	if err != nil {
		return
	}
	// log.Printf("p.Length=%d\n", p.Length)

	endPos := PKG_HEAD_BYTES + int(p.Length)
	if len(data) < endPos {
		// log.Println("need read more data")
		more = true
		return
	}

	// log.Printf("p.Code=%d\n", p.Code)

	err = binary.Read(buf, binary.LittleEndian, &p.Crc32)
	if err != nil {
		return
	}

	// log.Printf("slice[%d:%d]", PKG_HEAD_BYTES, endPos)

	p.Body = data[PKG_HEAD_BYTES:endPos]
	if endPos < len(data) {
		left = data[endPos:]
	}

	return
}
