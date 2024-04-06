package server

import (
	"context"
	"encoding/json"
	"iotvisual/processor/internal/domain/melody"
	"iotvisual/processor/internal/domain/messages"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (s *Server) MelodyEventHandler(ctx context.Context) mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		if !c.IsConnected() {
			s.Logger.Error().Stack().Ctx(ctx).Msg("MQTT client is not connected, cannot handle message")
			return
		}

		sound := messages.MessageSoundInput{}
		if err := json.Unmarshal(m.Payload(), &sound); err != nil {
			s.Logger.Err(err).Stack().Ctx(ctx).Msgf("JSON unmarshal sound: ")
			return
		}

		if _, ok := s.cache.Load(melody.ID(sound.Melody)); !ok {
			m, err := s.Queries.GetMelody(ctx, melody.ID(sound.Melody))
			if err != nil {
				s.Logger.Err(err).Stack().Ctx(ctx).Msgf("queries GetMelody: ")
				return
			}
			s.cache.Store(m.ID, m)
		}
	}
}
