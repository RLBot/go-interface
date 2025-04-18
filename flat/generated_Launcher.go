// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

/// Possible to launch Rocket League.
type Launcher byte

const (
	LauncherSteam    Launcher = 0
	LauncherEpic     Launcher = 1
	/// E.g. if you use Legendary.
	/// The game path is specified in the MatchConfiguration.
	LauncherCustom   Launcher = 2
	LauncherNoLaunch Launcher = 3
)

var EnumNamesLauncher = map[Launcher]string{
	LauncherSteam:    "Steam",
	LauncherEpic:     "Epic",
	LauncherCustom:   "Custom",
	LauncherNoLaunch: "NoLaunch",
}

var EnumValuesLauncher = map[string]Launcher{
	"Steam":    LauncherSteam,
	"Epic":     LauncherEpic,
	"Custom":   LauncherCustom,
	"NoLaunch": LauncherNoLaunch,
}

func (v Launcher) String() string {
	if s, ok := EnumNamesLauncher[v]; ok {
		return s
	}
	return "Launcher(" + strconv.FormatInt(int64(v), 10) + ")"
}
