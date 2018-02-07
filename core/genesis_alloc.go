// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xf8\xf2\xe1\x94\x0f*\xf2E|\x81\xb3o1\xf5ZQ#\xf8\xbe\xfb\u061d\xb3\xab\x8bg\xf2S(\xadJL\f4\x00\x00\xe1\x94}Wy\u966eRt\x8e;\xb6\x1d\xa7\xc6\xea\xb4Q|XG\x8b\x81\xee\xe7\xf2\u061c\xdf\x0fA\x00\x00\u2505i|B8>\xef\x01\x8e%\xf1\xf4N%\xb7(\xeaF\x06=\x8c\x017\xd6\xf9z\a\xde\xe4$\x9c\x00\x00\u1519'\xdf\xf4t\xf9\xfe\xeb\xd18;|\v'\xb8X\x10\xceI\xf3\x8bg\xf2S(\xadJL\f4\x00\x00\xe2\x94\xc7W\xd59\xb0\xc0\xad\x12x\xddN\x15\xbf\x14-k\xa0\x93x\xc1\x8c\x01\x03\xdd\xcf\xe5\xb19\xbe\x1e\x82\x00\x00\xe2\x94\xd7:%\x82\x13\x8a\x9dcN\xb8\xc0j\x1d\xe4\xfc\x0e\xc1\x11?\x88\x8c\x04\x0fw?\x96\xc4\xe6\xf8z\b\x00\x00\xe2\x94\xe0\x01JBh\xad={\x13\x15Q\xbeG\xbcn\xd7*i7\xe4\x8c\x02\x89\xaa\x87\xbe;\x10[LE\x00\x00"

const testnetAllocData = "\xe3\xe2\x94A\x10\xbd\x1f\xf0\xb7?\xa1,%\x9a\xcf9\xc9P'\u007f&g\x87\x8c O\xce^>%\x02a\x10\x00\x00\x00"
const rinkebyAllocData = "\xe3\xe2\x94A\x10\xbd\x1f\xf0\xb7?\xa1,%\x9a\xcf9\xc9P'\u007f&g\x87\x8c O\xce^>%\x02a\x10\x00\x00\x00"
