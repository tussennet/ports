PORTNAME=	docker-registry
PORTVERSION=	2.7.1
DISTVERSIONPREFIX=v
PORTREVISION=	1
CATEGORIES=	sysutils

MAINTAINER=	freebsd@yapsoft.it
COMMENT=	The Docker implementation of the OCI Distribution Specification.

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

PLIST_FILES=	bin/registry \
		etc/${PORTNAME}/config.yml.sample

USE_RC_SUBR=	docker_registry
USE_GITHUB=	yes

USES=		go

GH_ACCOUNT=	docker
GH_PROJECT=	distribution

GO_PKGNAME=	github.com/${GH_ACCOUNT}/${GH_PROJECT}
GO_TARGET=	github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd/registry

post-install:
	${MKDIR} ${STAGEDIR}${PREFIX}/etc/${PORTNAME}
	${INSTALL_DATA} ${WRKSRC}/cmd/registry/config-example.yml ${STAGEDIR}${PREFIX}/etc/${PORTNAME}/config.yml.sample

.include <bsd.port.mk>
