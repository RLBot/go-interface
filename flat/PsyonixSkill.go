// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import "strconv"

type PsyonixSkill byte

const (
	PsyonixSkillBeginner PsyonixSkill = 0
	PsyonixSkillRookie   PsyonixSkill = 1
	PsyonixSkillPro      PsyonixSkill = 2
	PsyonixSkillAllStar  PsyonixSkill = 3
)

var EnumNamesPsyonixSkill = map[PsyonixSkill]string{
	PsyonixSkillBeginner: "Beginner",
	PsyonixSkillRookie:   "Rookie",
	PsyonixSkillPro:      "Pro",
	PsyonixSkillAllStar:  "AllStar",
}

var EnumValuesPsyonixSkill = map[string]PsyonixSkill{
	"Beginner": PsyonixSkillBeginner,
	"Rookie":   PsyonixSkillRookie,
	"Pro":      PsyonixSkillPro,
	"AllStar":  PsyonixSkillAllStar,
}

func (v PsyonixSkill) String() string {
	if s, ok := EnumNamesPsyonixSkill[v]; ok {
		return s
	}
	return "PsyonixSkill(" + strconv.FormatInt(int64(v), 10) + ")"
}
