# $FreeBSD: head/devel/kodi-platform/Makefile 473636 2018-06-30 19:52:10Z decke $

PORTNAME=	kodi-platform
PORTVERSION=	20180302
CATEGORIES=	devel
PORTREVISION=	1

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Kodi platform support library

LICENSE=	GPLv2+

BUILD_DEPENDS=	${LOCALBASE}/share/kodi/cmake/KodiConfig.cmake:multimedia/kodi
LIB_DEPENDS=	libtinyxml.so:textproc/tinyxml \
		libp8-platform.so:devel/p8-platform

USES=		cmake

USE_GITHUB=	yes
GH_ACCOUNT=	xbmc
GH_TAGNAME=	e8574b8

USE_LDCONFIG=	yes

.include <bsd.port.mk>
