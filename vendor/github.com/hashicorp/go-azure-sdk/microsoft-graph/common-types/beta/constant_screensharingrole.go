package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScreenSharingRole string

const (
	ScreenSharingRole_Sharer ScreenSharingRole = "sharer"
	ScreenSharingRole_Viewer ScreenSharingRole = "viewer"
)

func PossibleValuesForScreenSharingRole() []string {
	return []string{
		string(ScreenSharingRole_Sharer),
		string(ScreenSharingRole_Viewer),
	}
}

func (s *ScreenSharingRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScreenSharingRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScreenSharingRole(input string) (*ScreenSharingRole, error) {
	vals := map[string]ScreenSharingRole{
		"sharer": ScreenSharingRole_Sharer,
		"viewer": ScreenSharingRole_Viewer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScreenSharingRole(input)
	return &out, nil
}
