package messages

import (
	"iotvisual/processor/internal/domain/device"
	"iotvisual/processor/internal/domain/session"
)

// MessageSoundInput событие получение звука из датчика.
type MessageSoundInput struct {
	// ID устройства (датчика).
	Device device.ID `yaml:"device"`
	// ID сессии.
	SessionUUID session.UUID `yaml:"session_uuid"`
	// Частота звука.
	Frequency float64 `yaml:"frequency"`
	// Длительность в миллисекундах.
	LengthMS int64 `yaml:"length_ms"`
}
