# Vamp Cloud command line client

Vamp Cloud Cli is a command line client written in Go and allows to easily interact with Vamp Cloud.

## Table of Contents

================

- [Vamp Cloud command line client](#vamp-cloud-command-line-client)
    - [Table of Contents](#table-of-contents)
    - [Development](#development)
    - [Build](#build)
    - [Installation](#installation)
    - [Usage](#usage)
        - [Create cluster](#create-cluster)
        - [Create application](#create-application)
        - [Create ingress](#create-ingress)  
        - [Describe cluster](#describe-cluster)  
        - [Describe application](#describe-application)  
        - [Get installation command](#get-installation-command)
        - [Watch release](#watch-release)                          

## Build

For docker build:

```shell
./build.sh
```

for local build:

```shell
./build.sh local
```

binaries will be place under the bin directory

## Installation

If you have binaries built locally:

For mac run:

```shell
./bin/vamp-darwin-amd64 --help
```

or copy the binaries to you /usr/local/bin/vamp folder.

If you have downloaded the binaries directly. Just copy the binary appropriate to you platform to the user binaries folder. For example for MacOs:

```shell
cp vamp-darwin-amd64 /usr/local/bin/vamp
chmod +x /usr/local/bin/vamp
```

Alternatively you can easily install the cli for MacOS or Linux by running

```shell
version=$(curl -s https://api.github.com/repos/magneticio/vamp-cloud-cli/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/') &&
  base=https://github.com/magneticio/vamp-cloud-cli/releases/download/$version &&
  curl -sL $base/vamp-$(uname -s)-$(uname -m) >/usr/local/bin/vamp &&
  chmod +x /usr/local/bin/vamp
```

Keep in mind this command might fail, give the fact that the repository is private.

For general users it is recommended to download the binary for your platform.
The latest release can be found here:
https://github.com/magneticio/vamp/releases/latest

You can verify your installation by running

```
vamp version
```

## Prerequisites

In order to use Vamp Cloud cli you need to provide the following environment variables

  - VAMP_CLOUD_HOST: the Vamp Cloud api host
  - VAMP_CLOUD_BASE_PATH: the base path for the Vamp Cloud api host
  - VAMP_CLOUD_API_KEY: the Vamp Cloud api key that you can obtaion from the UI


## Usage

Vamp Cloud cli can be used to monitor the status of ongoing releases on services in Vamp Cloud.
To reach that goal it offers commands to query and create different resources in Vamp Cloud.
Here's a list of the available commands:

### Create cluster

Allows to create a new cluster in the project associated to the specified api key.

Example:
```
vamp create cluster <cluster-name> --provider=<provider> --description=<description>
```

### Create application

Allows to create a new application in the specified cluster.

Example:
```
vamp create application <application_name> --cluster=<cluster-name> --namespace=<namespace> --ingress-type=<ingress-type>
```

### Create ingress

Allows to create a new ingress in the specified application.

Example:
```
vamp create ingress <domain-nam> --application=<application-name> --tls-secret=<tls-secret-name>
```

### Describe cluster

Describes the specified cluster in the project associated to the api key.

Example:
```
vamp-cloud-cli describe cluster <cluster-name>
```

Example output:
```
  NAME             PROVIDER     DESCRIPTION
 ---------------- ------------ ---------------
  <cluster-name>   <provider>   <description>

```

### Describe application

Describes the specified application in the project associated to the api key.

Example:
```
vamp-cloud-cli describe application <application-name>
```

Example output:
```
  NAME                 CLUSTER          NAMESPACE     INGRESS(ES)
 -------------------- ---------------- ------------- ---------------
  <application-name>   <cluster-name>   <namespace>   <domain-name>

```

### Get installation command

Retrieves the Release Agent installation command for the specified application in the project associated to the api key.

Example:
```
vamp-cloud-cli get token <application-name>
```

### Watch release

Retrieves the latest release for the specified application and service and checks its status every 30 seconds, until it finishes or fail.

Example:
```
vamp watch release <service-name> --application=<application-name>
```

Example output:
```
 NAME           TYPE             SOURCE             TARGET             STEP   STATUS     HEALTH
-------------- ---------------- ------------------ ------------------ ------ ---------- ----------
<service-name>  <release-type>   <source-version>   <target-version>   1      <status>   <health>
<service-name>  <release-type>   <source-version>   <target-version>   1      <status>   <health>
<service-name>  <release-type>   <source-version>   <target-version>   2      <status>   <health>
<service-name>  <release-type>   <source-version>   <target-version>   3      <status>   <health>
<service-name>  <release-type>   <source-version>   <target-version>   4      <status>   <health>
```