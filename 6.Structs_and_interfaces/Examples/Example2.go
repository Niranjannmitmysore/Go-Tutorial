//  Multiple inheritance
/* 1. By embedding a struct into another you have a mechanism similar to 
    multiple inheritance with non-virtual members.
2. Let's call "base" the struct embedded and "derived" the struct doing the embedding.
3. After embedding, the base fields and methods are directly available in the derived struct. 
    Internally a hidden field is created, named as the base-struct-name.
4. Base fields and methods are directly available as if they were declared in the derived struct, 
    but base fields and methods can be "shadowed" 
5. Shadowing means defining another field or method with the same name (and signature) of a 
    base field or method.
6. Once shadowed, the only way to access the base member is to use the hidden field named as the 
    base-struct-name. */

/*
Golang | Classic OOP
struct  | class with fields, only non-virtual methods
interface | class without fields, only virtual methods
embedding | multiple inheritance
*/



