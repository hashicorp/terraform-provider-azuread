package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsEdgeKioskType string

const (
	WindowsEdgeKioskType_FullScreen     WindowsEdgeKioskType = "fullScreen"
	WindowsEdgeKioskType_PublicBrowsing WindowsEdgeKioskType = "publicBrowsing"
)

func PossibleValuesForWindowsEdgeKioskType() []string {
	return []string{
		string(WindowsEdgeKioskType_FullScreen),
		string(WindowsEdgeKioskType_PublicBrowsing),
	}
}

func (s *WindowsEdgeKioskType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsEdgeKioskType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsEdgeKioskType(input string) (*WindowsEdgeKioskType, error) {
	vals := map[string]WindowsEdgeKioskType{
		"fullscreen":     WindowsEdgeKioskType_FullScreen,
		"publicbrowsing": WindowsEdgeKioskType_PublicBrowsing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsEdgeKioskType(input)
	return &out, nil
}
