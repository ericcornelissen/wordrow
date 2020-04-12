package cli

import "testing"


func TestArgContextToString(t *testing.T) {
  t.Run("contextInputFile", func(t *testing.T) {
    result := contextInputFile.String()
    if result == "" {
      t.Error("result should not be an empty string")
    }
  })
  t.Run("contextConfigFile", func(t *testing.T) {
    result := contextConfigFile.String()
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
}
