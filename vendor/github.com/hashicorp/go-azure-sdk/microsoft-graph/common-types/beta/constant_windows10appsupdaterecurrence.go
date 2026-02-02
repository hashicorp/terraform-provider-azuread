package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10AppsUpdateRecurrence string

const (
	Windows10AppsUpdateRecurrence_Daily   Windows10AppsUpdateRecurrence = "daily"
	Windows10AppsUpdateRecurrence_Monthly Windows10AppsUpdateRecurrence = "monthly"
	Windows10AppsUpdateRecurrence_None    Windows10AppsUpdateRecurrence = "none"
	Windows10AppsUpdateRecurrence_Weekly  Windows10AppsUpdateRecurrence = "weekly"
)

func PossibleValuesForWindows10AppsUpdateRecurrence() []string {
	return []string{
		string(Windows10AppsUpdateRecurrence_Daily),
		string(Windows10AppsUpdateRecurrence_Monthly),
		string(Windows10AppsUpdateRecurrence_None),
		string(Windows10AppsUpdateRecurrence_Weekly),
	}
}

func (s *Windows10AppsUpdateRecurrence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10AppsUpdateRecurrence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10AppsUpdateRecurrence(input string) (*Windows10AppsUpdateRecurrence, error) {
	vals := map[string]Windows10AppsUpdateRecurrence{
		"daily":   Windows10AppsUpdateRecurrence_Daily,
		"monthly": Windows10AppsUpdateRecurrence_Monthly,
		"none":    Windows10AppsUpdateRecurrence_None,
		"weekly":  Windows10AppsUpdateRecurrence_Weekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10AppsUpdateRecurrence(input)
	return &out, nil
}
