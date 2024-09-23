package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsRoleTrustEntityType string

const (
	AwsRoleTrustEntityType_CrossAccount AwsRoleTrustEntityType = "crossAccount"
	AwsRoleTrustEntityType_None         AwsRoleTrustEntityType = "none"
	AwsRoleTrustEntityType_Service      AwsRoleTrustEntityType = "service"
	AwsRoleTrustEntityType_Sso          AwsRoleTrustEntityType = "sso"
	AwsRoleTrustEntityType_WebIdentity  AwsRoleTrustEntityType = "webIdentity"
)

func PossibleValuesForAwsRoleTrustEntityType() []string {
	return []string{
		string(AwsRoleTrustEntityType_CrossAccount),
		string(AwsRoleTrustEntityType_None),
		string(AwsRoleTrustEntityType_Service),
		string(AwsRoleTrustEntityType_Sso),
		string(AwsRoleTrustEntityType_WebIdentity),
	}
}

func (s *AwsRoleTrustEntityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsRoleTrustEntityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsRoleTrustEntityType(input string) (*AwsRoleTrustEntityType, error) {
	vals := map[string]AwsRoleTrustEntityType{
		"crossaccount": AwsRoleTrustEntityType_CrossAccount,
		"none":         AwsRoleTrustEntityType_None,
		"service":      AwsRoleTrustEntityType_Service,
		"sso":          AwsRoleTrustEntityType_Sso,
		"webidentity":  AwsRoleTrustEntityType_WebIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsRoleTrustEntityType(input)
	return &out, nil
}
