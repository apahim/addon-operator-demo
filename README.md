# AddOn Operator

![](docs/images/addon-operator.png)

Example CR:

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
```

Example controller logs:

```
INFO controllers.AddOn Monitoring Stack already deployed {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the Namespace {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the CatalogSource {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the OperatorGroup {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the Subscription {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the Params ConfigMap {"addon": "addon-operator/prow-operator"}
INFO controllers.AddOn Creating the PD and DMS Secret {"addon": "addon-operator/prow-operator"}
```
