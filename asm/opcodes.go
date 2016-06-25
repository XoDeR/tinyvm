// Copyright 2016 Jeffrey Wilcke
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

package asm

type Op byte

const (
	Mov Op = iota
	Push
	Pop

	Jmp
	Jmpi
	Jmpn

	Gt
	Gteq
	Lt
	Lteq
	Eq
	Nq

	Call
	Ret

	Add
	Sub

	Dbg

	Stop
	Nop
)

var OpString = map[string]Op{
	"mov":  Mov,
	"push": Push,
	"pop":  Pop,
	"jmp":  Jmp,
	"jmpi": Jmpi,
	"jmpn": Jmpn,
	"gt":   Gt,
	"gteq": Gteq,
	"lt":   Lt,
	"lteq": Lteq,
	"eq":   Eq,
	"nq":   Nq,
	"add":  Add,
	"sub":  Sub,
	"call": Call,
	"ret":  Ret,
	"nop":  Nop,
	"stop": Stop,
	"dbg":  Dbg,
}
