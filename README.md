# ckad-study-apps
A set of simple applications to aid in studying for the Certified Kubernetes Application Develop exam.  The associated YAML manifests are quite basic and can be modified extensively to suite your studying needs.  For example:
* Configure to use Persistent Volumes
* Secrets
* Resource Requests and Limits
* Network Policies
* etc...

## Ambassador App
The ambassador app is a simple go webserver which gets and sets key value pairs through URL commands.  The Kubernetes deployment uses a ambassador pod running NGINX to proxy the connection to a seperate Redis deployment for state persistence.  The app can be deployed using the associated yaml definitions which use existing containers on the public dockerhub repository.  Alternatively, the source is available to build from scratch and push the resulting containers to where ever you choose.  The associated yaml configuration file specifies a service of type `loadbalancer`, if your Kubernetes deployment does not support a loadbalancer capability then adjust the yaml to change the service type to `ingress`.

If you choose to build from scratch then make sure the output binary is named main and it exists within the same directory as the Dockerfile.  Use the following commands to build:

```
cd ambassador_src
CGO_ENABLED=0 GOOS=linux go build -a -o main .
docker build -t <your repo name>/ambassadorweb:<tag> .
docker push <your repo name>/ambassadorweb:<tag>
# update kubernetes yaml manifest to point container image to new repository
kubectl apply -f k8sconfigs/ambassador-svc-dep.yaml
```

## Side Car App
The side car app is a simple go webserver which displays contents of a shared kubernetes volume.  A side car container runs a bash script to populate a timestamp into a file within the share volume for the webserver to read and display for HTTP get requests.  The associated yaml configuration file specifies a service of type `loadbalancer`, if your Kubernetes deployment does not support a loadbalancer capability then adjust the yaml to change the service type to `ingress`.

If you choose to build from scratch then make sure the output binary is named main and it exists within the same directory as the Dockerfile.  Use the following commands to build:

```
cd sidecar_src
CGO_ENABLED=0 GOOS=linux go build -a -o main .
docker build -t <your repo name>/sidecarweb:<tag> .
docker push <your repo name>/sidecarweb:<tag>
cd sidecar
docker build -t <your repo name>/sidecar-sc:<tag> .
docker push <your repo name>/sidecar-sc:<tag>
# update kubernetes yaml manifest to point container image to new repository
kubectl apply -f k8sconfigs/sidecar-svc-dep.yaml
```
