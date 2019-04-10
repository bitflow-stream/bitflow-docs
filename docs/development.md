# Development Process for Bitflow Repositories

This page describes the development process for the following repositories:

- [https://github.com/bitflow-stream/bitflow4j](https://github.com/bitflow-stream/bitflow4j)
- [https://github.com/bitflow-stream/python-bitflow](https://github.com/bitflow-stream/python-bitflow)
- [https://github.com/bitflow-stream/antlr-grammars](https://github.com/bitflow-stream/antlr-grammars)
- [https://github.com/bitflow-stream/go-bitflow-collector](https://github.com/bitflow-stream/go-bitflow-collector)
- [https://github.com/bitflow-stream/go-bitflow](https://github.com/bitflow-stream/go-bitflow)
- [https://github.com/bitflow-stream/bitflow-docs](https://github.com/bitflow-stream/bitflow-docs)
- [https://github.com/bitflow-stream/bitflow-coordinator](https://github.com/bitflow-stream/bitflow-coordinator)
- [https://github.com/bitflow-stream/bitflow-process-agent](https://github.com/bitflow-stream/bitflow-process-agent)
- [https://github.com/citlab/zerops-orchestrator](https://github.com/citlab/zerops-orchestrator) (private)
- [https://github.com/citlab/zerops-analysis](https://github.com/citlab/zerops-analysis) (private)
- [https://github.com/citlab/distributed-anomaly-injection](https://github.com/citlab/distributed-anomaly-injection) (private)

## Branches

The main branch is the `master` branch, there is no `develop` branch.
There are no direct pushes to the `master` branch, only pull requests from *feature branches* are allowed.
Pull requests must be reviewed and accepted by *1* other team member.
Everybody can review and accept pull requests.
Try to reduce code ownership and repository ownership.

Feature branches should follow the following naming scheme: `<owner>-<feature>`, where the first part is the name of the main contributor of the feature branch, and the second part is the purpose of the branch.

## Versioning

The Bitflow projects follow [semantic versioning](https://semver.org/).
Released version are marked through Git tags with names like `v0.0.1` or `v1.2.3`.
The major and minor versions are shared by all projects listed above, and indicate cross-project compatibility.
The patch (last part in the version string) is incremented individually in each repository.
This means, the `bitflow4j` repository version `v0.2.3` is compatible with the `go-bitflow` repository version `v0.2.12`.

In the future, the CI pipeline should support automatic version bumping.
This will work through special keywords in commit messages, which can be appended to the *merge commits* when accepting pull requests.
This CI server will then automatically increment the version in a version file (if necessary, this depends on the programming language), and add a Git version tag.
