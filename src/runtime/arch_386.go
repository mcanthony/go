// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	thechar        = '8'
	_BigEndian     = 0
	_CacheLineSize = 64
	_PhysPageSize  = goos_nacl*65536 + (1-goos_nacl)*4096 // 4k normally; 64k on NaCl
	_PCQuantum     = 1
	_Int64Align    = 4
	hugePageSize   = 1 << 21
	minFrameSize   = 0
)

type uintreg uint32
type intptr int32 // TODO(rsc): remove
