<h1 align='center'><span style='font-size:36px;'>Go Good First Issue</span></h1>

<div align='center' style='font-size:20px;'>Welcome to Go's good first issue list! These are issues that are great for new contributors to the Go community. It is updated every 24 hours.</div> <br>



<div align='center'>Last updated at April 27, 2024 01:49 UTC.</div>


## vmware-tanzu/secrets-manager <span style='color:#F1C40F'>(130 ⭐️)</span>

- [VSecM Init containar shall be able to (optionally) decrypt a file mounted to the workload, provided it can fech an AES or age decrption key](https://github.com/vmware-tanzu/secrets-manager/issues/765)

- [VSecM shall be able to create an AES key, or an age key pair and register it as secrets for a workload; so that if two workload share the same AES key, they can decrypt a shared encrypted file](https://github.com/vmware-tanzu/secrets-manager/issues/764)

- [vsecm sentinel should be (optionally) configured to require a user id at every command execution](https://github.com/vmware-tanzu/secrets-manager/issues/762)

- [VSecM Sidecar shall be able to interpolate secrets to a file that is templatized using go templating language](https://github.com/vmware-tanzu/secrets-manager/issues/760)

- [VSecM Sidecar should be able to (optionally) create one file per secret key](https://github.com/vmware-tanzu/secrets-manager/issues/761)

- [VSecM Init Container should be (optionally cofigurable) able to watch the changes in a volume or a file instead of pinging VSecM Safe](https://github.com/vmware-tanzu/secrets-manager/issues/759)

- [write tests: https://github.com/vmware-tanzu/secrets-manager/blob/main/app/sentinel/internal/safe/action_test.go](https://github.com/vmware-tanzu/secrets-manager/issues/874)

- [write tests: https://github.com/vmware-tanzu/secrets-manager/blob/main/app/sentinel/cmd/parse_test.go](https://github.com/vmware-tanzu/secrets-manager/issues/873)

- [write tests: https://github.com/vmware-tanzu/secrets-manager/blob/main/app/sentinel/cmd/help_test.go](https://github.com/vmware-tanzu/secrets-manager/issues/872)

- [write tests: https://github.com/vmware-tanzu/secrets-manager/blob/main/app/sentinel/background/initialization/validation_test.go](https://github.com/vmware-tanzu/secrets-manager/issues/871)

- [write tests: https://github.com/vmware-tanzu/secrets-manager/blob/main/app/sentinel/background/initialization/token_test.go](https://github.com/vmware-tanzu/secrets-manager/issues/870)

- [write tests: https://github.com/vmware-tanzu/secrets-manager/blob/main/app/sentinel/background/initialization/run_test.go](https://github.com/vmware-tanzu/secrets-manager/issues/869)

- [write tests: https://github.com/vmware-tanzu/secrets-manager/blob/main/app/sentinel/background/initialization/keystone_test.go](https://github.com/vmware-tanzu/secrets-manager/issues/868)


## TBD54566975/ftl <span style='color:#F1C40F'>(15 ⭐️)</span>

- [Validate that the secrets and config required by the schema exactly match defined secrets](https://github.com/TBD54566975/ftl/issues/1173)

- [Extend HTTP ingress to support `omitempty`](https://github.com/TBD54566975/ftl/issues/1254)

- [Think about whether events table serves the role of audit logs](https://github.com/TBD54566975/ftl/issues/1236)

- [Write a reaper that periodically cleans `$CACHE/ftl-runner`](https://github.com/TBD54566975/ftl/issues/1234)

- [Replace use of `docker` CLI with Docker API](https://github.com/TBD54566975/ftl/issues/1169)

- [Add `ftl dev --background-after-deploy` or alter existing `--background` to have this behaviour](https://github.com/TBD54566975/ftl/issues/1168)


## cloudnative-pg/cloudnative-pg <span style='color:#F1C40F'>(3.4K ⭐️)</span>

- [[Bug]: kubectl cnpg psql doesn't pass `--context` to kubectl](https://github.com/cloudnative-pg/cloudnative-pg/issues/4332)

- [[Feature]: support `allow_alter_system` from PostgreSQL 17](https://github.com/cloudnative-pg/cloudnative-pg/issues/4244)


## ollama/ollama <span style='color:#F1C40F'>(60.8K ⭐️)</span>

- [Document `OLLAMA_DEBUG` in `ollama serve` `-h` docs](https://github.com/ollama/ollama/issues/3401)


## radius-project/radius <span style='color:#F1C40F'>(1.4K ⭐️)</span>

- [migrate golang/mock to uber-go/mock](https://github.com/radius-project/radius/issues/7543)

- [rad env delete command removes environment and its deployed applications without confirming](https://github.com/radius-project/radius/issues/7529)


## lightningnetwork/lnd <span style='color:#F1C40F'>(7.5K ⭐️)</span>

- [[bug]: InboundFees are overwritten if they were not explicitly specified in lncli](https://github.com/lightningnetwork/lnd/issues/8614)

- [[feature]: Append the listchannels command with the short_chan_id representation (XXXX:XX:X) besides the 8 byte integer one.](https://github.com/lightningnetwork/lnd/issues/8650)


## cockroachdb/cockroach <span style='color:#F1C40F'>(29.1K ⭐️)</span>

- [o11y: unavailable ranges in console is coloured green](https://github.com/cockroachdb/cockroach/issues/122014)

- [server: bad certificate error spams the logs](https://github.com/cockroachdb/cockroach/issues/122622)


## jaegertracing/jaeger <span style='color:#F1C40F'>(19.4K ⭐️)</span>

- [Upgrade google.golang.org 1.63.0](https://github.com/jaegertracing/jaeger/issues/5330)

- [[Bug]: Jaeger not respecting JAEGER_DISABLED=true](https://github.com/jaegertracing/jaeger/issues/5385)

- [[test] Crossdock tests do not need to build Docker images for all architectures](https://github.com/jaegertracing/jaeger/issues/5350)


## numaproj/numaflow <span style='color:#F1C40F'>(917 ⭐️)</span>

- [Bypass transformer calls if source transformer is not used](https://github.com/numaproj/numaflow/issues/1612)


## gofr-dev/gofr <span style='color:#F1C40F'>(867 ⭐️)</span>

- [favicon.ico should be overridable ](https://github.com/gofr-dev/gofr/issues/474)

- [Enhance the DB logs ](https://github.com/gofr-dev/gofr/issues/472)

- [Improve details in debug log for Migrations](https://github.com/gofr-dev/gofr/issues/477)


## open-telemetry/opentelemetry-collector-contrib <span style='color:#F1C40F'>(2.6K ⭐️)</span>

- [[Weekly report] Action failed due to message body being too long](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/32655)

- [Update weekly report to include waiting-for-code-owners](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/32490)

- [[exporter/elasticsearchexporter] Push failures not reported in metrics](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/32302)


## openfga/openfga <span style='color:#F1C40F'>(2.3K ⭐️)</span>

- [Improve cyclomatic complexity of func (s *ServerContext) Run(ctx context.Context, config *serverconfig.Config)](https://github.com/openfga/openfga/issues/1575)

- [Replace `github.com/hashicorp/go-multierror` with `errors.Join`](https://github.com/openfga/openfga/issues/1556)


## containers/podman <span style='color:#F1C40F'>(21.7K ⭐️)</span>

- [`docker login` with docker.io creds  "successfully" logs-into `registry.fedoraproject.org` then fails to push to `registry-1.docker.io`](https://github.com/containers/podman/issues/22400)


## microsoft/retina <span style='color:#F1C40F'>(2.4K ⭐️)</span>

- [Repeated logs in plugins](https://github.com/microsoft/retina/issues/319)

- [Capture Manager should use contexts](https://github.com/microsoft/retina/issues/317)

- [log: fix "unsupported value" error in a log line](https://github.com/microsoft/retina/issues/309)

- [Feature Request: Outputting the .tar FileName as part of Retina Output please. ☕️](https://github.com/microsoft/retina/issues/286)

- [prevent the potential for misconfiguration of MetricsInterval](https://github.com/microsoft/retina/issues/284)

- [deps: bump DavidAnson/markdownlint-cli2-action from 9 to 16](https://github.com/microsoft/retina/pull/244)

- [Most or all metrics exported by Retina should be counters, not gauges](https://github.com/microsoft/retina/issues/237)

- [retina-operator shows up as unhealthy in Prometheus targets](https://github.com/microsoft/retina/issues/177)


## istio/istio <span style='color:#F1C40F'>(35.0K ⭐️)</span>

- [Ambient: should use `istio.io/dataplane-mode=ambient` for enrolling individual pods](https://github.com/istio/istio/issues/50355)

- [sidecar iptables - istio-specific rules should only be inserted into custom chains](https://github.com/istio/istio/issues/50532)

- [istioctl precheck: add unit tests](https://github.com/istio/istio/issues/50255)


## flyteorg/flyte <span style='color:#F1C40F'>(4.8K ⭐️)</span>

- [Add a securitySchemes to swagger.json](https://github.com/flyteorg/flyte/issues/5160)


## open-telemetry/opentelemetry-go <span style='color:#F1C40F'>(4.8K ⭐️)</span>

- [Add README generation to semconv tooling](https://github.com/open-telemetry/opentelemetry-go/issues/5255)


## kopia/kopia <span style='color:#F1C40F'>(6.3K ⭐️)</span>

- [Snapshot retention is deleting more snapshots as defined by policy in case system was pausing snapshots](https://github.com/kopia/kopia/issues/3758)


## devhatt/pet-dex-backend <span style='color:#F1C40F'>(36 ⭐️)</span>

- [Listagem de ONG](https://github.com/devhatt/pet-dex-backend/issues/110)

- [Consultar ONG](https://github.com/devhatt/pet-dex-backend/issues/111)

- [Criar referencias visuais para tarefas ](https://github.com/devhatt/pet-dex-backend/issues/116)

- [Update ONG](https://github.com/devhatt/pet-dex-backend/issues/112)

- [Criar Delete de ONG](https://github.com/devhatt/pet-dex-backend/issues/113)


## helm/helm <span style='color:#F1C40F'>(26.0K ⭐️)</span>

- [Additional YAML document header when rendering CRDs to stdout ](https://github.com/helm/helm/issues/12953)


## kubeflow/training-operator <span style='color:#F1C40F'>(1.4K ⭐️)</span>

- [Update pytorch launcher component in Kubeflow Pipelines repository](https://github.com/kubeflow/training-operator/issues/2068)

- [Add more AI/ML Training Examples](https://github.com/kubeflow/training-operator/issues/2040)


## osmosis-labs/osmosis <span style='color:#F1C40F'>(867 ⭐️)</span>

- [Delete osmomath.BigDec.ApproxRoot](https://github.com/osmosis-labs/osmosis/issues/8007)

- [Osmosis TokenFactory CLI is missing `MsgSetDenomMetadata`](https://github.com/osmosis-labs/osmosis/issues/7941)

- [[Bug]: TokenFactory CLI --help for `create-denom` is outdated as it includes `(cost osmo though!)`](https://github.com/osmosis-labs/osmosis/issues/7942)


## hashicorp/terraform-provider-aws <span style='color:#F1C40F'>(9.5K ⭐️)</span>

- [[Enhancement]: r/aws_redshift_cluster_snapshot: Set `arn` attribute from API response](https://github.com/hashicorp/terraform-provider-aws/issues/36915)

- [[Enhancement]: r/aws_quicksight_account_subscription: Add IAM Identity Center Instance ARN argument](https://github.com/hashicorp/terraform-provider-aws/issues/36782)


## cnoe-io/idpbuilder <span style='color:#F1C40F'>(117 ⭐️)</span>

- [[Doc] Rename the title of the section "Example commands" to "Advanced usage" as it concerns how to use the command "create"](https://github.com/cnoe-io/idpbuilder/issues/225)

- [[Bug]: Cannot git clone a gitea repository using ssh](https://github.com/cnoe-io/idpbuilder/issues/227)


## cloud-ark/kubeplus <span style='color:#F1C40F'>(607 ⭐️)</span>

- [Bulk upgrade test case](https://github.com/cloud-ark/kubeplus/issues/1230)

- [Follow Kubernetes naming convention for service instance names](https://github.com/cloud-ark/kubeplus/issues/1218)


## ethereum/go-ethereum <span style='color:#F1C40F'>(46.1K ⭐️)</span>

- [Beacon root contract in dev mode and private chains](https://github.com/ethereum/go-ethereum/issues/29539)


## solo-io/gloo <span style='color:#F1C40F'>(4.0K ⭐️)</span>

- [More detailed proxy latency docs](https://github.com/solo-io/gloo/issues/9333)


## kubernetes/minikube <span style='color:#F1C40F'>(28.4K ⭐️)</span>

- [`minikube mount` does not work on 6.1.0-20-cloud-amd64 - Detect when OS does not support 9p ](https://github.com/kubernetes/minikube/issues/18724)


## hashicorp/terraform <span style='color:#F1C40F'>(41.2K ⭐️)</span>

- [terraform test: produces invalid warning during cleanup phase](https://github.com/hashicorp/terraform/issues/35061)


## gnolang/gno <span style='color:#F1C40F'>(824 ⭐️)</span>

- [print all errors from `scanner.ErrorList` when calling `gno.ParseFile`](https://github.com/gnolang/gno/issues/1933)


## thanos-io/thanos <span style='color:#F1C40F'>(12.6K ⭐️)</span>

- [Ruler UI: `alert.query-template` is not honored inside the Rules UI](https://github.com/thanos-io/thanos/issues/7278)


## cilium/tetragon <span style='color:#F1C40F'>(3.3K ⭐️)</span>

- [doc: document why Tetragon will fail to load its BPF programs with operation not permitted](https://github.com/cilium/tetragon/issues/2265)


## external-secrets/external-secrets <span style='color:#F1C40F'>(3.9K ⭐️)</span>

- [Implement Pushsecrets `updatePolicy: IfNotExists` for AWS](https://github.com/external-secrets/external-secrets/issues/3380)

- [[Docs] GKE Workload Identity has changed](https://github.com/external-secrets/external-secrets/issues/3339)


## defang-io/defang <span style='color:#F1C40F'>(16 ⭐️)</span>

- [C# .NET sample](https://github.com/defang-io/defang/issues/294)


## grafana/mimir <span style='color:#F1C40F'>(3.7K ⭐️)</span>

- [mixin: networking overview renders useless panels for the gateway](https://github.com/grafana/mimir/issues/7911)


## elastic/beats <span style='color:#F1C40F'>(12.0K ⭐️)</span>

- [Elasticsearch output can report an incorrect count for active events](https://github.com/elastic/beats/issues/39146)


## glasskube/glasskube <span style='color:#F1C40F'>(845 ⭐️)</span>

- [Handle bootstrap errors gracefully when glasskube is not connected to the internet](https://github.com/glasskube/glasskube/issues/491)

- [Handle `v` prefix in versions in install and update command](https://github.com/glasskube/glasskube/issues/463)

- [CLI version completion when updating to specific version](https://github.com/glasskube/glasskube/issues/462)


## open-telemetry/opentelemetry-operator <span style='color:#F1C40F'>(1.1K ⭐️)</span>

- [Docs for enabling multi-instrumentation still point to enabling it via the --feature-gates flag](https://github.com/open-telemetry/opentelemetry-operator/issues/2855)


## cosmos/cosmos-sdk <span style='color:#F1C40F'>(5.9K ⭐️)</span>

- [Remove x/exp import](https://github.com/cosmos/cosmos-sdk/issues/19892)


## cebilon123/waffle <span style='color:#F1C40F'>(56 ⭐️)</span>

- [Add BPF provider ](https://github.com/cebilon123/waffle/issues/16)

- [Refactor how config and certificates are being read](https://github.com/cebilon123/waffle/issues/15)


## celestiaorg/celestia-app <span style='color:#F1C40F'>(324 ⭐️)</span>

- [Move markdown link check to run daily and Slack message failures](https://github.com/celestiaorg/celestia-app/issues/3324)

- [Use one term for MaxBlockSizeBytes](https://github.com/celestiaorg/celestia-app/issues/3322)


## CycloneDX/sbom-utility <span style='color:#F1C40F'>(62 ⭐️)</span>

- [Support SPDX in the "patch" command](https://github.com/CycloneDX/sbom-utility/issues/82)

- [TODO: Change Formulation and ModelCard schemas to use pointers](https://github.com/CycloneDX/sbom-utility/issues/79)


## kedacore/keda <span style='color:#F1C40F'>(7.8K ⭐️)</span>

- [Metrics Server: `klogr.New` is deprecated use `textlogger.NewLogger` instead](https://github.com/kedacore/keda/issues/5732)


## opentofu/opentofu <span style='color:#F1C40F'>(20.2K ⭐️)</span>

- [Broken links on website](https://github.com/opentofu/opentofu/issues/1563)

- [Migration guide index page doesn't have a title](https://github.com/opentofu/opentofu/issues/1565)

