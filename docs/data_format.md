# Bitflow Data Format

The Bitflow software components define multiple common data formats for exchanging streams of samples.
The data formats defined below can be stored in files or transmitted directly over the network.
In both cases, the receiver (or reader) usually automatically determines what data format is used (see below on details on how to determine the format).
In the TCP communication case, the listening socket can act both as a data source and a data sink, so the side that opens a connection has to know in advance, if a listening port will receive or transmit data.
Once a connection is established, the sending side simply starts transmitting the data stream in the same format that would be stored in a file.
In case of errors or if the data stream stops, the connection is closed.
There is no further error handling.

A data stream always starts with a header, which is followed by zero or more samples.
The header defines the names of the metrics, while the sample contains the metric values.

One sample contains the following information:
* A timestamp
* An arbitrary number of key-value string pairs (tags)
* A number of metric values (double precision) that matches the number of metric names defined in the header.

The data format defines, how the header and the samples are encoded as streams of bytes.
If the metrics provided by a data source change, an ongoing TCP connection must be closed and a new connection must be opened, which will send a new header defining the new metrics.
The same must be done when storing data in files: each file can only contain a single header.

**Note:** Some implementations of Bitflow (for example go-bitflow) support changing headers both in files and in TCP connections.
This should not be done to stay compatible with other implementations (bitflow4j and bitflow.py).
Changing headers are implemented by simply transmitting the byte stream of a new header instead of a sample.
The header data formats are defined in a way that the header can be distinguished from sample data.
Writing a CSV file with multiple headers breaks the semantics of regular CSV files.

## CSV data format

The CSV data format follows the basic CSV syntax. The header is followed by a number of samples, one sample per line:
```
time,tags,cpu,disk-io/all/io,disk-io/all/ioBytes,disk-io/all/ioTime
2017-11-09 13:51:09.877210495,experiment=cpu host=wally133,0,0,0,0
2017-11-09 13:51:10.377433859,experiment=cpu host=wally133,54.99999999927241,11.654148188577683,859247.4519964771,0,11.653989825577415
```

Comma (`,`) is used for separating fields and newline characters separate the header line and the individual samples.
The number of header fields must match the number of fields in each sample.
The first column name must always be `time`, the second column is always `tags`. This is used to identify the data format as CSV.
The `time` column contains the timestamp of each sample in the format `YYYY-MM-DD HH:mm:SS.sssssssss`.
The `tags` column contains the tags as string key-value pairs separated by `=` and single space characters.
There can be an arbitrary number of tags, including zero, in which case the field is empty.
There can also be an arbitray number of metrics. When there are no metrics, the CSV header is simply `time,tags`.

## Binary data format

A header in the binary format is encoded by writing the header fields, separated by newline characters.
The first header field is for the timestamp and is called `timB`. This is different from the CSV format, but has the same length to make the format discovery more efficient.
The second header field is `tags`, which is then followed by the metric names.
After all fields are written, an empty line signals the start of the first sample.
In other words, the end of the header is encoded as a double newline character (`\n\n`).

A binary sample starts with a capital `X` byte. This is to distinguish it from a repeated header.
After that, the timestamp is transmitted as a big-endian 8-byte unsigned integer value of the [Unix timestamp](https://en.wikipedia.org/wiki/Unix_time).
After the timestamp, the tags are transmitted in the same format as in the CSV format, concluded by a single newline character.
Following the tags, the actual metric data is encoded as big-endian IEEE 754 double-precision values.
The number of metric values must match the number of metric names defined in the header. The sample is not terminated by any additional byte sequence!
After the last byte of the last metric, an `X` byte indicates a new sample, while a `timB` byte sequence indicates an updated header.
Any other byte sequence would violate the marshalling format and should lead to a closed TCP connection or signal an invalid data file.

**Note:** Since some byte sequences are special in the two data formats, the tags must not contain any of the following characters: `, \n =`, in addition to the white space character. By convention, such characters are replaced by underscores.
