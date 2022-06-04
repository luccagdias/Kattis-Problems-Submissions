def get_rotated_half(half): 
    rotated_half = ''
    rotation_value = calculate_rotation_value(half)

    for character in half:
        result = (ord(character) - 65 + rotation_value) % 26 + 65
        rotated_half += chr(result)

    return rotated_half


def calculate_rotation_value(half):
    rotation_value = 0
    for character in half:
        rotation_value += ord(character) % 65

    return rotation_value


def merge_halves(rotated_first_half, rotated_second_half):
    result = ''
    for i in range(0, len(rotated_first_half)):
        result += chr((ord(rotated_first_half[i]) - 65 + ord(rotated_second_half[i]) % 65) % 26 + 65)

    return result

drm_message = input()
half =  int(len(drm_message) / 2)

rotated_first_half = get_rotated_half(drm_message[:half])
rotated_second_half = get_rotated_half(drm_message[half:])

result = merge_halves(rotated_first_half, rotated_second_half)
print(result)

