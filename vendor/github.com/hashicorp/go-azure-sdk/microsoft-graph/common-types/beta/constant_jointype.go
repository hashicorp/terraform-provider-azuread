package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinType string

const (
	JoinType_AzureADJoined       JoinType = "azureADJoined"
	JoinType_AzureADRegistered   JoinType = "azureADRegistered"
	JoinType_HybridAzureADJoined JoinType = "hybridAzureADJoined"
	JoinType_Unknown             JoinType = "unknown"
)

func PossibleValuesForJoinType() []string {
	return []string{
		string(JoinType_AzureADJoined),
		string(JoinType_AzureADRegistered),
		string(JoinType_HybridAzureADJoined),
		string(JoinType_Unknown),
	}
}

func (s *JoinType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJoinType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJoinType(input string) (*JoinType, error) {
	vals := map[string]JoinType{
		"azureadjoined":       JoinType_AzureADJoined,
		"azureadregistered":   JoinType_AzureADRegistered,
		"hybridazureadjoined": JoinType_HybridAzureADJoined,
		"unknown":             JoinType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JoinType(input)
	return &out, nil
}
