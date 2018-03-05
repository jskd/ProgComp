PY=python3

test-all:
	${PY} src/testing.py
test:
	${PY}  src/testing.py grp=${grp} number=${number}

webapp:
	${PY} web_app/Run.py

info:
	find ./ -type f -iname \*.py -o -type f -iname \*.css -o -type f -iname \*.html -o -type f -iname \*.json | xargs wc -l
