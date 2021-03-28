package cache

import (
	"errors"
	"fmt"
)


type Config struct {
	MaxSize int  `json:"maxSize"`
}


func (c *Config) Check() error {
	if c == nil {
		return errors.New("cache config is nil")
	}

	if c.MaxSize <= 0 {
		return errors.New(fmt.Sprintf("invalid MaxSize %d\n", c.MaxSize))
	}

	return nil
}
