// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TouchT struct {
	PlayerName string `json:"player_name"`
	GameSeconds float32 `json:"game_seconds"`
	Location *Vector3T `json:"location"`
	Normal *Vector3T `json:"normal"`
	Team uint32 `json:"team"`
	PlayerIndex uint32 `json:"player_index"`
}

func (t *TouchT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	playerNameOffset := flatbuffers.UOffsetT(0)
	if t.PlayerName != "" {
		playerNameOffset = builder.CreateString(t.PlayerName)
	}
	TouchStart(builder)
	TouchAddPlayerName(builder, playerNameOffset)
	TouchAddGameSeconds(builder, t.GameSeconds)
	locationOffset := t.Location.Pack(builder)
	TouchAddLocation(builder, locationOffset)
	normalOffset := t.Normal.Pack(builder)
	TouchAddNormal(builder, normalOffset)
	TouchAddTeam(builder, t.Team)
	TouchAddPlayerIndex(builder, t.PlayerIndex)
	return TouchEnd(builder)
}

func (rcv *Touch) UnPackTo(t *TouchT) {
	t.PlayerName = string(rcv.PlayerName())
	t.GameSeconds = rcv.GameSeconds()
	t.Location = rcv.Location(nil).UnPack()
	t.Normal = rcv.Normal(nil).UnPack()
	t.Team = rcv.Team()
	t.PlayerIndex = rcv.PlayerIndex()
}

func (rcv *Touch) UnPack() *TouchT {
	if rcv == nil {
		return nil
	}
	t := &TouchT{}
	rcv.UnPackTo(t)
	return t
}

type Touch struct {
	_tab flatbuffers.Table
}

func GetRootAsTouch(buf []byte, offset flatbuffers.UOffsetT) *Touch {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Touch{}
	x.Init(buf, n+offset)
	return x
}

func FinishTouchBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsTouch(buf []byte, offset flatbuffers.UOffsetT) *Touch {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Touch{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedTouchBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Touch) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Touch) Table() flatbuffers.Table {
	return rcv._tab
}

/// The name of the player involved with the touch.
func (rcv *Touch) PlayerName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The name of the player involved with the touch.
/// Seconds that had elapsed in the game when the touch occurred.
func (rcv *Touch) GameSeconds() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Seconds that had elapsed in the game when the touch occurred.
func (rcv *Touch) MutateGameSeconds(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

/// The point of contact for the touch.
func (rcv *Touch) Location(obj *Vector3) *Vector3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vector3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The point of contact for the touch.
/// The direction of the touch.
func (rcv *Touch) Normal(obj *Vector3) *Vector3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vector3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The direction of the touch.
/// The Team which the touch belongs to, 0 for blue 1 for orange.
func (rcv *Touch) Team() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// The Team which the touch belongs to, 0 for blue 1 for orange.
func (rcv *Touch) MutateTeam(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

/// The index of the player involved with the touch.
func (rcv *Touch) PlayerIndex() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// The index of the player involved with the touch.
func (rcv *Touch) MutatePlayerIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func TouchStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func TouchAddPlayerName(builder *flatbuffers.Builder, playerName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(playerName), 0)
}
func TouchAddGameSeconds(builder *flatbuffers.Builder, gameSeconds float32) {
	builder.PrependFloat32Slot(1, gameSeconds, 0.0)
}
func TouchAddLocation(builder *flatbuffers.Builder, location flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(location), 0)
}
func TouchAddNormal(builder *flatbuffers.Builder, normal flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(normal), 0)
}
func TouchAddTeam(builder *flatbuffers.Builder, team uint32) {
	builder.PrependUint32Slot(4, team, 0)
}
func TouchAddPlayerIndex(builder *flatbuffers.Builder, playerIndex uint32) {
	builder.PrependUint32Slot(5, playerIndex, 0)
}
func TouchEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
