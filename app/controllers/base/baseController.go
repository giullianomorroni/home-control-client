// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	ControllerBase provides all the boilerplate code for controllers.
	Intercept functions are implemented to handle getting a MongoDB
	session and to release that session.

	The ControllerBase also provides access to the services base object
	providing boilerplate code and state for the different services.
*/
package controllerBase

import (
	"github.com/robfig/revel"
)

//** TYPES

type (
	// BaseController contains common fields and behavior for all controllers
	BaseController struct {
		*revel.Controller
	}
)

//** INTERCEPT FUNCTIONS

// Before is called prior to the controller method
func (this *BaseController) Before() revel.Result {
	return nil
}

// After is called once the controller method completes
func (this *BaseController) After() revel.Result {
	return nil
}

// Panic is called is an panic is caught by the framework
func (this *BaseController) Panic() revel.Result {
	return nil
}

//** PUBLIC FUNCTIONS

// Base returns a pointer of the BaseController type
func (this *BaseController) Base() *BaseController {
	return this
}

