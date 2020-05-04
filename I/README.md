# Interface Segregation Principle (ISP)

```text
Clients should not be forced to depend on methods that they do not use. -- Robert C. Martin
```

This principle is about dealing with what Robert calls "fat" interfaces. We usually tend to pollute interfaces
by adding methods only for the benefit of one of its subtypes. This forces us to implement this new method in all of the 
rest of the clients, even if concept encapsulated in it has nothing to do with the new method, and deriving into an LSP violation
that can compromise maintenance and re-usability.
When clients are forced to depend on methods they don't use, they are attached to changes on those when they are necessary in their 
fellow subtypes. This can be seen as a coupling between all the clients of that abstraction. Since all the clients are coupled,
changes in one may lead to a recompiling, building or deploying of the other ones, even if they weren't modified. This is what have
defined as _viscosity_. If we go into a more macro situation this happens a lot within the mobile world, where the Android client and 
the IOS client are coupled by a poorly designed backend API. 

### Separation through multiple implementation 

In Go and in general OOPLs the process of segregating interfaces is a matter of isolating the concepts that we are working with and stretching
them as much as possible until we have an interface with the smallest amount of methods. The ISP acknowledges that different entities 
may use different interfaces that are not between the same conceptual boundaries.

### Examples

_Note: I took this example of a friend's workshop in one of my latest jobs_

- [Wrong](https://github.com/gonzispina/go-solid/tree/master/I/wrong): in this example we can see a single interface named _Person_ which has
four methods_ `Name`, `Eat`, `Sleep` and `Work`. This seems perfectly reasonable, but what happen
if our person is a unemployed person or a retired person who doesnt work? Well.. we will still have to implement the method
work, even if these subtypes won't use. This is a clear pollution among the subtypes.

- [Right](https://github.com/gonzispina/go-solid/tree/master/I/right): in the right example we have segregated the interface into `Person` and 
`SlaveSystem`. While our `McDonaldsEmployee` implements both, our `Retired` model only implements the `Person` interface.

  