// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A collection of values shown on the scoreboard (and a few more).
type ScoreInfoT struct {
	Score uint32 `json:"score"`
	Goals uint32 `json:"goals"`
	OwnGoals uint32 `json:"own_goals"`
	Assists uint32 `json:"assists"`
	Saves uint32 `json:"saves"`
	Shots uint32 `json:"shots"`
	Demolitions uint32 `json:"demolitions"`
}

func (t *ScoreInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateScoreInfo(builder, t.Score, t.Goals, t.OwnGoals, t.Assists, t.Saves, t.Shots, t.Demolitions)
}
func (rcv *ScoreInfo) UnPackTo(t *ScoreInfoT) {
	t.Score = rcv.Score()
	t.Goals = rcv.Goals()
	t.OwnGoals = rcv.OwnGoals()
	t.Assists = rcv.Assists()
	t.Saves = rcv.Saves()
	t.Shots = rcv.Shots()
	t.Demolitions = rcv.Demolitions()
}

func (rcv *ScoreInfo) UnPack() *ScoreInfoT {
	if rcv == nil {
		return nil
	}
	t := &ScoreInfoT{}
	rcv.UnPackTo(t)
	return t
}

type ScoreInfo struct {
	_tab flatbuffers.Struct
}

func (rcv *ScoreInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ScoreInfo) Table() flatbuffers.Table {
	return rcv._tab.Table
}

/// The accumulated score, roughly indicating how well a player performs.
func (rcv *ScoreInfo) Score() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
/// The accumulated score, roughly indicating how well a player performs.
func (rcv *ScoreInfo) MutateScore(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

/// Number of goals scored.
func (rcv *ScoreInfo) Goals() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
/// Number of goals scored.
func (rcv *ScoreInfo) MutateGoals(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

/// Number of own-goals scored.
func (rcv *ScoreInfo) OwnGoals() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
/// Number of own-goals scored.
func (rcv *ScoreInfo) MutateOwnGoals(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

/// Number of goals assisted.
func (rcv *ScoreInfo) Assists() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(12))
}
/// Number of goals assisted.
func (rcv *ScoreInfo) MutateAssists(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(12), n)
}

/// Number of shots saved.
func (rcv *ScoreInfo) Saves() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(16))
}
/// Number of shots saved.
func (rcv *ScoreInfo) MutateSaves(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(16), n)
}

/// Number of shots on opponent goal.
func (rcv *ScoreInfo) Shots() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(20))
}
/// Number of shots on opponent goal.
func (rcv *ScoreInfo) MutateShots(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(20), n)
}

/// Number of demolitions made.
func (rcv *ScoreInfo) Demolitions() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(24))
}
/// Number of demolitions made.
func (rcv *ScoreInfo) MutateDemolitions(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(24), n)
}

func CreateScoreInfo(builder *flatbuffers.Builder, score uint32, goals uint32, ownGoals uint32, assists uint32, saves uint32, shots uint32, demolitions uint32) flatbuffers.UOffsetT {
	builder.Prep(4, 28)
	builder.PrependUint32(demolitions)
	builder.PrependUint32(shots)
	builder.PrependUint32(saves)
	builder.PrependUint32(assists)
	builder.PrependUint32(ownGoals)
	builder.PrependUint32(goals)
	builder.PrependUint32(score)
	return builder.Offset()
}
