# PyTesting @ Version 0.1

## Requires:

 - Python 3
 - termcolor

## Installation:
```sh
  pip install termcolor
```

### Makefile rules

```make test``` :  Run the file testing.py for all group
```make test grp=group_name number=1``` : Run test 1 on group_name target

### Files description

```config.json``` : Paths configurations
```testing.py``` : Main script
```TestReport.py``` : Class TestReport
```results/group_name/rapport.html``` : Tests report html format
```tests/test_name/``` : Script test and informations
```tests/_tests_order.txt``` : Order of tests, indexed by lines numbers.
