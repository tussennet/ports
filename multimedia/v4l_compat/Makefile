# $FreeBSD$

PORTNAME=	v4l_compat
PORTVERSION=	4.17.6
CATEGORIES=	multimedia
MASTER_SITES=	LOCAL/decke
DISTNAME=	linux-${PORTVERSION}-dvb-headers

MAINTAINER=	multimedia@FreeBSD.org
COMMENT=	Linux Kernel headers needed for Video4Linux

LICENSE=	GPLv2 LGPL21
LICENSE_COMB=	multi

NO_BUILD=	yes
NO_ARCH=	yes

WRKSRC=		${WRKDIR}/linux
USES=		tar:xz

post-patch:
	# fix weird types, except a few headers which have compat code for these types.
	@${FIND} ${WRKDIR} -type f \( -iname "*.h" ! -iname "input.h" ! -iname "videodev2.h" \) | \
		${XARGS} -n 10 ${REINPLACE_CMD} \
		-e 's|linux/types.h|sys/types.h|g; s|linux/ioctl.h|sys/ioctl.h|g; s|__user||g; \
		s|__u8|uint8_t|g; s|__s8|int8_t|g; \
		s|__u16|uint16_t|g; s|__s16|int16_t|g; \
		s|__u32|uint32_t|g; s|__s32|int32_t|g; \
		s|__u64|uint64_t|g; s|__s64|int64_t|g'

do-install:
	${MKDIR} ${STAGEDIR}${PREFIX}/include/linux/ && \
	(cd ${WRKSRC} && ${COPYTREE_SHARE} "*.h dvb/*.h" ${STAGEDIR}${PREFIX}/include/linux/)

.include <bsd.port.mk>
