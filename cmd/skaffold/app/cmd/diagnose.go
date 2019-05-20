/*
Copyright 2019 The Skaffold Authors

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

package cmd

import (
	"fmt"
	"io"

	"github.com/GoogleContainerTools/skaffold/cmd/skaffold/app/cmd/commands"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/version"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	yaml "gopkg.in/yaml.v2"
)

// NewCmdDiagnose describes the CLI command to diagnose skaffold.
func NewCmdDiagnose(out io.Writer) *cobra.Command {
	return commands.
		New(out).
		WithDescription("diagnose", "Run a diagnostic on Skaffold").
		WithFlags(func(f *pflag.FlagSet) {
			f.StringVarP(&opts.ConfigurationFile, "filename", "f", "skaffold.yaml", "Filename or URL to the pipeline file")
			f.StringSliceVarP(&opts.Profiles, "profile", "p", nil, "Activate profiles by name")
		}).
		NoArgs(doDiagnose)
}

func doDiagnose(out io.Writer) error {
	runner, config, err := newRunner(opts)
	if err != nil {
		return errors.Wrap(err, "creating runner")
	}

	fmt.Fprintln(out, "Skaffold version:", version.Get().GitCommit)
	fmt.Fprintln(out, "Configuration version:", config.APIVersion)
	fmt.Fprintln(out, "Number of artifacts:", len(config.Build.Artifacts))

	if err := runner.DiagnoseArtifacts(out); err != nil {
		return errors.Wrap(err, "running diagnostic on artifacts")
	}

	color.Blue.Fprintln(out, "\nConfiguration")
	buf, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "marshalling configuration")
	}
	out.Write(buf)

	return nil
}
