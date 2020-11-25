package utilities

import (
	"testing"

	"github.com/logicmonitor/k8s-argus/pkg/constants"
	"github.com/stretchr/testify/assert"
	"github.com/vkumbhar94/lm-sdk-go/models"
)

var (
	deviceName             = "test-device"
	customPropertiesName1  = "name1"
	customPropertiesValue1 = "value1"
	customPropertiesName2  = "name2"
	customPropertiesValue2 = "value2"

	systemPropertiesName1  = "system-name1"
	systemPropertiesValue1 = "system-value1"
	systemPropertiesName2  = "system-name2"
	systemPropertiesValue2 = "system-value2"
)

func TestGetPropertyValue(t *testing.T) {
	device := getDevice()
	value := GetPropertyValue(device, customPropertiesName1)
	t.Logf("name=%s, value=%s", customPropertiesName1, value)
	value = GetPropertyValue(device, systemPropertiesName2)
	t.Logf("name=%s, value=%s", systemPropertiesName2, value)
	value = GetPropertyValue(device, "non-exist-name")
	t.Logf("name=%s, value=%s", "non-exist-name", value)
}

func TestGetDesiredDisplayNameByResourceAndConfig(t *testing.T) {
	t.Parallel()
	TestCases := []struct {
		testcasename                  string
		name                          string
		namespace                     string
		clusterName                   string
		resource                      string
		displayNameIncludeNamespace   bool
		displayNameIncludeClusterName bool
		expectedResult                string
	}{
		{
			testcasename:                  "Get displayName for pod when full name is disabled.",
			name:                          "pod-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.Pods,
			displayNameIncludeNamespace:   false,
			displayNameIncludeClusterName: false,
			expectedResult:                "pod-device-pod",
		},
		{
			testcasename:                  "Get displayName for pod when namespace is included",
			name:                          "pod-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.Pods,
			displayNameIncludeNamespace:   true,
			displayNameIncludeClusterName: false,
			expectedResult:                "pod-device-pod-default",
		},
		{
			testcasename:                  "Get displayName for pod when full name is enabled.",
			name:                          "pod-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.Pods,
			displayNameIncludeNamespace:   true,
			displayNameIncludeClusterName: true,
			expectedResult:                "pod-device-pod-default-cluster1",
		},
		{
			testcasename:                  "Get displayName for service when full name is disabled.",
			name:                          "service-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.Services,
			displayNameIncludeNamespace:   false,
			displayNameIncludeClusterName: false,
			expectedResult:                "service-device-svc",
		},
		{
			testcasename:                  "Get displayName for service when namespace is included",
			name:                          "service-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.Services,
			displayNameIncludeNamespace:   true,
			displayNameIncludeClusterName: false,
			expectedResult:                "service-device-svc-default",
		},
		{
			testcasename:                  "Get displayName for service when full name is enabled.",
			name:                          "service-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.Services,
			displayNameIncludeNamespace:   true,
			displayNameIncludeClusterName: true,
			expectedResult:                "service-device-svc-default-cluster1",
		},
		{
			testcasename:                  "Get displayName for deployment when full name is disabled.",
			name:                          "deloyment-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.Deployments,
			displayNameIncludeNamespace:   false,
			displayNameIncludeClusterName: false,
			expectedResult:                "deloyment-device-deploy",
		},
		{
			testcasename:                  "Get displayName for deloyment when namespace is included",
			name:                          "deloyment-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.Deployments,
			displayNameIncludeNamespace:   true,
			displayNameIncludeClusterName: false,
			expectedResult:                "deloyment-device-deploy-default",
		},
		{
			testcasename:                  "Get displayName for deloyment when full name is enabled.",
			name:                          "deloyment-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.Deployments,
			displayNameIncludeNamespace:   true,
			displayNameIncludeClusterName: true,
			expectedResult:                "deloyment-device-deploy-default-cluster1",
		},
		{
			testcasename:                  "Get displayName for hpa when full name is disabled.",
			name:                          "hpa-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.HorizontalPodAutoScalers,
			displayNameIncludeNamespace:   false,
			displayNameIncludeClusterName: false,
			expectedResult:                "hpa-device-hpa",
		},
		{
			testcasename:                  "Get displayName for hpa when namespace is included",
			name:                          "hpa-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.HorizontalPodAutoScalers,
			displayNameIncludeNamespace:   true,
			displayNameIncludeClusterName: false,
			expectedResult:                "hpa-device-hpa-default",
		},
		{
			testcasename:                  "Get displayName for hpa when full name is enabled.",
			name:                          "hpa-device",
			namespace:                     "default",
			clusterName:                   "cluster1",
			resource:                      constants.HorizontalPodAutoScalers,
			displayNameIncludeNamespace:   true,
			displayNameIncludeClusterName: true,
			expectedResult:                "hpa-device-hpa-default-cluster1",
		},
	}

	assert := assert.New(t)
	// nolint: dupl
	for _, testCase := range TestCases {
		t.Run(testCase.testcasename, func(t *testing.T) {
			result := GetDesiredDisplayNameByResourceAndConfig(testCase.name, testCase.namespace, testCase.clusterName, testCase.resource, testCase.displayNameIncludeNamespace, testCase.displayNameIncludeClusterName)

			// check expected evaluation result
			assert.Equal(testCase.expectedResult, result, "TestCase: \"%s\" \nResult: Expected evaluate \"%s\" but got \"%s\"", testCase.testcasename, testCase.expectedResult, result)
		})
	}
}

