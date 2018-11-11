package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/pkg/profile"
)

type ctxKey int

const scopeCtx ctxKey = iota

type Complex struct {
	Field1  string
	Field2  string
	Field3  string
	Field4  bool
	Field5  int32
	Field6  string
	Field7  float32
	Field8  string
	Field9  string
	Field10 string
	Field11 bool
	Field12 int32
	Field13 string
	Field14 float32
	Field15 string
	Field16 string
	Field17 string
	Field18 bool
	Field19 int32
	Field20 string
	Field21 float32

	FieldStruct1 Struct1
	FieldStruct2 []Struct2

	FieldFunc1 func() error
	FieldFunc2 func(string)
}

type Struct1 struct {
	Field1  string
	Field2  string
	Field3  string
	Field4  bool
	Field5  int32
	Field6  string
	Field7  float32
	Field8  string
	Field9  string
	Field10 string
	Field11 bool
	Field12 int32
}

type Struct2 struct {
	Field1 string
	Field2 string
	Field3 string
	Field4 bool
	Field5 int32
	Field6 string
	Field7 float32
}

func main() {
	p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	defer p.Stop()
	t := time.Now()
	for i := 0; i < 1000000; i++ {
		err := fn1()
		if err.Error() != "asd" {
			fmt.Println("not working")
		}
	}
	fmt.Println(time.Since(t))
}

func fn1() error {
	c := getComplex()
	return fn2(&c)
}

func fn2(c *Complex) error {
	return fn3(c)
}

func fn3(c *Complex) error {
	return fn4(c)
}

func fn4(c *Complex) error {
	return fn5(c)
}

func fn5(c *Complex) error {
	return errors.New("asd")
}

/*
	Context based
*/

func fn1Ctx() error {
	c := getComplex()
	ctx := context.Background()
	ctx = context.WithValue(ctx, scopeCtx, c)

	return fn2Ctx(ctx)
}

func fn2Ctx(ctx context.Context) error {
	return fn3Ctx(ctx)
}

func fn3Ctx(ctx context.Context) error {
	return fn4Ctx(ctx)
}

func fn4Ctx(ctx context.Context) error {
	return fn5Ctx(ctx)
}

func fn5Ctx(ctx context.Context) error {
	return errors.New("asd")
}

func getComplex() Complex {
	return Complex{
		Field1:  "test",
		Field2:  "test",
		Field3:  "test",
		Field4:  true,
		Field5:  12,
		Field6:  "test",
		Field7:  12.2,
		Field8:  "test",
		Field9:  "test",
		Field10: "test",
		Field11: true,
		Field12: 12,
		Field13: "test",
		Field14: 12.2,
		Field15: "test",
		Field16: "test",
		Field17: "test",
		Field18: false,
		Field19: 12,
		Field20: "test",
		Field21: 12.2,

		FieldStruct1: Struct1{
			Field1:  "test",
			Field2:  "test",
			Field3:  "test",
			Field4:  true,
			Field5:  12,
			Field6:  "test",
			Field7:  12.2,
			Field8:  "test",
			Field9:  "test",
			Field10: "test",
			Field11: true,
			Field12: 12,
		},
		FieldStruct2: []Struct2{
			{
				Field1: "test",
				Field2: "test",
				Field3: "test",
				Field4: true,
				Field5: 12,
				Field6: "test",
				Field7: 12.2,
			},
			{
				Field1: "test",
				Field2: "test",
				Field3: "test",
				Field4: true,
				Field5: 12,
				Field6: "test",
				Field7: 12.2,
			},
		},
		FieldFunc1: func() error {
			return nil
		},
		FieldFunc2: func(str string) {
			fmt.Println("hm...")
		},
	}
}
