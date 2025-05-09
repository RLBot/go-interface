// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A boolean value located in a separate struct allowing for optional floats elsewhere.
type BoolT struct {
	Val bool `json:"val"`
}

func (t *BoolT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateBool(builder, t.Val)
}
func (rcv *Bool) UnPackTo(t *BoolT) {
	t.Val = rcv.Val()
}

func (rcv *Bool) UnPack() *BoolT {
	if rcv == nil {
		return nil
	}
	t := &BoolT{}
	rcv.UnPackTo(t)
	return t
}

type Bool struct {
	_tab flatbuffers.Struct
}

func (rcv *Bool) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Bool) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Bool) Val() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Bool) MutateVal(n bool) bool {
	return rcv._tab.MutateBool(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func CreateBool(builder *flatbuffers.Builder, val bool) flatbuffers.UOffsetT {
	builder.Prep(1, 1)
	builder.PrependBool(val)
	return builder.Offset()
}
