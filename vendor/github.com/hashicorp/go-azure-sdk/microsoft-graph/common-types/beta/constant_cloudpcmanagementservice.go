package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCManagementService string

const (
	CloudPCManagementService_DevBox     CloudPCManagementService = "devBox"
	CloudPCManagementService_RpaBox     CloudPCManagementService = "rpaBox"
	CloudPCManagementService_Windows365 CloudPCManagementService = "windows365"
)

func PossibleValuesForCloudPCManagementService() []string {
	return []string{
		string(CloudPCManagementService_DevBox),
		string(CloudPCManagementService_RpaBox),
		string(CloudPCManagementService_Windows365),
	}
}

func (s *CloudPCManagementService) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCManagementService(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCManagementService(input string) (*CloudPCManagementService, error) {
	vals := map[string]CloudPCManagementService{
		"devbox":     CloudPCManagementService_DevBox,
		"rpabox":     CloudPCManagementService_RpaBox,
		"windows365": CloudPCManagementService_Windows365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCManagementService(input)
	return &out, nil
}
