
# Go SOLID
How to apply SOLID principles while gophing. But, what are these principles anyway?

## Agile Design

```text
“After reviewing the software development life cycle as I understood it, I concluded that the only software documentation that actually 
seems to satisfy the criteria of an engineering design is the source code listings.” -— Jack Reeves
```

First we need to talk about designing. The design of a code is the code itself. It is not a bunch of UML diagrams and documents, that 
may or may not help us to describe the code, but the code itself. If agility is build software in tiny increments, how can developers 
take the time to ensure that the code is flexible, maintainable and re-usable? 

In agile teams the big picture evolves along with the software, and the team makes its effort to improve the design to be as good it can 
be at that exact moment. That leads us to another question, how do we know if the design is good or not?

### Indentifying poorly designed code

```text
"Everything should be made as simple as possible, but not simpler." -- Albert Einstein
```

Well... thankfully [Robert C. Martin](https://en.wikipedia.org/wiki/Robert_C._Martin) already identified what, in his book _"Agile Software 
Development: Principles, Patterns, and Practices"_, calls the "smells" of a poor design. 

1. **Rigidity**: is the tendency for software to be difficult to change. Even when the change is small, in a rigid code, it will cause a 
cascade of subsequent changes in dependant modules or functions. That is why we usually end up saying _“It was a lot more complicated than 
I thought!”_

2.  **Fragility**: when a design is fragile, the slightest change can break a program in at least one place if not more, even though those 
places do not have a conceptual relationship with the portion of code we needed to change in the first place.

3. **Immobility**: we say that a code is immobile when we identify its re-usability in some other place, but the risk of taking it apart
from its original source is too big.

4. **Viscosity**: software viscosity can be noticed when we see one or more ways of making a change, but the "design preserving" ways are 
much more difficult to make than some "hacks" we usually do to achieve the same functionality.

5. **Needless Complexity**: here we refer to implementation complexity not business complexity which is really different.  Designs are 
unnecessary complex when they have implementations that are not useful yet. This is frequent when developers try to anticipate changes 
to requirements and add some facilities to deal with this potential changes. Preparing code for future changes may seem a good idea at 
first, but being ahead of too many contingencies, leads us to an implementation filled with constructs that are never used. 

6. **Needless Repetition**: cut and paste may be useful but they can be some of the dev's worst enemies. When the same code appears over 
and over again, in slightly different forms, the developers are missing an abstraction. Finding all the repetition and eliminating it 
with an appropriate abstraction **may not be high on their priority list, but it would go a long way toward making the system easier to 
understand and maintain.**

7. **Opacity**: Opacity is the tendency of a module to be difficult to understand. Code can be written in a clear and expressive manner 
or it can be written in an opaque and convoluted manner.

### SOLID: Keeping the design as good as it can be.

Robert not only identified this issues through his experience, but also the solutions to these problems as what he called the SOLID 
principles. SOLID stands for:
 - [Single Responsibility (SRP)](https://github.com/gonzispina/go-solid/tree/master/S)
 - [Open / Closed (OCP)](https://github.com/gonzispina/go-solid/tree/master/O)
 - [Liskov's Substitution (LSP)](https://github.com/gonzispina/go-solid/tree/master/L)
 - [Interface Segregation (ISP)](https://github.com/gonzispina/go-solid/tree/master/I)
 - [Dependency Inversion (DIP)](https://github.com/gonzispina/go-solid/tree/master/D)
 
 _NOTE: keep in mind that the examples of every principle are designed to isolate and expose the key idea of each principle. None of 
the examples combine all of them. Of course, if we applied every principle in each example the code would be much better, but that's not
 the point._

## References

- _"Agile Software Development, Principles, Patterns, and Practices"_. Robert C. Martin, 15 Oct. 2002. (p85-145) 
- _"Solid Go design"_ -- Dave Cheney, 20 Aug. 2016. ([Youtube](https://www.youtube.com/watch?v=zzAdEt3xZ1M)) 

## Further Reading

- Liskov, Barbara. Data Abstraction and Hierarchy. SIGPLAN Notices, 23,5 (May 1988).
- Meyer, Bertrand. Object-Oriented Software Construction, 2d ed. Upper Saddle River, NJ: Prentice Hall, 1997.



