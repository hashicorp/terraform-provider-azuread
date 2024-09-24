package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalySeverityOverview struct {
	// Indicates count of high severity anomalies which have been detected. Valid values -2147483648 to 2147483647
	HighSeverityAnomalyCount *int64 `json:"highSeverityAnomalyCount,omitempty"`

	// Indicates count of informational severity anomalies which have been detected. Valid values -2147483648 to 2147483647
	InformationalSeverityAnomalyCount *int64 `json:"informationalSeverityAnomalyCount,omitempty"`

	// Indicates count of low severity anomalies which have been detected. Valid values -2147483648 to 2147483647
	LowSeverityAnomalyCount *int64 `json:"lowSeverityAnomalyCount,omitempty"`

	// Indicates count of medium severity anomalies which have been detected. Valid values -2147483648 to 2147483647
	MediumSeverityAnomalyCount *int64 `json:"mediumSeverityAnomalyCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
