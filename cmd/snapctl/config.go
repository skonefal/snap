/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/intelsdi-x/snap/core/ctypes"
)

func getConfig(ctx *cli.Context) {
	ptyp := ctx.String("plugin-type")
	pname := ctx.String("plugin-name")
	pver := ctx.Int("plugin-version")
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	defer w.Flush()
	printFields(w, false, 0,
		"NAME",
		"VALUE",
		"TYPE",
	)
	r := pClient.GetPluginConfig(ptyp, pname, strconv.Itoa(pver))
	for k, v := range r.Table() {
		switch t := v.(type) {
		case ctypes.ConfigValueInt:
			printFields(w, false, 0, k, t.Value, t.Type())
		case ctypes.ConfigValueBool:
			printFields(w, false, 0, k, t.Value, t.Type())
		case ctypes.ConfigValueFloat:
			printFields(w, false, 0, k, t.Value, t.Type())
		case ctypes.ConfigValueStr:
			printFields(w, false, 0, k, t.Value, t.Type())
		}
	}
}
