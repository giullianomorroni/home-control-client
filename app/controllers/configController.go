package controllers

import (
	cb "home-control-client/app/controllers/base"
	"github.com/robfig/revel"
)

//** TYPES
type (
	Config struct {
		cb.BaseController
	}
)

func init() { }

func (this *Config) Index() revel.Result {
	return this.Render();
}

func (this *Config) NewDevice() revel.Result {
	return this.Render();
}
