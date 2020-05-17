# Ejemplos en Golang para Desarrollo de Aplicaciones Cliente-Servidor

## Contruir los ejemplo

```console
make docker
```

## Desplegar en Kubernetes

```console
kubectl apply -f deploy/deployment.yaml
kubectl apply -f deploy/service.yaml
```

## Acceder al servicio

```console
curl -X GET ku  http://localhost:8080/hospital
```

## Limpiar

```console
kubectl delete -f deploy/deployment.yaml
kubectl delete -f deploy/service.yaml
```
