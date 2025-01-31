# platform-hello-world-go

[![Platform CI/CD](https://github.com/uniphore/platform-hello-world-go/actions/workflows/cicd.yml/badge.svg)](https://github.com/uniphore/platform-hello-world-go/actions/workflows/cicd.yml)

A template for platform services written in Go

## Description

This repository includes a **production-ready** *Hello World* example that can
be used as a template/reference for new services written in Go.

This repository is a **template repository**. This means you can't fork it, but
you can still create a new repository based on it. This solution has several
advantages. A couple important advantages are that the commit history is not
carried over and the contributions graph works as expected. For more information, see the
[GitHub documentation](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template).

---

## Usage

Select the `Use this template` button at the top of a template repository home page on GitHub.com:
![github-template]

Alternatively, when creating a new repository, there is an option to select the appropriate template.

## Post-init steps

These steps are **mandatory**. You must follow them carefully.

### Bootstrap the new repository

1. Clone the new service repository

```shell
git clone https://github.com/uniphore/${NEW-SERVICE-REPOSITORY}
```

2. Go to the cloned repository

```
cd ${NEW-SERVICE-REPOSITORY}
```

3. Run the following script and follow with instructions

```shell
./bootstrap-repo.sh
```

**NOTE**: It's recommended to name the service based on the repository name. The
above script does just that. If a different service name is required, the
script accepts the new name as an optional argument.

4. Check that the ownership of the repository is properly configured in
   [service.datadog.yaml](service.datadog.yaml) and [CODEOWNERS](CODEOWNERS)

5. Merge new changes from the `bootstrap-repo` branch to the `main` branch via
   Pull Request

---

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
garden get test-result hello-world-go-image unit
```

## QA

QA is executed through the Platform CI/CD pipeline, and the outcomes can be accessed on the
[SonarQube dashboard](https://plat-sonar.uniphoredev.com/dashboard?id=platform-hello-world-go).

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

[github-template]: https://docs.github.com/assets/cb-100333/images/help/repository/use-this-template-button.png
