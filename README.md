<h1 align='center'><span style='font-size:36px;'>Go Good First Issue</span></h1>

<div align='center' style='font-size:20px;'>Welcome to Go's good first issue list! These are issues that are great for new contributors to the Go community. It is updated every 24 hours.</div> <br>



<div align='center'>Last updated at April 1, 2024 01:48 UTC.</div>


## vmware-tanzu/secrets-manager <span style='color:#F1C40F'>(124 ⭐️)</span>

- [VSecM Init containar shall be able to (optionally) decrypt a file mounted to the workload, provided it can fech an AES or age decrption key](https://github.com/vmware-tanzu/secrets-manager/issues/765)

- [VSecM shall be able to create an AES key, or an age key pair and register it as secrets for a workload; so that if two workload share the same AES key, they can decrypt a shared encrypted file](https://github.com/vmware-tanzu/secrets-manager/issues/764)

- [VSecM Init Container shall be able to wait on a Kubernetes secret](https://github.com/vmware-tanzu/secrets-manager/issues/763)

- [VSecM Sidecar shall be able to interpolate secrets to a file that is templatized using go templating language](https://github.com/vmware-tanzu/secrets-manager/issues/760)

- [VSecM Sidecar should be able to (optionally) create one file per secret key](https://github.com/vmware-tanzu/secrets-manager/issues/761)

- [VSecM Init Container should be (optionally cofigurable) able to watch the changes in a volume or a file instead of VSecM Safe](https://github.com/vmware-tanzu/secrets-manager/issues/759)

- [VSecM Init Container shall wait for a (configurable) grace period before exiting and yielding control](https://github.com/vmware-tanzu/secrets-manager/issues/766)

- [Nobody uses the `Stats()` function in VSecM Safe; have a /stats endpoint to use it. Sentinel can be used to query this endpoint, and the endpoint will display these stats in JSON form.](https://github.com/vmware-tanzu/secrets-manager/issues/634)

- [vsecm components (and SPIRE components too) should be able to have customizable labels and annotations; helm charts should have the option to provide it](https://github.com/vmware-tanzu/secrets-manager/issues/635)

- [Sentinel should do a basic validation of init commands before running them](https://github.com/vmware-tanzu/secrets-manager/issues/641)

- [`safe -l -e` should be able to get the secrets per workload (instead of listing all secrets) enable a command line option for that](https://github.com/vmware-tanzu/secrets-manager/issues/642)

- [write integration tests for multiple ns feature](https://github.com/vmware-tanzu/secrets-manager/issues/649)

- [write integration tests around this: https://github.com/vmware-tanzu/secrets-manager/pull/564#pullrequestreview-1920918083](https://github.com/vmware-tanzu/secrets-manager/issues/651)


## grafana/mimir <span style='color:#F1C40F'>(3.7K ⭐️)</span>

- [Make /ingester/shutdown react on POST method only](https://github.com/grafana/mimir/issues/7623)


## jaegertracing/jaeger <span style='color:#F1C40F'>(19.3K ⭐️)</span>

- [Fix breaking changes in upgrading OTEL Collector](https://github.com/jaegertracing/jaeger/issues/5302)


## pipe-cd/pipecd <span style='color:#F1C40F'>(945 ⭐️)</span>

- [Remove deprecated CloudProvider from model](https://github.com/pipe-cd/pipecd/issues/4806)


## apache/dubbo-go <span style='color:#F1C40F'>(4.6K ⭐️)</span>

- [Sample and test-case tasks](https://github.com/apache/dubbo-go/issues/2608)


## siglens/siglens <span style='color:#F1C40F'>(918 ⭐️)</span>

- [[BUG] Visualization tab doesn't update](https://github.com/siglens/siglens/issues/646)


## cockroachdb/cockroach <span style='color:#F1C40F'>(28.9K ⭐️)</span>

- [backupccl/console: don't report estimated time remaining if progress is less than <5%](https://github.com/cockroachdb/cockroach/issues/119873)


## NethermindEth/juno <span style='color:#F1C40F'>(348 ⭐️)</span>

- [Remove Goerli from Juno codebase](https://github.com/NethermindEth/juno/issues/1758)


## openbao/openbao <span style='color:#F1C40F'>(1.7K ⭐️)</span>

- [Fix Dependency Vulnerabilities For Alpha Release](https://github.com/openbao/openbao/issues/240)


## lightningnetwork/lnd <span style='color:#F1C40F'>(7.4K ⭐️)</span>

- [[bug]: lncli listpayments --count_total_payments should honor --creation_date_start --creation_date_end](https://github.com/lightningnetwork/lnd/issues/8530)


## argoproj/argo-workflows <span style='color:#F1C40F'>(14.2K ⭐️)</span>

- [Support displaying line number in object-editor](https://github.com/argoproj/argo-workflows/issues/12807)

- [local image build: support arm and amd64](https://github.com/argoproj/argo-workflows/issues/12750)


## open-telemetry/opentelemetry-collector-contrib <span style='color:#F1C40F'>(2.5K ⭐️)</span>

- [[processor/transform] Add debug log that prints current TransformContext](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/31912)


## ollama/ollama <span style='color:#F1C40F'>(51.8K ⭐️)</span>

- [Document `OLLAMA_DEBUG` in `ollama serve` `-h` docs](https://github.com/ollama/ollama/issues/3401)

- [Allow `api.Client` to be constructed using URL & http.Client](https://github.com/ollama/ollama/issues/2948)


## kubernetes/kubernetes <span style='color:#F1C40F'>(106.2K ⭐️)</span>

- [sig-windows-gce test jobs are failing consistently for a long time](https://github.com/kubernetes/kubernetes/issues/124047)


## kcp-dev/kcp <span style='color:#F1C40F'>(2.2K ⭐️)</span>

- [feature: add `kubectl ws -i`](https://github.com/kcp-dev/kcp/issues/3098)


## microsoft/retina <span style='color:#F1C40F'>(2.2K ⭐️)</span>

- [Duplicate import of flow library](https://github.com/microsoft/retina/issues/98)

- [Investigate and optimize slow Windows CodeQL runs](https://github.com/microsoft/retina/issues/200)

- [Make Signing and Signed-off-By requirements clear in Contributing docs and README ](https://github.com/microsoft/retina/issues/197)

- [duplicate windows metric: remove `windows_hns_stats`, keep `forward_count`](https://github.com/microsoft/retina/issues/92)

- [optimize build pipeline to minimize code duplication. ](https://github.com/microsoft/retina/issues/137)

- [remove irrelevant metric labels for dns and apiserver latency](https://github.com/microsoft/retina/issues/75)

- [API Server Latency buckets are skewed in large scale clusters](https://github.com/microsoft/retina/issues/82)

- [Support scheduling distributed captures](https://github.com/microsoft/retina/issues/88)

- [Evaluate security context/caps ](https://github.com/microsoft/retina/issues/93)

- [edge cases for MetricConfiguration CRD](https://github.com/microsoft/retina/issues/74)


## RamenDR/ramen <span style='color:#F1C40F'>(67 ⭐️)</span>

- [Cache addons resources](https://github.com/RamenDR/ramen/issues/1290)

- [rook-cluster: repository 'https://raw.githubusercontent.com/rook/rook/' not found](https://github.com/RamenDR/ramen/issues/1310)

- [Bad error handling in velero test when running velero backup/restore](https://github.com/RamenDR/ramen/issues/1281)

- [Upgrade csi-addons to 0.8.0](https://github.com/RamenDR/ramen/issues/1257)

- [Fake s3 store in tests doesn't use locks to store data in a map](https://github.com/RamenDR/ramen/issues/1289)

- [Drenv - create busybox with 2 pvcs in the samples](https://github.com/RamenDR/ramen/issues/1229)

- [Upgrade submariner to v0.17.0](https://github.com/RamenDR/ramen/issues/1263)

- [Upgrade olm to version v0.27.0](https://github.com/RamenDR/ramen/issues/1260)

- [ramenctl config times out waiting for policies to propagate to managed clusters](https://github.com/RamenDR/ramen/issues/1242)


## argoproj/argo-cd <span style='color:#F1C40F'>(15.9K ⭐️)</span>

- [CLI : add sync policy option `manual` for app create command to make it similar to options provided in Web console](https://github.com/argoproj/argo-cd/issues/17447)


## grendel-consulting/steampipe-plugin-kolide <span style='color:#F1C40F'>(1 ⭐️)</span>

- [GET /device_groups/{id}](https://github.com/grendel-consulting/steampipe-plugin-kolide/issues/30)

- [GET /device_groups/](https://github.com/grendel-consulting/steampipe-plugin-kolide/issues/29)

- [GET /auth_logs/{id}](https://github.com/grendel-consulting/steampipe-plugin-kolide/issues/28)

- [GET /auth_logs/](https://github.com/grendel-consulting/steampipe-plugin-kolide/issues/27)

- [GET /person_groups/{id}](https://github.com/grendel-consulting/steampipe-plugin-kolide/issues/26)

- [GET /person_groups/](https://github.com/grendel-consulting/steampipe-plugin-kolide/issues/25)


## flyteorg/flyte <span style='color:#F1C40F'>(4.7K ⭐️)</span>

- [[Core feature] Override task `secret_requests` using `with_overrides`](https://github.com/flyteorg/flyte/issues/5085)

- [[UI Feature] The UI should show that a task is pending due to `cache_serialize=True`](https://github.com/flyteorg/flyte/issues/5096)

- [[Core feature] Implement `pyflyte recover`](https://github.com/flyteorg/flyte/issues/5049)

- [[BUG] sql_alchemy plugin doesn't support pandas version > 2.1.4](https://github.com/flyteorg/flyte/issues/5031)


## elastic/cloudbeat <span style='color:#F1C40F'>(34 ⭐️)</span>

- [CIS K8s boolean cli arguments rules perform case sensitive check](https://github.com/elastic/cloudbeat/issues/2007)


## redis/rueidis <span style='color:#F1C40F'>(2.2K ⭐️)</span>

- [Fix code scanning alert - Incorrect conversion between integer types](https://github.com/redis/rueidis/issues/513)

- [feature request: adding an option to disable dial on creating client](https://github.com/redis/rueidis/issues/489)

- [Need to improve the test coverage on rueidiscompat](https://github.com/redis/rueidis/issues/487)


## cosmos/cosmos-sdk <span style='color:#F1C40F'>(5.9K ⭐️)</span>

- [Remove x/exp import](https://github.com/cosmos/cosmos-sdk/issues/19892)


## thanos-io/thanos <span style='color:#F1C40F'>(12.5K ⭐️)</span>

- [Tracing: Add missing sampler types](https://github.com/thanos-io/thanos/issues/7230)


## litmuschaos/litmus <span style='color:#F1C40F'>(4.2K ⭐️)</span>

- [Need to formatting text in probe card](https://github.com/litmuschaos/litmus/issues/4560)

- [Add pre-commit check to find and delete unused strings in strings.en.yaml file](https://github.com/litmuschaos/litmus/issues/4550)

- [Add Fuzzing test suites in Litmus](https://github.com/litmuschaos/litmus/issues/4548)

- [Save button not working when no "resilience" probe is added while creating experiments, this should throw an error toaster](https://github.com/litmuschaos/litmus/issues/4551)

- [In gitops, access key is added to a plain field, this should be changed to a hidden field like password](https://github.com/litmuschaos/litmus/issues/4549)


## kyverno/kyverno <span style='color:#F1C40F'>(5.0K ⭐️)</span>

- [[Feature] Boolean for disabling testing pod or increase the sleep time](https://github.com/kyverno/kyverno/issues/9866)


## karmada-io/karmada <span style='color:#F1C40F'>(4.1K ⭐️)</span>

- [Update the worng example comment in the e2e suite ](https://github.com/karmada-io/karmada/issues/4775)


## FerretDB/FerretDB <span style='color:#F1C40F'>(8.4K ⭐️)</span>

- [`checkdocs` should check if linked issues are still open](https://github.com/FerretDB/FerretDB/issues/4171)


## flipt-io/flipt <span style='color:#F1C40F'>(3.3K ⭐️)</span>

- [[FLI-938] Allow passing a starting dir `flipt validate` ](https://github.com/flipt-io/flipt/issues/2916)

- [[FLI-907] config: add ability to specify env vars](https://github.com/flipt-io/flipt/issues/2844)

- [[FLI-915] UI: Copy As Dropdown](https://github.com/flipt-io/flipt/issues/2863)


## k8sgpt-ai/k8sgpt-operator <span style='color:#F1C40F'>(244 ⭐️)</span>

- [[Feature]: Generate results CRD in related to issue namespace](https://github.com/k8sgpt-ai/k8sgpt-operator/issues/390)


## zitadel/zitadel <span style='color:#F1C40F'>(6.7K ⭐️)</span>

- [Instance settings title](https://github.com/zitadel/zitadel/issues/7632)

- [[Bug]: Translations overflow buttons in login UI](https://github.com/zitadel/zitadel/issues/7619)

- [[Bug]: Applications in a granted project not showing initially](https://github.com/zitadel/zitadel/issues/7613)

- [[Bug]: "Default settings" missing in console mobile view](https://github.com/zitadel/zitadel/issues/7574)


## koordinator-sh/koordinator <span style='color:#F1C40F'>(1.2K ⭐️)</span>

- [[proposal] Upgrade ginkgo to 2.x](https://github.com/koordinator-sh/koordinator/issues/1964)


## keploy/keploy <span style='color:#F1C40F'>(3.3K ⭐️)</span>

- [[bug]: TestResult summary overlap issue at certain media sizes](https://github.com/keploy/keploy/issues/1698)

- [[bug]: test mode panics in case of empty test case ](https://github.com/keploy/keploy/issues/1755)

- [[bug]:  incorrect routing on blog website ](https://github.com/keploy/keploy/issues/1733)

- [[bug]: table titles are not visible in dark mode](https://github.com/keploy/keploy/issues/1756)

- [[feature]: update all the frameworks logo on landing page](https://github.com/keploy/keploy/issues/1734)

- [[bug]: copy button doesn't copy to clipboard for Safari](https://github.com/keploy/keploy/issues/1706)

- [[bug]: Invalid DOM Properties Warnings](https://github.com/keploy/keploy/issues/1681)

- [[bug]: Unresponsive CLI one click command](https://github.com/keploy/keploy/issues/1672)

- [[refactor]: Star us on github button ](https://github.com/keploy/keploy/issues/1670)

- [[bug]  Command built from user-controlled sources](https://github.com/keploy/keploy/issues/1661)

- [[feature]:  Addition of version argument to installation script](https://github.com/keploy/keploy/issues/1639)

