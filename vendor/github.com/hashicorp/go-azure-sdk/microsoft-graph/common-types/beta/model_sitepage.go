package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseSitePage = SitePage{}

type SitePage struct {
	// Indicates the layout of the content in a given SharePoint page, including horizontal sections and vertical sections.
	CanvasLayout *CanvasLayout `json:"canvasLayout,omitempty"`

	// Indicates the promotion kind of the sitePage. The possible values are: microsoftReserved, page, newsPost,
	// unknownFutureValue.
	PromotionKind *PagePromotionType `json:"promotionKind,omitempty"`

	// Reactions information for the page.
	Reactions *ReactionsFacet `json:"reactions,omitempty"`

	// Determines whether or not to show comments at the bottom of the page.
	ShowComments nullable.Type[bool] `json:"showComments,omitempty"`

	// Determines whether or not to show recommended pages at the bottom of the page.
	ShowRecommendedPages nullable.Type[bool] `json:"showRecommendedPages,omitempty"`

	// Url of the sitePage's thumbnail image
	ThumbnailWebUrl nullable.Type[string] `json:"thumbnailWebUrl,omitempty"`

	// Title area on the SharePoint page.
	TitleArea *TitleArea `json:"titleArea,omitempty"`

	// Collection of webparts on the SharePoint page.
	WebParts *[]WebPart `json:"webParts,omitempty"`

	// Fields inherited from BaseSitePage

	// The name of the page layout of the page. The possible values are: microsoftReserved, article, home,
	// unknownFutureValue, newsLink. Use the Prefer: include-unknown-enum-members request header to get the following value
	// in this evolvable enum: newsLink.
	PageLayout *PageLayoutType `json:"pageLayout,omitempty"`

	// The publishing status and the MM.mm version of the page.
	PublishingState *PublicationFacet `json:"publishingState,omitempty"`

	// Title of the sitePage.
	Title nullable.Type[string] `json:"title,omitempty"`

	// Fields inherited from BaseItem

	// Identity of the user, device, or application that created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	CreatedByUser *User `json:"createdByUser,omitempty"`

	// Date and time of item creation. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the item.
	Description nullable.Type[string] `json:"description,omitempty"`

	// ETag for the item. Read-only.
	ETag nullable.Type[string] `json:"eTag,omitempty"`

	// Identity of the user, device, and application that last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	LastModifiedByUser *User `json:"lastModifiedByUser,omitempty"`

	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The name of the item. Read-write.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`

	// URL that either displays the resource in the browser (for Office file formats), or is a direct link to the file (for
	// other formats). Read-only.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

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

func (s SitePage) BaseSitePage() BaseBaseSitePageImpl {
	return BaseBaseSitePageImpl{
		PageLayout:           s.PageLayout,
		PublishingState:      s.PublishingState,
		Title:                s.Title,
		CreatedBy:            s.CreatedBy,
		CreatedByUser:        s.CreatedByUser,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		ETag:                 s.ETag,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedByUser:   s.LastModifiedByUser,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Name:                 s.Name,
		ParentReference:      s.ParentReference,
		WebUrl:               s.WebUrl,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SitePage) BaseItem() BaseBaseItemImpl {
	return BaseBaseItemImpl{
		CreatedBy:            s.CreatedBy,
		CreatedByUser:        s.CreatedByUser,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		ETag:                 s.ETag,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedByUser:   s.LastModifiedByUser,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Name:                 s.Name,
		ParentReference:      s.ParentReference,
		WebUrl:               s.WebUrl,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SitePage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SitePage{}

func (s SitePage) MarshalJSON() ([]byte, error) {
	type wrapper SitePage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SitePage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SitePage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sitePage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SitePage: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SitePage{}

func (s *SitePage) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CanvasLayout         *CanvasLayout         `json:"canvasLayout,omitempty"`
		PromotionKind        *PagePromotionType    `json:"promotionKind,omitempty"`
		Reactions            *ReactionsFacet       `json:"reactions,omitempty"`
		ShowComments         nullable.Type[bool]   `json:"showComments,omitempty"`
		ShowRecommendedPages nullable.Type[bool]   `json:"showRecommendedPages,omitempty"`
		ThumbnailWebUrl      nullable.Type[string] `json:"thumbnailWebUrl,omitempty"`
		TitleArea            *TitleArea            `json:"titleArea,omitempty"`
		PageLayout           *PageLayoutType       `json:"pageLayout,omitempty"`
		PublishingState      *PublicationFacet     `json:"publishingState,omitempty"`
		Title                nullable.Type[string] `json:"title,omitempty"`
		CreatedByUser        *User                 `json:"createdByUser,omitempty"`
		CreatedDateTime      *string               `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string] `json:"description,omitempty"`
		ETag                 nullable.Type[string] `json:"eTag,omitempty"`
		LastModifiedByUser   *User                 `json:"lastModifiedByUser,omitempty"`
		LastModifiedDateTime *string               `json:"lastModifiedDateTime,omitempty"`
		Name                 nullable.Type[string] `json:"name,omitempty"`
		ParentReference      *ItemReference        `json:"parentReference,omitempty"`
		WebUrl               nullable.Type[string] `json:"webUrl,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CanvasLayout = decoded.CanvasLayout
	s.PromotionKind = decoded.PromotionKind
	s.Reactions = decoded.Reactions
	s.ShowComments = decoded.ShowComments
	s.ShowRecommendedPages = decoded.ShowRecommendedPages
	s.ThumbnailWebUrl = decoded.ThumbnailWebUrl
	s.TitleArea = decoded.TitleArea
	s.CreatedByUser = decoded.CreatedByUser
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.ETag = decoded.ETag
	s.Id = decoded.Id
	s.LastModifiedByUser = decoded.LastModifiedByUser
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PageLayout = decoded.PageLayout
	s.ParentReference = decoded.ParentReference
	s.PublishingState = decoded.PublishingState
	s.Title = decoded.Title
	s.WebUrl = decoded.WebUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SitePage into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SitePage': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SitePage': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	if v, ok := temp["webParts"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling WebParts into list []json.RawMessage: %+v", err)
		}

		output := make([]WebPart, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWebPartImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'WebParts' for 'SitePage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.WebParts = &output
	}

	return nil
}
