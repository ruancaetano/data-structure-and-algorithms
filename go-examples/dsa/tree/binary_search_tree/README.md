# Binary Search Tree

Binary Search Tree is a node-based binary tree data structure which has the following properties:

The left subtree of a node contains only nodes with keys lesser than the node’s key.
The right subtree of a node contains only nodes with keys greater than the node’s key.
The left and right subtree each must also be a binary search tree.


## Insertion order notes

The first node you add to the BST (binary search tree) will always be the root and cannot be changed after initialization. This fact alone can explain pretty obviously why the ordering matters. Given N elements, you have already N possible (and more importantly, different) BST structures.

Example

```
nums = [1,5,10]


TREE1

1
 \
  5
   \
   10


TREE2

   5
 /   \
1    10

TREE3

     10
    /
   5
  /
1
```


Ordering will affect the balance of a tree (TREE2 is balanced, TREE1 and TREE2 are not)
balancing affects efficiency for BST operations, e.g. traversal/search:

perfectly balanced O(log_n) time
unbalanced: O(n) time (= worst case).