# Ejemplos en Golang para Desarrollo de Aplicaciones Cliente-Servidor

## Construir la imagen de Docker

```console
make docker
```

## Desplegar en Kubernetes

```console
kubectl apply -f deploy/backend.yaml
kubectl apply -f deploy/frontend.yaml
kubectl apply -f deploy/service.yaml
```

## Acceder al servicio

```console
curl -X GET ku  http://localhost:8080/hospital
```

## Limpiar

```console
kubectl delete -f deploy/backend.yaml
kubectl delete -f deploy/frontend.yaml
kubectl delete -f deploy/service.yaml
```
