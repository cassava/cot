// Copyright (c) 2017, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package internal

import (
	"github.com/Sirupsen/logrus"
)

type ConsoleFormatter struct{}

func (cf ConsoleFormatter) Format(e *logrus.Entry) ([]byte, error) {
	return []byte(e.Message + "\n"), nil
}
