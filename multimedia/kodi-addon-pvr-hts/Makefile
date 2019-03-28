# Created by: Bernhard Froehlich <decke@FreeBSD.org>
# $FreeBSD$

PORTNAME=	kodi-addon-pvr-hts
PORTVERSION=	4.4.16
CATEGORIES=	multimedia

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Tvheadend PVR addon for Kodi

LICENSE=	GPLv2+

BUILD_DEPENDS=	${LOCALBASE}/include/kodi/libXBMC_addon.h:multimedia/kodi-devel
LIB_DEPENDS=	libp8-platform.so:devel/p8-platform \
		libkodiplatform.so:devel/kodi-platform
RUN_DEPENDS=	kodi:multimedia/kodi-devel

USES=		cmake:noninja

USE_GITHUB=	yes
GH_ACCOUNT=	kodi-pvr
GH_PROJECT=	pvr.hts
GH_TAGNAME=	${PORTVERSION}-Leia

.include <bsd.port.mk>
