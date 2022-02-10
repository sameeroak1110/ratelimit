/* *****************************************************************************
Copyright (c) 2022, sameeroak1110 (sameeroak1110@gmail.com)
All rights reserved.
BSD 3-Clause License.

Package     : github.com/sameeroak1110/ratelimit
Filename    : github.com/sameeroak1110/ratelimit/types.go
File-type   : golang source code file

Compiler/Runtime: go version go1.17 linux/amd64

Version History
Version     : 1.0
Author      : sameer oak (sameeroak1110@gmail.com)
Description :
- Local and exported user defined data-types.
***************************************************************************** */
package ratelimit

import (
	"sync"
	"time"
)

type apiSource struct {
	key               string       // IP address is the key.
	precLock          *sync.Mutex  // Record lock.
	remainingAttempts int          // Current count of APIs hit in the valid RATELIMIT_WINDOW. Max allowed number of APIs is MAX_API_PER_RATELIMIT_WINDOW.
	lastAPITS         time.Time    // Time-stamp of the last API.
}
