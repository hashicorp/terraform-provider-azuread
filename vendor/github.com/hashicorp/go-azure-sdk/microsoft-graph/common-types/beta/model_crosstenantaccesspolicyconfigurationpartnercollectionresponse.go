package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = CrossTenantAccessPolicyConfigurationPartnerCollectionResponse{}

type CrossTenantAccessPolicyConfigurationPartnerCollectionResponse struct {
	Value *[]CrossTenantAccessPolicyConfigurationPartner `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CrossTenantAccessPolicyConfigurationPartnerCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = CrossTenantAccessPolicyConfigurationPartnerCollectionResponse{}

func (s CrossTenantAccessPolicyConfigurationPartnerCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper CrossTenantAccessPolicyConfigurationPartnerCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CrossTenantAccessPolicyConfigurationPartnerCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CrossTenantAccessPolicyConfigurationPartnerCollectionResponse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.crossTenantAccessPolicyConfigurationPartnerCollectionResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CrossTenantAccessPolicyConfigurationPartnerCollectionResponse: %+v", err)
	}

	return encoded, nil
}
