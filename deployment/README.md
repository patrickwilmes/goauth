# K8S Deployment

This folder contains example k8s configuration. Ensure to read carefully
through the configuration, and adapt it to your needs. The following
parts are contaiend in the deployment descriptor:
- Actual Deployment of the service
- k8s service
- Ingress
- HPA

To deploy the system ensure to reference the correct docker image from your registry, then:
```bash
kubectl apply -f deployment.yaml
```
