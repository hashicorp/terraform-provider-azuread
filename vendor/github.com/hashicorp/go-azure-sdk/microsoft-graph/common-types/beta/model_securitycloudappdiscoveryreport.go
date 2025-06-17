package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityCloudAppDiscoveryReport{}

type SecurityCloudAppDiscoveryReport struct {
	// Use 1 if the machine information is anonymized; otherwise use 0.
	AnonymizeMachineData *bool `json:"anonymizeMachineData,omitempty"`

	// Use 1 if the user information is anonymized; otherwise use 0.
	AnonymizeUserData *bool `json:"anonymizeUserData,omitempty"`

	// The date in the format specified. The Timestamp represents date and time information using ISO 8601 format and is
	// always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// A comment or description for the report.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the continuous report.
	DisplayName *string `json:"displayName,omitempty"`

	// Use 1 for a snapshot report; otherwise use 0.
	IsSnapshotReport *bool `json:"isSnapshotReport,omitempty"`

	// The date when the data was last received. The Timestamp represents date and time information using ISO 8601 format
	// and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastDataReceivedDateTime nullable.Type[string] `json:"lastDataReceivedDateTime,omitempty"`

	// The date when the continuous report was last modified. The Timestamp represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The applicable log data provider. Possible values are: barracuda, bluecoat, checkpoint, ciscoAsa, ciscoIronportProxy,
	// fortigate, paloAlto, squid, zscaler, mcafeeSwg, ciscoScanSafe, juniperSrx, sophosSg, websenseV75, websenseSiemCef,
	// machineZoneMeraki, squidNative, ciscoFwsm, microsoftIsaW3C, sonicwall, sophosCyberoam, clavister, customParser,
	// juniperSsg, zscalerQradar, juniperSrxSd, juniperSrxWelf, microsoftConditionalAppAccess, ciscoAsaFirepower,
	// genericCef, genericLeef, genericW3C, iFilter, checkpointXml, checkpointSmartViewTracker, barracudaNextGenFw,
	// barracudaNextGenFwWeblog, microsoftDefenderForEndpoint, zscalerCef, sophosXg, iboss, forcepoint, fortios,
	// ciscoIronportWsaIi, paloAltoLeef, forcepointLeef, stormshield, contentkeeper, ciscoIronportWsaIii, checkpointCef,
	// corrata, ciscoFirepowerV6, menloSecurityCef, watchguardXtm, openSystemsSecureWebGateway, wandera, unknownFutureValue.
	LogDataProvider *SecurityLogDataProvider `json:"logDataProvider,omitempty"`

	// The count of log files history.
	LogFileCount nullable.Type[int64] `json:"logFileCount,omitempty"`

	// The applicable receiver protocol. Possible values are: ftp, ftps, syslogUdp, syslogTcp, syslogTls,
	// unknownFutureValue.
	ReceiverProtocol *SecurityReceiverProtocol `json:"receiverProtocol,omitempty"`

	// The supported entity type. Possible values are: userName, ipAddress, machineName, other, unknown, unknownFutureValue.
	SupportedEntityTypes *[]SecurityEntityType `json:"supportedEntityTypes,omitempty"`

	// The supported traffic type. Possible values are: downloadedBytes, uploadedBytes, unknown, unknownFutureValue.
	SupportedTrafficTypes *[]SecurityTrafficType `json:"supportedTrafficTypes,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityCloudAppDiscoveryReport) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityCloudAppDiscoveryReport{}

func (s SecurityCloudAppDiscoveryReport) MarshalJSON() ([]byte, error) {
	type wrapper SecurityCloudAppDiscoveryReport
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityCloudAppDiscoveryReport: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityCloudAppDiscoveryReport: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.cloudAppDiscoveryReport"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityCloudAppDiscoveryReport: %+v", err)
	}

	return encoded, nil
}
