package mnet

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"log"
)

const (
	PKG_HEAD_BYTES = 8
)

type Packet struct {
	Body interface{}
}

type Payload struct {
	code   uint16
	length uint16 // body length
	crc32  uint32
	body   []byte
}

func (p *Payload) Encode() (data []byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, p.code)
	if err != nil {
		return
	}

	p.length = uint16(len(p.body))
	err = binary.Write(buf, binary.LittleEndian, p.length)
	if err != nil {
		return
	}

	p.crc32 = crc32.ChecksumIEEE(p.body)
	err = binary.Write(buf, binary.LittleEndian, p.crc32)
	if err != nil {
		return
	}

	err = binary.Write(buf, binary.LittleEndian, p.body)
	if err != nil {
		return
	}

	return buf.Bytes(), nil
}

func (p *Payload) Decode(data []byte) (err error, more bool, left []byte) {
	log.Printf("Decode: % x\n", data)
	more = false
	err = nil
	if len(data) < PKG_HEAD_BYTES {
		more = true
		return
	}

	buf := bytes.NewBuffer(data[0:PKG_HEAD_BYTES])
	err = binary.Read(buf, binary.LittleEndian, &p.code)
	if err != nil {
		return
	}

	err = binary.Read(buf, binary.LittleEndian, &p.length)
	if err != nil {
		return
	}
	log.Printf("p.length=%d\n", p.length)

	endPos := PKG_HEAD_BYTES + int(p.length)
	if len(data) < endPos {
		log.Println("need read more data")
		more = true
		return
	}

	log.Printf("p.code=%d\n", p.code)

	err = binary.Read(buf, binary.LittleEndian, &p.crc32)
	if err != nil {
		return
	}

	log.Printf("slice[%d:%d]", PKG_HEAD_BYTES, endPos)

	p.body = data[PKG_HEAD_BYTES:endPos]
	if endPos < len(data) {
		left = data[endPos:]
	}

	return
}
