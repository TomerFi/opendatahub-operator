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
)

func manifestsPath() odhtypes.ManifestInfo {
	return odhtypes.ManifestInfo{
		Path:       odhdeploy.DefaultManifestPath,
		ContextDir: ComponentName,
		SourcePath: "base",
	}
}
