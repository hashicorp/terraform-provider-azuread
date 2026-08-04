# GitHub Actions Workflows

## Naming

Workflow files are grouped by prefix:

| Prefix | Purpose |
|---|---|
| `pr-*` | Pull request automation (checks, labels, housekeeping) |
| `pr-waiting-response-*` | The `waiting-response` label machinery (see below) |
| `issue-*` | Issue automation |
| `milestone-*` | Milestone automation |

`build.yml` and `provider-release.yml` are HashiCorp Common Release Tooling (CRT)
files referenced by external release pipelines — do not rename them.

Prefixes describe what a workflow acts on, not how it is invoked. Three files are
reusable building blocks with no triggers of their own (`workflow_call` only):
`pr-save-artifacts.yaml`, `pr-comment-failure.yaml`, and `issue-remove-label.yaml`.

## The artifact relay pattern (why `pr-save-artifacts.yaml` exists)

Several workflows need to *modify* a PR (e.g. add the `waiting-response` label) —
but the workflow that detects the condition is not allowed to do so:

1. **PR checks can't write.** Checks run on `pull_request`, and for PRs from forks
   the `GITHUB_TOKEN` is read-only, because the workflow executes untrusted code.
2. **The privileged side is blind.** Workflows triggered by `workflow_run` execute
   in the base repo with write permissions, but their event payload does not
   reliably identify the PR — `github.event.workflow_run.pull_requests` is empty
   for fork PRs.

So the two halves communicate through an artifact:

```
pull_request workflow (untrusted, read-only, knows the PR)
  └─ uploads wr_actions/{ghowner,ghrepo,prnumber}.txt as an artifact
       └─ workflow_run workflow (trusted, write token, blind)
            └─ downloads the artifact → labels that PR
```

Concretely:

- The review flow: `pr-reviewed.yaml` uploads the reviewer/state artifact, and
  `pr-waiting-response-on-review.yaml` (triggered by `workflow_run` on its
  **display name**, "Pull Request Reviewed") downloads it and adds or removes
  the `waiting-response` label.
- Every job in `pr-checks-combined.yaml` calls `pr-save-artifacts.yaml` on
  failure, so a privileged `workflow_run` workflow can identify the failing PR.

**Security note**: artifact contents come from the untrusted side. Never trust
them for anything more dangerous than labeling — a malicious PR could upload an
arbitrary PR number.

## Renaming caveats

- `workflow_run` triggers reference workflow **display names** — renaming a
  workflow's `name:` silently breaks its consumers
  (`pr-waiting-response-on-review.yaml` depends on "Pull Request Reviewed").
- Reusable workflows (`pr-save-artifacts.yaml`, `pr-comment-failure.yaml`,
  `issue-remove-label.yaml`) are referenced by **file path** in `uses:` lines.
