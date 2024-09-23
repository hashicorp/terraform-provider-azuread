package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesPlatform string

const (
	SecurityKubernetesPlatform_Aks     SecurityKubernetesPlatform = "aks"
	SecurityKubernetesPlatform_Arc     SecurityKubernetesPlatform = "arc"
	SecurityKubernetesPlatform_Eks     SecurityKubernetesPlatform = "eks"
	SecurityKubernetesPlatform_Gke     SecurityKubernetesPlatform = "gke"
	SecurityKubernetesPlatform_Unknown SecurityKubernetesPlatform = "unknown"
)

func PossibleValuesForSecurityKubernetesPlatform() []string {
	return []string{
		string(SecurityKubernetesPlatform_Aks),
		string(SecurityKubernetesPlatform_Arc),
		string(SecurityKubernetesPlatform_Eks),
		string(SecurityKubernetesPlatform_Gke),
		string(SecurityKubernetesPlatform_Unknown),
	}
}

func (s *SecurityKubernetesPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesPlatform(input string) (*SecurityKubernetesPlatform, error) {
	vals := map[string]SecurityKubernetesPlatform{
		"aks":     SecurityKubernetesPlatform_Aks,
		"arc":     SecurityKubernetesPlatform_Arc,
		"eks":     SecurityKubernetesPlatform_Eks,
		"gke":     SecurityKubernetesPlatform_Gke,
		"unknown": SecurityKubernetesPlatform_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesPlatform(input)
	return &out, nil
}
