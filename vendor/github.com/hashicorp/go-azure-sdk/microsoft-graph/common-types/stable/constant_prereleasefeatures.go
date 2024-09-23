package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrereleaseFeatures string

const (
	PrereleaseFeatures_NotAllowed                  PrereleaseFeatures = "notAllowed"
	PrereleaseFeatures_SettingsAndExperimentations PrereleaseFeatures = "settingsAndExperimentations"
	PrereleaseFeatures_SettingsOnly                PrereleaseFeatures = "settingsOnly"
	PrereleaseFeatures_UserDefined                 PrereleaseFeatures = "userDefined"
)

func PossibleValuesForPrereleaseFeatures() []string {
	return []string{
		string(PrereleaseFeatures_NotAllowed),
		string(PrereleaseFeatures_SettingsAndExperimentations),
		string(PrereleaseFeatures_SettingsOnly),
		string(PrereleaseFeatures_UserDefined),
	}
}

func (s *PrereleaseFeatures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrereleaseFeatures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrereleaseFeatures(input string) (*PrereleaseFeatures, error) {
	vals := map[string]PrereleaseFeatures{
		"notallowed":                  PrereleaseFeatures_NotAllowed,
		"settingsandexperimentations": PrereleaseFeatures_SettingsAndExperimentations,
		"settingsonly":                PrereleaseFeatures_SettingsOnly,
		"userdefined":                 PrereleaseFeatures_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrereleaseFeatures(input)
	return &out, nil
}
