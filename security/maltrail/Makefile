# $FreeBSD: head/security/maltrail/Makefile 495777 2019-03-15 08:47:18Z decke $

PORTNAME=	maltrail
PORTVERSION=	0.14
CATEGORIES=	security python

MAINTAINER=	m.muenz@gmail.com
COMMENT=	Malicious traffic detection system, utilizing public (black)lists

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	${PYTHON_PKGNAMEPREFIX}pcapy>=0:net/py-pcapy@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}sqlite3>=0:databases/py-sqlite3@${PY_FLAVOR}

USES=		python:2.7,run

USE_GITHUB=	yes
GH_ACCOUNT=	stamparm

NO_BUILD=	yes
NO_ARCH=	yes

do-install:
		${MKDIR} ${STAGEDIR}${DATADIR} && \
		(cd ${WRKSRC} && ${COPYTREE_SHARE} "*" ${STAGEDIR}${DATADIR})

.include <bsd.port.mk>
