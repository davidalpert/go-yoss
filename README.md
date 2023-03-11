<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![License: GPL v3][license-shield]][license-url]
<!-- [![Issues][issues-shield]][issues-url] -->
<!-- [![Forks][forks-shield]][forks-url] -->
<!-- ![GitHub Contributors][contributors-shield] -->
<!-- ![GitHub Contributors Image][contributors-image-url] -->

<!-- PROJECT LOGO -->
<br />
<!-- vale Google.Headings = NO -->
<h1 align="center"><code>go-yoss</code></h1>
<!-- vale Google.Headings = YES -->

<p align="center">
  A command-line tool for yossing configuration into a remote configuration management store.
  <br />
  <a href="./README.md"><strong>README</strong></a>
  ·
  <a href="./CHANGELOG.md">CHANGELOG</a>
  .
  <a href="./CONTRIBUTING.md">CONTRIBUTING</a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-yoss">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-yoss/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-yoss/issues">Request Feature</a>
</p>

<details open="open">
  <summary><h2 style="display: inline-block">Table of contents</h2></summary>

- [About the project](#about-the-project)
  - [Built with](#built-with)
- [Getting started](#getting-started)
  - [Install](#install)
    - [`go install`](#go-install)
    - [Pre-compiled binaries](#pre-compiled-binaries)
  - [Verify your installation](#verify-your-installation)
  - [Uninstall](#uninstall)
- [Usage](#usage)
- [Troubleshooting](#troubleshooting)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

</details>

<!-- ABOUT THE PROJECT -->
## About the project

Managing configuration as code has significant benefits.

`yoss` offers a way to do just that:

- store configuration data as structured YAML documents
- perform convention-based merges and synchronize the diff with a remote configuration store

### Built with

* [Golang 1.16](https://golang.org/)
* [go-releaser](https://goreleaser.com/)

<!-- GETTING STARTED -->
## Getting started

### Install

#### `go install`

With a working golang installation at version >= 1.16 you can install or update with:

```
go install github.com/davidalpert/go-yoss/cmd/yoss@latest
```

#### Pre-compiled binaries

Visit the [Releases](https://github.com/davidalpert/go-yoss/releases) page to find binary packages pre-compiled for a variety of `GOOS` and `GOARCH` combinations:
1. Download an appropriate package for your `GOOS` and `GOARCH`;
1. Unzip it and put the binary in your path;

### Verify your installation

1. Check the version installed:
    ```
    yoss version
    ```

### Uninstall

- `go-yoss` ships with an `uninstall` sub-command which cleans up and removes itself:

    ```
    yoss uninstall
    ```

<!-- USAGE EXAMPLES -->
## Usage

- TODO; coming as the project nears v1.0

<!-- Troubleshooting -->
## Troubleshooting

If you run into trouble you can ask `yoss` to write some diagnostics to a log file by setting the following environment variables:

| Variable        | Default   | Description                                                      |
| --------------- | --------- | ---------------------------------------------------------------- |
| YOSS_LOG_LEVEL  | `"fatal"` | `"fatal"`, `"error"`, `"warning"`, `"warn"`, `"info"`, `"debug"` |
| YOSS_LOG_FORMAT | `"text"`  | `"text"` or `"json"`                                             |
| YOSS_LOG_FILE   | `""`      | path to a log file; when empty logs go to STDOUT                 |

<!-- ROADMAP -->
## Roadmap

<!-- vale Google.Parens = NO -->
See [open issues](https://github.com/davidalpert/go-yoss/issues) project board for a list of known issues and up-for-grabs tasks.
<!-- vale Google.Parens = YES -->

## Contributing

See the [CONTRIBUTING](CONTRIBUTING.md) guide for local development setup and contribution guidelines.

<!-- LICENSE -->
## License

Distributed under the GPU v3 License. See [LICENSE](LICENSE) for more information.

<!-- CONTACT -->
## Contact

David Alpert - [@davidalpert](https://twitter.com/davidalpert)

Project Link: [https://github.com/davidalpert/go-yoss](https://github.com/davidalpert/go-yoss)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/davidalpert/go-yoss
[contributors-image-url]: https://contrib.rocks/image?repo=davidalpert/go-yoss
[forks-shield]: https://img.shields.io/github/forks/davidalpert/go-yoss
[forks-url]: https://github.com/davidalpert/go-yoss/network/members
[issues-shield]: https://img.shields.io/github/issues/davidalpert/go-yoss
[issues-url]: https://github.com/davidalpert/go-yoss/issues
[license-shield]: https://img.shields.io/badge/License-GPLv3-blue.svg
[license-url]: https://www.gnu.org/licenses/gpl-3.0

