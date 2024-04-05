<h1 align='center'><span style='font-size:36px;'>Go Good First Issue</span></h1>

<div align='center' style='font-size:20px;'>Welcome to Go's good first issue list! These are issues that are great for new contributors to the Go community. It is updated every 24 hours.</div> <br>



<div align='center'>Last updated at April 5, 2024 01:48 UTC.</div>


## kopia/kopia <span style='color:#F1C40F'>(6.2K ⭐️)</span>

- [Snapshot retention is deleting more snapshots as defined by policy in case system was pausing snapshots](https://github.com/kopia/kopia/issues/3758)


## vmware-tanzu/secrets-manager <span style='color:#F1C40F'>(127 ⭐️)</span>

- [dynamic log levels: vsecm safe, and sentinel shall have APIs to update their log levels at runtime without killing the pods.](https://github.com/vmware-tanzu/secrets-manager/issues/708)

- [a cli to check if a secret exists (like, return true/false based on if there is an assigned secret to a workload](https://github.com/vmware-tanzu/secrets-manager/issues/706)

- [VSecM Init Container shall wait for a (configurable) grace period before exiting and yielding control](https://github.com/vmware-tanzu/secrets-manager/issues/766)

- [VSecM Init containar shall be able to (optionally) decrypt a file mounted to the workload, provided it can fech an AES or age decrption key](https://github.com/vmware-tanzu/secrets-manager/issues/765)

- [VSecM shall be able to create an AES key, or an age key pair and register it as secrets for a workload; so that if two workload share the same AES key, they can decrypt a shared encrypted file](https://github.com/vmware-tanzu/secrets-manager/issues/764)

- [VSecM Init Container shall be able to wait on a Kubernetes secret](https://github.com/vmware-tanzu/secrets-manager/issues/763)

- [VSecM Sidecar shall be able to interpolate secrets to a file that is templatized using go templating language](https://github.com/vmware-tanzu/secrets-manager/issues/760)

- [VSecM Sidecar should be able to (optionally) create one file per secret key](https://github.com/vmware-tanzu/secrets-manager/issues/761)

- [VSecM Init Container should be (optionally cofigurable) able to watch the changes in a volume or a file instead of VSecM Safe](https://github.com/vmware-tanzu/secrets-manager/issues/759)

- [Nobody uses the `Stats()` function in VSecM Safe; have a /stats endpoint to use it. Sentinel can be used to query this endpoint, and the endpoint will display these stats in JSON form.](https://github.com/vmware-tanzu/secrets-manager/issues/634)

- [vsecm components (and SPIRE components too) should be able to have customizable labels and annotations; helm charts should have the option to provide it](https://github.com/vmware-tanzu/secrets-manager/issues/635)

- [Sentinel should do a basic validation of init commands before running them](https://github.com/vmware-tanzu/secrets-manager/issues/641)

- [`safe -l -e` should be able to get the secrets per workload (instead of listing all secrets) enable a command line option for that](https://github.com/vmware-tanzu/secrets-manager/issues/642)


## flyteorg/flyte <span style='color:#F1C40F'>(4.7K ⭐️)</span>

- [Add a securitySchemes to swagger.json](https://github.com/flyteorg/flyte/issues/5160)

- [[UI Feature] The UI should show that a task is pending due to `cache_serialize=True`](https://github.com/flyteorg/flyte/issues/5096)

- [[Core feature] Override task `secret_requests` using `with_overrides`](https://github.com/flyteorg/flyte/issues/5085)

- [[Core feature] Implement `pyflyte recover`](https://github.com/flyteorg/flyte/issues/5049)

- [[BUG] sql_alchemy plugin doesn't support pandas version > 2.1.4](https://github.com/flyteorg/flyte/issues/5031)


## hashicorp/terraform-provider-kubernetes <span style='color:#F1C40F'>(1.5K ⭐️)</span>

- [SecurityContext `procMount` Support](https://github.com/hashicorp/terraform-provider-kubernetes/issues/2451)

- [Init container doesn't support restart_policy option (to transform it in a sidecar container)](https://github.com/hashicorp/terraform-provider-kubernetes/issues/2446)


## TBD54566975/ftl <span style='color:#F1C40F'>(15 ⭐️)</span>

- [Add `ftl dev --background-after-deploy` or alter existing `--background` to have this behaviour](https://github.com/TBD54566975/ftl/issues/1168)

- [Get rid of `all_{controllers,runners,ingress_routes}` from `GetStatus()`](https://github.com/TBD54566975/ftl/issues/1171)

- [`runInBackground` should return errors](https://github.com/TBD54566975/ftl/issues/1170)

- [Replace use of `docker` CLI with Docker API](https://github.com/TBD54566975/ftl/issues/1169)

- [Extract databases in go schema](https://github.com/TBD54566975/ftl/issues/1120)

- [Multiple usages of `config` usages show up in schema](https://github.com/TBD54566975/ftl/issues/1121)

- [Add quickstart to the readme and ftl help](https://github.com/TBD54566975/ftl/issues/1033)


## coder/coder <span style='color:#F1C40F'>(6.7K ⭐️)</span>

- [Deprecate gauge metrics named like counters](https://github.com/coder/coder/issues/12744)

- [AlecAivazis/survey is no longer maintained](https://github.com/coder/coder/issues/12720)


## istio/istio <span style='color:#F1C40F'>(34.9K ⭐️)</span>

- [istioctl precheck: add unit tests](https://github.com/istio/istio/issues/50255)


## hashicorp/terraform-provider-aws <span style='color:#F1C40F'>(9.4K ⭐️)</span>

- [[Enhancement]: r/aws_elasticache_serverless_cache: Support minimum cache usage limits](https://github.com/hashicorp/terraform-provider-aws/issues/36624)

- [[Docs]: Route53 Resolver Rule Resource Documentation Missing protocol Argument of the target_ip Object](https://github.com/hashicorp/terraform-provider-aws/issues/36439)

- [[Bug]: modification of cache_usage_limits (ecpu_per_second or data_storage maximum) in aws_elasticache_serverless_cache results in replacement](https://github.com/hashicorp/terraform-provider-aws/issues/36317)


## derailed/k9s <span style='color:#F1C40F'>(24.6K ⭐️)</span>

- [Make Icons / Emoji customizable](https://github.com/derailed/k9s/issues/2624)


## ollama/ollama <span style='color:#F1C40F'>(52.7K ⭐️)</span>

- [Document `OLLAMA_DEBUG` in `ollama serve` `-h` docs](https://github.com/ollama/ollama/issues/3401)

- [Allow `api.Client` to be constructed using URL & http.Client](https://github.com/ollama/ollama/issues/2948)


## redis/rueidis <span style='color:#F1C40F'>(2.2K ⭐️)</span>

- [Fix code scanning alert - Incorrect conversion between integer types](https://github.com/redis/rueidis/issues/513)


## Kong/kubernetes-ingress-controller <span style='color:#F1C40F'>(2.1K ⭐️)</span>

- [Get rid of `tcpMutex` and `tlsMutex` from integration tests](https://github.com/Kong/kubernetes-ingress-controller/issues/5757)


## buildpacks/pack <span style='color:#F1C40F'>(2.4K ⭐️)</span>

- [Remove validation of deprecated `io.buildpacks.stack.id`](https://github.com/buildpacks/pack/issues/2104)


## kubernetes/kubernetes <span style='color:#F1C40F'>(106.3K ⭐️)</span>

- [sig-windows-gce test jobs are failing consistently for a long time](https://github.com/kubernetes/kubernetes/issues/124047)


## microsoft/retina <span style='color:#F1C40F'>(2.3K ⭐️)</span>

- [Enrich flows based on non-primary Pod IPs](https://github.com/microsoft/retina/issues/224)

- [Fix timestamp of flows](https://github.com/microsoft/retina/issues/204)

- [Duplicate import of flow library](https://github.com/microsoft/retina/issues/98)

- [duplicate windows metric: remove `windows_hns_stats`, keep `forward_count`](https://github.com/microsoft/retina/issues/92)

- [optimize build pipeline to minimize code duplication. ](https://github.com/microsoft/retina/issues/137)

- [remove irrelevant metric labels for dns and apiserver latency](https://github.com/microsoft/retina/issues/75)

- [API Server Latency buckets are skewed in large scale clusters](https://github.com/microsoft/retina/issues/82)

- [Support scheduling distributed captures](https://github.com/microsoft/retina/issues/88)

- [Evaluate security context/caps ](https://github.com/microsoft/retina/issues/93)

- [edge cases for MetricConfiguration CRD](https://github.com/microsoft/retina/issues/74)


## kyverno/kyverno <span style='color:#F1C40F'>(5.1K ⭐️)</span>

- [[Bug] generate with synchronize: true does not detect removal of secret/cm keys](https://github.com/kyverno/kyverno/issues/9972)

- [[Feature] Boolean for disabling testing pod or increase the sleep time](https://github.com/kyverno/kyverno/issues/9866)


## argoproj/argo-workflows <span style='color:#F1C40F'>(14.2K ⭐️)</span>

- [local image build: support arm and amd64](https://github.com/argoproj/argo-workflows/issues/12750)


## kubernetes/ingress-nginx <span style='color:#F1C40F'>(16.6K ⭐️)</span>

- [strict-validate-path-type does not allow period/dot/. in Exact or Prefix path](https://github.com/kubernetes/ingress-nginx/issues/11176)


## Azure/azure-dev <span style='color:#F1C40F'>(351 ⭐️)</span>

- [Operations for a given service (e.g. restore) should only verify presence of dependency tools needed for that service](https://github.com/Azure/azure-dev/issues/3533)


## radius-project/radius <span style='color:#F1C40F'>(1.3K ⭐️)</span>

- [`rad group switch` doesn't print success message](https://github.com/radius-project/radius/issues/7395)

- [`rad group show` doesn't show the current group](https://github.com/radius-project/radius/issues/7391)


## kedacore/keda <span style='color:#F1C40F'>(7.7K ⭐️)</span>

- [Document which (observability) metrics are initialized and which not](https://github.com/kedacore/keda/issues/5609)


## ossf/scorecard <span style='color:#F1C40F'>(4.1K ⭐️)</span>

- [Feature: Probe whether repo has up-to-date CODEOWNERS](https://github.com/ossf/scorecard/issues/3931)


## kubernetes-sigs/aws-load-balancer-controller <span style='color:#F1C40F'>(3.7K ⭐️)</span>

- [add the ability to add a runtimeClassName parameter in values.yaml to enable gvisor or kata containers](https://github.com/kubernetes-sigs/aws-load-balancer-controller/issues/3628)

- [Helm chart does not allow --load-balancer-class flag to be set](https://github.com/kubernetes-sigs/aws-load-balancer-controller/issues/3629)


## numaproj/numaflow <span style='color:#F1C40F'>(907 ⭐️)</span>

- [Docs: Message Headers](https://github.com/numaproj/numaflow/issues/1598)

- [User Defined Source: Add functionality to tag messages](https://github.com/numaproj/numaflow/issues/1605)

- [Bypass transformer calls if source transformer is not used](https://github.com/numaproj/numaflow/issues/1612)


## EchoVault/EchoVault <span style='color:#F1C40F'>(88 ⭐️)</span>

- [Create Godoc documentation for Embedded API](https://github.com/EchoVault/EchoVault/issues/16)


## cloudnative-pg/cloudnative-pg <span style='color:#F1C40F'>(3.2K ⭐️)</span>

- [[Feature]: Generate hostname as service.namespace in postgresql uri/jdbc-uri in credentials secret](https://github.com/cloudnative-pg/cloudnative-pg/issues/4062)

- [[Feature]: Enable configuration of `wal_log_hints`](https://github.com/cloudnative-pg/cloudnative-pg/issues/4041)

- [[Feature]: Display PDB in `cnpg status`](https://github.com/cloudnative-pg/cloudnative-pg/issues/4019)


## filecoin-project/lotus <span style='color:#F1C40F'>(2.8K ⭐️)</span>

- [Option to back-fill events by re-execution of messages](https://github.com/filecoin-project/lotus/issues/11744)

- [Avoid disputing WindowedPoSt messages when the target Miner has no balance to pay rewards](https://github.com/filecoin-project/lotus/issues/11715)


## konveyor/analyzer-lsp <span style='color:#F1C40F'>(7 ⭐️)</span>

- [[Bug] Add openapi struct fields for dependency condition](https://github.com/konveyor/analyzer-lsp/issues/551)

- [[BUG] Adding Title and Description to capability](https://github.com/konveyor/analyzer-lsp/issues/550)


## openziti/ziti <span style='color:#F1C40F'>(1.9K ⭐️)</span>

- [adopt single-port default vars in ziti create config](https://github.com/openziti/ziti/issues/1838)


## pingcap/tidb <span style='color:#F1C40F'>(36.0K ⭐️)</span>

- [`TestDumpBinaryTime` failed in Egypt time](https://github.com/pingcap/tidb/issues/52345)


## open-telemetry/opentelemetry-collector <span style='color:#F1C40F'>(3.8K ⭐️)</span>

- [[component] Reject `component.ID`s that are too long](https://github.com/open-telemetry/opentelemetry-collector/issues/9872)


## a-h/templ <span style='color:#F1C40F'>(6.2K ⭐️)</span>

- [feature: add folding to LSP](https://github.com/a-h/templ/issues/650)

- [bug: backslashes in constant string attributes cannot be generated](https://github.com/a-h/templ/issues/624)


## dapr/dapr <span style='color:#F1C40F'>(23.2K ⭐️)</span>

- [Additional Dapr metrics properties when subscriber sends `status=DROP` ](https://github.com/dapr/dapr/issues/7610)


## crossplane/crossplane <span style='color:#F1C40F'>(8.6K ⭐️)</span>

- [`validate` doesn't do tilde expansion](https://github.com/crossplane/crossplane/issues/5487)

- [crossplane beta validate does not work with Configurations using composition function dependencies](https://github.com/crossplane/crossplane/issues/5491)


## litmuschaos/litmus <span style='color:#F1C40F'>(4.2K ⭐️)</span>

- [Need to formatting text in probe card](https://github.com/litmuschaos/litmus/issues/4560)

- [Add pre-commit check to find and delete unused strings in strings.en.yaml file](https://github.com/litmuschaos/litmus/issues/4550)

- [Add Fuzzing test suites in Litmus](https://github.com/litmuschaos/litmus/issues/4548)

- [Save button not working when no "resilience" probe is added while creating experiments, this should throw an error toaster](https://github.com/litmuschaos/litmus/issues/4551)

- [In gitops, access key is added to a plain field, this should be changed to a hidden field like password](https://github.com/litmuschaos/litmus/issues/4549)


## RamenDR/ramen <span style='color:#F1C40F'>(67 ⭐️)</span>

- [ramenctl config times out waiting for policies to propagate to managed clusters](https://github.com/RamenDR/ramen/issues/1242)

- [rook-cluster: repository 'https://raw.githubusercontent.com/rook/rook/' not found](https://github.com/RamenDR/ramen/issues/1310)

- [Cache addons resources](https://github.com/RamenDR/ramen/issues/1290)

- [Bad error handling in velero test when running velero backup/restore](https://github.com/RamenDR/ramen/issues/1281)

- [Fake s3 store in tests doesn't use locks to store data in a map](https://github.com/RamenDR/ramen/issues/1289)

- [Drenv - create busybox with 2 pvcs in the samples](https://github.com/RamenDR/ramen/issues/1229)

- [Upgrade olm to version v0.27.0](https://github.com/RamenDR/ramen/issues/1260)


## keploy/keploy <span style='color:#F1C40F'>(3.3K ⭐️)</span>

- [[feature]: automatically add test reports folder to gitignore ](https://github.com/keploy/keploy/issues/1775)

- [[bug]: Build Warnings on Image Optimization ](https://github.com/keploy/keploy/issues/1770)

- [[bug]: Table of Contents Positioning Issue on Medium Screens](https://github.com/keploy/keploy/issues/1771)

- [[bug]: test mode panics in case of empty test case ](https://github.com/keploy/keploy/issues/1755)

- [[bug]:  incorrect routing on blog website ](https://github.com/keploy/keploy/issues/1733)

- [[bug]: table titles are not visible in dark mode](https://github.com/keploy/keploy/issues/1756)

- [[feature]: update all the frameworks logo on landing page](https://github.com/keploy/keploy/issues/1734)

- [[bug]: copy button doesn't copy to clipboard for Safari](https://github.com/keploy/keploy/issues/1706)

- [[bug]: Invalid DOM Properties Warnings](https://github.com/keploy/keploy/issues/1681)

- [[bug]: Unresponsive CLI one click command](https://github.com/keploy/keploy/issues/1672)

- [[refactor]: Star us on github button ](https://github.com/keploy/keploy/issues/1670)

- [[bug]  Command built from user-controlled sources](https://github.com/keploy/keploy/issues/1661)


## grafana/mimir <span style='color:#F1C40F'>(3.7K ⭐️)</span>

- [Make /ingester/shutdown react on POST method only](https://github.com/grafana/mimir/issues/7623)

