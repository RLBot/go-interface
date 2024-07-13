// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SphereShapeT struct {
	Diameter float32 `json:"diameter"`
}

func (t *SphereShapeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	SphereShapeStart(builder)
	SphereShapeAddDiameter(builder, t.Diameter)
	return SphereShapeEnd(builder)
}

func (rcv *SphereShape) UnPackTo(t *SphereShapeT) {
	t.Diameter = rcv.Diameter()
}

func (rcv *SphereShape) UnPack() *SphereShapeT {
	if rcv == nil {
		return nil
	}
	t := &SphereShapeT{}
	rcv.UnPackTo(t)
	return t
}

type SphereShape struct {
	_tab flatbuffers.Table
}

func GetRootAsSphereShape(buf []byte, offset flatbuffers.UOffsetT) *SphereShape {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SphereShape{}
	x.Init(buf, n+offset)
	return x
}

func FinishSphereShapeBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsSphereShape(buf []byte, offset flatbuffers.UOffsetT) *SphereShape {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SphereShape{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedSphereShapeBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *SphereShape) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SphereShape) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SphereShape) Diameter() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *SphereShape) MutateDiameter(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func SphereShapeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SphereShapeAddDiameter(builder *flatbuffers.Builder, diameter float32) {
	builder.PrependFloat32Slot(0, diameter, 0.0)
}
func SphereShapeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
