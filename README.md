# URL Friendly Base64 Hashing

A simple and fast Go module to hash integers to URL friendly base64 encodings.
This algorithm converts each 6 bits of a number to one character meaning
it should create the miniumum possible base64 string length and is easily reproducable
in any programming language.

Usage:
```go
package main

import (
  "fmt"

  "github.com/nbwoodward/base64"
)

func main() {
  num := 314159
  hash, _ := base64.IntToString(num)
  intr, _ := base64.StringToInt(hash)

  fmt.Println(num, " -> ", hash, " -> ", intr)
  // Prints: "314159 -> 1cIL -> 314159"
}
```

### Limitations
 - Does not hash negative numbers
 - Cannot hash a number greater than the max 64 bit integer (9223372036854775807)

### The Algorithm
This algorithm takes each 6 bits of a number (which yields a number between
0 and 63) and maps that to an alphanumeric character or dash or underscore (62 alphanumeric
characters plus two URL friendly characters)

The algorithm uses bit shifting and map key for memory access to minimize the speed of conversion.
To run the benchmark tests:
```
git clone https://github.com/nbwoodward/base64
cd base64
go test -bench=.
```
