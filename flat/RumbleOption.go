// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

type RumbleOption byte

const (
	RumbleOptionNo_Rumble         RumbleOption = 0
	RumbleOptionDefault           RumbleOption = 1
	RumbleOptionSlow              RumbleOption = 2
	RumbleOptionCivilized         RumbleOption = 3
	RumbleOptionDestruction_Derby RumbleOption = 4
	RumbleOptionSpring_Loaded     RumbleOption = 5
	RumbleOptionSpikes_Only       RumbleOption = 6
	RumbleOptionSpike_Rush        RumbleOption = 7
	RumbleOptionHaunted_Ball_Beam RumbleOption = 8
	RumbleOptionTactical          RumbleOption = 9
)

var EnumNamesRumbleOption = map[RumbleOption]string{
	RumbleOptionNo_Rumble:         "No_Rumble",
	RumbleOptionDefault:           "Default",
	RumbleOptionSlow:              "Slow",
	RumbleOptionCivilized:         "Civilized",
	RumbleOptionDestruction_Derby: "Destruction_Derby",
	RumbleOptionSpring_Loaded:     "Spring_Loaded",
	RumbleOptionSpikes_Only:       "Spikes_Only",
	RumbleOptionSpike_Rush:        "Spike_Rush",
	RumbleOptionHaunted_Ball_Beam: "Haunted_Ball_Beam",
	RumbleOptionTactical:          "Tactical",
}

var EnumValuesRumbleOption = map[string]RumbleOption{
	"No_Rumble":         RumbleOptionNo_Rumble,
	"Default":           RumbleOptionDefault,
	"Slow":              RumbleOptionSlow,
	"Civilized":         RumbleOptionCivilized,
	"Destruction_Derby": RumbleOptionDestruction_Derby,
	"Spring_Loaded":     RumbleOptionSpring_Loaded,
	"Spikes_Only":       RumbleOptionSpikes_Only,
	"Spike_Rush":        RumbleOptionSpike_Rush,
	"Haunted_Ball_Beam": RumbleOptionHaunted_Ball_Beam,
	"Tactical":          RumbleOptionTactical,
}

func (v RumbleOption) String() string {
	if s, ok := EnumNamesRumbleOption[v]; ok {
		return s
	}
	return "RumbleOption(" + strconv.FormatInt(int64(v), 10) + ")"
}
