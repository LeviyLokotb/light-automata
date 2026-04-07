package config

type Config struct {
	WidthCells   int     `yaml:"width"`
	HeightCells  int     `yaml:"height"`
	ExposureRate float64 `yaml:"exposure_rate"`
	WaveMode     bool    `yaml:"wave_mode"`
	PixelSize    int     `yaml:"pixel_size"`
}

func NewDefault() Config {
	return Config{
		WidthCells:   400,
		HeightCells:  400,
		ExposureRate: 8e-4,
		WaveMode:     false,
		PixelSize:    1,
	}
}
