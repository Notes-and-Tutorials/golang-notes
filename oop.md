
# Docs
Is Go an object-oriented language?
Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).

Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java.

# Todd McLeod
"
Go is Object Oriented
Encapsulation
state ("fields")
behavior ("methods")
exported & unexported; viewable & not viewable
Reusability
inheritance ("embedded types")
Polymorphism
interfaces
Overriding
"promotion"
Traditional OOP
Classes
data structure describing a type of object
you can then create "instances"/"objects" from the class / blueprint
classes hold both:
state / data / fields
behavior / methods
public / private
Inheritance
In Go:
you don't create classes, you create a TYPE
you don't instantiate, you create a VALUE of a TYPE
User defined types
We can declare a new type
foo 
the underlying type of foo is int
int conversion
int(myAge) 
converting type foo to type int 
IT IS A BAD PRACTICE TO ALIAS TYPES 
one exception: 
if you need to attach methods to a type 
see the time package for an example of this godoc.org/time 
type Duration int64 
Duration has methods attached to it
Named types vs anonymous types
Anonymous types are indeterminate. They have not been declared as a type yet. The compiler has flexibility with anonymous types. You can assign an anonymous type to a variable declared as a certain type. If the assignment can occur, the compiler will figure it out; the compiler will do an implicit conversion. You cannot assign a named type to a different named type.
Padding & architectural alignment
Convention: logically organize your fields together. Readability & clarity trump performance as a design concern. Go will be performant. Go for readability first. However, if you are in a situation where you need to prioritize performance: lay the fields out from largest to smallest, eg, int 64, int64, float32, bool

"

