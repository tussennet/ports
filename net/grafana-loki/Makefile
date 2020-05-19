# $FreeBSD$

PORTNAME=	grafana-loki
DISTVERSIONPREFIX=v
DISTVERSION=	1.4.1
CATEGORIES=	net

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Like Prometheus, but for logs

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	grafana
GH_PROJECT=	loki

PLIST_FILES=	bin/loki \
		bin/logcli \
		bin/promtail

GO_TARGET=	./cmd/loki ./cmd/logcli ./cmd/promtail

GH_TUPLE=	\
		Azure:azure-pipeline-go:v0.2.2:azure_azure_pipeline_go/vendor/github.com/Azure/azure-pipeline-go \
		Azure:azure-sdk-for-go:v36.2.0:azure_azure_sdk_for_go/vendor/github.com/Azure/azure-sdk-for-go \
		Azure:azure-storage-blob-go:v0.8.0:azure_azure_storage_blob_go/vendor/github.com/Azure/azure-storage-blob-go \
		Azure:go-ansiterm:d6e3b3328b78:azure_go_ansiterm/vendor/github.com/Azure/go-ansiterm \
		Azure:go-autorest:3492b2aff503:azure_go_autorest_autorest/vendor/github.com/Azure/go-autorest \
		Azure:go-autorest:5bd9621f41a0:azure_go_autorest_date \
		Azure:go-autorest:logger/v0.1.0:azure_go_autorest_logger \
		Azure:go-autorest:tracing/v0.5.0:azure_go_autorest_tracing \
		BurntSushi:toml:v0.3.1:burntsushi_toml/vendor/github.com/BurntSushi/toml \
		Microsoft:go-winio:v0.4.12:microsoft_go_winio/vendor/github.com/Microsoft/go-winio \
		NYTimes:gziphandler:v1.1.1:nytimes_gziphandler/vendor/github.com/NYTimes/gziphandler \
		alecthomas:kingpin:v2.2.6:alecthomas_kingpin/vendor/gopkg.in/alecthomas/kingpin.v2 \
		alecthomas:template:fb15b899a751:alecthomas_template/vendor/github.com/alecthomas/template \
		alecthomas:units:f65c72e2690d:alecthomas_units/vendor/github.com/alecthomas/units \
		armon:go-metrics:v0.3.0:armon_go_metrics/vendor/github.com/armon/go-metrics \
		aws:aws-sdk-go:v1.25.48:aws_aws_sdk_go/vendor/github.com/aws/aws-sdk-go \
		beorn7:perks:v1.0.1:beorn7_perks/vendor/github.com/beorn7/perks \
		blang:semver:v3.5.1:blang_semver/vendor/github.com/blang/semver \
		bmatcuk:doublestar:v1.2.2:bmatcuk_doublestar/vendor/github.com/bmatcuk/doublestar \
		bradfitz:gomemcache:a41fca850d0b:bradfitz_gomemcache/vendor/github.com/bradfitz/gomemcache \
		c2h5oh:datasize:28bbd4740fee:c2h5oh_datasize/vendor/github.com/c2h5oh/datasize \
		census-instrumentation:opencensus-go:v0.22.2:census_instrumentation_opencensus_go/vendor/go.opencensus.io \
		cespare:xxhash:v1.1.0:cespare_xxhash/vendor/github.com/cespare/xxhash \
		cespare:xxhash:v2.1.1:cespare_xxhash_v2/vendor/github.com/cespare/xxhash/v2 \
		containerd:containerd:v1.3.2:containerd_containerd/vendor/github.com/containerd/containerd \
		containerd:fifo:a9fb20d87448:containerd_fifo/vendor/github.com/containerd/fifo \
		coreos:go-semver:v0.3.0:coreos_go_semver/vendor/github.com/coreos/go-semver \
		coreos:go-systemd:95778dfbb74e:coreos_go_systemd/vendor/github.com/coreos/go-systemd \
		coreos:pkg:399ea9e2e55f:coreos_pkg/vendor/github.com/coreos/pkg \
		cortexproject:cortex:v0.7.0-rc.0:cortexproject_cortex/vendor/github.com/cortexproject/cortex \
		davecgh:go-spew:v1.1.1:davecgh_go_spew/vendor/github.com/davecgh/go-spew \
		dgrijalva:jwt-go:v3.2.0:dgrijalva_jwt_go/vendor/github.com/dgrijalva/jwt-go \
		docker:distribution:v2.7.1:docker_distribution/vendor/github.com/docker/distribution \
		moby:moby:4760db040282:docker_docker/vendor/github.com/docker/docker \
		docker:go-connections:v0.4.0:docker_go_connections/vendor/github.com/docker/go-connections \
		docker:go-metrics:b84716841b82:docker_go_metrics/vendor/github.com/docker/go-metrics \
		docker:go-plugins-helpers:1e6269c305b8:docker_go_plugins_helpers/vendor/github.com/docker/go-plugins-helpers \
		docker:go-units:v0.4.0:docker_go_units/vendor/github.com/docker/go-units \
		dominikh:go-tools:v0.0.1-2019.2.3:dominikh_go_tools/vendor/honnef.co/go/tools \
		dustin:go-humanize:v1.0.0:dustin_go_humanize/vendor/github.com/dustin/go-humanize \
		edsrzf:mmap-go:v1.0.0:edsrzf_mmap_go/vendor/github.com/edsrzf/mmap-go \
		etcd-io:bbolt:v1.3.3:etcd_io_bbolt/vendor/go.etcd.io/bbolt \
		etcd-io:etcd:8f85f0dc2607:etcd_io_etcd/vendor/go.etcd.io/etcd \
		facette:natsort:2cd4dd1e2dcb:facette_natsort/vendor/github.com/facette/natsort \
		fatih:color:v1.7.0:fatih_color/vendor/github.com/fatih/color \
		fluent:fluent-bit-go:ea13c021720c:fluent_fluent_bit_go/vendor/github.com/fluent/fluent-bit-go \
		fsnotify:fsnotify:v1.4.7:fsnotify_fsnotify/vendor/github.com/fsnotify/fsnotify \
		fsouza:fake-gcs-server:v1.7.0:fsouza_fake_gcs_server/vendor/github.com/fsouza/fake-gcs-server \
		go-inf:inf:v0.9.1:go_inf_inf/vendor/gopkg.in/inf.v0 \
		go-kit:kit:v0.9.0:go_kit_kit/vendor/github.com/go-kit/kit \
		go-logfmt:logfmt:v0.4.0:go_logfmt_logfmt/vendor/github.com/go-logfmt/logfmt \
		go-tomb:tomb:dd632973f1e7:go_tomb_tomb/vendor/gopkg.in/tomb.v1 \
		go-yaml:yaml:v2.2.7:go_yaml_yaml/vendor/gopkg.in/yaml.v2 \
		gocql:gocql:95d072f1b5bb:gocql_gocql/vendor/github.com/gocql/gocql \
		gogo:googleapis:v1.1.0:gogo_googleapis/vendor/github.com/gogo/googleapis \
		gogo:protobuf:v1.3.1:gogo_protobuf/vendor/github.com/gogo/protobuf \
		gogo:status:v1.0.3:gogo_status/vendor/github.com/gogo/status \
		golang:appengine:v1.6.5:golang_appengine/vendor/google.golang.org/appengine \
		golang:crypto:e1110fd1c708:golang_crypto/vendor/golang.org/x/crypto \
		golang:exp:a1ab85dbe136:golang_exp/vendor/golang.org/x/exp \
		golang:groupcache:611e8accdfc9:golang_groupcache/vendor/github.com/golang/groupcache \
		golang:lint:16217165b5de:golang_lint/vendor/golang.org/x/lint \
		golang:net:aa69164e4478:golang_net/vendor/golang.org/x/net \
		golang:oauth2:0f29369cfe45:golang_oauth2/vendor/golang.org/x/oauth2 \
		golang:protobuf:v1.3.2:golang_protobuf/vendor/github.com/golang/protobuf \
		golang:snappy:v0.0.1:golang_snappy/vendor/github.com/golang/snappy \
		golang:sync:cd5d95a43a6e:golang_sync/vendor/golang.org/x/sync \
		golang:sys:e047566fdf82:golang_sys/vendor/golang.org/x/sys \
		golang:text:v0.3.2:golang_text/vendor/golang.org/x/text \
		golang:time:555d28b269f0:golang_time/vendor/golang.org/x/time \
		golang:tools:3a792d9c32b2:golang_tools/vendor/golang.org/x/tools \
		gomodule:redigo:v2.0.0:gomodule_redigo/vendor/github.com/gomodule/redigo \
		google:btree:v1.0.0:google_btree/vendor/github.com/google/btree \
		google:go-cmp:v0.4.0:google_go_cmp/vendor/github.com/google/go-cmp \
		google:go-genproto:c23dd37a84c9:google_go_genproto/vendor/google.golang.org/genproto \
		google:gofuzz:v1.0.0:google_gofuzz/vendor/github.com/google/gofuzz \
		google:uuid:v1.1.1:google_uuid/vendor/github.com/google/uuid \
		googleapis:gax-go:v2.0.5:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go \
		googleapis:gnostic:v0.3.1:googleapis_gnostic/vendor/github.com/googleapis/gnostic \
		googleapis:google-api-go-client:v0.14.0:googleapis_google_api_go_client/vendor/google.golang.org/api \
		googleapis:google-cloud-go:bigtable/v1.1.0:googleapis_google_cloud_go_bigtable \
		googleapis:google-cloud-go:storage/v1.3.0:googleapis_google_cloud_go_storage \
		googleapis:google-cloud-go:v0.49.0:googleapis_google_cloud_go/vendor/cloud.google.com/go \
		gophercloud:gophercloud:v0.6.0:gophercloud_gophercloud/vendor/github.com/gophercloud/gophercloud \
		gorilla:mux:v1.7.1:gorilla_mux/vendor/github.com/gorilla/mux \
		gorilla:websocket:v1.4.0:gorilla_websocket/vendor/github.com/gorilla/websocket \
		grafana:tail:0b54ddf21fe7:grafana_tail/vendor/github.com/hpcloud/tail \
		grpc-ecosystem:go-grpc-middleware:v1.1.0:grpc_ecosystem_go_grpc_middleware/vendor/github.com/grpc-ecosystem/go-grpc-middleware \
		grpc-ecosystem:go-grpc-prometheus:6af20e3a5340:grpc_ecosystem_go_grpc_prometheus/vendor/github.com/grpc-ecosystem/go-grpc-prometheus \
		grpc-ecosystem:grpc-gateway:v1.12.1:grpc_ecosystem_grpc_gateway/vendor/github.com/grpc-ecosystem/grpc-gateway \
		grpc-ecosystem:grpc-opentracing:8e809c8a8645:grpc_ecosystem_grpc_opentracing/vendor/github.com/grpc-ecosystem/grpc-opentracing \
		grpc:grpc-go:v1.25.1:grpc_grpc_go/vendor/google.golang.org/grpc \
		hailocab:go-hostpool:e80d13ce29ed:hailocab_go_hostpool/vendor/github.com/hailocab/go-hostpool \
		hashicorp:consul:v1.3.0:hashicorp_consul_api/vendor/github.com/hashicorp/consul \
		hashicorp:errwrap:v1.0.0:hashicorp_errwrap/vendor/github.com/hashicorp/errwrap \
		hashicorp:go-cleanhttp:v0.5.1:hashicorp_go_cleanhttp/vendor/github.com/hashicorp/go-cleanhttp \
		hashicorp:go-immutable-radix:v1.1.0:hashicorp_go_immutable_radix/vendor/github.com/hashicorp/go-immutable-radix \
		hashicorp:go-msgpack:v0.5.5:hashicorp_go_msgpack/vendor/github.com/hashicorp/go-msgpack \
		hashicorp:go-multierror:v1.0.0:hashicorp_go_multierror/vendor/github.com/hashicorp/go-multierror \
		hashicorp:go-rootcerts:v1.0.1:hashicorp_go_rootcerts/vendor/github.com/hashicorp/go-rootcerts \
		hashicorp:go-sockaddr:v1.0.2:hashicorp_go_sockaddr/vendor/github.com/hashicorp/go-sockaddr \
		hashicorp:golang-lru:v0.5.3:hashicorp_golang_lru/vendor/github.com/hashicorp/golang-lru \
		hashicorp:memberlist:v0.1.5:hashicorp_memberlist/vendor/github.com/hashicorp/memberlist \
		hashicorp:serf:v0.8.5:hashicorp_serf/vendor/github.com/hashicorp/serf \
		influxdata:go-syslog:v2.0.1:influxdata_go_syslog_v2/vendor/github.com/influxdata/go-syslog/v2 \
		jmespath:go-jmespath:c2b33e8439af:jmespath_go_jmespath/vendor/github.com/jmespath/go-jmespath \
		jonboulle:clockwork:v0.1.0:jonboulle_clockwork/vendor/github.com/jonboulle/clockwork \
		jpillora:backoff:v1.0.0:jpillora_backoff/vendor/github.com/jpillora/backoff \
		json-iterator:go:v1.1.9:json_iterator_go/vendor/github.com/json-iterator/go \
		jstemmer:go-junit-report:v0.9.1:jstemmer_go_junit_report/vendor/github.com/jstemmer/go-junit-report \
		klauspost:compress:v1.9.4:klauspost_compress/vendor/github.com/klauspost/compress \
		konsorten:go-windows-terminal-sequences:v1.0.2:konsorten_go_windows_terminal_sequences/vendor/github.com/konsorten/go-windows-terminal-sequences \
		kr:logfmt:b84e30acd515:kr_logfmt/vendor/github.com/kr/logfmt \
		kubernetes-sigs:yaml:v1.1.0:kubernetes_sigs_yaml/vendor/sigs.k8s.io/yaml \
		kubernetes:api:47f6de673b26:kubernetes_api/vendor/k8s.io/api \
		kubernetes:apimachinery:3c7067801da2:kubernetes_apimachinery/vendor/k8s.io/apimachinery \
		kubernetes:client-go:78d2af792bab:kubernetes_client_go/vendor/k8s.io/client-go \
		kubernetes:klog:v1.0.0:kubernetes_klog/vendor/k8s.io/klog \
		kubernetes:utils:6ca3b61696b6:kubernetes_utils/vendor/k8s.io/utils \
		leodido:ragel-machinery:299bdde78165:leodido_ragel_machinery/vendor/github.com/leodido/ragel-machinery \
		mattn:go-colorable:v0.0.9:mattn_go_colorable/vendor/github.com/mattn/go-colorable \
		mattn:go-ieproxy:7c0f6868bffe:mattn_go_ieproxy/vendor/github.com/mattn/go-ieproxy \
		mattn:go-isatty:v0.0.4:mattn_go_isatty/vendor/github.com/mattn/go-isatty \
		matttproud:golang_protobuf_extensions:v1.0.1:matttproud_golang_protobuf_extensions/vendor/github.com/matttproud/golang_protobuf_extensions \
		miekg:dns:v1.1.22:miekg_dns/vendor/github.com/miekg/dns \
		mitchellh:go-homedir:v1.1.0:mitchellh_go_homedir/vendor/github.com/mitchellh/go-homedir \
		mitchellh:mapstructure:v1.1.2:mitchellh_mapstructure/vendor/github.com/mitchellh/mapstructure \
		modern-go:concurrent:bacd9c7ef1dd:modern_go_concurrent/vendor/github.com/modern-go/concurrent \
		modern-go:reflect2:v1.0.1:modern_go_reflect2/vendor/github.com/modern-go/reflect2 \
		morikuni:aec:v1.0.0:morikuni_aec/vendor/github.com/morikuni/aec \
		mwitkow:go-conntrack:2f068394615f:mwitkow_go_conntrack/vendor/github.com/mwitkow/go-conntrack \
		oklog:ulid:v1.3.1:oklog_ulid/vendor/github.com/oklog/ulid \
		opencontainers:go-digest:v1.0.0-rc1:opencontainers_go_digest/vendor/github.com/opencontainers/go-digest \
		opencontainers:image-spec:v1.0.1:opencontainers_image_spec/vendor/github.com/opencontainers/image-spec \
		opentracing-contrib:go-grpc:4b5a12d3ff02:opentracing_contrib_go_grpc/vendor/github.com/opentracing-contrib/go-grpc \
		opentracing-contrib:go-stdlib:cf7a6c988dc9:opentracing_contrib_go_stdlib/vendor/github.com/opentracing-contrib/go-stdlib \
		opentracing:opentracing-go:2876d2018785:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go \
		pierrec:lz4:9085dacd1e1e:pierrec_lz4/vendor/github.com/pierrec/lz4 \
		pkg:errors:v0.8.1:pkg_errors/vendor/github.com/pkg/errors \
		pmezard:go-difflib:v1.0.0:pmezard_go_difflib/vendor/github.com/pmezard/go-difflib \
		prometheus:client_golang:v1.5.0:prometheus_client_golang/vendor/github.com/prometheus/client_golang \
		prometheus:client_model:v0.2.0:prometheus_client_model/vendor/github.com/prometheus/client_model \
		prometheus:common:v0.9.1:prometheus_common/vendor/github.com/prometheus/common \
		prometheus:procfs:v0.0.8:prometheus_procfs/vendor/github.com/prometheus/procfs \
		prometheus:prometheus:b90be6f32a33:prometheus_prometheus/vendor/github.com/prometheus/prometheus \
		rsc:binaryregexp:v0.2.0:rsc_binaryregexp/vendor/rsc.io/binaryregexp \
		samuel:go-zookeeper:2cc03de413da:samuel_go_zookeeper/vendor/github.com/samuel/go-zookeeper \
		sean-:seed:e2103e2c3529:sean__seed/vendor/github.com/sean-/seed \
		sercand:kuberesolver:v2.1.0:sercand_kuberesolver/vendor/github.com/sercand/kuberesolver \
		shurcooL:httpfs:8d4bc4ba7749:shurcool_httpfs/vendor/github.com/shurcooL/httpfs \
		shurcooL:vfsgen:6a9ea43bcacd:shurcool_vfsgen/vendor/github.com/shurcooL/vfsgen \
		sirupsen:logrus:v1.4.2:sirupsen_logrus/vendor/github.com/sirupsen/logrus \
		soheilhy:cmux:v0.1.4:soheilhy_cmux/vendor/github.com/soheilhy/cmux \
		spf13:pflag:v1.0.5:spf13_pflag/vendor/github.com/spf13/pflag \
		stretchr:objx:v0.2.0:stretchr_objx/vendor/github.com/stretchr/objx \
		stretchr:testify:v1.5.1:stretchr_testify/vendor/github.com/stretchr/testify \
		tmc:grpc-websocket-proxy:0ad062ec5ee5:tmc_grpc_websocket_proxy/vendor/github.com/tmc/grpc-websocket-proxy \
		tonistiigi:fifo:a9fb20d87448:tonistiigi_fifo \
		uber-go:atomic:v1.5.0:uber_go_atomic/vendor/go.uber.org/atomic \
		uber-go:multierr:v1.1.0:uber_go_multierr/vendor/go.uber.org/multierr \
		uber-go:zap:v1.10.0:uber_go_zap/vendor/go.uber.org/zap \
		uber:jaeger-client-go:v2.20.1:uber_jaeger_client_go/vendor/github.com/uber/jaeger-client-go \
		uber:jaeger-lib:v2.2.0:uber_jaeger_lib/vendor/github.com/uber/jaeger-lib \
		ugorji:go:v1.1.7:ugorji_go_codec/vendor/github.com/ugorji/go \
		weaveworks:common:2708ba4e60a4:weaveworks_common/vendor/github.com/weaveworks/common \
		weaveworks:promrus:v1.2.0:weaveworks_promrus/vendor/github.com/weaveworks/promrus \
		xiang90:probing:43a291ad63a2:xiang90_probing/vendor/github.com/xiang90/probing

