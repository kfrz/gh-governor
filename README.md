# Governor (`gh-governor`) 🕵🏻‍♂️

 A GitHub (`gh`) CLI extension to help audit and enforce governance requirements on **repositories** by defining and invoking policy-based **enforcers**.

<a href="https://github.com/kfrz/gh-governor/actions"><img alt="GitHub Workflow Status (with event)" src="https://img.shields.io/github/actions/workflow/status/kfrz/gh-governor/release.yml?event=push&style=flat-square"></a>
<a href="https://github.com/kfrz/gh-governor/releases"><img src="https://img.shields.io/github/release/kfrz/gh-governor.svg" alt="Latest Release"></a>
<a href="./LICENSE.txt"><img alt="GitHub" src="https://img.shields.io/github/license/kfrz/gh-governor"></a>
<a href="https://goreportcard.com/report/github.com/kfrz/gh-governor"><img alt="Go ReportCard" src="https://goreportcard.com/badge/github.com/kfrz/gh-governor"></a>

## 🔋 Features

* 🧐 | Interactive **search** for repos based on given *`Enforcers`* (CODEOWNERS, branch-protection, etc.)
* 📝 | Generate audit reports of governance status across one, many, or all repositories in an organization or account
* ☎️  | Remediate and act on audit results with custom workflows, create issues, open PRs, etc.
* 🔭 | Automate reporting and view details about a repository with a detailed governance audit analysis
* 📚 | Configure custom `Enforcers` allowing for fine-grained governance policy management

## 📦 Installation

1. Install the `gh` CLI - see the [installation](https://github.com/cli/cli#installation)

   *Installation requires a minimum version (2.0.0) of the the GitHub CLI that supports extensions.*

2. Install this extension:

   ```sh
   gh extension install kfrz/gh-governor
   ```

3. Ensure the `gh` CLI is authenticated:

    ```sh
    gh auth login
    ```

<details>
   <summary>Manual Installation</summary>

1. Clone the repo:

   ```sh
   # git
   git clone https://github.com/kfrz/gh-governor

   # GitHub CLI
   gh repo clone kfrz/gh-governor
   ```

2. `cd` into the repo directory:

   ```sh
   cd gh-governor
   ```

3. Install the extension:

   ```sh
   gh extension install .
   ```

 </details>

## ⚡️ Usage

Run `gh governor --help` for more info:

```
Usage:
  gh governor [flags]

Flags:
  -c, --config string   use this configuration file (default is $GH_GOVERNOR_CONFIG, or $XDG_CONFIG_HOME/gh-governor/config.yml)
      --debug           passing this flag will allow writing debug output to debug.log
  -h, --help            help for gh-governor
```

## ⚙️ Configuration

All configuration is provided within a `config.yml` file under the extension's directory (either `$XDG_CONFIG_HOME/gh-governor` or `~/.config/gh-governor/` or your OS config dir) or `$GH_GOVERNOR_CONFIG`.

An example `config.yml` file contains:

```yml
---
# General configuration for gh-governor
governor:
  debug: false
  dry_run: false
  enforce_all: true
  sign_commits: true
  verbose: false

# Enforcer configuration. Can be used to enable/disable enforcers, and configure their arguments.
# For more information on enforcers, see https://governor.github.io/docs/enforcers
enforcers:
    - name: "CodeownersEnforcer"
        enabled: true
        arguments:
        codeowner_is_team: true
        ownership_rules:
            - team: "@team"
              patterns:
                - "path/to/file"
                - "path/to/other/file"
            - team: "@my-other-team"
              patterns:
                - "path/to/file"
                - "path/to/other/file"
    - name: "DefaultBranchNameEnforcer"
        enabled: true
        arguments:
        default_branch_name: "main"
```

You can run `gh governor --config <path-to-file>` to run `gh-governor` against another config file.

### 🔐 Authentication

Typically, `gh-governor` uses the same authentication as the `gh` CLI. 

You can authenticate with `gh auth login` or `gh auth login --with-token` to use a Personal Access Token (PAT).


#### CI/CD Authentication

If you are using `gh-governor` in a CI/CD environment, you can use the `--with-token` flag to authenticate with a PAT.

#### Scope & Permissions

* The PAT must have the `repo` scope.
* The PAT must have the `read:org` scope if you are using the `--org` flag.

### 🗃 Configuring Enforcers

For **repositories**, the available default Enforcers are:

**`CodeownersEnforcer`**
| Argument      | Description                                                                     |
| ------------- | ------------------------------------------------------------------------------- |
| `CodeownerIsTeam`    | Boolean, whether or not the CODEOWNER is required to be a team or not. (default true) |

<br />

**`DefaultBranchNameEnforcer`**
| Argument      | Description                                                                     |
| ------------- | ------------------------------------------------------------------------------- |
| `DefaultBranchName` | The enforced name of the default branch (e.g. `main`)                     |

#### Contributing

`gh-governor` is an open source project and contributions are welcome!

Check out the [CONTRIBUTING](./CONTRIBUTING.md) guide to get started.

#### License & Authors

* [MIT License](./LICENSE)
* Authors: <a href="https://github.com/kfrz"> @kfrz </a>
