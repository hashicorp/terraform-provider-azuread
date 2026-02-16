package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftManagedDesktopType string

const (
	MicrosoftManagedDesktopType_NotManaged      MicrosoftManagedDesktopType = "notManaged"
	MicrosoftManagedDesktopType_PremiumManaged  MicrosoftManagedDesktopType = "premiumManaged"
	MicrosoftManagedDesktopType_StandardManaged MicrosoftManagedDesktopType = "standardManaged"
	MicrosoftManagedDesktopType_StarterManaged  MicrosoftManagedDesktopType = "starterManaged"
)

func PossibleValuesForMicrosoftManagedDesktopType() []string {
	return []string{
		string(MicrosoftManagedDesktopType_NotManaged),
		string(MicrosoftManagedDesktopType_PremiumManaged),
		string(MicrosoftManagedDesktopType_StandardManaged),
		string(MicrosoftManagedDesktopType_StarterManaged),
	}
}

func (s *MicrosoftManagedDesktopType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftManagedDesktopType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftManagedDesktopType(input string) (*MicrosoftManagedDesktopType, error) {
	vals := map[string]MicrosoftManagedDesktopType{
		"notmanaged":      MicrosoftManagedDesktopType_NotManaged,
		"premiummanaged":  MicrosoftManagedDesktopType_PremiumManaged,
		"standardmanaged": MicrosoftManagedDesktopType_StandardManaged,
		"startermanaged":  MicrosoftManagedDesktopType_StarterManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftManagedDesktopType(input)
	return &out, nil
}
