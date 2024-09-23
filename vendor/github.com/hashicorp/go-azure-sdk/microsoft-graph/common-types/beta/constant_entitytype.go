package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityType string

const (
	EntityType_Acronym      EntityType = "acronym"
	EntityType_Bookmark     EntityType = "bookmark"
	EntityType_ChatMessage  EntityType = "chatMessage"
	EntityType_Drive        EntityType = "drive"
	EntityType_DriveItem    EntityType = "driveItem"
	EntityType_Event        EntityType = "event"
	EntityType_ExternalItem EntityType = "externalItem"
	EntityType_List         EntityType = "list"
	EntityType_ListItem     EntityType = "listItem"
	EntityType_Message      EntityType = "message"
	EntityType_Person       EntityType = "person"
	EntityType_Qna          EntityType = "qna"
	EntityType_Site         EntityType = "site"
)

func PossibleValuesForEntityType() []string {
	return []string{
		string(EntityType_Acronym),
		string(EntityType_Bookmark),
		string(EntityType_ChatMessage),
		string(EntityType_Drive),
		string(EntityType_DriveItem),
		string(EntityType_Event),
		string(EntityType_ExternalItem),
		string(EntityType_List),
		string(EntityType_ListItem),
		string(EntityType_Message),
		string(EntityType_Person),
		string(EntityType_Qna),
		string(EntityType_Site),
	}
}

func (s *EntityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntityType(input string) (*EntityType, error) {
	vals := map[string]EntityType{
		"acronym":      EntityType_Acronym,
		"bookmark":     EntityType_Bookmark,
		"chatmessage":  EntityType_ChatMessage,
		"drive":        EntityType_Drive,
		"driveitem":    EntityType_DriveItem,
		"event":        EntityType_Event,
		"externalitem": EntityType_ExternalItem,
		"list":         EntityType_List,
		"listitem":     EntityType_ListItem,
		"message":      EntityType_Message,
		"person":       EntityType_Person,
		"qna":          EntityType_Qna,
		"site":         EntityType_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityType(input)
	return &out, nil
}
