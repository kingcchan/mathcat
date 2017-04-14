// Copyright 2016 Steven Oud. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.

package mathcat

import (
	"fmt"
	"math"
	"math/rand"
)

type Function struct {
	Arity int
	Fn    func(args []float64) float64
}

type Functions map[string]*Function

// Functions holds all the Function names that are available for use
var FunctionsList []string

var funcs = make(Functions)

func (f Functions) register(name string, Function *Function) {
	FunctionsList = append(FunctionsList, name)
	f[name] = Function
}

func RegisterFunc(name string, Function *Function) {
	funcs.register(name, Function)
}

func init() {
	funcs.register("abs", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Abs(args[0])
		},
	})
	funcs.register("ceil", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Ceil(args[0])
		},
	})
	funcs.register("floor", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Floor(args[0])
		},
	})
	funcs.register("sin", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Sin(args[0])
		},
	})
	funcs.register("cos", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Cos(args[0])
		},
	})
	funcs.register("tan", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Tan(args[0])
		},
	})
	funcs.register("asin", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Asin(args[0])
		},
	})
	funcs.register("acos", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Acos(args[0])
		},
	})
	funcs.register("atan", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Atan(args[0])
		},
	})
	funcs.register("log", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Log(args[0])
		},
	})
	funcs.register("max", &Function{
		Arity: 2,
		Fn: func(args []float64) float64 {
			return math.Max(args[0], args[1])
		},
	})
	funcs.register("min", &Function{
		Arity: 2,
		Fn: func(args []float64) float64 {
			return math.Min(args[0], args[1])
		},
	})
	funcs.register("sqrt", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return math.Sqrt(args[0])
		},
	})
	funcs.register("rand", &Function{
		Arity: 0,
		Fn: func(_ []float64) float64 {
			return rand.Float64()
		},
	})
	funcs.register("fact", &Function{
		Arity: 1,
		Fn: func(args []float64) float64 {
			return float64(Factorial(int64(args[0])))
		},
	})
	funcs.register("gcd", &Function{
		Arity: 2,
		Fn: func(args []float64) float64 {
			return Gcd(args[0], args[1])
		},
	})
	funcs.register("list", &Function{
		Arity: 0,
		Fn: func(_ []float64) float64 {
			for _, name := range FunctionsList {
				fmt.Printf(name + " ")
			}
			fmt.Println()
			return 0
		},
	})
}

// Factorial calculates the factorial of number n
func Factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}

// Gcd calculates the greatest common divisor of the numbers x and y
func Gcd(x, y float64) float64 {
	for y != 0 {
		x, y = y, math.Mod(x, y)
	}

	return x
}
