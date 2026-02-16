package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAgentHealthCheckErrorType string

const (
	CloudPCAgentHealthCheckErrorType_AgentCheckHeartbeatLost                           CloudPCAgentHealthCheckErrorType = "agentCheckHeartbeatLost"
	CloudPCAgentHealthCheckErrorType_AgentCheckNotExisted                              CloudPCAgentHealthCheckErrorType = "agentCheckNotExisted"
	CloudPCAgentHealthCheckErrorType_AgentCheckNotRunning                              CloudPCAgentHealthCheckErrorType = "agentCheckNotRunning"
	CloudPCAgentHealthCheckErrorType_AgentCheckOldVersion                              CloudPCAgentHealthCheckErrorType = "agentCheckOldVersion"
	CloudPCAgentHealthCheckErrorType_CommunicationCheckChannelDowngraded               CloudPCAgentHealthCheckErrorType = "communicationCheckChannelDowngraded"
	CloudPCAgentHealthCheckErrorType_CommunicationCheckNotAvailable                    CloudPCAgentHealthCheckErrorType = "communicationCheckNotAvailable"
	CloudPCAgentHealthCheckErrorType_DeviceStatusCheckNotAvailable                     CloudPCAgentHealthCheckErrorType = "deviceStatusCheckNotAvailable"
	CloudPCAgentHealthCheckErrorType_DeviceStatusCheckNotRunning                       CloudPCAgentHealthCheckErrorType = "deviceStatusCheckNotRunning"
	CloudPCAgentHealthCheckErrorType_FunctionalityCheckPowerShellNotRunnable           CloudPCAgentHealthCheckErrorType = "functionalityCheckPowerShellNotRunnable"
	CloudPCAgentHealthCheckErrorType_Healthy                                           CloudPCAgentHealthCheckErrorType = "healthy"
	CloudPCAgentHealthCheckErrorType_InstallationCheckFoundErrors                      CloudPCAgentHealthCheckErrorType = "installationCheckFoundErrors"
	CloudPCAgentHealthCheckErrorType_InstallationCheckMsiFileNotExecutable             CloudPCAgentHealthCheckErrorType = "installationCheckMsiFileNotExecutable"
	CloudPCAgentHealthCheckErrorType_InstallationFailedToDownloadMaterials             CloudPCAgentHealthCheckErrorType = "installationFailedToDownloadMaterials"
	CloudPCAgentHealthCheckErrorType_InternalAgentUnknownError                         CloudPCAgentHealthCheckErrorType = "internalAgentUnknownError"
	CloudPCAgentHealthCheckErrorType_NetworkAvailabilityCheckRequiredUrlsNotAccessible CloudPCAgentHealthCheckErrorType = "networkAvailabilityCheckRequiredUrlsNotAccessible"
	CloudPCAgentHealthCheckErrorType_ResourceAvailabilityCheckDiskSpaceNotEnough       CloudPCAgentHealthCheckErrorType = "resourceAvailabilityCheckDiskSpaceNotEnough"
)

func PossibleValuesForCloudPCAgentHealthCheckErrorType() []string {
	return []string{
		string(CloudPCAgentHealthCheckErrorType_AgentCheckHeartbeatLost),
		string(CloudPCAgentHealthCheckErrorType_AgentCheckNotExisted),
		string(CloudPCAgentHealthCheckErrorType_AgentCheckNotRunning),
		string(CloudPCAgentHealthCheckErrorType_AgentCheckOldVersion),
		string(CloudPCAgentHealthCheckErrorType_CommunicationCheckChannelDowngraded),
		string(CloudPCAgentHealthCheckErrorType_CommunicationCheckNotAvailable),
		string(CloudPCAgentHealthCheckErrorType_DeviceStatusCheckNotAvailable),
		string(CloudPCAgentHealthCheckErrorType_DeviceStatusCheckNotRunning),
		string(CloudPCAgentHealthCheckErrorType_FunctionalityCheckPowerShellNotRunnable),
		string(CloudPCAgentHealthCheckErrorType_Healthy),
		string(CloudPCAgentHealthCheckErrorType_InstallationCheckFoundErrors),
		string(CloudPCAgentHealthCheckErrorType_InstallationCheckMsiFileNotExecutable),
		string(CloudPCAgentHealthCheckErrorType_InstallationFailedToDownloadMaterials),
		string(CloudPCAgentHealthCheckErrorType_InternalAgentUnknownError),
		string(CloudPCAgentHealthCheckErrorType_NetworkAvailabilityCheckRequiredUrlsNotAccessible),
		string(CloudPCAgentHealthCheckErrorType_ResourceAvailabilityCheckDiskSpaceNotEnough),
	}
}

func (s *CloudPCAgentHealthCheckErrorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAgentHealthCheckErrorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAgentHealthCheckErrorType(input string) (*CloudPCAgentHealthCheckErrorType, error) {
	vals := map[string]CloudPCAgentHealthCheckErrorType{
		"agentcheckheartbeatlost":                 CloudPCAgentHealthCheckErrorType_AgentCheckHeartbeatLost,
		"agentchecknotexisted":                    CloudPCAgentHealthCheckErrorType_AgentCheckNotExisted,
		"agentchecknotrunning":                    CloudPCAgentHealthCheckErrorType_AgentCheckNotRunning,
		"agentcheckoldversion":                    CloudPCAgentHealthCheckErrorType_AgentCheckOldVersion,
		"communicationcheckchanneldowngraded":     CloudPCAgentHealthCheckErrorType_CommunicationCheckChannelDowngraded,
		"communicationchecknotavailable":          CloudPCAgentHealthCheckErrorType_CommunicationCheckNotAvailable,
		"devicestatuschecknotavailable":           CloudPCAgentHealthCheckErrorType_DeviceStatusCheckNotAvailable,
		"devicestatuschecknotrunning":             CloudPCAgentHealthCheckErrorType_DeviceStatusCheckNotRunning,
		"functionalitycheckpowershellnotrunnable": CloudPCAgentHealthCheckErrorType_FunctionalityCheckPowerShellNotRunnable,
		"healthy":                                           CloudPCAgentHealthCheckErrorType_Healthy,
		"installationcheckfounderrors":                      CloudPCAgentHealthCheckErrorType_InstallationCheckFoundErrors,
		"installationcheckmsifilenotexecutable":             CloudPCAgentHealthCheckErrorType_InstallationCheckMsiFileNotExecutable,
		"installationfailedtodownloadmaterials":             CloudPCAgentHealthCheckErrorType_InstallationFailedToDownloadMaterials,
		"internalagentunknownerror":                         CloudPCAgentHealthCheckErrorType_InternalAgentUnknownError,
		"networkavailabilitycheckrequiredurlsnotaccessible": CloudPCAgentHealthCheckErrorType_NetworkAvailabilityCheckRequiredUrlsNotAccessible,
		"resourceavailabilitycheckdiskspacenotenough":       CloudPCAgentHealthCheckErrorType_ResourceAvailabilityCheckDiskSpaceNotEnough,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAgentHealthCheckErrorType(input)
	return &out, nil
}
