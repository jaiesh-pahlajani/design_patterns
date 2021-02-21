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
