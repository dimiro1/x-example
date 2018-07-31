package app

import (
	"github.com/dimiro1/health"
	"github.com/dimiro1/health/url"
	"github.com/dimiro1/x/xhealth"
)

func GoogleHealthCheck() xhealth.CheckMapping {
	return xhealth.Register("google", url.NewChecker("http://www.google.com"))
}

func HelloFreshHealthCheck() xhealth.CheckMapping {
	return xhealth.Register("hellofresh", url.NewChecker("http://www.hellofresh.com"))
}

func CustomHealthCheck() xhealth.CheckMapping {
	return xhealth.Register("custom", health.CheckerFunc(func() health.Health {
		h := health.NewHealth()
		h.AddInfo("message", "Hello World")
		h.Up()
		return h
	}))
}
