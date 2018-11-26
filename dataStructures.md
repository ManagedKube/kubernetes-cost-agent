Data Structures
==================


# Node

Kubectl command:
```
kubectl get --raw /api/v1/nodes | jq
```

```json
{
  "metadata": {
    "name": "gke-gar-1-pool-1-dce25db1-2mm5",
    "selfLink": "/api/v1/nodes/gke-gar-1-pool-1-dce25db1-2mm5",
    "uid": "bddd62ec-e164-11e8-a351-42010a800049",
    "resourceVersion": "6279550",
    "creationTimestamp": "2018-11-06T01:39:03Z",
    "labels": {
      "beta.kubernetes.io/arch": "amd64",
      "beta.kubernetes.io/fluentd-ds-ready": "true",
      "beta.kubernetes.io/instance-type": "n1-standard-1",
      "beta.kubernetes.io/os": "linux",
      "cloud.google.com/gke-nodepool": "pool-1",
      "cloud.google.com/gke-os-distribution": "cos",
      "cloud.google.com/gke-preemptible": "true",
      "failure-domain.beta.kubernetes.io/region": "us-central1",
      "failure-domain.beta.kubernetes.io/zone": "us-central1-c",
      "kubernetes.io/hostname": "gke-gar-1-pool-1-dce25db1-2mm5"
    },
    "annotations": {
      "node.alpha.kubernetes.io/ttl": "0",
      "volumes.kubernetes.io/controller-managed-attach-detach": "true"
    }
  },
  "spec": {
    "podCIDR": "10.48.10.0/24",
    "providerID": "gce://managedkube/us-central1-c/gke-gar-1-pool-1-dce25db1-2mm5",
    "externalID": "125028673476447880"
  },
  "status": {
    "capacity": {
      "cpu": "1",
      "ephemeral-storage": "98868448Ki",
      "hugepages-2Mi": "0",
      "memory": "3787616Ki",
      "pods": "110"
    },
    "allocatable": {
      "cpu": "940m",
      "ephemeral-storage": "47093746742",
      "hugepages-2Mi": "0",
      "memory": "2702176Ki",
      "pods": "110"
    },
    "conditions": [
      {
        "type": "FrequentUnregisterNetDevice",
        "status": "False",
        "lastHeartbeatTime": "2018-11-07T00:22:17Z",
        "lastTransitionTime": "2018-11-06T01:44:04Z",
        "reason": "UnregisterNetDevice",
        "message": "node is functioning properly"
      },
      {
        "type": "KernelDeadlock",
        "status": "False",
        "lastHeartbeatTime": "2018-11-07T00:22:17Z",
        "lastTransitionTime": "2018-11-06T01:39:02Z",
        "reason": "KernelHasNoDeadlock",
        "message": "kernel has no deadlock"
      },
      {
        "type": "NetworkUnavailable",
        "status": "False",
        "lastHeartbeatTime": "2018-11-06T01:39:16Z",
        "lastTransitionTime": "2018-11-06T01:39:16Z",
        "reason": "RouteCreated",
        "message": "RouteController created a route"
      },
      {
        "type": "OutOfDisk",
        "status": "False",
        "lastHeartbeatTime": "2018-11-07T00:22:40Z",
        "lastTransitionTime": "2018-11-06T01:39:03Z",
        "reason": "KubeletHasSufficientDisk",
        "message": "kubelet has sufficient disk space available"
      },
      {
        "type": "MemoryPressure",
        "status": "False",
        "lastHeartbeatTime": "2018-11-07T00:22:40Z",
        "lastTransitionTime": "2018-11-06T01:39:03Z",
        "reason": "KubeletHasSufficientMemory",
        "message": "kubelet has sufficient memory available"
      },
      {
        "type": "DiskPressure",
        "status": "False",
        "lastHeartbeatTime": "2018-11-07T00:22:40Z",
        "lastTransitionTime": "2018-11-06T01:39:03Z",
        "reason": "KubeletHasNoDiskPressure",
        "message": "kubelet has no disk pressure"
      },
      {
        "type": "PIDPressure",
        "status": "False",
        "lastHeartbeatTime": "2018-11-07T00:22:40Z",
        "lastTransitionTime": "2018-11-06T01:39:03Z",
        "reason": "KubeletHasSufficientPID",
        "message": "kubelet has sufficient PID available"
      },
      {
        "type": "Ready",
        "status": "True",
        "lastHeartbeatTime": "2018-11-07T00:22:40Z",
        "lastTransitionTime": "2018-11-06T01:39:23Z",
        "reason": "KubeletReady",
        "message": "kubelet is posting ready status. AppArmor enabled"
      }
    ],
    "addresses": [
      {
        "type": "InternalIP",
        "address": "10.128.0.4"
      },
      {
        "type": "ExternalIP",
        "address": "35.192.99.255"
      },
      {
        "type": "Hostname",
        "address": "gke-gar-1-pool-1-dce25db1-2mm5"
      }
    ],
    "daemonEndpoints": {
      "kubeletEndpoint": {
        "Port": 10250
      }
    },
    "nodeInfo": {
      "machineID": "591e5f514cbe63fcb11fc5ce2c340202",
      "systemUUID": "591E5F51-4CBE-63FC-B11F-C5CE2C340202",
      "bootID": "60b5dcea-8934-49e7-8d79-dea8d75a32f5",
      "kernelVersion": "4.14.56+",
      "osImage": "Container-Optimized OS from Google",
      "containerRuntimeVersion": "docker://17.3.2",
      "kubeletVersion": "v1.10.7-gke.2",
      "kubeProxyVersion": "v1.10.7-gke.2",
      "operatingSystem": "linux",
      "architecture": "amd64"
    },
    "images": [
      {
        "names": [
          "gcr.io/managedkube/kubernetes-cost-attribution@sha256:11ee06509513238919f5f1dabe2bbb170687ebbcbaa5ae10f32e6957ba3e90b5",
          "gcr.io/managedkube/kubernetes-cost-attribution:1.0.0"
        ],
        "sizeBytes": 927852951
      },
      {
        "names": [
          "gcr.io/stackdriver-agents/stackdriver-logging-agent@sha256:a33f69d0034fdce835a1eb7df8a051ea74323f3fc30d911bbd2e3f2aef09fc93",
          "gcr.io/stackdriver-agents/stackdriver-logging-agent:0.3-1.5.34-1-k8s-1"
        ],
        "sizeBytes": 554981103
      },
      {
        "names": [
          "k8s.gcr.io/node-problem-detector@sha256:f95cab985c26b2f46e9bd43283e0bfa88860c14e0fb0649266babe8b65e9eb2b",
          "k8s.gcr.io/node-problem-detector:v0.4.1"
        ],
        "sizeBytes": 286572743
      },
      {
        "names": [
          "grafana/grafana@sha256:263023526eff9a8875a7b9f33abb6cfff1f5057543f1dba7cb5822959f920dd9",
          "grafana/grafana:5.1.2"
        ],
        "sizeBytes": 238158725
      },
      {
        "names": [
          "k8s.gcr.io/fluentd-elasticsearch@sha256:b8c94527b489fb61d3d81ce5ad7f3ddbb7be71e9620a3a36e2bede2f2e487d73",
          "k8s.gcr.io/fluentd-elasticsearch:v2.0.4"
        ],
        "sizeBytes": 135716379
      },
      {
        "names": [
          "gcr.io/google-containers/fluentd-gcp-scaler@sha256:bfd8ffbadf5cbfc7cd0944f5c13aaa8da421e3ab2322d610e64c4d7de9424c29",
          "gcr.io/google-containers/fluentd-gcp-scaler:0.3"
        ],
        "sizeBytes": 115128950
      },
      {
        "names": [
          "gcr.io/google_containers/kube-proxy:v1.10.7-gke.2",
          "k8s.gcr.io/kube-proxy:v1.10.7-gke.2"
        ],
        "sizeBytes": 103121873
      },
      {
        "names": [
          "k8s.gcr.io/kubernetes-dashboard-amd64@sha256:dc4026c1b595435ef5527ca598e1e9c4343076926d7d62b365c44831395adbd0",
          "k8s.gcr.io/kubernetes-dashboard-amd64:v1.8.3"
        ],
        "sizeBytes": 102319441
      },
      {
        "names": [
          "k8s.gcr.io/event-exporter@sha256:12549978ffdbe1be958b98df3030038e25cc5dea81ccbca253a8be3781f28a0e",
          "k8s.gcr.io/event-exporter:v0.2.1"
        ],
        "sizeBytes": 94193305
      },
      {
        "names": [
          "k8s.gcr.io/kube-addon-manager@sha256:3519273916ba45cfc9b318448d4629819cb5fbccbb0822cce054dd8c1f68cb60",
          "k8s.gcr.io/kube-addon-manager:v8.6"
        ],
        "sizeBytes": 78384272
      },
      {
        "names": [
          "k8s.gcr.io/heapster-amd64@sha256:fc33c690a3a446de5abc24b048b88050810a58b9e4477fa763a43d7df029301a",
          "k8s.gcr.io/heapster-amd64:v1.5.3"
        ],
        "sizeBytes": 75318342
      },
      {
        "names": [
          "k8s.gcr.io/rescheduler@sha256:66a900b01c70d695e112d8fa7779255640aab77ccc31f2bb661e6c674fe0d162",
          "k8s.gcr.io/rescheduler:v0.3.1"
        ],
        "sizeBytes": 74659350
      },
      {
        "names": [
          "k8s.gcr.io/ingress-gce-glbc-amd64@sha256:31d36bbd9c44caffa135fc78cf0737266fcf25e3cf0cd1c2fcbfbc4f7309cc52",
          "k8s.gcr.io/ingress-gce-glbc-amd64:v1.1.1"
        ],
        "sizeBytes": 67801919
      },
      {
        "names": [
          "k8s.gcr.io/prometheus-to-sd@sha256:71b2389fc4931b1cc3bb27ba1873361c346950558dd2220beb02ab6b31d3a651",
          "k8s.gcr.io/prometheus-to-sd:v0.2.4"
        ],
        "sizeBytes": 58688412
      },
      {
        "names": [
          "gcr.io/google-containers/prometheus-to-sd@sha256:be220ec4a66275442f11d420033c106bb3502a3217a99c806eef3cf9858788a2",
          "gcr.io/google-containers/prometheus-to-sd:v0.2.3"
        ],
        "sizeBytes": 55342106
      },
      {
        "names": [
          "k8s.gcr.io/cpvpa-amd64@sha256:cfe7b0a11c9c8e18c87b1eb34fef9a7cbb8480a8da11fc2657f78dbf4739f869",
          "k8s.gcr.io/cpvpa-amd64:v0.6.0"
        ],
        "sizeBytes": 51785854
      },
      {
        "names": [
          "k8s.gcr.io/cluster-proportional-autoscaler-amd64@sha256:003f98d9f411ddfa6ff6d539196355e03ddd69fa4ed38c7ffb8fec6f729afe2d",
          "k8s.gcr.io/cluster-proportional-autoscaler-amd64:1.1.2-r2"
        ],
        "sizeBytes": 49648481
      },
      {
        "names": [
          "k8s.gcr.io/k8s-dns-kube-dns-amd64@sha256:b99fc3eee2a9f052f7eb4cc00f15eb12fc405fa41019baa2d6b79847ae7284a8",
          "k8s.gcr.io/k8s-dns-kube-dns-amd64:1.14.10"
        ],
        "sizeBytes": 49549457
      },
      {
        "names": [
          "k8s.gcr.io/ip-masq-agent-amd64@sha256:723cd85422e644427dd00c0d2b1ece9e618f3c1789543f8d68edceb65ef610f5",
          "k8s.gcr.io/ip-masq-agent-amd64:v2.0.2"
        ],
        "sizeBytes": 48581568
      },
      {
        "names": [
          "k8s.gcr.io/metrics-server-amd64@sha256:49a9f12f7067d11f42c803dbe61ed2c1299959ad85cb315b25ff7eef8e6b8892",
          "k8s.gcr.io/metrics-server-amd64:v0.2.1"
        ],
        "sizeBytes": 42541759
      },
      {
        "names": [
          "k8s.gcr.io/k8s-dns-sidecar-amd64@sha256:4f1ab957f87b94a5ec1edc26fae50da2175461f00afecf68940c4aa079bd08a4",
          "k8s.gcr.io/k8s-dns-sidecar-amd64:1.14.10"
        ],
        "sizeBytes": 41635309
      },
      {
        "names": [
          "k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64@sha256:bbb2a290a568125b3b996028958eb773f33b5b87a6b37bf38a28f8b62dddb3c8",
          "k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.14.10"
        ],
        "sizeBytes": 40372149
      },
      {
        "names": [
          "k8s.gcr.io/addon-resizer@sha256:507aa9845ecce1fdde4d61f530c802f4dc2974c700ce0db7730866e442db958d",
          "k8s.gcr.io/addon-resizer:1.8.1"
        ],
        "sizeBytes": 32968591
      },
      {
        "names": [
          "quay.io/coreos/kube-state-metrics@sha256:6685342bbd645754b1aabdd9b663691109ec680645af261288289e62571ac201",
          "quay.io/coreos/kube-state-metrics:v1.4.0"
        ],
        "sizeBytes": 28151762
      },
      {
        "names": [
          "prom/node-exporter@sha256:55302581333c43d540db0e144cf9e7735423117a733cdec27716d87254221086",
          "prom/node-exporter:v0.16.0"
        ],
        "sizeBytes": 22915749
      },
      {
        "names": [
          "k8s.gcr.io/kube-state-metrics@sha256:15c89813ccd3d426c2023df9c1ab2edff33a279ac9779be488a3e4463c555739",
          "k8s.gcr.io/kube-state-metrics:v1.3.1"
        ],
        "sizeBytes": 22163232
      },
      {
        "names": [
          "k8s.gcr.io/metadata-proxy@sha256:5be758058e67b578f7041498e2daca46ccd7426bc602d35ed0f50bd4a3659d50",
          "k8s.gcr.io/metadata-proxy:v0.1.10"
        ],
        "sizeBytes": 8953717
      },
      {
        "names": [
          "k8s.gcr.io/defaultbackend@sha256:865b0c35e6da393b8e80b7e3799f777572399a4cff047eb02a81fa6e7a48ed4b",
          "k8s.gcr.io/defaultbackend:1.4"
        ],
        "sizeBytes": 4844064
      },
      {
        "names": [
          "k8s.gcr.io/busybox@sha256:545e6a6310a27636260920bc07b994a299b6708a1b26910cfefd335fdfb60d2b",
          "k8s.gcr.io/busybox:1.27"
        ],
        "sizeBytes": 1129289
      },
      {
        "names": [
          "k8s.gcr.io/busybox@sha256:4bdd623e848417d96127e16037743f0cd8b528c026e9175e22a84f639eca58ff",
          "k8s.gcr.io/busybox:1.24"
        ],
        "sizeBytes": 1113554
      },
      {
        "names": [
          "asia.gcr.io/google_containers/pause-amd64@sha256:163ac025575b775d1c0f9bf0bdd0f086883171eb475b5068e7defa4ca9e76516",
          "eu.gcr.io/google_containers/pause-amd64@sha256:163ac025575b775d1c0f9bf0bdd0f086883171eb475b5068e7defa4ca9e76516",
          "gcr.io/google_containers/pause-amd64@sha256:163ac025575b775d1c0f9bf0bdd0f086883171eb475b5068e7defa4ca9e76516",
          "asia.gcr.io/google_containers/pause-amd64:3.0",
          "eu.gcr.io/google_containers/pause-amd64:3.0"
        ],
        "sizeBytes": 746888
      },
      {
        "names": [
          "k8s.gcr.io/pause-amd64@sha256:59eec8837a4d942cc19a52b8c09ea75121acc38114a2c68b98983ce9356b8610",
          "k8s.gcr.io/pause-amd64:3.1"
        ],
        "sizeBytes": 742472
      }
    ]
  }
}
```


