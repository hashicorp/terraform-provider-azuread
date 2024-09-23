package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesPublishingSingleSignOn struct {
	// The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
	KerberosSignOnSettings *KerberosSignOnSettings `json:"kerberosSignOnSettings,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos,
	// aadHeaderBased,pingHeaderBased, oAuthToken.
	SingleSignOnMode *SingleSignOnMode `json:"singleSignOnMode,omitempty"`
}
