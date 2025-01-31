# uzng-service-status

[![Platform CI/CD](https://github.com/uniphore/uzng-service-status/actions/workflows/cicd.yml/badge.svg)](https://github.com/uniphore/uzng-service-status/actions/workflows/cicd.yml)

This service has been created from a template:
[platform-hello-world-go](https://github.com/uniphore/platform-hello-world-go)

## Requirements

Detailed instructions can be found in our internal documentation:
[Setting up Platform Development Environment](https://uniphore.atlassian.net/wiki/spaces/PlatEng/pages/2093744464/Setting+up+Platform+Development+Environment).

## Building

### Local machine

Use the `make` rules for local development, such as:

```shell
make build
```

### Remote cluster

```shell
garden build
```

## Testing

### Local machine

Use the `make` rules for local testing, such as:

* Run all tests and generate reports

```shell
make test
```

### Remote cluster

```shell
garden test
```

Test results can be retrieved using:

```shell
garden get test-result uzng-service-status-image unit
```

## QA

QA is executed through the Platform CI/CD pipeline, and the outcomes can be accessed on the
[SonarQube dashboard](https://plat-sonar.uniphoredev.com/dashboard?id=uzng-service-status).

## Deployment

### Remote cluster

```shell
garden deploy
```

### Continuous development

Use the `--watch` argument to hot syncing all local changes to the cluster,
for example:

```shell
garden deploy --watch
```

### Production and staging

Before deploying a new service or app to staging or production,
it must first be added to the [Flux clusters configuration](https://github.com/uniphore/platform-flux).
If there is no Flux configuration for the service, the deploy job in the CI/CD pipeline will have no effect.

For more information about Flux configuration, please refer to the [application provisioning](https://platform.cloud.uniphore.com/application_provisioning.html) documentation.

