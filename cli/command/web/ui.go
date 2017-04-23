// Copyright © 2017 blacktop <https://github.com/blacktop>
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

package web

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/maliceio/malice/cli"
	"github.com/maliceio/malice/cli/command"
	"github.com/spf13/cobra"
)

type uiOptions struct {
	all bool
}

// NewUICommand returns a new cobra ui command for plugins
func NewUICommand(maliceCli *command.MaliceCli) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "ui [OPTIONS]",
		Short: "Start Web UI",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			output, err := runUI(maliceCli, opts)
			if err != nil {
				return err
			}
			if output != "" {
				fmt.Fprintln(maliceCli.Out(), output)
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.all, "all", "a", false, "Show all plugin ui")

	return cmd
}

func runUI(maliceCli *command.MaliceCli, opts uiOptions) (output string, err error) {

	// if !opts.force && !command.PromptForConfirmation(maliceCli.In(), maliceCli.Out(), warning) {
	// 	return
	// }

	ui, err := maliceCli.Client().PluginUI(context.Background(), uiFilters)
	if err != nil {
		return
	}

	if len(ui) > 0 {
		output = "Plugin UI:\n"
		// TODO: format ui output
		output = ui
	}

	return
}