package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTimeDeviceMembershipTargetResult struct {
	// A list of validation status of the memberships targetted to profile. This collection can contain a maximum of 1
	// elements.
	EnrollmentTimeDeviceMembershipTargetValidationStatuses *[]EnrollmentTimeDeviceMembershipTargetStatus `json:"enrollmentTimeDeviceMembershipTargetValidationStatuses,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates if validations succeeded for the device membership target. When 'true', the device membership target
	// validation found no issues. When 'false', the device membership target validation found issues. default - false
	ValidationSucceeded *bool `json:"validationSucceeded,omitempty"`
}
