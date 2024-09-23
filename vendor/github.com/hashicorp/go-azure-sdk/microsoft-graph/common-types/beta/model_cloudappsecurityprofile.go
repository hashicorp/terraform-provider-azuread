package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudAppSecurityProfile{}

type CloudAppSecurityProfile struct {
	AzureSubscriptionId    nullable.Type[string]           `json:"azureSubscriptionId,omitempty"`
	AzureTenantId          nullable.Type[string]           `json:"azureTenantId,omitempty"`
	CreatedDateTime        nullable.Type[string]           `json:"createdDateTime,omitempty"`
	DeploymentPackageUrl   nullable.Type[string]           `json:"deploymentPackageUrl,omitempty"`
	DestinationServiceName nullable.Type[string]           `json:"destinationServiceName,omitempty"`
	IsSigned               nullable.Type[bool]             `json:"isSigned,omitempty"`
	LastModifiedDateTime   nullable.Type[string]           `json:"lastModifiedDateTime,omitempty"`
	Manifest               nullable.Type[string]           `json:"manifest,omitempty"`
	Name                   nullable.Type[string]           `json:"name,omitempty"`
	PermissionsRequired    *ApplicationPermissionsRequired `json:"permissionsRequired,omitempty"`
	Platform               nullable.Type[string]           `json:"platform,omitempty"`
	PolicyName             nullable.Type[string]           `json:"policyName,omitempty"`
	Publisher              nullable.Type[string]           `json:"publisher,omitempty"`
	RiskScore              nullable.Type[string]           `json:"riskScore,omitempty"`
	Tags                   *[]string                       `json:"tags,omitempty"`
	Type                   nullable.Type[string]           `json:"type,omitempty"`
	VendorInformation      *SecurityVendorInformation      `json:"vendorInformation,omitempty"`

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

func (s CloudAppSecurityProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudAppSecurityProfile{}

func (s CloudAppSecurityProfile) MarshalJSON() ([]byte, error) {
	type wrapper CloudAppSecurityProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudAppSecurityProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudAppSecurityProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudAppSecurityProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudAppSecurityProfile: %+v", err)
	}

	return encoded, nil
}
