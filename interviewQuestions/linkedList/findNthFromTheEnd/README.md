# Challenge

## findNthFromTheEnd.go
Given a singly-linked list of n-elements, find the nth emement from the end.

So the problem is that once you've passed any given point, you cant go back because you've lost the address.
## Data
the findNthFromTheEnd.dat file is a csv on one row. 
The parser just creates a singly linked list 
eg:
a,s,d,f,g 
creates a list like
a->s->d->g

## Solution
The accepted solution uses a look head pointer that moves ahead n spaces, then the actual pointer is set to the head of the linked list. Then you move each pointer forward one node at a time until the look ahead pointer reaches the end of the linked list.

## Alternate solution 1
copy the linked list to an array and use array math to find the nth element.
