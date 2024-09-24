package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationContextDetail string

const (
	AuthenticationContextDetail_NotApplicable       AuthenticationContextDetail = "notApplicable"
	AuthenticationContextDetail_PreviouslySatisfied AuthenticationContextDetail = "previouslySatisfied"
	AuthenticationContextDetail_Required            AuthenticationContextDetail = "required"
)

func PossibleValuesForAuthenticationContextDetail() []string {
	return []string{
		string(AuthenticationContextDetail_NotApplicable),
		string(AuthenticationContextDetail_PreviouslySatisfied),
		string(AuthenticationContextDetail_Required),
	}
}

func (s *AuthenticationContextDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationContextDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationContextDetail(input string) (*AuthenticationContextDetail, error) {
	vals := map[string]AuthenticationContextDetail{
		"notapplicable":       AuthenticationContextDetail_NotApplicable,
		"previouslysatisfied": AuthenticationContextDetail_PreviouslySatisfied,
		"required":            AuthenticationContextDetail_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationContextDetail(input)
	return &out, nil
}
