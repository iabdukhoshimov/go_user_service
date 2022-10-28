# Udevs microservice helm chart

## Updating chart

```shell
helm package --version x.x.x .
curl --request POST --form 'chart=@microservice-X.X.X.tgz' --user @YOUR_USERNAME:@YOUR_ACCESS_TOKEN https://gitlab.udevs.io/api/v4/projects/476/packages/helm/api/stable/charts
```

## Installation to the cluster

```shell
helm repo add --username @YOUR_USERNAME --password @YOUR_ACCESS_TOKEN microservice https://gitlab.udevs.io/api/v4/projects/476/packages/helm/stable
helm repo update
```

## Updating via `helm push`

### After adding charts to your cluster you can update charts using helm push extension

```shell
helm push microservice-X.X.X.tgz microservice
```

## Releasing service with another values.yaml

```shell
helm install @YOUR_SERVICE_NAME --values .helm/values.yaml microservice/microservice

# ex
helm install user-service --values .helm/values.yaml microservice/microservice
```