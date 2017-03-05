# Rabbitmqbeat

Rabbitmqbeat is a beat based on metricbeat which was generated with metricbeat/metricset generator.


## Getting started

To compile your beat run `make`. Then you can run the following command to see the first output:

```
rabbitmqbeat -e -d "*"
```

## Configuation

The configuration settings can be found in the `rabbitmqbeat.yml` file. In there you can configure which RabbitMQ and Elasticsearch hosts to use.