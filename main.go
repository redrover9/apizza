// Copyright © 2019 Harrison Brown harrybrown98@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/harrybrwn/apizza/cmd"
	"github.com/harrybrwn/apizza/pkg/errs"
)

func main() {
	// TODO:
	// Make the config dir movable
	//
	// config_dir := os.Getenv("APIZZA_CONFIG")
	// if config_dir == "" {
	// 		config_dir = ".apizza"
	// }
	// cmd.Execute(os.Args[1:], config_dir)
	err := cmd.Execute(os.Args[1:], ".apizza")
	if err != nil {
		errs.Handle(err.Err, err.Msg, err.Code)
	}
}