# Pod

```json
{
  "metadata": {
    "name": "cost-attribution-prometheus-5dd645756b-w4n7s",
    "generateName": "cost-attribution-prometheus-5dd645756b-",
    "namespace": "test-ns",
    "selfLink": "/api/v1/namespaces/test-ns/pods/cost-attribution-prometheus-5dd645756b-w4n7s",
    "uid": "0e393372-e1e6-11e8-a351-42010a800049",
    "resourceVersion": "6235965",
    "creationTimestamp": "2018-11-06T17:04:43Z",
    "labels": {
      "app": "prometheus",
      "pod-template-hash": "1882013126"
    },
    "ownerReferences": [
      {
        "apiVersion": "extensions/v1beta1",
        "kind": "ReplicaSet",
        "name": "cost-attribution-prometheus-5dd645756b",
        "uid": "9cb308c3-e11b-11e8-a351-42010a800049",
        "controller": true,
        "blockOwnerDeletion": true
      }
    ]
  },
  "spec": {
    "volumes": [
      {
        "name": "config",
        "configMap": {
          "name": "cost-attribution-prometheus-config",
          "defaultMode": 420
        }
      },
      {
        "name": "datadir",
        "persistentVolumeClaim": {
          "claimName": "cost-attribution-data"
        }
      },
      {
        "name": "test-deployment-controller-serviceaccount-a970-token-qhz9x",
        "secret": {
          "secretName": "test-deployment-controller-serviceaccount-a970-token-qhz9x",
          "defaultMode": 420
        }
      }
    ],
    "initContainers": [
      {
        "name": "init-directory",
        "image": "busybox",
        "command": [
          "sh",
          "-c",
          "mkdir -p /data/prometheus-data; chmod -R 777 /data/prometheus-data;"
        ],
        "resources": {},
        "volumeMounts": [
          {
            "name": "datadir",
            "mountPath": "/data"
          },
          {
            "name": "test-deployment-controller-serviceaccount-a970-token-qhz9x",
            "readOnly": true,
            "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
          }
        ],
        "terminationMessagePath": "/dev/termination-log",
        "terminationMessagePolicy": "File",
        "imagePullPolicy": "Always"
      }
    ],
    "containers": [
      {
        "name": "prometheus",
        "image": "prom/prometheus:v2.2.1",
        "args": [
          "--config.file=/etc/prometheus/config/config.yml",
          "--web.enable-lifecycle",
          "--storage.tsdb.path=/data/prometheus-data",
          "--storage.tsdb.retention=60d"
        ],
        "ports": [
          {
            "name": "web",
            "containerPort": 9090,
            "protocol": "TCP"
          }
        ],
        "resources": {
          "limits": {
            "cpu": "4",
            "memory": "4000Mi"
          },
          "requests": {
            "cpu": "250m",
            "memory": "250Mi"
          }
        },
        "volumeMounts": [
          {
            "name": "config",
            "mountPath": "/etc/prometheus/config"
          },
          {
            "name": "datadir",
            "mountPath": "/data"
          },
          {
            "name": "test-deployment-controller-serviceaccount-a970-token-qhz9x",
            "readOnly": true,
            "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
          }
        ],
        "livenessProbe": {
          "httpGet": {
            "path": "/",
            "port": 9090,
            "scheme": "HTTP"
          },
          "initialDelaySeconds": 20,
          "timeoutSeconds": 5,
          "periodSeconds": 10,
          "successThreshold": 1,
          "failureThreshold": 3
        },
        "readinessProbe": {
          "httpGet": {
            "path": "/",
            "port": 9090,
            "scheme": "HTTP"
          },
          "initialDelaySeconds": 20,
          "timeoutSeconds": 5,
          "periodSeconds": 5,
          "successThreshold": 1,
          "failureThreshold": 3
        },
        "terminationMessagePath": "/dev/termination-log",
        "terminationMessagePolicy": "File",
        "imagePullPolicy": "Always"
      }
    ],
    "restartPolicy": "Always",
    "terminationGracePeriodSeconds": 300,
    "dnsPolicy": "ClusterFirst",
    "serviceAccountName": "test-deployment-controller-serviceaccount-a970",
    "serviceAccount": "test-deployment-controller-serviceaccount-a970",
    "nodeName": "gke-gar-1-pool-1-acc63c03-7zvd",
    "securityContext": {},
    "schedulerName": "default-scheduler",
    "tolerations": [
      {
        "key": "node.kubernetes.io/not-ready",
        "operator": "Exists",
        "effect": "NoExecute",
        "tolerationSeconds": 300
      },
      {
        "key": "node.kubernetes.io/unreachable",
        "operator": "Exists",
        "effect": "NoExecute",
        "tolerationSeconds": 300
      }
    ]
  },
  "status": {
    "phase": "Running",
    "conditions": [
      {
        "type": "Initialized",
        "status": "True",
        "lastProbeTime": null,
        "lastTransitionTime": "2018-11-06T19:22:27Z"
      },
      {
        "type": "Ready",
        "status": "True",
        "lastProbeTime": null,
        "lastTransitionTime": "2018-11-06T19:22:54Z"
      },
      {
        "type": "PodScheduled",
        "status": "True",
        "lastProbeTime": null,
        "lastTransitionTime": "2018-11-06T19:22:08Z"
      }
    ],
    "hostIP": "10.128.0.2",
    "podIP": "10.48.11.3",
    "startTime": "2018-11-06T19:22:08Z",
    "initContainerStatuses": [
      {
        "name": "init-directory",
        "state": {
          "terminated": {
            "exitCode": 0,
            "reason": "Completed",
            "startedAt": "2018-11-06T19:22:27Z",
            "finishedAt": "2018-11-06T19:22:27Z",
            "containerID": "docker://2f0e77847b64861395bb47e1dd8f1e6bea6103f96f588f4057c7830239c9826f"
          }
        },
        "lastState": {},
        "ready": true,
        "restartCount": 0,
        "image": "busybox:latest",
        "imageID": "docker-pullable://busybox@sha256:915f390a8912e16d4beb8689720a17348f3f6d1a7b659697df850ab625ea29d5",
        "containerID": "docker://2f0e77847b64861395bb47e1dd8f1e6bea6103f96f588f4057c7830239c9826f"
      }
    ],
    "containerStatuses": [
      {
        "name": "prometheus",
        "state": {
          "running": {
            "startedAt": "2018-11-06T19:22:31Z"
          }
        },
        "lastState": {},
        "ready": true,
        "restartCount": 0,
        "image": "prom/prometheus:v2.2.1",
        "imageID": "docker-pullable://prom/prometheus@sha256:129e16b08818a47259d972767fd834d84fb70ca11b423cc9976c9bce9b40c58f",
        "containerID": "docker://c713962194489f04f00d44f69e575217ab91bc30328555533a4fc059d8091139"
      }
    ],
    "qosClass": "Burstable"
  }
}
```

