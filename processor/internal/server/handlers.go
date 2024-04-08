package server

import (
	"context"
	"encoding/json"
	"iotvisual/processor/internal/domain/melody"
	"iotvisual/processor/internal/domain/messages"
	"iotvisual/processor/internal/pkg/cacher"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (s *Server) MelodyEventHandler(ctx context.Context) mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		if !c.IsConnected() {
			s.Logger.Error().Stack().Ctx(ctx).Msg("MQTT client is not connected, cannot handle message")
			return
		}

		input := messages.MessageSoundInput{}
		if err := json.Unmarshal(m.Payload(), &input); err != nil {
			s.Logger.Err(err).Stack().Ctx(ctx).Msgf("JSON unmarshal input: ")
			return
		}

		// TODO: написать обработку индексов нот.
		// Можно передавать с сообщением, можно создать мапу, в которой хранить:
		// Key: device, Value: {Melody, Session, Index}.
		sound, status := s.cache.LoadNthNote(melody.ID(input.Melody), 0)
		switch status {
		case cacher.NoteLoadStatusMelodyDoesNotExist:
			m, err := s.Queries.GetMelody(ctx, melody.ID(input.Melody))
			if err != nil {
				s.Logger.Err(err).Stack().Ctx(ctx).Msgf("queries GetMelody: ")
				return
			}
			s.cache.Store(m.ID, m)
			// If index out of range : fallthrough ???
		case cacher.NoteLoadStatusNoteIndexOutOfRange:
			s.Logger.Warn().Stack().Ctx(ctx).Msg("Note index out of bounds, recording aborted")
			return
		case cacher.NoteLoadStatusOK:
			break
		}
	}
}
