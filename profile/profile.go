package profile

import (
	"path/filepath"
	"strings"

	"github.com/arbal/alfred-brave-history/utils"
)

const (
	braveConfigPath = "~/Library/Application Support/BraveSoftware/Brave-Browser"
)

func GetProfilePath(profile string) (string, error) {
	var path string
	if strings.Contains(profile, "/") {
		path = filepath.Clean(profile)
	} else {
		if profile == "" {
			profile = "Default"
		}
		path = filepath.Join(braveConfigPath, profile)
	}

	return utils.ExpandTilde(path)
}
