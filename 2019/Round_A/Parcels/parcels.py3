import numpy as np


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
    map = []
    for j in range(r):
        line = input()
        map_row = [int(s) for s in line]
        map.append(map_row)

    map = np.array(map)
    original = np.zeros(4 * r * c).reshape(r * c, 4)

    for j in range(r):
        for k in range(c):
            original[j * c + k] = np.array(
                [j, k, map[j, k], 0 if map[j, k] == 1 else -1])

    for j in range(r * c):
        if original[j, 2] == 1:
            for k in range(r * c):
                dist = manhattan(
                    original[j, 0], original[j, 1], original[k, 0], original[k, 1])
                if dist < original[k, 3] or original[k, 3] < 0:
                    original[k, 3] = dist

    ans = np.max(original[:, 3])

    if ans > 0:
        for j in range(r * c):
            if original[j, 2] == 0:
                solution = original.copy()
                for k in range(r * c):
                    dist = manhattan(
                        original[j, 0], original[j, 1], original[k, 0], original[k, 1])
                    if dist < solution[k, 3]:
                        solution[k, 3] = dist
                min_delivery = np.max(solution[:, 3])
                if min_delivery < ans:
                    ans = min_delivery

    # Print the answer
    print("Case #{}: {}".format(i + 1, int(ans)))
