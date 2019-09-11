# $FreeBSD$

PORTNAME=	drone
DISTVERSIONPREFIX=	v
DISTVERSION=	1.3.1
CATEGORIES=	devel

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Container-Native, Continuous Delivery Platform

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_TUPLE=	docker:go-docker:v1.0.0:godocker/vendor/docker.io/go-docker \
		99designs:httpsignatures-go:88528bf4ca7e:httpsignaturesgo/vendor/github.com/99designs/httpsignatures-go \
		Azure:azure-pipeline-go:v0.2.1:azurepipelinego/vendor/github.com/Azure/azure-pipeline-go \
		Azure:azure-storage-blob-go:v0.7.0:azurestorageblobgo/vendor/github.com/Azure/azure-storage-blob-go \
		alecthomas:template:a0175ee3bccc:template/vendor/github.com/alecthomas/template \
		alecthomas:units:2efee857e7cf:units/vendor/github.com/alecthomas/units \
		asaskevich:govalidator:ccb8e960c48f:govalidator/vendor/github.com/asaskevich/govalidator \
		aws:aws-sdk-go:v1.15.57:awssdkgo/vendor/github.com/aws/aws-sdk-go \
		beorn7:perks:3a771d992973:perks/vendor/github.com/beorn7/perks \
		bmatcuk:doublestar:v1.1.1:doublestar/vendor/github.com/bmatcuk/doublestar \
		buildkite:yaml:v2.1.0:yaml/vendor/github.com/buildkite/yaml \
		coreos:go-semver:v0.2.0:gosemver/vendor/github.com/coreos/go-semver \
		davecgh:go-spew:v1.1.1:gospew/vendor/github.com/davecgh/go-spew \
		dchest:authcookie:fbdef6e99866:authcookie/vendor/github.com/dchest/authcookie \
		dchest:uniuri:8902c56451e9:uniuri/vendor/github.com/dchest/uniuri \
		dgrijalva:jwt-go:v3.2.0:jwtgo/vendor/github.com/dgrijalva/jwt-go \
		docker:distribution:v2.7.1:distribution/vendor/github.com/docker/distribution \
		docker:go-connections:v0.3.0:goconnections/vendor/github.com/docker/go-connections \
		docker:go-units:v0.3.3:gounits/vendor/github.com/docker/go-units \
		drone:drone-go:v1.0.6:dronedronego/vendor/github.com/drone/drone-go \
		drone:drone-runtime:v1.1.0:dronedroneruntime/vendor/github.com/drone/drone-runtime \
		drone:drone-ui:v0.8.1:dronedroneui/vendor/github.com/drone/drone-ui \
		drone:drone-yaml:ad8ad9816fbf:dronedroneyaml/vendor/github.com/drone/drone-yaml \
		drone:envsubst:v1.0.1:envsubst/vendor/github.com/drone/envsubst \
		drone:go-license:v1.0.2:golicense/vendor/github.com/drone/go-license \
		drone:go-login:2a4df4f242a2:dronegologin/vendor/github.com/drone/go-login \
		drone:go-scm:fda52b1e0829:dronegoscm/vendor/github.com/drone/go-scm \
		drone:signal:v1.0.0:signal/vendor/github.com/drone/signal \
		dustin:go-humanize:v1.0.0:gohumanize/vendor/github.com/dustin/go-humanize \
		ghodss:yaml:v1.0.0:ghodssyaml/vendor/github.com/ghodss/yaml \
		go-chi:chi:v3.3.3:chi/vendor/github.com/go-chi/chi \
		go-chi:cors:v1.0.0:cors/vendor/github.com/go-chi/cors \
		go-ini:ini:v1.39.0:ini/vendor/github.com/go-ini/ini \
		go-sql-driver:mysql:v1.4.0:mysql/vendor/github.com/go-sql-driver/mysql \
		gogo:protobuf:100ba4e88506:protobuf/vendor/github.com/gogo/protobuf \
		golang:mock:v1.3.1:golangmock/vendor/github.com/golang/mock \
		golang:protobuf:v1.2.0:golangprotobuf/vendor/github.com/golang/protobuf \
		google:btree:4030bb1f1f0c:btree/vendor/github.com/google/btree \
		google:go-cmp:v0.3.0:googlegocmp/vendor/github.com/google/go-cmp \
		google:go-jsonnet:v0.12.1:gojsonnet/vendor/github.com/google/go-jsonnet \
		google:gofuzz:24818f796faf:gofuzz/vendor/github.com/google/gofuzz \
		google:wire:v0.2.1:wire/vendor/github.com/google/wire \
		googleapis:gnostic:v0.2.0:gnostic/vendor/github.com/googleapis/gnostic \
		gorhill:cronexpr:a557574d6c02:cronexpr/vendor/github.com/gorhill/cronexpr \
		gosimple:slug:v1.3.0:slug/vendor/github.com/gosimple/slug \
		gregjones:httpcache:c63ab54fda8f:httpcache/vendor/github.com/gregjones/httpcache \
		h2non:gock:v1.0.10:gock/vendor/github.com/h2non/gock \
		h2non:parth:b4df798d6542:parth/vendor/github.com/h2non/parth \
		hashicorp:errwrap:v1.0.0:errwrap/vendor/github.com/hashicorp/errwrap \
		hashicorp:go-cleanhttp:v0.5.0:gocleanhttp/vendor/github.com/hashicorp/go-cleanhttp \
		hashicorp:go-multierror:v1.0.0:gomultierror/vendor/github.com/hashicorp/go-multierror \
		hashicorp:go-retryablehttp:e651d75abec6:goretryablehttp/vendor/github.com/hashicorp/go-retryablehttp \
		hashicorp:go-rootcerts:v1.0.1:gorootcerts/vendor/github.com/hashicorp/go-rootcerts \
		hashicorp:golang-lru:v0.5.0:golanglru/vendor/github.com/hashicorp/golang-lru \
		hashicorp:nomad:134391155854:nomad/vendor/github.com/hashicorp/nomad \
		imdario:mergo:v0.3.6:mergo/vendor/github.com/imdario/mergo \
		jmespath:go-jmespath:0b12d6b521d8:gojmespath/vendor/github.com/jmespath/go-jmespath \
		jmoiron:sqlx:0dae4fefe7c0:sqlx/vendor/github.com/jmoiron/sqlx \
		joho:godotenv:v1.3.0:godotenv/vendor/github.com/joho/godotenv \
		json-iterator:go:v1.1.5:go/vendor/github.com/json-iterator/go \
		kelseyhightower:envconfig:v1.3.0:envconfig/vendor/github.com/kelseyhightower/envconfig \
		lib:pq:v1.0.0:pq/vendor/github.com/lib/pq \
		mattn:go-ieproxy:91bb50d98149:goieproxy/vendor/github.com/mattn/go-ieproxy \
		mattn:go-isatty:v0.0.4:goisatty/vendor/github.com/mattn/go-isatty \
		mattn:go-sqlite3:v1.9.0:gosqlite3/vendor/github.com/mattn/go-sqlite3 \
		matttproud:golang_protobuf_extensions:v1.0.1:golangprotobufextensions/vendor/github.com/matttproud/golang_protobuf_extensions \
		mitchellh:go-homedir:v1.1.0:mitchellhgohomedir/vendor/github.com/mitchellh/go-homedir \
		modern-go:concurrent:bacd9c7ef1dd:concurrent/vendor/github.com/modern-go/concurrent \
		modern-go:reflect2:4b7aa43c6742:reflect2/vendor/github.com/modern-go/reflect2 \
		natessilva:dag:7194b8dcc5c4:dag/vendor/github.com/natessilva/dag \
		opencontainers:go-digest:v1.0.0-rc1:godigest/vendor/github.com/opencontainers/go-digest \
		opencontainers:image-spec:v1.0.1:imagespec/vendor/github.com/opencontainers/image-spec \
		oxtoacart:bpool:4e1c5567d7c2:bpool/vendor/github.com/oxtoacart/bpool \
		peterbourgon:diskv:v2.0.1:diskv/vendor/github.com/peterbourgon/diskv \
		pkg:errors:v0.8.0:errors/vendor/github.com/pkg/errors \
		prometheus:client_golang:v0.8.0:clientgolang/vendor/github.com/prometheus/client_golang \
		prometheus:client_model:5c3871d89910:clientmodel/vendor/github.com/prometheus/client_model \
		prometheus:common:c7de2306084e:common/vendor/github.com/prometheus/common \
		prometheus:procfs:185b4288413d:procfs/vendor/github.com/prometheus/procfs \
		rainycape:unidecode:cb7f23ec59be:unidecode/vendor/github.com/rainycape/unidecode \
		robfig:cron:b41be1df6967:cron/vendor/github.com/robfig/cron \
		segmentio:ksuid:v1.0.2:ksuid/vendor/github.com/segmentio/ksuid \
		sergi:go-diff:v1.0.0:godiff/vendor/github.com/sergi/go-diff \
		sirupsen:logrus:44067abb194b:logrus/vendor/github.com/sirupsen/logrus \
		spf13:pflag:v1.0.3:pflag/vendor/github.com/spf13/pflag \
		unrolled:secure:4b6b7cf51606:secure/vendor/github.com/unrolled/secure \
		vinzenz:yaml:91409cdd725d:vinzenzyaml/vendor/github.com/vinzenz/yaml \
		golang:crypto:c2843e01d9a2:golangcrypto/vendor/golang.org/x/crypto \
		golang:net:d8887717615a:golangnet/vendor/golang.org/x/net \
		golang:oauth2:d668ce993890:oauth2/vendor/golang.org/x/oauth2 \
		golang:sync:112230192c58:golangsync/vendor/golang.org/x/sync \
		golang:sys:d0b11bdaac8a:golangsys/vendor/golang.org/x/sys \
		golang:text:v0.3.0:text/vendor/golang.org/x/text \
		golang:time:85acf8d2951c:time/vendor/golang.org/x/time \
		alecthomas:kingpin:v2.2.6:kingpin/vendor/gopkg.in/alecthomas/kingpin.v2 \
		h2non:gock:v1.0.14:h2nongock/vendor/gopkg.in/h2non/gock.v1 \
		go-inf:inf:v0.9.1:infv0/vendor/gopkg.in/inf.v0 \
		go-yaml:yaml:v2.2.2:goyamlyaml/vendor/gopkg.in/yaml.v2 \
		kubernetes:api:d04500c8c3dd:api/vendor/k8s.io/api \
		kubernetes:apimachinery:eb8c8024849b:apimachinery/vendor/k8s.io/apimachinery \
		kubernetes:client-go:v10.0.0:clientgo/vendor/k8s.io/client-go \
		kubernetes:klog:v0.1.0:klog/vendor/k8s.io/klog \
		kubernetes-sigs:yaml:v1.1.0:kubernetesyaml/vendor/sigs.k8s.io/yaml

GO_BUILDFLAGS=	-tags oss
GO_TARGET=	github.com/drone/drone/cmd/drone-server

PLIST_FILES=	bin/drone-server

.include <bsd.port.mk>
