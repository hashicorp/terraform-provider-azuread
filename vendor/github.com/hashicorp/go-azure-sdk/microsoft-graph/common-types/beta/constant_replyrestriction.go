package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplyRestriction string

const (
	ReplyRestriction_AuthorAndModerators ReplyRestriction = "authorAndModerators"
	ReplyRestriction_Everyone            ReplyRestriction = "everyone"
)

func PossibleValuesForReplyRestriction() []string {
	return []string{
		string(ReplyRestriction_AuthorAndModerators),
		string(ReplyRestriction_Everyone),
	}
}

func (s *ReplyRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplyRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplyRestriction(input string) (*ReplyRestriction, error) {
	vals := map[string]ReplyRestriction{
		"authorandmoderators": ReplyRestriction_AuthorAndModerators,
		"everyone":            ReplyRestriction_Everyone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplyRestriction(input)
	return &out, nil
}
