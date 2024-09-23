package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsentRequestFilterByCurrentUserOptions string

const (
	ConsentRequestFilterByCurrentUserOptions_Reviewer ConsentRequestFilterByCurrentUserOptions = "reviewer"
)

func PossibleValuesForConsentRequestFilterByCurrentUserOptions() []string {
	return []string{
		string(ConsentRequestFilterByCurrentUserOptions_Reviewer),
	}
}

func (s *ConsentRequestFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConsentRequestFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConsentRequestFilterByCurrentUserOptions(input string) (*ConsentRequestFilterByCurrentUserOptions, error) {
	vals := map[string]ConsentRequestFilterByCurrentUserOptions{
		"reviewer": ConsentRequestFilterByCurrentUserOptions_Reviewer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConsentRequestFilterByCurrentUserOptions(input)
	return &out, nil
}
