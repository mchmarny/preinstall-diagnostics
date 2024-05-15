package resources

import "github.com/run-ai/preinstall-diagnostics/internal/registry"

const (
	ResourceNameDefault = "runai-diag"

	rbacAPIGroup   = "rbac.authorization.k8s.io"
	rbacAPIVersion = "v1"
	rbacGV         = rbacAPIGroup + "/" + rbacAPIVersion

	coreAPIVersion = "v1"

	appsAPIGroup   = "apps"
	appsAPIVersion = "v1"
	appsGV         = appsAPIGroup + "/" + appsAPIVersion
)

var (
	defaultImage = registry.RunAIDiagnosticsImage
)
