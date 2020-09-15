package kube

import (
	"fmt"
	// "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/kube"
	ovntest "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/testing"
	clientgo_mock "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/testing/mocks/k8s.io/client-go/kubernetes"
	core_mock "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/testing/mocks/k8s.io/client-go/kubernetes/typed/core/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestKube_Events(t *testing.T) {
	mockInterface := new(clientgo_mock.Interface)
	mockCoreInterface := new(core_mock.CoreV1Interface)
	// mockPodInterface := new(core_mock.PodInterface)
	k := &Kube{}
	tests := []struct {
		desc              string
		expectedErr       bool
		errorMatch        error
		onRetArgs         *ovntest.TestifyMockHelper
		onRetMockCoreArgs *ovntest.TestifyMockHelper
		onRetMockPodArgs  *ovntest.TestifyMockHelper
	}{
		{
			desc:              "Positive test code path",
			expectedErr:       true,
			errorMatch:        fmt.Errorf("required CNI variable missing"),
			onRetArgs:         &ovntest.TestifyMockHelper{"CoreV1", []string{}, []interface{}{mockCoreInterface}},
			onRetMockCoreArgs: &ovntest.TestifyMockHelper{"Events", []string{}, []interface{}{}},
		},
	}
	for i, tc := range tests {

		t.Run(fmt.Sprintf("%d:%s", i, tc.desc), func(t *testing.T) {
			if tc.onRetArgs != nil {
				call := mockInterface.On(tc.onRetArgs.OnCallMethodName)
				for _, arg := range tc.onRetArgs.OnCallMethodArgType {
					call.Arguments = append(call.Arguments, mock.AnythingOfType(arg))
				}
				for _, elem := range tc.onRetArgs.RetArgList {
					call.ReturnArguments = append(call.ReturnArguments, elem)
				}
				call.Once()
			}
			if tc.onRetMockCoreArgs != nil {
				mockCall := mockInterface.On(tc.onRetArgs.OnCallMethodName)
				for _, arg := range tc.onRetMockCoreArgs.OnCallMethodArgType {
					mockCall.Arguments = append(mockCall.Arguments, mock.AnythingOfType(arg))
				}
				for _, elem := range tc.onRetMockCoreArgs.RetArgList {
					mockCall.ReturnArguments = append(mockCall.ReturnArguments, elem)
				}
				mockCall.Once()
			}
			e := k.Events()

			assert.Nil(t, e)
			mockInterface.AssertExpectations(t)
			mockCoreInterface.AssertExpectations(t)

		})
	}
}
