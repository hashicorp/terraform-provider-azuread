package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryKeyState struct {
	// A Windows registry hive : HKEYCURRENTCONFIG HKEYCURRENTUSER HKEYLOCALMACHINE/SAM HKEYLOCALMACHINE/Security
	// HKEYLOCALMACHINE/Software HKEYLOCALMACHINE/System HKEY_USERS/.Default. Possible values are: unknown, currentConfig,
	// currentUser, localMachineSam, localMachineSecurity, localMachineSoftware, localMachineSystem, usersDefault.
	Hive *RegistryHive `json:"hive,omitempty"`

	// Current (i.e. changed) registry key (excludes HIVE).
	Key nullable.Type[string] `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Previous (i.e. before changed) registry key (excludes HIVE).
	OldKey nullable.Type[string] `json:"oldKey,omitempty"`

	// Previous (i.e. before changed) registry key value data (contents).
	OldValueData nullable.Type[string] `json:"oldValueData,omitempty"`

	// Previous (i.e. before changed) registry key value name.
	OldValueName nullable.Type[string] `json:"oldValueName,omitempty"`

	// Operation that changed the registry key name and/or value. Possible values are: unknown, create, modify, delete.
	Operation *RegistryOperation `json:"operation,omitempty"`

	// Process ID (PID) of the process that modified the registry key (process details will appear in the alert 'processes'
	// collection).
	ProcessId nullable.Type[int64] `json:"processId,omitempty"`

	// Current (i.e. changed) registry key value data (contents).
	ValueData nullable.Type[string] `json:"valueData,omitempty"`

	// Current (i.e. changed) registry key value name
	ValueName nullable.Type[string] `json:"valueName,omitempty"`

	// Registry key value type REGBINARY REGDWORD REGDWORDLITTLEENDIAN REGDWORDBIGENDIANREGEXPANDSZ REGLINK REGMULTISZ
	// REGNONE REGQWORD REGQWORDLITTLEENDIAN REG_SZ Possible values are: unknown, binary, dword, dwordLittleEndian,
	// dwordBigEndian, expandSz, link, multiSz, none, qword, qwordlittleEndian, sz.
	ValueType *RegistryValueType `json:"valueType,omitempty"`
}
