# design_patterns

**Solid Principles**

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