// Copyright 2018 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Code generated by "stringer -type=Mutation"; DO NOT EDIT

package reflow

import "fmt"

const _Mutation_name = "IncrDecrCached"

var _Mutation_index = [...]uint8{0, 4, 8, 14}

func (i Mutation) String() string {
	if i < 0 || i >= Mutation(len(_Mutation_index)-1) {
		return fmt.Sprintf("Mutation(%d)", i)
	}
	return _Mutation_name[_Mutation_index[i]:_Mutation_index[i+1]]
}