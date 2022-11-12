password1 = input()
password2 = input()

distinct_sequences = 1

for i in range(4):
    if password1[i] == password2[i]:
        distinct_sequences *= 1
    else:
        distinct_sequences *= 2

print(distinct_sequences)
