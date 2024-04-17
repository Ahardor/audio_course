package messages

import (
	"fmt"
	"iotvisual/processor/internal/domain/device"
	"iotvisual/processor/internal/domain/session"
)

// MessageSoundOutput выходное сообщение обработчика.
type MessageSoundOutput struct {
	// ID устройства (датчика).
	Device device.ID `json:"device"`
	// ID сессии.
	SessionUUID session.UUID `json:"session_uuid"`
	// Название мелодии.
	Melody string `json:"melody"`
	// Ожидаемая нота.
	ExpectedNote string `json:"expected_note"`
	// Ожидаемая частота.
	ExpectedFrequency float64 `json:"expected_frequency"`
	// Фактическая частота.
	ActualFrequency float64 `json:"actual_frequency"`
	// Фактическая нота.
	ActualNote string `json:"actual_note"`
	// Ожидаемая длительность в секундах.
	ExpectedLengthSeconds float64 `json:"expected_length_seconds"`
	// Фактическая длительность в секундах.
	ActualLengthSeconds float64 `json:"actual_length_seconds"`
}

// String реализация интерфейса Stringer.
func (m MessageSoundOutput) String() string {
	s := fmt.Sprintf("\nDevice: %v\nSessionUUID: %v\nMelody: %v\n\tExpectedNote: %v\n\tActualNote: %v\n\tExpectedDuration: %v\n\tActualDuration: %v\n",
		m.Device, m.SessionUUID, m.Melody, m.ExpectedNote, m.ActualNote, m.ExpectedLengthSeconds, m.ActualLengthSeconds,
	)
	return s
}
