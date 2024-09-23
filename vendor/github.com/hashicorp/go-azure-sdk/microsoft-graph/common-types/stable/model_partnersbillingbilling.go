package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PartnersBillingBilling{}

type PartnersBillingBilling struct {
	// Represents metadata for the exported data.
	Manifests *[]PartnersBillingManifest `json:"manifests,omitempty"`

	// Represents an operation to export the billing data of a partner.
	Operations *[]PartnersBillingOperation `json:"operations,omitempty"`

	Reconciliation *PartnersBillingBillingReconciliation `json:"reconciliation,omitempty"`
	Usage          *PartnersBillingAzureUsage            `json:"usage,omitempty"`

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

func (s PartnersBillingBilling) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnersBillingBilling{}

func (s PartnersBillingBilling) MarshalJSON() ([]byte, error) {
	type wrapper PartnersBillingBilling
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnersBillingBilling: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnersBillingBilling: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partners.billing.billing"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnersBillingBilling: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PartnersBillingBilling{}

func (s *PartnersBillingBilling) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Manifests      *[]PartnersBillingManifest            `json:"manifests,omitempty"`
		Reconciliation *PartnersBillingBillingReconciliation `json:"reconciliation,omitempty"`
		Usage          *PartnersBillingAzureUsage            `json:"usage,omitempty"`
		Id             *string                               `json:"id,omitempty"`
		ODataId        *string                               `json:"@odata.id,omitempty"`
		ODataType      *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Manifests = decoded.Manifests
	s.Reconciliation = decoded.Reconciliation
	s.Usage = decoded.Usage
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PartnersBillingBilling into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["operations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Operations into list []json.RawMessage: %+v", err)
		}

		output := make([]PartnersBillingOperation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPartnersBillingOperationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Operations' for 'PartnersBillingBilling': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Operations = &output
	}

	return nil
}
