min = -10.0
max = 10.0
step = 0.0001
places = 3

current = min
while current <= max:
    print('{0} {1}'.format(current, round(current, places)))
    current = current + step