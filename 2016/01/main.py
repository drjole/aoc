x, y = 0, 0
current_direction = 0
visited = set()
visited_twice = None

with open("input.txt") as io:
    for instruction in io.read().split(", "):
        direction = instruction[0]
        if direction == "R":
            current_direction += 1
        elif direction == "L":
            current_direction -= 1
        current_direction %= 4

        steps = int(instruction[1:])
        for i in range(steps):
            if current_direction == 0:
                y += 1
            elif current_direction == 1:
                x += 1
            elif current_direction == 2:
                y -= 1
            elif current_direction == 3:
                x -= 1
            if visited_twice is None and (x, y) in visited:
                visited_twice = (x, y)
            visited.add((x, y))

print(abs(x) + abs(y))
print(abs(visited_twice[0]) + abs(visited_twice[1]))
