package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppContentFileUploadState string

const (
	MobileAppContentFileUploadState_AzureStorageUriRenewalFailed   MobileAppContentFileUploadState = "azureStorageUriRenewalFailed"
	MobileAppContentFileUploadState_AzureStorageUriRenewalPending  MobileAppContentFileUploadState = "azureStorageUriRenewalPending"
	MobileAppContentFileUploadState_AzureStorageUriRenewalSuccess  MobileAppContentFileUploadState = "azureStorageUriRenewalSuccess"
	MobileAppContentFileUploadState_AzureStorageUriRenewalTimedOut MobileAppContentFileUploadState = "azureStorageUriRenewalTimedOut"
	MobileAppContentFileUploadState_AzureStorageUriRequestFailed   MobileAppContentFileUploadState = "azureStorageUriRequestFailed"
	MobileAppContentFileUploadState_AzureStorageUriRequestPending  MobileAppContentFileUploadState = "azureStorageUriRequestPending"
	MobileAppContentFileUploadState_AzureStorageUriRequestSuccess  MobileAppContentFileUploadState = "azureStorageUriRequestSuccess"
	MobileAppContentFileUploadState_AzureStorageUriRequestTimedOut MobileAppContentFileUploadState = "azureStorageUriRequestTimedOut"
	MobileAppContentFileUploadState_CommitFileFailed               MobileAppContentFileUploadState = "commitFileFailed"
	MobileAppContentFileUploadState_CommitFilePending              MobileAppContentFileUploadState = "commitFilePending"
	MobileAppContentFileUploadState_CommitFileSuccess              MobileAppContentFileUploadState = "commitFileSuccess"
	MobileAppContentFileUploadState_CommitFileTimedOut             MobileAppContentFileUploadState = "commitFileTimedOut"
	MobileAppContentFileUploadState_Error                          MobileAppContentFileUploadState = "error"
	MobileAppContentFileUploadState_Success                        MobileAppContentFileUploadState = "success"
	MobileAppContentFileUploadState_TransientError                 MobileAppContentFileUploadState = "transientError"
	MobileAppContentFileUploadState_Unknown                        MobileAppContentFileUploadState = "unknown"
)

func PossibleValuesForMobileAppContentFileUploadState() []string {
	return []string{
		string(MobileAppContentFileUploadState_AzureStorageUriRenewalFailed),
		string(MobileAppContentFileUploadState_AzureStorageUriRenewalPending),
		string(MobileAppContentFileUploadState_AzureStorageUriRenewalSuccess),
		string(MobileAppContentFileUploadState_AzureStorageUriRenewalTimedOut),
		string(MobileAppContentFileUploadState_AzureStorageUriRequestFailed),
		string(MobileAppContentFileUploadState_AzureStorageUriRequestPending),
		string(MobileAppContentFileUploadState_AzureStorageUriRequestSuccess),
		string(MobileAppContentFileUploadState_AzureStorageUriRequestTimedOut),
		string(MobileAppContentFileUploadState_CommitFileFailed),
		string(MobileAppContentFileUploadState_CommitFilePending),
		string(MobileAppContentFileUploadState_CommitFileSuccess),
		string(MobileAppContentFileUploadState_CommitFileTimedOut),
		string(MobileAppContentFileUploadState_Error),
		string(MobileAppContentFileUploadState_Success),
		string(MobileAppContentFileUploadState_TransientError),
		string(MobileAppContentFileUploadState_Unknown),
	}
}

func (s *MobileAppContentFileUploadState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppContentFileUploadState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppContentFileUploadState(input string) (*MobileAppContentFileUploadState, error) {
	vals := map[string]MobileAppContentFileUploadState{
		"azurestorageurirenewalfailed":   MobileAppContentFileUploadState_AzureStorageUriRenewalFailed,
		"azurestorageurirenewalpending":  MobileAppContentFileUploadState_AzureStorageUriRenewalPending,
		"azurestorageurirenewalsuccess":  MobileAppContentFileUploadState_AzureStorageUriRenewalSuccess,
		"azurestorageurirenewaltimedout": MobileAppContentFileUploadState_AzureStorageUriRenewalTimedOut,
		"azurestorageurirequestfailed":   MobileAppContentFileUploadState_AzureStorageUriRequestFailed,
		"azurestorageurirequestpending":  MobileAppContentFileUploadState_AzureStorageUriRequestPending,
		"azurestorageurirequestsuccess":  MobileAppContentFileUploadState_AzureStorageUriRequestSuccess,
		"azurestorageurirequesttimedout": MobileAppContentFileUploadState_AzureStorageUriRequestTimedOut,
		"commitfilefailed":               MobileAppContentFileUploadState_CommitFileFailed,
		"commitfilepending":              MobileAppContentFileUploadState_CommitFilePending,
		"commitfilesuccess":              MobileAppContentFileUploadState_CommitFileSuccess,
		"commitfiletimedout":             MobileAppContentFileUploadState_CommitFileTimedOut,
		"error":                          MobileAppContentFileUploadState_Error,
		"success":                        MobileAppContentFileUploadState_Success,
		"transienterror":                 MobileAppContentFileUploadState_TransientError,
		"unknown":                        MobileAppContentFileUploadState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppContentFileUploadState(input)
	return &out, nil
}
