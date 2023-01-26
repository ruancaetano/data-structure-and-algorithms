# Hash Table

Hash Table, Map, HashMap, Dictionary or Associative are all the names of the same data structure. It is one of the most commonly used data structures.

The Hash table data structure stores elements in key-value pairs where

- Key = used for indexing the values
- Value = data that are associated with keys.

## Hashing

In a hash table a new index is processed using the keys. And, the element corresponding to that key is stored in the index. This process is called hashing.


## Hash Collision
When the hash function generates the same index for multiple keys, there will be a conflict (what value to be stored in that index). This is called a hash collision.

We can resolve the hash collision using resolution by:

- Chaining: Uses one of a linked list of pointers at each index, so whenever there is a collision, the item in question is added to the end of the linked list at that index

- Open Addressing: It only uses an array, whenever a collision happens, it searches the next index for the given key using the following Strategies Linear/Quadratic Probing or Double Hashing