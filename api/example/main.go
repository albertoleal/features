// Copyright 2015 Features authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/albertoleal/features/api"
	"github.com/albertoleal/features/engine/memory"
)

func main() {
	api := api.NewApi(memory.New())
	api.Run()
}
