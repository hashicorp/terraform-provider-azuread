package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TimeZoneBase = CustomTimeZone{}

type CustomTimeZone struct {
	// The time offset of the time zone from Coordinated Universal Time (UTC). This value is in minutes. Time zones that are
	// ahead of UTC have a positive offset; time zones that are behind UTC have a negative offset.
	Bias nullable.Type[int64] `json:"bias,omitempty"`

	// Specifies when the time zone switches from standard time to daylight saving time.
	DaylightOffset *DaylightTimeZoneOffset `json:"daylightOffset,omitempty"`

	// Specifies when the time zone switches from daylight saving time to standard time.
	StandardOffset StandardTimeZoneOffset `json:"standardOffset"`

	// Fields inherited from TimeZoneBase

	// The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized
	// Time Zone' for a custom time zone.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CustomTimeZone) TimeZoneBase() BaseTimeZoneBaseImpl {
	return BaseTimeZoneBaseImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomTimeZone{}

func (s CustomTimeZone) MarshalJSON() ([]byte, error) {
	type wrapper CustomTimeZone
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomTimeZone: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomTimeZone: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customTimeZone"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomTimeZone: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CustomTimeZone{}

func (s *CustomTimeZone) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Bias           nullable.Type[int64]    `json:"bias,omitempty"`
		DaylightOffset *DaylightTimeZoneOffset `json:"daylightOffset,omitempty"`
		Name           nullable.Type[string]   `json:"name,omitempty"`
		ODataId        *string                 `json:"@odata.id,omitempty"`
		ODataType      *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Bias = decoded.Bias
	s.DaylightOffset = decoded.DaylightOffset
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CustomTimeZone into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["standardOffset"]; ok {
		impl, err := UnmarshalStandardTimeZoneOffsetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StandardOffset' for 'CustomTimeZone': %+v", err)
		}
		s.StandardOffset = impl
	}

	return nil
}
