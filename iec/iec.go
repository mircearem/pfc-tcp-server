package iec

import (
	"bytes"
	"encoding/binary"
)

type Primitive uint8

// @TODO add all data types from IEC61131
const (
	BOOL Primitive = iota // 0
	WORD                  // 1
	INT                   // 2
	REAL                  // 3
)

type Var struct {
	Primitive Primitive
	Payload   []byte
}

func NewVar(t uint8, payload []byte) *Var {
	var size int
	switch t {
	case 0:
		size = 1
	case 1, 2:
		size = 2
	case 3:
		size = 4
	}
	return &Var{
		Primitive: Primitive(t),
		Payload:   payload[:size],
	}
}

func (v *Var) Decode() (any, error) {
	buf := bytes.NewReader(v.Payload)
	switch v.Primitive {
	// IEC BOOL primitive
	case 0:
		var res bool
		if err := binary.Read(buf, binary.LittleEndian, &res); err != nil {
			return nil, err
		}
		return res, nil
	// IEC WORD primitive
	case 1:
		var res uint16
		if err := binary.Read(buf, binary.LittleEndian, &res); err != nil {
			return nil, err
		}
		return res, nil
	// IEC INT primitive
	case 2:
		var res int16
		if err := binary.Read(buf, binary.LittleEndian, &res); err != nil {
			return nil, err
		}
		return res, nil
	// IEC REAL primitive
	case 3:
		var res float32
		if err := binary.Read(buf, binary.LittleEndian, &res); err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, nil
}

func (v *Var) Type() string {
	switch v.Primitive {
	case 0:
		return "BOOL"
	case 1:
		return "WORD"
	case 2:
		return "INT"
	case 3:
		return "REAL"
	default:
		return ""
	}
}
