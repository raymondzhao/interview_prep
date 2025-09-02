from typing import List

class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # create adjacency list to represent prereq map to dfs
        prereqList = {i:[] for i in range(numCourses)}
        for course,prereq in prerequisites:
            prereqList[course].append(prereq)
        
        # create visited set to make sure we dont have cycles
        visited = set()

        # create recursive func
        def dfs(course):
            # base case - if we are visiting a course that we have visited, we have a cycle
            if course in visited:
                return False

            # if the course has no prereqs then we can complete it which should be true
            if prereqList[course] == []:
                return True

            visited.add(course)
            for prereq in prereqList[course]:
                # if we cannot complete the prereq course, we can't complete the current course
                if not dfs(prereq):
                    return False

            # if we can complete this loop, then it means we can complete this course.
            # since we have stopped visiting the course, we can remove it from the set
            visited.remove(course)

            # since we know it can be visited, we can set the prereq map to empty list
            # dfs wont need to repeat work since we know it can be completed
            prereqList[course] = []
            return True
        
        # iterate through every course bc what if we have many different graphs that aren't connected
        for course in range(numCourses):
            if not dfs(course):
                return False
            
        return True