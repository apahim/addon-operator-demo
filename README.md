# AddOn Operator

![](docs/images/addon-operator.png)


## Installing


```yaml
apiVersion: addon.example.com/v1alpha1
kind: AddOn
metadata:
  name: prow-operator
  namespace: addon-operator
spec:
  targetNamespace: prow-operator
  catalogSourceImage: quay.io/osd-addons/prow-operator-index:71ac363
  operatorGroupInstallMode: SingleNamespace
  subscriptionChannel: alpha
  parameters:
    key1: value1
    key2: value2
    key3: value3
  monitoringStack: true
  upgradeSchedule: 0 10 * * 1-5
  delete: false
```

Logs:

```
INFO controllers.AddOn Monitoring Stack already deployed {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the Namespace {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the CatalogSource {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the OperatorGroup {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the Subscription {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the Params ConfigMap {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the PD and DMS Secret {"addon": "addon-operator/prow-operator"}
```

## Updating

```diff
apiVersion: addon.example.com/v1alpha1
kind: AddOn
metadata:
  name: prow-operator
  namespace: addon-operator
spec:
  targetNamespace: prow-operator
- catalogSourceImage: quay.io/osd-addons/prow-operator-index:71ac363
+ catalogSourceImage: quay.io/osd-addons/prow-operator-index:ae6c761
  operatorGroupInstallMode: SingleNamespace
  subscriptionChannel: alpha
  parameters:
    key1: value1
    key2: value2
    key3: value3
  monitoringStack: true
  upgradeSchedule: 0 10 * * 1-5
  delete: false
```

Logs:

```
INFO controllers.AddOn Updating CatalogSource {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Upgrade pending. Waiting for the next upgrade time slot {"addon": "addon-operator/prow-operator"}
...
INFO controllers.AddOn Approving pending upgrade {"addon": "addon-operator/prow-operator"}
```

## Deleting

```diff
apiVersion: addon.example.com/v1alpha1
kind: AddOn
metadata:
  name: prow-operator
  namespace: addon-operator
spec:
  targetNamespace: prow-operator
  catalogSourceImage: quay.io/osd-addons/prow-operator-index:ae6c761
  operatorGroupInstallMode: SingleNamespace
  subscriptionChannel: alpha
  parameters:
    key1: value1
    key2: value2
    key3: value3
  monitoringStack: true
  upgradeSchedule: 0 10 * * 1-5
- delete: false
+ delete: true
```

Logs:

```
INFO controllers.AddOn Deleting the Subscription {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Deleting the CatalogSource {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Deleting the OperatorGroup {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Deleting the PD and DMS Secret {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Deleting the Params ConfigMap {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Deleting the Namespace {"addon": "addon-operator/prow-operator"}
```
