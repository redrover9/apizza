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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Order pizza from dominos",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cached, err := cmd.Flags().GetBool("cached"); cached && err == nil {
			fmt.Println("this is where you would see previous orders and saved orders")
			return nil
		} else if err != nil {
			return err
		}

		print("under constuction!")
		return nil
	},
}

func init() {
	orderCmd.Flags().BoolP("cached", "c", false, "show the previously cached and saved orders")
	rootCmd.AddCommand(orderCmd)
}
