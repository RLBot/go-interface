// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Rect3DT struct {
	Anchor *RenderAnchorT `json:"anchor"`
	Width float32 `json:"width"`
	Height float32 `json:"height"`
	Color *ColorT `json:"color"`
}

func (t *Rect3DT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	anchorOffset := t.Anchor.Pack(builder)
	Rect3DStart(builder)
	Rect3DAddAnchor(builder, anchorOffset)
	Rect3DAddWidth(builder, t.Width)
	Rect3DAddHeight(builder, t.Height)
	colorOffset := t.Color.Pack(builder)
	Rect3DAddColor(builder, colorOffset)
	return Rect3DEnd(builder)
}

func (rcv *Rect3D) UnPackTo(t *Rect3DT) {
	t.Anchor = rcv.Anchor(nil).UnPack()
	t.Width = rcv.Width()
	t.Height = rcv.Height()
	t.Color = rcv.Color(nil).UnPack()
}

func (rcv *Rect3D) UnPack() *Rect3DT {
	if rcv == nil {
		return nil
	}
	t := &Rect3DT{}
	rcv.UnPackTo(t)
	return t
}

type Rect3D struct {
	_tab flatbuffers.Table
}

func GetRootAsRect3D(buf []byte, offset flatbuffers.UOffsetT) *Rect3D {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Rect3D{}
	x.Init(buf, n+offset)
	return x
}

func FinishRect3DBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsRect3D(buf []byte, offset flatbuffers.UOffsetT) *Rect3D {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Rect3D{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedRect3DBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Rect3D) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Rect3D) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Rect3D) Anchor(obj *RenderAnchor) *RenderAnchor {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(RenderAnchor)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Screen-space size such that width=0.1 is 10% of window width.
func (rcv *Rect3D) Width() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Screen-space size such that width=0.1 is 10% of window width.
func (rcv *Rect3D) MutateWidth(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

/// Screen-space size such that height=0.1 is 10% of window height.
func (rcv *Rect3D) Height() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Screen-space size such that height=0.1 is 10% of window height.
func (rcv *Rect3D) MutateHeight(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *Rect3D) Color(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Color)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func Rect3DStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func Rect3DAddAnchor(builder *flatbuffers.Builder, anchor flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(anchor), 0)
}
func Rect3DAddWidth(builder *flatbuffers.Builder, width float32) {
	builder.PrependFloat32Slot(1, width, 0.0)
}
func Rect3DAddHeight(builder *flatbuffers.Builder, height float32) {
	builder.PrependFloat32Slot(2, height, 0.0)
}
func Rect3DAddColor(builder *flatbuffers.Builder, color flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(color), 0)
}
func Rect3DEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
