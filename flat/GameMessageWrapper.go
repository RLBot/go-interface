// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GameMessageWrapperT struct {
	Message *GameMessageT `json:"message"`
}

func (t *GameMessageWrapperT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	messageOffset := t.Message.Pack(builder)

	GameMessageWrapperStart(builder)
	if t.Message != nil {
		GameMessageWrapperAddMessageType(builder, t.Message.Type)
	}
	GameMessageWrapperAddMessage(builder, messageOffset)
	return GameMessageWrapperEnd(builder)
}

func (rcv *GameMessageWrapper) UnPackTo(t *GameMessageWrapperT) {
	messageTable := flatbuffers.Table{}
	if rcv.Message(&messageTable) {
		t.Message = rcv.MessageType().UnPack(messageTable)
	}
}

func (rcv *GameMessageWrapper) UnPack() *GameMessageWrapperT {
	if rcv == nil {
		return nil
	}
	t := &GameMessageWrapperT{}
	rcv.UnPackTo(t)
	return t
}

type GameMessageWrapper struct {
	_tab flatbuffers.Table
}

func GetRootAsGameMessageWrapper(buf []byte, offset flatbuffers.UOffsetT) *GameMessageWrapper {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GameMessageWrapper{}
	x.Init(buf, n+offset)
	return x
}

func FinishGameMessageWrapperBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsGameMessageWrapper(buf []byte, offset flatbuffers.UOffsetT) *GameMessageWrapper {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GameMessageWrapper{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedGameMessageWrapperBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *GameMessageWrapper) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GameMessageWrapper) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GameMessageWrapper) MessageType() GameMessage {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return GameMessage(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *GameMessageWrapper) MutateMessageType(n GameMessage) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *GameMessageWrapper) Message(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func GameMessageWrapperStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func GameMessageWrapperAddMessageType(builder *flatbuffers.Builder, messageType GameMessage) {
	builder.PrependByteSlot(0, byte(messageType), 0)
}
func GameMessageWrapperAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(message), 0)
}
func GameMessageWrapperEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
