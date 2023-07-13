// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"time"

	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func emptySynchronizationSecretKeyStringValuePair(in []interface{}) *[]msgraph.SynchronizationSecretKeyStringValuePair {
	result := make([]msgraph.SynchronizationSecretKeyStringValuePair, 0)

	for _, raw := range in {
		if raw == nil {
			continue
		}
		item := raw.(map[string]interface{})

		result = append(result, msgraph.SynchronizationSecretKeyStringValuePair{
			Key:   utils.String(item["key"].(string)),
			Value: utils.String(""),
		})
	}

	return &result
}

func expandSynchronizationSecretKeyStringValuePair(in []interface{}) *[]msgraph.SynchronizationSecretKeyStringValuePair {
	result := make([]msgraph.SynchronizationSecretKeyStringValuePair, 0)

	for _, raw := range in {
		if raw == nil {
			continue
		}
		item := raw.(map[string]interface{})

		result = append(result, msgraph.SynchronizationSecretKeyStringValuePair{
			Key:   utils.String(item["key"].(string)),
			Value: utils.String(item["value"].(string)),
		})
	}

	return &result
}

func flattenSynchronizationSchedule(in *msgraph.SynchronizationSchedule) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	expiration := ""
	if v := in.Expiration; v != nil {
		expiration = v.Format(time.RFC3339)
	}
	return []map[string]interface{}{{
		"expiration": expiration,
		"interval":   in.Interval,
		"state":      in.State,
	}}
}

func flattenSynchronizationSecretKeyStringValuePair(in *[]msgraph.SynchronizationSecretKeyStringValuePair, current []interface{}) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	credentials := make([]interface{}, 0)
	for _, item := range *in {
		value := item.Value
		if *value == "*" && current != nil {
			// Use value from state if API returns * indicating sensitive data
			for _, raw := range current {
				if raw == nil {
					continue
				}
				currentItem := raw.(map[string]interface{})
				if currentItem["key"].(string) == *item.Key {
					value = utils.String(currentItem["value"].(string))
				}
			}
		}
		credential := map[string]interface{}{
			"key":   item.Key,
			"value": value,
		}
		credentials = append(credentials, credential)
	}

	return credentials
}

func allCredentialsRemoved(in []msgraph.SynchronizationSecretKeyStringValuePair, current []msgraph.SynchronizationSecretKeyStringValuePair) bool {
	for _, item := range in {
		for _, itemCurrent := range current {
			if *item.Key == *itemCurrent.Key {
				return false
			}
		}
	}
	return true
}
