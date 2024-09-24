package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAccessType string

const (
	AzureAccessType_Private AzureAccessType = "private"
	AzureAccessType_Public  AzureAccessType = "public"
)

func PossibleValuesForAzureAccessType() []string {
	return []string{
		string(AzureAccessType_Private),
		string(AzureAccessType_Public),
	}
}

func (s *AzureAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureAccessType(input string) (*AzureAccessType, error) {
	vals := map[string]AzureAccessType{
		"private": AzureAccessType_Private,
		"public":  AzureAccessType_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureAccessType(input)
	return &out, nil
}
