// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PredictionSliceT struct {
	GameSeconds float32 `json:"game_seconds"`
	Physics *PhysicsT `json:"physics"`
}

func (t *PredictionSliceT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	physicsOffset := t.Physics.Pack(builder)
	PredictionSliceStart(builder)
	PredictionSliceAddGameSeconds(builder, t.GameSeconds)
	PredictionSliceAddPhysics(builder, physicsOffset)
	return PredictionSliceEnd(builder)
}

func (rcv *PredictionSlice) UnPackTo(t *PredictionSliceT) {
	t.GameSeconds = rcv.GameSeconds()
	t.Physics = rcv.Physics(nil).UnPack()
}

func (rcv *PredictionSlice) UnPack() *PredictionSliceT {
	if rcv == nil {
		return nil
	}
	t := &PredictionSliceT{}
	rcv.UnPackTo(t)
	return t
}

type PredictionSlice struct {
	_tab flatbuffers.Table
}

func GetRootAsPredictionSlice(buf []byte, offset flatbuffers.UOffsetT) *PredictionSlice {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PredictionSlice{}
	x.Init(buf, n+offset)
	return x
}

func FinishPredictionSliceBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPredictionSlice(buf []byte, offset flatbuffers.UOffsetT) *PredictionSlice {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PredictionSlice{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPredictionSliceBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *PredictionSlice) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PredictionSlice) Table() flatbuffers.Table {
	return rcv._tab
}

/// The moment in game time that this prediction corresponds to.
/// This corresponds to 'secondsElapsed' in the GameInfo table.
func (rcv *PredictionSlice) GameSeconds() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// The moment in game time that this prediction corresponds to.
/// This corresponds to 'secondsElapsed' in the GameInfo table.
func (rcv *PredictionSlice) MutateGameSeconds(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

/// The predicted location and motion of the object.
func (rcv *PredictionSlice) Physics(obj *Physics) *Physics {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Physics)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The predicted location and motion of the object.
func PredictionSliceStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PredictionSliceAddGameSeconds(builder *flatbuffers.Builder, gameSeconds float32) {
	builder.PrependFloat32Slot(0, gameSeconds, 0.0)
}
func PredictionSliceAddPhysics(builder *flatbuffers.Builder, physics flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(physics), 0)
}
func PredictionSliceEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
