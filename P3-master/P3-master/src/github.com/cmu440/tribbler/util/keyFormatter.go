package util

import (
	"fmt"
	"math"
	"math/rand"
)

// Get formatted key for a user
// Example: bob => bob:userID
func FormatUserKey(userID string) string {
	return fmt.Sprintf("%s:usrid", userID)
}

// Get formatted key for a user's subscription list
// Example: bob => bob:sublist
func FormatSubListKey(userID string) string {
	return fmt.Sprintf("%s:sublist", userID)
}

// Get formatted key when posting a new tribble
// Example: bob makes post at time <timeNano> => bob:post_<timeNano>_<rand64>
// timeNano (in hex) should already be unique, but <rand64> breaks any ties
func FormatPostKey(userID string, timeNano int64) string {
	return fmt.Sprintf("%s:post_%x_%x", userID, timeNano,
		math.Float64bits(rand.Float64()))
}

// Get formatted key for a user's list of tribble keys
// Example: bob => bob:triblist
func FormatTribListKey(userID string) string {
	return fmt.Sprintf("%s:triblist", userID)
}
