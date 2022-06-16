volume = 7


def change_volume(request):
    global volume
    
    if request == 'Skru op!' and volume < 10:
        volume += 1

    if request == 'Skru ned!' and volume > 0:
        volume -= 1


number_of_requests = int(input())

for i in range(number_of_requests):
    request = input()
    change_volume(request)


print(volume)
