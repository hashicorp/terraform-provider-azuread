package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroundingEntityType string

const (
	GroundingEntityType_Drive     GroundingEntityType = "drive"
	GroundingEntityType_DriveItem GroundingEntityType = "driveItem"
	GroundingEntityType_List      GroundingEntityType = "list"
	GroundingEntityType_ListItem  GroundingEntityType = "listItem"
	GroundingEntityType_Site      GroundingEntityType = "site"
)

func PossibleValuesForGroundingEntityType() []string {
	return []string{
		string(GroundingEntityType_Drive),
		string(GroundingEntityType_DriveItem),
		string(GroundingEntityType_List),
		string(GroundingEntityType_ListItem),
		string(GroundingEntityType_Site),
	}
}

func (s *GroundingEntityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroundingEntityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroundingEntityType(input string) (*GroundingEntityType, error) {
	vals := map[string]GroundingEntityType{
		"drive":     GroundingEntityType_Drive,
		"driveitem": GroundingEntityType_DriveItem,
		"list":      GroundingEntityType_List,
		"listitem":  GroundingEntityType_ListItem,
		"site":      GroundingEntityType_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroundingEntityType(input)
	return &out, nil
}
