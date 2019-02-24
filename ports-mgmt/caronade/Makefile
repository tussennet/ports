# $FreeBSD$

PORTNAME=	caronade
#DISTVERSIONPREFIX=v
DISTVERSION=	master
CATEGORIES=	ports-mgmt
MASTER_SITES=	https://code.bluelife.at/decke/caronade/archive/
DISTFILES=	${DISTVERSIONPREFIX}${DISTVERSION}${DISTVERSIONSUFFIX}${EXTRACT_SUFX}
DIST_SUBDIR=	${PORTNAME}

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Small and light tool to help with FreeBSD Ports CI

LICENSE=	BSD2CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go
USE_RC_SUBR=	caronade

USE_GITHUB=	nodefault
GH_TUPLE=	rakyll:statik:v0.1.5:statik/vendor/github.com/rakyll/statik \
		go-yaml:yaml:v2.2.2:yaml/vendor/gopkg.in/yaml.v2

GO_PKGNAME=	code.bluelife.at/decke/caronade
MAKE_ENV+=	GOFLAGS=-mod=vendor

post-install:
		${MKDIR} ${STAGEDIR}${ETCDIR}
		${INSTALL_DATA} ${WRKSRC}/caronade.yaml ${STAGEDIR}${ETCDIR}/caronade.yaml.sample

		${MKDIR} ${STAGEDIR}${PREFIX}/${PORTNAME}
		(cd ${WRKSRC} && ${COPYTREE_SHARE} work ${STAGEDIR}${PREFIX}/${PORTNAME})

		${MKDIR} ${STAGEDIR}${PREFIX}/${PORTNAME}/builds

.include <bsd.port.mk>
