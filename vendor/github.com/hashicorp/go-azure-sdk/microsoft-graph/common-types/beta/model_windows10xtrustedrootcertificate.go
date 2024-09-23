package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementResourceAccessProfileBase = Windows10XTrustedRootCertificate{}

type Windows10XTrustedRootCertificate struct {
	// File name to display in UI.
	CertFileName nullable.Type[string] `json:"certFileName,omitempty"`

	// Possible values for the Certificate Destination Store.
	DestinationStore *CertificateDestinationStore `json:"destinationStore,omitempty"`

	// Trusted Root Certificate
	TrustedRootCertificate *string `json:"trustedRootCertificate,omitempty"`

	// Fields inherited from DeviceManagementResourceAccessProfileBase

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceManagementResourceAccessProfileAssignment `json:"assignments,omitempty"`

	// DateTime profile was created
	CreationDateTime nullable.Type[string] `json:"creationDateTime,omitempty"`

	// Profile description
	Description nullable.Type[string] `json:"description,omitempty"`

	// Profile display name
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime profile was last modified
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Scope Tags
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Version of the profile
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

func (s Windows10XTrustedRootCertificate) DeviceManagementResourceAccessProfileBase() BaseDeviceManagementResourceAccessProfileBaseImpl {
	return BaseDeviceManagementResourceAccessProfileBaseImpl{
		Assignments:          s.Assignments,
		CreationDateTime:     s.CreationDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		RoleScopeTagIds:      s.RoleScopeTagIds,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Windows10XTrustedRootCertificate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10XTrustedRootCertificate{}

func (s Windows10XTrustedRootCertificate) MarshalJSON() ([]byte, error) {
	type wrapper Windows10XTrustedRootCertificate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10XTrustedRootCertificate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10XTrustedRootCertificate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10XTrustedRootCertificate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10XTrustedRootCertificate: %+v", err)
	}

	return encoded, nil
}
