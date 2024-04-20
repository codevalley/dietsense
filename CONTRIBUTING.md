# Contributing to DietSense

We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:
- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## We Develop with Github

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

## We Use [Github Flow](https://guides.github.com/introduction/flow/index.html), So All Code Changes Happen Through Pull Requests

Pull requests are the best way to propose changes to the codebase. We actively welcome your pull requests:

1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

## Any contributions you make will be under the MIT Software License

In short, when you submit code changes, your submissions are understood to be under the same [MIT License](http://opensource.org/licenses/MIT) that covers the project. Feel free to contact the maintainers if that's a concern.

## Report bugs using Github's [issues](https://github.com/codevalley/dietsense/issues)

We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/yourusername/dietsense/issues/new); it's that easy!

## Write bug reports with detail, background, and sample code

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can.
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

People *love* thorough bug reports. I'm not even kidding.

## Use a Consistent Coding Style

* 2 spaces for indentation rather than tabs
* You can try running `go fmt` for automatically formatting code in a standard way
* Make sure to run `go vet` to check for issues that the compiler does not catch
* Test your code as you go.

## Contributing Guidelines

### Configuration File Policy

- The `config.yaml` file in the repository is a template for development purposes. It contains non-sensitive, generic default values.
- **Do Not Commit Changes to `config.yaml`**: If you need to update this file (e.g., to add new configuration parameters), you must discuss the changes with the project maintainers and obtain approval before committing these changes.
- **Local Changes**: Developers should configure their local development environments without altering the tracked `config.yaml` file. Use the command `git update-index --assume-unchanged config.yaml` to ignore local changes to this file.
- **Commit Hooks**: This repository uses pre-commit hooks to prevent accidental changes to `config.yaml`. If you believe your changes are necessary, refer to the discussion and approval process outlined above.

Please adhere to these guidelines to maintain the integrity and security of the project configuration.

## License

By contributing, you agree that your contributions will be licensed under its MIT License.

## References

This document was adapted from the open-source contribution guidelines for [Facebook's Draft](https://github.com/facebook/draft-js/blob/master/CONTRIBUTING.md)