# PersistentVolumeClaim

Kubectl:
```
kubectl get --raw /api/v1/persistentvolumes | jq
```

```
{
  "kind": "PersistentVolumeList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/persistentvolumes",
    "resourceVersion": "406721"
  },
  "items": [
    {
      "metadata": {
        "name": "pvc-d065fcbe-edcf-11e8-b20f-42010a800020",
        "selfLink": "/api/v1/persistentvolumes/pvc-d065fcbe-edcf-11e8-b20f-42010a800020",
        "uid": "d36f9ff8-edcf-11e8-b20f-42010a800020",
        "resourceVersion": "6809",
        "creationTimestamp": "2018-11-21T20:55:49Z",
        "labels": {
          "failure-domain.beta.kubernetes.io/region": "us-central1",
          "failure-domain.beta.kubernetes.io/zone": "us-central1-b"
        },
        "annotations": {
          "kubernetes.io/createdby": "gce-pd-dynamic-provisioner",
          "pv.kubernetes.io/bound-by-controller": "yes",
          "pv.kubernetes.io/provisioned-by": "kubernetes.io/gce-pd"
        }
      },
      "spec": {
        "capacity": {
          "storage": "50Gi"
        },
        "gcePersistentDisk": {
          "pdName": "gke-gar-2-702b8e2e-dyn-pvc-d065fcbe-edcf-11e8-b20f-42010a800020",
          "fsType": "ext4"
        },
        "accessModes": [
          "ReadWriteOnce"
        ],
        "claimRef": {
          "kind": "PersistentVolumeClaim",
          "namespace": "kubernetes-cost-attribution",
          "name": "cost-attribution-data",
          "uid": "d065fcbe-edcf-11e8-b20f-42010a800020",
          "apiVersion": "v1",
          "resourceVersion": "6750"
        },
        "persistentVolumeReclaimPolicy": "Delete",
        "storageClassName": "standard"
      },
      "status": {
        "phase": "Bound"
      }
    },
    {
      "metadata": {
        "name": "pvc-fd986382-eddb-11e8-910e-42010a800036",
        "selfLink": "/api/v1/persistentvolumes/pvc-fd986382-eddb-11e8-910e-42010a800036",
        "uid": "0130e784-eddc-11e8-b20f-42010a800020",
        "resourceVersion": "19168",
        "creationTimestamp": "2018-11-21T22:23:00Z",
        "labels": {
          "failure-domain.beta.kubernetes.io/region": "us-central1",
          "failure-domain.beta.kubernetes.io/zone": "us-central1-b"
        },
        "annotations": {
          "kubernetes.io/createdby": "gce-pd-dynamic-provisioner",
          "pv.kubernetes.io/bound-by-controller": "yes",
          "pv.kubernetes.io/provisioned-by": "kubernetes.io/gce-pd"
        }
      },
      "spec": {
        "capacity": {
          "storage": "50Gi"
        },
        "gcePersistentDisk": {
          "pdName": "gke-gar-2-702b8e2e-dyn-pvc-fd986382-eddb-11e8-910e-42010a800036",
          "fsType": "ext4"
        },
        "accessModes": [
          "ReadWriteOnce"
        ],
        "claimRef": {
          "kind": "PersistentVolumeClaim",
          "namespace": "test-ns",
          "name": "cost-attribution-data",
          "uid": "fd986382-eddb-11e8-910e-42010a800036",
          "apiVersion": "v1",
          "resourceVersion": "19104"
        },
        "persistentVolumeReclaimPolicy": "Delete",
        "storageClassName": "standard"
      },
      "status": {
        "phase": "Bound"
      }
    }
  ]
}
```
# Service

