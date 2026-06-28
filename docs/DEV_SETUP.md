# Developer setup

## Primary clone

Use **`/Users/fbelchi/github/agentic-travel`** as the main working copy. Day-to-day work happens on feature branches; **`main`** tracks `origin/main`.

```bash
cd /Users/fbelchi/github/agentic-travel
git fetch origin
git checkout main
git pull --ff-only
```

Build a CLI from its module directory, for example:

```bash
cd travelodge-cli && go build ./...
```

## Worktrees

This repo may use additional [git worktrees](https://git-scm.com/docs/git-worktree) for parallel agent or QA work (for example under `/private/tmp/`). **Only one worktree may have `main` checked out at a time.** If `git checkout main` fails with *main is already used by worktree*, either:

- finish or stash WIP in the other worktree and run `git worktree remove <path>`, or
- stay on a feature branch in the primary clone and `git fetch origin main` without checking out `main` locally.

Do not force-push. Stash or commit unpushed work before removing a worktree.

## Recovering WIP

Stashes from worktree cleanup are listed with `git stash list`. Pop the relevant entry when continuing that work on a feature branch.
