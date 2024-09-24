package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceType string

const (
	SecurityKubernetesServiceType_ClusterIP    SecurityKubernetesServiceType = "clusterIP"
	SecurityKubernetesServiceType_ExternalName SecurityKubernetesServiceType = "externalName"
	SecurityKubernetesServiceType_LoadBalancer SecurityKubernetesServiceType = "loadBalancer"
	SecurityKubernetesServiceType_NodePort     SecurityKubernetesServiceType = "nodePort"
	SecurityKubernetesServiceType_Unknown      SecurityKubernetesServiceType = "unknown"
)

func PossibleValuesForSecurityKubernetesServiceType() []string {
	return []string{
		string(SecurityKubernetesServiceType_ClusterIP),
		string(SecurityKubernetesServiceType_ExternalName),
		string(SecurityKubernetesServiceType_LoadBalancer),
		string(SecurityKubernetesServiceType_NodePort),
		string(SecurityKubernetesServiceType_Unknown),
	}
}

func (s *SecurityKubernetesServiceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesServiceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesServiceType(input string) (*SecurityKubernetesServiceType, error) {
	vals := map[string]SecurityKubernetesServiceType{
		"clusterip":    SecurityKubernetesServiceType_ClusterIP,
		"externalname": SecurityKubernetesServiceType_ExternalName,
		"loadbalancer": SecurityKubernetesServiceType_LoadBalancer,
		"nodeport":     SecurityKubernetesServiceType_NodePort,
		"unknown":      SecurityKubernetesServiceType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceType(input)
	return &out, nil
}
