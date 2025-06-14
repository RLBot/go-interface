package gointerface

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"os"

	flat "github.com/RLBot/go-interface/flat"
	flatbuffers "github.com/google/flatbuffers/go"
)

type RLBotConnection struct {
	conn    net.Conn
	builder flatbuffers.Builder
}

func Connect(addr string) (RLBotConnection, error) {
	conn, err := net.Dial("tcp", addr)
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.SetNoDelay(true)
	}
	return RLBotConnection{conn, *flatbuffers.NewBuilder(65536)}, err
}

func (conn RLBotConnection) Initialize(default_agent_id string, wants_ball_predictions bool, wants_comms bool) (*flat.MatchConfigurationT, *flat.FieldInfoT, *flat.ControllableTeamInfoT, error) {
	agent_id := os.Getenv("RLBOT_AGENT_ID")
	if agent_id == "" {
		agent_id = default_agent_id
	}

	err := conn.SendPacket(&flat.ConnectionSettingsT{
		AgentId:              agent_id,
		WantsBallPredictions: wants_ball_predictions,
		WantsComms:           wants_comms,
		CloseBetweenMatches:  true,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	var match_config *flat.MatchConfigurationT
	var field_info *flat.FieldInfoT
	var controllables *flat.ControllableTeamInfoT

	for {
		packet, err := conn.RecvPacket()
		if err != nil {
			return nil, nil, nil, err
		}

		switch packet := packet.Value.(type) {
		case *flat.MatchConfigurationT:
			match_config = packet
		case *flat.FieldInfoT:
			field_info = packet
		case *flat.ControllableTeamInfoT:
			controllables = packet
		}

		if match_config != nil && field_info != nil && controllables != nil {
			break
		}
	}

	err = conn.SendPacket(&flat.InitCompleteT{})
	if err != nil {
		return nil, nil, nil, err
	}

	return match_config, field_info, controllables, nil
}

func (conn RLBotConnection) SendPacket(msg any) error {
	var packetType flat.InterfaceMessage

	switch msg.(type) {
	case *flat.StartCommandT:
		packetType = flat.InterfaceMessageStartCommand
	case *flat.MatchConfigurationT:
		packetType = flat.InterfaceMessageMatchConfiguration
	case *flat.PlayerInputT:
		packetType = flat.InterfaceMessagePlayerInput
	case *flat.DesiredGameStateT:
		packetType = flat.InterfaceMessageDesiredGameState
	case *flat.RenderGroupT:
		packetType = flat.InterfaceMessageRenderGroup
	case *flat.RemoveRenderGroupT:
		packetType = flat.InterfaceMessageRemoveRenderGroup
	case *flat.MatchCommT:
		packetType = flat.InterfaceMessageMatchComm
	case *flat.ConnectionSettingsT:
		packetType = flat.InterfaceMessageConnectionSettings
	case *flat.StopCommandT:
		packetType = flat.InterfaceMessageStopCommand
	case *flat.SetLoadoutT:
		packetType = flat.InterfaceMessageSetLoadout
	case *flat.InitCompleteT:
		packetType = flat.InterfaceMessageInitComplete
	case *flat.RenderingStatusT:
		packetType = flat.InterfaceMessageRenderingStatus
	default:
		return errors.New("unsupported packet type")
	}

	packet := flat.InterfacePacketT{
		Message: &flat.InterfaceMessageT{
			Type:  packetType,
			Value: msg,
		},
	}

	conn.builder.Reset()
	conn.builder.Finish(packet.Pack(&conn.builder))
	packetPayload := conn.builder.FinishedBytes()

	packetLen := len(packetPayload)
	finalBuf := make([]byte, 2)

	binary.BigEndian.PutUint16(finalBuf, uint16(packetLen))
	finalBuf = append(finalBuf, packetPayload...)

	_, err := conn.conn.Write(finalBuf)
	return err
}

func (conn RLBotConnection) RecvPacket() (*flat.CoreMessageT, error) {
	buffer := make([]byte, 2)

	// Read packetLen
	_, err := io.ReadFull(conn.conn, buffer)
	if err != nil {
		return nil, err
	}
	packetLen := binary.BigEndian.Uint16(buffer)

	// Read packetPayload
	buffer = make([]byte, packetLen)
	_, err = io.ReadFull(conn.conn, buffer)
	if err != nil {
		return nil, err
	}

	return flat.GetRootAsCorePacket(buffer, 0).UnPack().Message, nil
}
