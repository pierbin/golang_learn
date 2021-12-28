# golang_learn
It is a personal learn golang lab.

Golang has 5 differences with other Object-Oriented languages.
1. The first things is you should take care of types in Go.

2. The biggest difference is that GO does not provide Classes, it has Structs. They are basically stripped down classes, 
bundles of data with methods. 

Struct can have many attributes, and you can attach methods that any instances of that struct can call. 
Methods, can also be public or private, as well as attributes using both UpperCamelCase for public and 
lowerCamelCase for privates.

The key difference between GO structs and classes is that it does not allow inheritance.
They do not have inherent constructors as classes do and that interfaces are implemented implicitly, 
which mean that you do not have to specify that a class implements an interface, if it implements the methods
of the interface, it implies it implements the interface as well.

3. Golang does not quite really supports Inheritance, you could use inheritance but what you would be using deep down is
Composition, which is the recommended way. 

Composition is basically using a class that uses some other classes, 
like assembling the different components of a PC and making them work together. 
Composition in GO is normally implemented using dependency injection, 
in which you pass the dependencies or components to a class instead of letting the class instantiate it itself.
This can make the code much cleaner since you can track the dependencies to a common source instead of having them all
over the place. It also helps you to build unit tests since you it’s easier to mock the dependencies.

4. GO is missing many of the main data structures common in other languages like built-in trees and sets.

First, you’ve got arrays and slices who behave in a very similar way with the main difference being that arrays are
immutable. Second, you have structs. Finally, there’s the map, the map is pretty much like any other map or dictionary 
in any other language, you have a key-value data structure.

5. Error handling is also really different in GO.

First, because GO lacks one key component in most programming languages: the try-catch statement. 
So GO does not let you wrap your entire code into a single try-catch to screw up as much as you’d like.
If you think that a certain part of your code is risky or error-prone, 
you should handle that part individually with its one conditions and proper response.
So the way that GO does this is by having two responses on its functions, 
normally the standard says that you should return an error alongside your response and wherever you invoke that function,
you’d either get that error and validate it with an if statement,
or you’ll just ignore it using a blank identifier (_) to move on.