package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingProviderType string

const (
	OnlineMeetingProviderType_SkypeForBusiness OnlineMeetingProviderType = "skypeForBusiness"
	OnlineMeetingProviderType_SkypeForConsumer OnlineMeetingProviderType = "skypeForConsumer"
	OnlineMeetingProviderType_TeamsForBusiness OnlineMeetingProviderType = "teamsForBusiness"
	OnlineMeetingProviderType_Unknown          OnlineMeetingProviderType = "unknown"
)

func PossibleValuesForOnlineMeetingProviderType() []string {
	return []string{
		string(OnlineMeetingProviderType_SkypeForBusiness),
		string(OnlineMeetingProviderType_SkypeForConsumer),
		string(OnlineMeetingProviderType_TeamsForBusiness),
		string(OnlineMeetingProviderType_Unknown),
	}
}

func (s *OnlineMeetingProviderType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingProviderType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingProviderType(input string) (*OnlineMeetingProviderType, error) {
	vals := map[string]OnlineMeetingProviderType{
		"skypeforbusiness": OnlineMeetingProviderType_SkypeForBusiness,
		"skypeforconsumer": OnlineMeetingProviderType_SkypeForConsumer,
		"teamsforbusiness": OnlineMeetingProviderType_TeamsForBusiness,
		"unknown":          OnlineMeetingProviderType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingProviderType(input)
	return &out, nil
}
