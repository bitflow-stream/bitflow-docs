# The Bitflow Project

Bitflow is an open source project developed at the (Complex and Distributed IT Systems Group (CIT))[https://www.cit.tu-berlin.de] at the (Technical University of Berlin)[https://www.tu.berlin/].
Bitflow is a lightweight distributed datastream processing engine that works on Kubernetes and features a domain specific language for expressing operator graphs (Bitflowscript) and a flexible dataflow graph model.
The core functionallity is written in Go, with plugins available in Java and Python.
Bitflow allows to read data from sources and process them in arbitrary ways.
Compared to Big Data frameworks such as Apache Flink, Bitflow does not assume to run on dedicated hardware or resources.
It is designed to run alongside other workload and minimizes its influence on the "main workload" through intelligent scheduling and by limiting the resources assigned to stream processing containers.

The documentation for this project is work in progress. Missing documentation:

0. Usage guideline, Getting started
1. Pipelines
2. Naming conventions in implementations:
    1. Naming of processing steps and parameters
    2. Naming of types in the code
3. Processing steps supported in all implementations:
    1. Filters
    2. Forks
4. Command line flags supported in all implementations
    1. Examples on how to start CLI tools. Docs for individual tools.
