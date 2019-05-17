# $FreeBSD$

PORTNAME=	wazuh
DISTVERSIONPREFIX=	v
DISTVERSION=	3.9.0
CATEGORIES=	security
PKGNAMESUFFIX=	-agent

MAINTAINER=	m.muenz@gmail.com
COMMENT=	Security tool to monitor and check logs and intrusions

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	curl:ftp/curl
RUN_DEPENDS=	curl:ftp/curl

USES=		gmake perl5 readline ssl

USE_GITHUB=	yes

USERS=		ossec ossecm ossecr
GROUPS=		ossec

do-build:
		@cd ${WRKSRC}/src && ${GMAKE} deps && ${GMAKE}

do-install:
		@cd ${WRKSRC} && ${SH} install.sh

.include <bsd.port.mk>
