package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

func assignManager(ctx context.Context, client *msgraph.UsersClient, directoryObjectsClient *msgraph.DirectoryObjectsClient, userId, managerId string) error {
	if managerId != "" {
		managerObject, _, err := directoryObjectsClient.Get(ctx, managerId, odata.Query{})
		if err != nil {
			return err
		}
		if managerObject == nil {
			return errors.New("managerObject was nil")
		}
		// TODO: remove this workaround for https://github.com/hashicorp/terraform-provider-azuread/issues/588
		//if managerObject.ODataId == nil {
		//	return tf.ErrorDiagF(errors.New("ODataId was nil"), "Could not retrieve manager principal object %q", managerId)
		//}
		managerObject.ODataId = (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
			client.BaseClient.Endpoint, client.BaseClient.TenantId, managerId)))

		manager := msgraph.User{
			DirectoryObject: *managerObject,
		}

		if _, err = client.AssignManager(ctx, userId, manager); err != nil {
			return err
		}
	} else {
		if _, err := client.DeleteManager(ctx, userId); err != nil {
			return err
		}
	}

	return nil
}

type QueryOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	ResultCount      bool
}

func read(ctx context.Context, client *msgraph.UsersClient, searchValue string, qOptions QueryOptions,
	queryPropertyName string, schemaPropertyName string, propertyLogSynonym string) (*msgraph.User, diag.Diagnostics) {
	query := odata.Query{
		Filter: fmt.Sprintf("%v eq '%s'",
			queryPropertyName,
			utils.EscapeSingleQuote(searchValue)),
	}
	if qOptions.ConsistencyLevel != nil {
		query.ConsistencyLevel = *qOptions.ConsistencyLevel
	}
	if qOptions.ResultCount {
		query.Count = qOptions.ResultCount
	}
	users, _, err := client.List(ctx, query)
	if err != nil {
		return nil, tf.ErrorDiagF(err, "Finding user with %v: %q", propertyLogSynonym, searchValue)
	}
	if users == nil {
		return nil, tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
	}
	count := len(*users)
	if count > 1 {
		return nil, tf.ErrorDiagPathF(nil, schemaPropertyName, "More than one user found with %v: %q", propertyLogSynonym, searchValue)
	} else if count == 0 {
		return nil, tf.ErrorDiagPathF(nil, schemaPropertyName, "User with %v %q was not found", propertyLogSynonym, searchValue)
	}

	return &(*users)[0], nil
}

func listRead(ctx context.Context, client *msgraph.UsersClient, searchElements []interface{}, ignoreMissing bool, qOptions QueryOptions,
	queryPropertyName string, schemaPropertyName string, propertyLogSynonym string) ([]msgraph.User, diag.Diagnostics) {
	users := []msgraph.User{}
	for _, v := range searchElements {
		query := odata.Query{
			Filter: fmt.Sprintf("%v eq '%s'",
				queryPropertyName,
				utils.EscapeSingleQuote(v.(string))),
		}
		if qOptions.ConsistencyLevel != nil {
			query.ConsistencyLevel = *qOptions.ConsistencyLevel
		}
		if qOptions.ResultCount {
			query.Count = qOptions.ResultCount
		}
		result, _, err := client.List(ctx, query)
		if err != nil {
			return users, tf.ErrorDiagF(err, "Finding user with %v: %q", propertyLogSynonym, v)
		}
		if result == nil {
			return users, tf.ErrorDiagF(errors.New("API returned nil result"), "Bad API Response")
		}
		count := len(*result)
		if count > 1 {
			return users, tf.ErrorDiagPathF(nil, schemaPropertyName, "More than one user found with %v: %q", propertyLogSynonym, v)
		} else if count == 0 {
			if ignoreMissing {
				continue
			}
			return users, tf.ErrorDiagPathF(nil, schemaPropertyName, "User with %v %q was not found", propertyLogSynonym, v)
		}
		users = append(users, (*result)[0])
	}
	return users, nil
}
