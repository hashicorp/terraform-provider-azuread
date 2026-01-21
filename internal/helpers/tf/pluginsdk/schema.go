// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package pluginsdk

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Schema = schema.Schema

const (
	TypeInvalid = schema.TypeInvalid
	TypeBool    = schema.TypeBool
	TypeInt     = schema.TypeInt
	TypeFloat   = schema.TypeFloat
	TypeString  = schema.TypeString
	TypeList    = schema.TypeList
	TypeMap     = schema.TypeMap
	TypeSet     = schema.TypeSet
)

const (
	SchemaConfigModeAuto  = schema.SchemaConfigModeAuto
	SchemaConfigModeAttr  = schema.SchemaConfigModeAttr
	SchemaConfigModeBlock = schema.SchemaConfigModeBlock
)

type SchemaDiffSuppressFunc = schema.SchemaDiffSuppressFunc
type SchemaDefaultFunc = schema.SchemaDefaultFunc

// EnvDefaultFunc is a helper function that returns the value of the
// given environment variable, if one exists, or the default value
// otherwise.
func EnvDefaultFunc(k string, dv interface{}) SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(k); v != "" {
			return v, nil
		}

		return dv, nil
	}
}

// MultiEnvDefaultFunc is a helper function that returns the value of the first
// environment variable in the given list that returns a non-empty value. If
// none of the environment variables return a value, the default value is
// returned.
func MultiEnvDefaultFunc(ks []string, dv interface{}) SchemaDefaultFunc {
	return func() (interface{}, error) {
		for _, k := range ks {
			if v := os.Getenv(k); v != "" {
				return v, nil
			}
		}
		return dv, nil
	}
}
