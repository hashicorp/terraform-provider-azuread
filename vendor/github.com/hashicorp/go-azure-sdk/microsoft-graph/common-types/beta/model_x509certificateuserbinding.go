package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateUserBinding struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The priority of the binding. Microsoft Entra ID uses the binding with the highest priority. This value must be a
	// non-negative integer and unique in the collection of objects in the certificateUserBindings property of an
	// x509CertificateAuthenticationMethodConfiguration object. Required
	Priority *int64 `json:"priority,omitempty"`

	// The affinity level of the username binding rule. The possible values are: low, high, unknownFutureValue.
	TrustAffinityLevel *X509CertificateAffinityLevel `json:"trustAffinityLevel,omitempty"`

	// Defines the Microsoft Entra user property of the user object to use for the binding. The possible values are:
	// userPrincipalName, onPremisesUserPrincipalName, email. Required.
	UserProperty nullable.Type[string] `json:"userProperty,omitempty"`

	// The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name.
	X509CertificateField nullable.Type[string] `json:"x509CertificateField,omitempty"`
}
