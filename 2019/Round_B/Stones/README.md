# Important notes

## Solutions

|      Language       | Test set 1 | Test set 2 |
| :-----------------: | :--------: | :--------: |
| [GoLang](stones.go) |    PASS    |    PASS    |

## Language specific

### GoLang

The GoLang version running on `KickStart` (v1.7) does not support
`sort.SliceStable()`, so a custom type is required so that `sort.Sort()` can be
used.
