package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSignInRecommendationScope string

const (
	UserSignInRecommendationScope_Application UserSignInRecommendationScope = "application"
	UserSignInRecommendationScope_Tenant      UserSignInRecommendationScope = "tenant"
)

func PossibleValuesForUserSignInRecommendationScope() []string {
	return []string{
		string(UserSignInRecommendationScope_Application),
		string(UserSignInRecommendationScope_Tenant),
	}
}

func (s *UserSignInRecommendationScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserSignInRecommendationScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserSignInRecommendationScope(input string) (*UserSignInRecommendationScope, error) {
	vals := map[string]UserSignInRecommendationScope{
		"application": UserSignInRecommendationScope_Application,
		"tenant":      UserSignInRecommendationScope_Tenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserSignInRecommendationScope(input)
	return &out, nil
}
