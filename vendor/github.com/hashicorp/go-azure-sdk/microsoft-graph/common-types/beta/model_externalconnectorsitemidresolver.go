package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExternalConnectorsUrlToItemResolverBase = ExternalConnectorsItemIdResolver{}

type ExternalConnectorsItemIdResolver struct {
	// Pattern that specifies how to form the ID of the external item that the URL represents. The named groups from the
	// regular expression in urlPattern within the urlMatchInfo can be referenced by inserting the group name inside curly
	// brackets.
	ItemId *string `json:"itemId,omitempty"`

	// Configurations to match and resolve URL.
	UrlMatchInfo *ExternalConnectorsUrlMatchInfo `json:"urlMatchInfo,omitempty"`

	// Fields inherited from ExternalConnectorsUrlToItemResolverBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The priority which defines the sequence in which the urlToItemResolverBase instances are evaluated.
	Priority nullable.Type[int64] `json:"priority,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ExternalConnectorsItemIdResolver) ExternalConnectorsUrlToItemResolverBase() BaseExternalConnectorsUrlToItemResolverBaseImpl {
	return BaseExternalConnectorsUrlToItemResolverBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Priority:  s.Priority,
	}
}

var _ json.Marshaler = ExternalConnectorsItemIdResolver{}

func (s ExternalConnectorsItemIdResolver) MarshalJSON() ([]byte, error) {
	type wrapper ExternalConnectorsItemIdResolver
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalConnectorsItemIdResolver: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalConnectorsItemIdResolver: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalConnectors.itemIdResolver"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalConnectorsItemIdResolver: %+v", err)
	}

	return encoded, nil
}
