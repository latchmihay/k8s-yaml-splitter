# k8s-yaml-splitter
It takes a combined kubernetes yaml config and splits it into multiple files in a folder of your choosing

# Usage

```console
Usage: ./k8s-yaml-splitter /path/to/combined-k8s.yaml /path/to/output/dir
Usage Dry Run: ./k8s-yaml-splitter /path/to/combined-k8s.yaml /path/to/output/dir -d
```

# Example:
#### Dry Run:
```console
# ./k8s-yaml-splitter example/combined-k8s.yaml /tmp/k8s-split/ -d
Found! type: Secret | apiVersion: v1 | name: grafana | namespace: istio-system
==> DryRun: Writing /tmp/k8s-split/Secret-grafana.yaml
Found! type: Secret | apiVersion: v1 | name: kiali | namespace: istio-system
==> DryRun: Writing /tmp/k8s-split/Secret-kiali.yaml
Found! type: ConfigMap | apiVersion: v1 | name: istio-galley-configuration | namespace: istio-system
==> DryRun: Writing /tmp/k8s-split/ConfigMap-istio-galley-configuration.yaml
```

#### Normal Run:
```console
# ./k8s-yaml-splitter example/combined-k8s.yaml /tmp/k8s-split/
Found! type: Secret | apiVersion: v1 | name: grafana | namespace: istio-system
* Writing /tmp/k8s-split/Secret-grafana.yaml
* Wrote 250 bytes to /tmp/k8s-split/Secret-grafana.yaml
Found! type: Secret | apiVersion: v1 | name: kiali | namespace: istio-system
* Writing /tmp/k8s-split/Secret-kiali.yaml
* Wrote 242 bytes to /tmp/k8s-split/Secret-kiali.yaml
Found! type: ConfigMap | apiVersion: v1 | name: istio-galley-configuration | namespace: istio-system
* Writing /tmp/k8s-split/ConfigMap-istio-galley-configuration.yaml
* Wrote 3308 bytes to /tmp/k8s-split/ConfigMap-istio-galley-configuration.yaml
```

#### Piped in 
> will create a `.k8s-global-obejcts` folder in your current directory and split the yamls in there

```console
# cat example/combined-k8s.yaml | ./k8s-yaml-splitter
Found! type: Secret | apiVersion: v1 | name: grafana | namespace: istio-system
* Writing /Users/lnm0811/go/src/github.com/latchmihay/k8s-yaml-splitter/.k8s-yaml-splitter/Secret-grafana.yaml
* Wrote 250 bytes to /Users/madcricket/go/src/github.com/latchmihay/k8s-yaml-splitter/.k8s-yaml-splitter/Secret-grafana.yaml
Found! type: Secret | apiVersion: v1 | name: kiali | namespace: istio-system
* Writing /Users/lnm0811/go/src/github.com/latchmihay/k8s-yaml-splitter/.k8s-yaml-splitter/Secret-kiali.yaml
* Wrote 242 bytes to /Users/madcricket/go/src/github.com/latchmihay/k8s-yaml-splitter/.k8s-yaml-splitter/Secret-kiali.yaml
Found! type: ConfigMap | apiVersion: v1 | name: istio-galley-configuration | namespace: istio-system
* Writing /Users/lnm0811/go/src/github.com/latchmihay/k8s-yaml-splitter/.k8s-yaml-splitter/ConfigMap-istio-galley-configuration.yaml
* Wrote 3308 bytes to /Users/madcricket/go/src/github.com/latchmihay/k8s-yaml-splitter/.k8s-yaml-splitter/ConfigMap-istio-galley-configuration.yaml
```