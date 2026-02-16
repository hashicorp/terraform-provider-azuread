package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessAssociation = NetworkaccessAssociatedBranch{}

type NetworkaccessAssociatedBranch struct {
	// Identifier for the branch.
	BranchId *string `json:"branchId,omitempty"`

	// Fields inherited from NetworkaccessAssociation

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessAssociatedBranch) NetworkaccessAssociation() BaseNetworkaccessAssociationImpl {
	return BaseNetworkaccessAssociationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessAssociatedBranch{}

func (s NetworkaccessAssociatedBranch) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessAssociatedBranch
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessAssociatedBranch: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessAssociatedBranch: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.associatedBranch"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessAssociatedBranch: %+v", err)
	}

	return encoded, nil
}
