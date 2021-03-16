# design_patterns

**Solid Principles**
https://docs.google.com/document/d/1Dr63HIAsFJnxyW_89IKBQhH5-t96n3nV-rRvZq9ZGPw/edit

**Builder**
1. A builder is a separate component used for building an object.
2. To make builder fluent, return the receiver allows chaining.
3. Different facets of an object can be built with different builders working in tandem via a common struct.

**Factory**
1. A factory function (a.k.a. constructor) is a helper function for making struct instances.
2. A factory is any entity that can take care of object creation
3. Can be a function or a dedicated struct

**Prototype**
1. Complicated objects aren't designed form scratch
2. An existing design is a prototype.
3. We make a copy of the prototype and customize it

**Singleton**
1. For some components it only makes sense to have one instance in the system i.e. Database repository. 
2. Construction call is expensive , only do it once.
3. Need to take care of lazy instantiation
4. Adhere to DIP: depend on interfaces, not concrete types. 

**Adapter**
1. Determine the API you have and the API you need
2. Create a component which aggregates(has a pointer to) the adaptee
3. Intermediate representations can pile up: use caching and other optimizations

**Bridge**
1. Decouple abstraction from implementation
2. Both can exist as hierarchies
3. A stronger form of encapsulation

**Composite**
1. Objects use other objects fields/methods through embedding/composition
2. Composition lets us make compound objects
    - A mathematical expression composed of simple expressions 
    - A shape group made of several different shapes
3. Composite design pattern is used to treat both single (scalar) and composite objects uniformly
    - Foo and []Foo have common APIs
4. A mechanism for treating individual objects and composition of objects in a uniform manner.
5. Some composed and singular objects need similar/identical behaviours
6. Composed design pattern lets us treat both types of objects uniformly
7. Iteration supported with iterator design pattern

**Decorator**
1. Want to augment an object with additional functionality
2. Do not want to rewrite or alter existing code(OCP)
3. Want to keep new functionality separate(SRP)
4. Need to be able to interact with existing structures
5. Inshort facilitates the addition of behaviours to individual objects through embedding

**Facade**
1. Provides a simple easy to use interface over a large sophisticated body of code
2. Provides a simplified API over a set of components
3. Optional to expose internal details through facade
4. May allows users to escalate to use more complex APIs if they need to. 