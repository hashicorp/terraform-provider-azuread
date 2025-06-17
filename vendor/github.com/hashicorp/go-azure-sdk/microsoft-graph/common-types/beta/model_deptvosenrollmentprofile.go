package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EnrollmentProfile = DepTvOSEnrollmentProfile{}

type DepTvOSEnrollmentProfile struct {

	// Fields inherited from EnrollmentProfile

	// Configuration endpoint url to use for Enrollment
	ConfigurationEndpointUrl nullable.Type[string] `json:"configurationEndpointUrl,omitempty"`

	// Description of the profile
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the profile
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
	EnableAuthenticationViaCompanyPortal *bool `json:"enableAuthenticationViaCompanyPortal,omitempty"`

	// Indicates that Company Portal is required on setup assistant enrolled devices
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`

	// Indicates if the profile requires user authentication
	RequiresUserAuthentication *bool `json:"requiresUserAuthentication,omitempty"`

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

func (s DepTvOSEnrollmentProfile) EnrollmentProfile() BaseEnrollmentProfileImpl {
	return BaseEnrollmentProfileImpl{
		ConfigurationEndpointUrl:             s.ConfigurationEndpointUrl,
		Description:                          s.Description,
		DisplayName:                          s.DisplayName,
		EnableAuthenticationViaCompanyPortal: s.EnableAuthenticationViaCompanyPortal,
		RequireCompanyPortalOnSetupAssistantEnrolledDevices: s.RequireCompanyPortalOnSetupAssistantEnrolledDevices,
		RequiresUserAuthentication:                          s.RequiresUserAuthentication,
		Id:                                                  s.Id,
		ODataId:                                             s.ODataId,
		ODataType:                                           s.ODataType,
	}
}

func (s DepTvOSEnrollmentProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DepTvOSEnrollmentProfile{}

func (s DepTvOSEnrollmentProfile) MarshalJSON() ([]byte, error) {
	type wrapper DepTvOSEnrollmentProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DepTvOSEnrollmentProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DepTvOSEnrollmentProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.depTvOSEnrollmentProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DepTvOSEnrollmentProfile: %+v", err)
	}

	return encoded, nil
}
