package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSContentCachingType string

const (
	MacOSContentCachingType_NotConfigured     MacOSContentCachingType = "notConfigured"
	MacOSContentCachingType_SharedContentOnly MacOSContentCachingType = "sharedContentOnly"
	MacOSContentCachingType_UserContentOnly   MacOSContentCachingType = "userContentOnly"
)

func PossibleValuesForMacOSContentCachingType() []string {
	return []string{
		string(MacOSContentCachingType_NotConfigured),
		string(MacOSContentCachingType_SharedContentOnly),
		string(MacOSContentCachingType_UserContentOnly),
	}
}

func (s *MacOSContentCachingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSContentCachingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSContentCachingType(input string) (*MacOSContentCachingType, error) {
	vals := map[string]MacOSContentCachingType{
		"notconfigured":     MacOSContentCachingType_NotConfigured,
		"sharedcontentonly": MacOSContentCachingType_SharedContentOnly,
		"usercontentonly":   MacOSContentCachingType_UserContentOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSContentCachingType(input)
	return &out, nil
}
