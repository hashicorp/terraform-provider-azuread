package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountAppSyncStatus string

const (
	AndroidManagedStoreAccountAppSyncStatus_AndroidForWorkApiError AndroidManagedStoreAccountAppSyncStatus = "androidForWorkApiError"
	AndroidManagedStoreAccountAppSyncStatus_CredentialsNotValid    AndroidManagedStoreAccountAppSyncStatus = "credentialsNotValid"
	AndroidManagedStoreAccountAppSyncStatus_ManagementServiceError AndroidManagedStoreAccountAppSyncStatus = "managementServiceError"
	AndroidManagedStoreAccountAppSyncStatus_None                   AndroidManagedStoreAccountAppSyncStatus = "none"
	AndroidManagedStoreAccountAppSyncStatus_Success                AndroidManagedStoreAccountAppSyncStatus = "success"
	AndroidManagedStoreAccountAppSyncStatus_UnknownError           AndroidManagedStoreAccountAppSyncStatus = "unknownError"
)

func PossibleValuesForAndroidManagedStoreAccountAppSyncStatus() []string {
	return []string{
		string(AndroidManagedStoreAccountAppSyncStatus_AndroidForWorkApiError),
		string(AndroidManagedStoreAccountAppSyncStatus_CredentialsNotValid),
		string(AndroidManagedStoreAccountAppSyncStatus_ManagementServiceError),
		string(AndroidManagedStoreAccountAppSyncStatus_None),
		string(AndroidManagedStoreAccountAppSyncStatus_Success),
		string(AndroidManagedStoreAccountAppSyncStatus_UnknownError),
	}
}

func (s *AndroidManagedStoreAccountAppSyncStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAccountAppSyncStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAccountAppSyncStatus(input string) (*AndroidManagedStoreAccountAppSyncStatus, error) {
	vals := map[string]AndroidManagedStoreAccountAppSyncStatus{
		"androidforworkapierror": AndroidManagedStoreAccountAppSyncStatus_AndroidForWorkApiError,
		"credentialsnotvalid":    AndroidManagedStoreAccountAppSyncStatus_CredentialsNotValid,
		"managementserviceerror": AndroidManagedStoreAccountAppSyncStatus_ManagementServiceError,
		"none":                   AndroidManagedStoreAccountAppSyncStatus_None,
		"success":                AndroidManagedStoreAccountAppSyncStatus_Success,
		"unknownerror":           AndroidManagedStoreAccountAppSyncStatus_UnknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAccountAppSyncStatus(input)
	return &out, nil
}
