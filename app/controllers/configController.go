package controllers

import (
	cb "home-control-client/app/controllers/base"
	model "home-control-client/app/models/config"
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
	accountEmail := model.ReadAccountEmail()
	wifiName := model.ReadWiFiName()
	return this.Render(accountEmail, wifiName);
}

func (this *Config) NewDevice() revel.Result {
	return this.Render();
}
