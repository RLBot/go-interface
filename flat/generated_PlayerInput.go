// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A player index and the controller state of that player.
/// Used to indicate what the player is doing this tick.
type PlayerInputT struct {
	PlayerIndex uint32 `json:"player_index"`
	ControllerState *ControllerStateT `json:"controller_state"`
}

func (t *PlayerInputT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	PlayerInputStart(builder)
	PlayerInputAddPlayerIndex(builder, t.PlayerIndex)
	controllerStateOffset := t.ControllerState.Pack(builder)
	PlayerInputAddControllerState(builder, controllerStateOffset)
	return PlayerInputEnd(builder)
}

func (rcv *PlayerInput) UnPackTo(t *PlayerInputT) {
	t.PlayerIndex = rcv.PlayerIndex()
	t.ControllerState = rcv.ControllerState(nil).UnPack()
}

func (rcv *PlayerInput) UnPack() *PlayerInputT {
	if rcv == nil {
		return nil
	}
	t := &PlayerInputT{}
	rcv.UnPackTo(t)
	return t
}

type PlayerInput struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayerInput(buf []byte, offset flatbuffers.UOffsetT) *PlayerInput {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PlayerInput{}
	x.Init(buf, n+offset)
	return x
}

func FinishPlayerInputBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPlayerInput(buf []byte, offset flatbuffers.UOffsetT) *PlayerInput {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PlayerInput{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPlayerInputBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *PlayerInput) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PlayerInput) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PlayerInput) PlayerIndex() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerInput) MutatePlayerIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *PlayerInput) ControllerState(obj *ControllerState) *ControllerState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ControllerState)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func PlayerInputStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PlayerInputAddPlayerIndex(builder *flatbuffers.Builder, playerIndex uint32) {
	builder.PrependUint32Slot(0, playerIndex, 0)
}
func PlayerInputAddControllerState(builder *flatbuffers.Builder, controllerState flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(controllerState), 0)
}
func PlayerInputEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
