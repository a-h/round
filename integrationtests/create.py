from decimal import *

min = Decimal('-10.0')
max = Decimal('10.0')
step = Decimal('0.0001')
places = 3

current = min
while current <= max:
    print('{0} {1}'.format(current, round(current, places)))
    current = current + step