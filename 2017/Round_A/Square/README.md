# Important notes

## Solutions

|      Language       | Test set 1 | Test set 2 |
| :-----------------: | :--------: | :--------: |
| [GoLang](square.go) |    PASS    |    PASS    |

## About the problem

- It is easy to calculate regular rectangles (aligning along axes), so it is
  important to figure out that for a regular rectangle whose width is `a`, there
  are exactly `a-1` irregular rectangles that can be placed inside the regular
  rectangle.
- Mul results can easily exceed max int, so every time a mul is done, a modulo
  should follow.
- When there is subtraction, modulo results can be negative, this should be
  handled before output.
- `r*c` and `r+c` can be calculated and stored in advance.
