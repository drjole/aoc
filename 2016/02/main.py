def first(instructions):
    code = ''
    x, y = 1, 1
    for line in instructions:
        for direction in line:
            if direction == "U" and y > 0:
                y -= 1
            if direction == "R" and x < 2:
                x += 1
            if direction == "D" and y < 2:
                y += 1
            if direction == "L" and x > 0:
                x -= 1
        code += "123456789"[3 * y + x]
    return code


def second(instructions):
    code = ''
    x, y = 0, 2
    invalid = {
        (0, 0),
        (0, 1),
        (0, 3),
        (0, 4),
        (1, 0),
        (1, 4),
        (3, 0),
        (3, 4),
        (4, 0),
        (4, 1),
        (4, 3),
        (4, 4),
    }
    for line in instructions:
        for direction in line:
            if direction == "U" and y > 0 and (x, y - 1) not in invalid:
                y -= 1
            if direction == "R" and x < 4 and (x + 1, y) not in invalid:
                x += 1
            if direction == "D" and y < 4 and (x, y + 1) not in invalid:
                y += 1
            if direction == "L" and x > 0 and (x - 1, y) not in invalid:
                x -= 1
        code += "__1___234_56789_ABC___D__"[5 * y + x]
    return code


with open("input.txt") as io:
    lines = list(io.readlines())

print(first(lines))
print(second(lines))
