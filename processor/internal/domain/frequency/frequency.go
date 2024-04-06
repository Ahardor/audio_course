package frequency

import "math"

const (
	// FrequencyBaseUnit описывает базовые единицы (base points) представления частоты.
	// В 1 Hz содержится FrequencyBaseUnit базовых единиц.
	// Для работы с частотами подразумевается округление частоты до 2 знаков после запятой.
	// Data source: https://mixbutton.com/mixing-articles/music-note-to-frequency-chart/.
	FrequencyBaseUnit = 100
	// OctavesCount количество октав.
	OctavesCount = 9
)

// Frequency - частота ноты для октавы в базовых единицах.
type Frequency int

// Octave - октава.
type Octave int

// List - набор частот, характерных для ноты в 0-8 октавах.
type List [OctavesCount]Frequency

// FrequencyFromFloat переводит частоту из float64 в base points.
func FromFloat(fr float64) Frequency {
	return Frequency(math.RoundToEven(fr * FrequencyBaseUnit))
}

// IsApproximatelyEqual определяет, лежит ли частота target в отрезке [freq - delta; freq + delta].
func IsApproximatelyEqual(target, freq Frequency, delta int) bool {
	return target >= freq-Frequency(delta) && target <= freq+Frequency(delta)
}
