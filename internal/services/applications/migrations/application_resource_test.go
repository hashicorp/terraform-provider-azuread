package migrations

import (
	"context"
	"reflect"
	"testing"
)

func TestResourceApplicationInstanceStateUpgradeV0(t *testing.T) {
	cases := []struct {
		TestName        string
		StateV0         map[string]interface{}
		StateV1Expected map[string]interface{}
	}{
		{
			TestName: "Empty group_membership_claims",
			StateV0: map[string]interface{}{
				"group_membership_claims": "",
				"public_client":           false,
			},
			StateV1Expected: map[string]interface{}{
				"group_membership_claims":        []string{},
				"fallback_public_client_enabled": false,
			},
		},
		{
			TestName: "Single group_membership_claims value",
			StateV0: map[string]interface{}{
				"group_membership_claims": "All",
				"public_client":           false,
			},
			StateV1Expected: map[string]interface{}{
				"group_membership_claims":        []string{"All"},
				"fallback_public_client_enabled": false,
			},
		},
		{
			TestName: "Multiple group_membership_claims values",
			StateV0: map[string]interface{}{
				"group_membership_claims": "ApplicationGroup,DirectoryRole,SecurityGroup",
				"public_client":           false,
			},
			StateV1Expected: map[string]interface{}{
				"group_membership_claims":        []string{"ApplicationGroup", "DirectoryRole", "SecurityGroup"},
				"fallback_public_client_enabled": false,
			},
		},
		{
			TestName: "public_client",
			StateV0: map[string]interface{}{
				"group_membership_claims": "",
				"public_client":           true,
			},
			StateV1Expected: map[string]interface{}{
				"group_membership_claims":        []string{},
				"fallback_public_client_enabled": true,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			newState, err := ResourceApplicationInstanceStateUpgradeV0(context.TODO(), tc.StateV0, nil)
			if err != nil {
				t.Fatalf("migrating state from v0 to v1: %+v", err)
			}

			if !reflect.DeepEqual(newState, tc.StateV1Expected) {
				t.Fatalf("migration failed. expected:\n\n%#v\n\nactual:\n\n%#v", tc.StateV1Expected, newState)
			}
		})
	}

}
