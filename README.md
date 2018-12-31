# env-secrets
[![Build Status](https://travis-ci.org/sganon/env-secrets.svg?branch=master)](https://travis-ci.org/sganon/env-secrets)
[![codecov](https://codecov.io/gh/sganon/env-secrets/branch/master/graph/badge.svg)](https://codecov.io/gh/sganon/env-secrets)

Securely fetch and set env secrets needed for a project using Bitwarden or 1password CLI

![](img/env-secrets_bw.gif)

## Prerequisites

Each of the available password managers (1password and Bitwarden) need to have their CLI installed, ofc you need to install only the one you will use. For Bitwarden installation instructions are located [here](https://github.com/bitwarden/cli#downloadinstall), and for 1passsword [here](https://support.1password.com/command-line-getting-started/#set-up-the-command-line-tool).

Env-secrets is only a wrapper around the CLI so you'll need to be sure you can use their CLI as it is before using env-secrets. You'll surely need to login, here are the instructions for [Bitwarden](https://help.bitwarden.com/article/cli/#session-management) and [1password](https://support.1password.com/command-line-getting-started/#get-started-with-the-command-line-tool).

## Usage

Both commands works the same way: you first need to organize your secrets into folders for Bitwarden and tags for 1password.

**NB**: For now I've only tested with *Secure Notes* in Bitwarden, and *Password* in 1password as item type.

Then you should be able to do:
```shell
env-secrets [bw || 1p] [--domain=<OP_DOMAIN>] <FOLDERS/TAGS> ...
```

**NB**: The op command needs an additional flag `--domain` corresponding to the subdomain of your vault's URL.

## Install

For now you'll need to have go installed and build binary yourself. I will make release in the future.
