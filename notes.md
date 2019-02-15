
## Reading documentation
### package.Identifier 
ex: `fmt.Println()`
`func Println(a ...interface{})(n int, err error)`
`(a ...interface{})` what it accepts as arguments  
  - `...` means it takes a variadic parameter (unlimited number)
  - `interface{}` of any type (empty interface)  
`(n int, err error)` what it returns  
> Don't have to catch returns for all functions
 
### short declaration operator
`:=`
`x := 42`
both declares variable and assigns a value (of a certain type) [in other words, initialize]
After to resign just use `=`

### var keyword
`var x = 42`
- can use inside func body
- must use outside of func body (scope would be package level )
`var z int`
-  `keyword identifier type` (same structure for everything)
- declare variable z
- type set as int
- auto assign the zero value

## Random
- no `===`
