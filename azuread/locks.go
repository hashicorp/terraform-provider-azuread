package azuread

// handle the case of using the same name for different kinds of resources
func azureADLockByName(name string, resourceType string) {
	armMutexKV.Lock(resourceType + "." + name)
}

func azureADUnlockByName(name string, resourceType string) {
	armMutexKV.Unlock(resourceType + "." + name)
}
