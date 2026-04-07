package config

type Config struct {
	WidthCells   int
	HeightCells  int
	ExposureRate float64
	TimeSpeed    float64
	WaveMode     bool
}

func NewDefault() Config {
	return Config{
		WidthCells:   400,
		HeightCells:  400,
		ExposureRate: 8e-4,
		TimeSpeed:    1,
		WaveMode:     false,
	}
}
