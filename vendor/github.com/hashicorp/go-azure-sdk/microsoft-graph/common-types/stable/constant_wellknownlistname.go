package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WellknownListName string

const (
	WellknownListName_DefaultList   WellknownListName = "defaultList"
	WellknownListName_FlaggedEmails WellknownListName = "flaggedEmails"
	WellknownListName_None          WellknownListName = "none"
)

func PossibleValuesForWellknownListName() []string {
	return []string{
		string(WellknownListName_DefaultList),
		string(WellknownListName_FlaggedEmails),
		string(WellknownListName_None),
	}
}

func (s *WellknownListName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWellknownListName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWellknownListName(input string) (*WellknownListName, error) {
	vals := map[string]WellknownListName{
		"defaultlist":   WellknownListName_DefaultList,
		"flaggedemails": WellknownListName_FlaggedEmails,
		"none":          WellknownListName_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WellknownListName(input)
	return &out, nil
}
