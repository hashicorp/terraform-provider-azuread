package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserNewMessageRestriction string

const (
	UserNewMessageRestriction_Everyone             UserNewMessageRestriction = "everyone"
	UserNewMessageRestriction_EveryoneExceptGuests UserNewMessageRestriction = "everyoneExceptGuests"
	UserNewMessageRestriction_Moderators           UserNewMessageRestriction = "moderators"
)

func PossibleValuesForUserNewMessageRestriction() []string {
	return []string{
		string(UserNewMessageRestriction_Everyone),
		string(UserNewMessageRestriction_EveryoneExceptGuests),
		string(UserNewMessageRestriction_Moderators),
	}
}

func (s *UserNewMessageRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserNewMessageRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserNewMessageRestriction(input string) (*UserNewMessageRestriction, error) {
	vals := map[string]UserNewMessageRestriction{
		"everyone":             UserNewMessageRestriction_Everyone,
		"everyoneexceptguests": UserNewMessageRestriction_EveryoneExceptGuests,
		"moderators":           UserNewMessageRestriction_Moderators,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserNewMessageRestriction(input)
	return &out, nil
}
