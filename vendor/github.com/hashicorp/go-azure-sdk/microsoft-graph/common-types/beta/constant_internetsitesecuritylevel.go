package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InternetSiteSecurityLevel string

const (
	InternetSiteSecurityLevel_High        InternetSiteSecurityLevel = "high"
	InternetSiteSecurityLevel_Medium      InternetSiteSecurityLevel = "medium"
	InternetSiteSecurityLevel_MediumHigh  InternetSiteSecurityLevel = "mediumHigh"
	InternetSiteSecurityLevel_UserDefined InternetSiteSecurityLevel = "userDefined"
)

func PossibleValuesForInternetSiteSecurityLevel() []string {
	return []string{
		string(InternetSiteSecurityLevel_High),
		string(InternetSiteSecurityLevel_Medium),
		string(InternetSiteSecurityLevel_MediumHigh),
		string(InternetSiteSecurityLevel_UserDefined),
	}
}

func (s *InternetSiteSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInternetSiteSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInternetSiteSecurityLevel(input string) (*InternetSiteSecurityLevel, error) {
	vals := map[string]InternetSiteSecurityLevel{
		"high":        InternetSiteSecurityLevel_High,
		"medium":      InternetSiteSecurityLevel_Medium,
		"mediumhigh":  InternetSiteSecurityLevel_MediumHigh,
		"userdefined": InternetSiteSecurityLevel_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InternetSiteSecurityLevel(input)
	return &out, nil
}
