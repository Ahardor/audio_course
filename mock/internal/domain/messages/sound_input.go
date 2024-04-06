package messages

import (
	"iotvisual/mock/internal/domain/device"
	"iotvisual/mock/internal/domain/session"
)

// MessageSoundInput событие получение звука из датчика.
type MessageSoundInput struct {
	// ID устройства (датчика).
	Device device.ID `yaml:"device"`
	// ID сессии.
	SessionUUID session.UUID `yaml:"session_uuid"`
	// Название мелодии.
	Melody string `yaml:"melody"`
	// Частота звука.
	Frequency float64 `yaml:"frequency"`
	// Длительность в миллисекундах.
	LengthMS int64 `yaml:"length_ms"`
}