post-extract:
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/bigtable
	@${RLN} ${WRKSRC_googleapis_google_cloud_go_bigtable}/bigtable ${WRKSRC}/vendor/cloud.google.com/go/bigtable
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/storage
	@${RLN} ${WRKSRC_googleapis_google_cloud_go_storage}/storage ${WRKSRC}/vendor/cloud.google.com/go/storage
	@${RM} -r ${WRKSRC}/vendor/github.com/Azure/go-autorest/autorest/date
	@${RLN} ${WRKSRC_azure_go_autorest_date}/autorest/date ${WRKSRC}/vendor/github.com/Azure/go-autorest/autorest/date
	@${RM} -r ${WRKSRC}/vendor/github.com/Azure/go-autorest/logger
	@${RLN} ${WRKSRC_azure_go_autorest_logger}/logger ${WRKSRC}/vendor/github.com/Azure/go-autorest/logger
	@${RM} -r ${WRKSRC}/vendor/github.com/Azure/go-autorest/tracing
	@${RLN} ${WRKSRC_azure_go_autorest_tracing}/tracing ${WRKSRC}/vendor/github.com/Azure/go-autorest/tracing
	@${MKDIR} ${WRKSRC}/vendor/github.com/tonistiigi
	@${RLN} ${WRKSRC_containerd_fifo} ${WRKSRC}/vendor/github.com/tonistiigi/fifo
	@${MKDIR} ${WRKSRC}/vendor/gopkg.in
	@${RLN} ${WRKSRC_fsnotify_fsnotify} ${WRKSRC}/vendor/gopkg.in/fsnotify.v1
	@${MKDIR} ${WRKSRC}/vendor/gopkg.in/fsnotify
	@${RLN} ${WRKSRC_fsnotify_fsnotify} ${WRKSRC}/vendor/gopkg.in/fsnotify/fsnotify.v1

.include <bsd.port.mk>
