package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAppStartLayoutTileSize string

const (
	WindowsAppStartLayoutTileSize_Hidden WindowsAppStartLayoutTileSize = "hidden"
	WindowsAppStartLayoutTileSize_Large  WindowsAppStartLayoutTileSize = "large"
	WindowsAppStartLayoutTileSize_Medium WindowsAppStartLayoutTileSize = "medium"
	WindowsAppStartLayoutTileSize_Small  WindowsAppStartLayoutTileSize = "small"
	WindowsAppStartLayoutTileSize_Wide   WindowsAppStartLayoutTileSize = "wide"
)

func PossibleValuesForWindowsAppStartLayoutTileSize() []string {
	return []string{
		string(WindowsAppStartLayoutTileSize_Hidden),
		string(WindowsAppStartLayoutTileSize_Large),
		string(WindowsAppStartLayoutTileSize_Medium),
		string(WindowsAppStartLayoutTileSize_Small),
		string(WindowsAppStartLayoutTileSize_Wide),
	}
}

func (s *WindowsAppStartLayoutTileSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAppStartLayoutTileSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAppStartLayoutTileSize(input string) (*WindowsAppStartLayoutTileSize, error) {
	vals := map[string]WindowsAppStartLayoutTileSize{
		"hidden": WindowsAppStartLayoutTileSize_Hidden,
		"large":  WindowsAppStartLayoutTileSize_Large,
		"medium": WindowsAppStartLayoutTileSize_Medium,
		"small":  WindowsAppStartLayoutTileSize_Small,
		"wide":   WindowsAppStartLayoutTileSize_Wide,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAppStartLayoutTileSize(input)
	return &out, nil
}
