package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AppConsentRequest{}

type AppConsentRequest struct {
	// The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby.
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// The identifier of the application. Required. Supports $filter (eq only) and $orderby.
	AppId string `json:"appId"`

	// A list of pending scopes waiting for approval. Required.
	PendingScopes []AppConsentRequestScope `json:"pendingScopes"`

	// A list of pending user consent requests. Supports $filter (eq).
	UserConsentRequests *[]UserConsentRequest `json:"userConsentRequests,omitempty"`

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

func (s AppConsentRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AppConsentRequest{}

func (s AppConsentRequest) MarshalJSON() ([]byte, error) {
	type wrapper AppConsentRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppConsentRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppConsentRequest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appConsentRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppConsentRequest: %+v", err)
	}

	return encoded, nil
}
