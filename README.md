#   Exercise: Design a Baggage Concierge

First part initial design of store and retrieve (30mins) Second part optimization of storage and/or retrieval (30 mins)

### Prompt
In this exercise we’ll be implementing an airport baggage concierge. A customer at an airport who has a long layover comes to this concierge to leave their bags and receives a ticket. Later the customer retrieves the bag with the same ticket. We can start with a single line and a single customer at a time.

There are three main constraints:
-   Bags have two sizes: carry-ons (small) and checked bags (large)
-   Storage bins store either 2 carry-ons or 1 checked bag
-   Bins have a finite length: 100

Please implement the interface to fulfill the requirements above.

```go
type BaggageConcierge interface {
    Store(bag Bag) Ticket
    Retrieve(ticket Ticket) Bag
}
```


### First part initial design of store and retrieve (30mins)
Guiding questions:
-   How would you design the Bag object?
-   How would Bag be related to Ticket?
-   How do you figure out storage area?
-   How do you handle error cases, i.e. it’s full, or full for only one bag?

###  Second part optimization of storage and/or retrieval (30 mins)
Guiding questions:
-   How would you test your algorithm
-   If we scale the bins to non-trivial size, how does your algorithm performance change over time?
-   Can you discuss better data structures for these operations?

Optional: What if there were four lines operating at the same time, how would you change your code to handle this?



## * How to Test [ Unit Test ] 

```go
$   git clone https://github.com/otyang/concierge
$   cd concierge/
$   go test -v
$   go test -bench .
```

- 3rd line runs unit test
- 4th runs benchmark test


## * Can you discuss better data structures for these operations?

Here are some better data structure choices for the given operations in the Baggage Concierge:

1. **Queue or Priority Queue:** For storing bags, the FIFO - (First In First Out) or a priority queue structure be used as it is efficient. It gives room for a better organization of the bags based on their size, availability of storage bins,  constant-time insertion and removal of bag operations, all this are beneficial for managing the bag storage.

2. **Heap:** If the orders in which the bags are stored are important, a min-heap or max-heap can be used. The heap will be based on "bagSize" - with the smallest bags first. This enables efficient access to the smallest or largest bags and ensures proper organization in the storage bins.

3. **Map:** A map can be used to keep track of the bag and its associated ticket. The map can store the ticket ID as the key and the bag as the value. This allows for efficient retrieval of bags based on their ticket ID with constant time complexity (`O(1)`).


4. **Segment Tree or Fenwick Tree:** This datastructure can be used in case where storage bins need be arranged on their capacity or availability giving room for more flexibility.


5. **Linked List:** A linked list data structure could be used to represent the storage bins, where each node of the linked list corresponds to a storage bin. This can provide constant-time insertion and removal of bags in the storage bins and dynamic resizing when the number of bins changes.

###   Summary:
The choice of much better data structures for these kind of operations are to a large extent dependent on salient factors such as 
-   expected number of bags,
-   frequency of bag storage and retrieval, 
-   dynamic resizing or organization,
-   concurrent access and synchronization (in cases where user need interact with bags simultaneously) etc

The right data structures and algorithms can significantly impact the performance, scalability, and maintainability of the Baggage Concierge system as it evolves over time.


