package state

import (
	"context"
	"testing"

	"github.com/kyma-project/serverless/components/operator/api/v1alpha1"
	"github.com/kyma-project/serverless/components/operator/internal/chart"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestUpdateNetworkPoliciesStatus(t *testing.T) {
	scheme := runtime.NewScheme()
	tests := []struct {
		name                  string
		enableNetworkPolicies bool
		expectedStatus        string
	}{
		{
			name:                  "NetworkPolicies enabled",
			enableNetworkPolicies: true,
			expectedStatus:        "True",
		},
		{
			name:                  "NetworkPolicies disabled",
			enableNetworkPolicies: false,
			expectedStatus:        "False",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.TODO()
			s := &systemState{
				instance: v1alpha1.Serverless{
					Spec: v1alpha1.ServerlessSpec{
						EnableNetworkPolicies: tt.enableNetworkPolicies,
					},
				},
				flagsBuilder: chart.NewFlagsBuilder(),
			}
			c := fake.NewClientBuilder().WithScheme(scheme).Build()
			r := &reconciler{log: zap.NewNop().Sugar(), k8s: k8s{client: c, EventRecorder: record.NewFakeRecorder(5)}}
			next, result, err := sFnOptionalDependencies(ctx, r, s)

			require.Nil(t, err)
			require.Nil(t, result)
			requireEqualFunc(t, sFnApplyResources, next)
			status := s.instance.Status
			assert.Equal(t, tt.expectedStatus, status.NetworkPoliciesEnabled)
		})
	}
}
