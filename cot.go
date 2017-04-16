// Copyright (c) 2017, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package cot

import (
	"github.com/Sirupsen/logrus"
)

// log is the logger for the entire library
var log = logrus.New()

// Logger returns the internal logrus Logger for this package.
//
// This can be set or used as is, for example:
//
//	*(cot.Logger()) = logrus.New()
//	cot.Logger().SetLevel(logrus.DebugLevel)
func Logger() *logrus.Logger {
	return log
}
