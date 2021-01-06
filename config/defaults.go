package config

func GetDefaults() config {
	defaults := config{}
	defaults.WebServerPort = 8080
	defaults.RTMPServerPort = 1935
	defaults.VideoSettings.ChunkLengthInSeconds = 4
	defaults.Files.MaxNumberInPlaylist = 5
	defaults.YP.Enabled = false
	defaults.YP.YPServiceURL = "https://yp.owncast.online"
	defaults.DatabaseFilePath = "data/owncast.db"
	defaults.DisableUpgradeChecks = false

	defaultQuality := StreamQuality{
		IsAudioPassthrough: true,
		VideoBitrate:       1200,
		EncoderPreset:      "veryfast",
		Framerate:          24,
	}
	defaults.VideoSettings.StreamQualities = []StreamQuality{defaultQuality}

	return defaults
}
