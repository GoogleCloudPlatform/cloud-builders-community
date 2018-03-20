# Dataflow

## Introduction

[Google Cloud Dataflow](https://cloud.google.com/dataflow/), based on [Apache Beam](https://beam.apache.org/), is a fully-managed service for transforming and enriching data in stream (real time) and batch (historical) modes with equal reliability and expressiveness.  Developers and Data Scientists use Dataflow to process large amounts of data without managing complex cluster infrastructure.

Google Container Builder offers a number of advantages for Cloud Dataflow developers:
* Small workloads (which run in 1 `n1-standard-1` virtual machine) can take advantage of the [free tier](https://cloud.google.com/container-builder/pricing), which provides 120 free build-minutes per day
* Pipelines get all the benefits of containerization, including a consistent environment and integration with your CI/CD pipeline
* 
