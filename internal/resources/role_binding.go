package resources

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RoleBinding = rbacv1.ClusterRoleBinding{
	TypeMeta: metav1.TypeMeta{
		APIVersion: rbacGV,
		Kind:       "RoleBinding",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      ResourceNameDefault,
		Namespace: ResourceNameDefault,
		Labels: map[string]string{
			ResourceNameDefault: "",
		},
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: rbacAPIGroup,
		Kind:     Role.Kind,
		Name:     Role.Name,
	},
	Subjects: []rbacv1.Subject{
		{
			Kind:      ServiceAccount.Kind,
			Name:      ServiceAccount.Name,
			Namespace: ServiceAccount.Namespace,
		},
	},
}
