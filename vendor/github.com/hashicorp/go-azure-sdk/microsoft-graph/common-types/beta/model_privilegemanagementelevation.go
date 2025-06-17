package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegeManagementElevation{}

type PrivilegeManagementElevation struct {
	// The certificate payload of the application. This is computed by hashing the certificate information on the client.
	// Example: 32c220482c68413fbf8290e3b1e49b0a85901cfcd62ab0738760568a2a6e8a50
	CertificatePayload nullable.Type[string] `json:"certificatePayload,omitempty"`

	// The company name of the application. This value is set by the creator of the application. Example: Microsoft
	// Corporation
	CompanyName nullable.Type[string] `json:"companyName,omitempty"`

	// The Intune deviceId. Unique identifier for the managed device. Example: 92ce5047-9553-4731-817f-9b401a999a1b
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The name associated with the device in the intune database. Example: JOHNDOE-LAPTOP.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Indicates the type of elevation occured
	ElevationType *PrivilegeManagementElevationType `json:"elevationType,omitempty"`

	// The date and time when the application was elevated. Example:2014-01-01T00:00:00Z
	EventDateTime *string `json:"eventDateTime,omitempty"`

	// The file description of the application. This value is set by the creator of the application. Example: Editor of
	// multiple coding languages.
	FileDescription nullable.Type[string] `json:"fileDescription,omitempty"`

	// The full file path of the application including the filename and file extension. Example: C:/Program Files/vscode.exe
	FilePath nullable.Type[string] `json:"filePath,omitempty"`

	// The version of the application. This value is set by the creator of the application. Example: 6.2211.1035.1000
	FileVersion nullable.Type[string] `json:"fileVersion,omitempty"`

	// The sha256 hash of the application. Example: 32c220482c68413fbf8290e3b1e49b0a85901cfcd62ab0738760568a2a6e8a57
	Hash nullable.Type[string] `json:"hash,omitempty"`

	// The internal name of the application. This value is set by the creator of the application. Example: VS code
	InternalName nullable.Type[string] `json:"internalName,omitempty"`

	// The justification to elevate the application. This is an input by the user when the privilegeManagementElevationType
	// is of type userConfirmedElevation or support approved elevation. This will be null in all other scenarios. The length
	// is capped at 256 char, enforced on the client side. Example: To install debug tool..
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// The name of parent process associated with the elevated process. This is always populated for both parent and child
	// process types
	ParentProcessName nullable.Type[string] `json:"parentProcessName,omitempty"`

	// Unique Identifier of the policy configured to run the application with elevated access
	PolicyId nullable.Type[string] `json:"policyId,omitempty"`

	// The name of the policy configured to run the application in elevated access
	PolicyName nullable.Type[string] `json:"policyName,omitempty"`

	// Indicates the type of elevated process
	ProcessType *PrivilegeManagementProcessType `json:"processType,omitempty"`

	// The product name of the application. This value is set by the creator of the application. Example: Visual Studio
	ProductName nullable.Type[string] `json:"productName,omitempty"`

	// The result of the elevation action with 0 being success, and everything else being exit code if the elevation was
	// unsuccessful. The value will always be 0 on all unmanaged elevation. Example: 0. Valid values 0 to 2147483647
	Result *int64 `json:"result,omitempty"`

	// Unique identifier of the rule configured to run the application with elevated access
	RuleId nullable.Type[string] `json:"ruleId,omitempty"`

	// To identify if the elevation is initiated by system or user interaction
	SystemInitiatedElevation *bool `json:"systemInitiatedElevation,omitempty"`

	// The User Principal Name of the user who performed the elevation. Example: john@domain.com
	Upn nullable.Type[string] `json:"upn,omitempty"`

	// The type of user account on Windows that was used to performed the elevation.
	UserType *PrivilegeManagementEndUserType `json:"userType,omitempty"`

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

func (s PrivilegeManagementElevation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegeManagementElevation{}

func (s PrivilegeManagementElevation) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegeManagementElevation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegeManagementElevation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegeManagementElevation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegeManagementElevation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegeManagementElevation: %+v", err)
	}

	return encoded, nil
}
