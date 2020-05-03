# Liskov Substitution Principle (LSP)

```text
What is wanted here is something like the following substitution property: If for each object of type S there is an object of 
type T such that for all programs P defined in terms of T, the behavior of P is unchanged when is substituted for then S is a 
subtype of T. -- Barbara Liskov
```

This principle was coined by [Barbara Liskov](https://www.youtube.com/watch?v=_jTc1BTFdIo) in 1988. Simplifying what she wrote in 
his paper this what this principle means is that **every subtype of an entity should be replaceable by its base entity**.

_"The importance of this principle becomes obvious when you consider the consequences of violating it. Presume that we have a function f that takes, as its argument, a pointer or reference to some base class B. Presume also
that there is some derivative D of B which, when passed to f in the guise of B, causes f to misbehave. Then D violates the LSP. Clearly D is Fragile in the presence of f.
The authors of f will be tempted to put in some kind of test for D so that f can behave properly when a D is
passed to it. This test violates the OCP because now f is not closed to all the various derivatives of B. Such tests are
a code smell that are the result of inexperienced developers (or, what’s worse, developers in a hurry) reacting to
LSP violations."_

OK, but in Go we don't have sub types because we don't have inheritance, yet we have interfaces and composition. The next piece of
code should clarify this concept to us.

```go
package wonderful_package
import "reflect"

// This is similar to the go's io.Reader abstraction
type Reader interface {
     Read(buf []byte) (n int, err error)
     Close() (err error)
}

type SomeReader struct {}
func (a *SomeReader) Close() (err error) { 
    // Do stuff
    return err
}

func (s *SomeReader) Read(buf []byte) (n int, err error) {
    // Do some stuff 
    return n, err
}

type AnotherReader struct {}
func (a *AnotherReader) Close() (err error) { 
    // Do another stuff
    return err
}

func (a *AnotherReader) Read(buf []byte) (n int, err error) {
    // Do stuff 
    err = a.Close() // Side effect: closing itself!!!   
    return n, err
}

func StringToStream(r Reader, data string) (err error) {
    buf := make([]byte, len(data))
    for i := 0; i < len(data); i++ {
        buf = append(buf, data[i])
    }
    
    _, err = r.Read(buf)
    if err != nil {
        return err
    }
    
    if reflect.TypeOf(r).String() == "wonderful_package.AnotherReader" {
        return nil
    } else if reflect.TypeOf(r).String() == "wonderful_package.SomeReader" {
        err = r.Close()
        return err
    } else {
        return err
    }
}
``` 

In this example we can see how the struct _AnotherReader_ violates the LSP, forcing us to violate the OCP in the _StringToStream_ function.
The later is know fragile and is forced to know every subtype of the _Reader_ abstraction, so it is not closed for modification.
When we see an "if/else" chain, like the one present in _StringToStream_, where types are being used to determine the behaviour it
is a clear violation of the LSP principle. 
Personally I've verified lot of times that if we take an approach "more functional" while we write code, even if we are doing OOP,
we tend to avoid this pitfalls.
Robert in his book says _"A model, viewed in isolation, cannot be meaningfully validated."_ and he uses an example of a square
inheriting methods and properties from a rectangle to explain this: We all know that from a mathematical point of view, a
square is a specific case of a rectangle where the height and width are equal. But from a behavioural point of view
they are not the same. If we'd try to change the width of a square with a method _SetWidth_ inherited from the rectangle we also
would be changing its height otherwise it wouldn't be a mathematical square anymore. It does not happen the same with a rectangle. Since
the square subtype and the rectangle base type are not the interchangeable they violate the LSP and will lead to a violation in the OCP
in there clients.

### Designing by contract

[Bertrand Meyer](https://es.wikipedia.org/wiki/Bertrand_Meyer) created a technique called _designing by contract_. Contracts are
"reasonable assumptions" of what the behaviour of an entity should be. For example, a reasonable assumption for our previous example
should be _"A Reader should continue open after its Read method is called"_. Some of the ways Robert propose to us to create those
contracts is unit testing:

```text
By thoroughly testing the behavior of a class, the unit tests make the behavior of the class clear. Authors of client code will want 
to review the unit tests so that they know what to reasonably assume about the classes they are using.
```

