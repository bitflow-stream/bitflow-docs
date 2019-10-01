
# Naming Conventions in Bitflow

## Glossary of General Terms

The following terms are the core of the Bitflow execution and programming model.
They should be used throughout all implementations.
The purpose of this glossary is to increase consistency between the implementations and avoid confusion.
The names of classes and other types in the implementations should follow these terms as much as possible.
Of course the implementations can have class hierarchies or helper types with different names.

- **Metric** (or value): Named floating point value
- **Sample**: Primary element of a data stream. Consists of a timestamp, a list of metrics and a string dictionary of tags.
- **Header**: A header contains meta information that is shared by a collection of samples. The most important information is a list of strings that contains the metric names.
- **Processing step** (or step): Entity that receives a stream of samples for arbitrary processing. It can forward arbitrary samples to its successor (also called subsequent processing step, or simply output).
    - Similar terms that should *NOT* be used: algorithm, pipeline step, processor, filter
- **Data source** (or source): Entity that generates samples, typically by reading and parsing some external data like a file or network stream
    - Similar terms that should *NOT* be used: input, data input, collector
- **Data sink** (or sink): Entity that receives samples and outputs them externally, like to a file or network stream.
    - Important: a data sink behaves semantically like a no-op processing step. After outputting a sample, it must be forwarded to the subsequent processing step. This implies that a data sink can be inserted anywhere in a pipeline, not just at the end.
    - Similar terms that should *NOT* be used: output, data output
- **Pipeline**: An ordered list of processing steps and/or data sinks, with an optional data source.
    - Note: A pipeline without a data source is usually not useful.
- **Fork** (or fork step or fork processing step): A processing step that splits incoming samples based on certain criteria and forwards them to different sub-pipelines.
- **Marshaller**: Entity that can transform a byte stream into a sample, and vice-versa.
    - Similar terms that should *NOT* be used: parser, formatter
- **Batch processing step** (or batch step): A processing step that handles a list of samples, instead of processing samples individually.
    - Similar terms that should *NOT* be used: window
- **Sub-Pipeline**: A pipeline inside a fork. A sub-pipeline is almost equal to a top-level pipeline, except that no separate data input can be defined. The input samples are automatically given by the fork that surrounds the sub-pipeline. Any samples that reach the end of a sub-pipeline must be forwarded to the processing step that follows the surrounding fork.

## Naming of Processing Steps

Processing steps in Bitflow Scripts should have imperative, command-like names. They should be lower case and use hyphens as word separators.
Examples:

- **`output-files`** instead of `file-output` or `OutputFiles`
- **`drop-samples`** instead of `sample-dropper` or `DropSamples`
- **`fork-tags`** instead of `tag-fork` or `ForkTags`

Classes or types that implement processing steps should avoid the suffix **ProcessingStep** or **Step** as much as possible.

In implementations that automatically generates the names of processing steps from class names, should use the following general procedure:

- In case there is a `ProcessingStep` or `Step` suffix, remove it (e.g. `DropSamplesStep` to `DropSamples`)
- Split the camel-case class name into parts and convert to lower case (e.g. `DropSamples` to `drop` and `samples`)
- Combine the resulting terms with hyphens (e.g. `drop-samples`)

The above procedure implies that class names of processing step classes should follow the same imperative, command-like naming convention as the processing steps in Bitflow Script.

## Core Data Sources, Sinks and Marshallers

Bitflow support various sources and sinks. A data source is a way a bitflow program can obtain samples from. The follow ones should be support by all implementations:
- **listen**: binds a tcp socket and waits for clients to connect and send data. Multiple senders are 
accepted. Maximum number of peers might be configurable with global parameter **`--source-listen-max-peers`**
- **download**: connects to a given endpoint and starts downloading data. If connection failes, the sources tries to reconnect constantly. 
- **file**: reads one or more files located on the disk in seqential order.  
- **console**: reads from stdio of the bitflow program.

Apart from source also sink types are defined in bitflow:
- **listen**: binds a tcp socket and waits for clients to connect. After the connection is established the sink-listen begins to send data to the connected peer(s). Sink-listen should supports mutiple peers. Sink-listen might also be able to buffer samples and send them to peers connecting later in time. The number of buffered samples might be configurable using bitflow-script args of the sink step.
- **send**: nnects to a given endpoint and starts sending data. If connection failes, the sources tries to reconnect constantly. After reconnection, the stream will continue at the after the last send sample. 
- **file**: writes samples into a file. 
- **console**: writes samples to stdio.

Further bitflow supports two data formats used to transport and store data.
- **csv**: is a csv conform data format if it is written to file. Each stream contains at least ns accurate timestamp, key-value string pairs called tags and optional double parameters.
- **bin**: binary represantation of a datastream. This improves lightness and reduces network overhead.  

These formats are defined in more detail in section data formats. 

## Sample Templates

Many processing steps take *sample templates* as parameters. A sample template is a string with placeholders, which are replaced by tag values of a sample. A sample template is always evaluated in the context of an individual sample.
The syntax for a placeholder is `${a}`.

For example, if samples should be output in different files based on their `data_type` tag, the output file can be defined as `./output-${data_type}.csv`. Now a sample with the tag `data_type=raw` will be output to the file `./output-raw.csv`, while a sample with the tag `data_type=aggregated` will be output to the file `./output-aggregated.csv`.

Other examples:
- `./${data_type}/${pump}.bin`

## Core Processing Steps

The following is a list of processing steps that should be supported in all language implementations of the Bitflow Script, with exactly the given names.
The parameters should also be supported, as much as possible.

- **`noop()`**: A processing step that does nothing
- **`drop()`**: A processing step that drops all incoming samples
- **`strip()`**: A processing step that removes all metrics
- **`strip-tags()`**: A processing step that removes all tags
- **`filter()`**: A processing step that only forwards samples that meet certain requirements
    - **TODO**: define parameters and semantics of the filter. Additional filter processing steps, e.g. filter-expression?
- **`batch`**: A batch processing step opens a new subpipeline only supporting batch-steps. Batch-pipelines forwards multiple samples at the same time to corresponding batch steps.
- **`fork-tag(tag=)`**: A fork processing step that forks the samples based on a tag value
- **`fork-template(template=)`**: A fork processing step that forks the samples based on a *sample template*
- **`set-tags(a=b, x=y, ...)`**: A processing step that takes arbitrary parameters and sets the given parameters as tags (possibly overwriting existing tags). Both the key and value parts should be evaluated as *sample templates* before setting the resulting value.

## Common Command-Line Options

The following command line flags should be supported by all implementations of Bitflow Script:

- Positional arguments: any non-flag arguments should be concatenated (with an additional space character as separator) and used as the Bitflow Script to be executed.
- **`[-s,--script] <bitflow script>`**: Alternative to providing the script as positional argument. The script is given as a parameter to this flag. Cannot be combined with positional arguments.
- **`[-f, --file] <bitflow script file>`**: Alternative to positional arguments or `-s`. Define a file, from which the Bitflow Script is read.
- **`[-v, --verbose]`**: Set log level to Debug (default is Info)
- **`-q`**: Set log level to Warning
- **`-qq`**: Set log level to Severe
- **`[-c,--capabilities]`**: Output a JSON-formatted list of all supported processing steps
    - TODO: define JSON data format
- **`[-p, --plugin] <plugin>`**: Load a Bitflow-Script plugin that contains new processing steps and/or data sources/sinks. The plugin mechanism and file type is implementation specific.
- **`-log <file>`**: Write all logging output to the given file (or other destination, if the parameter is a URL)
