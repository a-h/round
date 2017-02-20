# Integration Tests

This project includes a Python test file `create.py` which can be used to create inputs for the `main.go` test harness.

## Testing Process

```
python create.py > python2.txt
python3 create.py > python3.txt
```

```
go run main.go -in python2.txt -round AwayFromZero -places 2
go run main.go -in python3.txt -round ToEven -places 2
```

The `create.py` file can be modified to produce different sizes of test suites, e.g. 4dp.

```
min = -10.0
max = 10.0
step = 0.0001
places = 3

current = min
while current <= max:
    print('{0} {1}'.format(current, round(current, places)))
    current = current + step
```

For 3dp, then the test suite needs to change.

```
go run main.go -in python2.txt -round AwayFromZero -places 3
go run main.go -in python3.txt -round ToEven -places 3
```