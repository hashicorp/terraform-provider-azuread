package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Process struct {
	// User account identifier (user account context the process ran under) for example, AccountName, SID, and so on.
	AccountName nullable.Type[string] `json:"accountName,omitempty"`

	// The full process invocation commandline including all parameters.
	CommandLine nullable.Type[string] `json:"commandLine,omitempty"`

	// Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Complex type containing file hashes (cryptographic and location-sensitive).
	FileHash *FileHash `json:"fileHash,omitempty"`

	// The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system.
	IntegrityLevel *ProcessIntegrityLevel `json:"integrityLevel,omitempty"`

	// True if the process is elevated.
	IsElevated nullable.Type[bool] `json:"isElevated,omitempty"`

	// The name of the process' Image file.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ParentProcessCreatedDateTime nullable.Type[string] `json:"parentProcessCreatedDateTime,omitempty"`

	// The Process ID (PID) of the parent process.
	ParentProcessId nullable.Type[int64] `json:"parentProcessId,omitempty"`

	// The name of the image file of the parent process.
	ParentProcessName nullable.Type[string] `json:"parentProcessName,omitempty"`

	// Full path, including filename.
	Path nullable.Type[string] `json:"path,omitempty"`

	// The Process ID (PID) of the process.
	ProcessId nullable.Type[int64] `json:"processId,omitempty"`
}
