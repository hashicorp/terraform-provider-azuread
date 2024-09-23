package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDashboardCardContentSource struct {
	// The configuration for the bot source. Required if sourceType is set to bot.
	BotConfiguration *TeamsAppDashboardCardBotConfiguration `json:"botConfiguration,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the type of source that powers the content of the dashboard card. The possible values are: bot,
	// unknownFutureValue.
	SourceType *TeamsAppDashboardCardSourceType `json:"sourceType,omitempty"`
}
