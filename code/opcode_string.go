// Code generated by "stringer -type Opcode"; DO NOT EDIT.

package code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpConstant-1]
	_ = x[OpPop-2]
	_ = x[OpAdd-3]
	_ = x[OpSub-4]
	_ = x[OpMul-5]
	_ = x[OpDiv-6]
}

const _Opcode_name = "OpConstantOpPopOpAddOpSubOpMulOpDiv"

var _Opcode_index = [...]uint8{0, 10, 15, 20, 25, 30, 35}

func (i Opcode) String() string {
	i -= 1
	if i >= Opcode(len(_Opcode_index)-1) {
		return "Opcode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
