/* *****************************************************************************
Copyright (c) 2022, sameeroak1110 (sameeroak1110@gmail.com)
All rights reserved.
BSD 3-Clause License.

Package     : github.com/sameeroak1110/ratelimit
Filename    : github.com/sameeroak1110/ratelimit/ratelimit.go
File-type   : golang source code file

Compiler/Runtime: go version go1.17 linux/amd64

Version History
Version     : 1.0
Author      : sameer oak (sameeroak1110@gmail.com)
Description :
- Initializes ratelimit package.
***************************************************************************** */
package ratelimit

func Init() {
	apiSourceMap = make(map[string]*apiSource)
}
