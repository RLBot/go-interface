// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Specification for 'painted' items. See https://wiki.rlbot.org/botmaking/bot-customization/
type LoadoutPaintT struct {
	CarPaintId uint32 `json:"car_paint_id"`
	DecalPaintId uint32 `json:"decal_paint_id"`
	WheelsPaintId uint32 `json:"wheels_paint_id"`
	BoostPaintId uint32 `json:"boost_paint_id"`
	AntennaPaintId uint32 `json:"antenna_paint_id"`
	HatPaintId uint32 `json:"hat_paint_id"`
	TrailsPaintId uint32 `json:"trails_paint_id"`
	GoalExplosionPaintId uint32 `json:"goal_explosion_paint_id"`
}

func (t *LoadoutPaintT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	LoadoutPaintStart(builder)
	LoadoutPaintAddCarPaintId(builder, t.CarPaintId)
	LoadoutPaintAddDecalPaintId(builder, t.DecalPaintId)
	LoadoutPaintAddWheelsPaintId(builder, t.WheelsPaintId)
	LoadoutPaintAddBoostPaintId(builder, t.BoostPaintId)
	LoadoutPaintAddAntennaPaintId(builder, t.AntennaPaintId)
	LoadoutPaintAddHatPaintId(builder, t.HatPaintId)
	LoadoutPaintAddTrailsPaintId(builder, t.TrailsPaintId)
	LoadoutPaintAddGoalExplosionPaintId(builder, t.GoalExplosionPaintId)
	return LoadoutPaintEnd(builder)
}

func (rcv *LoadoutPaint) UnPackTo(t *LoadoutPaintT) {
	t.CarPaintId = rcv.CarPaintId()
	t.DecalPaintId = rcv.DecalPaintId()
	t.WheelsPaintId = rcv.WheelsPaintId()
	t.BoostPaintId = rcv.BoostPaintId()
	t.AntennaPaintId = rcv.AntennaPaintId()
	t.HatPaintId = rcv.HatPaintId()
	t.TrailsPaintId = rcv.TrailsPaintId()
	t.GoalExplosionPaintId = rcv.GoalExplosionPaintId()
}

func (rcv *LoadoutPaint) UnPack() *LoadoutPaintT {
	if rcv == nil {
		return nil
	}
	t := &LoadoutPaintT{}
	rcv.UnPackTo(t)
	return t
}

type LoadoutPaint struct {
	_tab flatbuffers.Table
}

func GetRootAsLoadoutPaint(buf []byte, offset flatbuffers.UOffsetT) *LoadoutPaint {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LoadoutPaint{}
	x.Init(buf, n+offset)
	return x
}

func FinishLoadoutPaintBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsLoadoutPaint(buf []byte, offset flatbuffers.UOffsetT) *LoadoutPaint {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LoadoutPaint{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedLoadoutPaintBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *LoadoutPaint) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LoadoutPaint) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LoadoutPaint) CarPaintId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadoutPaint) MutateCarPaintId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *LoadoutPaint) DecalPaintId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadoutPaint) MutateDecalPaintId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *LoadoutPaint) WheelsPaintId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadoutPaint) MutateWheelsPaintId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *LoadoutPaint) BoostPaintId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadoutPaint) MutateBoostPaintId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *LoadoutPaint) AntennaPaintId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadoutPaint) MutateAntennaPaintId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *LoadoutPaint) HatPaintId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadoutPaint) MutateHatPaintId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func (rcv *LoadoutPaint) TrailsPaintId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadoutPaint) MutateTrailsPaintId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(16, n)
}

func (rcv *LoadoutPaint) GoalExplosionPaintId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadoutPaint) MutateGoalExplosionPaintId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(18, n)
}

func LoadoutPaintStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func LoadoutPaintAddCarPaintId(builder *flatbuffers.Builder, carPaintId uint32) {
	builder.PrependUint32Slot(0, carPaintId, 0)
}
func LoadoutPaintAddDecalPaintId(builder *flatbuffers.Builder, decalPaintId uint32) {
	builder.PrependUint32Slot(1, decalPaintId, 0)
}
func LoadoutPaintAddWheelsPaintId(builder *flatbuffers.Builder, wheelsPaintId uint32) {
	builder.PrependUint32Slot(2, wheelsPaintId, 0)
}
func LoadoutPaintAddBoostPaintId(builder *flatbuffers.Builder, boostPaintId uint32) {
	builder.PrependUint32Slot(3, boostPaintId, 0)
}
func LoadoutPaintAddAntennaPaintId(builder *flatbuffers.Builder, antennaPaintId uint32) {
	builder.PrependUint32Slot(4, antennaPaintId, 0)
}
func LoadoutPaintAddHatPaintId(builder *flatbuffers.Builder, hatPaintId uint32) {
	builder.PrependUint32Slot(5, hatPaintId, 0)
}
func LoadoutPaintAddTrailsPaintId(builder *flatbuffers.Builder, trailsPaintId uint32) {
	builder.PrependUint32Slot(6, trailsPaintId, 0)
}
func LoadoutPaintAddGoalExplosionPaintId(builder *flatbuffers.Builder, goalExplosionPaintId uint32) {
	builder.PrependUint32Slot(7, goalExplosionPaintId, 0)
}
func LoadoutPaintEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
