package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Integer Types
	var int8Num int8 = 127                   // Max value for int8
	var int16Num int16 = 32767               // Max value for int16
	var int32Num int32 = 2147483647          // Max value for int32
	var int64Num int64 = 9223372036854775807 // Max value for int64
	var intNum int = 100                     // Default int (depends on system: 32-bit or 64-bit)

	var uint8Num uint8 = 255                    // Max value for uint8
	var uint16Num uint16 = 65535                // Max value for uint16
	var uint32Num uint32 = 4294967295           // Max value for uint32
	var uint64Num uint64 = 18446744073709551615 // Max value for uint64
	var uintNum uint = 200                      // Default uint (depends on system)

	// Floating-Point Numbers
	var float32Num float32 = 3.1415927         // Approx 6-7 decimal precision
	var float64Num float64 = 3.141592653589793 // High precision (15-16 decimals, recommended)

	// Complex Numbers
	var complex64Num complex64 = complex(1.5, -2.5)    // float32 real & imaginary
	var complex128Num complex128 = complex(3.14, 2.71) // float64 real & imaginary

	// Byte and Rune Types
	var charByte byte = 'A' // ASCII character (alias for uint8)
	var charRune rune = '✓' // Unicode character (alias for int32)

	// Printing all values
	fmt.Println("Integer Types:")
	fmt.Println("int8:", int8Num, "int16:", int16Num, "int32:", int32Num, "int64:", int64Num, "int:", intNum)
	fmt.Println("uint8:", uint8Num, "uint16:", uint16Num, "uint32:", uint32Num, "uint64:", uint64Num, "uint:", uintNum)

	fmt.Println("\nFloating-Point Types:")
	fmt.Println("float32:", float32Num, "float64:", float64Num)

	fmt.Println("\nComplex Numbers:")
	fmt.Println("complex64:", complex64Num, "complex128:", complex128Num)
	fmt.Println("Real Part of complex128:", real(complex128Num), "Imaginary Part:", imag(complex128Num))

	fmt.Println("\nCharacter Types:")
	fmt.Println("byte:", charByte, "as char:", string(charByte))
	fmt.Println("rune:", charRune, "as char:", string(charRune))

	// Additional Math Functions for Complex Numbers
	fmt.Println("\nComplex Number Operations:")
	fmt.Println("Magnitude (Abs):", cmplx.Abs(complex128Num))
	fmt.Println("Phase (Angle in Radians):", cmplx.Phase(complex128Num))
	fmt.Println("Complex Conjugate:", cmplx.Conj(complex128Num))
}
