# CS255 Examples

Example problems for the University of Warwick's [CS255](https://warwick.ac.uk/fac/sci/dcs/teaching/modules/cs255/) coursework

## Solutions

Every problem here has at least one task 2/3 solution. Problems <=65 have at least one task 1 solution.

Some problems may provide their solution underneath the module list - you might need to adapt the code that processes example problems to deal with this:

If you haven't messed with the coursework skeleton much, you should be able to add
```python
if len(moduleList) == 25:
  break
```
after line 37 of ReaderWriter.py to handle this.

## Tests

A test to check the problems are vaguely valid is provided in `checkformat.go`. Simply run:

```bash
go run checkformat.go
```