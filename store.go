package tgtk4

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type MediaItem struct {
	Path      string    `json:"path"`
	Mode      string    `json:"mode"`
	Timestamp time.Time `json:"timestamp"`
	Tags      []string  `json:"tags,omitempty"`
}

type MediaStore struct {
	Dir    string
	DbPath string
	Items  []MediaItem
}

func NewMediaStore(dir, configName string) *MediaStore {
	os.MkdirAll(dir, 0755)
	s := &MediaStore{
		Dir:    dir,
		DbPath: filepath.Join(ConfigDir(configName), "store.json"),
	}
	s.Load()
	s.Prune()
	return s
}

func (s *MediaStore) Load() {
	data, err := os.ReadFile(s.DbPath)
	if err != nil {
		s.Items = []MediaItem{}
		return
	}
	if err := json.Unmarshal(data, &s.Items); err != nil {
		s.Items = []MediaItem{}
	}
}

func (s *MediaStore) Save() {
	os.MkdirAll(filepath.Dir(s.DbPath), 0755)
	data, err := json.MarshalIndent(s.Items, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(s.DbPath, data, 0644)
}

func (s *MediaStore) Prune() {
	valid := []MediaItem{}
	for _, item := range s.Items {
		if _, err := os.Stat(item.Path); err == nil {
			valid = append(valid, item)
		}
	}
	if len(valid) != len(s.Items) {
		s.Items = valid
		s.Save()
	}
}

func (s *MediaStore) Add(path, mode string) {
	s.Items = append(s.Items, MediaItem{
		Path:      path,
		Mode:      mode,
		Timestamp: time.Now(),
	})
	s.Save()
}

func (s *MediaStore) RemoveBatch(paths []string) error {
	pathMap := make(map[string]bool)
	for _, p := range paths {
		os.Remove(p)
		pathMap[p] = true
	}

	filtered := []MediaItem{}
	for _, item := range s.Items {
		if !pathMap[item.Path] {
			filtered = append(filtered, item)
		}
	}
	s.Items = filtered
	s.Save()
	return nil
}

func (s *MediaStore) Sorted() []MediaItem {
	sorted := make([]MediaItem, len(s.Items))
	copy(sorted, s.Items)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Timestamp.After(sorted[j].Timestamp)
	})
	return sorted
}

func (s *MediaStore) UpdateTags(path string, tags []string) {
	for i := range s.Items {
		if s.Items[i].Path == path {
			s.Items[i].Tags = tags
			break
		}
	}
	s.Save()
}
