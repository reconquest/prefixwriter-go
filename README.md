# Prefix Writer

Implements simple writer, that will prefix each incoming line. It's guaranteed,
that prefix will be prepended only once per each line, no matter how many calls
to Write is done for writing that line.

# Behavior
```
NewPrefixWriter  Returns Writer With Prefix
PrefixWriter     Write Nothing At Empty Data
PrefixWriter     Write First Byte With Prefix
PrefixWriter     Not Inserts Prefix On Two Writes Without Newline
PrefixWriter     Adds Prefix On Each New Line
PrefixWriter     Inserts Prefix When Newline Comes First
```

Generated with [loverage](https://github.com/kovetskiy/loverage).

# Documentation

See reference at [godoc.org](http://godoc.org/github.com/reconquest/go-prefixwriter).
