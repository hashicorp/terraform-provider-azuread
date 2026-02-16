package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAllowedCombinationsResult struct {
	// Information about why the updateAllowedCombinations action was successful or failed.
	AdditionalInformation nullable.Type[string] `json:"additionalInformation,omitempty"`

	// References to existing Conditional Access policies that use this authentication strength.
	ConditionalAccessReferences *[]string `json:"conditionalAccessReferences,omitempty"`

	// The list of current authentication method combinations allowed by the authentication strength.
	CurrentCombinations *[]AuthenticationMethodModes `json:"currentCombinations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The list of former authentication method combinations allowed by the authentication strength before they were updated
	// through the updateAllowedCombinations action.
	PreviousCombinations *[]AuthenticationMethodModes `json:"previousCombinations,omitempty"`
}
