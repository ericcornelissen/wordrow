package common

import "github.com/ericcornelissen/wordrow/internal/logger"

// MapFrom creates a map from a
func MapFrom(froms []string, to string) map[string]string {
	mapping := make(map[string]string, len(froms))
	for _, from := range froms {
		mapping[from] = to
	}

	return mapping
}

// MergeMaps merges the maps `target` and `other` into map `target`. Keys
// present in both `target` and `other` will end up with the value of `other`.
func MergeMaps(target, other map[string]string) {
	for key, value := range other {
		if oldValue, present := target[key]; present {
			logger.Debugf("Overwriting '%s': from '%s' to '%s'", key, oldValue, value)
		}

		target[key] = value
	}
}
