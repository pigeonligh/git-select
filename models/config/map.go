package config

import (
	"errors"

	"github.com/pigeonligh/git-select/models/opts"
)

// Map is
type Map struct {
	Data map[string]Config `json:"data"`
}

// CheckTag is
func (m *Map) CheckTag(tag string) bool {
	if tag == "" {
		return true
	}
	_, ok := m.Data[tag]
	return ok
}

// Add is
func (m *Map) Add(opt opts.ConfigOpts) error {
	if m.CheckTag(opt.Tag) {
		return errors.New("Tag " + opt.Tag + " exists")
	}
	m.Data[opt.Tag] = Config{
		Name:    opt.Name,
		Email:   opt.Email,
		KeyPath: opt.KeyPath,
	}
	return SaveConfig(m)
}

// Remove is
func (m *Map) Remove(tag string) error {
	if m.CheckTag(tag) == false {
		return errors.New("tag " + tag + " does not exist.")
	}
	delete(m.Data, tag)
	return SaveConfig(m)
}

// Select is
func (m *Map) Select(tag string, global bool) error {
	config, ok := m.Data[tag]
	if ok == false {
		return errors.New("tag " + tag + " does not exist.")
	}
	return config.Setup(global)
}
