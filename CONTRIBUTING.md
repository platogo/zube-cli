# Contributing

## Setup

### Manual

1. Fork and clone the repo
1. Run `make` to setup the project
1. Create a branch for your PR with `git checkout -b your-branch-name`

> To keep `master` branch pointing to remote repository and make pull requests from branches on your fork. To do this, run:
>
> ```sh
> git remote add upstream https://github.com/platogo/zube-cli.git
> git fetch upstream
> git branch --set-upstream-to=upstream/master master
> ```

You can quickly install a locally built release by running

```sh
make install
```

It is assumed you already have familiarity with `Go` and popular CLI libraries, such as

- [Cobra](https://github.com/spf13/cobra-cli)
- [Viper](https://github.com/spf13/viper)

New commands should be generated using `cobra-cli add`

## Pull Request Guidelines

Don't worry if you get any of the below wrong, or if you don't know how. We'll gladly help out.

### Title

Make sure the title starts with a semantic prefix:

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- **refactor**: A code change that neither fixes a bug nor adds a feature
- **perf**: A code change that improves performance
- **test**: Adding missing tests or correcting existing tests
- **build**: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
- **ci**: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
- **chore**: Other changes that don't modify src or test files
- **revert**: Reverts a previous commit

### Testing

Once you submit your pull request it will automatically be tested. Be sure to check the results of the test and fix any issues that arise.

It's also a good idea to consider if your change should include additional tests. This is highly recommended for new features or bug-fixes. For example, it's good practice to create a test for each bug you fix which ensures that we don't regress the code in the future.
