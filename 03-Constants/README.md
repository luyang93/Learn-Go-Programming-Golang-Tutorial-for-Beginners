# Constants

- Naming convention
- Typed constants
- Untyped constants
- Enumerated constants
- Enumeration expressions

---

- Immutable, but can be shadowed
- Replaces by the compiler at compile time
    - Value must be calculable at compile time
- Named like variables
    - pascalCase for exported constants
    - camelCase for internal constants
- Typed constants work like immutable variables
    - Can interoperate only with same type
- Untyped constants work like literals
    - Can interoperate wth similar types
- Enumerated constants
    - Special symbol iota allows related constants to be created easily
    - Iota starts at 0 in each const block and increments by one
    - Watch out of constant values that match zero values for variables
- Enumerated expressions
    - Operations that can be determined at compile time are allowed
        - Arithmetic
        - Bitwise operations
        - Bitshift