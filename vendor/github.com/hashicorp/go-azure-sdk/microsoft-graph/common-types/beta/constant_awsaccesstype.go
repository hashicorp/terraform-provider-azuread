package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsAccessType string

const (
	AwsAccessType_CrossAccount AwsAccessType = "crossAccount"
	AwsAccessType_Private      AwsAccessType = "private"
	AwsAccessType_Public       AwsAccessType = "public"
	AwsAccessType_Restricted   AwsAccessType = "restricted"
)

func PossibleValuesForAwsAccessType() []string {
	return []string{
		string(AwsAccessType_CrossAccount),
		string(AwsAccessType_Private),
		string(AwsAccessType_Public),
		string(AwsAccessType_Restricted),
	}
}

func (s *AwsAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsAccessType(input string) (*AwsAccessType, error) {
	vals := map[string]AwsAccessType{
		"crossaccount": AwsAccessType_CrossAccount,
		"private":      AwsAccessType_Private,
		"public":       AwsAccessType_Public,
		"restricted":   AwsAccessType_Restricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsAccessType(input)
	return &out, nil
}
