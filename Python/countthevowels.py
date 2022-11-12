text = input()

total_vowels = 0
vowels = ['A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u']
for letter in text:
    if letter in vowels:
        total_vowels += 1

print(total_vowels)

