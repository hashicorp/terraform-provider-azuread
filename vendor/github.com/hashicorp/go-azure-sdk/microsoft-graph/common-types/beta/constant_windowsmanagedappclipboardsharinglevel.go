package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppClipboardSharingLevel string

const (
	WindowsManagedAppClipboardSharingLevel_AnyDestinationAnySource WindowsManagedAppClipboardSharingLevel = "anyDestinationAnySource"
	WindowsManagedAppClipboardSharingLevel_None                    WindowsManagedAppClipboardSharingLevel = "none"
)

func PossibleValuesForWindowsManagedAppClipboardSharingLevel() []string {
	return []string{
		string(WindowsManagedAppClipboardSharingLevel_AnyDestinationAnySource),
		string(WindowsManagedAppClipboardSharingLevel_None),
	}
}

func (s *WindowsManagedAppClipboardSharingLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedAppClipboardSharingLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedAppClipboardSharingLevel(input string) (*WindowsManagedAppClipboardSharingLevel, error) {
	vals := map[string]WindowsManagedAppClipboardSharingLevel{
		"anydestinationanysource": WindowsManagedAppClipboardSharingLevel_AnyDestinationAnySource,
		"none":                    WindowsManagedAppClipboardSharingLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppClipboardSharingLevel(input)
	return &out, nil
}
