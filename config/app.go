package config

import "goapitest/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"name":     config.Env("APP_NAME", "gotest"),
			"env":      config.Env("APP_ENV", "production"),
			"debug":    config.Env("APP_DEBUG", false),
			"port":     config.Env("APP_PORT", 3000),
			"key":      config.Env("APP_KEY", ""),
			"url":      config.Env("APP_URL", ""),
			"timezone": config.Env("TIMEZONE", ""),
		}
	})
}
