---
title: Support Matrix
weight: 60
description: "Table of supported Talos Linux versions and respective platforms."
---

| Talos Version                                                                                                  | 1.1                                | 1.0                                |
|----------------------------------------------------------------------------------------------------------------|------------------------------------|------------------------------------|
| Release Date                                                                                                   | 2022-06-22                         | 2022-03-29 (1.0.0)                 |
| End of Community Support                                                                                       | 1.2.0 release (2022-09-01, TBD)    | 1.1.0 release (2022-06-22)         |
| Enterprise Support                                                                                             | [offered by Sidero Labs Inc.](https://www.siderolabs.com/support/) | [offered by Sidero Labs Inc.](https://www.siderolabs.com/support/) |
| Kubernetes                                                                                                     | 1.24, 1.23, 1.22                   | 1.23, 1.22, 1.21                   |
| Architecture                                                                                                   | amd64, arm64                       | amd64, arm64                       |
| **Platforms**                                                                                                  |                                    |                                    |
| - cloud                                                                                                        | AWS, GCP, Azure, Digital Ocean, Hetzner, OpenStack, Oracle Cloud, Scaleway, Vultr, Upcloud | AWS, GCP, Azure, Digital Ocean, Hetzner, OpenStack, Oracle Cloud, Scaleway, Vultr, Upcloud |
| - bare metal                                                                                                   | x86: BIOS, UEFI; arm64: UEFI; boot: ISO, PXE, disk image | x86: BIOS, UEFI; arm64: UEFI; boot: ISO, PXE, disk image |
| - virtualized                                                                                                  | VMware, Hyper-V, KVM, Proxmox, Xen | VMware, Hyper-V, KVM, Proxmox, Xen |
| - SBCs                                                                                                         | Banana Pi M64, Jetson Nano, Libre Computer Board ALL-H3-CC, Pine64, Pine64 Rock64, Radxa ROCK Pi 4c, Raspberry Pi 4B | Banana Pi M64, Jetson Nano, Libre Computer Board ALL-H3-CC, Pine64, Pine64 Rock64, Radxa ROCK Pi 4c, Raspberry Pi 4B |
| - local                                                                                                        | Docker, QEMU                       | Docker, QEMU                       |
| **Cluster API**                                                                                                |                                    |                                    |
| [CAPI Bootstrap Provider Talos](https://github.com/siderolabs/cluster-api-bootstrap-provider-talos)            | >= 0.5.4                           | >= 0.5.3                           |
| [CAPI Control Plane Provider Talos](https://github.com/siderolabs/cluster-api-control-plane-provider-talos)    | >= 0.4.6                           | >= 0.4.5                           |
| [Sidero](https://www.sidero.dev/)                                                                              | >= 0.5.1                           | >= 0.5.0                           |
| **UI**                                                                                                         |                                    |                                    |
| [Theila](https://github.com/siderolabs/theila)                                                                 | ✓                                  | ✓                                  |

## Platform Tiers

* Tier 1: Automated tests, high-priority fixes.
* Tier 2: Tested from time to time, medium-priority bugfixes.
* Tier 3: Not tested by core Talos team, community tested.

### Tier 1

* Metal
* AWS
* GCP

### Tier 2

* Azure
* Digital Ocean
* OpenStack
* VMWare

### Tier 3

* Hetzner
* nocloud
* Oracle Cloud
* Scaleway
* Vultr
* Upcloud
