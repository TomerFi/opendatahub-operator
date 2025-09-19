package toolhiveoperator

import (
	"context"
	"fmt"

	componentApi "github.com/opendatahub-io/opendatahub-operator/v2/api/components/v1alpha1"
	odhtypes "github.com/opendatahub-io/opendatahub-operator/v2/pkg/controller/types"
	odhdeploy "github.com/opendatahub-io/opendatahub-operator/v2/pkg/deploy"
)

func initialize(_ context.Context, rr *odhtypes.ReconciliationRequest) error {
	rr.Manifests = append(rr.Manifests, manifestsPath())
	return nil
}

func devFlags(ctx context.Context, rr *odhtypes.ReconciliationRequest) error {
	toolhiveoperator, ok := rr.Instance.(*componentApi.ToolHiveOperator)
	if !ok {
		return fmt.Errorf("resource instance %v is not a componentApi.ToolHiveOperator)", rr.Instance)
	}

	if toolhiveoperator.Spec.DevFlags == nil {
		return nil
	}

	// If dev flags are set, update default manifests path
	if len(toolhiveoperator.Spec.DevFlags.Manifests) != 0 {
		manifestConfig := toolhiveoperator.Spec.DevFlags.Manifests[0]
		if err := odhdeploy.DownloadManifests(ctx, ComponentName, manifestConfig); err != nil {
			return err
		}
		if manifestConfig.SourcePath != "" {
			rr.Manifests[0].Path = odhdeploy.DefaultManifestPath
			rr.Manifests[0].ContextDir = ComponentName
			rr.Manifests[0].SourcePath = manifestConfig.SourcePath
		}
	}
	return nil
}
