# $FreeBSD$

PORTNAME=	drone-runner-exec
DISTVERSIONPREFIX=	v
DISTVERSION=	1.0.0-beta.5
CATEGORIES=	devel

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Drone pipeline runner that executes builds on the host machine

# LICENSE: Polyform-Noncommercial-1.0.0 OR Polyform-Small-Business-1.0.0 OR Polyform-Free-Trial-1.0.0
#LICENSE=	NONE
#LICENSE_FILE=	${WRKSRC}/LICENSE.md

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	drone-runners
GH_TUPLE=	99designs:basicauth-go:2a93ba0f464d:basicauthgo/vendor/github.com/99designs/basicauth-go \
		99designs:httpsignatures-go:88528bf4ca7e:httpsignaturesgo/vendor/github.com/99designs/httpsignatures-go \
		alecthomas:template:a0175ee3bccc:template/vendor/github.com/alecthomas/template \
		alecthomas:units:2efee857e7cf:units/vendor/github.com/alecthomas/units \
		bmatcuk:doublestar:v1.1.1:doublestar/vendor/github.com/bmatcuk/doublestar \
		buildkite:yaml:v2.1.0:yaml/vendor/github.com/buildkite/yaml \
		coreos:go-semver:v0.3.0:gosemver/vendor/github.com/coreos/go-semver \
		dchest:uniuri:8902c56451e9:uniuri/vendor/github.com/dchest/uniuri \
		docker:go-units:v0.4.0:gounits/vendor/github.com/docker/go-units \
		drone:drone-git:5b0a12bcb36c:dronedronegit/vendor/github.com/drone/drone-git \
		drone:drone-go:4d6116b897ba:dronedronego/vendor/github.com/drone/drone-go \
		drone:envsubst:v1.0.2:envsubst/vendor/github.com/drone/envsubst \
		drone:runner-go:v1.2.1:dronerunnergo/vendor/github.com/drone/runner-go \
		drone:signal:v1.0.0:signal/vendor/github.com/drone/signal \
		golang:mock:v1.3.1:mock/vendor/github.com/golang/mock \
		google:go-cmp:v0.3.0:gocmp/vendor/github.com/google/go-cmp \
		gosimple:slug:v1.5.0:slug/vendor/github.com/gosimple/slug \
		hashicorp:errwrap:v1.0.0:errwrap/vendor/github.com/hashicorp/errwrap \
		hashicorp:go-multierror:v1.0.0:gomultierror/vendor/github.com/hashicorp/go-multierror \
		joho:godotenv:v1.3.0:godotenv/vendor/github.com/joho/godotenv \
		kardianos:service:v1.0.0:service/vendor/github.com/kardianos/service \
		kelseyhightower:envconfig:v1.3.0:envconfig/vendor/github.com/kelseyhightower/envconfig \
		konsorten:go-windows-terminal-sequences:v1.0.1:gowindowsterminalsequences/vendor/github.com/konsorten/go-windows-terminal-sequences \
		mattn:go-isatty:v0.0.8:goisatty/vendor/github.com/mattn/go-isatty \
		natessilva:dag:7194b8dcc5c4:dag/vendor/github.com/natessilva/dag \
		orandin:lumberjackrus:v1.0.1:lumberjackrus/vendor/github.com/orandin/lumberjackrus \
		rainycape:unidecode:cb7f23ec59be:unidecode/vendor/github.com/rainycape/unidecode \
		sirupsen:logrus:v1.4.2:sirupsenlogrus/vendor/github.com/sirupsen/logrus \
		golang:crypto:cc06ce4a13d4:crypto/vendor/golang.org/x/crypto \
		golang:net:eb5bcb51f2a3:net/vendor/golang.org/x/net \
		golang:sync:112230192c58:sync/vendor/golang.org/x/sync \
		golang:sys:953cdadca894:golangsys/vendor/golang.org/x/sys \
		golang:text:v0.3.0:text/vendor/golang.org/x/text \
		alecthomas:kingpin:v2.2.6:kingpin/vendor/gopkg.in/alecthomas/kingpin.v2 \
		natefinch:lumberjack:v2.0.0:lumberjack/vendor/gopkg.in/natefinch/lumberjack.v2 \
		go-yaml:yaml:v2.2.2:goyamlyaml/vendor/gopkg.in/go-yaml/yaml

PLIST_FILES=	bin/${PORTNAME}

.include <bsd.port.mk>
