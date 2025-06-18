package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAutopilotConfiguration struct {
	// Indicates the number of minutes allowed for the Autopilot application to apply the device preparation profile (DPP)
	// configurations to the device. If the Autopilot application doesn't finish within the specified time
	// (applicationTimeoutInMinutes), the application error is added to the statusDetail property of the cloudPC object. The
	// supported value is an integer between 10 and 360. Required.
	ApplicationTimeoutInMinutes int64 `json:"applicationTimeoutInMinutes"`

	// The unique identifier (ID) of the Autopilot device preparation profile (DPP) that links a Windows Autopilot device
	// preparation policy to ensure that devices are ready for users after provisioning. Required.
	DevicePreparationProfileId string `json:"devicePreparationProfileId"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether the access to the device is allowed when the application of Autopilot device preparation profile
	// (DPP) configurations fails or times out. If true, the status of the device is failed and the device is unable to
	// access; otherwise, the status of the device is provisionedWithWarnings and the device is allowed to access. The
	// default value is false. Required.
	OnFailureDeviceAccessDenied bool `json:"onFailureDeviceAccessDenied"`
}
