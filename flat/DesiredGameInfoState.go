// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DesiredGameInfoStateT struct {
	WorldGravityZ *FloatT `json:"world_gravity_z"`
	GameSpeed *FloatT `json:"game_speed"`
	Paused *BoolT `json:"paused"`
	EndMatch *BoolT `json:"end_match"`
}

func (t *DesiredGameInfoStateT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	DesiredGameInfoStateStart(builder)
	worldGravityZOffset := t.WorldGravityZ.Pack(builder)
	DesiredGameInfoStateAddWorldGravityZ(builder, worldGravityZOffset)
	gameSpeedOffset := t.GameSpeed.Pack(builder)
	DesiredGameInfoStateAddGameSpeed(builder, gameSpeedOffset)
	pausedOffset := t.Paused.Pack(builder)
	DesiredGameInfoStateAddPaused(builder, pausedOffset)
	endMatchOffset := t.EndMatch.Pack(builder)
	DesiredGameInfoStateAddEndMatch(builder, endMatchOffset)
	return DesiredGameInfoStateEnd(builder)
}

func (rcv *DesiredGameInfoState) UnPackTo(t *DesiredGameInfoStateT) {
	t.WorldGravityZ = rcv.WorldGravityZ(nil).UnPack()
	t.GameSpeed = rcv.GameSpeed(nil).UnPack()
	t.Paused = rcv.Paused(nil).UnPack()
	t.EndMatch = rcv.EndMatch(nil).UnPack()
}

func (rcv *DesiredGameInfoState) UnPack() *DesiredGameInfoStateT {
	if rcv == nil {
		return nil
	}
	t := &DesiredGameInfoStateT{}
	rcv.UnPackTo(t)
	return t
}

type DesiredGameInfoState struct {
	_tab flatbuffers.Table
}

func GetRootAsDesiredGameInfoState(buf []byte, offset flatbuffers.UOffsetT) *DesiredGameInfoState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DesiredGameInfoState{}
	x.Init(buf, n+offset)
	return x
}

func FinishDesiredGameInfoStateBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsDesiredGameInfoState(buf []byte, offset flatbuffers.UOffsetT) *DesiredGameInfoState {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DesiredGameInfoState{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedDesiredGameInfoStateBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *DesiredGameInfoState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DesiredGameInfoState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DesiredGameInfoState) WorldGravityZ(obj *Float) *Float {
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

func (rcv *DesiredGameInfoState) GameSpeed(obj *Float) *Float {
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

func (rcv *DesiredGameInfoState) Paused(obj *Bool) *Bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Bool)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DesiredGameInfoState) EndMatch(obj *Bool) *Bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Bool)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DesiredGameInfoStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DesiredGameInfoStateAddWorldGravityZ(builder *flatbuffers.Builder, worldGravityZ flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(worldGravityZ), 0)
}
func DesiredGameInfoStateAddGameSpeed(builder *flatbuffers.Builder, gameSpeed flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(gameSpeed), 0)
}
func DesiredGameInfoStateAddPaused(builder *flatbuffers.Builder, paused flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(paused), 0)
}
func DesiredGameInfoStateAddEndMatch(builder *flatbuffers.Builder, endMatch flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(endMatch), 0)
}
func DesiredGameInfoStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
