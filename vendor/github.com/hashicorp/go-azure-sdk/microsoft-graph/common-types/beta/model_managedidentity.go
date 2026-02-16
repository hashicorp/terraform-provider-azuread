package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIdentity struct {
	// The ARM resource ID of the Azure resource associated with the managed identity used for sign in.
	AssociatedResourceId nullable.Type[string] `json:"associatedResourceId,omitempty"`

	// The unique ID of the federated token.
	FederatedTokenId nullable.Type[string] `json:"federatedTokenId,omitempty"`

	// The issuer of the federated token.
	FederatedTokenIssuer nullable.Type[string] `json:"federatedTokenIssuer,omitempty"`

	// The possible values are: none, userAssigned, systemAssigned, unknownFutureValue.
	MsiType *MsiType `json:"msiType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
