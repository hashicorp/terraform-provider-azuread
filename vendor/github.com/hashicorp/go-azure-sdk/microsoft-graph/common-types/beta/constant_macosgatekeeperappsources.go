package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSGatekeeperAppSources string

const (
	MacOSGatekeeperAppSources_Anywhere                           MacOSGatekeeperAppSources = "anywhere"
	MacOSGatekeeperAppSources_MacAppStore                        MacOSGatekeeperAppSources = "macAppStore"
	MacOSGatekeeperAppSources_MacAppStoreAndIdentifiedDevelopers MacOSGatekeeperAppSources = "macAppStoreAndIdentifiedDevelopers"
	MacOSGatekeeperAppSources_NotConfigured                      MacOSGatekeeperAppSources = "notConfigured"
)

func PossibleValuesForMacOSGatekeeperAppSources() []string {
	return []string{
		string(MacOSGatekeeperAppSources_Anywhere),
		string(MacOSGatekeeperAppSources_MacAppStore),
		string(MacOSGatekeeperAppSources_MacAppStoreAndIdentifiedDevelopers),
		string(MacOSGatekeeperAppSources_NotConfigured),
	}
}

func (s *MacOSGatekeeperAppSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSGatekeeperAppSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSGatekeeperAppSources(input string) (*MacOSGatekeeperAppSources, error) {
	vals := map[string]MacOSGatekeeperAppSources{
		"anywhere":                           MacOSGatekeeperAppSources_Anywhere,
		"macappstore":                        MacOSGatekeeperAppSources_MacAppStore,
		"macappstoreandidentifieddevelopers": MacOSGatekeeperAppSources_MacAppStoreAndIdentifiedDevelopers,
		"notconfigured":                      MacOSGatekeeperAppSources_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSGatekeeperAppSources(input)
	return &out, nil
}
