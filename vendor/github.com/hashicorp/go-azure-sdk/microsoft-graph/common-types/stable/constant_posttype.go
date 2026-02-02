package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PostType string

const (
	PostType_Quick     PostType = "quick"
	PostType_Regular   PostType = "regular"
	PostType_Strategic PostType = "strategic"
)

func PossibleValuesForPostType() []string {
	return []string{
		string(PostType_Quick),
		string(PostType_Regular),
		string(PostType_Strategic),
	}
}

func (s *PostType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePostType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePostType(input string) (*PostType, error) {
	vals := map[string]PostType{
		"quick":     PostType_Quick,
		"regular":   PostType_Regular,
		"strategic": PostType_Strategic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PostType(input)
	return &out, nil
}
