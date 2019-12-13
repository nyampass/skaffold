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

package deploy

import (
	"context"
	"fmt"
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/constants"
	deploy "github.com/GoogleContainerTools/skaffold/pkg/skaffold/deploy/kubectl"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/event"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/runner/runcontext"
)

// CloudRunDeployer deploys workflows using CloudRun CLI.
type CloudRunDeployer struct {
	originalImages     []build.Artifact
	workingDir         string
	CloudRun           deploy.CLI
	insecureRegistries map[string]bool
}

// NewCloudRunDeployer returns a new CloudRunDeployer for a DeployConfig filled
// with the needed configuration for `CloudRun apply`
func NewCloudRunDeployer(runCtx *runcontext.RunContext) *CloudRunDeployer {

	fmt.Println("in Cloud Run!!!!!!!!!!!!!!!!!")

	return &CloudRunDeployer{
		workingDir:         runCtx.WorkingDir,
		insecureRegistries: runCtx.InsecureRegistries,
	}
}

func (k *CloudRunDeployer) Labels() map[string]string {
	return map[string]string{
		constants.Labels.Deployer: "CloudRun",
	}
}

// Deploy templates the provided manifests with a simple `find and replace` and
// runs `CloudRun apply` on those manifests
func (k *CloudRunDeployer) Deploy(ctx context.Context, out io.Writer, builds []build.Artifact, labellers []Labeller) *Result {
	event.DeployInProgress()
	// manifests, err := k.renderManifests(ctx, out, builds)

	// if err != nil {
	// 	event.DeployFailed(err)
	// 	return NewDeployErrorResult(err)
	// }

	// if len(manifests) == 0 {
	// 	event.DeployComplete()
	// 	return NewDeploySuccessResult(nil)
	// }

	// manifests, err = manifests.SetLabels(merge(labellers...))
	// if err != nil {
	// 	event.DeployFailed(err)
	// 	return NewDeployErrorResult(errors.Wrap(err, "setting labels in manifests"))
	// }

	// namespaces, err := manifests.CollectNamespaces()
	// if err != nil {
	// 	event.DeployInfoEvent(errors.Wrap(err, "could not fetch deployed resource namespace. "+
	// 		"This might cause port-forward and deploy health-check to fail."))
	// }

	// if err := k.CloudRun.Apply(ctx, textio.NewPrefixWriter(out, " - "), manifests); err != nil {
	// 	event.DeployFailed(err)
	// 	return NewDeployErrorResult(errors.Wrap(err, "CloudRun error"))
	// }

	// event.DeployComplete()
	return NewDeploySuccessResult(nil)
}

func (k *CloudRunDeployer) Dependencies() ([]string, error) {
	return nil, nil
}

// Cleanup deletes what was deployed by calling Deploy.
func (k *CloudRunDeployer) Cleanup(ctx context.Context, out io.Writer) error {
	// manifests, err := k.readManifests(ctx)
	// if err != nil {
	// 	return errors.Wrap(err, "reading manifests")
	// }

	// // revert remote manifests
	// // TODO(dgageot): That seems super dangerous and I don't understand
	// // why we need to update resources just before we delete them.
	// if len(k.RemoteManifests) > 0 {
	// 	var rm deploy.ManifestList
	// 	for _, m := range k.RemoteManifests {
	// 		manifest, err := k.readRemoteManifest(ctx, m)
	// 		if err != nil {
	// 			return errors.Wrap(err, "get remote manifests")
	// 		}
	// 		rm = append(rm, manifest)
	// 	}

	// 	upd, err := rm.ReplaceImages(k.originalImages)
	// 	if err != nil {
	// 		return errors.Wrap(err, "replacing with originals")
	// 	}

	// 	if err := k.CloudRun.Apply(ctx, out, upd); err != nil {
	// 		return errors.Wrap(err, "apply original")
	// 	}
	// }

	// if err := k.CloudRun.Delete(ctx, textio.NewPrefixWriter(out, " - "), manifests); err != nil {
	// 	return errors.Wrap(err, "delete")
	// }

	return nil
}

func (k *CloudRunDeployer) Render(ctx context.Context, out io.Writer, builds []build.Artifact, filepath string) error {
	return nil
}
