package cli

import "testing"

func TestArgContextToString(t *testing.T) {
	t.Run("contextDefault", func(t *testing.T) {
		result := contextDefault.String()
		if result == "" {
			t.Error("result should not be an empty string")
		}
	})
	t.Run("contextMapFile", func(t *testing.T) {
		result := contextMapFile.String()
		if result == "" {
			t.Error("result should not be an empty string")
		}
	})
	t.Run("contextMapping", func(t *testing.T) {
		result := contextMapping.String()
		if result == "" {
			t.Error("result should not be an empty string")
		}
	})
}
