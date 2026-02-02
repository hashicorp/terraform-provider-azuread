package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceRemoteAction string

const (
	ManagedDeviceRemoteAction_ActivateDeviceEsim                        ManagedDeviceRemoteAction = "activateDeviceEsim"
	ManagedDeviceRemoteAction_CollectDiagnostics                        ManagedDeviceRemoteAction = "collectDiagnostics"
	ManagedDeviceRemoteAction_CustomTextNotification                    ManagedDeviceRemoteAction = "customTextNotification"
	ManagedDeviceRemoteAction_Delete                                    ManagedDeviceRemoteAction = "delete"
	ManagedDeviceRemoteAction_Deprovision                               ManagedDeviceRemoteAction = "deprovision"
	ManagedDeviceRemoteAction_Disable                                   ManagedDeviceRemoteAction = "disable"
	ManagedDeviceRemoteAction_FullScan                                  ManagedDeviceRemoteAction = "fullScan"
	ManagedDeviceRemoteAction_InitiateDeviceAttestation                 ManagedDeviceRemoteAction = "initiateDeviceAttestation"
	ManagedDeviceRemoteAction_InitiateMobileDeviceManagementKeyRecovery ManagedDeviceRemoteAction = "initiateMobileDeviceManagementKeyRecovery"
	ManagedDeviceRemoteAction_InitiateOnDemandProactiveRemediation      ManagedDeviceRemoteAction = "initiateOnDemandProactiveRemediation"
	ManagedDeviceRemoteAction_MoveDeviceToOrganizationalUnit            ManagedDeviceRemoteAction = "moveDeviceToOrganizationalUnit"
	ManagedDeviceRemoteAction_QuickScan                                 ManagedDeviceRemoteAction = "quickScan"
	ManagedDeviceRemoteAction_RebootNow                                 ManagedDeviceRemoteAction = "rebootNow"
	ManagedDeviceRemoteAction_Reenable                                  ManagedDeviceRemoteAction = "reenable"
	ManagedDeviceRemoteAction_Retire                                    ManagedDeviceRemoteAction = "retire"
	ManagedDeviceRemoteAction_SetDeviceName                             ManagedDeviceRemoteAction = "setDeviceName"
	ManagedDeviceRemoteAction_SignatureUpdate                           ManagedDeviceRemoteAction = "signatureUpdate"
	ManagedDeviceRemoteAction_SyncDevice                                ManagedDeviceRemoteAction = "syncDevice"
	ManagedDeviceRemoteAction_Wipe                                      ManagedDeviceRemoteAction = "wipe"
)

func PossibleValuesForManagedDeviceRemoteAction() []string {
	return []string{
		string(ManagedDeviceRemoteAction_ActivateDeviceEsim),
		string(ManagedDeviceRemoteAction_CollectDiagnostics),
		string(ManagedDeviceRemoteAction_CustomTextNotification),
		string(ManagedDeviceRemoteAction_Delete),
		string(ManagedDeviceRemoteAction_Deprovision),
		string(ManagedDeviceRemoteAction_Disable),
		string(ManagedDeviceRemoteAction_FullScan),
		string(ManagedDeviceRemoteAction_InitiateDeviceAttestation),
		string(ManagedDeviceRemoteAction_InitiateMobileDeviceManagementKeyRecovery),
		string(ManagedDeviceRemoteAction_InitiateOnDemandProactiveRemediation),
		string(ManagedDeviceRemoteAction_MoveDeviceToOrganizationalUnit),
		string(ManagedDeviceRemoteAction_QuickScan),
		string(ManagedDeviceRemoteAction_RebootNow),
		string(ManagedDeviceRemoteAction_Reenable),
		string(ManagedDeviceRemoteAction_Retire),
		string(ManagedDeviceRemoteAction_SetDeviceName),
		string(ManagedDeviceRemoteAction_SignatureUpdate),
		string(ManagedDeviceRemoteAction_SyncDevice),
		string(ManagedDeviceRemoteAction_Wipe),
	}
}

func (s *ManagedDeviceRemoteAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceRemoteAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceRemoteAction(input string) (*ManagedDeviceRemoteAction, error) {
	vals := map[string]ManagedDeviceRemoteAction{
		"activatedeviceesim":        ManagedDeviceRemoteAction_ActivateDeviceEsim,
		"collectdiagnostics":        ManagedDeviceRemoteAction_CollectDiagnostics,
		"customtextnotification":    ManagedDeviceRemoteAction_CustomTextNotification,
		"delete":                    ManagedDeviceRemoteAction_Delete,
		"deprovision":               ManagedDeviceRemoteAction_Deprovision,
		"disable":                   ManagedDeviceRemoteAction_Disable,
		"fullscan":                  ManagedDeviceRemoteAction_FullScan,
		"initiatedeviceattestation": ManagedDeviceRemoteAction_InitiateDeviceAttestation,
		"initiatemobiledevicemanagementkeyrecovery": ManagedDeviceRemoteAction_InitiateMobileDeviceManagementKeyRecovery,
		"initiateondemandproactiveremediation":      ManagedDeviceRemoteAction_InitiateOnDemandProactiveRemediation,
		"movedevicetoorganizationalunit":            ManagedDeviceRemoteAction_MoveDeviceToOrganizationalUnit,
		"quickscan":                                 ManagedDeviceRemoteAction_QuickScan,
		"rebootnow":                                 ManagedDeviceRemoteAction_RebootNow,
		"reenable":                                  ManagedDeviceRemoteAction_Reenable,
		"retire":                                    ManagedDeviceRemoteAction_Retire,
		"setdevicename":                             ManagedDeviceRemoteAction_SetDeviceName,
		"signatureupdate":                           ManagedDeviceRemoteAction_SignatureUpdate,
		"syncdevice":                                ManagedDeviceRemoteAction_SyncDevice,
		"wipe":                                      ManagedDeviceRemoteAction_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceRemoteAction(input)
	return &out, nil
}
