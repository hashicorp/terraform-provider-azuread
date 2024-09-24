package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftLauncherDockPresence string

const (
	MicrosoftLauncherDockPresence_Disabled      MicrosoftLauncherDockPresence = "disabled"
	MicrosoftLauncherDockPresence_Hide          MicrosoftLauncherDockPresence = "hide"
	MicrosoftLauncherDockPresence_NotConfigured MicrosoftLauncherDockPresence = "notConfigured"
	MicrosoftLauncherDockPresence_Show          MicrosoftLauncherDockPresence = "show"
)

func PossibleValuesForMicrosoftLauncherDockPresence() []string {
	return []string{
		string(MicrosoftLauncherDockPresence_Disabled),
		string(MicrosoftLauncherDockPresence_Hide),
		string(MicrosoftLauncherDockPresence_NotConfigured),
		string(MicrosoftLauncherDockPresence_Show),
	}
}

func (s *MicrosoftLauncherDockPresence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftLauncherDockPresence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftLauncherDockPresence(input string) (*MicrosoftLauncherDockPresence, error) {
	vals := map[string]MicrosoftLauncherDockPresence{
		"disabled":      MicrosoftLauncherDockPresence_Disabled,
		"hide":          MicrosoftLauncherDockPresence_Hide,
		"notconfigured": MicrosoftLauncherDockPresence_NotConfigured,
		"show":          MicrosoftLauncherDockPresence_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftLauncherDockPresence(input)
	return &out, nil
}
