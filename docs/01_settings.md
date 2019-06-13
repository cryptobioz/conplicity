# Settings


## Global options

* `--blacklist` allows you to blacklist volumes based on their name. Those volumes will not be backed up.
* `--whitelist` allows you to whitelist volumes based on their name. Only those volumes will be backed up.

## Manager options

```
Start Bivac backup manager

Usage:
  bivac manager [flags]

Flags:
      --agent.image string                        Agent's Docker image. [BIVAC_AGENT_IMAGE] (default "camptocamp/bivac:2.0")
      --cattle.accesskey string                   The Cattle access key. [CATTLE_ACCESS_KEY]
      --cattle.secretkey string                   The Cattle secret key. [CATTLE_SECRET_KEY]
      --cattle.url string                         The Cattle URL. [CATTLE_URL]
      --docker.endpoint string                    Docker endpoint. [BIVAC_DOCKER_ENDPOINT] (default "unix:///var/run/docker.sock")
  -h, --help                                      help for manager
      --kubernetes.agent-service-account string   Specify service account for agents. [KUBERNETES_AGENT_SERVICE_ACCOUNT]
      --kubernetes.all-namespaces                 Backup volumes of all namespaces. [KUBERNETES_ALL_NAMESPACES]
      --kubernetes.kubeconfig string              Path to your kuberconfig file. [KUBERNETES_KUBECONFIG]
      --kubernetes.namespace string               Namespace where you want to run Bivac. [KUBERNETES_NAMESPACE]
      --log.server string                         Manager's API address that will receive logs from agents. [BIVAC_LOG_SERVER]
  -o, --orchestrator string                       Orchestrator on which Bivac should connect to. [BIVAC_ORCHESTRATOR]
      --parallel.count int                        The count of agents to run in parallel [BIVAC_PARALLEL_COUNT] (default 2)
      --providers.config string                   Configuration file for providers. [BIVAC_PROVIDERS_CONFIG] (default "/providers-config.default.toml")
      --refresh.rate string                       The volume list refresh rate. [BIVAC_REFRESH_RATE] (default "10m")
      --restic.forget.args string                 Restic forget arguments. [RESTIC_FORGET_ARGS] (default "--group-by host --keep-daily 15 --prune")
      --retry.count int                           Retry to backup the volume if something goes wrong with Bivac. [BIVAC_RETRY_COUNT]
      --server.address string                     Address to bind on. [BIVAC_SERVER_ADDRESS] (default "0.0.0.0:8182")
      --server.psk string                         Pre-shared key. [BIVAC_SERVER_PSK]
  -r, --target.url string                         The target URL to push the backups to. [BIVAC_TARGET_URL]
      --whitelist.annotations                     Require pvc whitelist annotation [BIVAC_WHITELIST_ANNOTATION]

Global Flags:
  -b, --blacklist string   Do not backup blacklisted volumes. [BIVAC_BLACKLIST] [BIVAC_VOLUMES_BLACKLIST]
  -v, --verbose            Enable verbose output [BIVAC_VERBOSE]
  -w, --whitelist string   Only backup whitelisted volumes. [BIVAC_WHITELIST] [BIVAC_VOLUMES_WHITELIST]
```

* `--agent.image` is the Docker image that will be used by the Bivac manager to backup a volume. To avoid incompatibilities, it should be the same image as the manager.
* `--cattle.*` are the authentication parameters used by Rancher v1.6 to log in the cluster.
* `--docker.endpoint` is the endpoint used to connect to the Docker API.
* `--kubernetes.agent-service-account` is the Kubernetes Service Account that should be used by the Bivac agents. It might be useful if you have, for example, special security restrictions on your cluster.
* `--kubernetes.all-namespaces`
