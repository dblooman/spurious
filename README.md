# Spurious-Go

Go version of the Spurious CLI and Server combined into a single binary

## Description

Spurious is a toolset allowing development against a subset of AWS resources, locally.

The services are run as Docker containers, and Spurious manages their lifecycle and
linking so all you have to worry about is using the services.

To use Spurious, you'll need to change the endpoint and port for each AWS service to those provided by Spurious.

There are a number of supporting libraries that ease the configuration of the AWS SDKs.

## Supported services

Currently the following AWS services are supported by Spurious:

- S3 ([fake-s3](https://github.com/jubos/fake-s3))
- SQS ([fake_sqs](https://github.com/iain/fake_sqs))
- DynamoDB ([DynamoDB Local](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tools.DynamoDBLocal.html))
- ElastiCache ([fake_elasticache](https://github.com/stevenjack/fake_elasticache))
- Spurious Browser ([spurious-browser](https://github.com/stevenjack/spurious-browser))

> The following services are actively in development:

- SNS
- CloudFormation (Allow you to create resources that there are already services for in Spurious).

## Requirements

Spurious works on the following platforms:

- OSX
- Linux
- Windows

### Dependencies

Spurious requires the following to be installed and started to work correctly:

- Docker 1.10.*

## Installation

### Quick install

To install, use `go get`:

```bash
$ go get -d github.com/spurious-io/spurious-go
```

### Manual install

#### Docker

Each of the local services are run inside a Docker container, so without Docker, Spurious won't work.

Please ensure you are running docker-engine on Linux or Docker for Mac/Windows.  Older versions using docker-machine or boot2docker will not work.

#### Spurious

## Usage

Run the following commands to get the containers up and running:

```bash
spurious init
spurious start
```

You should now have six containers running, which you can check with:

```bash
docker ps
```

### GUI

One of the services started by Spurious is the [browser](https://www.github.com/spurious-io/spurious-browser). This allows you to interact and manage the fake services from a graphical interface.

To access the browser service, enter the following command:

```bash
spurious ports
```

This should display output similar to:

```bash
Service                      Host                         Port   Browser link
spurious-browser            browser.spurious.localhost  33103 http://browser.spurious.localhost:33103 <--- Link to browser
spurious-elasticache-docker localhost                   33104 -
spurious-elasticache        localhost                   33105 -
spurious-memcached          localhost                   33101 -
spurious-dynamo             dynamodb.spurious.localhost 33102 http://dynamodb.spurious.localhost:33102
spurious-s3                 s3.spurious.localhost       33100 http://s3.spurious.localhost:33100
spurious-sqs                sqs.spurious.localhost      33099 http://sqs.spurious.localhost:33099
```

You'll find the browser link next to the service `spurious-browser`.

### Using the containers

Once the containers are up and running, they're assigned random port numbers from Docker which are available on the ip address of the VM used to run the containers. To make the discovery of these ports simpler there's the following command:

```bash
spurious ports
```

This will return a list of host and port details for each of the Spurious containers.

### SDK Helpers

Once the containers are running you'll need to wire up the SDK to point to the correct endpoints and port numbers. Here's an example using the Ruby SDK:

```ruby
require 'json'

port_config = JSON.parse(`spurious ports --json`)

AWS.config(
  :region              => 'eu-west-1',
  :use_ssl             => false,
  :access_key_id       => "access",
  :secret_access_key   => "secret",
  :dynamo_db_endpoint  => port_config['spurious-dynamo']['Host'],
  :dynamo_db_port      => port_config['spurious-dynamo']['HostPort'],
  :sqs_endpoint        => port_config['spurious-sqs']['Host'],
  :sqs_port            => port_config['spurious-sqs']['HostPort'],
  :s3_endpoint         => port_config['spurious-s3']['Host'],
  :s3_port             => port_config['spurious-s3']['HostPort'],
  :s3_force_path_style => true
)
```

There are also helpers available for the different flavours of the AWS SDK:

* [Ruby](https://github.com/spurious-io/ruby-awssdk-helper)
* [Clojure](https://github.com/Integralist/spurious-clojure-aws-sdk-helper)
* [Javascript](https://github.com/spurious-io/js-aws-sdk-helper)

## Contribution

1. Fork ([https://github.com/spurious-io/spurious-go/fork](https://github.com/spurious-io/spurious-go/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[DaveBlooman](https://github.com/DaveBlooman)
