package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionType string

const (
	CloudPCOnPremisesConnectionType_AzureADJoin       CloudPCOnPremisesConnectionType = "azureADJoin"
	CloudPCOnPremisesConnectionType_HybridAzureADJoin CloudPCOnPremisesConnectionType = "hybridAzureADJoin"
)

func PossibleValuesForCloudPCOnPremisesConnectionType() []string {
	return []string{
		string(CloudPCOnPremisesConnectionType_AzureADJoin),
		string(CloudPCOnPremisesConnectionType_HybridAzureADJoin),
	}
}

func (s *CloudPCOnPremisesConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOnPremisesConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOnPremisesConnectionType(input string) (*CloudPCOnPremisesConnectionType, error) {
	vals := map[string]CloudPCOnPremisesConnectionType{
		"azureadjoin":       CloudPCOnPremisesConnectionType_AzureADJoin,
		"hybridazureadjoin": CloudPCOnPremisesConnectionType_HybridAzureADJoin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOnPremisesConnectionType(input)
	return &out, nil
}
