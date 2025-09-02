import collections
from typing import List


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        # we are performing a bfs since we are looking for adjacents
        # base case - when no grid we return 0 bc no islands
        if not grid:
            return 0
        
        rows,cols = len(grid), len(grid[0])
        visited = set()
        islands = 0

        def bfs(r,c):
            # need a queue to track what node to visit next
            q = collections.deque()

            # add source vertex
            visited.add((r,c))
            q.append((r,c))

            while q:
                # pop current node
                row,column = q.popleft() # this could be changed to pop() to make it dfs

                # check adjacent nodes
                directions = [[1,0], [-1,0], [0,1], [0,-1]]
                for dr,dc in directions:
                    r,c = row + dr, column + dc

                    # check if range of the grid, if the new spot is land, and to make sure we haven't visited the node
                    if r in range(rows) and c in range(cols) and grid[r][c] == "1" and (r,c) not in visited:
                        # add to q to continue search
                        q.append((r,c))
                        # add to visited so we dont accidentally visit again
                        visited.add((r,c))


        # go through the grid to find a 1, make sure to check if that one is not an island we found before
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == "1" and (r,c) not in visited:
                    bfs(r,c)
                    islands += 1

        return islands