# priority queue

I ran into the Priority Queue concept when trying to solve the LeetCode problem [Find K-closet elements](https://leetcode.com/problems/find-k-closest-elements)

A priority queue is an **Abstract Data Type (ADT)** that operates similar to a normal queue except that each element has a certain priority. The priority of the elements in priority queue determine the order in which elements are removed from the queue.

Priority queue only supports **comparable data**, meaning the data inserted into the priority queue must be able to be ordered in some way either from least to greatest or greatest to the least. This is so that we are able to assign relative priorities to each element.

Heap forms the canonical underlying data structurefor priority queue, but heaps are not priority queues (ADT), meaning we can implement priority queue using any data structure. The priority queue in this code is implemented as Binary Heap.

Priority queues most used in the certain problems such as finding the "next best" or "next worst" element, or algorithms such as Best First Search, Minimum Spanning Tree, Dijkstra's Shorted Path...

### references
Priority Queue Explained: https://www.youtube.com/watch?v=wptevk0bshY