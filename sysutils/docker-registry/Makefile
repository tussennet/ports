# $FreeBSD$

PORTNAME=	docker-registry
PORTVERSION=	2.7.1
DISTVERSIONPREFIX=v
CATEGORIES=	sysutils

MAINTAINER=	freebsd@yapsoft.it
COMMENT=	Docker implementation of the OCI Distribution Specification

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	docker
GH_PROJECT=	distribution

GO_PKGNAME=	github.com/${GH_ACCOUNT}/${GH_PROJECT}
GO_TARGET=	github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd/registry

PLIST_FILES=	bin/registry \
		${ETCDIR}/config.yml.sample

#USE_RC_SUBR=	docker_registry

post-install:
	${MKDIR} ${STAGEDIR}${ETCDIR}
	${INSTALL_DATA} ${WRKSRC}/cmd/registry/config-example.yml ${STAGEDIR}${PREFIX}/etc/${PORTNAME}/config.yml.sample

.include <bsd.port.mk>
