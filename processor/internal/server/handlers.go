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

		sound, status := s.cache.LoadNthNote(melody.ID(input.Melody), input.SerialNumber)
		switch status {
		case cacher.NoteLoadStatusMelodyDoesNotExist:
			m, err := s.Queries.GetMelody(ctx, melody.ID(input.Melody))
			if err != nil {
				s.Logger.Err(err).Stack().Ctx(ctx).Msgf("queries GetMelody: ")
				return
			}
			if len(m.Sounds) > input.SerialNumber {
				s.cache.Store(m.ID, m)
				sound = m.Sounds[input.SerialNumber]
				break
			}
			fallthrough
		case cacher.NoteLoadStatusNoteIndexOutOfRange:
			s.Logger.Warn().Stack().Ctx(ctx).Msg("Note index out of bounds, recording aborted")
			return
		case cacher.NoteLoadStatusOK:
			break
		}

		note := s.noteTable.FindNote(input.Frequency)
		// TODO: добавить погрешность к длительности ноты.
		// Поправить фактическую длительность нот (см. беседу проекта в ТГ).
		output := messages.MessageSoundOutput{
			Device:           input.Device,
			Melody:           input.Melody,
			SessionUUID:      input.SessionUUID,
			ExpectedNote:     sound.Note.GetNote(),
			ActualNote:       note.GetNote(),
			ExpectedLengthMS: sound.DurationMS,
			ActualLengthMS:   input.LengthMS,
		}

		bytes, err := json.Marshal(output)
		if err != nil {
			s.Logger.Err(err).Stack().Ctx(ctx).Msgf("JSON marshal output: ")
			return
		}

		token := c.Publish("sound/note/record", 0, false, bytes)
		s.Logger.Debug().Msgf("Sending processed note to RT-database: %t, with error: %v", token.Wait(), token.Error())
	}
}
