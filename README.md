# Prefix Writer

Implements simple writer, that will prefix each incoming line. It's guaranteed,
that prefix will be prepended only once per each line, no matter how many calls
to Write is done for writing that line.

# Behavior
```
New           Returns Writer With Prefix
Writer  Write Nothing At Empty Data
Writer  Write First Byte With Prefix
Writer  Not Inserts Prefix On Two Writes Without Newline
Writer  Adds Prefix On Each New Line
Writer  Inserts Prefix When Newline Comes First
```

Generated with [loverage](https://github.com/kovetskiy/loverage).

# Documentation

See reference at [godoc.org](http://godoc.org/github.com/reconquest/prefixwriter-go).
