// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Notification when the local player is spectating another player.
type PlayerSpectateT struct {
	PlayerIndex uint32 `json:"player_index"`
}

func (t *PlayerSpectateT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	PlayerSpectateStart(builder)
	PlayerSpectateAddPlayerIndex(builder, t.PlayerIndex)
	return PlayerSpectateEnd(builder)
}

func (rcv *PlayerSpectate) UnPackTo(t *PlayerSpectateT) {
	t.PlayerIndex = rcv.PlayerIndex()
}

func (rcv *PlayerSpectate) UnPack() *PlayerSpectateT {
	if rcv == nil {
		return nil
	}
	t := &PlayerSpectateT{}
	rcv.UnPackTo(t)
	return t
}

type PlayerSpectate struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayerSpectate(buf []byte, offset flatbuffers.UOffsetT) *PlayerSpectate {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PlayerSpectate{}
	x.Init(buf, n+offset)
	return x
}

func FinishPlayerSpectateBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPlayerSpectate(buf []byte, offset flatbuffers.UOffsetT) *PlayerSpectate {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PlayerSpectate{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPlayerSpectateBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *PlayerSpectate) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PlayerSpectate) Table() flatbuffers.Table {
	return rcv._tab
}

/// index of the player that is being spectated. Will be -1 if not spectating anyone.
func (rcv *PlayerSpectate) PlayerIndex() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// index of the player that is being spectated. Will be -1 if not spectating anyone.
func (rcv *PlayerSpectate) MutatePlayerIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func PlayerSpectateStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func PlayerSpectateAddPlayerIndex(builder *flatbuffers.Builder, playerIndex uint32) {
	builder.PrependUint32Slot(0, playerIndex, 0)
}
func PlayerSpectateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
