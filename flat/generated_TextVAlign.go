// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

/// Vertical text alignment.
type TextVAlign byte

const (
	TextVAlignTop    TextVAlign = 0
	TextVAlignCenter TextVAlign = 1
	TextVAlignBottom TextVAlign = 2
)

var EnumNamesTextVAlign = map[TextVAlign]string{
	TextVAlignTop:    "Top",
	TextVAlignCenter: "Center",
	TextVAlignBottom: "Bottom",
}

var EnumValuesTextVAlign = map[string]TextVAlign{
	"Top":    TextVAlignTop,
	"Center": TextVAlignCenter,
	"Bottom": TextVAlignBottom,
}

func (v TextVAlign) String() string {
	if s, ok := EnumNamesTextVAlign[v]; ok {
		return s
	}
	return "TextVAlign(" + strconv.FormatInt(int64(v), 10) + ")"
}
