// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	rbacv1 "k8s.io/api/rbac/v1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/rbac/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// ClusterRoleInterface is an autogenerated mock type for the ClusterRoleInterface type
type ClusterRoleInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, clusterRole, opts
func (_m *ClusterRoleInterface) Apply(ctx context.Context, clusterRole *v1.ClusterRoleApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1.ClusterRole, error) {
	ret := _m.Called(ctx, clusterRole, opts)

	var r0 *rbacv1.ClusterRole
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ClusterRoleApplyConfiguration, metav1.ApplyOptions) *rbacv1.ClusterRole); ok {
		r0 = rf(ctx, clusterRole, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.ClusterRole)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.ClusterRoleApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, clusterRole, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, clusterRole, opts
func (_m *ClusterRoleInterface) Create(ctx context.Context, clusterRole *rbacv1.ClusterRole, opts metav1.CreateOptions) (*rbacv1.ClusterRole, error) {
	ret := _m.Called(ctx, clusterRole, opts)

	var r0 *rbacv1.ClusterRole
	if rf, ok := ret.Get(0).(func(context.Context, *rbacv1.ClusterRole, metav1.CreateOptions) *rbacv1.ClusterRole); ok {
		r0 = rf(ctx, clusterRole, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.ClusterRole)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *rbacv1.ClusterRole, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, clusterRole, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *ClusterRoleInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *ClusterRoleInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *ClusterRoleInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*rbacv1.ClusterRole, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *rbacv1.ClusterRole
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *rbacv1.ClusterRole); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.ClusterRole)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *ClusterRoleInterface) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.ClusterRoleList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *rbacv1.ClusterRoleList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *rbacv1.ClusterRoleList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.ClusterRoleList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *ClusterRoleInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*rbacv1.ClusterRole, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rbacv1.ClusterRole
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *rbacv1.ClusterRole); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.ClusterRole)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, clusterRole, opts
func (_m *ClusterRoleInterface) Update(ctx context.Context, clusterRole *rbacv1.ClusterRole, opts metav1.UpdateOptions) (*rbacv1.ClusterRole, error) {
	ret := _m.Called(ctx, clusterRole, opts)

	var r0 *rbacv1.ClusterRole
	if rf, ok := ret.Get(0).(func(context.Context, *rbacv1.ClusterRole, metav1.UpdateOptions) *rbacv1.ClusterRole); ok {
		r0 = rf(ctx, clusterRole, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.ClusterRole)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *rbacv1.ClusterRole, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, clusterRole, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *ClusterRoleInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
