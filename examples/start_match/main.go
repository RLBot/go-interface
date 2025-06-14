package main

import (
	"os"

	RLBot "github.com/RLBot/go-interface"
	RLBotFlat "github.com/RLBot/go-interface/flat"
)

func main() {
	println("Connecting...")
	conn, err := RLBot.Connect("127.0.0.1:23234")
	if err != nil {
		panic("could not connect to rlbot core")
	}

	dir, err := os.Getwd()
	if err != nil {
		panic("could not get current working directory")
	}

	println("Sending MatchConfiguration packet...")
	conn.SendPacket(&RLBotFlat.MatchConfigurationT{
		PlayerConfigurations: []*RLBotFlat.PlayerConfigurationT{{
			Variety: &RLBotFlat.PlayerClassT{
				Type:  RLBotFlat.PlayerClassCustomBot,
				Value: &RLBotFlat.CustomBotT{
					RootDir:    dir + "/examples/atba",
					RunCommand: "go run main.go",
					AgentId:    "rlbot/go-example-bot",
					Name:       "Go Example",
					Loadout:    &RLBotFlat.PlayerLoadoutT{},
				},
			},
			Team:       0,
		}, {
			Variety: &RLBotFlat.PlayerClassT{
				Type:  RLBotFlat.PlayerClassHuman,
				Value: &RLBotFlat.HumanT{},
			},
			Team:    1,
		}},
		GameMode:   RLBotFlat.GameModeSoccar,
		GameMapUpk: "UtopiaStadium_P",
		Mutators: &RLBotFlat.MutatorSettingsT{
			MatchLength: RLBotFlat.MatchLengthMutatorUnlimited,
		},
		ExistingMatchBehavior: RLBotFlat.ExistingMatchBehaviorRestart,
		EnableRendering:       RLBotFlat.DebugRenderingAlwaysOff,
		AutoStartAgents:       true,
		WaitForAgents:         true,
	})
}
