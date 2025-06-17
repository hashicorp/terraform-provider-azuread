package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = FederatedIdentityCredential{}

type FederatedIdentityCredential struct {
	// The audience that can appear in the external token. This field is mandatory and should be set to
	// api://AzureADTokenExchange for Microsoft Entra ID. It says what Microsoft identity platform should accept in the aud
	// claim in the incoming token. This value represents Microsoft Entra ID in your external identity provider and has no
	// fixed value across identity providers - you may need to create a new application registration in your identity
	// provider to serve as the audience of this token. This field can only accept a single value and has a limit of 600
	// characters. Required.
	Audiences []string `json:"audiences"`

	// Nullable. Defaults to null if not set. Enables the use of claims matching expressions against specified claims. If
	// claimsMatchingExpression is defined, subject must be null. For the list of supported expression syntax and claims,
	// visit the Flexible FIC reference.
	ClaimsMatchingExpression *FederatedIdentityExpression `json:"claimsMatchingExpression,omitempty"`

	// The un-validated, user-provided description of the federated identity credential. It has a limit of 600 characters.
	// Optional.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The URL of the external identity provider and must match the issuer claim of the external token being exchanged. The
	// combination of the values of issuer and subject must be unique on the app. It has a limit of 600 characters.
	// Required.
	Issuer string `json:"issuer"`

	// The unique identifier for the federated identity credential, which has a limit of 120 characters and must be URL
	// friendly. It is immutable once created. Alternate key. Required. Not nullable. Supports $filter (eq).
	Name string `json:"name"`

	// Nullable. Defaults to null if not set. The identifier of the external software workload within the external identity
	// provider. Like the audience value, it has no fixed format, as each identity provider uses their own - sometimes a
	// GUID, sometimes a colon delimited identifier, sometimes arbitrary strings. The value here must match the sub claim
	// within the token presented to Microsoft Entra ID. The combination of issuer and subject must be unique on the app. It
	// has a limit of 600 characters. If subject is defined, claimsMatchingExpression must be null. Supports $filter (eq).
	Subject nullable.Type[string] `json:"subject,omitempty"`

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

func (s FederatedIdentityCredential) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = FederatedIdentityCredential{}

func (s FederatedIdentityCredential) MarshalJSON() ([]byte, error) {
	type wrapper FederatedIdentityCredential
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FederatedIdentityCredential: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FederatedIdentityCredential: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.federatedIdentityCredential"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FederatedIdentityCredential: %+v", err)
	}

	return encoded, nil
}
