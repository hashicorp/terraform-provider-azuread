package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessSessionControls struct {
	// Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session
	// control.
	ApplicationEnforcedRestrictions *ApplicationEnforcedRestrictionsSessionControl `json:"applicationEnforcedRestrictions,omitempty"`

	// Session control to apply cloud app security.
	CloudAppSecurity *CloudAppSecuritySessionControl `json:"cloudAppSecurity,omitempty"`

	// Session control for continuous access evaluation settings.
	ContinuousAccessEvaluation *ContinuousAccessEvaluationSessionControl `json:"continuousAccessEvaluation,omitempty"`

	// Session control that determines whether it's acceptable for Microsoft Entra ID to extend existing sessions based on
	// information collected prior to an outage or not.
	DisableResilienceDefaults nullable.Type[bool] `json:"disableResilienceDefaults,omitempty"`

	// Session control to link to Global Secure Access security profiles or filtering profiles.
	GlobalSecureAccessFilteringProfile *GlobalSecureAccessFilteringProfileSessionControl `json:"globalSecureAccessFilteringProfile,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Session control to define whether to persist cookies or not. All apps should be selected for this session control to
	// work correctly.
	PersistentBrowser *PersistentBrowserSessionControl `json:"persistentBrowser,omitempty"`

	// Session control to require sign in sessions to be bound to a device.
	SecureSignInSession *SecureSignInSessionControl `json:"secureSignInSession,omitempty"`

	// Session control to enforce signin frequency.
	SignInFrequency *SignInFrequencySessionControl `json:"signInFrequency,omitempty"`
}
