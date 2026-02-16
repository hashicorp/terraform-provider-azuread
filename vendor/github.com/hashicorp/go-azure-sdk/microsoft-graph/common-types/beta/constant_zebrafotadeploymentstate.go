package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentState string

const (
	ZebraFotaDeploymentState_Canceled        ZebraFotaDeploymentState = "canceled"
	ZebraFotaDeploymentState_Completed       ZebraFotaDeploymentState = "completed"
	ZebraFotaDeploymentState_CreateFailed    ZebraFotaDeploymentState = "createFailed"
	ZebraFotaDeploymentState_Created         ZebraFotaDeploymentState = "created"
	ZebraFotaDeploymentState_InProgress      ZebraFotaDeploymentState = "inProgress"
	ZebraFotaDeploymentState_PendingCancel   ZebraFotaDeploymentState = "pendingCancel"
	ZebraFotaDeploymentState_PendingCreation ZebraFotaDeploymentState = "pendingCreation"
)

func PossibleValuesForZebraFotaDeploymentState() []string {
	return []string{
		string(ZebraFotaDeploymentState_Canceled),
		string(ZebraFotaDeploymentState_Completed),
		string(ZebraFotaDeploymentState_CreateFailed),
		string(ZebraFotaDeploymentState_Created),
		string(ZebraFotaDeploymentState_InProgress),
		string(ZebraFotaDeploymentState_PendingCancel),
		string(ZebraFotaDeploymentState_PendingCreation),
	}
}

func (s *ZebraFotaDeploymentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaDeploymentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaDeploymentState(input string) (*ZebraFotaDeploymentState, error) {
	vals := map[string]ZebraFotaDeploymentState{
		"canceled":        ZebraFotaDeploymentState_Canceled,
		"completed":       ZebraFotaDeploymentState_Completed,
		"createfailed":    ZebraFotaDeploymentState_CreateFailed,
		"created":         ZebraFotaDeploymentState_Created,
		"inprogress":      ZebraFotaDeploymentState_InProgress,
		"pendingcancel":   ZebraFotaDeploymentState_PendingCancel,
		"pendingcreation": ZebraFotaDeploymentState_PendingCreation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentState(input)
	return &out, nil
}
