package misc

import "encoding/binary"

func SignExt(i uint64, n uint32) int64 {
	return (((int64)(i) << (64 - (n))) >> (64 - (n)))
}

func Imm22(sign uint64, imm5c uint64, imm9d uint64, imm7b uint64) int64 {
	return SignExt(sign<<21|imm5c<<16|imm9d<<7|imm7b, 22)
}

func Imm14(sign uint64, imm6d uint64, imm7b uint64) int64 {
	return SignExt(sign<<13|imm6d<<7|imm7b, 14)
}

func IntPow(x int64, n int64) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	y := IntPow(x, n/2)

	if n%2 == 0 {
		return y * y
	}

	return x * y * y
}

func ZeroExt(value int64, pos int64) int64 {
	and := IntPow(2, pos) - 1

	return value & and
}

func BytesToInt64(bytes []byte, count int64) int64 {
	switch count {
	case 1:
		return int64(bytes[0])
	case 2:
		return int64(binary.LittleEndian.Uint16(bytes))
	case 4:
		return int64(binary.LittleEndian.Uint32(bytes))
	case 8:
		return int64(binary.LittleEndian.Uint64(bytes))
	}

	return 0
}
