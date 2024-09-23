package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkSyncStatus string

const (
	AndroidForWorkSyncStatus_AndroidForWorkApiError AndroidForWorkSyncStatus = "androidForWorkApiError"
	AndroidForWorkSyncStatus_CredentialsNotValid    AndroidForWorkSyncStatus = "credentialsNotValid"
	AndroidForWorkSyncStatus_ManagementServiceError AndroidForWorkSyncStatus = "managementServiceError"
	AndroidForWorkSyncStatus_None                   AndroidForWorkSyncStatus = "none"
	AndroidForWorkSyncStatus_Success                AndroidForWorkSyncStatus = "success"
	AndroidForWorkSyncStatus_UnknownError           AndroidForWorkSyncStatus = "unknownError"
)

func PossibleValuesForAndroidForWorkSyncStatus() []string {
	return []string{
		string(AndroidForWorkSyncStatus_AndroidForWorkApiError),
		string(AndroidForWorkSyncStatus_CredentialsNotValid),
		string(AndroidForWorkSyncStatus_ManagementServiceError),
		string(AndroidForWorkSyncStatus_None),
		string(AndroidForWorkSyncStatus_Success),
		string(AndroidForWorkSyncStatus_UnknownError),
	}
}

func (s *AndroidForWorkSyncStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkSyncStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkSyncStatus(input string) (*AndroidForWorkSyncStatus, error) {
	vals := map[string]AndroidForWorkSyncStatus{
		"androidforworkapierror": AndroidForWorkSyncStatus_AndroidForWorkApiError,
		"credentialsnotvalid":    AndroidForWorkSyncStatus_CredentialsNotValid,
		"managementserviceerror": AndroidForWorkSyncStatus_ManagementServiceError,
		"none":                   AndroidForWorkSyncStatus_None,
		"success":                AndroidForWorkSyncStatus_Success,
		"unknownerror":           AndroidForWorkSyncStatus_UnknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkSyncStatus(input)
	return &out, nil
}
