package ffmpeg

import "time"

// Config ...
type Config struct {
	FfmpegBinPath   string
	FfprobeBinPath  string
	ProgressEnabled bool
	Verbose         bool
	Timeout         time.Duration
}
