package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileLobApp interface {
	Entity
	MobileApp
	MobileLobApp() BaseMobileLobAppImpl
}

var _ MobileLobApp = BaseMobileLobAppImpl{}

type BaseMobileLobAppImpl struct {
	// The internal committed content version.
	CommittedContentVersion nullable.Type[string] `json:"committedContentVersion,omitempty"`

	// The list of content versions for this app. This property is read-only.
	ContentVersions *[]MobileAppContent `json:"contentVersions,omitempty"`

	// The name of the main Lob application file.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The total size, including all uploaded files. This property is read-only.
	Size *int64 `json:"size,omitempty"`

	// Fields inherited from MobileApp

	// The list of group assignments for this mobile app.
	Assignments *[]MobileAppAssignment `json:"assignments,omitempty"`

	// The list of categories for this app.
	Categories *[]MobileAppCategory `json:"categories,omitempty"`

	// The date and time the app was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The total number of dependencies the child app has.
	DependentAppCount *int64 `json:"dependentAppCount,omitempty"`

	// The description of the app.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The developer of the app.
	Developer nullable.Type[string] `json:"developer,omitempty"`

	// The admin provided or imported title of the app.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The more information Url.
	InformationUrl nullable.Type[string] `json:"informationUrl,omitempty"`

	// The value indicating whether the app is assigned to at least one group.
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`

	// The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *MimeContent `json:"largeIcon,omitempty"`

	// The date and time the app was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Notes for the app.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The owner of the app.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// The privacy statement Url.
	PrivacyInformationUrl nullable.Type[string] `json:"privacyInformationUrl,omitempty"`

	// The publisher of the app.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Indicates the publishing state of an app.
	PublishingState *MobileAppPublishingState `json:"publishingState,omitempty"`

	// List of relationships for this mobile app.
	Relationships *[]MobileAppRelationship `json:"relationships,omitempty"`

	// List of scope tag ids for this mobile app.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The total number of apps this app is directly or indirectly superseded by. This property is read-only.
	SupersededAppCount *int64 `json:"supersededAppCount,omitempty"`

	// The total number of apps this app directly or indirectly supersedes. This property is read-only.
	SupersedingAppCount *int64 `json:"supersedingAppCount,omitempty"`

	// The upload state.
	UploadState *int64 `json:"uploadState,omitempty"`

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

func (s BaseMobileLobAppImpl) MobileLobApp() BaseMobileLobAppImpl {
	return s
}

func (s BaseMobileLobAppImpl) MobileApp() BaseMobileAppImpl {
	return BaseMobileAppImpl{
		Assignments:           s.Assignments,
		Categories:            s.Categories,
		CreatedDateTime:       s.CreatedDateTime,
		DependentAppCount:     s.DependentAppCount,
		Description:           s.Description,
		Developer:             s.Developer,
		DisplayName:           s.DisplayName,
		InformationUrl:        s.InformationUrl,
		IsAssigned:            s.IsAssigned,
		IsFeatured:            s.IsFeatured,
		LargeIcon:             s.LargeIcon,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Notes:                 s.Notes,
		Owner:                 s.Owner,
		PrivacyInformationUrl: s.PrivacyInformationUrl,
		Publisher:             s.Publisher,
		PublishingState:       s.PublishingState,
		Relationships:         s.Relationships,
		RoleScopeTagIds:       s.RoleScopeTagIds,
		SupersededAppCount:    s.SupersededAppCount,
		SupersedingAppCount:   s.SupersedingAppCount,
		UploadState:           s.UploadState,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s BaseMobileLobAppImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MobileLobApp = RawMobileLobAppImpl{}

// RawMobileLobAppImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMobileLobAppImpl struct {
	mobileLobApp BaseMobileLobAppImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawMobileLobAppImpl) MobileLobApp() BaseMobileLobAppImpl {
	return s.mobileLobApp
}

func (s RawMobileLobAppImpl) MobileApp() BaseMobileAppImpl {
	return s.mobileLobApp.MobileApp()
}

func (s RawMobileLobAppImpl) Entity() BaseEntityImpl {
	return s.mobileLobApp.Entity()
}

var _ json.Marshaler = BaseMobileLobAppImpl{}

func (s BaseMobileLobAppImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMobileLobAppImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMobileLobAppImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMobileLobAppImpl: %+v", err)
	}

	delete(decoded, "size")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileLobApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMobileLobAppImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseMobileLobAppImpl{}

func (s *BaseMobileLobAppImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CommittedContentVersion nullable.Type[string]     `json:"committedContentVersion,omitempty"`
		ContentVersions         *[]MobileAppContent       `json:"contentVersions,omitempty"`
		FileName                nullable.Type[string]     `json:"fileName,omitempty"`
		Size                    *int64                    `json:"size,omitempty"`
		Assignments             *[]MobileAppAssignment    `json:"assignments,omitempty"`
		Categories              *[]MobileAppCategory      `json:"categories,omitempty"`
		CreatedDateTime         *string                   `json:"createdDateTime,omitempty"`
		DependentAppCount       *int64                    `json:"dependentAppCount,omitempty"`
		Description             nullable.Type[string]     `json:"description,omitempty"`
		Developer               nullable.Type[string]     `json:"developer,omitempty"`
		DisplayName             nullable.Type[string]     `json:"displayName,omitempty"`
		InformationUrl          nullable.Type[string]     `json:"informationUrl,omitempty"`
		IsAssigned              *bool                     `json:"isAssigned,omitempty"`
		IsFeatured              *bool                     `json:"isFeatured,omitempty"`
		LargeIcon               *MimeContent              `json:"largeIcon,omitempty"`
		LastModifiedDateTime    *string                   `json:"lastModifiedDateTime,omitempty"`
		Notes                   nullable.Type[string]     `json:"notes,omitempty"`
		Owner                   nullable.Type[string]     `json:"owner,omitempty"`
		PrivacyInformationUrl   nullable.Type[string]     `json:"privacyInformationUrl,omitempty"`
		Publisher               nullable.Type[string]     `json:"publisher,omitempty"`
		PublishingState         *MobileAppPublishingState `json:"publishingState,omitempty"`
		RoleScopeTagIds         *[]string                 `json:"roleScopeTagIds,omitempty"`
		SupersededAppCount      *int64                    `json:"supersededAppCount,omitempty"`
		SupersedingAppCount     *int64                    `json:"supersedingAppCount,omitempty"`
		UploadState             *int64                    `json:"uploadState,omitempty"`
		Id                      *string                   `json:"id,omitempty"`
		ODataId                 *string                   `json:"@odata.id,omitempty"`
		ODataType               *string                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CommittedContentVersion = decoded.CommittedContentVersion
	s.ContentVersions = decoded.ContentVersions
	s.FileName = decoded.FileName
	s.Size = decoded.Size
	s.Assignments = decoded.Assignments
	s.Categories = decoded.Categories
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DependentAppCount = decoded.DependentAppCount
	s.Description = decoded.Description
	s.Developer = decoded.Developer
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.InformationUrl = decoded.InformationUrl
	s.IsAssigned = decoded.IsAssigned
	s.IsFeatured = decoded.IsFeatured
	s.LargeIcon = decoded.LargeIcon
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Notes = decoded.Notes
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Owner = decoded.Owner
	s.PrivacyInformationUrl = decoded.PrivacyInformationUrl
	s.Publisher = decoded.Publisher
	s.PublishingState = decoded.PublishingState
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.SupersededAppCount = decoded.SupersededAppCount
	s.SupersedingAppCount = decoded.SupersedingAppCount
	s.UploadState = decoded.UploadState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseMobileLobAppImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["relationships"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Relationships into list []json.RawMessage: %+v", err)
		}

		output := make([]MobileAppRelationship, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMobileAppRelationshipImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Relationships' for 'BaseMobileLobAppImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Relationships = &output
	}

	return nil
}

func UnmarshalMobileLobAppImplementation(input []byte) (MobileLobApp, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileLobApp into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidLobApp") {
		var out AndroidLobApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidLobApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosLobApp") {
		var out IosLobApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosLobApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSDmgApp") {
		var out MacOSDmgApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSDmgApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSLobApp") {
		var out MacOSLobApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSLobApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSPkgApp") {
		var out MacOSPkgApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSPkgApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobApp") {
		var out Win32LobApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAppX") {
		var out WindowsAppX
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAppX: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsMobileMSI") {
		var out WindowsMobileMSI
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsMobileMSI: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81AppX") {
		var out WindowsPhone81AppX
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81AppX: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhoneXAP") {
		var out WindowsPhoneXAP
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhoneXAP: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUniversalAppX") {
		var out WindowsUniversalAppX
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUniversalAppX: %+v", err)
		}
		return out, nil
	}

	var parent BaseMobileLobAppImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMobileLobAppImpl: %+v", err)
	}

	return RawMobileLobAppImpl{
		mobileLobApp: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
