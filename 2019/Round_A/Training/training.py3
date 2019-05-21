import numpy as np

# Read T (Number of test cases)
t = int(input())

# Deal with each test case
for i in range(t):
    # Read N (Number of students) & P (Number of players)
    line = input()
    [n, p] = [int(s) for s in line.split(' ')]

    # Read skill of each student
    line = input()
    s = np.array([int(s) for s in line.split(' ')])

    # Sort skill in descending order
    s.sort()
    s = np.flip(s)

    # Calculate accumulative sum
    skill_sum = np.append(np.zeros(1), np.cumsum(s))

    # Set original value for the answer
    min_hours = -1

    for j in range(n - p + 1):
        # Calculate current skill
        skill_current = skill_sum[j + p] - skill_sum[j]

        # Calculate target skill
        skill_target = p * s[j]

        # Calculate hours required
        hours = int(skill_target - skill_current)

        # Update answer with new minimal
        if hours < min_hours or min_hours < 0:
            min_hours = hours

        # Early break
        if min_hours == 0:
            break

    # Print the answer
    print("Case #{}: {}".format(i + 1, min_hours))
