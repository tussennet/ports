# $FreeBSD$

PORTNAME=	dvb-apps
PORTVERSION=	3d43b280298c
CATEGORIES=	multimedia
MASTER_SITES=	LOCAL/decke

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Linux DVB API applications and utilities

LICENSE=	LGPL21
LICENSE_FILE=	${WRKSRC}/COPYING.LGPL

BUILD_DEPENDS=	${LOCALBASE}/include/linux/dvb/version.h:multimedia/v4l_compat

USES=		tar:bzip2 gmake iconv

USE_LDCONFIG=	yes

MAKE_ARGS+=	prefix="${PREFIX}"
CFLAGS+=	-I${LOCALBASE}/include -Wno-int-conversion
LDFLAGS+=	-L${LOCALBASE}/lib

post-patch:
	@${REINPLACE_CMD} 's|/usr/include|${LOCALBASE}/include|' \
		${WRKSRC}/Makefile \
		${WRKSRC}/util/av7110_loadkeys/generate-keynames.sh \
		${WRKSRC}/util/szap/Makefile
	@${REINPLACE_CMD} 's|make|gmake|' \
		${WRKSRC}/test/Makefile
	@${REINPLACE_CMD} 's|<malloc.h>|<stdlib.h>|' \
		${WRKSRC}/lib/libdvbcfg/dvbcfg_scanfile.c \
		${WRKSRC}/lib/libdvbcfg/dvbcfg_zapchannel.c
	@${REINPLACE_CMD} 's|<linux/types.h>|<sys/types.h>|' \
		${WRKSRC}/util/dib3000-watch/dib3000-watch.c \
		${WRKSRC}/util/scan/lnb.c \
		${WRKSRC}/lib/libdvbsec/dvbsec_api.c \
		${WRKSRC}/lib/libdvbsec/dvbsec_cfg.c \
		${WRKSRC}/lib/libdvbapi/dvbvideo.c \
		${WRKSRC}/test/test_vevent.c \
		${WRKSRC}/test/test_stillimage.c \
		${WRKSRC}/test/test_av.c

.include <bsd.port.mk>
