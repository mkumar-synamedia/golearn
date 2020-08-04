package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"testing"
)

type ldeAULocIdx struct {
	// Allocation Unit Type / Magic
	auType uint64
	// Allocation Unit Checksum
	auChecksum uint32
	// Start Index of the local directory entry location
	startIndex uint32
	// Directory replica count
	dirReplicaCount uint8
	// Reserving 3 bytes for byte alignment
	reserved1 [3]byte
	// First unused offset
	firstUnusedOffset uint32
	// Reserving 8 bytes for byte alignment
	reserved2 [8]byte
	// Local Directory Entry AU Location Sets
	//ldeAULocationSet [][]uint64
}

func BenchmarkHandMarshal(b *testing.B) {
	l := ldeAULocIdx{}
	buf := make([]byte, 32)
	for n := 0; n < b.N; n++ {
		binary.LittleEndian.PutUint64(buf[0:], l.auType)
		offset := 8
		binary.LittleEndian.PutUint32(buf[offset:], l.auChecksum)
		offset += 4
		binary.LittleEndian.PutUint32(buf[offset:], l.startIndex)
		offset += 4
		PutUint8(buf[offset:], l.dirReplicaCount)
		offset += 1
		// add reserved bytes
		offset += 3
		binary.LittleEndian.PutUint32(buf[offset:], l.firstUnusedOffset)
		copy(buf, []byte{})
	}
}

func BenchmarkBinaryLibMarshal(b *testing.B) {
	l := ldeAULocIdx{}
	buf := new(bytes.Buffer)
	for n := 0; n < b.N; n++ {
		binary.Write(buf, binary.LittleEndian, l)
		buf.Reset()
	}
}

func BenchmarkGobLibMarshal(b *testing.B) {
	l := ldeAULocIdx{}
	buf := new(bytes.Buffer)
	for n := 0; n < b.N; n++ {
		gob.NewEncoder(buf).Encode(l)
		buf.Reset()
	}
}

func PutUint8(b []byte, v uint8) {
	_ = b[0] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
}

func main() {

}
