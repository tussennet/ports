# Created by: Andrew Pantyukhin <infofarmer@FreeBSD.org>
# $FreeBSD: head/deskutils/py-vobject/Makefile 472884 2018-06-20 17:05:41Z mat $

PORTNAME=	vobject
PORTVERSION=	0.9.5
CATEGORIES=	deskutils python
PKGNAMEPREFIX=	${PYTHON_PKGNAMEPREFIX}

MAINTAINER=	jun-g@daemonfreaks.com
COMMENT=	ICalendar library

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE-2.0.txt

RUN_DEPENDS=	${PYTHON_PKGNAMEPREFIX}dateutil>=2.4.0:devel/py-dateutil@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}six>0:devel/py-six@${PY_FLAVOR}

NO_ARCH=	yes
USES=		python
USE_PYTHON=	distutils autoplist

USE_GITHUB=	yes
GH_ACCOUNT=	eventable

do-test:
	@cd ${WRKSRC} && ${PYTHON_CMD} tests.py

.include <bsd.port.mk>
