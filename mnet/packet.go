package mnet

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
)

const (
	PKG_HEAD_BYTES = 8
)

type Packet struct {
	Body interface{}
}

type Payload struct {
	length uint16 // body length
	cmd    uint16
	crc32  uint32
	body   []byte
}

func (p *Payload) Encode() (data []byte, err error) {
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, p.length)
	if err != nil {
		return
	}

	err = binary.Write(buf, binary.LittleEndian, p.cmd)
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
	more = false
	err = nil
	if len(data) < PKG_HEAD_BYTES {
		more = true
		return
	}

	buf := bytes.NewBuffer(data[0 : PKG_HEAD_BYTES-1])
	err = binary.Read(buf, binary.LittleEndian, &p.length)
	if err != nil {
		return
	}

	if len(data) < PKG_HEAD_BYTES+int(p.length) {
		more = true
		return
	}

	err = binary.Read(buf, binary.LittleEndian, &p.cmd)
	if err != nil {
		return
	}

	err = binary.Read(buf, binary.LittleEndian, &p.crc32)
	if err != nil {
		return
	}

	p.body = data[PKG_HEAD_BYTES : p.length-PKG_HEAD_BYTES]
	left = data[PKG_HEAD_BYTES+p.length:]
	return
}
