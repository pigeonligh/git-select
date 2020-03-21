package config

import (
	"os/exec"
)

// Config is
type Config struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	KeyPath string `json:"key"`
}

// Setup is
func (c *Config) Setup(global bool) error {
	var cmd *exec.Cmd
	var err error

	if global {
		cmd = exec.Command("git", "config", "--global", "user.name", c.Name)
	} else {
		cmd = exec.Command("git", "config", "user.name", c.Name)
	}
	if err = cmd.Run(); err != nil {
		return err
	}

	if global {
		cmd = exec.Command("git", "config", "--global", "user.email", c.Email)
	} else {
		cmd = exec.Command("git", "config", "user.email", c.Email)
	}
	if err = cmd.Run(); err != nil {
		return err
	}

	if global {
		cmd = exec.Command("git", "config", "--global", "core.sshCommand", "ssh -i "+c.KeyPath+" -F /dev/null")
	} else {
		cmd = exec.Command("git", "config", "core.sshCommand", "ssh -i "+c.KeyPath+" -F /dev/null")
	}
	if err = cmd.Run(); err != nil {
		return err
	}
	return nil
}
