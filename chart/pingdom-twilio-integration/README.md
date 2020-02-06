# Pingdom Twilio integration

## Introduction
pingdom-twilio-integration sends messages to contact groups defined in configuration.

## Prerequisites

-  kubernetes 1.11+

## Installing the Chart

To install the chart with the release name `pingdom-twilio-integration`:

```bash
$ helm install docplanner/pingdom-twilio-integration --name pingdom-twilio-integration --values=my-values.yaml
```

## Uninstalling the Chart

To uninstall/delete the `pingdom-twilio-integration` deployment:

```bash
$ helm delete pingdom-twilio-integration
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the pingdom-twilio-integration chart and their default values.

| Parameter                                  | Description                               | Default                            |
| ------------------------------------------ | ----------------------------------------- | ---------------------------------- |
| `replicaCount` | replica count | `1`|
| `updateStrategy` | Deployment update strategy | `type: RollingUpdate` |
| `image.repository` | Docker image repo | `docplanner/pingdom-twilio-integration`|
| `image.tag` | Docker image tag | `latest`|
| `image.pullPolicy` | Docker image pull policy| `IfNotPresent`|
| `resources` | Resources | `{}`|
| `nodeSelector` | NodeSelector | `{}`|
| `tolerations` | Tolerations | `[]`|
| `affinity` | Affinity | `{}`|
| `ingress.enabled` | Ingress enabled | `false` |
| `ingress.annotations` | Ingress annotations | `{}` |
| `ingress.path` | Ingress path | `/` |
| `ingress.hosts` | Ingress hosts | `[]` |
| `ingress.tls` | Ingress TLS | `[]` |
| `config` | Configuration  | `` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```bash
$ helm install --name pingdom-twilio-integration --set ingress.tag=v1.0.0 docplanner/pingdom-twilio-integration
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart.