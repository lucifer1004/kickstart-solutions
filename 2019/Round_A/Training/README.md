# Important notes

## Solutions

|          Language          | Test set 1 | Test set 2 |
|:--------------------------:|:----------:|:----------:|
|  [Python3](training.py3)   |    PASS    |  **TLE**   |
| [Pypy2](training.pypy2.py) |    PASS    |  **TLE**   |
|   [GoLang](training.go)    |    PASS    |    PASS    |

## About the problem

An important step is to calculate accumulative sum of the skill array in advance, so that calculation can be saved when calculating partial sum.

## Language specific

### GoLang

`Scanner` has a default buffer size of 64K, which is insufficient for Test set 2, so we need to set a custom buffer size.