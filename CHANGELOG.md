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
  <a href="./README.md">README</a>
  ·
  <a href="./CHANGELOG.md"><strong>CHANGELOG</strong></a>
  .
  <a href="./CONTRIBUTING.md">CONTRIBUTING</a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-yoss">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-yoss/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-yoss/issues">Request Feature</a>
</p>

## Changelog


<a name="v0.1.0"></a>
## [v0.1.0] - 2023-03-12
### Build
- release build needs ruby if we are running features
- feature-flag vale so we don't need to run it inside goreleaser
- goreleaser can't call make without a Makefile
- yoss builds need go 1.18

### Code Refactoring
- reduce duplication
- normalize provider arg across get and sync

### Features
- fetch a config value out of AWS Parameter Store
- sync merged configs into AWS Parameter Store
- merge a folder of config files by convention
- merge one file into another to validate a merge
- merge in goconfig CLI from v1.0.0 of go-deep-merge


<a name="v0.0.1"></a>
## v0.0.1 - 2023-03-11
### Build
- initial CLI framework


[Unreleased]: https://github.com/davidalpert/go-yoss/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/davidalpert/go-yoss/compare/v0.0.1...v0.1.0
[license-shield]: https://img.shields.io/badge/License-GPLv3-blue.svg
[license-url]: https://www.gnu.org/licenses/gpl-3.0