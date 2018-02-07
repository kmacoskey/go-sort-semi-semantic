Sort a list of versions the same way the concourse s3 resource does.

Reads a list of versions from stdin that are \n delimited and prints the list with the same semi-semantic versioning applied.

Example Usage:
```
aws s3 ls s3://madlib-concourse/release/gpdb5/ | tr -s ' ' | cut -d' ' -f 4 | ./go-sort-semi-semantic
```
