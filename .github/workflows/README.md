# GitHub Actions Workflows

## Naming

Workflow files are grouped by prefix:

| Prefix | Purpose |
| --- | --- |
| `pr-*` | Pull request automation (checks, labels, housekeeping) |
| `pr-check-*` | Standalone PR checks that cannot live inside `pr-checks-combined.yaml` (special runners, OIDC auth, or their own path filters) |
| `pr-waiting-response-*` | The `waiting-response` label machinery (see below) |
| `issue-*` | Issue automation |
| `milestone-*` | Milestone automation |
| `workflow-*` | Linting of the workflow files themselves (actionlint, yamllint) |
| `teamcity-*` | TeamCity (acceptance test infrastructure) configuration and triggers |
| `main-checks.yaml` | Tree-state CI checks re-run on every push to `main`/`release/**` (catches semantic merge conflicts); failures surface as commit checks, no PR machinery |

`build.yml` and `provider-release.yml` are HashiCorp Common Release Tooling (CRT)
files referenced by external release pipelines — do not rename them.

Prefixes describe what a workflow acts on, not how it is invoked. Four files are
reusable building blocks with no triggers of their own (`workflow_call` only):
`pr-save-artifacts.yaml`, `pr-comment-failure.yaml`,
`pr-comment-failure-outdated.yaml`, and `issue-remove-label.yaml`.

## The artifact relay pattern (why `pr-save-artifacts.yaml` exists)

Several workflows need to *modify* a PR (e.g. add the `waiting-response` label when
CI fails) — but the workflow that detects the condition is not allowed to do so:

1. **PR checks can't write.** Checks run on `pull_request`, and for PRs from forks
   the `GITHUB_TOKEN` is read-only, because the workflow executes untrusted code.
2. **The privileged side is blind.** Workflows triggered by `workflow_run` execute
   in the base repo with write permissions, but their event payload does not
   reliably identify the PR — `github.event.workflow_run.pull_requests` is empty
   for fork PRs.

So the two halves communicate through an artifact:

```
pull_request workflow (untrusted, read-only, knows the PR)
  └─ on failure: pr-save-artifacts.yaml uploads wr_actions/{ghowner,ghrepo,prnumber}.txt
       └─ workflow_run workflow (trusted, write token, blind)
            └─ downloads the artifact → labels that PR
```

Concretely:

- Every job in `pr-checks-combined.yaml` and `teamcity-test.yaml` calls
  `pr-save-artifacts.yaml` on failure.
- `pr-waiting-response-on-ci-fail.yaml` triggers on those workflows' **display
  names** (`workflow_run` matches names, not filenames), downloads the artifact,
  and adds the `waiting-response` label.
- The same pattern powers the review flow: `pr-reviewed.yaml` uploads the
  reviewer/state artifact, `pr-waiting-response-on-review.yaml` applies it.

**When adding a new PR check**: if its failure should mark the PR
`waiting-response`, the check must (a) call `pr-save-artifacts.yaml` on failure
and (b) have its display name added to the `workflow_run` list in
`pr-waiting-response-on-ci-fail.yaml`. One without the other does nothing.

## The CI failure comment

Each check in `pr-checks-combined.yaml` (and `teamcity-test.yaml`) has an `-on-fail` job calling
`pr-comment-failure.yaml` and an `-on-success` job calling
`pr-comment-failure-outdated.yaml`, both passing the same `check_name`. Failures
are collected into a single "CI Build Failures for `<sha>`" comment per commit
with per-check remediation guidance (the table in `pr-comment-failure.yaml`);
when a check later passes, its entry is struck through and the comment is
minimized once every entry is fixed. The `check_name` strings are the join key
between the two files and the guidance table - keep all three in sync.

**Security note**: artifact contents come from the untrusted side. Never trust
them for anything more dangerous than labeling — a malicious PR could upload an
arbitrary PR number.

## Renaming caveats

- `workflow_run` triggers reference workflow **display names** — renaming a
  workflow's `name:` silently breaks its consumers (`pr-waiting-response-on-ci-fail.yaml`,
  `pr-waiting-response-on-review.yaml`).
- Reusable workflows (`pr-save-artifacts.yaml`, `pr-comment-failure*.yaml`,
  `issue-remove-label.yaml`) are referenced by **file path** in `uses:` lines.
- `teamcity-test.yaml` has a self-referencing `paths:` trigger that must be
  updated when the file is renamed.

## The `/test` command

`teamcity-run-tests-on-comment.yaml` lets members of the `terraform-azure` team
(or the `TCTEST_ALLOWED_USERS` secret) trigger TeamCity acceptance test runs by
commenting `/test [-f] [-s service] <TestPrefix>` on a PR. It dispatches via
[katbyte/tctest](https://github.com/katbyte/tctest) with `POST_GITHUB_COMMENT=true`
and the PR head SHA as `TRACKING_ID`, and the TeamCity build posts the results
back as a PR comment (see `.teamcity/scripts/post_github_comment.sh`). Needs the
`TCTEST_SERVER`, `TCTEST_TOKEN_TC` and `GH_MEMEBERSHIP_CHECK_TOKEN` secrets.

## TeamCity labels

`teamcity-remove-label-on-commit.yaml` swaps the `teamcity-passed` /
`teamcity-failed` / `teamcity-new-failure` labels (applied by the TeamCity
build's GitHub comment step, see `.teamcity/scripts/post_github_comment.sh`)
for `teamcity-outdated` when new commits are pushed, so stale acceptance test
results are not trusted. The four labels must exist in the repository.
