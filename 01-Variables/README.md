# Variables

- Variable declaration
    - var foo int
    - var foo int = 42
    - foo := 43
- Redeclaration and shadowing, can't redeclare variables, but can shadow them
- All variables must be used
- Visibility
    - lower case first letter for package scope
    - upper case first letter to export
    - no private scope
- Naming conventions
    - Pascal or camelCase
        - Capitalize acronyms(HTTP, URL)
    - As short as reasonable
        - longer names for longer lives
- Type conversions
    - destinationType(variable)
    - use strconv package for strings