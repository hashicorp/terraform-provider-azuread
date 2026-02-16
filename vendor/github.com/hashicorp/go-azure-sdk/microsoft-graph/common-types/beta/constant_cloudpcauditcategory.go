package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditCategory string

const (
	CloudPCAuditCategory_CloudPC CloudPCAuditCategory = "cloudPC"
)

func PossibleValuesForCloudPCAuditCategory() []string {
	return []string{
		string(CloudPCAuditCategory_CloudPC),
	}
}

func (s *CloudPCAuditCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAuditCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAuditCategory(input string) (*CloudPCAuditCategory, error) {
	vals := map[string]CloudPCAuditCategory{
		"cloudpc": CloudPCAuditCategory_CloudPC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditCategory(input)
	return &out, nil
}
