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

var addonControllerTemplate = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: addon-controller
  name: addon-controller-{{.SHARD}}
  namespace: projectsveltos
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: addon-controller
  template:
    metadata:
      annotations:
        kubectl.kubernetes.io/default-container: controller
      labels:
        control-plane: addon-controller
    spec:
      containers:
      - args:
        - --diagnostics-address=:8443
        - --report-mode=0
        - --shard-key={{.SHARD}}
        - --v=5
        - --version=v0.44.0
        command:
        - /manager
        env:
        - name: GOMEMLIMIT
          valueFrom:
            resourceFieldRef:
              resource: limits.memory
        - name: GOMAXPROCS
          valueFrom:
            resourceFieldRef:
              resource: limits.cpu
        image: docker.io/projectsveltos/addon-controller:v0.44.0
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /healthz
            port: healthz
            scheme: HTTP
          initialDelaySeconds: 15
          periodSeconds: 20
        name: controller
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
          requests:
            memory: 256Mi
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
            - ALL
          seccompProfile:
            type: RuntimeDefault
        volumeMounts:
        - mountPath: /tmp
          name: tmp
      securityContext:
        runAsNonRoot: true
      serviceAccountName: addon-controller
      terminationGracePeriodSeconds: 10
      volumes:
      - emptyDir: {}
        name: tmp
`)
