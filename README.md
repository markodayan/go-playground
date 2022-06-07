# Notes

Run tests with:

```
go test -v
```

### Insertion Sort

- Always start at index of 2nd element (first `i` loop)
- Inner Loop starts at `i - 1` and decrements `j` until and including 0
- Simply swap if `j+1`-indexed value is less than `j`-indexed array value

```
                             _
    i (starting)              |
    |                         |  i = 1, j loops from index 0 (once)
 j  |                         |  value 5 is swapped with value 2 (index 0 <-> index 1)
 |  |                         |
[5, 2, 4, 6, 1, 3]           _|

                             _
       i = 2                  |
       |                      |  i = 2, j loops from 1 to 0 (2 times)
 <- j  |                      |  value 4 is swapped with value 5 (index 1 <-> index 2)
    |  |                      |
[2, 5, 4, 6, 1, 3]           _|


                             _
          i = 3               |
          |                   |  i = 3, j loops from 2 to 0 (3 times)
    <- j  |                   |  no swaps needed
       |  |                   |
[2, 4, 5, 6, 1, 3]           _|

                             _
             i = 4            |
             |                |  i = 4, j loops from 3 to 0 (4 times)
       <- j  |                |  4 swaps occur, sinking the value 1 to index 0
          |  |                |
[2, 4, 5, 6, 1, 3]           _|

                             _
                i = 5         |
                |             |  i = 5, j loops from 4 to 0 (5 times)
          <- j  |             |  3 swaps occur, sinking the value 3 to index 2
             |  |             |
[2, 4, 5, 6, 1, 3]           _|


Final Sorted Array -> [1,2,3,4,5,6]

```