```
$ kubectl get --raw /api/v1/services | jq
```

```json
{
  "kind": "ServiceList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/services",
    "resourceVersion": "793822"
  },
  "items": [
    {
      "metadata": {
        "name": "kubernetes",
        "namespace": "default",
        "selfLink": "/api/v1/namespaces/default/services/kubernetes",
        "uid": "b145823f-edc9-11e8-b20f-42010a800020",
        "resourceVersion": "6",
        "creationTimestamp": "2018-11-21T20:11:55Z",
        "labels": {
          "component": "apiserver",
          "provider": "kubernetes"
        }
      },
      "spec": {
        "ports": [
          {
            "name": "https",
            "protocol": "TCP",
            "port": 443,
            "targetPort": 443
          }
        ],
        "clusterIP": "10.59.240.1",
        "type": "ClusterIP",
        "sessionAffinity": "ClientIP",
        "sessionAffinityConfig": {
          "clientIP": {
            "timeoutSeconds": 10800
          }
        }
      },
      "status": {
        "loadBalancer": {}
      }
    },
    {
      "metadata": {
        "name": "default-http-backend",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/services/default-http-backend",
        "uid": "c4c6d3a8-edc9-11e8-b20f-42010a800020",
        "resourceVersion": "278",
        "creationTimestamp": "2018-11-21T20:12:28Z",
        "labels": {
          "addonmanager.kubernetes.io/mode": "Reconcile",
          "k8s-app": "glbc",
          "kubernetes.io/cluster-service": "true",
          "kubernetes.io/name": "GLBCDefaultBackend"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"k8s-app\":\"glbc\",\"kubernetes.io/cluster-service\":\"true\",\"kubernetes.io/name\":\"GLBCDefaultBackend\"},\"name\":\"default-http-backend\",\"namespace\":\"kube-system\"},\"spec\":{\"ports\":[{\"name\":\"http\",\"port\":80,\"protocol\":\"TCP\",\"targetPort\":8080}],\"selector\":{\"k8s-app\":\"glbc\"},\"type\":\"NodePort\"}}\n"
        }
      },
      "spec": {
        "ports": [
          {
            "name": "http",
            "protocol": "TCP",
            "port": 80,
            "targetPort": 8080,
            "nodePort": 30889
          }
        ],
        "selector": {
          "k8s-app": "glbc"
        },
        "clusterIP": "10.59.246.64",
        "type": "NodePort",
        "sessionAffinity": "None",
        "externalTrafficPolicy": "Cluster"
      },
      "status": {
        "loadBalancer": {}
      }
    },
    {
      "metadata": {
        "name": "heapster",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/services/heapster",
        "uid": "c58fa9db-edc9-11e8-b20f-42010a800020",
        "resourceVersion": "299",
        "creationTimestamp": "2018-11-21T20:12:29Z",
        "labels": {
          "addonmanager.kubernetes.io/mode": "Reconcile",
          "kubernetes.io/cluster-service": "true",
          "kubernetes.io/name": "Heapster"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"kubernetes.io/cluster-service\":\"true\",\"kubernetes.io/name\":\"Heapster\"},\"name\":\"heapster\",\"namespace\":\"kube-system\"},\"spec\":{\"ports\":[{\"port\":80,\"targetPort\":8082}],\"selector\":{\"k8s-app\":\"heapster\"}}}\n"
        }
      },
      "spec": {
        "ports": [
          {
            "protocol": "TCP",
            "port": 80,
            "targetPort": 8082
          }
        ],
        "selector": {
          "k8s-app": "heapster"
        },
        "clusterIP": "10.59.254.38",
        "type": "ClusterIP",
        "sessionAffinity": "None"
      },
      "status": {
        "loadBalancer": {}
      }
    },
    {
      "metadata": {
        "name": "kube-dns",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/services/kube-dns",
        "uid": "c6108f81-edc9-11e8-b20f-42010a800020",
        "resourceVersion": "315",
        "creationTimestamp": "2018-11-21T20:12:30Z",
        "labels": {
          "addonmanager.kubernetes.io/mode": "Reconcile",
          "k8s-app": "kube-dns",
          "kubernetes.io/cluster-service": "true",
          "kubernetes.io/name": "KubeDNS"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"k8s-app\":\"kube-dns\",\"kubernetes.io/cluster-service\":\"true\",\"kubernetes.io/name\":\"KubeDNS\"},\"name\":\"kube-dns\",\"namespace\":\"kube-system\"},\"spec\":{\"clusterIP\":\"10.59.240.10\",\"ports\":[{\"name\":\"dns\",\"port\":53,\"protocol\":\"UDP\"},{\"name\":\"dns-tcp\",\"port\":53,\"protocol\":\"TCP\"}],\"selector\":{\"k8s-app\":\"kube-dns\"}}}\n"
        }
      },
      "spec": {
        "ports": [
          {
            "name": "dns",
            "protocol": "UDP",
            "port": 53,
            "targetPort": 53
          },
          {
            "name": "dns-tcp",
            "protocol": "TCP",
            "port": 53,
            "targetPort": 53
          }
        ],
        "selector": {
          "k8s-app": "kube-dns"
        },
        "clusterIP": "10.59.240.10",
        "type": "ClusterIP",
        "sessionAffinity": "None"
      },
      "status": {
        "loadBalancer": {}
      }
    },
    {
      "metadata": {
        "name": "kubernetes-dashboard",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/services/kubernetes-dashboard",
        "uid": "c5faea74-edc9-11e8-b20f-42010a800020",
        "resourceVersion": "312",
        "creationTimestamp": "2018-11-21T20:12:30Z",
        "labels": {
          "addonmanager.kubernetes.io/mode": "Reconcile",
          "k8s-app": "kubernetes-dashboard",
          "kubernetes.io/cluster-service": "true"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"k8s-app\":\"kubernetes-dashboard\",\"kubernetes.io/cluster-service\":\"true\"},\"name\":\"kubernetes-dashboard\",\"namespace\":\"kube-system\"},\"spec\":{\"ports\":[{\"port\":443,\"targetPort\":8443}],\"selector\":{\"k8s-app\":\"kubernetes-dashboard\"}}}\n"
        }
      },
      "spec": {
        "ports": [
          {
            "protocol": "TCP",
            "port": 443,
            "targetPort": 8443
          }
        ],
        "selector": {
          "k8s-app": "kubernetes-dashboard"
        },
        "clusterIP": "10.59.253.235",
        "type": "ClusterIP",
        "sessionAffinity": "None"
      },
      "status": {
        "loadBalancer": {}
      }
    },
    {
      "metadata": {
        "name": "metrics-server",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/services/metrics-server",
        "uid": "cbba6ed6-edc9-11e8-b20f-42010a800020",
        "resourceVersion": "382",
        "creationTimestamp": "2018-11-21T20:12:39Z",
        "labels": {
          "addonmanager.kubernetes.io/mode": "Reconcile",
          "kubernetes.io/cluster-service": "true",
          "kubernetes.io/name": "Metrics-server"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"addonmanager.kubernetes.io/mode\":\"Reconcile\",\"kubernetes.io/cluster-service\":\"true\",\"kubernetes.io/name\":\"Metrics-server\"},\"name\":\"metrics-server\",\"namespace\":\"kube-system\"},\"spec\":{\"ports\":[{\"port\":443,\"protocol\":\"TCP\",\"targetPort\":\"https\"}],\"selector\":{\"k8s-app\":\"metrics-server\"}}}\n"
        }
      },
      "spec": {
        "ports": [
          {
            "protocol": "TCP",
            "port": 443,
            "targetPort": "https"
          }
        ],
        "selector": {
          "k8s-app": "metrics-server"
        },
        "clusterIP": "10.59.247.162",
        "type": "ClusterIP",
        "sessionAffinity": "None"
      },
      "status": {
        "loadBalancer": {}
      }
    },
    {
      "metadata": {
        "name": "cost-attribution-grafana",
        "namespace": "kubernetes-cost-attribution",
        "selfLink": "/api/v1/namespaces/kubernetes-cost-attribution/services/cost-attribution-grafana",
        "uid": "d11aa4fb-edcf-11e8-b20f-42010a800020",
        "resourceVersion": "6967",
        "creationTimestamp": "2018-11-21T20:55:45Z",
        "labels": {
          "app": "cost-attribution-grafana"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"cost-attribution-grafana\"},\"name\":\"cost-attribution-grafana\",\"namespace\":\"kubernetes-cost-attribution\"},\"spec\":{\"ports\":[{\"name\":\"web\",\"port\":80,\"targetPort\":3000}],\"selector\":{\"app\":\"cost-attribution-grafana\"},\"type\":\"LoadBalancer\"}}\n"
        }
      },
      "spec": {
        "ports": [
          {
            "name": "web",
            "protocol": "TCP",
            "port": 80,
            "targetPort": 3000,
            "nodePort": 30392
          }
        ],
        "selector": {
          "app": "cost-attribution-grafana"
        },
        "clusterIP": "10.59.254.168",
        "type": "LoadBalancer",
        "sessionAffinity": "None",
        "externalTrafficPolicy": "Cluster"
      },
      "status": {
        "loadBalancer": {
          "ingress": [
            {
              "ip": "35.193.118.19"
            }
          ]
        }
      }
    },
    {
      "metadata": {
        "name": "cost-attribution-mk-agent",
        "namespace": "kubernetes-cost-attribution",
        "selfLink": "/api/v1/namespaces/kubernetes-cost-attribution/services/cost-attribution-mk-agent",
        "uid": "d0ea86b0-edcf-11e8-b20f-42010a800020",
        "resourceVersion": "6771",
        "creationTimestamp": "2018-11-21T20:55:45Z",
        "labels": {
          "app": "cost-attribution-mk-agent"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{\"prometheus.io/port\":\"9101\",\"prometheus.io/scrape\":\"true\"},\"labels\":{\"app\":\"cost-attribution-mk-agent\"},\"name\":\"cost-attribution-mk-agent\",\"namespace\":\"kubernetes-cost-attribution\"},\"spec\":{\"ports\":[{\"name\":\"web\",\"port\":9101}],\"selector\":{\"app\":\"cost-attribution-mk-agent\"}}}\n",
          "prometheus.io/port": "9101",
          "prometheus.io/scrape": "true"
        }
      },
      "spec": {
        "ports": [
          {
            "name": "web",
            "protocol": "TCP",
            "port": 9101,
            "targetPort": 9101
          }
        ],
        "selector": {
          "app": "cost-attribution-mk-agent"
        },
        "clusterIP": "10.59.247.179",
        "type": "ClusterIP",
        "sessionAffinity": "None"
      },
      "status": {
        "loadBalancer": {}
      }
    },
    {
      "metadata": {
        "name": "cost-attribution-prometheus",
        "namespace": "kubernetes-cost-attribution",
        "selfLink": "/api/v1/namespaces/kubernetes-cost-attribution/services/cost-attribution-prometheus",
        "uid": "d0a56407-edcf-11e8-b20f-42010a800020",
        "resourceVersion": "6757",
        "creationTimestamp": "2018-11-21T20:55:44Z",
        "labels": {
          "app": "prometheus"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{\"prometheus.io/scrape\":\"true\"},\"labels\":{\"app\":\"prometheus\"},\"name\":\"cost-attribution-prometheus\",\"namespace\":\"kubernetes-cost-attribution\"},\"spec\":{\"ports\":[{\"name\":\"prometheus\",\"port\":9090}],\"selector\":{\"app\":\"prometheus\"}}}\n",
          "prometheus.io/scrape": "true"
        }
      },
      "spec": {
        "ports": [
          {
            "name": "prometheus",
            "protocol": "TCP",
            "port": 9090,
            "targetPort": 9090
          }
        ],
        "selector": {
          "app": "prometheus"
        },
        "clusterIP": "10.59.255.166",
        "type": "ClusterIP",
        "sessionAffinity": "None"
      },
      "status": {
        "loadBalancer": {}
      }
    },
    {
      "metadata": {
        "name": "cost-attribution-grafana",
        "namespace": "test-ns",
        "selfLink": "/api/v1/namespaces/test-ns/services/cost-attribution-grafana",
        "uid": "fdac3833-eddb-11e8-910e-42010a800036",
        "resourceVersion": "19276",
        "creationTimestamp": "2018-11-21T22:22:54Z",
        "labels": {
          "app": "cost-attribution-grafana",
          "app.kubernetes.io/name": "test-deployment"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"cost-attribution-grafana\",\"app.kubernetes.io/name\":\"test-deployment\"},\"name\":\"cost-attribution-grafana\",\"namespace\":\"test-ns\",\"ownerReferences\":[{\"apiVersion\":\"app.k8s.io/v1beta1\",\"blockOwnerDeletion\":true,\"kind\":\"Application\",\"name\":\"test-deployment\",\"uid\":\"f99fc1e9-eddb-11e8-96bc-42010a800022\"}]},\"spec\":{\"ports\":[{\"name\":\"web\",\"port\":80,\"targetPort\":3000}],\"selector\":{\"app\":\"cost-attribution-grafana\"},\"type\":\"LoadBalancer\"}}\n"
        },
        "ownerReferences": [
          {
            "apiVersion": "app.k8s.io/v1beta1",
            "kind": "Application",
            "name": "test-deployment",
            "uid": "f99fc1e9-eddb-11e8-96bc-42010a800022",
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "ports": [
          {
            "name": "web",
            "protocol": "TCP",
            "port": 80,
            "targetPort": 3000,
            "nodePort": 31633
          }
        ],
        "selector": {
          "app": "cost-attribution-grafana"
        },
        "clusterIP": "10.59.243.238",
        "type": "LoadBalancer",
        "sessionAffinity": "None",
        "externalTrafficPolicy": "Cluster"
      },
      "status": {
        "loadBalancer": {
          "ingress": [
            {
              "ip": "35.192.152.172"
            }
          ]
        }
      }
    },
    {
      "metadata": {
        "name": "cost-attribution-mk-agent",
        "namespace": "test-ns",
        "selfLink": "/api/v1/namespaces/test-ns/services/cost-attribution-mk-agent",
        "uid": "fda336a4-eddb-11e8-910e-42010a800036",
        "resourceVersion": "19110",
        "creationTimestamp": "2018-11-21T22:22:54Z",
        "labels": {
          "app": "cost-attribution-mk-agent",
          "app.kubernetes.io/name": "test-deployment"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{\"prometheus.io/port\":\"9101\",\"prometheus.io/scrape\":\"true\"},\"labels\":{\"app\":\"cost-attribution-mk-agent\",\"app.kubernetes.io/name\":\"test-deployment\"},\"name\":\"cost-attribution-mk-agent\",\"namespace\":\"test-ns\",\"ownerReferences\":[{\"apiVersion\":\"app.k8s.io/v1beta1\",\"blockOwnerDeletion\":true,\"kind\":\"Application\",\"name\":\"test-deployment\",\"uid\":\"f99fc1e9-eddb-11e8-96bc-42010a800022\"}]},\"spec\":{\"ports\":[{\"name\":\"web\",\"port\":9101}],\"selector\":{\"app\":\"cost-attribution-mk-agent\"}}}\n",
          "prometheus.io/port": "9101",
          "prometheus.io/scrape": "true"
        },
        "ownerReferences": [
          {
            "apiVersion": "app.k8s.io/v1beta1",
            "kind": "Application",
            "name": "test-deployment",
            "uid": "f99fc1e9-eddb-11e8-96bc-42010a800022",
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "ports": [
          {
            "name": "web",
            "protocol": "TCP",
            "port": 9101,
            "targetPort": 9101
          }
        ],
        "selector": {
          "app": "cost-attribution-mk-agent"
        },
        "clusterIP": "10.59.245.162",
        "type": "ClusterIP",
        "sessionAffinity": "None"
      },
      "status": {
        "loadBalancer": {}
      }
    },
    {
      "metadata": {
        "name": "cost-attribution-prometheus",
        "namespace": "test-ns",
        "selfLink": "/api/v1/namespaces/test-ns/services/cost-attribution-prometheus",
        "uid": "fd9ba05f-eddb-11e8-910e-42010a800036",
        "resourceVersion": "19106",
        "creationTimestamp": "2018-11-21T22:22:54Z",
        "labels": {
          "app": "prometheus",
          "app.kubernetes.io/name": "test-deployment"
        },
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{\"prometheus.io/scrape\":\"true\"},\"labels\":{\"app\":\"prometheus\",\"app.kubernetes.io/name\":\"test-deployment\"},\"name\":\"cost-attribution-prometheus\",\"namespace\":\"test-ns\",\"ownerReferences\":[{\"apiVersion\":\"app.k8s.io/v1beta1\",\"blockOwnerDeletion\":true,\"kind\":\"Application\",\"name\":\"test-deployment\",\"uid\":\"f99fc1e9-eddb-11e8-96bc-42010a800022\"}]},\"spec\":{\"ports\":[{\"name\":\"prometheus\",\"port\":9090}],\"selector\":{\"app\":\"prometheus\"}}}\n",
          "prometheus.io/scrape": "true"
        },
        "ownerReferences": [
          {
            "apiVersion": "app.k8s.io/v1beta1",
            "kind": "Application",
            "name": "test-deployment",
            "uid": "f99fc1e9-eddb-11e8-96bc-42010a800022",
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "ports": [
          {
            "name": "prometheus",
            "protocol": "TCP",
            "port": 9090,
            "targetPort": 9090
          }
        ],
        "selector": {
          "app": "prometheus"
        },
        "clusterIP": "10.59.249.39",
        "type": "ClusterIP",
        "sessionAffinity": "None"
      },
      "status": {
        "loadBalancer": {}
      }
    }
  ]
}
```
