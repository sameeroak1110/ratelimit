/* *****************************************************************************
Copyright (c) 2022, sameeroak1110 (sameeroak1110@gmail.com)
All rights reserved.
BSD 3-Clause License.

Package     : github.com/sameeroak1110/ratelimit
Filename    : github.com/sameeroak1110/ratelimit/data.go
File-type   : golang source code file

Compiler/Runtime: go version go1.17 linux/amd64

Version History
Version     : 1.0
Author      : sameer oak (sameeroak1110@gmail.com)
Description :
- Local and exported data.
***************************************************************************** */
package ratelimit

import (
	"sync"
)

const pkgname string = "ratelimit"

//const RATELIMIT_WINDOW             int = 5    // 3 seconds.
//const MAX_API_PER_RATELIMIT_WINDOW int = 3    // max number of APIs can be invoked from IP in RATELIMIT_WINDOW.
const RATELIMIT_WINDOW             int = 60   // 60 seconds.
const MAX_API_PER_RATELIMIT_WINDOW int = 100  // max number of APIs can be invoked from IP in RATELIMIT_WINDOW.

var apiSourceMap map[string]*apiSource
var cacheLock sync.RWMutex       // cache store-lock. rd store lock/unlock and wr store lock/unlock operations.
