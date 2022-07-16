# Graphs

- Nodes = `Vertices`
- Links = `Edges`

Graphs are abstract data types, and so there are many different ways to implement it:

- Directed Graph
- Undirected Graph
- Directed Cyclic Graph
- Directed Weighted Graph
- Trees are graphs with certain restrictions
- Directional, cyclic, unweighted graph (an example of graph with several properties)

<u>There are 2 common ways to represent graphs:</u>

1. <b>Adjacency List</b>: Graph is expressed as a list of vertices and each vertice has a list holding the neigbouring vertices.
2. <b>Adjacency Matrix</b>: Graph is expressed with a 2D array (rows, columns -> from, to vertices)

### Adjacency List vs Adjacency Matrix

- If the <u>graph is sparse</u>, then an adjacency list will let us save space.
- If <u>speed of edge lookup</u> is important, then the adjacency matrix is much faster (constant time `O(1)`). Adjacency list would be O(n).
- If want to <u>add vertex to graph</u>, it is easy in an adjacency list `O(1)`. `O(n^2)` for adjacency matrix.
- If want to <u>remove vertex from graph</u>, it is easy in an adjacency list `up to O(E)` (E = total number of edges). `O(n^2)` for adjacency matrix.
- <u>Adding edge</u> -> quick in adjacency list `O(1)`. Also quick in adjacency matrix `O(1)`.
- <u>Removing edge</u> -> quicker in adjacency matrix `O(1)`. Will take `O(V)` time in adjacency list (because we need to traverse to find the edge).
