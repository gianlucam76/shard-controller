// Generated by *go generate* - DO NOT EDIT
/*
Copyright 2023. projectsveltos.io. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sharding

var sveltosClusterManagerTemplate = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: sc-manager
  name: sc-manager-{{.SHARD}}
  namespace: projectsveltos
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: sc-manager
  template:
    metadata:
      annotations:
        kubectl.kubernetes.io/default-container: manager
      labels:
        control-plane: sc-manager
    spec:
      containers:
      - args:
        - --diagnostics-address=:8443
        - --shard-key={{.SHARD}}
        - --v=5
        command:
        - /manager
        image: docker.io/projectsveltos/sveltoscluster-manager@sha256:b50d50aa5bbe61810ea04adff875ff30f00dd4adcce285176a82e38492bcee8d
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /healthz
            port: healthz
            scheme: HTTP
          initialDelaySeconds: 15
          periodSeconds: 20
        name: manager
        ports:
        - containerPort: 8443
          name: metrics
          protocol: TCP
        - containerPort: 9440
          name: healthz
          protocol: TCP
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /readyz
            port: healthz
            scheme: HTTP
          initialDelaySeconds: 5
          periodSeconds: 10
        resources:
          limits:
            cpu: 500m
            memory: 512Mi
          requests:
            cpu: 10m
            memory: 128Mi
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
            - ALL
      securityContext:
        runAsNonRoot: true
      serviceAccountName: sc-manager
      terminationGracePeriodSeconds: 10
`)
