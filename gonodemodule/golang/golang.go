package main

import "C"

//export Hello
func Hello() *C.char {
	return C.CString("Hello world!")
}

// required to build
func main() {
}

// https://github.com/charlieduong94/node-golang-native-addon-experiment
// maybe switch to https://www.electronjs.org/docs/tutorial/using-native-node-modules
