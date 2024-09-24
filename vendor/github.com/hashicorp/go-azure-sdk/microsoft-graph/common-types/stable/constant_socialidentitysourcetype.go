package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SocialIdentitySourceType string

const (
	SocialIdentitySourceType_Facebook SocialIdentitySourceType = "facebook"
)

func PossibleValuesForSocialIdentitySourceType() []string {
	return []string{
		string(SocialIdentitySourceType_Facebook),
	}
}

func (s *SocialIdentitySourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSocialIdentitySourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSocialIdentitySourceType(input string) (*SocialIdentitySourceType, error) {
	vals := map[string]SocialIdentitySourceType{
		"facebook": SocialIdentitySourceType_Facebook,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SocialIdentitySourceType(input)
	return &out, nil
}
