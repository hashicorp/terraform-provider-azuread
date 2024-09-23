package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContextDetails struct {
	// Nullable. Specifies the text to use in a user experience to display a link the the associated plannerPlanContext. If
	// null, applications should display the link with a custom text based on the displayLinkType property.
	CustomLinkText nullable.Type[string] `json:"customLinkText,omitempty"`

	// Specifies how an application should display the link to the associated plannerPlanContext. Applications may choose to
	// provide customized text, description, icons, or other experiences based on the type of the link. Possible values are:
	// teamsTab, sharePointPage, meetingNotes, loopPage, project, other, unknownFutureValue.
	DisplayLinkType *PlannerPlanContextType `json:"displayLinkType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the state of the associated plannerPlanContext.
	State *PlannerContextState `json:"state,omitempty"`

	// URL of the user experience represented by the associated plannerPlanContext.
	Url nullable.Type[string] `json:"url,omitempty"`
}
