# PyTesting @ Version 0.1

## Requires:

 - Python 3
 - termcolor

## Installation:
```sh
  pip install termcolor
  pip install cherrypy
```

### Makefile rules

```make test-all``` :  Run the file testing.py for all group
```make test grp=group_name number=1``` : Run test 1 on group_name target
```make webapp``` : Run the CherryPy serveur at http://127.0.0.1:8080

### Files description

```config.json``` : Paths configurations
```testing.py``` : Main script
```TestReport.py``` : Class TestReport
```results/group_name/rapport.html``` : Tests report html format
```tests/test_name/``` : Script test and informations
```tests/_tests_order.txt``` : Order of tests, indexed by lines numbers.
```web_app/``` : Web application directory (View stats and historic)
