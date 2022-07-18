# Pointer 

Pointer is a variable which store the address of other variable. 

### Declaring of Pointer 
```
var y *int 

// It will accept the address of int value only.

```

## Initialization of Pointer

```
var x = 100
var pointer *int = &x
```

## Important Point on Pointer :)
- The default value of a pointer is always a nil

```
import "fmt"
func main() {
   var s *int 
   fmt.Println(" s", s)  // nil
}
```

- On Declaring type of pointer it always store that type of variable. 

- To overcome above point remove type of pointer . It act like var interface. 

```
var y = &x 
```

## Derefreshing Pointer 

The derefreshing of pointer is by using * 

```
x= 100
y = &x 
fmt.Println(y) // address of x
fmt.Println(*y)  // value of x
```

### Important Point on Derefreshing 

- We can change the value of Pointer address 

```
x = 100
y = &x
*y = 200
fmt.Println("x :", x)  // x : 200
```

## Pointer to Go in function 

 In GO we can pass the pointer in function as like variable 

 ```
 func print(a *int) {
    // change the value of a 
    *a = 30
 }
 var x = 100
 var a *int = &x 
 print(a)  // 30
 ```

 ## Pointer to Struct 

 We can use pointer to struct by using & 

 ```
type Studen struct{
    Name string 
    Id   int
}
var x := Student{"Surya", 1}
pts := &x
fmt.Println(pts)  // &{Surya 1}
fmt.Println((*pts).Name)   // Surya

// Golang allows (*pts).Name as pts.Name

fmt.Println(pts.Name)  // Surya
 ```
It also allow to change the value from Pointer variable 

```
pts.Name = "Sai"
fmt.Println(x.Name) // Sai
```

## Pointer to Pointer In Go 

The pointer to pointer in Go is give the address value of that pointer. 

We can understand with the example 
```
var x = 100
fmt.Println("Value of x :", x)
var ptr1 = &x
fmt.Println("Value of ptr1", ptr1)
var ptr2 = &ptr1
fmt.Println("value of Address at ptr2: *ptr2", *ptr2) // value of Address at ptr2: *ptr2 0xc000014140
fmt.Println("*(value of address of ptr2)", **ptr2)  // *(value of address of ptr2) 100

```
