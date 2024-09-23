package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppDataTransferLevel string

const (
	WindowsManagedAppDataTransferLevel_AllApps WindowsManagedAppDataTransferLevel = "allApps"
	WindowsManagedAppDataTransferLevel_None    WindowsManagedAppDataTransferLevel = "none"
)

func PossibleValuesForWindowsManagedAppDataTransferLevel() []string {
	return []string{
		string(WindowsManagedAppDataTransferLevel_AllApps),
		string(WindowsManagedAppDataTransferLevel_None),
	}
}

func (s *WindowsManagedAppDataTransferLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedAppDataTransferLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedAppDataTransferLevel(input string) (*WindowsManagedAppDataTransferLevel, error) {
	vals := map[string]WindowsManagedAppDataTransferLevel{
		"allapps": WindowsManagedAppDataTransferLevel_AllApps,
		"none":    WindowsManagedAppDataTransferLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppDataTransferLevel(input)
	return &out, nil
}
