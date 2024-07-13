// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BoxShapeT struct {
	Length float32 `json:"length"`
	Width float32 `json:"width"`
	Height float32 `json:"height"`
}

func (t *BoxShapeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	BoxShapeStart(builder)
	BoxShapeAddLength(builder, t.Length)
	BoxShapeAddWidth(builder, t.Width)
	BoxShapeAddHeight(builder, t.Height)
	return BoxShapeEnd(builder)
}

func (rcv *BoxShape) UnPackTo(t *BoxShapeT) {
	t.Length = rcv.Length()
	t.Width = rcv.Width()
	t.Height = rcv.Height()
}

func (rcv *BoxShape) UnPack() *BoxShapeT {
	if rcv == nil {
		return nil
	}
	t := &BoxShapeT{}
	rcv.UnPackTo(t)
	return t
}

type BoxShape struct {
	_tab flatbuffers.Table
}

func GetRootAsBoxShape(buf []byte, offset flatbuffers.UOffsetT) *BoxShape {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BoxShape{}
	x.Init(buf, n+offset)
	return x
}

func FinishBoxShapeBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsBoxShape(buf []byte, offset flatbuffers.UOffsetT) *BoxShape {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BoxShape{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedBoxShapeBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *BoxShape) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BoxShape) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BoxShape) Length() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *BoxShape) MutateLength(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func (rcv *BoxShape) Width() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *BoxShape) MutateWidth(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *BoxShape) Height() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *BoxShape) MutateHeight(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func BoxShapeStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func BoxShapeAddLength(builder *flatbuffers.Builder, length float32) {
	builder.PrependFloat32Slot(0, length, 0.0)
}
func BoxShapeAddWidth(builder *flatbuffers.Builder, width float32) {
	builder.PrependFloat32Slot(1, width, 0.0)
}
func BoxShapeAddHeight(builder *flatbuffers.Builder, height float32) {
	builder.PrependFloat32Slot(2, height, 0.0)
}
func BoxShapeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
