package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationKeyOrigin string

const (
	ApplicationKeyOrigin_Application      ApplicationKeyOrigin = "application"
	ApplicationKeyOrigin_ServicePrincipal ApplicationKeyOrigin = "servicePrincipal"
)

func PossibleValuesForApplicationKeyOrigin() []string {
	return []string{
		string(ApplicationKeyOrigin_Application),
		string(ApplicationKeyOrigin_ServicePrincipal),
	}
}

func (s *ApplicationKeyOrigin) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationKeyOrigin(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationKeyOrigin(input string) (*ApplicationKeyOrigin, error) {
	vals := map[string]ApplicationKeyOrigin{
		"application":      ApplicationKeyOrigin_Application,
		"serviceprincipal": ApplicationKeyOrigin_ServicePrincipal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationKeyOrigin(input)
	return &out, nil
}
