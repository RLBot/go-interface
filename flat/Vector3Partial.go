// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Vector3PartialT struct {
	X *FloatT `json:"x"`
	Y *FloatT `json:"y"`
	Z *FloatT `json:"z"`
}

func (t *Vector3PartialT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	Vector3PartialStart(builder)
	xOffset := t.X.Pack(builder)
	Vector3PartialAddX(builder, xOffset)
	yOffset := t.Y.Pack(builder)
	Vector3PartialAddY(builder, yOffset)
	zOffset := t.Z.Pack(builder)
	Vector3PartialAddZ(builder, zOffset)
	return Vector3PartialEnd(builder)
}

func (rcv *Vector3Partial) UnPackTo(t *Vector3PartialT) {
	t.X = rcv.X(nil).UnPack()
	t.Y = rcv.Y(nil).UnPack()
	t.Z = rcv.Z(nil).UnPack()
}

func (rcv *Vector3Partial) UnPack() *Vector3PartialT {
	if rcv == nil {
		return nil
	}
	t := &Vector3PartialT{}
	rcv.UnPackTo(t)
	return t
}

type Vector3Partial struct {
	_tab flatbuffers.Table
}

func GetRootAsVector3Partial(buf []byte, offset flatbuffers.UOffsetT) *Vector3Partial {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Vector3Partial{}
	x.Init(buf, n+offset)
	return x
}

func FinishVector3PartialBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsVector3Partial(buf []byte, offset flatbuffers.UOffsetT) *Vector3Partial {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Vector3Partial{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedVector3PartialBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Vector3Partial) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vector3Partial) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Vector3Partial) X(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Float)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Vector3Partial) Y(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Float)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Vector3Partial) Z(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Float)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func Vector3PartialStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func Vector3PartialAddX(builder *flatbuffers.Builder, x flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(x), 0)
}
func Vector3PartialAddY(builder *flatbuffers.Builder, y flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(y), 0)
}
func Vector3PartialAddZ(builder *flatbuffers.Builder, z flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(z), 0)
}
func Vector3PartialEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
