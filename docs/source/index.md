# Welcome to the Bitflow documentation

Bitflow is an open source project developed at the Complex and Distributed IT Systems Group at the technical university of Berlin. It was first mentioned to be a framework to fast and easily implement and test machine learning based anomaly detection algorithms for automated cloud monitoring. The core functionallity is written in Java, GO and Python. Bitflow allows to read data from sources and process them in arbitrary ways. Compared to big data Frameworks like Apache Flink, Bitflow does not require dedicated hardware, its mentioned to run beside the actual workload.


1. [Sources, Sinks the data format](data_format.md)
2. Pipelines
3. Processing steps
  1. Filter
  2. Forks
  3. Anomaly Detection
4. [Bitflow-Script](script.md)

## Implementations 

- [go-bitflow](https://github.com/bitflow-stream/go-bitflow)
- [bitflow4j](https://github.com/bitflow-stream/bitflow4j)
- [python-bitflow](https://github.com/bitflow-stream/python-bitflow)

## Related Projects

- [bitflow-collector](https://github.com/bitflow-stream/antlr-grammars)
- [bitflow-attlr-grammar](https://github.com/bitflow-stream/antlr-grammars)
- [bitflow-coordinator](https://github.com/bitflow-stream/bitflow-coordinator)
- [bitflow-process-agent](https://github.com/bitflow-stream/bitflow-process-agent)