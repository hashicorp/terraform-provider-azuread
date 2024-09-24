package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegeManagementElevationRequest{}

type PrivilegeManagementElevationRequest struct {
	// Details of the application which is being requested to elevate, allowing the admin to understand the identity of the
	// application. It includes file info such as FilePath, FileHash, FilePublisher, and etc. Returned by default.
	// Read-only.
	ApplicationDetail *ElevationRequestApplicationDetail `json:"applicationDetail,omitempty"`

	// The device name used to initiate the elevation request. For example: 'cotonso-laptop'. Returned by default.
	// Read-only.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The date and time when the elevation request was submitted/created. The value cannot be modified and is automatically
	// populated when the elevation request is submitted/created. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Returned by default. Read-only.
	RequestCreatedDateTime *string `json:"requestCreatedDateTime,omitempty"`

	// Expiration set for the request when it was created, regardless of approved or denied status. For example:
	// '2023-08-03T14:24:22Z'. Returned by default. Returned by default. Read-only.
	RequestExpiryDateTime nullable.Type[string] `json:"requestExpiryDateTime,omitempty"`

	// Justification provided by the end user for the elevation request. For example :'Need to elevate to install microsoft
	// word'. Read-only.
	RequestJustification nullable.Type[string] `json:"requestJustification,omitempty"`

	// The date and time when the elevation request was either submitted/created or approved/denied. The value cannot be
	// modified and is automatically populated. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Returned by default. Read-only.
	RequestLastModifiedDateTime *string `json:"requestLastModifiedDateTime,omitempty"`

	// The Azure Active Directory (AAD) identifier of the end user who is requesting this elevation. For example:
	// 'F1A57311-B9EB-45B7-9415-8555E68EDC9E'. Returned by default. Read-only.
	RequestedByUserId nullable.Type[string] `json:"requestedByUserId,omitempty"`

	// The User Principal Name (UPN) of the end user who requested this elevation. For example: 'user1@contoso.com'.
	// Returned by default. Read-only.
	RequestedByUserPrincipalName nullable.Type[string] `json:"requestedByUserPrincipalName,omitempty"`

	// The Intune Device Identifier of the managed device used to initiate the elevation request. For example:
	// '90F5F6E8-CA09-4811-97F6-4D0DD532D916'. Returned by default. Read-only.
	RequestedOnDeviceId nullable.Type[string] `json:"requestedOnDeviceId,omitempty"`

	// This is the Azure Active Directory (AAD) user id of the administrator who approved or denied the request. For
	// example: 'F1A57311-B9EB-45B7-9415-8555E68EDC9E'. This field would be String.Empty before the request is either
	// approved or denied. Read-only.
	ReviewCompletedByUserId nullable.Type[string] `json:"reviewCompletedByUserId,omitempty"`

	// This is the User Principal Name (UPN) of the administrator who approved or denied the request. For example:
	// 'admin@contoso.com'. This field would be String.Empty before the request is either approved or denied. Read-only.
	ReviewCompletedByUserPrincipalName nullable.Type[string] `json:"reviewCompletedByUserPrincipalName,omitempty"`

	// The DateTime for which the request was approved or denied. For example, midnight UTC on August 3rd, 2023 would look
	// like this: '2023-08-03T00:00:00Z'. Read-only.
	ReviewCompletedDateTime nullable.Type[string] `json:"reviewCompletedDateTime,omitempty"`

	// An optional justification provided by approver at approval or denied time. This field will be String.Empty if
	// approver decides to not provide a justification. For example: 'Run this installer today'
	ReviewerJustification nullable.Type[string] `json:"reviewerJustification,omitempty"`

	// Indicates state of elevation request
	Status *ElevationRequestState `json:"status,omitempty"`

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

func (s PrivilegeManagementElevationRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegeManagementElevationRequest{}

func (s PrivilegeManagementElevationRequest) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegeManagementElevationRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegeManagementElevationRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegeManagementElevationRequest: %+v", err)
	}

	delete(decoded, "applicationDetail")
	delete(decoded, "deviceName")
	delete(decoded, "requestCreatedDateTime")
	delete(decoded, "requestExpiryDateTime")
	delete(decoded, "requestJustification")
	delete(decoded, "requestLastModifiedDateTime")
	delete(decoded, "requestedByUserId")
	delete(decoded, "requestedByUserPrincipalName")
	delete(decoded, "requestedOnDeviceId")
	delete(decoded, "reviewCompletedByUserId")
	delete(decoded, "reviewCompletedByUserPrincipalName")
	delete(decoded, "reviewCompletedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegeManagementElevationRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegeManagementElevationRequest: %+v", err)
	}

	return encoded, nil
}
