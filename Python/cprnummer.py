cpr_number = input()

i = 0
sum_of_digits = 0
corresponding_values = [4, 3, 2, 7, 6, 5, 4, 3, 2, 1]
for digit in cpr_number:
    if digit == '-':
        continue

    sum_of_digits += int(digit) * corresponding_values[i]
    i += 1

print(1) if sum_of_digits % 11 == 0 else print(0)
