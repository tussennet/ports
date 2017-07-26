# Created by: dave@dogwood.com
# $FreeBSD: head/net/olsrd/Makefile 434144 2017-02-14 22:24:59Z tijl $

PORTNAME=	olsrd
PORTVERSION=	0.9.6.2
CATEGORIES=	net

MAINTAINER=	koue@chaosophia.net
COMMENT=	OLSR routing daemon

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/license.txt

MAKE_ARGS=	VERBOSE=1
USES=		bison gmake tar:bzip2
USE_LDCONFIG=	yes
USE_RC_SUBR=	olsrd

USE_GITHUB=	yes
GH_ACCOUNT=	OLSR
GH_TAGNAME=	v${PORTVERSION}

ALL_TARGET=	build_all
INSTALL_TARGET=	install_all

OPTIONS_DEFINE=	MULTI_IF
MULTI_IF_DESC=	IPv4 multiple interface support

MULTI_IF_LIB_DEPENDS=	libnet.so:net/libnet
MULTI_IF_MAKE_ARGS=	LIBNET=1

post-patch:
	@${REINPLACE_CMD} -e 's|/sbin/ldconfig|:|g' ${WRKSRC}/make/Makefile.fbsd
	@${CP} ${WRKSRC}/src/cfgparser/cfgfile_gen.c \
		${WRKSRC}/src/cfgparser/olsrd_conf.h \
		${WRKSRC}/lib/httpinfo/src/

post-install:
	@${STRIP_CMD} ${STAGEDIR}${PREFIX}/sbin/${PORTNAME}
	@${STRIP_CMD} ${STAGEDIR}${PREFIX}/lib/*.so.*

.include <bsd.port.mk>
