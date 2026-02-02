package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResponseType string

const (
	ResponseType_Accepted            ResponseType = "accepted"
	ResponseType_Declined            ResponseType = "declined"
	ResponseType_None                ResponseType = "none"
	ResponseType_NotResponded        ResponseType = "notResponded"
	ResponseType_Organizer           ResponseType = "organizer"
	ResponseType_TentativelyAccepted ResponseType = "tentativelyAccepted"
)

func PossibleValuesForResponseType() []string {
	return []string{
		string(ResponseType_Accepted),
		string(ResponseType_Declined),
		string(ResponseType_None),
		string(ResponseType_NotResponded),
		string(ResponseType_Organizer),
		string(ResponseType_TentativelyAccepted),
	}
}

func (s *ResponseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResponseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResponseType(input string) (*ResponseType, error) {
	vals := map[string]ResponseType{
		"accepted":            ResponseType_Accepted,
		"declined":            ResponseType_Declined,
		"none":                ResponseType_None,
		"notresponded":        ResponseType_NotResponded,
		"organizer":           ResponseType_Organizer,
		"tentativelyaccepted": ResponseType_TentativelyAccepted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResponseType(input)
	return &out, nil
}
