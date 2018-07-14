# Map-processor

Library creates pipeline which is defined in yaml file. Pipeline creates map of properties which are then processed by different stages (here stages are called sinks).

Pipeline must start with source, and then can contain one or more sinks, example file can be found in `pipe.yml`.

## List of sources

- ssm - reads properties from AWS Parameter store and passes them as map, has 3 arguments
  - region from which ssm parameters should be read
  - bool flag which tells if secure strings should be decrypted
  - path ssm parameter path which should be recursively queried

## List of pipelines

- print - prints to std out map of properties which was passed to it