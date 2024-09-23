package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditActivityOperationType string

const (
	CloudPCAuditActivityOperationType_Create CloudPCAuditActivityOperationType = "create"
	CloudPCAuditActivityOperationType_Delete CloudPCAuditActivityOperationType = "delete"
	CloudPCAuditActivityOperationType_Patch  CloudPCAuditActivityOperationType = "patch"
)

func PossibleValuesForCloudPCAuditActivityOperationType() []string {
	return []string{
		string(CloudPCAuditActivityOperationType_Create),
		string(CloudPCAuditActivityOperationType_Delete),
		string(CloudPCAuditActivityOperationType_Patch),
	}
}

func (s *CloudPCAuditActivityOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAuditActivityOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAuditActivityOperationType(input string) (*CloudPCAuditActivityOperationType, error) {
	vals := map[string]CloudPCAuditActivityOperationType{
		"create": CloudPCAuditActivityOperationType_Create,
		"delete": CloudPCAuditActivityOperationType_Delete,
		"patch":  CloudPCAuditActivityOperationType_Patch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditActivityOperationType(input)
	return &out, nil
}
