# Introduction

Bivac is tool which allows you to backup your containers volumes deployed on Docker Engine, Cattle or Kubernetes.

## Architecture

### Volumes management

The Bivac manager starts a goroutine which constantly refresh a _volume-to-backup_ list. Every 10 minutes (`$BIVAC_REFRESH_RATE`), this list is updated by retrieving the volume list from the orchestrator's API.
If in the process, a new volume is discovered, Bivac will try to open its remote Restic repository and retrieve the last backup date.

### Backup process

The backup process is divided in two steps.

The first step is to detect the data type stored in the targeted volume. For example, if the volume is a PostgreSQL datadir, then we shouldn't backup the raw files but rather a dump of the database.

To address this issue, Bivac detects a volume type based on the providers configuration file ([providers-config.default.toml](https://raw.githubusercontent.com/camptocamp/bivac/master/providers-config.default.toml)) and can run _pre-commands_ in the container which mounts a volume before backing it up. Keeping the same example, Bivac will execute the pre-command `mkdir -p $volume/backups && pg_dumpall --clean -Upostgres > $volume/backups/all.sql` and then, only backs up the subdirectory `/backups`.

The second step is the _putting-local-data-to-a-remote-location_ process.

If and once the pre-commands have been executed, the manager deploys a Bivac agent container _via_ the orchestrator's API. This container mounts the volume to back up and execute the following operations:

* Initialize a remote Restic repository if it doesn't exist.
* Backup the files.
* Cleanup the old snapshots.

Once it's done, the optional _post-commands_ are executed.
