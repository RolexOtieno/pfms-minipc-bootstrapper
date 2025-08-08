/*
	Validates if the MiniPC making the request is on the allowed list.

Maintains a list of authorized MiniPC deviceIds.
Returns true or false based on the ID.
*/
package main

// IsAuthorizedDevice checks if the device is in the allowlist
func IsAuthorizedDevice(deviceID string) bool {
	authorizedDevices := map[string]bool{
		"MINIPC_123456": true,
		"MINIPC_654321": true,
		"MINIPC_999999": true,
	}
	return authorizedDevices[deviceID]
}
