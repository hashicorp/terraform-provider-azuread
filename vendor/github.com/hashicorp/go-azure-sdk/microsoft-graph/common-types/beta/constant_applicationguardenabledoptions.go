package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGuardEnabledOptions string

const (
	ApplicationGuardEnabledOptions_EnabledForEdge          ApplicationGuardEnabledOptions = "enabledForEdge"
	ApplicationGuardEnabledOptions_EnabledForEdgeAndOffice ApplicationGuardEnabledOptions = "enabledForEdgeAndOffice"
	ApplicationGuardEnabledOptions_EnabledForOffice        ApplicationGuardEnabledOptions = "enabledForOffice"
	ApplicationGuardEnabledOptions_NotConfigured           ApplicationGuardEnabledOptions = "notConfigured"
)

func PossibleValuesForApplicationGuardEnabledOptions() []string {
	return []string{
		string(ApplicationGuardEnabledOptions_EnabledForEdge),
		string(ApplicationGuardEnabledOptions_EnabledForEdgeAndOffice),
		string(ApplicationGuardEnabledOptions_EnabledForOffice),
		string(ApplicationGuardEnabledOptions_NotConfigured),
	}
}

func (s *ApplicationGuardEnabledOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationGuardEnabledOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationGuardEnabledOptions(input string) (*ApplicationGuardEnabledOptions, error) {
	vals := map[string]ApplicationGuardEnabledOptions{
		"enabledforedge":          ApplicationGuardEnabledOptions_EnabledForEdge,
		"enabledforedgeandoffice": ApplicationGuardEnabledOptions_EnabledForEdgeAndOffice,
		"enabledforoffice":        ApplicationGuardEnabledOptions_EnabledForOffice,
		"notconfigured":           ApplicationGuardEnabledOptions_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationGuardEnabledOptions(input)
	return &out, nil
}
