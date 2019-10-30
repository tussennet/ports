# $FreeBSD$

PORTNAME=	bluelife-config-graz
PORTVERSION=	2019.10.30
PORTREVISION=	0
CATEGORIES=	misc
MASTER_SITES=	#
DISTFILES=	#

MAINTAINER=	decke@bluelife.at
COMMENT=	Common FreeBSD Server Configuration for bluelife.at Network

LICENSE=	CC0-1.0

NO_BUILD=	yes
NO_ARCH=	yes

do-install:
	${MKDIR} ${STAGEDIR}
	${CP} -pr ${FILESDIR}/* ${STAGEDIR}/

.include <bsd.port.mk>
