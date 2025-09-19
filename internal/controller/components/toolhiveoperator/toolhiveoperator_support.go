package toolhiveoperator

import (
	"path"

	componentApi "github.com/opendatahub-io/opendatahub-operator/v2/api/components/v1alpha1"
	"github.com/opendatahub-io/opendatahub-operator/v2/internal/controller/status"
	odhtypes "github.com/opendatahub-io/opendatahub-operator/v2/pkg/controller/types"
	odhdeploy "github.com/opendatahub-io/opendatahub-operator/v2/pkg/deploy"
)

var (
	paramsPath = path.Join(odhdeploy.DefaultManifestPath, ComponentName, "base")

	imageParamMap = map[string]string{
		"toolhive-operator-image": "RELATED_IMAGE_ODH_TOOLHIVE_OPERATOR_IMAGE",
		"toolhive-proxy-image":    "RELATED_IMAGE_ODH_TOOLHIVE_PROXY_IMAGE",
	}

	conditionTypes = []string{
		status.ConditionDeploymentsAvailable,
	}
)

const (
	ComponentName = componentApi.ToolHiveOperatorComponentName

	ReadyConditionType = componentApi.ToolHiveOperatorKind + status.ReadySuffix

	// LegacyComponentName is the name of the component that is assigned to deployments
	// via Kustomize. Since a deployment selector is immutable, we can't upgrade existing
	// deployment to the new component name, so keep it around till we figure out a solution.
	LegacyComponentName = "toolhiveoperator"
)

func manifestsPath() odhtypes.ManifestInfo {
	return odhtypes.ManifestInfo{
		Path:       odhdeploy.DefaultManifestPath,
		ContextDir: ComponentName,
		SourcePath: "base",
	}
}
