package config_test

import (
	"github.com/jasonmccuen/go-splash/internal/config"
	"testing"
)

func TestInitConfig(t *testing.T) {
	c := config.InitConfig()

	if c.SplashHost != "localhost" {
		t.Errorf("SplashHost does not match")
	}
}