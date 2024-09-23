package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserActivity{}

type UserActivity struct {
	// Required. URL used to launch the activity in the best native experience represented by the appId. Might launch a
	// web-based app if no native app exists.
	ActivationUrl string `json:"activationUrl"`

	// Required. URL for the domain representing the cross-platform identity mapping for the app. Mapping is stored either
	// as a JSON file hosted on the domain or configurable via Windows Dev Center. The JSON file is named
	// cross-platform-app-identifiers and is hosted at root of your HTTPS domain, either at the top level domain or include
	// a sub domain. For example: https://contoso.com or https://myapp.contoso.com but NOT
	// https://myapp.contoso.com/somepath. You must have a unique file and domain (or sub domain) per cross-platform app
	// identity. For example, a separate file and domain is needed for Word vs. PowerPoint.
	ActivitySourceHost string `json:"activitySourceHost"`

	// Required. The unique activity ID in the context of the app - supplied by caller and immutable thereafter.
	AppActivityId string `json:"appActivityId"`

	// Optional. Short text description of the app used to generate the activity for use in cases when the app is not
	// installed on the user’s local device.
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// Optional. Used in the event the content can be rendered outside of a native or web-based app experience (for example,
	// a pointer to an item in an RSS feed).
	ContentUrl nullable.Type[string] `json:"contentUrl,omitempty"`

	// Set by the server. DateTime in UTC when the object was created on the server.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Set by the server. DateTime in UTC when the object expired on the server.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// Optional. URL used to launch the activity in a web-based app, if available.
	FallbackUrl nullable.Type[string] `json:"fallbackUrl,omitempty"`

	// Optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
	HistoryItems *[]ActivityHistoryItem `json:"historyItems,omitempty"`

	// Set by the server. DateTime in UTC when the object was modified on the server.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
	Status *Status `json:"status,omitempty"`

	// Optional. The timezone in which the user's device used to generate the activity was located at activity creation
	// time; values supplied as Olson IDs in order to support cross-platform representation.
	UserTimezone nullable.Type[string] `json:"userTimezone,omitempty"`

	VisualElements *VisualInfo `json:"visualElements,omitempty"`

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

func (s UserActivity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserActivity{}

func (s UserActivity) MarshalJSON() ([]byte, error) {
	type wrapper UserActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserActivity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserActivity: %+v", err)
	}

	return encoded, nil
}
