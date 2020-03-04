# $FreeBSD: head/sysutils/sensu-go/Makefile 509241 2019-08-18 21:48:11Z zi $

PORTNAME=	sensu-go
DISTVERSIONPREFIX=v
DISTVERSION=	5.18.0
CATEGORIES=	sysutils

MAINTAINER=	jp+ports@supplntr.io
COMMENT=	Ephemeral infrastructure and distributed applications monitoring

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
USE_GITHUB=	yes
GH_ACCOUNT=	sensu

GH_TUPLE=	\
		AlecAivazis:survey:v1.4.0:alecaivazis_survey/vendor/gopkg.in/AlecAivazis/survey.v1 \
		AlecAivazis:survey:v1.4.1:alecaivazis_survey_1/vendor/github.com/AlecAivazis/survey \
		Azure:go-ansiterm:d6e3b3328b78:azure_go_ansiterm/vendor/github.com/Azure/go-ansiterm \
		NYTimes:gziphandler:5032c8878b9d:nytimes_gziphandler/vendor/github.com/NYTimes/gziphandler \
		StackExchange:wmi:b12b22c5341f:stackexchange_wmi/vendor/github.com/StackExchange/wmi \
		andybalholm:brotli:v1.0.0:andybalholm_brotli/vendor/github.com/andybalholm/brotli \
		ash2k:stager:6e9c7b0eacd4:ash2k_stager/vendor/github.com/ash2k/stager \
		atlassian:gostatsd:af796620006e:atlassian_gostatsd/vendor/github.com/atlassian/gostatsd \
		beorn7:perks:v1.0.1:beorn7_perks/vendor/github.com/beorn7/perks \
		cespare:xxhash:v2.1.0:cespare_xxhash/vendor/github.com/cespare/xxhash/v2 \
		coreos:bbolt:v1.3.3:coreos_bbolt/vendor/github.com/coreos/bbolt \
		coreos:etcd:v3.3.17:coreos_etcd/vendor/github.com/coreos/etcd \
		coreos:go-semver:v0.3.0:coreos_go_semver/vendor/github.com/coreos/go-semver \
		coreos:go-systemd:fd7a80b32e1f:coreos_go_systemd/vendor/github.com/coreos/go-systemd \
		coreos:pkg:399ea9e2e55f:coreos_pkg/vendor/github.com/coreos/pkg \
		dave:jennifer:d8bdbdbee4e1:dave_jennifer/vendor/github.com/dave/jennifer \
		davecgh:go-spew:v1.1.1:davecgh_go_spew/vendor/github.com/davecgh/go-spew \
		dgrijalva:jwt-go:v3.2.0:dgrijalva_jwt_go/vendor/github.com/dgrijalva/jwt-go \
		moby:moby:cbde00b44273:moby_moby/vendor/github.com/docker/docker \
		dsnet:compress:v0.0.1:dsnet_compress/vendor/github.com/dsnet/compress \
		echlebek:crock:v1.0.1:echlebek_crock/vendor/github.com/echlebek/crock \
		echlebek:timeproxy:v1.0.0:echlebek_timeproxy/vendor/github.com/echlebek/timeproxy \
		emicklei:proto:v1.1.0:emicklei_proto/vendor/github.com/emicklei/proto \
		etcd-io:bbolt:v1.3.2:etcd_io_bbolt/vendor/go.etcd.io/bbolt \
		fsnotify:fsnotify:v1.4.7:fsnotify_fsnotify/vendor/github.com/fsnotify/fsnotify \
		ghodss:yaml:v1.0.0:ghodss_yaml/vendor/github.com/ghodss/yaml \
		go-ole:go-ole:de8695c8edbf:go_ole_go_ole/vendor/github.com/go-ole/go-ole \
		go-resty:resty:v2.1.0:go_resty_resty/vendor/github.com/go-resty/resty/v2 \
		go-sourcemap:sourcemap:v1.0.5:go_sourcemap_sourcemap/vendor/gopkg.in/sourcemap.v1 \
		go-yaml:yaml:v2.2.4:go_yaml_yaml/vendor/gopkg.in/yaml.v2 \
		gogo:protobuf:v1.3.1:gogo_protobuf/vendor/github.com/gogo/protobuf \
		golang:crypto:87dc89f01550:golang_crypto/vendor/golang.org/x/crypto \
		golang:groupcache:404acd9df4cc:golang_groupcache/vendor/github.com/golang/groupcache \
		golang:net:da9a3fd4c582:golang_net/vendor/golang.org/x/net \
		golang:protobuf:v1.3.2:golang_protobuf/vendor/github.com/golang/protobuf \
		golang:snappy:v0.0.1:golang_snappy/vendor/github.com/golang/snappy \
		golang:sys:4c7a9d0fe056:golang_sys/vendor/golang.org/x/sys \
		golang:text:v0.3.2:golang_text/vendor/golang.org/x/text \
		golang:time:c4c64cad1fd0:golang_time/vendor/golang.org/x/time \
		google:btree:v1.0.0:google_btree/vendor/github.com/google/btree \
		google:go-genproto:548a555dbc03:google_go_genproto/vendor/google.golang.org/genproto \
		google:uuid:v1.1.1:google_uuid/vendor/github.com/google/uuid \
		gorilla:context:1ea25387ff6f:gorilla_context/vendor/github.com/gorilla/context \
		gorilla:mux:v1.6.2:gorilla_mux/vendor/github.com/gorilla/mux \
		gorilla:websocket:v1.4.1:gorilla_websocket/vendor/github.com/gorilla/websocket \
		graph-gophers:dataloader:78139374585c:graph_gophers_dataloader/vendor/github.com/graph-gophers/dataloader \
		graphql-go:graphql:2e2b648ecbe4:graphql_go_graphql/vendor/github.com/graphql-go/graphql \
		grpc-ecosystem:go-grpc-middleware:v1.1.0:grpc_ecosystem_go_grpc_middleware/vendor/github.com/grpc-ecosystem/go-grpc-middleware \
		grpc-ecosystem:go-grpc-prometheus:v1.2.0:grpc_ecosystem_go_grpc_prometheus/vendor/github.com/grpc-ecosystem/go-grpc-prometheus \
		grpc-ecosystem:grpc-gateway:v1.11.3:grpc_ecosystem_grpc_gateway/vendor/github.com/grpc-ecosystem/grpc-gateway \
		grpc:grpc-go:v1.24.0:grpc_grpc_go/vendor/google.golang.org/grpc \
		gxed:GoEndian:0f5c6873267e:gxed_goendian/vendor/github.com/gxed/GoEndian \
		gxed:eventfd:80a92cca79a8:gxed_eventfd/vendor/github.com/gxed/eventfd \
		h2non:filetype:v1.0.3:h2non_filetype/vendor/gopkg.in/h2non/filetype.v1 \
		hashicorp:go-version:v1.2.0:hashicorp_go_version/vendor/github.com/hashicorp/go-version \
		hashicorp:hcl:v1.0.0:hashicorp_hcl/vendor/github.com/hashicorp/hcl \
		inconshreveable:mousetrap:v1.0.0:inconshreveable_mousetrap/vendor/github.com/inconshreveable/mousetrap \
		ipfs:go-log:7ecd3df29a4a:ipfs_go_log/vendor/github.com/ipfs/go-log \
		jbenet:go-reuseport:15a1cd37f050:jbenet_go_reuseport/vendor/github.com/jbenet/go-reuseport \
		jonboulle:clockwork:v0.1.0:jonboulle_clockwork/vendor/github.com/jonboulle/clockwork \
		json-iterator:go:v1.1.7:json_iterator_go/vendor/github.com/json-iterator/go \
		klauspost:compress:v1.9.2:klauspost_compress/vendor/github.com/klauspost/compress \
		klauspost:pgzip:v1.2.1:klauspost_pgzip/vendor/github.com/klauspost/pgzip \
		konsorten:go-windows-terminal-sequences:v1.0.2:konsorten_go_windows_terminal_sequences/vendor/github.com/konsorten/go-windows-terminal-sequences \
		kubernetes-sigs:yaml:v1.1.0:kubernetes_sigs_yaml/vendor/sigs.k8s.io/yaml \
		libp2p:go-reuseport:15a1cd37f050:libp2p_go_reuseport/vendor/github.com/libp2p/go-reuseport \
		libp2p:go-sockaddr:f3e9f73a53d1:libp2p_go_sockaddr/vendor/github.com/libp2p/go-sockaddr \
		magiconair:properties:v1.8.0:magiconair_properties/vendor/github.com/magiconair/properties \
		mattn:go-colorable:v0.0.9:mattn_go_colorable/vendor/github.com/mattn/go-colorable \
		mattn:go-isatty:v0.0.2:mattn_go_isatty/vendor/github.com/mattn/go-isatty \
		mattn:go-runewidth:v0.0.2:mattn_go_runewidth/vendor/github.com/mattn/go-runewidth \
		matttproud:golang_protobuf_extensions:v1.0.1:matttproud_golang_protobuf_extensions/vendor/github.com/matttproud/golang_protobuf_extensions \
		mgutz:ansi:9520e82c474b:mgutz_ansi/vendor/github.com/mgutz/ansi \
		mholt:archiver:44285f7ed244:mholt_archiver/vendor/github.com/mholt/archiver/v3 \
		mitchellh:go-homedir:v1.1.0:mitchellh_go_homedir/vendor/github.com/mitchellh/go-homedir \
		mitchellh:mapstructure:v1.1.2:mitchellh_mapstructure/vendor/github.com/mitchellh/mapstructure \
		modern-go:concurrent:bacd9c7ef1dd:modern_go_concurrent/vendor/github.com/modern-go/concurrent \
		modern-go:reflect2:v1.0.1:modern_go_reflect2/vendor/github.com/modern-go/reflect2 \
		nwaples:rardecode:v1.0.0:nwaples_rardecode/vendor/github.com/nwaples/rardecode \
		olekukonko:tablewriter:d4647c9c7a84:olekukonko_tablewriter/vendor/github.com/olekukonko/tablewriter \
		opentracing:opentracing-go:v1.1.0:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go \
		pelletier:go-toml:v1.2.0:pelletier_go_toml/vendor/github.com/pelletier/go-toml \
		pierrec:lz4:v3.0.1:pierrec_lz4/vendor/github.com/pierrec/lz4/v3 \
		pmezard:go-difflib:v1.0.0:pmezard_go_difflib/vendor/github.com/pmezard/go-difflib \
		prometheus:client_golang:v1.2.0:prometheus_client_golang/vendor/github.com/prometheus/client_golang \
		prometheus:client_model:14fe0d1b01d4:prometheus_client_model/vendor/github.com/prometheus/client_model \
		prometheus:common:v0.7.0:prometheus_common/vendor/github.com/prometheus/common \
		prometheus:procfs:v0.0.5:prometheus_procfs/vendor/github.com/prometheus/procfs \
		robertkrimen:otto:15f95af6e78d:robertkrimen_otto/vendor/github.com/robertkrimen/otto \
		robfig:cron:v3.0.0:robfig_cron/vendor/github.com/robfig/cron/v3 \
		sensu:lasr:v1.2.1:sensu_lasr/vendor/github.com/sensu/lasr \
		shirou:gopsutil:8048a2e9c577:shirou_gopsutil/vendor/github.com/shirou/gopsutil \
		shirou:w32:bb4de0191aa4:shirou_w32/vendor/github.com/shirou/w32 \
		sirupsen:logrus:v1.4.2:sirupsen_logrus/vendor/github.com/sirupsen/logrus \
		soheilhy:cmux:v0.1.4:soheilhy_cmux/vendor/github.com/soheilhy/cmux \
		spf13:afero:v1.1.2:spf13_afero/vendor/github.com/spf13/afero \
		spf13:cast:v1.3.0:spf13_cast/vendor/github.com/spf13/cast \
		spf13:cobra:v0.0.5:spf13_cobra/vendor/github.com/spf13/cobra \
		spf13:jwalterweatherman:v1.0.0:spf13_jwalterweatherman/vendor/github.com/spf13/jwalterweatherman \
		spf13:pflag:v1.0.5:spf13_pflag/vendor/github.com/spf13/pflag \
		spf13:viper:v1.4.0:spf13_viper/vendor/github.com/spf13/viper \
		stretchr:objx:v0.1.1:stretchr_objx/vendor/github.com/stretchr/objx \
		stretchr:testify:v1.4.0:stretchr_testify/vendor/github.com/stretchr/testify \
		tmc:grpc-websocket-proxy:0ad062ec5ee5:tmc_grpc_websocket_proxy/vendor/github.com/tmc/grpc-websocket-proxy \
		uber-go:atomic:v1.4.0:uber_go_atomic/vendor/go.uber.org/atomic \
		uber-go:multierr:v1.2.0:uber_go_multierr/vendor/go.uber.org/multierr \
		uber-go:zap:v1.10.0:uber_go_zap/vendor/go.uber.org/zap \
		ulikunitz:xz:v0.5.6:ulikunitz_xz/vendor/github.com/ulikunitz/xz \
		whyrusleeping:go-logging:0457bb6b88fc:whyrusleeping_go_logging/vendor/github.com/whyrusleeping/go-logging \
		willf:pad:b3d780601022:willf_pad/vendor/github.com/willf/pad \
		xi2:xz:48954b6210f8:xi2_xz/vendor/github.com/xi2/xz \
		xiang90:probing:43a291ad63a2:xiang90_probing/vendor/github.com/xiang90/probing

