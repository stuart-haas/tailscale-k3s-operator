# Tailscale K3s Operator

A Kubernetes operator that automatically discovers and configures Tailscale nodes as K3s agents, enabling seamless cluster expansion over your Tailscale network. This operator watches for new Tailscale machines and automatically provisions them as K3s agents, making it easy to grow your Kubernetes cluster across your Tailscale network.

## Features

- 🔄 Automatic discovery of new Tailscale nodes
- 🤖 Automated K3s agent provisioning
- 🔒 Secure node communication over Tailscale network
- 🏷️ Tag-based node selection
- 📊 Status tracking of configured nodes
- 🔑 OAuth-based Tailscale API integration

## Use Cases

- Automatically expand your K3s cluster as new Tailscale nodes join
- Manage distributed Kubernetes clusters across multiple locations
- Seamlessly integrate edge devices into your K3s cluster
- Create self-healing cluster infrastructure

## Prerequisites

- Kubernetes cluster (v1.19+)
- Tailscale account and API access
- K3s server running and accessible
- kubectl installed and configured

## Development

- [kubebuilder](https://book.kubebuilder.io/quick-start)
