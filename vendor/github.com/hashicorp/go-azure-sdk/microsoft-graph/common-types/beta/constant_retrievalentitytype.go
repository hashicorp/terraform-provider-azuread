package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetrievalEntityType string

const (
	RetrievalEntityType_Drive        RetrievalEntityType = "drive"
	RetrievalEntityType_DriveItem    RetrievalEntityType = "driveItem"
	RetrievalEntityType_ExternalItem RetrievalEntityType = "externalItem"
	RetrievalEntityType_List         RetrievalEntityType = "list"
	RetrievalEntityType_ListItem     RetrievalEntityType = "listItem"
	RetrievalEntityType_Site         RetrievalEntityType = "site"
)

func PossibleValuesForRetrievalEntityType() []string {
	return []string{
		string(RetrievalEntityType_Drive),
		string(RetrievalEntityType_DriveItem),
		string(RetrievalEntityType_ExternalItem),
		string(RetrievalEntityType_List),
		string(RetrievalEntityType_ListItem),
		string(RetrievalEntityType_Site),
	}
}

func (s *RetrievalEntityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRetrievalEntityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRetrievalEntityType(input string) (*RetrievalEntityType, error) {
	vals := map[string]RetrievalEntityType{
		"drive":        RetrievalEntityType_Drive,
		"driveitem":    RetrievalEntityType_DriveItem,
		"externalitem": RetrievalEntityType_ExternalItem,
		"list":         RetrievalEntityType_List,
		"listitem":     RetrievalEntityType_ListItem,
		"site":         RetrievalEntityType_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetrievalEntityType(input)
	return &out, nil
}
