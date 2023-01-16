// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package generated

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type VerificationAddEthAddressBody struct {
	_tab flatbuffers.Table
}

func GetRootAsVerificationAddEthAddressBody(buf []byte, offset flatbuffers.UOffsetT) *VerificationAddEthAddressBody {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &VerificationAddEthAddressBody{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsVerificationAddEthAddressBody(buf []byte, offset flatbuffers.UOffsetT) *VerificationAddEthAddressBody {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &VerificationAddEthAddressBody{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *VerificationAddEthAddressBody) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *VerificationAddEthAddressBody) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *VerificationAddEthAddressBody) Address(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *VerificationAddEthAddressBody) AddressLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *VerificationAddEthAddressBody) AddressBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *VerificationAddEthAddressBody) MutateAddress(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *VerificationAddEthAddressBody) EthSignature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *VerificationAddEthAddressBody) EthSignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *VerificationAddEthAddressBody) EthSignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *VerificationAddEthAddressBody) MutateEthSignature(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *VerificationAddEthAddressBody) BlockHash(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *VerificationAddEthAddressBody) BlockHashLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *VerificationAddEthAddressBody) BlockHashBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *VerificationAddEthAddressBody) MutateBlockHash(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func VerificationAddEthAddressBodyStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func VerificationAddEthAddressBodyAddAddress(builder *flatbuffers.Builder, address flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(address), 0)
}
func VerificationAddEthAddressBodyStartAddressVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func VerificationAddEthAddressBodyAddEthSignature(builder *flatbuffers.Builder, ethSignature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(ethSignature), 0)
}
func VerificationAddEthAddressBodyStartEthSignatureVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func VerificationAddEthAddressBodyAddBlockHash(builder *flatbuffers.Builder, blockHash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(blockHash), 0)
}
func VerificationAddEthAddressBodyStartBlockHashVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func VerificationAddEthAddressBodyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}