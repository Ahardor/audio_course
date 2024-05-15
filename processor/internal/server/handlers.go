package server

import (
	"context"
	"encoding/json"
	"fmt"
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
		fmt.Println("NOTE: ", sound.Note)
		fmt.Println("OCTAVE: ", sound.Octave)
		output := messages.MessageSoundOutput{
			Device:                input.Device,
			Melody:                input.Melody,
			SessionUUID:           input.SessionUUID,
			ExpectedNote:          sound.Note.GetNote(),
			ActualNote:            note.GetNote(),
			ExpectedLengthSeconds: float64(sound.DurationMS) / 1000,
			ActualLengthSeconds:   float64(input.LengthMS) / 1000,
			ExpectedFrequency:     float64(s.noteTable.GetFrequency(sound.Note, sound.Octave)) / 100,
			ActualFrequency:       input.Frequency,
		}

		bytes, err := json.Marshal(output)
		if err != nil {
			s.Logger.Err(err).Stack().Ctx(ctx).Msgf("JSON marshal output: ")
			return
		}

		token := c.Publish("sound/note/record", 1, false, bytes)
		s.Logger.Debug().Msgf("Sending processed note to RT-database: %t, with error: %v", token.Wait(), token.Error())
		s.Logger.Debug().Msgf("Sent message: %s", output.String())
	}
}
