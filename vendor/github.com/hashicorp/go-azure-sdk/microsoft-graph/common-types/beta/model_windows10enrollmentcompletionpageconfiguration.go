package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceEnrollmentConfiguration = Windows10EnrollmentCompletionPageConfiguration{}

type Windows10EnrollmentCompletionPageConfiguration struct {
	// When TRUE, allows device reset on installation failure. When false, reset is blocked. The default is false.
	AllowDeviceResetOnInstallFailure *bool `json:"allowDeviceResetOnInstallFailure,omitempty"`

	// When TRUE, allows the user to continue using the device on installation failure. When false, blocks the user on
	// installation failure. The default is false.
	AllowDeviceUseOnInstallFailure *bool `json:"allowDeviceUseOnInstallFailure,omitempty"`

	// When TRUE, allows log collection on installation failure. When false, log collection is not allowed. The default is
	// false.
	AllowLogCollectionOnInstallFailure *bool `json:"allowLogCollectionOnInstallFailure,omitempty"`

	// When TRUE, ESP (Enrollment Status Page) installs all required apps targeted during technician phase and ignores any
	// failures for non-blocking apps. When FALSE, ESP fails on any error during app install. The default is false.
	AllowNonBlockingAppInstallation *bool `json:"allowNonBlockingAppInstallation,omitempty"`

	// When TRUE, blocks user from retrying the setup on installation failure. When false, user is allowed to retry. The
	// default is false.
	BlockDeviceSetupRetryByUser *bool `json:"blockDeviceSetupRetryByUser,omitempty"`

	// The custom error message to show upon installation failure. Max length is 10000. example: 'Setup could not be
	// completed. Please try again or contact your support person for help.'
	CustomErrorMessage nullable.Type[string] `json:"customErrorMessage,omitempty"`

	// When TRUE, disables showing installation progress for first user post enrollment. When false, enables showing
	// progress. The default is false.
	DisableUserStatusTrackingAfterFirstUser *bool `json:"disableUserStatusTrackingAfterFirstUser,omitempty"`

	// The installation progress timeout in minutes. Default is 60 minutes.
	InstallProgressTimeoutInMinutes nullable.Type[int64] `json:"installProgressTimeoutInMinutes,omitempty"`

	// Allows quality updates installation during OOBE
	InstallQualityUpdates *bool `json:"installQualityUpdates,omitempty"`

	// Selected applications to track the installation status. It is in the form of an array of GUIDs.
	SelectedMobileAppIds *[]string `json:"selectedMobileAppIds,omitempty"`

	// When TRUE, shows installation progress to user. When false, hides installation progress. The default is false.
	ShowInstallationProgress *bool `json:"showInstallationProgress,omitempty"`

	// When TRUE, installation progress is tracked for only Autopilot enrollment scenarios. When false, other scenarios are
	// tracked as well. The default is false.
	TrackInstallProgressForAutopilotOnly *bool `json:"trackInstallProgressForAutopilotOnly,omitempty"`

	// Fields inherited from DeviceEnrollmentConfiguration

	// The list of group assignments for the device configuration profile
	Assignments *[]EnrollmentConfigurationAssignment `json:"assignments,omitempty"`

	// Created date time in UTC of the device enrollment configuration
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the device enrollment configuration
	Description nullable.Type[string] `json:"description,omitempty"`

	// Describes the TemplateFamily for the Template entity
	DeviceEnrollmentConfigurationType *DeviceEnrollmentConfigurationType `json:"deviceEnrollmentConfigurationType,omitempty"`

	// The display name of the device enrollment configuration
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Last modified date time in UTC of the device enrollment configuration
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject
	// only to the configuration with the lowest priority value.
	Priority *int64 `json:"priority,omitempty"`

	// Optional role scope tags for the enrollment restrictions.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The version of the device enrollment configuration
	Version *int64 `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s Windows10EnrollmentCompletionPageConfiguration) DeviceEnrollmentConfiguration() BaseDeviceEnrollmentConfigurationImpl {
	return BaseDeviceEnrollmentConfigurationImpl{
		Assignments:                       s.Assignments,
		CreatedDateTime:                   s.CreatedDateTime,
		Description:                       s.Description,
		DeviceEnrollmentConfigurationType: s.DeviceEnrollmentConfigurationType,
		DisplayName:                       s.DisplayName,
		LastModifiedDateTime:              s.LastModifiedDateTime,
		Priority:                          s.Priority,
		RoleScopeTagIds:                   s.RoleScopeTagIds,
		Version:                           s.Version,
		Id:                                s.Id,
		ODataId:                           s.ODataId,
		ODataType:                         s.ODataType,
	}
}

func (s Windows10EnrollmentCompletionPageConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10EnrollmentCompletionPageConfiguration{}

func (s Windows10EnrollmentCompletionPageConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows10EnrollmentCompletionPageConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10EnrollmentCompletionPageConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10EnrollmentCompletionPageConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10EnrollmentCompletionPageConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10EnrollmentCompletionPageConfiguration: %+v", err)
	}

	return encoded, nil
}
