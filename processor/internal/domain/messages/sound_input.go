package messages

import (
	"iotvisual/processor/internal/domain/device"
	"iotvisual/processor/internal/domain/session"
)

// MessageSoundInput событие получение звука из датчика.
type MessageSoundInput struct {
	// ID устройства (датчика).
	Device device.ID `json:"device"`
	// ID сессии.
	SessionUUID session.UUID `json:"session_uuid"`
	// Название мелодии.
	Melody string `json:"melody"`
	// Частота звука.
	Frequency float64 `json:"frequency"`
	// Длительность в миллисекундах.
	LengthMS int64 `json:"length_ms"`
	// Порядковый номер в мелодии.
	SerialNumber int `json:"serial_number"`
}
