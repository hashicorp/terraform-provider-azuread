package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnenoteEntitySchemaObjectModel = OnenotePage{}

type OnenotePage struct {
	// The page's HTML content.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The URL for the page's HTML content. Read-only.
	ContentUrl nullable.Type[string] `json:"contentUrl,omitempty"`

	// The unique identifier of the application that created the page. Read-only.
	CreatedByAppId nullable.Type[string] `json:"createdByAppId,omitempty"`

	// The date and time when the page was last modified. The timestamp represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The indentation level of the page. Read-only.
	Level nullable.Type[int64] `json:"level,omitempty"`

	// Links for opening the page. The oneNoteClientURL link opens the page in the OneNote native client if it 's installed.
	// The oneNoteWebUrl link opens the page in OneNote on the web. Read-only.
	Links *PageLinks `json:"links,omitempty"`

	// The order of the page within its parent section. Read-only.
	Order nullable.Type[int64] `json:"order,omitempty"`

	// The notebook that contains the page. Read-only.
	ParentNotebook *Notebook `json:"parentNotebook,omitempty"`

	// The section that contains the page. Read-only.
	ParentSection *OnenoteSection `json:"parentSection,omitempty"`

	// The title of the page.
	Title nullable.Type[string] `json:"title,omitempty"`

	UserTags *[]string `json:"userTags,omitempty"`

	// Fields inherited from OnenoteEntitySchemaObjectModel

	// The date and time when the page was created. The timestamp represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Fields inherited from OnenoteEntityBaseModel

	// The endpoint where you can get details about the page. Read-only.
	Self nullable.Type[string] `json:"self,omitempty"`

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

func (s OnenotePage) OnenoteEntitySchemaObjectModel() BaseOnenoteEntitySchemaObjectModelImpl {
	return BaseOnenoteEntitySchemaObjectModelImpl{
		CreatedDateTime: s.CreatedDateTime,
		Self:            s.Self,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s OnenotePage) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return BaseOnenoteEntityBaseModelImpl{
		Self:      s.Self,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s OnenotePage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnenotePage{}

func (s OnenotePage) MarshalJSON() ([]byte, error) {
	type wrapper OnenotePage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnenotePage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnenotePage: %+v", err)
	}

	delete(decoded, "contentUrl")
	delete(decoded, "createdByAppId")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "level")
	delete(decoded, "links")
	delete(decoded, "order")
	delete(decoded, "parentNotebook")
	delete(decoded, "parentSection")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onenotePage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnenotePage: %+v", err)
	}

	return encoded, nil
}
