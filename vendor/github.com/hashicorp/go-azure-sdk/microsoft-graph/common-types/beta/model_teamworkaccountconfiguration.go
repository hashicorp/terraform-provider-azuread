package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkAccountConfiguration struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The account used to sync the calendar.
	OnPremisesCalendarSyncConfiguration *TeamworkOnPremisesCalendarSyncConfiguration `json:"onPremisesCalendarSyncConfiguration,omitempty"`

	// The supported client for Teams Rooms devices. The possible values are: unknown, skypeDefaultAndTeams,
	// teamsDefaultAndSkype, skypeOnly, teamsOnly, unknownFutureValue.
	SupportedClient *TeamworkSupportedClient `json:"supportedClient,omitempty"`
}
