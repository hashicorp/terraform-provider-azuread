package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelDeploymentMode string

const (
	MicrosoftTunnelDeploymentMode_PodRootful         MicrosoftTunnelDeploymentMode = "podRootful"
	MicrosoftTunnelDeploymentMode_PodRootless        MicrosoftTunnelDeploymentMode = "podRootless"
	MicrosoftTunnelDeploymentMode_StandaloneRootful  MicrosoftTunnelDeploymentMode = "standaloneRootful"
	MicrosoftTunnelDeploymentMode_StandaloneRootless MicrosoftTunnelDeploymentMode = "standaloneRootless"
)

func PossibleValuesForMicrosoftTunnelDeploymentMode() []string {
	return []string{
		string(MicrosoftTunnelDeploymentMode_PodRootful),
		string(MicrosoftTunnelDeploymentMode_PodRootless),
		string(MicrosoftTunnelDeploymentMode_StandaloneRootful),
		string(MicrosoftTunnelDeploymentMode_StandaloneRootless),
	}
}

func (s *MicrosoftTunnelDeploymentMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftTunnelDeploymentMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftTunnelDeploymentMode(input string) (*MicrosoftTunnelDeploymentMode, error) {
	vals := map[string]MicrosoftTunnelDeploymentMode{
		"podrootful":         MicrosoftTunnelDeploymentMode_PodRootful,
		"podrootless":        MicrosoftTunnelDeploymentMode_PodRootless,
		"standalonerootful":  MicrosoftTunnelDeploymentMode_StandaloneRootful,
		"standalonerootless": MicrosoftTunnelDeploymentMode_StandaloneRootless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTunnelDeploymentMode(input)
	return &out, nil
}
