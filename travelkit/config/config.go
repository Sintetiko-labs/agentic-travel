// Package config handles on-disk state under ~/.{slug}/.
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// Dir returns the config directory for envPrefix (e.g. MELIA → ~/.melia).
func Dir(envPrefix string) string {
	slug := strings.ToLower(envPrefix)
	if d := os.Getenv(envPrefix + "_CONFIG_DIR"); d != "" {
		return d
	}
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return "." + slug
	}
	return filepath.Join(home, "."+slug)
}

// Load reads a JSON file from the config directory.
func Load(envPrefix, name string, v any) error {
	b, err := os.ReadFile(filepath.Join(Dir(envPrefix), name))
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// Save writes v as indented JSON into the config directory.
func Save(envPrefix, name string, v any) error {
	dir := Dir(envPrefix)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, name), b, 0o600)
}
