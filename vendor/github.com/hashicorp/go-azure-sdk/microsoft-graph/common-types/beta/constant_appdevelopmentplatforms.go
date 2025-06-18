package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppDevelopmentPlatforms string

const (
	AppDevelopmentPlatforms_DeveloperPortal AppDevelopmentPlatforms = "developerPortal"
)

func PossibleValuesForAppDevelopmentPlatforms() []string {
	return []string{
		string(AppDevelopmentPlatforms_DeveloperPortal),
	}
}

func (s *AppDevelopmentPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppDevelopmentPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppDevelopmentPlatforms(input string) (*AppDevelopmentPlatforms, error) {
	vals := map[string]AppDevelopmentPlatforms{
		"developerportal": AppDevelopmentPlatforms_DeveloperPortal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppDevelopmentPlatforms(input)
	return &out, nil
}
