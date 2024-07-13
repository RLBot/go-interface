// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

type GameSpeedOption byte

const (
	GameSpeedOptionDefault   GameSpeedOption = 0
	GameSpeedOptionSlo_Mo    GameSpeedOption = 1
	GameSpeedOptionTime_Warp GameSpeedOption = 2
)

var EnumNamesGameSpeedOption = map[GameSpeedOption]string{
	GameSpeedOptionDefault:   "Default",
	GameSpeedOptionSlo_Mo:    "Slo_Mo",
	GameSpeedOptionTime_Warp: "Time_Warp",
}

var EnumValuesGameSpeedOption = map[string]GameSpeedOption{
	"Default":   GameSpeedOptionDefault,
	"Slo_Mo":    GameSpeedOptionSlo_Mo,
	"Time_Warp": GameSpeedOptionTime_Warp,
}

func (v GameSpeedOption) String() string {
	if s, ok := EnumNamesGameSpeedOption[v]; ok {
		return s
	}
	return "GameSpeedOption(" + strconv.FormatInt(int64(v), 10) + ")"
}
