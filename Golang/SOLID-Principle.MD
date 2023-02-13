# Solid principal in Golang

1. **S**ingle Responsibility principle
1. **O**pen close principle
1. **L**iskov Substitution Principle
1. **I**nterface Segregation Principle
1. **D**ependency Inversion Principle


## Single Responsibilty principle

> A class should have one, and only one, reason to change. –Robert C Martin

This different from saying, "class should do one thing"? Why "one reason to change"?
Which is more distressing? 
    1. You own code is causing your code to change?
    1. Thrid party code is causing your code to change?

```Golang
func (u user) showPost() []Post{
    posts := u.fetchPosts()
    return thridparty.Sort(posts)
}
```

```Golang
func (u user) showPost() []Post{
    posts := u.fetchPosts()
    return my.Sort(posts)
}
```

### Coupling and Cohesion
Coupling is simply a word that describes two things moving together--a movement in one induces a movement in another.

In the context of software, cohesion is the property of describing pieces of code are naturally attracted to one another. They are coherent. They fit well together.

Go has package, and related code reside in a package.

```Go
import (
    net/http
    os/exec
    encoding/json
)
```
Import statement creates a source level coupling between two packages.

Don't use package names like,
```Go
import (
    server // which server? http, https, grpc, 
    common // what is common ?
    util   // utilites for ?
    miscellenous // what is miscelleneous? what is it's scope
)
```
In effect each Go package is itself a small Go program, a single unit of change, with a single responsibility.

## Open Close principle

> Software entities should be open for extension, but closed for modification. –Bertrand Meyer, Object-Oriented Software Construction

```Go
package main

type A struct {
        year int
}
func (a A) Greet() { fmt.Println("Hello ", a.year) }

type B struct {
        A
}
func (b B) Greet() { fmt.Println("Welcome ", b.year) }

func main() {
        var a A
        a.year = 2023
        var b B
        b.year = 2023
        a.Greet() // Hello 2023
        b.Greet() // Welcome 2023
}
```
type A is embeedded in type B. B doesn't have year but it can access it from A. Like in the code.

If b.Greet is not written you will access a.Greet.

```Go
package main

type Cat struct {
        Name string
}
func (c Cat) Legs() int { return 4 }
func (c Cat) PrintLegs() {
        fmt.Printf("I have %d legs\n", c.Legs())
}

type OctoCat struct {
        Cat
}
func (o OctoCat) Legs() int { return 5 }

func main() {
        var octo OctoCat
        fmt.Println(octo.Legs()) // 5
        octo.PrintLegs()         // I have 4 legs
}
```
It says, OctoCat has 4 legs. Because, PrintLegs is calling Legs method of cat returning 4. This can be avoided if you have,
```Go
func PrintLegs(c Cat) {
        fmt.Printf("I have %d legs\n", c.Legs())
}
```
In this case, without Cat, this function won't work. 

OR we can modify to method to accent legs. 
```Go
package main

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Legs() int { return 4 }
func (c Cat) PrintLegs(legs int) {
	fmt.Printf("I have %d legs\n", legs)
}

type OctoCat struct {
	Cat
}

func (o OctoCat) Legs() int { return 5 }

func main() {
	var octo OctoCat
	fmt.Println(octo.Legs())    // 5
	octo.PrintLegs(octo.Legs()) // I have 5 legs
}
```

## Liskov Substitution Principle

Roughly, it states that, two types are substitutable if they exhibit behaviour such that the caller is unable to tell the difference.
In class based language, this principle is used with Abstract classes and inheritance.

Go has embedding and implicit inheritance. Interface is implemented just by implementing the methods of interface.
```Go
type Reader interface {
        // Read reads up to len(buf) bytes into buf.
        Read(buf []byte) (n int, err error)
}
```
Read has byte parameter. Byte can represent string, network stream, or any thing. 
Details, https://gobyexample.com/interfaces


## Interface Segregation Principle
> Clients should not be forced to depend on methods they do not use. –Robert C. Martin

```Go
package main

import (
	"fmt"
)

type AnimalWorld interface {
	Run(a Animal)
	Swim(a Animal)
}

type Dog struct {
	name string
}

func (d Dog) Run() {
	fmt.Println(d.name, "is running")
}
func (d Dog) Swim() {
	fmt.Println(d.name, "is Swimming")
}

func main() {
	d := Dog{name: "Rocky"}
	d.Run()
	d.Swim()
}
```
Better solution would be to have simpler interface,
```Go
type Animal interface {
    Run()
}
type Fish interface{
    Swim()
}

type Amphibian interface{
    Animal
    Fish
}
```

> A great rule of thumb for Go is accept interfaces, return structs. –Jack Lindamood

## Dependency Inversion Principle

> High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions. –Robert C. Martin

We have seen that in Go, we have discrete packages. Each package has single responsbility. 

Code should define dependencies in terms of interfaces. Those interfaces should be factored to describe only the behaviour those functions require. 

In other words, there shouldn’t be much left to do.

Certainly the context of Go, it is implemeted by the structure of your import graph.

In Go, your import graph must be acyclic. A failure to respect this acyclic requirement is grounds for a compilation failure, but more gravely represents a serious error in design.

Import graph of a well designed Go program should be a wide, and relatively flat, rather than tall and narrow.

More https://parserdigital.com/dependency-injection-and-unit-testing-in-golang/

# Reference
1. https://dave.cheney.net/2016/08/20/solid-go-design
