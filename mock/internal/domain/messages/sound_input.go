package messages

import (
	"iotvisual/mock/internal/domain/device"
	"iotvisual/mock/internal/domain/session"
)

// MessageSoundInput событие получение звука из датчика.
type MessageSoundInput struct {
	// ID устройства (датчика).
	Device device.ID `yaml:"device" json:"device"`
	// ID сессии.
	SessionUUID session.UUID `yaml:"session_uuid" json:"session_uuid"`
	// Название мелодии.
	Melody string `yaml:"melody" json:"melody"`
	// Частота звука.
	Frequency float64 `yaml:"frequency" json:"frequency"`
	// Длительность в миллисекундах.
	LengthMS int64 `yaml:"length_ms" json:"length_ms"`
	// Порядковый номер в мелодии.
	SerialNumber int `yaml:"serial_number" json:"serial_number"`
}