GO_TARGET=

OPTIONS_DEFINE=		SENSU_AGENT SENSU_CLI SENSU_BE
OPTIONS_DEFAULT=	SENSU_AGENT
OPTIONS_SUB=		yes

SENSU_AGENT_DESC=	Install sensu-agent daemon
SENSU_AGENT_VARS=	GO_TARGET+=./cmd/sensu-agent \
			USE_RC_SUBR+=sensu-agent

SENSU_CLI_DESC=		Install sensuctl binary
SENSU_CLI_VARS=		GO_TARGET+=./cmd/sensuctl

SENSU_BE_DESC=		Install sensu-backend daemon (unsupported)
SENSU_BE_VARS=		GO_TARGET+=./cmd/sensu-backend
SENSU_BE_BROKEN=	Currently not tested under FreeBSD

.include <bsd.port.options.mk>

post-patch:
	@${REINPLACE_CMD} '/git rev/d' ${WRKSRC}/build.sh
#	@${REINPLACE_CMD} '/exec.Command("git"/d' ${WRKSRC}/version/cmd/version/version.go

#do-build:
#.for GO_TARGET in ${GO_TARGETS}
#	@(cd ${GO_WRKSRC}; ${SETENV} ${MAKE_ENV} ${GO_ENV} ./build.sh build_${GO_TARGET})
#.endfor

#do-install:
#.for GO_BIN_TARGET in ${GO_BIN_TARGETS}
#	${INSTALL_PROGRAM} ${GO_WRKSRC}/target/freebsd-${ARCH:S/^i//}/${GO_BIN_TARGET} ${STAGEDIR}${PREFIX}/bin
#.endfor

post-install-SENSU_AGENT-on:
	${INSTALL_DATA} ${PATCHDIR}/sensu-agent.yml.sample \
		${STAGEDIR}${PREFIX}/etc/sensu-agent.yml.sample

post-install-SENSU_BE-on:
	${INSTALL_DATA} ${PATCHDIR}/sensu-backend.yml.sample \
		${STAGEDIR}${PREFIX}/etc/sensu-backend.yml.sample

.include <bsd.port.mk>
