# Single Responsibility Principle (SRP)

```text
A class should have only one reason to change. -- Robert C. Martin
```

This principle comes to save us from fragility. The opposite to a fragile
model is a ductile model and that is what we want to achieve, ductility.
As Robert says in his book:

_"If a class has more than one responsibility, then the responsibilities 
become coupled. Changes to one responsibility may impair or inhibit 
the ability of the class to meet the others. This kind of coupling 
leads to fragile designs that break in unexpected ways when changed."_

Since we don't have classes in Go, we first need to identify which
is the most basic unit we could use to wrap a concept and that are packages. 
Go has no classes but it has packages. 

### Naming
So in this context, we can define responsibility as "a reason to change". If
we can think of more than one reason to change a package, that package
has more than one responsibility. The most obvious example of this are
packages named _"utils"_ or _"common"_.
So the best way to apply this principle is to start with a good and specific
name for the package. An example of good naming in the Go's standard
library are the packages `encoding/json` and `encoding/xml`.
When we read _encoding_ we obviously expect something that lets us encode
and decode, but from where to where? 
And that's where _json_ and _xml_ appear. We can se the **cohesiveness** between
the _encoding_ and _json_ parts but we can also see that the _encoding_ package 
is not **coupled** to one and only notation. If we go even further, the _xml_ package
is not coupled with the _json_ package either, and changes in one of them probably won't
have repercussions in the other.

Stretching the packages to be as small as possible and then combining the to build 
a bigger model is what this principle is about.

### Examples

- [Wrong](https://github.com/gonzispina/go-solid/tree/master/S/wrong) in this example we can see how the persistence layer is coupled with 
to the business logic of the user's implementation. 

- [Right](https://github.com/gonzispina/go-solid/tree/master/S/right): simple but obvious, we moved the persistence logic to its
own package and isolated the business logic in its own package.