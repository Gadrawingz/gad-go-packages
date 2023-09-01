# Packages in Go

> In Go, a package is a directory with one or more go code files with .go extension. The package group related features together which makes it easier to maintain, reuse and share functionality.
(In go, every package is defined with a different name that closely 
resembles the functionality it performs. For example, the most used 
package "fmt" short for format contains string formatting and printing 
functions which are used in other go files.)

> Every go file belongs to a package. The first line in every file will 
be the package declaration as package **packageName**. 

> A package or a member belonging to a package can be exported to another package. Other packages can import and reuse the functionality of the exported package.


## Advantages of packages
- Packages organizes related code and functionality together.
  (This makes it easier to maintain and reuse code.)
- Packages speed up the compilation process by compiling only the parts 
that have changed. For example, the "fmt" package is not compiled every time the program changes. 
- It reduces the naming conflict. You can have functions with the same names in different packages.

## Types of Packages
____________________
In go language, There are two types of packages, namely:
- **Executable package:** (it is an executable and will contain the main() function. 
- **Utility package:** is a package that is not executable but provides utility 
functionality and other assets to the executable package.

The starting point of a go program is the main package. 
It is a special package that is used in programs that are compiled to be 
executable. The main() function is the entry point of an executable program.

### Example of executable package
_________________________________
'''
package main

import "fmt"

func main() {

fmt.Println ("Hello World")

}

## How to import a package in Go?

### Example 1: Using multiple import statements:
import "fmt"
import "math"
import "math/rand"

> Here, on import "math/rand", rand is nested inside the math package 
> as a subdirectory.

### Example 2: Using Factored import statements:
import(
"fmt"
"math"
"math/rand"
)


## Last statement
## Exported names in a Package

> Any variable, type, or function that starts with capital letter is exported 
> and is visible outside the package. Anything that does not start with a 
> capital letter is not exported and is available only within the package.

### When you import a package, you can only access the exported names.

**Note:** Go language consists of many code standard library packages which 
are often used in all programs. But you can build your custom package in go.

