#   Exercise: Design a Baggage Concierge

First part initial design of store and retrieve (30mins) Second part optimization of storage and/or retrieval (30 mins)

# Prompt

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


## First part initial design of store and retrieve (30mins)
Guiding questions:
-   How would you design the Bag object?
-   How would Bag be related to Ticket?
-   How do you figure out storage area?
-   How do you handle error cases, i.e. it’s full, or full for only one bag?

##  Second part optimization of storage and/or retrieval (30 mins)
Guiding questions:
-   How would you test your algorithm
-   If we scale the bins to non-trivial size, how does your algorithm performance change over time?
-   Can you discuss better data structures for these operations?

Optional: What if there were four lines operating at the same time, how would you change your code to handle this?