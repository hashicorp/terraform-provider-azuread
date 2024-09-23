package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateScenarios string

const (
	TemplateScenarios_EmergingThreats  TemplateScenarios = "emergingThreats"
	TemplateScenarios_New              TemplateScenarios = "new"
	TemplateScenarios_ProtectAdmins    TemplateScenarios = "protectAdmins"
	TemplateScenarios_RemoteWork       TemplateScenarios = "remoteWork"
	TemplateScenarios_SecureFoundation TemplateScenarios = "secureFoundation"
	TemplateScenarios_ZeroTrust        TemplateScenarios = "zeroTrust"
)

func PossibleValuesForTemplateScenarios() []string {
	return []string{
		string(TemplateScenarios_EmergingThreats),
		string(TemplateScenarios_New),
		string(TemplateScenarios_ProtectAdmins),
		string(TemplateScenarios_RemoteWork),
		string(TemplateScenarios_SecureFoundation),
		string(TemplateScenarios_ZeroTrust),
	}
}

func (s *TemplateScenarios) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTemplateScenarios(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTemplateScenarios(input string) (*TemplateScenarios, error) {
	vals := map[string]TemplateScenarios{
		"emergingthreats":  TemplateScenarios_EmergingThreats,
		"new":              TemplateScenarios_New,
		"protectadmins":    TemplateScenarios_ProtectAdmins,
		"remotework":       TemplateScenarios_RemoteWork,
		"securefoundation": TemplateScenarios_SecureFoundation,
		"zerotrust":        TemplateScenarios_ZeroTrust,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TemplateScenarios(input)
	return &out, nil
}
