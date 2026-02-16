package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAuditData interface {
	SecurityAuditData() BaseSecurityAuditDataImpl
}

var _ SecurityAuditData = BaseSecurityAuditDataImpl{}

type BaseSecurityAuditDataImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSecurityAuditDataImpl) SecurityAuditData() BaseSecurityAuditDataImpl {
	return s
}

var _ SecurityAuditData = RawSecurityAuditDataImpl{}

// RawSecurityAuditDataImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityAuditDataImpl struct {
	securityAuditData BaseSecurityAuditDataImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawSecurityAuditDataImpl) SecurityAuditData() BaseSecurityAuditDataImpl {
	return s.securityAuditData
}

func UnmarshalSecurityAuditDataImplementation(input []byte) (SecurityAuditData, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityAuditData into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.aadRiskDetectionAuditRecord") {
		var out SecurityAadRiskDetectionAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAadRiskDetectionAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.aedAuditRecord") {
		var out SecurityAedAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAedAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.aiAppInteractionAuditRecord") {
		var out SecurityAiAppInteractionAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAiAppInteractionAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.aipFileDeleted") {
		var out SecurityAipFileDeleted
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAipFileDeleted: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.aipHeartBeat") {
		var out SecurityAipHeartBeat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAipHeartBeat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.aipProtectionActionLogRequest") {
		var out SecurityAipProtectionActionLogRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAipProtectionActionLogRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.aipScannerDiscoverEvent") {
		var out SecurityAipScannerDiscoverEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAipScannerDiscoverEvent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.aipSensitivityLabelActionLogRequest") {
		var out SecurityAipSensitivityLabelActionLogRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAipSensitivityLabelActionLogRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.airAdminActionInvestigationData") {
		var out SecurityAirAdminActionInvestigationData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAirAdminActionInvestigationData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.airInvestigationData") {
		var out SecurityAirInvestigationData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAirInvestigationData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.airManualInvestigationData") {
		var out SecurityAirManualInvestigationData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAirManualInvestigationData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.attackSimAdminAuditRecord") {
		var out SecurityAttackSimAdminAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAttackSimAdminAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.auditSearchAuditRecord") {
		var out SecurityAuditSearchAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAuditSearchAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.azureActiveDirectoryAccountLogonAuditRecord") {
		var out SecurityAzureActiveDirectoryAccountLogonAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAzureActiveDirectoryAccountLogonAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.azureActiveDirectoryAuditRecord") {
		var out SecurityAzureActiveDirectoryAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAzureActiveDirectoryAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.azureActiveDirectoryBaseAuditRecord") {
		var out SecurityAzureActiveDirectoryBaseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAzureActiveDirectoryBaseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.azureActiveDirectoryStsLogonAuditRecord") {
		var out SecurityAzureActiveDirectoryStsLogonAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAzureActiveDirectoryStsLogonAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.campaignAuditRecord") {
		var out SecurityCampaignAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCampaignAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.caseAuditRecord") {
		var out SecurityCaseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCaseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.caseInvestigation") {
		var out SecurityCaseInvestigation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCaseInvestigation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cdpColdCrawlStatusRecord") {
		var out SecurityCdpColdCrawlStatusRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCdpColdCrawlStatusRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cdpContentExplorerAggregateRecord") {
		var out SecurityCdpContentExplorerAggregateRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCdpContentExplorerAggregateRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cdpDlpSensitiveAuditRecord") {
		var out SecurityCdpDlpSensitiveAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCdpDlpSensitiveAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cdpDlpSensitiveEndpointAuditRecord") {
		var out SecurityCdpDlpSensitiveEndpointAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCdpDlpSensitiveEndpointAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cdpLogRecord") {
		var out SecurityCdpLogRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCdpLogRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cdpOcrBillingRecord") {
		var out SecurityCdpOcrBillingRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCdpOcrBillingRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cdpResourceScopeChangeEventRecord") {
		var out SecurityCdpResourceScopeChangeEventRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCdpResourceScopeChangeEventRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cernerSMSLinkRecord") {
		var out SecurityCernerSMSLinkRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCernerSMSLinkRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cernerSMSSettingsUpdateRecord") {
		var out SecurityCernerSMSSettingsUpdateRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCernerSMSSettingsUpdateRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cernerSMSUnlinkRecord") {
		var out SecurityCernerSMSUnlinkRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCernerSMSUnlinkRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceConnectorAuditRecord") {
		var out SecurityComplianceConnectorAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceConnectorAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDLMExchangeAuditRecord") {
		var out SecurityComplianceDLMExchangeAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDLMExchangeAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDLMSharePointAuditRecord") {
		var out SecurityComplianceDLMSharePointAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDLMSharePointAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpApplicationsAuditRecord") {
		var out SecurityComplianceDlpApplicationsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpApplicationsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpApplicationsClassificationAuditRecord") {
		var out SecurityComplianceDlpApplicationsClassificationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpApplicationsClassificationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpBaseAuditRecord") {
		var out SecurityComplianceDlpBaseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpBaseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpClassificationBaseAuditRecord") {
		var out SecurityComplianceDlpClassificationBaseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpClassificationBaseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpClassificationBaseCdpRecord") {
		var out SecurityComplianceDlpClassificationBaseCdpRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpClassificationBaseCdpRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpEndpointAuditRecord") {
		var out SecurityComplianceDlpEndpointAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpEndpointAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpEndpointDiscoveryAuditRecord") {
		var out SecurityComplianceDlpEndpointDiscoveryAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpEndpointDiscoveryAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpExchangeAuditRecord") {
		var out SecurityComplianceDlpExchangeAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpExchangeAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpExchangeClassificationAuditRecord") {
		var out SecurityComplianceDlpExchangeClassificationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpExchangeClassificationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpExchangeClassificationCdpRecord") {
		var out SecurityComplianceDlpExchangeClassificationCdpRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpExchangeClassificationCdpRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpExchangeDiscoveryAuditRecord") {
		var out SecurityComplianceDlpExchangeDiscoveryAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpExchangeDiscoveryAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpSharePointAuditRecord") {
		var out SecurityComplianceDlpSharePointAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpSharePointAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpSharePointClassificationAuditRecord") {
		var out SecurityComplianceDlpSharePointClassificationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpSharePointClassificationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceDlpSharePointClassificationExtendedAuditRecord") {
		var out SecurityComplianceDlpSharePointClassificationExtendedAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceDlpSharePointClassificationExtendedAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceManagerActionRecord") {
		var out SecurityComplianceManagerActionRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceManagerActionRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceSupervisionBaseAuditRecord") {
		var out SecurityComplianceSupervisionBaseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceSupervisionBaseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.complianceSupervisionExchangeAuditRecord") {
		var out SecurityComplianceSupervisionExchangeAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityComplianceSupervisionExchangeAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.consumptionResourceAuditRecord") {
		var out SecurityConsumptionResourceAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityConsumptionResourceAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.copilotInteractionAuditRecord") {
		var out SecurityCopilotInteractionAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCopilotInteractionAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.coreReportingSettingsAuditRecord") {
		var out SecurityCoreReportingSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCoreReportingSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cortanaBriefingAuditRecord") {
		var out SecurityCortanaBriefingAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCortanaBriefingAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cpsCommonPolicyAuditRecord") {
		var out SecurityCpsCommonPolicyAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCpsCommonPolicyAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.cpsPolicyConfigAuditRecord") {
		var out SecurityCpsPolicyConfigAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCpsPolicyConfigAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.crmBaseAuditRecord") {
		var out SecurityCrmBaseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCrmBaseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.crmEntityOperationAuditRecord") {
		var out SecurityCrmEntityOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCrmEntityOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.customerKeyServiceEncryptionAuditRecord") {
		var out SecurityCustomerKeyServiceEncryptionAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCustomerKeyServiceEncryptionAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataCenterSecurityBaseAuditRecord") {
		var out SecurityDataCenterSecurityBaseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataCenterSecurityBaseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataCenterSecurityCmdletAuditRecord") {
		var out SecurityDataCenterSecurityCmdletAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataCenterSecurityCmdletAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataGovernanceAuditRecord") {
		var out SecurityDataGovernanceAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataGovernanceAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataInsightsRestApiAuditRecord") {
		var out SecurityDataInsightsRestApiAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataInsightsRestApiAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataLakeExportOperationAuditRecord") {
		var out SecurityDataLakeExportOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataLakeExportOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dataShareOperationAuditRecord") {
		var out SecurityDataShareOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDataShareOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.defaultAuditData") {
		var out SecurityDefaultAuditData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDefaultAuditData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.defenderSecurityAlertBaseRecord") {
		var out SecurityDefenderSecurityAlertBaseRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDefenderSecurityAlertBaseRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.deleteCertificateRecord") {
		var out SecurityDeleteCertificateRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDeleteCertificateRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.disableConsentRecord") {
		var out SecurityDisableConsentRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDisableConsentRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.discoveryAuditRecord") {
		var out SecurityDiscoveryAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDiscoveryAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dlpEndpointAuditRecord") {
		var out SecurityDlpEndpointAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDlpEndpointAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dlpSensitiveInformationTypeCmdletRecord") {
		var out SecurityDlpSensitiveInformationTypeCmdletRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDlpSensitiveInformationTypeCmdletRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dlpSensitiveInformationTypeRulePackageCmdletRecord") {
		var out SecurityDlpSensitiveInformationTypeRulePackageCmdletRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDlpSensitiveInformationTypeRulePackageCmdletRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.downloadCertificateRecord") {
		var out SecurityDownloadCertificateRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDownloadCertificateRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.dynamics365BusinessCentralAuditRecord") {
		var out SecurityDynamics365BusinessCentralAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDynamics365BusinessCentralAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.enableConsentRecord") {
		var out SecurityEnableConsentRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEnableConsentRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.epicSMSLinkRecord") {
		var out SecurityEpicSMSLinkRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEpicSMSLinkRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.epicSMSSettingsUpdateRecord") {
		var out SecurityEpicSMSSettingsUpdateRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEpicSMSSettingsUpdateRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.epicSMSUnlinkRecord") {
		var out SecurityEpicSMSUnlinkRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEpicSMSUnlinkRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.exchangeAdminAuditRecord") {
		var out SecurityExchangeAdminAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityExchangeAdminAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.exchangeAggregatedMailboxAuditRecord") {
		var out SecurityExchangeAggregatedMailboxAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityExchangeAggregatedMailboxAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.exchangeAggregatedOperationRecord") {
		var out SecurityExchangeAggregatedOperationRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityExchangeAggregatedOperationRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.exchangeMailboxAuditBaseRecord") {
		var out SecurityExchangeMailboxAuditBaseRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityExchangeMailboxAuditBaseRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.exchangeMailboxAuditGroupRecord") {
		var out SecurityExchangeMailboxAuditGroupRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityExchangeMailboxAuditGroupRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.exchangeMailboxAuditRecord") {
		var out SecurityExchangeMailboxAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityExchangeMailboxAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fhirBaseUrlAddRecord") {
		var out SecurityFhirBaseUrlAddRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFhirBaseUrlAddRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fhirBaseUrlApproveRecord") {
		var out SecurityFhirBaseUrlApproveRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFhirBaseUrlApproveRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fhirBaseUrlDeleteRecord") {
		var out SecurityFhirBaseUrlDeleteRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFhirBaseUrlDeleteRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.fhirBaseUrlUpdateRecord") {
		var out SecurityFhirBaseUrlUpdateRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFhirBaseUrlUpdateRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.healthcareSignalRecord") {
		var out SecurityHealthcareSignalRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHealthcareSignalRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostedRpaAuditRecord") {
		var out SecurityHostedRpaAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostedRpaAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hrSignalAuditRecord") {
		var out SecurityHrSignalAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHrSignalAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hygieneEventRecord") {
		var out SecurityHygieneEventRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHygieneEventRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.informationBarrierPolicyApplicationAuditRecord") {
		var out SecurityInformationBarrierPolicyApplicationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityInformationBarrierPolicyApplicationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.informationWorkerProtectionAuditRecord") {
		var out SecurityInformationWorkerProtectionAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityInformationWorkerProtectionAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.insiderRiskScopedUserInsightsRecord") {
		var out SecurityInsiderRiskScopedUserInsightsRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityInsiderRiskScopedUserInsightsRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.insiderRiskScopedUsersRecord") {
		var out SecurityInsiderRiskScopedUsersRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityInsiderRiskScopedUsersRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.irmSecurityAlertRecord") {
		var out SecurityIrmSecurityAlertRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIrmSecurityAlertRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.irmUserDefinedDetectionRecord") {
		var out SecurityIrmUserDefinedDetectionRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIrmUserDefinedDetectionRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.kaizalaAuditRecord") {
		var out SecurityKaizalaAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityKaizalaAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.labelAnalyticsAggregateAuditRecord") {
		var out SecurityLabelAnalyticsAggregateAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityLabelAnalyticsAggregateAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.labelContentExplorerAuditRecord") {
		var out SecurityLabelContentExplorerAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityLabelContentExplorerAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.largeContentMetadataAuditRecord") {
		var out SecurityLargeContentMetadataAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityLargeContentMetadataAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.m365ComplianceConnectorAuditRecord") {
		var out SecurityM365ComplianceConnectorAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityM365ComplianceConnectorAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.m365DAADAuditRecord") {
		var out SecurityM365DAADAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityM365DAADAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mS365DCustomDetectionAuditRecord") {
		var out SecurityMS365DCustomDetectionAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMS365DCustomDetectionAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mS365DIncidentAuditRecord") {
		var out SecurityMS365DIncidentAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMS365DIncidentAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mS365DSuppressionRuleAuditRecord") {
		var out SecurityMS365DSuppressionRuleAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMS365DSuppressionRuleAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mailSubmissionData") {
		var out SecurityMailSubmissionData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMailSubmissionData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.managedServicesAuditRecord") {
		var out SecurityManagedServicesAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityManagedServicesAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.managedTenantsAuditRecord") {
		var out SecurityManagedTenantsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityManagedTenantsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mapgAlertsAuditRecord") {
		var out SecurityMapgAlertsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMapgAlertsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mapgOnboardAuditRecord") {
		var out SecurityMapgOnboardAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMapgOnboardAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mapgPolicyAuditRecord") {
		var out SecurityMapgPolicyAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMapgPolicyAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mcasAlertsAuditRecord") {
		var out SecurityMcasAlertsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMcasAlertsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mdaDataSecuritySignalRecord") {
		var out SecurityMdaDataSecuritySignalRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMdaDataSecuritySignalRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mdatpAuditRecord") {
		var out SecurityMdatpAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMdatpAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mdcEventsRecord") {
		var out SecurityMdcEventsRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMdcEventsRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mdiAuditRecord") {
		var out SecurityMdiAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMdiAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.meshWorldsAuditRecord") {
		var out SecurityMeshWorldsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMeshWorldsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoft365BackupBackupItemAuditRecord") {
		var out SecurityMicrosoft365BackupBackupItemAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoft365BackupBackupItemAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoft365BackupBackupPolicyAuditRecord") {
		var out SecurityMicrosoft365BackupBackupPolicyAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoft365BackupBackupPolicyAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoft365BackupRestoreItemAuditRecord") {
		var out SecurityMicrosoft365BackupRestoreItemAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoft365BackupRestoreItemAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoft365BackupRestoreTaskAuditRecord") {
		var out SecurityMicrosoft365BackupRestoreTaskAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoft365BackupRestoreTaskAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftDefenderExpertsBaseAuditRecord") {
		var out SecurityMicrosoftDefenderExpertsBaseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftDefenderExpertsBaseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftDefenderExpertsXDRAuditRecord") {
		var out SecurityMicrosoftDefenderExpertsXDRAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftDefenderExpertsXDRAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftFlowAuditRecord") {
		var out SecurityMicrosoftFlowAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftFlowAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftFormsAuditRecord") {
		var out SecurityMicrosoftFormsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftFormsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftGraphDataConnectConsent") {
		var out SecurityMicrosoftGraphDataConnectConsent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftGraphDataConnectConsent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftGraphDataConnectOperation") {
		var out SecurityMicrosoftGraphDataConnectOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftGraphDataConnectOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftPurviewDataMapOperationRecord") {
		var out SecurityMicrosoftPurviewDataMapOperationRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftPurviewDataMapOperationRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftPurviewMetadataPolicyOperationRecord") {
		var out SecurityMicrosoftPurviewMetadataPolicyOperationRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftPurviewMetadataPolicyOperationRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftPurviewPolicyOperationRecord") {
		var out SecurityMicrosoftPurviewPolicyOperationRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftPurviewPolicyOperationRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftPurviewPrivacyAuditEvent") {
		var out SecurityMicrosoftPurviewPrivacyAuditEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftPurviewPrivacyAuditEvent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftStreamAuditRecord") {
		var out SecurityMicrosoftStreamAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftStreamAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftTeamsAdminAuditRecord") {
		var out SecurityMicrosoftTeamsAdminAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftTeamsAdminAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftTeamsAnalyticsAuditRecord") {
		var out SecurityMicrosoftTeamsAnalyticsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftTeamsAnalyticsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftTeamsAuditRecord") {
		var out SecurityMicrosoftTeamsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftTeamsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftTeamsDeviceAuditRecord") {
		var out SecurityMicrosoftTeamsDeviceAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftTeamsDeviceAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftTeamsRetentionLabelActionAuditRecord") {
		var out SecurityMicrosoftTeamsRetentionLabelActionAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftTeamsRetentionLabelActionAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftTeamsSensitivityLabelActionAuditRecord") {
		var out SecurityMicrosoftTeamsSensitivityLabelActionAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftTeamsSensitivityLabelActionAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.microsoftTeamsShiftsAuditRecord") {
		var out SecurityMicrosoftTeamsShiftsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMicrosoftTeamsShiftsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelExchangeItemAuditRecord") {
		var out SecurityMipAutoLabelExchangeItemAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelExchangeItemAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelItemAuditRecord") {
		var out SecurityMipAutoLabelItemAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelItemAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelPolicyAuditRecord") {
		var out SecurityMipAutoLabelPolicyAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelPolicyAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelProgressFeedbackAuditRecord") {
		var out SecurityMipAutoLabelProgressFeedbackAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelProgressFeedbackAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelSharePointItemAuditRecord") {
		var out SecurityMipAutoLabelSharePointItemAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelSharePointItemAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelSharePointPolicyLocationAuditRecord") {
		var out SecurityMipAutoLabelSharePointPolicyLocationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelSharePointPolicyLocationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelSimulationSharePointCompletionRecord") {
		var out SecurityMipAutoLabelSimulationSharePointCompletionRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelSimulationSharePointCompletionRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelSimulationSharePointProgressRecord") {
		var out SecurityMipAutoLabelSimulationSharePointProgressRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelSimulationSharePointProgressRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelSimulationStatisticsRecord") {
		var out SecurityMipAutoLabelSimulationStatisticsRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelSimulationStatisticsRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipAutoLabelSimulationStatusRecord") {
		var out SecurityMipAutoLabelSimulationStatusRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipAutoLabelSimulationStatusRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipExactDataMatchAuditRecord") {
		var out SecurityMipExactDataMatchAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipExactDataMatchAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipLabelAnalyticsAuditRecord") {
		var out SecurityMipLabelAnalyticsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipLabelAnalyticsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.mipLabelAuditRecord") {
		var out SecurityMipLabelAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMipLabelAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.msdeGeneralSettingsAuditRecord") {
		var out SecurityMsdeGeneralSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMsdeGeneralSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.msdeIndicatorsSettingsAuditRecord") {
		var out SecurityMsdeIndicatorsSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMsdeIndicatorsSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.msdeResponseActionsAuditRecord") {
		var out SecurityMsdeResponseActionsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMsdeResponseActionsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.msdeRolesSettingsAuditRecord") {
		var out SecurityMsdeRolesSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMsdeRolesSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.msticNationStateNotificationRecord") {
		var out SecurityMsticNationStateNotificationRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMsticNationStateNotificationRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.multiStageDispositionAuditRecord") {
		var out SecurityMultiStageDispositionAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMultiStageDispositionAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.myAnalyticsSettingsAuditRecord") {
		var out SecurityMyAnalyticsSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMyAnalyticsSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.officeNativeAuditRecord") {
		var out SecurityOfficeNativeAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityOfficeNativeAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.omePortalAuditRecord") {
		var out SecurityOmePortalAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityOmePortalAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.onPremisesFileShareScannerDlpAuditRecord") {
		var out SecurityOnPremisesFileShareScannerDlpAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityOnPremisesFileShareScannerDlpAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.onPremisesScannerDlpAuditRecord") {
		var out SecurityOnPremisesScannerDlpAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityOnPremisesScannerDlpAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.onPremisesSharePointScannerDlpAuditRecord") {
		var out SecurityOnPremisesSharePointScannerDlpAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityOnPremisesSharePointScannerDlpAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.oneDriveAuditRecord") {
		var out SecurityOneDriveAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityOneDriveAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.owaGetAccessTokenForResourceAuditRecord") {
		var out SecurityOwaGetAccessTokenForResourceAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityOwaGetAccessTokenForResourceAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.peopleAdminSettingsAuditRecord") {
		var out SecurityPeopleAdminSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPeopleAdminSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.physicalBadgingSignalAuditRecord") {
		var out SecurityPhysicalBadgingSignalAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPhysicalBadgingSignalAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.plannerCopyPlanAuditRecord") {
		var out SecurityPlannerCopyPlanAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPlannerCopyPlanAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.plannerPlanAuditRecord") {
		var out SecurityPlannerPlanAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPlannerPlanAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.plannerPlanListAuditRecord") {
		var out SecurityPlannerPlanListAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPlannerPlanListAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.plannerRosterAuditRecord") {
		var out SecurityPlannerRosterAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPlannerRosterAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.plannerRosterSensitivityLabelAuditRecord") {
		var out SecurityPlannerRosterSensitivityLabelAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPlannerRosterSensitivityLabelAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.plannerTaskAuditRecord") {
		var out SecurityPlannerTaskAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPlannerTaskAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.plannerTaskListAuditRecord") {
		var out SecurityPlannerTaskListAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPlannerTaskListAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.plannerTenantSettingsAuditRecord") {
		var out SecurityPlannerTenantSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPlannerTenantSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerAppsAuditAppRecord") {
		var out SecurityPowerAppsAuditAppRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerAppsAuditAppRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerAppsAuditPlanRecord") {
		var out SecurityPowerAppsAuditPlanRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerAppsAuditPlanRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerAppsAuditResourceRecord") {
		var out SecurityPowerAppsAuditResourceRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerAppsAuditResourceRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerBiAuditRecord") {
		var out SecurityPowerBiAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerBiAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerBiDlpAuditRecord") {
		var out SecurityPowerBiDlpAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerBiDlpAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerPagesSiteAuditRecord") {
		var out SecurityPowerPagesSiteAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerPagesSiteAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerPlatformAdminDlpAuditRecord") {
		var out SecurityPowerPlatformAdminDlpAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerPlatformAdminDlpAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerPlatformAdminEnvironmentAuditRecord") {
		var out SecurityPowerPlatformAdminEnvironmentAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerPlatformAdminEnvironmentAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerPlatformAdministratorActivityRecord") {
		var out SecurityPowerPlatformAdministratorActivityRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerPlatformAdministratorActivityRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerPlatformLockboxResourceAccessRequestAuditRecord") {
		var out SecurityPowerPlatformLockboxResourceAccessRequestAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerPlatformLockboxResourceAccessRequestAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerPlatformLockboxResourceCommandAuditRecord") {
		var out SecurityPowerPlatformLockboxResourceCommandAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerPlatformLockboxResourceCommandAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.powerPlatformServiceActivityAuditRecord") {
		var out SecurityPowerPlatformServiceActivityAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPowerPlatformServiceActivityAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.privacyDataMatchAuditRecord") {
		var out SecurityPrivacyDataMatchAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPrivacyDataMatchAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.privacyDataMinimizationRecord") {
		var out SecurityPrivacyDataMinimizationRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPrivacyDataMinimizationRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.privacyDigestEmailRecord") {
		var out SecurityPrivacyDigestEmailRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPrivacyDigestEmailRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.privacyOpenAccessAuditRecord") {
		var out SecurityPrivacyOpenAccessAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPrivacyOpenAccessAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.privacyPortalAuditRecord") {
		var out SecurityPrivacyPortalAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPrivacyPortalAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.privacyRemediationActionRecord") {
		var out SecurityPrivacyRemediationActionRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPrivacyRemediationActionRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.privacyRemediationRecord") {
		var out SecurityPrivacyRemediationRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPrivacyRemediationRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.privacyTenantAuditHistoryRecord") {
		var out SecurityPrivacyTenantAuditHistoryRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPrivacyTenantAuditHistoryRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.projectAuditRecord") {
		var out SecurityProjectAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProjectAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.projectForTheWebAssignedToMeSettingsAuditRecord") {
		var out SecurityProjectForTheWebAssignedToMeSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProjectForTheWebAssignedToMeSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.projectForTheWebProjectAuditRecord") {
		var out SecurityProjectForTheWebProjectAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProjectForTheWebProjectAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.projectForTheWebProjectSettingsAuditRecord") {
		var out SecurityProjectForTheWebProjectSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProjectForTheWebProjectSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.projectForTheWebRoadmapAuditRecord") {
		var out SecurityProjectForTheWebRoadmapAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProjectForTheWebRoadmapAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.projectForTheWebRoadmapItemAuditRecord") {
		var out SecurityProjectForTheWebRoadmapItemAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProjectForTheWebRoadmapItemAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.projectForTheWebRoadmapSettingsAuditRecord") {
		var out SecurityProjectForTheWebRoadmapSettingsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProjectForTheWebRoadmapSettingsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.projectForTheWebTaskAuditRecord") {
		var out SecurityProjectForTheWebTaskAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProjectForTheWebTaskAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.publicFolderAuditRecord") {
		var out SecurityPublicFolderAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPublicFolderAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.purviewInsiderRiskAlertsRecord") {
		var out SecurityPurviewInsiderRiskAlertsRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPurviewInsiderRiskAlertsRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.purviewInsiderRiskCasesRecord") {
		var out SecurityPurviewInsiderRiskCasesRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityPurviewInsiderRiskCasesRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.quarantineAuditRecord") {
		var out SecurityQuarantineAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityQuarantineAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.recordsManagementAuditRecord") {
		var out SecurityRecordsManagementAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRecordsManagementAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.retentionPolicyAuditRecord") {
		var out SecurityRetentionPolicyAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRetentionPolicyAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.scoreEvidence") {
		var out SecurityScoreEvidence
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityScoreEvidence: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.scorePlatformGenericAuditRecord") {
		var out SecurityScorePlatformGenericAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityScorePlatformGenericAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.scriptRunAuditRecord") {
		var out SecurityScriptRunAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityScriptRunAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.searchAuditRecord") {
		var out SecuritySearchAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySearchAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.securityComplianceAlertRecord") {
		var out SecuritySecurityComplianceAlertRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySecurityComplianceAlertRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.securityComplianceCenterEOPCmdletAuditRecord") {
		var out SecuritySecurityComplianceCenterEOPCmdletAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySecurityComplianceCenterEOPCmdletAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.securityComplianceInsightsAuditRecord") {
		var out SecuritySecurityComplianceInsightsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySecurityComplianceInsightsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.securityComplianceRBACAuditRecord") {
		var out SecuritySecurityComplianceRBACAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySecurityComplianceRBACAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.securityComplianceUserChangeAuditRecord") {
		var out SecuritySecurityComplianceUserChangeAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySecurityComplianceUserChangeAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sharePointAppPermissionOperationAuditRecord") {
		var out SecuritySharePointAppPermissionOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySharePointAppPermissionOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sharePointAuditRecord") {
		var out SecuritySharePointAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySharePointAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sharePointCommentOperationAuditRecord") {
		var out SecuritySharePointCommentOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySharePointCommentOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sharePointContentTypeOperationAuditRecord") {
		var out SecuritySharePointContentTypeOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySharePointContentTypeOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sharePointESignatureAuditRecord") {
		var out SecuritySharePointESignatureAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySharePointESignatureAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sharePointFieldOperationAuditRecord") {
		var out SecuritySharePointFieldOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySharePointFieldOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sharePointFileOperationAuditRecord") {
		var out SecuritySharePointFileOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySharePointFileOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sharePointListOperationAuditRecord") {
		var out SecuritySharePointListOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySharePointListOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.sharePointSharingOperationAuditRecord") {
		var out SecuritySharePointSharingOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySharePointSharingOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.skypeForBusinessBaseAuditRecord") {
		var out SecuritySkypeForBusinessBaseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySkypeForBusinessBaseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.skypeForBusinessCmdletsAuditRecord") {
		var out SecuritySkypeForBusinessCmdletsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySkypeForBusinessCmdletsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.skypeForBusinessPSTNUsageAuditRecord") {
		var out SecuritySkypeForBusinessPSTNUsageAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySkypeForBusinessPSTNUsageAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.skypeForBusinessUsersBlockedAuditRecord") {
		var out SecuritySkypeForBusinessUsersBlockedAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySkypeForBusinessUsersBlockedAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.smsCreatePhoneNumberRecord") {
		var out SecuritySmsCreatePhoneNumberRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySmsCreatePhoneNumberRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.smsDeletePhoneNumberRecord") {
		var out SecuritySmsDeletePhoneNumberRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySmsDeletePhoneNumberRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.supervisoryReviewDayXInsightsAuditRecord") {
		var out SecuritySupervisoryReviewDayXInsightsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySupervisoryReviewDayXInsightsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.syntheticProbeAuditRecord") {
		var out SecuritySyntheticProbeAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySyntheticProbeAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.teamsEasyApprovalsAuditRecord") {
		var out SecurityTeamsEasyApprovalsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTeamsEasyApprovalsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.teamsHealthcareAuditRecord") {
		var out SecurityTeamsHealthcareAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTeamsHealthcareAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.teamsUpdatesAuditRecord") {
		var out SecurityTeamsUpdatesAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTeamsUpdatesAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.tenantAllowBlockListAuditRecord") {
		var out SecurityTenantAllowBlockListAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTenantAllowBlockListAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.threatFinderAuditRecord") {
		var out SecurityThreatFinderAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityThreatFinderAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.threatIntelligenceAtpContentData") {
		var out SecurityThreatIntelligenceAtpContentData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityThreatIntelligenceAtpContentData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.threatIntelligenceMailData") {
		var out SecurityThreatIntelligenceMailData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityThreatIntelligenceMailData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.threatIntelligenceUrlClickData") {
		var out SecurityThreatIntelligenceUrlClickData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityThreatIntelligenceUrlClickData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.todoAuditRecord") {
		var out SecurityTodoAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityTodoAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.uamOperationAuditRecord") {
		var out SecurityUamOperationAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUamOperationAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.unifiedGroupAuditRecord") {
		var out SecurityUnifiedGroupAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUnifiedGroupAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.unifiedSimulationMatchedItemAuditRecord") {
		var out SecurityUnifiedSimulationMatchedItemAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUnifiedSimulationMatchedItemAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.unifiedSimulationSummaryAuditRecord") {
		var out SecurityUnifiedSimulationSummaryAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUnifiedSimulationSummaryAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.uploadCertificateRecord") {
		var out SecurityUploadCertificateRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUploadCertificateRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.urbacAssignmentAuditRecord") {
		var out SecurityUrbacAssignmentAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUrbacAssignmentAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.urbacEnableStateAuditRecord") {
		var out SecurityUrbacEnableStateAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUrbacEnableStateAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.urbacRoleAuditRecord") {
		var out SecurityUrbacRoleAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUrbacRoleAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.userTrainingAuditRecord") {
		var out SecurityUserTrainingAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUserTrainingAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vfamBasePolicyAuditRecord") {
		var out SecurityVfamBasePolicyAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVfamBasePolicyAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vfamCreatePolicyAuditRecord") {
		var out SecurityVfamCreatePolicyAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVfamCreatePolicyAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vfamDeletePolicyAuditRecord") {
		var out SecurityVfamDeletePolicyAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVfamDeletePolicyAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vfamUpdatePolicyAuditRecord") {
		var out SecurityVfamUpdatePolicyAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVfamUpdatePolicyAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vivaGoalsAuditRecord") {
		var out SecurityVivaGoalsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVivaGoalsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vivaLearningAdminAuditRecord") {
		var out SecurityVivaLearningAdminAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVivaLearningAdminAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vivaLearningAuditRecord") {
		var out SecurityVivaLearningAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVivaLearningAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vivaPulseAdminAuditRecord") {
		var out SecurityVivaPulseAdminAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVivaPulseAdminAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vivaPulseOrganizerAuditRecord") {
		var out SecurityVivaPulseOrganizerAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVivaPulseOrganizerAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vivaPulseReportAuditRecord") {
		var out SecurityVivaPulseReportAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVivaPulseReportAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.vivaPulseResponseAuditRecord") {
		var out SecurityVivaPulseResponseAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityVivaPulseResponseAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.wdatpAlertsAuditRecord") {
		var out SecurityWdatpAlertsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityWdatpAlertsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.windows365CustomerLockboxAuditRecord") {
		var out SecurityWindows365CustomerLockboxAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityWindows365CustomerLockboxAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.workplaceAnalyticsAuditRecord") {
		var out SecurityWorkplaceAnalyticsAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityWorkplaceAnalyticsAuditRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.yammerAuditRecord") {
		var out SecurityYammerAuditRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityYammerAuditRecord: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityAuditDataImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityAuditDataImpl: %+v", err)
	}

	return RawSecurityAuditDataImpl{
		securityAuditData: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
