package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsImportanceScore string

const (
	ExternalConnectorsImportanceScore_High     ExternalConnectorsImportanceScore = "high"
	ExternalConnectorsImportanceScore_Low      ExternalConnectorsImportanceScore = "low"
	ExternalConnectorsImportanceScore_Medium   ExternalConnectorsImportanceScore = "medium"
	ExternalConnectorsImportanceScore_VeryHigh ExternalConnectorsImportanceScore = "veryHigh"
)

func PossibleValuesForExternalConnectorsImportanceScore() []string {
	return []string{
		string(ExternalConnectorsImportanceScore_High),
		string(ExternalConnectorsImportanceScore_Low),
		string(ExternalConnectorsImportanceScore_Medium),
		string(ExternalConnectorsImportanceScore_VeryHigh),
	}
}

func (s *ExternalConnectorsImportanceScore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsImportanceScore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsImportanceScore(input string) (*ExternalConnectorsImportanceScore, error) {
	vals := map[string]ExternalConnectorsImportanceScore{
		"high":     ExternalConnectorsImportanceScore_High,
		"low":      ExternalConnectorsImportanceScore_Low,
		"medium":   ExternalConnectorsImportanceScore_Medium,
		"veryhigh": ExternalConnectorsImportanceScore_VeryHigh,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsImportanceScore(input)
	return &out, nil
}
