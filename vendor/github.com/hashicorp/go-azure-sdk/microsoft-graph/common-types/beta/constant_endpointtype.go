package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointType string

const (
	EndpointType_Default                   EndpointType = "default"
	EndpointType_SkypeForBusiness          EndpointType = "skypeForBusiness"
	EndpointType_SkypeForBusinessVoipPhone EndpointType = "skypeForBusinessVoipPhone"
	EndpointType_Voicemail                 EndpointType = "voicemail"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointType_Default),
		string(EndpointType_SkypeForBusiness),
		string(EndpointType_SkypeForBusinessVoipPhone),
		string(EndpointType_Voicemail),
	}
}

func (s *EndpointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointType(input string) (*EndpointType, error) {
	vals := map[string]EndpointType{
		"default":                   EndpointType_Default,
		"skypeforbusiness":          EndpointType_SkypeForBusiness,
		"skypeforbusinessvoipphone": EndpointType_SkypeForBusinessVoipPhone,
		"voicemail":                 EndpointType_Voicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointType(input)
	return &out, nil
}
