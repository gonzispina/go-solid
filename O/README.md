# Open-Closed Principle (OCP)

```text
Software entities should be open for extension, but closed for modification. –- Bertrand Meyer, Object-Oriented Software Construction
```

When a single change to a program results in a cascade of changes to dependent modules, the design smells of Rigidity. If the OCP
is applied well, in the future we will only have to add new code instead of modifying the one that is already working.

- **Open for extension**: As the requirements of the business change, we implement our entities and extend their behaviour
with new capabilities.

- **Closed for modification**: This extension should not result in changes in the source code of the package that is being extended

This is simple but it is not that intuitive. How do we add new behaviour without touching the source code? Since we don't have classes
in Go, we don't have inheritance or abstract classes. But we have something much more powerful and that's **composition**. To achieve
this we must compose structs and interfaces into bigger abstractions. Creating abstractions of concepts by creating interfaces and then
creating multiple implementations of it is another way. 
The OCP is at the heart of object-oriented design and attaching to this principle is what will give us the greatest benefits of OOP. 
Just using an OOPL won't be sufficient to satisfy this principle. In general, no matter how “closed” a module is, there will always be
some kind of change against which it is not closed. There is no model that can satisfy all the contexts and that is where the difficulty
of this principle resides.

Sometimes it is a good idea not to try creating the abstractions. Writing the code as we would not expect any change on it until
the necessity of changing emerges. This will make us avoid _Needless Complexity_ in our code until that moment and guarantee 
that the abstraction we'll write is going to be the one that suits the model.  

## Examples

- [Wrong](https://github.com/gonzispina/go-solid/tree/master/O/wrong): in this example we see an un-extensible transaction implementation. 
New requirements to this portion of code will end up with us adding methods and validations in it as it is done in the _cancel_ method.

- [Right](https://github.com/gonzispina/go-solid/tree/master/O/right): here we can see the principle applied. There is an abstraction as 
interface to represent our transaction, a base implementation which is closed and has the common behaviour of every transaction, and two 
more structs implementing the class and being a composition of the _BaseTransaction_. It is clearly visible how the _NationalTransaction_
is extending the base implementation by adding the cancel method. Further changes in this kind of code would result in the creation of new 
structs instead of modifying the base one.
 
