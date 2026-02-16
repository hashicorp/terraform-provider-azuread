package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OrganizationSettings{}

type OrganizationSettings struct {
	// Contains the properties that are configured by an administrator as a tenant-level privacy control whether to identify
	// duplicate contacts among a user's contacts list and suggest the user to merge those contacts to have a cleaner
	// contacts list. List contactInsights returns the settings to display or return contact insights in an organization.
	ContactInsights *InsightsSettings `json:"contactInsights,omitempty"`

	// Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived
	// insights, between a user and other items in Microsoft 365, such as documents or sites. List itemInsights returns the
	// settings to display or return item insights in an organization.
	ItemInsights *InsightsSettings `json:"itemInsights,omitempty"`

	MicrosoftApplicationDataAccess *MicrosoftApplicationDataAccessSettings `json:"microsoftApplicationDataAccess,omitempty"`

	// Contains the properties that are configured by an administrator for the visibility of a list of people relevant and
	// working with a user in Microsoft 365. List peopleInsights returns the settings to display or return people insights
	// in an organization.
	PeopleInsights *InsightsSettings `json:"peopleInsights,omitempty"`

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

func (s OrganizationSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OrganizationSettings{}

func (s OrganizationSettings) MarshalJSON() ([]byte, error) {
	type wrapper OrganizationSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OrganizationSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OrganizationSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.organizationSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OrganizationSettings: %+v", err)
	}

	return encoded, nil
}
