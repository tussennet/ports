# Created by: Rozhuk Ivan 2018 <rozhuk.im@gmail.com>
# $FreeBSD: head/multimedia/kodi-addon-peripheral-joystick/Makefile 488341 2018-12-25 20:25:39Z tcberner $

PORTNAME=	kodi-addon-peripheral-joystick
PORTVERSION=	1.4.9
CATEGORIES=	multimedia

MAINTAINER=	rozhuk.im@gmail.com
COMMENT=	Joystick support for Kodi

LICENSE=	GPLv2+

BUILD_DEPENDS=	${LOCALBASE}/include/kodi/libXBMC_addon.h:multimedia/kodi
LIB_DEPENDS=	libp8-platform.so:devel/p8-platform \
		libkodiplatform.so:devel/kodi-platform \
		libudev.so:devel/libudev-devd
RUN_DEPENDS=	kodi:multimedia/kodi

USES=		cmake:noninja

USE_GITHUB=	yes
GH_ACCOUNT=	xbmc
GH_PROJECT=	peripheral.joystick
GH_TAGNAME=	${PORTVERSION}-Leia

post-extract:
	@${MKDIR} ${WRKSRC}/src/linux/
	@${CP} -f ${FILESDIR}/input.h ${WRKSRC}/src/linux/
	@${CP} -f ${FILESDIR}/input-event-codes.h ${WRKSRC}/src/linux/
	@${CP} -f ${FILESDIR}/joystick.h ${WRKSRC}/src/linux/
	@${CP} -f ${FILESDIR}/types.h ${WRKSRC}/src/linux/

post-patch:
	@${REINPLACE_CMD} -e \
	    's|if(HAVE_LINUX_JOYSTICK_H)|if(1)|g ; \
	     s|if(UDEV_FOUND)|if(1)|g' \
		${WRKSRC}/CMakeLists.txt
	@${REINPLACE_CMD} -e 's|#define __stat64 stat64|#define __stat64 stat|g' \
		${WRKSRC}/src/utils/CommonIncludes.h

.include <bsd.port.mk>
