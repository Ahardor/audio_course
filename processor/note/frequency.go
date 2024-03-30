package note

const (
	// FrequencyBaseUnit описывает базовые единицы представления частоты.
	// В 1 Hz содержится FrequencyBaseUnit базовых единиц.
	// Для работы с частотами подразумевается округление частоты до 2 знаков после запятой.
	// Data source: https://mixbutton.com/mixing-articles/music-note-to-frequency-chart/.
	FrequencyBaseUnit = 100
	// OctavesCount количество октав.
	OctavesCount = 9
)

// Frequency - частота ноты для октавы в базовых единицах.
type Frequency int

// FrequencyList - набор частот, характерных для ноты в 0-8 октавах.
type FrequencyList [OctavesCount]Frequency
