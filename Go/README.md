### **Variables in Go**  

#### **1. `var` (Variable Declaration)**
- Used to declare **mutable** variables (values can be changed).
- The type is **optional** (Go infers it automatically).
- Can be declared at both **global and local** levels.
- Example:  
  ```go
  var number = 100      // Type inferred as int
  var name string = "Go" // Explicit type declaration
  ```

#### **2. `const` (Constant Declaration)**
- Used to declare **immutable** values.
- The value must be known **at compile-time**.
- Example:  
  ```go
  const pi = 3.14159
  ```

#### **3. `:=` (Short Variable Declaration)**
- Used **only inside functions** for concise variable declaration.
- The type is always **inferred** automatically.
- Example:  
  ```go
  varId := 22 // Equivalent to: var varId int = 22
  ```

---

### **Key Differences Between `var`, `const`, and `:=`**  

| Keyword | Purpose                | Mutability | Type Inference | Scope | Reassignable | Redeclarable | Example          |
|---------|------------------------|------------|----------------|-------|--------------|--------------|------------------|
| `var`   | Declare a variable      | Mutable    | Optional       | Global & Local | ✅ Yes | ❌ No (same scope) | `var x int = 10` |
| `const` | Declare a constant      | Immutable  | No             | Global & Local | ❌ No | ❌ No | `const pi = 3.14`|
| `:=`    | Shorthand for variable  | Mutable    | Yes (inferred) | Local (Inside functions only) | ✅ Yes | ❌ No (same scope) | `x := 10` |

---

### **Go Data Types**  

Go provides different categories of types:  

#### **1. Basic Types**  
| Type      | Description                           | Example          |
|-----------|--------------------------------------|------------------|
| `bool`    | Boolean (true or false)             | `var isActive bool = true` |
| `string`  | Textual data                        | `var message string = "Hello"` |
| `int`     | Integer (default `int` size)        | `var age int = 25` |
| `int8`    | 8-bit integer (-128 to 127)         | `var num int8 = 120` |
| `int16`   | 16-bit integer                      | `var num int16 = 32000` |
| `int32`   | 32-bit integer                      | `var num int32 = 2147483647` |
| `int64`   | 64-bit integer                      | `var num int64 = 9223372036854775807` |
| `uint`    | Unsigned integer (default size)     | `var count uint = 5` |
| `float32` | 32-bit floating point               | `var price float32 = 19.99` |
| `float64` | 64-bit floating point (default)     | `var gpa float64 = 3.9` |
| `complex64`  | Complex number with float32 real and imaginary parts | `var c complex64 = 3 + 4i` |
| `complex128` | Complex number with float64 real and imaginary parts | `var c complex128 = 5 + 12i` |

---

#### **2. Composite Types**  
| Type      | Description                           | Example          |
|-----------|--------------------------------------|------------------|
| `array`   | Fixed-length collection of elements | `var arr [5]int` |
| `slice`   | Dynamically sized array-like type   | `var s []int = []int{1, 2, 3}` |
| `map`     | Key-value data structure            | `var m map[string]int` |
| `struct`  | Custom data structure               | `type Person struct {name string; age int}` |
| `pointer` | Stores memory address of a variable | `var p *int` |

---

#### **3. Special Types**  
| Type      | Description                           | Example          |
|-----------|--------------------------------------|------------------|
| `interface{}` | Empty interface (any type)      | `var anyData interface{}` |
| `chan`     | Channel for goroutines             | `var ch chan int` |
| `func`     | Function type                      | `var f func(int) int` |

---