func TestGetNameWithResourceType(t *testing.T) {
	t.Parallel()
	TestCases := []struct {
		testcasename   string
		name           string
		resource       string
		expectedResult string
	}{
		{
			testcasename:   "Get name for Pod",
			name:           "pod-device",
			resource:       constants.Pods,
			expectedResult: "pod-device-pod",
		},
		{
			testcasename:   "Get name for Service",
			name:           "service-device",
			resource:       constants.Services,
			expectedResult: "service-device-svc",
		},
		{
			testcasename:   "Get name for Deployment",
			name:           "deployment-device",
			resource:       constants.Deployments,
			expectedResult: "deployment-device-deploy",
		},
		{
			testcasename:   "Get name for Node",
			name:           "node-device",
			resource:       constants.Nodes,
			expectedResult: "node-device-node",
		},
		{
			testcasename:   "Get name for HPA",
			name:           "hpa-device",
			resource:       constants.HorizontalPodAutoScalers,
			expectedResult: "hpa-device-hpa",
		},
	}
	assert := assert.New(t)
	// nolint: dupl
	for _, testCase := range TestCases {
		t.Run(testCase.testcasename, func(t *testing.T) {
			result := GetNameWithResourceType(testCase.name, testCase.resource)

			// check expected evaluation result
			assert.Equal(testCase.expectedResult, result, "TestCase: \"%s\" \nResult: Expected evaluate \"%s\" but got \"%s\"", testCase.testcasename, testCase.expectedResult, result)
		})
	}
}

func TestGetConflictCategoryByResourceType(t *testing.T) {
	t.Parallel()
	TestCases := []struct {
		testcasename   string
		resource       string
		expectedResult string
	}{
		{
			testcasename:   "Get conflict category for Pod",
			resource:       constants.Pods,
			expectedResult: constants.PodConflictCategory,
		},
		{
			testcasename:   "Get conflict category for Service",
			resource:       constants.Services,
			expectedResult: constants.ServiceConflictCategory,
		},
		{
			testcasename:   "Get conflict category for Deployment",
			resource:       constants.Deployments,
			expectedResult: constants.DeploymentConflictCategory,
		},
		{
			testcasename:   "Get conflict category for Node",
			resource:       constants.Nodes,
			expectedResult: constants.NodeConflictCategory,
		},
		{
			testcasename:   "Get conflict category for HPA",
			resource:       constants.HorizontalPodAutoScalers,
			expectedResult: constants.HorizontalPodAutoscalerConflictCategory,
		},
	}
	assert := assert.New(t)
	// nolint: dupl
	for _, testCase := range TestCases {
		t.Run(testCase.testcasename, func(t *testing.T) {
			result := GetConflictCategoryByResourceType(testCase.resource)

			// check expected evaluation result
			assert.Equal(testCase.expectedResult, result, "TestCase: \"%s\" \nResult: Expected evaluate \"%s\" but got \"%s\"", testCase.testcasename, testCase.expectedResult, result)
		})
	}
}

func getDevice() *models.Device {
	return &models.Device{
		Name:        &deviceName,
		DisplayName: &deviceName,
		CustomProperties: []*models.NameAndValue{
			{
				Name:  &customPropertiesName1,
				Value: &customPropertiesValue1,
			}, {
				Name:  &customPropertiesName2,
				Value: &customPropertiesValue2,
			},
		},
		SystemProperties: []*models.NameAndValue{
			{
				Name:  &systemPropertiesName1,
				Value: &systemPropertiesValue1,
			}, {
				Name:  &systemPropertiesName2,
				Value: &systemPropertiesValue2,
			},
		},
	}
}
