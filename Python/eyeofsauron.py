tower = input()

affixes = tower.split('()')
prefix, suffix = affixes[0], affixes[1]

print('correct' if prefix == suffix else 'fix')
