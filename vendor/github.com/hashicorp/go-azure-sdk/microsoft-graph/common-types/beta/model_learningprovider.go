package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = LearningProvider{}

type LearningProvider struct {
	// The display name that appears in Viva Learning. Required.
	DisplayName string `json:"displayName"`

	// Indicates whether a provider can ingest learning course activity records. The default value is false. Set to true to
	// make learningCourseActivities available for this provider.
	IsCourseActivitySyncEnabled nullable.Type[bool] `json:"isCourseActivitySyncEnabled,omitempty"`

	// Learning catalog items for the provider.
	LearningContents *[]LearningContent `json:"learningContents,omitempty"`

	LearningCourseActivities *[]LearningCourseActivity `json:"learningCourseActivities,omitempty"`

	// Authentication URL to access the courses for the provider. Optional.
	LoginWebUrl nullable.Type[string] `json:"loginWebUrl,omitempty"`

	// The long logo URL for the dark mode that needs to be a publicly accessible image. This image would be saved to the
	// blob storage of Viva Learning for rendering within the Viva Learning app. Required.
	LongLogoWebUrlForDarkTheme string `json:"longLogoWebUrlForDarkTheme"`

	// The long logo URL for the light mode that needs to be a publicly accessible image. This image would be saved to the
	// blob storage of Viva Learning for rendering within the Viva Learning app. Required.
	LongLogoWebUrlForLightTheme string `json:"longLogoWebUrlForLightTheme"`

	// The square logo URL for the dark mode that needs to be a publicly accessible image. This image would be saved to the
	// blob storage of Viva Learning for rendering within the Viva Learning app. Required.
	SquareLogoWebUrlForDarkTheme string `json:"squareLogoWebUrlForDarkTheme"`

	// The square logo URL for the light mode that needs to be a publicly accessible image. This image would be saved to the
	// blob storage of Viva Learning for rendering within the Viva Learning app. Required.
	SquareLogoWebUrlForLightTheme string `json:"squareLogoWebUrlForLightTheme"`

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

func (s LearningProvider) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LearningProvider{}

func (s LearningProvider) MarshalJSON() ([]byte, error) {
	type wrapper LearningProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LearningProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LearningProvider: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.learningProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LearningProvider: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &LearningProvider{}

func (s *LearningProvider) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName                   string                `json:"displayName"`
		IsCourseActivitySyncEnabled   nullable.Type[bool]   `json:"isCourseActivitySyncEnabled,omitempty"`
		LearningContents              *[]LearningContent    `json:"learningContents,omitempty"`
		LoginWebUrl                   nullable.Type[string] `json:"loginWebUrl,omitempty"`
		LongLogoWebUrlForDarkTheme    string                `json:"longLogoWebUrlForDarkTheme"`
		LongLogoWebUrlForLightTheme   string                `json:"longLogoWebUrlForLightTheme"`
		SquareLogoWebUrlForDarkTheme  string                `json:"squareLogoWebUrlForDarkTheme"`
		SquareLogoWebUrlForLightTheme string                `json:"squareLogoWebUrlForLightTheme"`
		Id                            *string               `json:"id,omitempty"`
		ODataId                       *string               `json:"@odata.id,omitempty"`
		ODataType                     *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.IsCourseActivitySyncEnabled = decoded.IsCourseActivitySyncEnabled
	s.LearningContents = decoded.LearningContents
	s.LoginWebUrl = decoded.LoginWebUrl
	s.LongLogoWebUrlForDarkTheme = decoded.LongLogoWebUrlForDarkTheme
	s.LongLogoWebUrlForLightTheme = decoded.LongLogoWebUrlForLightTheme
	s.SquareLogoWebUrlForDarkTheme = decoded.SquareLogoWebUrlForDarkTheme
	s.SquareLogoWebUrlForLightTheme = decoded.SquareLogoWebUrlForLightTheme
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LearningProvider into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["learningCourseActivities"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling LearningCourseActivities into list []json.RawMessage: %+v", err)
		}

		output := make([]LearningCourseActivity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalLearningCourseActivityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'LearningCourseActivities' for 'LearningProvider': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.LearningCourseActivities = &output
	}

	return nil
}
