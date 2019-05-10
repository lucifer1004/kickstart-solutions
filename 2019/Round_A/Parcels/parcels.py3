import numpy as np

# Constants
dire = [[0, -1], [-1, 0], [0, 1], [1, 0]]


def manhattan(r1, c1, r2, c2):
    """
    Calculates Manhattan distance of given (r1, c1) and (r2, c2)
    :param r1: row number of the first square
    :param c1: column number of the first square
    :param r2: row number of the second square
    :param c2: column number of the second square
    :return: number. The Manhattan distance from (r1, c1) to (r2, c2)
    """
    return abs(r1 - r2) + abs(c1 - c2)


# Read T (Number of test cases)
t = int(input())

# Deal with each test case
for i in range(t):

    # Read R (Number of rows) & C (Number of columns)
    line = input()
    [r, c] = [int(s) for s in line.split(' ')]

    # Read the grid
    grid = []
    for j in range(r):
        line = input()
        grid_row = [int(s) for s in line]
        grid.append(grid_row)

    grid = np.array(grid)
    dist = np.zeros((r, c))
    visited = np.zeros((r, c))
    queue = []

    # BFS to get current delivery distance of every square
    for j in range(r):
        for k in range(c):
            if grid[j, k] == 1:
                dist[j, k] = 0
                visited[j, k] = 1
                queue.append([j, k, 0])

    head = 0
    tail = len(queue)

    while head < tail:
        head_r = queue[head][0]
        head_c = queue[head][1]
        dist[head_r, head_c] = queue[head][2]
        for j in range(len(dire)):
            target_r = head_r + dire[j][0]
            target_c = head_c + dire[j][1]
            if 0 <= target_r < r and 0 <= target_c < c:
                if visited[target_r, target_c] == 0:
                    queue.append([target_r, target_c, queue[head][2] + 1])
                    visited[target_r, target_c] = 1
                    dist[head_r, head_c] = queue[head][2]
                    tail += 1
        head += 1

    ans = np.max(dist)

    if ans > 0:
        for j in range(r * c):
            if queue[j][2] > 0:
                new_office_r = queue[j][0]
                new_office_c = queue[j][1]
                solution = dist.copy()
                for k in range(r * c):
                    curr_r = queue[k][0]
                    curr_c = queue[k][1]
                    new_dist = manhattan(
                        new_office_r, new_office_c, curr_r, curr_c)
                    if new_dist < solution[curr_r, curr_c]:
                        solution[curr_r, curr_c] = new_dist
                min_delivery = np.max(solution)
                if min_delivery < ans:
                    ans = min_delivery

    # Print the answer
    print("Case #{}: {}".format(i + 1, int(ans)))
