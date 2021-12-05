package pointer

import "github.com/PumpkinSeed/memory-try/types"

//go:noinline
func Test() {
	var s = &types.Test{}

	test(s)
}

//go:noinline
func TestWithAlloc() *types.Test {
	s := wrappedTest()
	s.Field1 = "test"
	return s
}

//go:noinline
func wrappedTest() *types.Test {
	var s = &types.Test{}

	test(s)
	s.Field12 = "test"
	return s
}

//go:noinline
func test(s *types.Test) {
	setField1(s)
	setField2(s)
	setField3(s)
	setField4(s)
	setField5(s)
	setField6(s)
	setField7(s)
	setField8(s)
	setField9(s)
	setField10(s)
	setField11(s)
	setField12(s)
	setField13(s)
	setField14(s)
	setField15(s)
	setField16(s)
	setField17(s)
	setField18(s)
	setField19N1(s)
	setField19N2(s)
	setField19N3(s)
	setField19N4(s)
	setField19N5(s)
	setField19N6(s)
	setField19N7(s)
	setField19N8(s)
	setField19N9(s)
	setField19N10(s)
	setField19N11(s)
	setField19N12(s)
	setField19N13(s)
	setField19N14(s)
	setField19N15(s)
	setField19N16(s)
	setField19N17(s)
	setField19N18(s)
}

//go:noinline
func setField1(s *types.Test) {
	s.Field1 = "random"
}

//go:noinline
func setField2(s *types.Test) {
	s.Field2 = 41243
}

//go:noinline
func setField3(s *types.Test) {
	s.Field3 = true
}

//go:noinline
func setField4(s *types.Test) {
	s.Field4 = "random"
}

//go:noinline
func setField5(s *types.Test) {
	s.Field5 = 5342
}

//go:noinline
func setField6(s *types.Test) {
	s.Field6 = "random"
}

//go:noinline
func setField7(s *types.Test) {
	s.Field7 = 42344
}

//go:noinline
func setField8(s *types.Test) {
	s.Field8 = "random"
}

//go:noinline
func setField9(s *types.Test) {
	s.Field9 = 66553
}

//go:noinline
func setField10(s *types.Test) {
	s.Field10 = "random"
}

//go:noinline
func setField11(s *types.Test) {
	s.Field11 = "random"
}

//go:noinline
func setField12(s *types.Test) {
	s.Field12 = "random"
}

//go:noinline
func setField13(s *types.Test) {
	s.Field13 = 7765
}

//go:noinline
func setField14(s *types.Test) {
	s.Field14 = "random"
}

//go:noinline
func setField15(s *types.Test) {
	s.Field15 = "random"
}

//go:noinline
func setField16(s *types.Test) {
	s.Field16 = "random"
}

//go:noinline
func setField17(s *types.Test) {
	s.Field17 = "random"
}

//go:noinline
func setField18(s *types.Test) {
	s.Field18 = "random"
}

//go:noinline
func setField19N1(s *types.Test) {
	s.Field19.Field1 = "random"
}

//go:noinline
func setField19N2(s *types.Test) {
	s.Field19.Field2 = 345345
}

//go:noinline
func setField19N3(s *types.Test) {
	s.Field19.Field3 = true
}

//go:noinline
func setField19N4(s *types.Test) {
	s.Field19.Field4 = "random"
}

//go:noinline
func setField19N5(s *types.Test) {
	s.Field19.Field5 = 6541234
}

//go:noinline
func setField19N6(s *types.Test) {
	s.Field19.Field6 = "random"
}

//go:noinline
func setField19N7(s *types.Test) {
	s.Field19.Field7 = 56554
}

//go:noinline
func setField19N8(s *types.Test) {
	s.Field19.Field8 = "random"
}

//go:noinline
func setField19N9(s *types.Test) {
	s.Field19.Field9 = 6654
}

//go:noinline
func setField19N10(s *types.Test) {
	s.Field19.Field10 = "random"
}

//go:noinline
func setField19N11(s *types.Test) {
	s.Field19.Field11 = "random"
}

//go:noinline
func setField19N12(s *types.Test) {
	s.Field19.Field12 = "random"
}

//go:noinline
func setField19N13(s *types.Test) {
	s.Field19.Field13 = 99554
}

//go:noinline
func setField19N14(s *types.Test) {
	s.Field19.Field14 = "random"
}

//go:noinline
func setField19N15(s *types.Test) {
	s.Field19.Field15 = "random"
}

//go:noinline
func setField19N16(s *types.Test) {
	s.Field19.Field16 = "random"
}

//go:noinline
func setField19N17(s *types.Test) {
	s.Field19.Field17 = "random"
}

//go:noinline
func setField19N18(s *types.Test) {
	s.Field19.Field18 = "random"
}