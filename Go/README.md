### **Variable in go**:
1. **`var`**:  
   - Used to declare **variables** (with optional explicit types). These can be **mutable**, meaning their value can be changed later.
   - Example: `var number = 100`

2. **`const`**:  
   - Used to declare **constants**, which cannot be changed once assigned.  
   - Constants must have a value at compile-time and are often used for fixed values.
   - Example: `const pi = 3.14159`

3. **`:=`**:  
   - **Shorthand** for declaring and initializing variables **inside functions**. The type is inferred automatically based on the assigned value.
   - Example: `varId := 22`

### **Key Differences**:
| Keyword | Purpose                | Mutability | Type Inference | Example          |
|---------|------------------------|------------|----------------|------------------|
| `var`   | Declare a variable      | Mutable    | Optional       | `var x int = 10` |
| `const` | Declare a constant      | Immutable  | No             | `const pi = 3.14`|
| `:=`    | Shorthand for variable  | Mutable    | Yes (inferred) | `x := 10`        |
