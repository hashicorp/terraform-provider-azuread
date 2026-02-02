package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyCorrelationGroupFeature struct {
	// Indicates the device's feature type. Possible values are: manufacturer, model, osVersion, application or driver.
	DeviceFeatureType *UserExperienceAnalyticsAnomalyDeviceFeatureType `json:"deviceFeatureType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specific metric values that describe the features of the given device feature type.
	Values *[]string `json:"values,omitempty"`
}
