package main

import "unsafe"

var allocs = map[int32][]byte{}

//export alloc
func alloc(size int32) int32 {
	if size <= 0 {
		return 0
	}
	buf := make([]byte, size)
	ptr := int32(uintptr(unsafe.Pointer(&buf[0])))
	allocs[ptr] = buf
	return ptr
}

//export dealloc
func dealloc(ptr int32) {
	delete(allocs, ptr)
}

//export validate
func validate(platePtr, plateLen, countryPtr, countryLen int32) int32 {
	if IsValid(string(view(platePtr, plateLen)), string(view(countryPtr, countryLen))) {
		return 1
	}
	return 0
}

func view(ptr, length int32) []byte {
	b, ok := allocs[ptr]
	if !ok || length <= 0 {
		return nil
	}
	if int(length) > len(b) {
		length = int32(len(b))
	}
	return b[:length:length]
}

func main() {}
