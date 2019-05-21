# mig-controller

## Prerequisites

 - golang compiler (tested @ 1.11.5)
 - kubebuilder (tested @ 1.0.7)
 - dep (tested @ v0.5.0)
 - velero (tested @ v0.11.0) installed on both clusters involved in migration

---

## Quick-start

__1. Identify a pair of running OpenShift clusters to migrate workloads between__

You'll be able to use mig-controller and mig-ui to move workloads from a _source_ to a _destination_ cluster. You'll need cluster-admin permissions on both clusters.

---

__2. Deploy Velero to both the _source_ and _destination_ OpenShift clusters__

```

# Download bash script to deploy Velero along with required plugins
$ wget https://github.com/fusor/mig-controller/blob/master/hack/deploy/deploy_velero.sh

# Login to source cluster, run deploy_velero.sh against it
$ oc login https://my-source-cluster:8443
$ bash deploy_velero.sh

Deploying Velero...
[...]

# Login to destination cluster, run deploy_velero.sh against it
$ oc login https://my-destination-cluster:8443
$ bash deploy_velero.sh

Deploying Velero...
[...]
```

---

__2. Deploy _mig-controller_ and _mig-ui_ to one of the two involved clusters__

```
# Download bash script to deploy latest mig-controller and mig-ui images
wget https://github.com/fusor/mig-controller/blob/master/hack/deploy/deploy_mig.sh

# Login to cluster of your choice where controller + UI will run
oc login https://my-cluster:8443
bash deploy_mig.sh

Deploying mig-controller...
[...]

Deploying mig-ui...
[...]

```

---

__3. Create Mig CRs to describe the Migration that will be performed__

Before mig-controller can run a Migration, you'll need to provide it with:
 - Coordinates & auth info for 2 OpenShift clusters (source + destination)
 - A list of namespaces to be migrated
 - Storage to use for the migration

 These items can be specified by "Mig" CRs. For the source of truth on what will be accepted in CR fields, see the appropriate _types.go_ file.

- [MigPlan](https://github.com/fusor/mig-controller/blob/master/pkg/apis/migration/v1alpha1/migplan_types.go)
- [MigCluster](https://github.com/fusor/mig-controller/blob/master/pkg/apis/migration/v1alpha1/migcluster_types.go)
- [MigStorage](https://github.com/fusor/mig-controller/blob/master/pkg/apis/migration/v1alpha1/migstorage_types.go)
- [MigMigration](https://github.com/fusor/mig-controller/blob/master/pkg/apis/migration/v1alpha1/migmigration_types.go)
- [Cluster](https://github.com/kubernetes/cluster-registry/blob/master/pkg/apis/clusterregistry/v1alpha1/types.go)


*__To make it easier to run your first Migration with mig-controller__*, we've published a set of annotated sample CRs that you can walk through and fill out values on. The first step will be to run `make samples`.

```
make samples
# [... sample CR content will be copied to 'migsamples' dir]
```

**_Inspect and edit each of the files in the 'migsamples' directory, making changes as needed._** Much of the content in these sample files can stay unchanged, but you'll need to provide information such as:
 
 - Remote cluster URL (in the `Cluster` resource)
 - S3 bucket coordinates and access/secret key
 - List of namespaces to be migrated from the source to the destination cluster

After modifying resource yaml, create the resources on the OpenShift cluster where the controller is running.

```
# Option 1: Create everything in a single command
oc apply -f migsamples

# ------------------------------------------------

# Option 2: Create resources individually
cd migsamples

# Source cluster definition 
# Note: no coordinates/auth needed when MigCluster has 'isHostCluster: true'
oc apply -f mig-cluster-local.yaml

# Destination cluster definition, coordinates, auth details
oc apply -f cluster-aws.yaml
oc apply -f sa-secret-aws.yaml
oc apply -f mig-cluster-aws.yaml

# Describes where to store data during Migration, storage auth details
# Note: the contents of mig-storage-creds.yaml will be used to overwrite Velero cloud-credentials
oc apply -f mig-storage-creds.yaml
oc apply -f mig-storage.yaml

# Describes which clusters, storage, and namespaces should be to run a Migration
oc apply -f mig-plan.yaml

# Declares that a Migration operation should be run 
oc apply -f mig-migration.yaml
```

- See [config/samples](https://github.com/fusor/mig-controller/tree/master/config/samples) CR samples. It is _highly_ recommended to run `make samples` to copy these to the .gitignore'd 'migsamples' before filling out cluster details (URLs + SA tokens).

### Creating a Service Account (SA) token

For mig-controller to perform migration actions on a remote cluster, you'll need to provide:
- The remote OpenShift cluster URL
- A valid Service Account (SA) token granting 'cluster-admin' access to the remote cluster


To configure the SA token, run the following on the remote cluster:
```bash
# Create a new service account in the mig ns
oc create namespace mig
oc create sa -n mig mig
# Grant the 'mig' service account cluster-admin (cluster level root privileges, use with caution!)
oc adm policy add-cluster-role-to-user cluster-admin system:serviceaccount:mig:mig
# Get the ServiceAccount token in a base64-encoded format to put in the remote MigCluster spec
oc sa get-token -n mig mig|base64 -w 0

```
Use the base64-encoded SA token from the last command output to fill in `migsamples/sa-secret-aws.yaml`