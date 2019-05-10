import functools

# Read T (Number of test cases)
t = int(input())

# Deal with each test case
for i in range(t):
    # Read N (Number of students) & P (Number of players)
    line = raw_input()
    [n, p] = [int(s) for s in line.split(' ')]

    # Read skill of each student
    line = raw_input()
    s = [int(s) for s in line.split(' ')]

    # Sort skill in descending order
    s.sort(reverse=True)

    # Calculate accumulative sum
    skill_sum = functools.reduce(
        lambda x, y: x + [x[-1] + y] if len(x) > 0 else [y], s, [])
    skill_sum = [0] + skill_sum

    # Set original value for the answer
    min_hours = -1

    for j in range(n - p + 1):
        # Calculate current skill
        skill_current = skill_sum[j + p] - skill_sum[j]

        # Calculate target skill
        skill_target = p * s[j]

        # Calculate hours required
        hours = skill_target - skill_current

        # Update answer with new minimal
        if hours < min_hours or min_hours < 0:
            min_hours = hours

    # Print the answer
    print("Case #{}: {}".format(i + 1, min_hours))
