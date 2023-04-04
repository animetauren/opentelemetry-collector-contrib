# Contributing

If you would like to contribute please read OpenTelemetry Collector [contributing
guidelines](https://github.com/open-telemetry/opentelemetry-collector/blob/main/CONTRIBUTING.md) before you begin your
work.

<<<<<<< HEAD
## Pull-request title
=======
## Target audiences

The OpenTelemetry Collector has two main target audiences:

1. End-users, aiming to use an OpenTelemetry Collector binary.
1. Collector distributions, consuming the APIs exposed by the OpenTelemetry core repository. Distributions can be an
official OpenTelemetry community project, such as the OpenTelemetry Collector "core" and "contrib", or external
distributions, such as other open-source projects building on top of the Collector or vendor-specific distributions.

### End-users

End-users are the target audience for our binary distributions, as made available via the
[opentelemetry-collector-releases](https://github.com/open-telemetry/opentelemetry-collector-releases) repository. To
them, stability in the behavior is important, be it runtime or configuration. They are more numerous and harder to get
in touch with, making our changes to the collector more disruptive to them than to other audiences. As a general rule,
whenever you are developing OpenTelemetry Collector components (extensions, receivers, processors, exporters), you
should have end-users' interests in mind. Similarly, changes to code within packages like `config` will have an impact
on this audience. Make sure to cause minimal disruption when doing changes here.

### Collector distributions

In this capacity, the opentelemetry-collector repository's public Go types, functions, and interfaces act as an API for
other projects. In addition to the end-user aspect mentioned above, this audience also cares about API compatibility,
making them susceptible to our refactorings, even though such changes wouldn't cause any impact to end-users. See the
"Breaking changes" in this document for more information on how to perform changes affecting this audience.

This audience might use tools like the
[opentelemetry-collector-builder](https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder) as
part of their delivery pipeline. Be mindful that changes there might cause disruption to this audience.

## How to structure PRs to get expedient reviews?
>>>>>>> upstream/main

The title for your pull-request should contain the component type and name in brackets, plus a short statement for your
change. For instance:

    [processor/tailsampling] fix AND policy

<<<<<<< HEAD
## Changelog

Pull requests that contain user-facing changes will require a changelog entry. Keep in mind the following types of users:
1. Those who are consuming the telemetry exported from the collector
2. Those who are deploying or otherwise managing the collector or its configuration
3. Those who are depending on APIs exported from collector packages
4. Those who are contributing to the repository
=======
Components comprise of exporters, extensions, receivers, and processors. The key criteria to implementing a component is to:

* Implement the `component.Component` interface
* Provide a configuration structure which defines the configuration of the component
* Provide the implementation which performs the component operation

For more details on components, see the [Adding New Components](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/CONTRIBUTING.md#adding-new-components) document and the tutorial [Building a Trace Receiver](https://opentelemetry.io/docs/collector/trace-receiver/) which provides a detailed example of building a component.

When adding a new component to the OpenTelemetry Collector, ensure that any configuration structs used by the component include fields with the `configopaque.String` type for sensitive data. This ensures that the data is masked when serialized to prevent accidental exposure.

When submitting a component to the community, consider breaking it down into separate PRs as follows:

* First PR should include the overall structure of the new component:
  * Readme, configuration, and factory implementation usually using the helper
    factory structs.
  * This PR is usually trivial to review, so the size limit does not apply to
    it.
* Second PR should include the concrete implementation of the component. If the
  size of this PR is larger than the recommended size consider splitting it in
  multiple PRs.
* Last PR should enable the new component and add it to the `otelcontribcol`
  binary by updating the `components.go` file. The component must be enabled
  only after sufficient testing, and there is enough confidence in the
  stability and quality of the component.
* Once a new component has been added to the executable, please add the component
  to the [OpenTelemetry.io registry](https://github.com/open-telemetry/opentelemetry.io#adding-a-project-to-the-opentelemetry-registry).
* intra-repository `replace` statements in `go.mod` files can be automatically inserted by running `make crosslink`. For more information
  on the `crosslink` tool see the README [here](https://github.com/open-telemetry/opentelemetry-go-build-tools/tree/main/crosslink).
>>>>>>> upstream/main

The changelog is primarily directed at the first three groups but it is sometimes appropriate to include important updates relevant only to the forth group.

If a changelog entry is not required, a maintainer or approver will add the `Skip Changelog` label to the pull request.

**Examples**

Changelog entry required:
- Changes to the configuration of the collector or any component
- Changes to the telemetry emitted from and/or processed by the collector
- Changes to the prerequisites or assumptions for running a collector
- Changes to an API exported by a collector package
- Meaningful changes to the performance of the collector

Judgement call:
- Major changes to documentation
- Major changes to tests or test frameworks
- Changes to developer tooling in the repo

No changelog entry:
- Typical documentation updates
- Refactorings with no meaningful change in functionality
- Most changes to tests
- Chores, such as enabling linters, or minor changes to the CI process

### Adding a Changelog Entry

The [CHANGELOG.md](./CHANGELOG.md) file in this repo is autogenerated from `.yaml` files in the `./.chloggen` directory.

Your pull-request should add a new `.yaml` file to this directory. The name of your file must be unique since the last release.

During the collector release process, all `./.chloggen/*.yaml` files are transcribed into `CHANGELOG.md` and then deleted.

**Recommended Steps**
1. Create an entry file using `make chlog-new`. This generates a file based on your current branch (e.g. `./.chloggen/my-branch.yaml`)
2. Fill in all fields in the new file
3. Run `make chlog-validate` to ensure the new file is valid
4. Commit and push the file

Alternately, copy `./.chloggen/TEMPLATE.yaml`, or just create your file from scratch.

## Portable Code

In order to ensure compatibility with different operating systems, code should be portable. Below are some guidelines to follow when writing portable code:

* Avoid using platform-specific libraries, features etc. Please opt for portable multi-platform solutions. 

* Avoid hard-coding platform-specific values. Use environment variables or configuration files for storing platform-specific values.

    For example, avoid using hard-coded file path
    ```
    filePath := "C:\Users\Bob\Documents\sampleData.csv"
    ```

    Instead environment variable or configuration file can be used.
    ```
    filePath := os.Getenv("DATA_FILE_PATH")
    ```
    or
    ```
    filePath := Configuration.Get("data_file_path")
    ```

* Be mindful of 
  - Standard file systems and file paths such as forward slashes (/) instead of backward slashes (\\) in Windows. Use the [`path/filepath` package](https://pkg.go.dev/path/filepath) when working with filepaths. 
  - Consistent line ending formats such as Unix (LF) or Windows (CRLF).

* Test your implementation thoroughly on different platforms if possible and fix any issues. 

With above guidelines, you can write code that is more portable and easier to maintain across different platforms. 

## Adding New Components

**Before** any code is written, [open an
issue](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/new?assignees=&labels=new+component&template=new_component.md&title=New%20component)
providing the following information:

* Who's the sponsor for your component. A sponsor is an approver who will be in charge of being the official reviewer of
  the code and become a code owner for the component. For vendor-specific components, it's good to have a volunteer
  sponsor. If you can't find one, we'll assign one in a round-robin fashion. A vendor-specific component directly interfaces
  with a vendor-specific API and is expected to be maintained by a representative of the same vendor. For non-vendor specific
  components, having a sponsor means that your use case has been validated.
* Some information about your component, such as the reasoning behind it, use-cases, telemetry data types supported, and
  anything else you think is relevant for us to make a decision about accepting the component.
* The configuration options your component will accept. This will help us understand what it does and have an idea of
  how the implementation might look like.

Components comprise of exporters, extensions, receivers, and processors. The key criteria to implementing a component is to:

* Implement the [component.Component](https://pkg.go.dev/go.opentelemetry.io/collector/component#Component) interface
* Provide a configuration structure which defines the configuration of the component
* Provide the implementation which performs the component operation

Familiarize yourself with the interface of the component that you want to write, and use existing implementations as reference.
[Building a Trace Receiver](https://opentelemetry.io/docs/collector/trace-receiver/) tutorial provides a detailed example of building a component.

*NOTICE:* The Collector is in Beta stage and as such the interfaces may undergo breaking changes. Component creators
must be available to update or review their components when such changes happen, otherwise the component will be
excluded from the default builds.

Generally, maintenance of components is the responsibility of contributors who authored them. If the original author or
some other contributor does not maintain the component it may be excluded from the default build. The component **will**
be excluded if it causes build problems, has failing tests or otherwise causes problems to the rest of the repository
and the rest of contributors.

- Create your component under the proper folder and use Go standard package naming recommendations.
- Use a boiler-plate Makefile that just references the one at top level, ie.: `include ../../Makefile.Common` - this
  allows you to build your component with required build configurations for the contrib repo while avoiding building the
  full repo during development.
- Each component has its own go.mod file. This allows custom builds of the collector to take a limited sets of
  dependencies - so run `go mod` commands as appropriate for your component.
- Implement the needed interface on your component by importing the appropriate component from the core repo. Follow the
  pattern of existing components regarding config and factory source files and tests.
- Implement your component as appropriate. Provide end-to-end tests (or mock backend/client as appropriate). Target is
  to get 80% or more of code coverage.
- Add a README.md on the root of your component describing its configuration and usage, likely referencing some of the
  yaml files used in the component tests. We also suggest that the yaml files used in tests have comments for all
  available configuration settings so users can copy and modify them as needed.
- Run `make crosslink` to update intra-repository dependencies. It will add a `replace` directive to `go.mod` file of every intra-repository dependant. This is necessary for your component to be included in the contrib executable.
- Add your component to `versions.yaml`.
- All components included in the distribution must be included in [`cmd/otelcontribcol/builder-config.yaml`](./cmd/otelcontribcol/builder-config.yaml) 
  and in the respective testing harnesses. To align with the test goal of the project, components must be testable within the framework defined within
  the folder. If a component can not be properly tested within the existing framework, it must increase the non testable
  components number with a comment within the PR explaining as to why it can not be tested.
- Add the sponsor for your component and yourself to a new line for your component in the
  [`.github/CODEOWNERS`](./.github/CODEOWNERS) file.
- Run `make generate-gh-issue-templates` to add your component to the dropdown list in the issue templates.

### Releasing New Components
After a component has been approved and merged, and has been enabled in `internal/components/`, it must be added to the
[OpenTelemetry Collector Contrib's release manifest.yaml](https://github.com/open-telemetry/opentelemetry-collector-releases/blob/main/distributions/otelcol-contrib/manifest.yaml)
to be included in the distributed otelcol-contrib binaries and docker images.

### Rotating sponsors

The following GitHub users are the currently available sponsors, either by being an approver or a maintainer of the contrib repository. The list is ordered based on a random sort of the list of sponsors done live at the Collector SIG meeting on 27-Apr-2022 and serves as the seed for the round-robin selection of sponsors, as described in the section above.

* [@dashpole](https://github.com/dashpole)
* [@TylerHelmuth](https://github.com/TylerHelmuth)
* [@djaglowski](https://github.com/djaglowski)
* [@codeboten](https://github.com/codeboten)
* [@Aneurysm9](https://github.com/Aneurysm9)
* [@mx-psi](https://github.com/mx-psi)
* [@dmitryax](https://github.com/dmitryax)
* [@evan-bradley](https://github.com/evan-bradley)
* [@MovieStoreGuy](https://github.com/MovieStoreGuy)
* [@bogdandrutu](https://github.com/bogdandrutu)
* [@jpkrohling](https://github.com/jpkrohling)

Whenever a sponsor is picked from the top of this list, please move them to the bottom.

## Adding metrics to existing receivers
Following these steps for contributing additional metrics to existing receivers.
 - Read instructions [here](https://github.com/open-telemetry/opentelemetry-collector/blob/main/CONTRIBUTING.md#fork) on how to
   fork, build and create PRs. The only difference is to change repository name from `opentelemetry-collector` to `opentelemetry-collector-contrib`
 - Edit `metadata.yaml` of your metrics receiver to add new metrics, e.g.: `redisreceiver/metadata.yaml`
 - To generate new metrics on top of this updated YAML file.
   - Run `cd receiver/redisreceiver`
   - Run `go generate ./...`
- Review the changed files and merge the changes into your forked repo.
- Create PR from Github web console following the instructions above.

## General Recommendations
Below are some recommendations that apply to typical components. These are not rigid rules and there are exceptions but
in general try to follow them.

- Avoid introducing batching, retries or worker pools directly on receivers and exporters. Typically, these are general
  cases that can be better handled via processors (that also can be reused by other receivers and exporters).
- When implementing exporters try to leverage the exporter helpers from the core repo, see [exporterhelper
  package](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/exporterhelper). This will
  ensure that the exporter provides [zPages](https://opencensus.io/zpages/) and a standard set of metrics.
- `replace` statements in `go.mod` files can be automatically inserted by running `make crosslink`. For more information
  on the `crosslink` tool see the README [here](https://github.com/open-telemetry/opentelemetry-go-build-tools/tree/main/crosslink).

## Issue Triaging

To help provide a consistent process for seeing issues through to completion, this section details some guidelines and
definitions to keep in mind when triaging issues.

### Roles

Determining the root cause of issues is a shared responsibility between those with triager permissions, code owners,
OpenTelemetry community members, issue authors, and anyone else who would like to contribute.

#### Triagers

Contributors with [triager](https://github.com/open-telemetry/opentelemetry-collector-contrib/#contributing) permissions can help move
issues along by adding missing component labels, which help organize issues and trigger automations to notify code owners. They can
also use their familiarity with the Collector and its components to investigate issues themselves. Alternatively, they may point issue
authors to another resource or someone else who may know more.

#### Code Owners

In many cases, the code owners for an issue are the best resource to help determine the root cause of a bug or whether an enhancement
is fit to be added to a component. Code owners will be notified by repository automations when:

- a component label is added to an issue
- an issue is opened
- the issue becomes stale

Code owners may not have triager permissions on the repository,
so they can help triage through investigation and by participating in discussions. They can also help organize issues by
[adding labels via comments](#adding-labels-via-comments).

#### Community Members

Community members or interested parties are welcome to help triage issues by investigating the root cause of bugs, adding input for
features they would like to see, or participating in design discussions.

### Triage process

Triaging an issue requires getting the issue into a state where there is enough information available on the issue or understanding
between the involved parties to allow work to begin or for the issue to be closed. Facilitating this may involve, but is not limited to:

- Determining whether the issue is related to the code or documentation, or whether the issue can be resolved without any changes.
- Ensuring that a bug can be reproduced, and if possible, the behavior can be traced back to the offending code or documentation.
- Determining whether a feature request belongs in a component, should be accomplished through other means, or isn't appropriate for a component at this time.
- Guiding any interested parties to another person or resource that may be more knowledgeable about an issue.
- Suggesting an issue for discussion at a SIG meeting if a synchronous discussion would be more productive.

#### Issue assignment

Issues are assigned for someone to work on by a triager when someone volunteers to work on an issue. Assignment is intended to prevent duplicate work by making it visible who is
working on a particular task. A person who is assigned to the issue may be assigned to help triage the issue and implement it, or can be assigned after the issue has already been
triaged and is ready for work. If someone who is assigned to an issue is no longer able to work on it, they may request to be unassigned from the issue.

### Label Definitions

| Label                | When to apply                                                                                                                                                                                           |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `bug`                | Something that is advertised or intended to work isn't working as expected.                                                                                                                             |
| `enhancement`        | Something that isn't an advertised feature that would be useful to users or maintainers.                                                                                                                |
| `flaky test`         | A test unexpectedly failed during CI, showing that there is a problem with the tests or test setup that is causing the tests to intermittently fail.                                                    |
| `good first issue`   | Implementing this issue would not require specialized or in-depth knowledge about the component and is ideal for a new or first-time contributor to take.                                               |
| `help wanted`        | The code owners for this component do not expect to have time to work on it soon, and would welcome help from contributors.                                                                             |
| `needs discussion`   | This issue needs more input from the maintainers or community before work can be started.                                                                                                               |
| `needs triage`       | This label is added automatically, and can be removed when a triager or code owner deems that an issue is either ready for work or should not need any work.                                            |
| `waiting for author` | Can be applied when input is required from the author before the issue can move any further.                                                                                                            |
| `priority:p0`        | A critical security vulnerability or Collector panic using a default or common configuration unrelated to a specific component.                                                                         |
| `priority:p1`        | An urgent issue that should be worked on quickly, before most other issues.                                                                                                                             |
| `priority:p2`        | A standard bug or enhancement.                                                                                                                                                                          |
| `priority:p3`        | A technical improvement, lower priority bug, or other minor issue. Generally something that is considered a "nice to have."                                                                               |
| `release:blocker`    | This issue must be resolved before the next Collector version can be released.                                                                                                                          |
| `Sponsor Needed`     | A new component has been proposed, but implementation is not ready to begin. This can be because a sponsor has not yet been decided, or because some details on the component still need to be decided. |
| `Accepted Component` | A sponsor has elected to take on a component and implementation is ready to begin.                                                                                                                      |
| `Vendor Specific Component` | This should be applied to any component proposal where the functionality for the component is particular to a vendor. |

### Adding Labels via Comments

In order to facilitate proper label usage and to empower Code Owners, you are able to add labels to issues via comments. To add a label through a comment, post a new comment on an issue starting with `/label`, followed by a space-separated list of your desired labels. Supported labels come from the table below, or correspond to a component defined in the [CODEOWNERS file](.github/CODEOWNERS).

The following general labels are supported:

| Label                | Label in Comment     |
|----------------------|----------------------|
| `good first issue`   | `good-first-issue`   |
| `help wanted`        | `help-wanted`        |
| `needs discussion`   | `needs-discussion`   |
| `needs triage`       | `needs-triage`       |
| `waiting for author` | `waiting-for-author` |

To delete a label, prepend the label with `-`. Note that you must make a new comment to modify labels; you cannot edit an existing comment.

Example label comment:

```
/label receiver/prometheus help-wanted -exporter/prometheus
```

## Becoming a Code Owner

A Code Owner is responsible for a component within Collector Contrib, as indicated by the [CODEOWNERS file](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/.github/CODEOWNERS). That responsibility includes maintaining the component, responding to issues, and reviewing pull requests.

Sometimes a component may be in need of a new or additional Code Owner. A few reasons this situation may arise would be:
- The component was never assigned a Code Owner.
- A previous Code Owner stepped down.
- An existing Code Owner has become unresponsive. See [unmaintained stability status](https://github.com/open-telemetry/opentelemetry-collector#unmaintained).
- The existing Code Owners are actively looking for new Code Owners to help.

If you would like to help and become a Code Owner you must meet the following requirements:

<<<<<<< HEAD
1. [Be a member of the OpenTelemetry organization.](https://github.com/open-telemetry/community/blob/main/community-membership.md#member)
2. (Code Owner Discretion) It is best to have resolved an issue related to the component, contributed directly to the component, and/or review component PRs. How much interaction with the component is required before becoming a Code Owner is up to any existing Code Owners.
=======
1. [git](https://git-scm.com/)
2. [go](https://golang.org/) (version 1.18 and up)
3. [make](https://www.gnu.org/software/make/)
4. [docker](https://www.docker.com/)
>>>>>>> upstream/main

Code Ownership is ultimately up to the judgement of the existing Code Owners and Collector Contrib Maintainers. Meeting the above requirements is not a guarantee to be granted Code Ownership.

To become a Code Owner, open a PR with the CODEOWNERS file modified, adding your GitHub username to the component's row. Be sure to tag the existing Code Owners, if any, within the PR to ensure they receive a notification.

### Makefile Guidelines

When adding or modifying the `Makefile`'s in this repository, consider the following design guidelines.

Make targets are organized according to whether they apply to the entire repository, or only to an individual module.
The [Makefile](./Makefile) SHOULD contain "repo-level" targets. (i.e. targets that apply to the entire repo.)
Likewise, `Makefile.Common` SHOULD contain "module-level" targets. (i.e. targets that apply to one module at a time.)
Each module should have a `Makefile` at its root that includes `Makefile.Common`.

#### Module-level targets

<<<<<<< HEAD
Module-level targets SHOULD NOT act on nested modules. For example, running `make lint` at the root of the repo will
*only* evaluate code that is part of the `go.opentelemetry.io/collector` module. This excludes nested modules such as
`go.opentelemetry.io/collector/component`.

Each module-level target SHOULD have a corresponding repo-level target. For example, `make golint` will run `make lint`
in each module. In this way, the entire repository is covered. The root `Makefile` contains some "for each module" targets
that can wrap a module-level target into a repo-level target.

#### Repo-level targets
=======
```shell
$ make
```

## Creating a PR
>>>>>>> upstream/main

Whenever reasonable, targets SHOULD be implemented as module-level targets (and wrapped with a repo-level target).
However, there are many valid justifications for implementing a standalone repo-level target.

1. The target naturally applies to the repo as a whole. (e.g. Building the collector.)
2. Interaction between modules would be problematic.
3. A necessary tool does not provide a mechanism for scoping its application. (e.g. `porto` cannot be limited to a specific module.)
4. The "for each module" pattern would result in incomplete coverage of the codebase. (e.g. A target that scans all file, not just `.go` files.)

<<<<<<< HEAD
#### Default targets

The default module-level target (i.e. running `make` in the context of an individual module), should run a substantial set of module-level
targets for an individual module. Ideally, this would include *all* module-level targets, but exceptions should be made if a particular
target would result in unacceptable latency in the local development loop.

The default repo-level target (i.e. running `make` at the root of the repo) should meaningfully validate the entire repo. This should include
running the default common target for each module as well as additional repo-level targets.
=======
### Commit Messages

Use descriptive commit messages. Here are [some recommendations](https://cbea.ms/git-commit/)
on how to write good commit messages.
When creating PRs GitHub will automatically copy commit messages into the PR description,
so it is a useful habit to write good commit messages before the PR is created.
Also, unless you actually want to tell a story with multiple commits make sure to squash
into a single commit before creating the PR.

When maintainers merge PRs with multiple commits, they will be squashed and GitHub will
concatenate all commit messages right before you hit the "Confirm squash and merge"
button. Maintainers must make sure to edit this concatenated message to make it right before merging.
In some cases, if the commit messages are lacking the easiest approach to have at
least something useful is copy/pasting the PR description into the commit message box
before merging (but see above paragraph about writing good commit messages in the first place).

## General Notes

This project uses Go 1.18.* and [Github Actions.](https://github.com/features/actions)

It is recommended to run `make gofmt all` before submitting your PR

The dependencies are managed with `go mod` if you work with the sources under your
`$GOPATH` you need to set the environment variable `GO111MODULE=on`.

## Coding Guidelines

Although OpenTelemetry project as a whole is still in Alpha stage we consider
OpenTelemetry Collector to be close to production quality and the quality bar
for contributions is set accordingly. Contributions must have readable code written
with maintainability in mind (if in doubt check [Effective Go](https://golang.org/doc/effective_go.html)
for coding advice). The code must adhere to the following robustness principles that
are important for software that runs autonomously and continuously without direct
interaction with a human (such as this Collector).

### Naming convention

To keep naming patterns consistent across the project, naming patterns are enforced to make intent clear by:

- Methods that return a variable that uses the zero value or values provided via the method MUST have the prefix `New`. For example:
  - `func NewKinesisExporter(kpl aws.KinesisProducerLibrary)` allocates a variable that uses
    the variables passed on creation.
  - `func NewKeyValueBuilder()` SHOULD allocate internal variables to a safe zero value.
- Methods that return a variable that uses non zero value(s) that impacts business logic MUST use the prefix `NewDefault`. For example:
  - `func NewDefaultKinesisConfig()` would return a configuration that is the suggested default
    and can be updated without concern of a race condition.
- Methods that act upon an input variable MUST have a signature that reflect concisely the logic being done. For example:
  - `func FilterAttributes(attrs []Attribute, match func(attr Attribute) bool) []Attribute` MUST only filter attributes out of the passed input
    slice and return a new slice with values that `match` returns true. It may not do more work than what the method name implies, ie, it
    must not key a global history of all the slices that have been filtered.
- Methods that get the value of a field i.e. a getterMethod MUST use uppercase first alphabet and NOT `get` prefix. For example:
  - `func (p *Person) Name() string {return p.name} ` Name (with an uppercase N,exported) method is used here to get the value of the name field and not `getName`.The use of upper-case names for export provides the hook to discriminate the field from the method.
- Methods that set the value of a field i.e. a setterMethod MUST use a `set` prefix. For example:
  - `func (p *Person) SetName(newName string) {p.name = newName}` SetName method here sets the value of the name field.
- Variable assigned in a package's global scope that is preconfigured with a default set of values MUST use `Default` as the prefix. For example:
  - `var DefaultMarshallers = map[string]pdata.Marshallers{...}` is defined with an exporters package that allows for converting an encoding name,
    `zipkin`, and return the preconfigured marshaller to be used in the export process.
- Types that are specific to a signal MUST be worded with the signal used as an adjective, i.e. `SignalType`. For example:
  - `type TracesSink interface {...}`
- Types that deal with multiple signal types should use the relationship between the signals to describe the type, e.g. `SignalToSignalType` or `SignalAndSignalType`. For example:
  - `type TracesToTracesFunc func(...) ...`
- Functions dealing with specific signals or signal-specific types MUST be worded with the signal or type as a direct object, i.e. `VerbSignal`, or `VerbType` where `Type` is the full name of the type including the signal name. For example:
  - `func ConsumeTraces(...) {...}`
  - `func CreateTracesExport(...) {...}`
  - `func CreateTracesToTracesFunc(...) {...}`

### Recommended Libraries / Defaults

In order to simplify developing within the project, library recommendations have been set
and should be followed.

| Scenario 	 | Recommended                   	                | Rationale                                                                                                                  |
|------------|------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------|
| Hashing  	 | ["hashing/fnv"](https://pkg.go.dev/hash/fnv) 	 | The project adopted this as the default hashing method due to the efficiency and is reasonable for non cryptographic use 	 |
| Testing  	 | Use `t.Parallel()` where possible            	 | Enabling more test to be run in parallel will speed up the feedback process when working on the project.                 	 |


Within the project, there are some packages that are yet to follow the recommendations and are being address, however, any new code should adhere to the recommendations.


### Startup Error Handling

Verify configuration during startup and fail fast if the configuration is invalid.
This will bring the attention of a human to the problem as it is more typical for humans
to notice problems when the process is starting as opposed to problems that may arise
sometime (potentially long time) after process startup. Monitoring systems are likely
to automatically flag processes that exit with failure during startup, making it
easier to notice the problem. The Collector should print a reasonable log message to
explain the problem and exit with a non-zero code. It is acceptable to crash the process
during startup if there is no good way to exit cleanly but do your best to log and
exit cleanly with a process exit code.

### Propagate Errors to the Caller

Do not crash or exit outside the `main()` function, e.g. via `log.Fatal` or `os.Exit`,
even during startup. Instead, return detailed errors to be handled appropriately
by the caller. The code in packages other than `main` may be imported and used by
third-party applications, and they should have full control over error handling
and process termination.

### Do not Crash after Startup

Do not crash or exit the Collector process after the startup sequence is finished.
A running Collector typically contains data that is received but not yet exported further
(e.g. is stored in the queues and other processors). Crashing or exiting the Collector
process will result in losing this data since typically the receiver has
already acknowledged the receipt for this data and the senders of the data will
not send that data again.

### Bad Input Handling

Do not crash on bad input in receivers or elsewhere in the pipeline.
[Crash-only software](https://en.wikipedia.org/wiki/Crash-only_software)
is valid in certain cases; however, this is not a correct approach for Collector (except
during startup, see above). The reason is that many senders from which Collector
receives data have built-in automatic retries of the _same_ data if no
acknowledgment is received from the Collector. If you crash on bad input
chances are high that after the Collector is restarted it will see the same
data in the input and will crash again. This will likely result in infinite
crashing loop if you have automatic retries in place.

Typically bad input when detected in a receiver should be reported back to the
sender. If it is elsewhere in the pipeline it may be too late to send a response
to the sender (particularly in processors which are not synchronously processing
data). In either case it is recommended to keep a metric that counts bad input data.

### Error Handling and Retries

Be rigorous in error handling. Don't ignore errors. Think carefully about each
error and decide if it is a fatal problem or a transient problem that may go away
when retried. Fatal errors should be logged or recorded in an internal metric to
provide visibility to users of the Collector. For transient errors come up with a
retrying strategy and implement it. Typically you will
want to implement retries with some sort of exponential back-off strategy. For
connection or sending retries use jitter for back-off intervals to avoid overwhelming
your destination when network is restored or the destination is recovered.
[Exponential Backoff](https://github.com/cenkalti/backoff) is a good library that
provides all this functionality.

### Logging

Log your component startup and shutdown, including successful outcomes (but don't
overdo it, keep the number of success message to a minimum).
This can help to understand the context of failures if they occur elsewhere after
your code is successfully executed.

Use logging carefully for events that can happen frequently to avoid flooding
the logs. Avoid outputting logs per a received or processed data item since this can
amount to very large number of log entries (Collector is designed to process
many thousands of spans and metrics per second). For such high-frequency events
instead of logging consider adding an internal metric and increment it when
the event happens.

Make log message human readable and also include data that is needed for easier
understanding of what happened and in what context.

### Executing External Processes

The components should avoid executing arbitrary external processes with arbitrary command
line arguments based on user input, including input received from the network or input
read from the configuration file. Failure to follow this rule can result in arbitrary
remote code execution, compelled by malicious actors that can craft the input.

The following limitations are recommended:
- If an external process needs to be executed limit and hard-code the location where
  the executable file may be located, instead of allowing the input to dictate the
  full path to the executable.
- If possible limit the name of the executable file to be one from a hard-coded
  list defined at compile time.
- If command line arguments need to be passed to the process do not take the arguments
  from the user input directly. Instead, compose the command line arguments indirectly,
  if necessary, deriving the value from the user input. Limit as much as possible the
  possible space of values for command line arguments.

### Observability

Out of the box, your users should be able to observe the state of your component.
The collector exposes an OpenMetrics endpoint at `http://localhost:8888/metrics`
where your data will land.

When using the regular helpers, you should have some metrics added around key
events automatically. For instance, exporters should have `otelcol_exporter_sent_spans`
tracked without your exporter doing anything.

### Resource Usage

Limit usage of CPU, RAM or other resources that the code can use. Do not write code
that consumes resources in an uncontrolled manner. For example if you have a queue
that can contain unprocessed messages always limit the size of the queue unless you
have other ways to guarantee that the queue will be consumed faster than items are
added to it.

Performance test the code for both normal use-cases under acceptable load and also for
abnormal use-cases when the load exceeds acceptable many times. Ensure that
your code performs predictably under abnormal use. For example if the code
needs to process received data and cannot keep up with the receiving rate it is
not acceptable to keep allocating more memory for received data until the Collector
runs out of memory. Instead have protections for these situations, e.g. when hitting
resource limits drop the data and record the fact that it was dropped in a metric
that is exposed to users.

### Graceful Shutdown

Collector does not yet support graceful shutdown but we plan to add it. All components
must be ready to shutdown gracefully via `Shutdown()` function that all component
interfaces require. If components contain any temporary data they need to process
and export it out of the Collector before shutdown is completed. The shutdown process
will have a maximum allowed duration so put a limit on how long your shutdown
operation can take.

### Unit Tests

Cover important functionality with unit tests. We require that contributions
do not decrease overall code coverage of the codebase - this is aligned with our
goal to increase coverage over time. Keep track of execution time for your unit
tests and try to keep them as short as possible.

#### Testing Library Recommendations

To keep testing practices consistent across the project, it is advised to use these libraries under
these circumstances:

- For assertions to validate expectations, use `"github.com/stretchr/testify/assert"`
- For assertions that are required to continue the test, use `"github.com/stretchr/testify/require"`
- For mocking external resources, use `"github.com/stretchr/testify/mock"`
- For validating HTTP traffic interactions, `"net/http/httptest"`

### Integration Testing

Integration testing is encouraged throughout the project, container images can be used in order to facilitate
a local version. In their absence, it is strongly advised to mock the integration.

### Using CGO

Using CGO is prohibited due to the lack of portability and complexity
that comes with managing external libaries with different operating systems and configurations.
However, if the package MUST use CGO, this should be explicitly called out within the readme
with clear instructions on how to install the required libraries.
Furthermore, if your package requires CGO, it MUST be able to compile and operate in a no op mode
or report a warning back to the collector with a clear error saying CGO is required to work.

### Breaking changes

Whenever possible, we adhere to semver as our minimum standards. Even before v1, we strive to not break compatibility
without a good reason. Hence, when a change is known to cause a breaking change, it MUST be clearly marked in the
changelog, and SHOULD include a line instructing users how to move forward.

We also strive to perform breaking changes in two stages, deprecating it first (`vM.N`) and breaking it in a subsequent
version (`vM.N+1`).

- when we need to remove something, we MUST mark a feature as deprecated in one version, and MAY remove it in a
  subsequent one
- when renaming or refactoring types, functions, or attributes, we MUST create the new name and MUST deprecate the old
  one in one version (step 1), and MAY remove it in a subsequent version (step 2). For simple renames, the old name
  SHALL call the new one.
- when a feature is being replaced in favor of an existing one, we MUST mark a feature as deprecated in one version, and
  MAY remove it in a subsequent one.

Deprecation notice SHOULD contain a version starting from which the deprecation takes effect for tracking purposes. For
example, if `GetFoo` function is going to be deprecated in `v0.45.0` version, it gets the following godoc line:

```golang
package test

// Deprecated: [v0.45.0] Use MustDoFoo instead.
func DoFoo() {}
```

When deprecating a feature affecting end-users, consider first deprecating the feature in one version, then hiding it
behind a [feature
gate](https://github.com/open-telemetry/opentelemetry-collector/blob/6b5a3d08a96bfb41a5e121b34f592a1d5c6e0435/service/featuregate/)
in a subsequent version, and eventually removing it after yet another version. This is how it would look like, considering
that each of the following steps is done in a separate version:

1. Mark the feature as deprecated, add a short lived feature gate with the feature enabled by default
1. Change the feature gate to disable the feature by default, deprecating the gate at the same time
1. Remove the feature and the gate

#### Example #1 - Renaming a function

1. Current version `v0.N` has `func GetFoo() Bar`
1. We now decided that `GetBar` is a better name. As such, on `v0.N+1` we add a new `func GetBar() Bar` function,
   changing the existing `func GetFoo() Bar` to be an alias of the new function. Additionally, a log entry with a
   warning is added to the old function, along with an entry to the changelog.
1. On `v0.N+2`, we MAY remove `func GetFoo() Bar`.

#### Example #2 - Changing the return values of a function

1. Current version `v0.N` has `func GetFoo() Foo`
1. We now need to also return an error. We do it by creating a new function that will be equivalent to the existing one,
   so that current users can easily migrate to that: `func MustGetFoo() Foo`, which panics on errors. We release this in
   `v0.N+1`, deprecating the existing `func GetFoo() Foo` with it, adding an entry to the changelog and perhaps a log
   entry with a warning.
1. On `v0.N+2`, we change `func GetFoo() Foo` to `func GetFoo() (Foo, error)`.

#### Example #3 - Changing the arguments of a function

1. Current version `v0.N` has `func GetFoo() Foo`
1. We now decided to do something that might be blocking as part of `func GetFoo() Foo`, so, we start accepting a
   context: `func GetFooWithContext(context.Context) Foo`. We release this in `v0.N+1`, deprecating the existing `func
   GetFoo() Foo` with it, adding an entry to the changelog and perhaps a log entry with a warning. The existing `func
   GetFoo() Foo` is changed to call `func GetFooWithContext(context.Background()) Foo`.
1. On `v0.N+2`, we change `func GetFoo() Foo` to `func GetFoo(context.Context) Foo` if desired or remove it entirely if
   needed.

### Specification Tracking

The [OpenTelemetry Specification](https://github.com/open-telemetry/opentelemetry-specification) can be a rapidly
moving target at times.  While it may seem efficient to get an early start on implementing new features or
functionality under development in the specification, this can also lead to significant churn and a risk that
changes in the specification can result in breaking changes to the implementation.  For this reason it is the
policy of the Collector SIG to not implement, or accept implementations of, new or changed specification language
prior to inclusion in a stable release of the specification.

## Changelog

An entry into the changelog is required for the following reasons:

- Changes made to the behaviour of the component
- Changes to the configuration
- Changes to default settings
- New components being added

It is reasonable to omit an entry to the changelog under these circuimstances:

- Updating test to remove flakiness or improve coverage
- Updates to the CI/CD process

If there is some uncertainty with regards to if a changelog entry is needed, the recomendation is to create
an entry to in the event that the change is important to the project consumers.

### Adding a Changelog Entry

The [CHANGELOG.md](./CHANGELOG.md) file in this repo is autogenerated from `.yaml` files in the `./.chloggen` directory.

Your pull-request should add a new `.yaml` file to this directory. The name of your file must be unique since the last release.

During the collector release process, all `./chloggen/*.yaml` files are transcribed into `CHANGELOG.md` and then deleted.

**Recommended Steps**
1. Create an entry file using `make chlog-new`. This generates a file based on your current branch (e.g. `./.chloggen/my-branch.yaml`)
2. Fill in all fields in the new file
3. Run `make chlog-validate` to ensure the new file is valid
4. Commit and push the file

Alternately, copy `./.chloggen/TEMPLATE.yaml`, or just create your file from scratch.

## Release

See [release](docs/release.md) for details.

## Contributing Images
If you are adding any new images, please use [Excalidraw](https://excalidraw.com). It's a free and open source web application and doesn't require any account to get started. Once you've created the design, while exporting the image, make sure to tick **"Embed scene into exported file"** option. This allows the image to be imported in an editable format for other contributors later.

## Common Issues

Build fails due to dependency issues, e.g.

```sh
go: github.com/golangci/golangci-lint@v1.31.0 requires
	github.com/tommy-muehle/go-mnd@v1.3.1-0.20200224220436-e6f9a994e8fa: invalid pseudo-version: git fetch --unshallow -f origin in /root/go/pkg/mod/cache/vcs/053b1e985f53e43f78db2b3feaeb7e40a2ae482c92734ba3982ca463d5bf19ce: exit status 128:
	fatal: git fetch-pack: expected shallow list
 ```

`go env GOPROXY` should return `https://proxy.golang.org,direct`. If it does not, set it as an environment variable:

`export GOPROXY=https://proxy.golang.org,direct`

### Makefile Guidelines

When adding or modifying the `Makefile`'s in this repository, consider the following design guidelines.

Make targets are organized according to whether they apply to the entire repository, or only to an individual module.
The [Makefile](./Makefile) SHOULD contain "repo-level" targets. (i.e. targets that apply to the entire repo.)
Likewise, `Makefile.Common` SHOULD contain "module-level" targets. (i.e. targets that apply to one module at a time.)
Each module should have a `Makefile` at its root that includes `Makefile.Common`.

#### Module-level targets

Module-level targets SHOULD NOT act on nested modules. For example, running `make lint` at the root of the repo will
_only_ evaluate code that is part of the `go.opentelemetry.io/collector` module. This excludes nested modules such as
`go.opentelemetry.io/collector/component`.

Each module-level target SHOULD have a corresponding repo-level target. For example, `make golint` will run `make lint`
in each module. In this way, the entire repository is covered. The root `Makefile` contains some "for each module" targets
that can wrap a module-level target into a repo-level target.

#### Repo-level targets

Whenever reasonable, targets SHOULD be implemented as module-level targets (and wrapped with a repo-level target).
However, there are many valid justifications for implementing a standalone repo-level target.

1. The target naturally applies to the repo as a whole. (e.g. Building the collector.)
2. Interaction between modules would be problematic.
3. A necessary tool does not provide a mechanism for scoping its application. (e.g. `porto` cannot be limited to a specific module.)
4. The "for each module" pattern would result in incomplete coverage of the codebase. (e.g. A target that scans all file, not just `.go` files.)

#### Default targets

The default module-level target (i.e. running `make` in the context of an individual module), should run a substantial set of module-level
targets for an individual module. Ideally, this would include *all* module-level targets, but exceptions should be made if a particular
target would result in unacceptable latency in the local development loop.

The default repo-level target (i.e. running `make` at the root of the repo) should meaningfully validate the entire repo. This should include
running the default common target for each module as well as additional repo-level targets.

## Exceptions

While the rules in this and other documents in this repository is what we strive to follow, we acknowledge that rules may be
unfeasible to be applied in some situations. Exceptions to the rules
on this and other documents are acceptable if consensus can be obtained from approvers in the pull request they are proposed.
A reason for requesting the exception MUST be given in the pull request. Until unanimity is obtained, approvers and maintainers are
encouraged to discuss the issue at hand. If a consensus (unanimity) cannot be obtained, the maintainers are then tasked in getting a
decision using its regular means (voting, TC help, ...).
>>>>>>> upstream/main
