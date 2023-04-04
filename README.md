---

<p align="center">
  <strong>
    <a href="https://opentelemetry.io/docs/collector/getting-started/">Getting Started</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/CONTRIBUTING.md">Getting Involved</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://cloud-native.slack.com/archives/C01N6P7KR6W">Getting In Touch</a>
  </strong>
</p>

<p align="center">
<<<<<<< HEAD
  <a href="https://github.com/open-telemetry/opentelemetry-collector-contrib/actions/workflows/build-and-test.yml?query=branch%3Amain">
    <img alt="Build Status" src="https://img.shields.io/github/actions/workflow/status/open-telemetry/opentelemetry-collector-contrib/build-and-test.yml?branch=main&style=for-the-badge">
  </a>
  <a href="https://goreportcard.com/report/github.com/open-telemetry/opentelemetry-collector-contrib">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/open-telemetry/opentelemetry-collector-contrib?style=for-the-badge">
  </a>
  <a href="https://codecov.io/gh/open-telemetry/opentelemetry-collector-contrib/branch/main/">
    <img alt="Codecov Status" src="https://img.shields.io/codecov/c/github/open-telemetry/opentelemetry-collector-contrib?style=for-the-badge">
=======
  <a href="https://github.com/open-telemetry/opentelemetry-collector/actions/workflows/build-and-test.yml?query=branch%3Amain">
    <img alt="Build Status" src="https://img.shields.io/github/actions/workflow/status/open-telemetry/opentelemetry-collector/build-and-test.yml?branch=main&style=for-the-badge">
  </a>
  <a href="https://goreportcard.com/report/github.com/open-telemetry/opentelemetry-collector">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/open-telemetry/opentelemetry-collector?style=for-the-badge">
  </a>
  <a href="https://codecov.io/gh/open-telemetry/opentelemetry-collector/branch/main/">
    <img alt="Codecov Status" src="https://img.shields.io/codecov/c/github/open-telemetry/opentelemetry-collector?style=for-the-badge">
>>>>>>> upstream/main
  </a>
  <a href="https://github.com/open-telemetry/opentelemetry-collector-contrib/releases">
    <img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/v/release/open-telemetry/opentelemetry-collector-contrib?include_prereleases&style=for-the-badge">
  </a>
</p>

<p align="center">
  <strong>
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/vision.md">Vision</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/design.md">Design</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
<<<<<<< HEAD
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/monitoring.md">Monitoring</a>
=======
    <a href="https://opentelemetry.io/docs/collector/configuration/">Configuration</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="docs/monitoring.md">Monitoring</a>
>>>>>>> upstream/main
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/performance.md">Performance</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
<<<<<<< HEAD
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/security-best-practices.md">Security</a>
=======
    <a href="docs/security-best-practices.md">Security</a>
>>>>>>> upstream/main
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/roadmap.md">Roadmap</a>
  </strong>
</p>

---

# OpenTelemetry Collector Contrib

