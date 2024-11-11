// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PhysicsT struct {
	Location *Vector3T `json:"location"`
	Rotation *RotatorT `json:"rotation"`
	Velocity *Vector3T `json:"velocity"`
	AngularVelocity *Vector3T `json:"angular_velocity"`
}

func (t *PhysicsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreatePhysics(builder, t.Location.X, t.Location.Y, t.Location.Z, t.Rotation.Pitch, t.Rotation.Yaw, t.Rotation.Roll, t.Velocity.X, t.Velocity.Y, t.Velocity.Z, t.AngularVelocity.X, t.AngularVelocity.Y, t.AngularVelocity.Z)
}
func (rcv *Physics) UnPackTo(t *PhysicsT) {
	t.Location = rcv.Location(nil).UnPack()
	t.Rotation = rcv.Rotation(nil).UnPack()
	t.Velocity = rcv.Velocity(nil).UnPack()
	t.AngularVelocity = rcv.AngularVelocity(nil).UnPack()
}

func (rcv *Physics) UnPack() *PhysicsT {
	if rcv == nil {
		return nil
	}
	t := &PhysicsT{}
	rcv.UnPackTo(t)
	return t
}

type Physics struct {
	_tab flatbuffers.Struct
}

func (rcv *Physics) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Physics) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Physics) Location(obj *Vector3) *Vector3 {
	if obj == nil {
		obj = new(Vector3)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+0)
	return obj
}
func (rcv *Physics) Rotation(obj *Rotator) *Rotator {
	if obj == nil {
		obj = new(Rotator)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+12)
	return obj
}
func (rcv *Physics) Velocity(obj *Vector3) *Vector3 {
	if obj == nil {
		obj = new(Vector3)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+24)
	return obj
}
func (rcv *Physics) AngularVelocity(obj *Vector3) *Vector3 {
	if obj == nil {
		obj = new(Vector3)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+36)
	return obj
}

func CreatePhysics(builder *flatbuffers.Builder, location_x float32, location_y float32, location_z float32, rotation_pitch float32, rotation_yaw float32, rotation_roll float32, velocity_x float32, velocity_y float32, velocity_z float32, angular_velocity_x float32, angular_velocity_y float32, angular_velocity_z float32) flatbuffers.UOffsetT {
	builder.Prep(4, 48)
	builder.Prep(4, 12)
	builder.PrependFloat32(angular_velocity_z)
	builder.PrependFloat32(angular_velocity_y)
	builder.PrependFloat32(angular_velocity_x)
	builder.Prep(4, 12)
	builder.PrependFloat32(velocity_z)
	builder.PrependFloat32(velocity_y)
	builder.PrependFloat32(velocity_x)
	builder.Prep(4, 12)
	builder.PrependFloat32(rotation_roll)
	builder.PrependFloat32(rotation_yaw)
	builder.PrependFloat32(rotation_pitch)
	builder.Prep(4, 12)
	builder.PrependFloat32(location_z)
	builder.PrependFloat32(location_y)
	builder.PrependFloat32(location_x)
	return builder.Offset()
}
