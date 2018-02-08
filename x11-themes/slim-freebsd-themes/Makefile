# $FreeBSD$

PORTNAME=	slim-freebsd-themes
PORTVERSION=	1.0
CATEGORIES=	x11-themes
MASTER_SITES=	LOCAL/decke:local
DISTFILES=	slim-freebsd-beastie.tar.gz:local \
		slim-freebsd-simple.tar.gz:local \
		slim-freebsd-black-theme.tar.gz

MAINTAINER=	decke@FreeBSD.org
COMMENT=	FreeBSD Theme pack for SLiM login app

LICENSE=	CC-BY-SA-3.0

USE_GITHUB=	yes
GH_TUPLE=	lebarondemerde:slim-freebsd-black-theme:1.0

NO_BUILD=	yes
NO_WRKSUBDIR=	yes
SHAREMODE=	644
THEMEDIR=	${PREFIX}/share/slim/themes

do-install:
	${MKDIR} ${STAGEDIR}${THEMEDIR}
	(cd ${WRKSRC} && ${COPYTREE_SHARE} . ${STAGEDIR}${THEMEDIR}/)

.include <bsd.port.mk>
