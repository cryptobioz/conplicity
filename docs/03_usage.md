# Usage

To help you to manage your backups, Bivac comes with a CLI client.

## Backup a volume

By default, Bivac will backup your volumes every day but you may want to manually force a backup of a volume.

To do so, you must first list the volume to retrieve the one you want to backup:

```bash
$ bivac volumes
ID              Name            Hostname      Mountpoint      LastBackupDate          LastBackupStatus        Backing up
mysql           mysql           testing       /var/lib/mysql  2019-06-13 01:33:44     Success                 false
ssh_config      ssh_config      testing       /etc/ssh        2019-06-13 01:43:12     Success                 false
```

Let's say you want to backup the volume `ssh_config`:

```bash
$ bivac backup ssh_config
Backing up `ssh_config'...
ID: ssh_config
Name: ssh_sshconfig
Mountpoint: /etc/ssh
Backup date: 2019-06-13 09:35:38
Backup status: Success
Logs:

        testInit
        init
        backup          [0]
                        Files:           0 new,     0 changed,    11 unmodified
                        Dirs:            0 new,     1 changed,     0 unmodified
                        Added to the repo: 702 B

                        processed 11 files, 299.375 KiB in 0:01
                        snapshot 1c21ee5b saved

        forget          [0] Applying Policy: keep the last 15 daily snapshots
                        snapshots for (host [testing]):

                        keep 15 snapshots:
                        ID        Time                 Host                                       Tags        Reasons         Paths
                        --------------------------------------------------------------------------------------------------------------
                        1c21ee5b  2019-06-13 09:35:32  testing                                                daily snapshot  /etc/ssh
                        --------------------------------------------------------------------------------------------------------------
                        1 snapshots

                        repository contains 18 packs (44 blobs) with 317.585 KiB
                        processed 44 blobs: 0 duplicate blobs, 0B duplicate
                        load all snapshots
                        find data that is still in use for 15 snapshots
                        [0:00] 100.00%  15 / 15 snapshots

                        found 42 of 44 data blobs still in use, removing 2 blobs
                        will remove 0 invalid files
                        will delete 1 packs and rewrite 0 packs, this frees 763B
                        counting files in repo
                        [0:00] 100.00%  17 / 17 packs

                        finding old index files
                        saved new indexes as [f027febd]
                        remove 2 old index files
                        [0:00] 100.00%  1 / 1 packs deleted

                        done

```

Congratulations, you've successfully backed up your volume! 🍾 🎉

## Restore a volume

To restore a backed up volume, you can run the following command which will restore the latest snapshot in the volume:

```bash
$ bivac restore canary
Restoring `canary'...
ID: canary
Name: canary
Mountpoint: /var/lib/docker/volumes/canary/_data
Backup date: 2019-06-13 07:56:36
Backup status: Success
Logs:
         	
restore  	[0] restoring <Snapshot 15583d4b of [/var/lib/docker/volumes/canary/_data] at 2019-06-13 07:56:13.905600644 +0000 UTC by root@testing> to /var/lib/docker/volumes/canary/_data/h3bf5TfCxKtisKYF
snapshots	[0] [{"time":"2019-06-13T07:56:13.905600644Z","tree":"e6790a6cf2fd100d01b3bcac795c8787411b0879c85d60514f109403d26890bf","paths":["/var/lib/docker/volumes/canary/_data"],"hostname":"testing","username":"root","id":"15583d4b11605ec552be08fd1fd76d7549aefa0104ab4111f629737d5c7f7a17","short_id":"15583d4b"}]
```

## Manage a remote Restic repository

If you want to list volume's snapshots or retrieve some stats, you will have to use Restic and Bivac provides a good abstraction to do it.

Let's say you have volume called `canary` and you want to list the associate snapshots, then you'll simply run:

```bash
$ bivac restic --volume canary snapshots
ID        Time                 Host                      Tags        Paths
-------------------------------------------------------------------------------------------
9d22678e  2019-01-13 03:35:01  canary                                /mnt/geoserver_geodata
-------------------------------------------------------------------------------------------
1 snapshots

```

In case, you'd like to run a more complex command, you must use `--` as follow:

```bash
$ bivac restic --volume canary -- forget --prune --keep-daily 15
```

## Troubleshooting

### _My backup failed because the remote repository is locked._

The first thing to do is to check the date and the user who created the lock. From these informations, you should be able to determine if the lock is "legit" (a backup is running) or if it's a remnant of a forgotten backup. If you think it's safe to remove it, then you can run:

```bash
$ bivac backup [VOLUME_ID] --force
```

With the option `--force`, Bivac will unlock the Restic repository before doing a backup.
