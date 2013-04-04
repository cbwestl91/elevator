package elevdriver

/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"

func IoInit() bool {
	return bool(int(C.io_init())!=1)
}

func Set_bit(channel int) {
	C.io_set_bit(C.int(channel))
}

func Clear_bit(channel int) {
	C.io_clear_bit(C.int(channel))
}

func Write_analog(channel, value int) {
	C.io_write_analog(C.int(channel), C.int(value))
}

func Read_bit(channel int) bool {
	return int(C.io_read_bit(C.int(channel))) != 0
}

func Read_analog(channel int) int {
	return int(C.io_read_analog(C.int(channel)))
}
