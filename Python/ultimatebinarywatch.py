hour = input()

binaries = []

for s in hour:
    n = int(s)
    n_binary = int(bin(n)[2:])

    binaries.append(f'{n_binary:04d}')


for i in range(4):
    for j in range(4):
        binary = binaries[j]

        if binary[i] == '0':
            print('.', end='')
        else:
            print('*', end='')

        if j == 1:
            print('   ', end='')

        if j != 1 and j !=3:
            print(' ', end='')

    print()
