# $FreeBSD$

PORTNAME=	wazuh
DISTVERSIONPREFIX=	v
DISTVERSION=	3.9.0
CATEGORIES=	security
PKGNAMESUFFIX=	-agent
DISTFILES=	cJSON.tar.gz src_cpython.tar.gz curl.tar.gz libdb.tar.gz libffi.tar.gz \
		libyaml.tar.gz openssl.tar.gz procps.tar.gz sqlite.tar.gz  zlib.tar.gz \
		audit-userspace.tar.gz msgpack.tar.gz
MASTER_SITES=	https://packages.wazuh.com/deps/3.9/
DIST_SUBDIR=	${PORTNAME}-${DISTVERSION}
EXTRACT_ONLY=	${DISTNAME}${EXTRACT_SUFX}


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

# extract all extra distfiles in src/external
post-extract:
	@for file in ${DISTFILES}; do \
		if ! (cd ${WRKSRC}/src/external && ${EXTRACT_CMD} ${EXTRACT_BEFORE_ARGS} ${_DISTDIR}/$$file ${EXTRACT_AFTER_ARGS}); \
		then \
			exit 1; \
		fi; \
	done

do-build:
		@cd ${WRKSRC}/src && ${GMAKE}

do-install:
		@cd ${WRKSRC} && ${SH} install.sh

.include <bsd.port.mk>
