package tgtk4

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Colors Colors `toml:"colors"`
}

func ConfigDir(name string) string {
	if d := os.Getenv("XDG_CONFIG_HOME"); d != "" {
		return filepath.Join(d, name)
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", name)
}

func LoadConfig[T any](name string, cfg *T) error {
	path := filepath.Join(ConfigDir(name), "config.toml")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(path), 0755)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		return toml.NewEncoder(f).Encode(cfg)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return toml.Unmarshal(data, cfg)
}

func SaveConfig[T any](name string, cfg *T) error {
	path := filepath.Join(ConfigDir(name), "config.toml")
	os.MkdirAll(filepath.Dir(path), 0755)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(cfg)
}
