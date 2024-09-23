package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KerberosSignOnSettings struct {
	// The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the
	// connector can present delegated credentials.
	KerberosServicePrincipalName nullable.Type[string] `json:"kerberosServicePrincipalName,omitempty"`

	// The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with
	// different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName,
	// userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
	KerberosSignOnMappingAttributeType *KerberosSignOnMappingAttributeType `json:"kerberosSignOnMappingAttributeType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
