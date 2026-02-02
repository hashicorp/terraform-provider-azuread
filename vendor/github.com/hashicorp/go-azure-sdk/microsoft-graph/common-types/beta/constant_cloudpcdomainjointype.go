package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDomainJoinType string

const (
	CloudPCDomainJoinType_AzureADJoin       CloudPCDomainJoinType = "azureADJoin"
	CloudPCDomainJoinType_HybridAzureADJoin CloudPCDomainJoinType = "hybridAzureADJoin"
)

func PossibleValuesForCloudPCDomainJoinType() []string {
	return []string{
		string(CloudPCDomainJoinType_AzureADJoin),
		string(CloudPCDomainJoinType_HybridAzureADJoin),
	}
}

func (s *CloudPCDomainJoinType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDomainJoinType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDomainJoinType(input string) (*CloudPCDomainJoinType, error) {
	vals := map[string]CloudPCDomainJoinType{
		"azureadjoin":       CloudPCDomainJoinType_AzureADJoin,
		"hybridazureadjoin": CloudPCDomainJoinType_HybridAzureADJoin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDomainJoinType(input)
	return &out, nil
}
