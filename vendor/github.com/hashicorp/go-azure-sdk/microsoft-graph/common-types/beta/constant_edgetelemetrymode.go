package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeTelemetryMode string

const (
	EdgeTelemetryMode_Internet            EdgeTelemetryMode = "internet"
	EdgeTelemetryMode_Intranet            EdgeTelemetryMode = "intranet"
	EdgeTelemetryMode_IntranetAndInternet EdgeTelemetryMode = "intranetAndInternet"
	EdgeTelemetryMode_NotConfigured       EdgeTelemetryMode = "notConfigured"
)

func PossibleValuesForEdgeTelemetryMode() []string {
	return []string{
		string(EdgeTelemetryMode_Internet),
		string(EdgeTelemetryMode_Intranet),
		string(EdgeTelemetryMode_IntranetAndInternet),
		string(EdgeTelemetryMode_NotConfigured),
	}
}

func (s *EdgeTelemetryMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdgeTelemetryMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdgeTelemetryMode(input string) (*EdgeTelemetryMode, error) {
	vals := map[string]EdgeTelemetryMode{
		"internet":            EdgeTelemetryMode_Internet,
		"intranet":            EdgeTelemetryMode_Intranet,
		"intranetandinternet": EdgeTelemetryMode_IntranetAndInternet,
		"notconfigured":       EdgeTelemetryMode_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdgeTelemetryMode(input)
	return &out, nil
}
