package structure

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
)

func SuppressJsonDiff(k, old, new string, d *schema.ResourceData) bool {
	oldMap, err := ExpandJsonFromString(old)
	if err != nil {
		return false
	}

	fmt.Printf("[MMMM]: %s", oldMap)

	newMap, err := ExpandJsonFromString(new)
	if err != nil {
		return false
	}
	fmt.Printf("[NNNN]: %s", newMap)

	return reflect.DeepEqual(oldMap, newMap)
}
