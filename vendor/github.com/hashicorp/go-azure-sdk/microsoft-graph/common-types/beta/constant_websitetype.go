package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebsiteType string

const (
	WebsiteType_Blog    WebsiteType = "blog"
	WebsiteType_Home    WebsiteType = "home"
	WebsiteType_Other   WebsiteType = "other"
	WebsiteType_Profile WebsiteType = "profile"
	WebsiteType_Work    WebsiteType = "work"
)

func PossibleValuesForWebsiteType() []string {
	return []string{
		string(WebsiteType_Blog),
		string(WebsiteType_Home),
		string(WebsiteType_Other),
		string(WebsiteType_Profile),
		string(WebsiteType_Work),
	}
}

func (s *WebsiteType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebsiteType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebsiteType(input string) (*WebsiteType, error) {
	vals := map[string]WebsiteType{
		"blog":    WebsiteType_Blog,
		"home":    WebsiteType_Home,
		"other":   WebsiteType_Other,
		"profile": WebsiteType_Profile,
		"work":    WebsiteType_Work,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebsiteType(input)
	return &out, nil
}
