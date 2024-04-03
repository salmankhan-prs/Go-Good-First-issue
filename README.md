<h1 align='center'><span style='font-size:36px;'>Go Good First Issue</span></h1>

<div align='center' style='font-size:20px;'>Welcome to Go's good first issue list! These are issues that are great for new contributors to the Go community. It is updated every 24 hours.</div> <br>



<div align='center'>Last updated at April 3, 2024 01:49 UTC.</div>


## kubernetes/kubernetes <span style='color:#F1C40F'>(106.3K ⭐️)</span>

- [sig-windows-gce test jobs are failing consistently for a long time](https://github.com/kubernetes/kubernetes/issues/124047)


## ollama/ollama <span style='color:#F1C40F'>(52.3K ⭐️)</span>

- [Document `OLLAMA_DEBUG` in `ollama serve` `-h` docs](https://github.com/ollama/ollama/issues/3401)

- [Allow `api.Client` to be constructed using URL & http.Client](https://github.com/ollama/ollama/issues/2948)


## purpleidea/mgmt <span style='color:#F1C40F'>(3.4K ⭐️)</span>

- [Add a test for go mod madness](https://github.com/purpleidea/mgmt/issues/749)


## cockroachdb/cockroach <span style='color:#F1C40F'>(28.9K ⭐️)</span>

- [backupccl/console: don't report estimated time remaining if progress is less than <5%](https://github.com/cockroachdb/cockroach/issues/119873)


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

- [[UI Feature] The UI should show that a task is pending due to `cache_serialize=True`](https://github.com/flyteorg/flyte/issues/5096)

- [[Core feature] Override task `secret_requests` using `with_overrides`](https://github.com/flyteorg/flyte/issues/5085)

- [[Core feature] Implement `pyflyte recover`](https://github.com/flyteorg/flyte/issues/5049)

- [[BUG] sql_alchemy plugin doesn't support pandas version > 2.1.4](https://github.com/flyteorg/flyte/issues/5031)


## a-h/templ <span style='color:#F1C40F'>(6.2K ⭐️)</span>

- [cmd: add `--notify` to `generate --watch` in cli](https://github.com/a-h/templ/issues/656)

- [feature: add folding to LSP](https://github.com/a-h/templ/issues/650)

- [bug: backslashes in constant string attributes cannot be generated](https://github.com/a-h/templ/issues/624)


## grafana/mimir <span style='color:#F1C40F'>(3.7K ⭐️)</span>

- [Make /ingester/shutdown react on POST method only](https://github.com/grafana/mimir/issues/7623)


## konveyor/analyzer-lsp <span style='color:#F1C40F'>(7 ⭐️)</span>

- [[Bug] Add openapi struct fields for dependency condition](https://github.com/konveyor/analyzer-lsp/issues/551)

- [[BUG] Adding Title and Description to capability](https://github.com/konveyor/analyzer-lsp/issues/550)


## Frank-Mayer/ohmygosh <span style='color:#F1C40F'>(1 ⭐️)</span>

- [add curly braces syntax for variables](https://github.com/Frank-Mayer/ohmygosh/issues/5)

- [add pipe all `|&`](https://github.com/Frank-Mayer/ohmygosh/issues/4)

- [/dev/zero implementation is wrong](https://github.com/Frank-Mayer/ohmygosh/issues/3)


## zitadel/zitadel <span style='color:#F1C40F'>(6.7K ⭐️)</span>

- [Instance settings title](https://github.com/zitadel/zitadel/issues/7632)

- [[Bug]: Translations overflow buttons in login UI](https://github.com/zitadel/zitadel/issues/7619)

- [[Bug]: Applications in a granted project not showing initially](https://github.com/zitadel/zitadel/issues/7613)

- [[Bug]: "Default settings" missing in console mobile view](https://github.com/zitadel/zitadel/issues/7574)


## hashicorp/terraform-provider-aws <span style='color:#F1C40F'>(9.4K ⭐️)</span>

- [[Enhancement]: Data Sources for Cognito User Pool](https://github.com/hashicorp/terraform-provider-aws/issues/36087)

- [[Enhancement]: r/aws_elasticache_serverless_cache: Support minimum cache usage limits](https://github.com/hashicorp/terraform-provider-aws/issues/36624)

- [[Docs]: Route53 Resolver Rule Resource Documentation Missing protocol Argument of the target_ip Object](https://github.com/hashicorp/terraform-provider-aws/issues/36439)

- [[Bug]: modification of cache_usage_limits (ecpu_per_second or data_storage maximum) in aws_elasticache_serverless_cache results in replacement](https://github.com/hashicorp/terraform-provider-aws/issues/36317)


## kyverno/kyverno <span style='color:#F1C40F'>(5.0K ⭐️)</span>

- [[Feature] Boolean for disabling testing pod or increase the sleep time](https://github.com/kyverno/kyverno/issues/9866)


## coder/coder <span style='color:#F1C40F'>(6.7K ⭐️)</span>

- [Deprecate gauge metrics named like counters](https://github.com/coder/coder/issues/12744)

- [AlecAivazis/survey is no longer maintained](https://github.com/coder/coder/issues/12720)

- [bug: ci workflow being cancelled on `main`](https://github.com/coder/coder/issues/12532)


## stacklok/minder <span style='color:#F1C40F'>(181 ⭐️)</span>

- [The repository registered field is not set](https://github.com/stacklok/minder/issues/2628)

- [Update the ListRemoteRepositoriesFromProvider response with a field saying which repo is registered and which not](https://github.com/stacklok/minder/issues/2627)

- [Move to log/slog logging](https://github.com/stacklok/minder/issues/2502)


## devhatt/pet-dex-backend <span style='color:#F1C40F'>(35 ⭐️)</span>

- [Melhor implementação das migrations](https://github.com/devhatt/pet-dex-backend/issues/96)


## prysmaticlabs/prysm <span style='color:#F1C40F'>(3.3K ⭐️)</span>

- [Multiple network flags should prevent the BN to start.](https://github.com/prysmaticlabs/prysm/issues/13801)


## osmosis-labs/osmosis <span style='color:#F1C40F'>(852 ⭐️)</span>

- [CI job to auto replace assetlists at some interval](https://github.com/osmosis-labs/osmosis/issues/7859)

- [Broken Links Job Gets Rate Limited](https://github.com/osmosis-labs/osmosis/issues/7877)


## matheusgomes28/urchin <span style='color:#F1C40F'>(23 ⭐️)</span>

- [Abstract image file access ](https://github.com/matheusgomes28/urchin/issues/65)

- [Correct return errors from Post endpoints](https://github.com/matheusgomes28/urchin/issues/67)

- [Add system tests for the image upload endpoint](https://github.com/matheusgomes28/urchin/issues/69)

- [Plz increase test coverage 🧪 ](https://github.com/matheusgomes28/urchin/issues/43)

- [Remove the image folder configuration](https://github.com/matheusgomes28/urchin/issues/66)

- [Remove test files from coverage reports](https://github.com/matheusgomes28/urchin/issues/61)

- [Serve Proper Error Pages](https://github.com/matheusgomes28/urchin/issues/58)


## getporter/porter <span style='color:#F1C40F'>(1.1K ⭐️)</span>

- [add brew tap for porter](https://github.com/getporter/porter/issues/3045)


## open-telemetry/opentelemetry-collector <span style='color:#F1C40F'>(3.8K ⭐️)</span>

- [[component] Reject `component.ID`s that are too long](https://github.com/open-telemetry/opentelemetry-collector/issues/9872)

- [[CI] Changes to workflows are not tested on PRs](https://github.com/open-telemetry/opentelemetry-collector/issues/9676)


## microsoft/retina <span style='color:#F1C40F'>(2.3K ⭐️)</span>

- [Investigate and optimize slow Windows CodeQL runs](https://github.com/microsoft/retina/issues/200)

- [Make Signing and Signed-off-By requirements clear in Contributing docs and README ](https://github.com/microsoft/retina/issues/197)

- [Fix timestamp of flows](https://github.com/microsoft/retina/issues/204)

- [Duplicate import of flow library](https://github.com/microsoft/retina/issues/98)

- [duplicate windows metric: remove `windows_hns_stats`, keep `forward_count`](https://github.com/microsoft/retina/issues/92)

- [optimize build pipeline to minimize code duplication. ](https://github.com/microsoft/retina/issues/137)

- [remove irrelevant metric labels for dns and apiserver latency](https://github.com/microsoft/retina/issues/75)

- [API Server Latency buckets are skewed in large scale clusters](https://github.com/microsoft/retina/issues/82)

- [Support scheduling distributed captures](https://github.com/microsoft/retina/issues/88)

- [Evaluate security context/caps ](https://github.com/microsoft/retina/issues/93)

- [edge cases for MetricConfiguration CRD](https://github.com/microsoft/retina/issues/74)


## jozu-ai/kitops <span style='color:#F1C40F'>(44 ⭐️)</span>

- [CLI UX: Sorting commands & Flags](https://github.com/jozu-ai/kitops/issues/76)

- [CLI UX: Pimp help commands](https://github.com/jozu-ai/kitops/issues/77)

- [CLI UX: Customize warning, errors, special messages](https://github.com/jozu-ai/kitops/issues/78)

- [Support shell completions for commands](https://github.com/jozu-ai/kitops/issues/79)


## argoproj/argo-cd <span style='color:#F1C40F'>(15.9K ⭐️)</span>

- [CLI : add sync policy option `manual` for app create command to make it similar to options provided in Web console](https://github.com/argoproj/argo-cd/issues/17447)


## kubeflow/katib <span style='color:#F1C40F'>(1.4K ⭐️)</span>

- [Whether the hyperparameter search algorithm will refer to the value of additionalMetricNames](https://github.com/kubeflow/katib/issues/2286)


## argoproj/argo-workflows <span style='color:#F1C40F'>(14.2K ⭐️)</span>

- [Support displaying line numbers in object-editor](https://github.com/argoproj/argo-workflows/issues/12807)

- [local image build: support arm and amd64](https://github.com/argoproj/argo-workflows/issues/12750)


## syself/cluster-api-provider-hetzner <span style='color:#F1C40F'>(491 ⭐️)</span>

- [Condition indicating that HetznerBareMetalHost is deleting](https://github.com/syself/cluster-api-provider-hetzner/issues/1229)


## Azure/azure-dev <span style='color:#F1C40F'>(349 ⭐️)</span>

- [Operations for a given service (e.g. restore) should only verify presence of dependency tools needed for that service](https://github.com/Azure/azure-dev/issues/3533)


## ossf/scorecard <span style='color:#F1C40F'>(4.1K ⭐️)</span>

- [Feature: Probe whether repo has up-to-date CODEOWNERS](https://github.com/ossf/scorecard/issues/3931)

- [Dangerous Workflow: some user input are not being detected as untrusted input.](https://github.com/ossf/scorecard/issues/3915)


## lightningnetwork/lnd <span style='color:#F1C40F'>(7.4K ⭐️)</span>

- [[bug]: lncli listpayments --count_total_payments should honor --creation_date_start --creation_date_end](https://github.com/lightningnetwork/lnd/issues/8530)


## gnolang/gno <span style='color:#F1C40F'>(814 ⭐️)</span>

- [bug: no error handling for invalid label](https://github.com/gnolang/gno/issues/1749)


## tektoncd/pipeline <span style='color:#F1C40F'>(8.3K ⭐️)</span>

- [doc: Make it clearer what the `auth` documentation is about](https://github.com/tektoncd/pipeline/issues/7737)


## kubernetes-sigs/jobset <span style='color:#F1C40F'>(90 ⭐️)</span>

- [Add unit tests to jobset success policy functions](https://github.com/kubernetes-sigs/jobset/issues/491)

- [Add pod controller unit tests](https://github.com/kubernetes-sigs/jobset/issues/478)


## radius-project/radius <span style='color:#F1C40F'>(1.3K ⭐️)</span>

- [`rad group switch` doesn't print success message](https://github.com/radius-project/radius/issues/7395)

- [`rad group show` doesn't show the current group](https://github.com/radius-project/radius/issues/7391)


## ksctl/ksctl <span style='color:#F1C40F'>(229 ⭐️)</span>

- [[Enhancement Proposal]: e2e mock test lack functionality tests](https://github.com/ksctl/ksctl/issues/316)


## ysugimoto/falco <span style='color:#F1C40F'>(88 ⭐️)</span>

- [[BUG] Error field of simulator output is an empty object](https://github.com/ysugimoto/falco/issues/292)

- [[Feature] Access to functional subroutine return values in tests](https://github.com/ysugimoto/falco/issues/288)

- [[BUG] Documentation checker tool misses variables containing `{NAME}`](https://github.com/ysugimoto/falco/issues/287)

