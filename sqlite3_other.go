// Copyright (C) 2014 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
// +build !windows

package sqlite3

/*
#cgo CFLAGS: -I. -DHAVE_USLEEP=1 -DTHREADSAFE=1
#cgo linux LDFLAGS: -ldl
#cgo LDFLAGS: -lpthread
*/
import "C"
