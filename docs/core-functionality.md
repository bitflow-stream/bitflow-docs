# Core functionality

This section describes the cores features of bitflow.

## Batch 

Batch is one of the few keywords in the bitflow script language. Batch is used as a simple processing steps inside a script and must be combined with a subpipeline like **`"input.csv" -> batch(sample-window-size=5){ avg() } -> out.csv`**

(TODO) Required batch modes:
- flush after input closed (entire file)
- separation tag (with optional timeout)
- flush after fixed number of samples
- flush after fixed period of time

## Fork

TODO: describe fork semantics (distributors, etc)