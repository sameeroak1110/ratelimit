/* *****************************************************************************
Copyright (c) 2022, sameeroak1110 (sameeroak1110@gmail.com)
All rights reserved.
BSD 3-Clause License.

Package     : github.com/sameeroak1110/ratelimit
Filename    : github.com/sameeroak1110/ratelimit/process.go
File-type   : golang source code file

Compiler/Runtime: go version go1.17 linux/amd64

Version History
Version     : 1.0
Author      : sameer oak (sameeroak1110@gmail.com)
Description :
- Rate limiter process.
***************************************************************************** */
package ratelimit

import (
	"sync"
	"time"

	"github.com/sameeroak1110/logger"
)


/* *****************************************************************************
Description :
- Identifies if the IP source is in the permissible limits (max number of API hits in a window) imposed by ratelimiter.

Arguments   :
1> srcIP string: Source IP wherefrom API has been invoked.

Return value:
1> bool: true if source IP is in the permissible limits (max number of API hits in a window) imposed by ratelimiter.
false otherwise.

Additional note: NA
***************************************************************************** */
func GetStatus(srcIP string) bool {
	nowTS := time.Now().In(time.UTC)

	pSrcRec, isOK := getAPISourceRec(srcIP)
	if !isOK {  // the sourceIP appears for the 1st time. need to add it the map
		cacheLock.Lock()
		defer cacheLock.Unlock()

		prec := &apiSource {
			key: srcIP,
			remainingAttempts: MAX_API_PER_RATELIMIT_WINDOW - 1,  // MAX_API_PER_RATELIMIT_WINDOW - 1 since this very req is one amongst the MAX_API_PER_RATELIMIT_WINDOW
			lastAPITS: nowTS,
		}
		prec.precLock = &sync.Mutex{}
		apiSourceMap[srcIP] = prec
		return true
	}

	// here, means the record is available in the cache and it's in the locked state.
	defer func() {
		pSrcRec.precLock.Unlock()  // record is unlocked.
	}()

	diffTS := nowTS.Sub(pSrcRec.lastAPITS)
	if int(diffTS.Seconds()) <= RATELIMIT_WINDOW {
		logger.Log(pkgname, logger.DEBUG, "1> diff seconds: %d,  remaining attempts: %d", int(diffTS.Seconds()), pSrcRec.remainingAttempts)
		if pSrcRec.remainingAttempts <= 0 {
			pSrcRec.lastAPITS = nowTS  // penalty, as source-IP hit beyond permissible limit of no. of APIs per ratelimit-window. it's likely the API has been hit with malicious intent.
			return false
		}
		pSrcRec.remainingAttempts = pSrcRec.remainingAttempts - 1
		return true
	} else {
		pSrcRec.remainingAttempts = MAX_API_PER_RATELIMIT_WINDOW - 1  // MAX_API_PER_RATELIMIT_WINDOW - 1 since this very req is one amongst the MAX_API_PER_RATELIMIT_WINDOW
		pSrcRec.lastAPITS = nowTS
		logger.Log(pkgname, logger.DEBUG, "2> diff seconds: %d,  remaining attempts: %d", int(diffTS.Seconds()), pSrcRec.remainingAttempts)
		return true
	}

	return true
}
