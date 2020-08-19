# [KUSTOMIZE](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/)

## For customizing kubernetes configurations

1. Generate resources
2. Setting cross-cutting fields
3. Composing and customizing

### Generate resources

Kustomize had `secretGenerator` and `configMapGenerator` which generates secrets and configmap because the source of truth is generally external to the cluster.

```zsh
kubectl kustomize ./example/generate-resources
```

### Setting cross-cutting fields

Use cases for this are:
- Set same namespace for all the resources
- Add name prefix/suffix, labels and annotations

```zsh
kubectl kustomize ./example/cross-cutting
```

### Composing and customizing

Composing multiple k8s resources into one.

For example below we combine deployment and service together in one `kustomization.yaml` file
```zsh
# Create a kustomization.yaml composing them
cat <<EOF >./kustomization.yaml
resources:
- deployment.yaml
- service.yaml
EOF

kubectl kustomize ./  # It is well understood, hence did not add the example.
```

Customizing k8s resources.
Small patches are only possible. For example setting new replica set or cpu limit.

```zsh
# Create a patch increase_replicas.yaml
cat <<EOF > increase_replicas.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  replicas: 3
EOF

cat <<EOF >./kustomization.yaml
resources:
- deployment.yaml
patchesStrategicMerge:
- increase_replicas.yaml
- set_memory.yaml
EOF
```

## Bases and Overlays

Example:
Let's have a directory named `base` with kustomization.yaml

The below two `dev` and `prod` can create and overlay from the `base` and customize further in the three ways mentioned above.

**Note**: `base` doesn't have any information of it's overlays.

```zsh
mkdir dev
cat <<EOF > dev/kustomization.yaml
bases:
- ../base
namePrefix: dev-
EOF

mkdir prod
cat <<EOF > prod/kustomization.yaml
bases:
- ../base
namePrefix: prod-
EOF
```

## Apply/View/Delete using Kustomization

It works with kustomization directory only.

Above commands are self explanatory.
```zsh
kubectl apply -k <kustomization directory>/

kubectl get -k ./

kubectl describe -k ./

kubectl diff -k ./

kubectl delete -k ./
```


