// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package synchronization

import (
	"net/http"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

const servicePrincipalResourceName = "azuread_service_principal"

func synchronizationRetryFunc() client.RequestRetryFunc {
	return func(resp *http.Response, o *odata.OData) (bool, error) {
		return response.WasConflict(resp) || response.WasStatusCode(resp, http.StatusForbidden), nil
	}
}

func emptySynchronizationSecretKeyStringValuePair(in []interface{}) *[]stable.SynchronizationSecretKeyStringValuePair {
	result := make([]stable.SynchronizationSecretKeyStringValuePair, 0)

	for _, raw := range in {
		if raw == nil {
			continue
		}
		item := raw.(map[string]interface{})

		result = append(result, stable.SynchronizationSecretKeyStringValuePair{
			Key:   pointer.To(stable.SynchronizationSecret(item["key"].(string))),
			Value: nullable.Value(""),
		})
	}

	return &result
}

func expandSynchronizationSecretKeyStringValuePair(in []interface{}) *[]stable.SynchronizationSecretKeyStringValuePair {
	result := make([]stable.SynchronizationSecretKeyStringValuePair, 0)

	for _, raw := range in {
		if raw == nil {
			continue
		}
		item := raw.(map[string]interface{})

		result = append(result, stable.SynchronizationSecretKeyStringValuePair{
			Key:   pointer.To(stable.SynchronizationSecret(item["key"].(string))),
			Value: nullable.Value(item["value"].(string)),
		})
	}

	return &result
}

func expandSynchronizationJobApplicationParameters(in []interface{}) *[]stable.SynchronizationJobApplicationParameters {
	result := make([]stable.SynchronizationJobApplicationParameters, 0)

	for _, raw := range in {
		if raw == nil {
			continue
		}
		item := raw.(map[string]interface{})

		result = append(result, stable.SynchronizationJobApplicationParameters{
			Subjects: expandSynchronizationJobSubject(item["subject"].([]interface{})),
			RuleId:   nullable.Value(item["rule_id"].(string)),
		})
	}

	return &result
}

func expandSynchronizationJobSubject(in []interface{}) *[]stable.SynchronizationJobSubject {
	result := make([]stable.SynchronizationJobSubject, 0)
	for _, raw := range in {
		if raw == nil {
			continue
		}
		item := raw.(map[string]interface{})

		result = append(result, stable.SynchronizationJobSubject{
			ObjectId:       nullable.Value(item["object_id"].(string)),
			ObjectTypeName: nullable.Value(item["object_type_name"].(string)),
		})
	}

	return &result
}

func flattenSynchronizationSchedule(in *stable.SynchronizationSchedule) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"expiration": in.Expiration.GetOrZero(),
		"interval":   pointer.From(in.Interval),
		"state":      pointer.From(in.State),
	}}
}

func flattenSynchronizationSecretKeyStringValuePair(in *[]stable.SynchronizationSecretKeyStringValuePair, current []interface{}) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	credentials := make([]interface{}, 0)
	for _, item := range *in {
		value := item.Value.GetOrZero()
		if value == "*" && current != nil {
			// Use value from state if API returns * indicating sensitive data
			for _, raw := range current {
				if raw == nil {
					continue
				}
				currentItem := raw.(map[string]interface{})
				if currentItem["key"].(string) == string(pointer.From(item.Key)) {
					value = currentItem["value"].(string)
				}
			}
		}
		credential := map[string]interface{}{
			"key":   pointer.From(item.Key),
			"value": value,
		}
		credentials = append(credentials, credential)
	}

	return credentials
}

func allCredentialsRemoved(in []stable.SynchronizationSecretKeyStringValuePair, current []stable.SynchronizationSecretKeyStringValuePair) bool {
	for _, item := range in {
		for _, itemCurrent := range current {
			if pointer.From(item.Key) == pointer.From(itemCurrent.Key) {
				return false
			}
		}
	}

	return true
}
