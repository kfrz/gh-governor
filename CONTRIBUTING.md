# Contributing to gh-governor

## Development

1. Use Golang version >= 1.20.7
1. Fork or clone this repository (https://github.com/kfrz/gh-governor)

    ```sh
    gh repo clone kfrz/gh-governor
    ```
1. Create a feature branch
1. Make your changes, test, lint, and format your code
1. Commit your changes 
1. Rebase your local changes against the master branch
1. Create new Pull Request
1. Remediate any issues found by the CI workflow checks
1. Address any review comments
1. Ship it!  🚢

Bugs, feature requests and comments are more than welcome. Please submit an [issue](https://github.com/kfrz/gh-governor).

### Style & Conventions

#### Commit Messages

This project is configured to generate changelogs using [git-chglog](https://github.com/git-chglog/git-chglog). To ensure changelogs are generated correctly, please follow the commit message guidelines below.

* All commits must be signed. To sign commits, run `git config --global commit.gpgsign true` and ensure you have a GPG key configured. For more information on signing commits, see [GitHub's documentation](https://docs.github.com/en/github/authenticating-to-github/signing-commits).

* Commit content should follow Conventional Commit guidelines. See [conventionalcommits.org](https://www.conventionalcommits.org/en/v1.0.0/) for more information. In general, commit messages should be structured as follows:

    ```
    <type>[optional scope]: <description>

    [optional body]

    [optional footer(s)]
    ```
* Types include: 
            `build`, `ci`, `chore`, `docs`, `feat`, `fix`, `perf`, `refactor`, `revert`, `style`, `test`.
* Keep commits focused and atomic, and use the body to provide additional context. If a commit closes an issue, include `Closes #<issue number>` in the footer.

#### pre-commit

`.pre-commit-config.yaml` is included to help with linting and formatting. To install the pre-commit hooks, ensure you have [pre-commit](https://pre-commit.com) installed then run:

```sh
pre-commit install
# or, selectively install hook types
pre-commit install --hook-type commit-msg
```

### Architectural Decisions

- The goal of gh-governor is to be lightweight and safe.
- gh-governor should be able to be run in a dry-run mode to preview changes.
- Any major changes should be discussed in an issue before being implemented.

### Code of Conduct

Be kind.