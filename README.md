# Vamp Cloud command line client

Vamp Cloud Cli is a command line client written in Go and allows to easily interact with Vamp Cloud.

## Table of Contents

================

- [Vamp Cloud command line client](#vamp-cloud-command-line-client)
    - [Table of Contents](#table-of-contents)
    - [Development](#development)
    - [Build](#build)
    - [Installation](#installation)

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