<<<<<<< HEAD
This is a repository for OpenTelemetry Collector components that are not suitable for the  [core repository](https://github.com/open-telemetry/opentelemetry-collector) of the collector. 
=======
The OpenTelemetry Collector offers a vendor-agnostic implementation on how to
receive, process and export telemetry data. In addition, it removes the need
to run, operate and maintain multiple agents/collectors in order to support
open-source telemetry data formats (e.g. Jaeger, Prometheus, etc.) to
multiple open-source or commercial back-ends.
>>>>>>> upstream/main

The official distributions, core and contrib, are available as part of the [opentelemetry-collector-releases](https://github.com/open-telemetry/opentelemetry-collector-releases) repository. Some of the components in this repository are part of the "core" distribution, such as the Jaeger and Prometheus components, but most of the components here are only available as part of the "contrib" distribution. Users of the OpenTelemetry Collector are also encouraged to build their own custom distributions with the [OpenTelemetry Collector Builder](https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder), using the components they need from the core repository, the contrib repository, and possibly third-party or internal repositories.

Each component has its own support levels, as defined in the following sections. For each signal that a component supports, there's a stability level, setting the right expectations. It is possible then that a component will be **Stable** for traces but **Alpha** for metrics and **Development** for logs.

## Stability levels

Stability level for components in this repository follow the [definitions](https://github.com/open-telemetry/opentelemetry-collector#stability-levels) from the OpenTelemetry Collector repository.

## Gated features

Some features are hidden behind feature gates before they are part of the main code path for the component. Note that the feature gates themselves might be at different [lifecycle stages](https://github.com/open-telemetry/opentelemetry-collector/tree/main/featuregate#feature-lifecycle).

## Support

Each component is supported either by the community of OpenTelemetry Collector Contrib maintainers, as defined by the GitHub group [@open-telemetry/collector-contrib-maintainer](https://github.com/orgs/open-telemetry/teams/collector-contrib-maintainer), or by specific vendors. See the individual README files for information about the specific components.

The OpenTelemetry Collector Contrib maintainers may at any time downgrade specific components, including vendor-specific ones, if they are deemed unmaintained or if they pose a risk to the repository and/or binary distribution.

Even though the OpenTelemetry Collector Contrib maintainers are ultimately responsible for the components hosted here, actual support will likely be provided by individual contributors, typically a code owner for the specific component.

## Stability levels

The collector components and implementation are in different stages of stability, and usually split between
functionality and configuration. The status for each component is available in the README file for the component. While
we intend to provide high-quality components as part of this repository, we acknowledge that not all of them are ready
for prime time. As such, each component should list its current stability level for each telemetry signal, according to
the following definitions:

### Development

Not all pieces of the component are in place yet and it might not be available as part of any distributions yet. Bugs and performance issues should be reported, but it is likely that the component owners might not give them much attention. Your feedback is still desired, especially when it comes to the user-experience (configuration options, component observability, technical implementation details, ...). Configuration options might break often depending on how things evolve. The component should not be used in production.

### Alpha

The component is ready to be used for limited non-critical workloads and the authors of this component would welcome your feedback. Bugs and performance problems should be reported, but component owners might not work on them right away. The configuration options might change often without backwards compatibility guarantees.

### Beta

Same as Alpha, but the configuration options are deemed stable. While there might be breaking changes between releases, component owners should try to minimize them. A component at this stage is expected to have had exposure to non-critical production workloads already during its **Alpha** phase, making it suitable for broader usage.

### Stable

The component is ready for general availability. Bugs and performance problems should be reported and there's an expectation that the component owners will work on them. Breaking changes, including configuration options and the component's output are not expected to happen without prior notice, unless under special circumstances.

### Deprecated

The component is planned to be removed in a future version and no further support will be provided. Note that new issues will likely not be worked on. When a component enters "deprecated" mode, it is expected to exist for at least two minor releases. See the component's readme file for more details on when a component will cease to exist.

### Unmaintained

A component identified as unmaintained does not have an active code owner. Such component may have never been assigned a code owner or a previously active code owner has not responded to requests for feedback within 6 weeks of being contacted. Issues and pull requests for unmaintained components will be labelled as such. After 6 months of being unmaintained, these components will be removed from official distribution. Components that are unmaintained are actively seeking contributors to become code owners.

## Compatibility

When used as a library, the OpenTelemetry Collector attempts to track the currently supported versions of Go, as [defined by the Go team](https://go.dev/doc/devel/release#policy).
Removing support for an unsupported Go version is not considered a breaking change.

Starting with the release of Go 1.18, support for Go versions on the OpenTelemetry Collector will be updated as follows:

1. The first release after the release of a new Go minor version `N` will add build and tests steps for the new Go minor version.
2. The first release after the release of a new Go minor version `N` will remove support for Go version `N-2`.

Official OpenTelemetry Collector distro binaries may be built with any supported Go version.

## Contributing

See the [Contributing Guide](CONTRIBUTING.md) for details.

<<<<<<< HEAD
Triagers ([@open-telemetry/collector-contrib-triagers](https://github.com/orgs/open-telemetry/teams/collector-contrib-triagers))

- [Andrzej Stencel](https://github.com/astencel-sumo), Sumo Logic
- [Benedikt Bongartz](https://github.com/frzifus), Red Hat
- [Goutham Veeramachaneni](https://github.com/gouthamve), Grafana
- Actively seeking contributors to triage issues

Emeritus Triagers:
=======
Here is a list of community roles with current and previous members:
>>>>>>> upstream/main

- Triagers ([@open-telemetry/collector-triagers](https://github.com/orgs/open-telemetry/teams/collector-triagers)):

<<<<<<< HEAD
Approvers ([@open-telemetry/collector-contrib-approvers](https://github.com/orgs/open-telemetry/teams/collector-contrib-approvers)):

- [Anthony Mirabella](https://github.com/Aneurysm9), AWS
- [Antoine Toulme](https://github.com/atoulme), Splunk
- [David Ashpole](https://github.com/dashpole), Google
- [Evan Bradley](https://github.com/evan-bradley), Dynatrace
- [Ruslan Kovalov](https://github.com/kovrus), Grafana Labs
- [Sean Marciniak](https://github.com/MovieStoreGuy), Atlassian
- [Ziqi Zhao](https://github.com/fatsheep9146), Alibaba

Emeritus Approvers:
- [Przemek Maciolek](https://github.com/pmm-sumo)

Maintainers ([@open-telemetry/collector-contrib-maintainer](https://github.com/orgs/open-telemetry/teams/collector-contrib-maintainer)):

- [Alex Boten](https://github.com/codeboten), Lightstep
- [Bogdan Drutu](https://github.com/bogdandrutu), Splunk
- [Daniel Jaglowski](https://github.com/djaglowski), observIQ
- [Dmitrii Anoshin](https://github.com/dmitryax), Splunk
- [Juraci Paixão Kröhling](https://github.com/jpkrohling), Grafana Labs
- [Pablo Baeyens](https://github.com/mx-psi), DataDog
- [Tyler Helmuth](https://github.com/TylerHelmuth), Honeycomb

Emeritus Maintainers
- [Tigran Najaryan](https://github.com/tigrannajaryan), Splunk
=======
  - Actively seeking contributors to triage issues

- Emeritus Triagers:

   - [Andrew Hsu](https://github.com/andrewhsu), Lightstep
   - [Alolita Sharma](https://github.com/alolita), Apple
   - [Punya Biswal](https://github.com/punya), Google
   - [Steve Flanders](https://github.com/flands), Splunk

- Approvers ([@open-telemetry/collector-approvers](https://github.com/orgs/open-telemetry/teams/collector-approvers)):
>>>>>>> upstream/main

   - [Anthony Mirabella](https://github.com/Aneurysm9), AWS
   - [Daniel Jaglowski](https://github.com/djaglowski), observIQ
   - [Juraci Paixão Kröhling](https://github.com/jpkrohling), Grafana Labs
   - [Pablo Baeyens](https://github.com/mx-psi), DataDog

- Emeritus Approvers:

   - [James Bebbington](https://github.com/james-bebbington), Google
   - [Jay Camp](https://github.com/jrcamp), Splunk
   - [Nail Islamov](https://github.com/nilebox), Google
   - [Owais Lone](https://github.com/owais), Splunk
   - [Rahul Patel](https://github.com/rghetia), Google
   - [Steven Karis](https://github.com/sjkaris), Splunk
   - [Yang Song](https://github.com/songy23), Google

- Maintainers ([@open-telemetry/collector-maintainers](https://github.com/orgs/open-telemetry/teams/collector-maintainers)):

   - [Alex Boten](https://github.com/codeboten), Lightstep
   - [Bogdan Drutu](https://github.com/BogdanDrutu), Splunk
   - [Dmitrii Anoshin](https://github.com/dmitryax), Splunk

- Emeritus Maintainers:

   - [Paulo Janotti](https://github.com/pjanotti), Splunk
   - [Tigran Najaryan](https://github.com/tigrannajaryan), Splunk

Learn more about roles in [Community membership](https://github.com/open-telemetry/community/blob/main/community-membership.md).
In addition to what is described at the organization-level, the SIG Collector requires all core approvers to take part in rotating
the role of the [release manager](./docs/release.md#release-manager).

## PRs and Reviews

When creating a PR please following the process [described
here](https://github.com/open-telemetry/opentelemetry-collector/blob/main/CONTRIBUTING.md#how-to-structure-prs-to-get-expedient-reviews).

News PRs will be automatically associated with the reviewers based on
[CODEOWNERS](.github/CODEOWNERS). PRs will be also automatically assigned to one of the
maintainers or approvers for facilitation.

The facilitator is responsible for helping the PR author and reviewers to make progress
or if progress cannot be made for closing the PR.

If the reviewers do not have approval rights the facilitator is also responsible
for the official approval that is required for the PR to be merged and if the facilitator
is a maintainer they are responsible for merging the PR as well.

The facilitator is not required to perform a thorough review, but they are encouraged to
enforce Collector best practices and consistency across the codebase and component
behavior. The facilitators will typically rely on codeowner's detailed review of the code
when making the final approval decision. 

We recommend maintainers and approvers to keep an eye on the
[project board](https://github.com/orgs/open-telemetry/projects/3). All newly created
PRs are automatically added to this board. (If you don't see the PR on the board you
may need to add it manually by setting the Project field in the PR view).
