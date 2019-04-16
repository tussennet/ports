# $FreeBSD$

PORTNAME=	stubby
PORTVERSION=	0.2.6
DISTVERSIONPREFIX=	v
CATEGORIES=	dns ipv6

MAINTAINER=	decke@FreeBSD.org
COMMENT=	DNS Privacy stub resolver

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/COPYING

LIB_DEPENDS=	libgetdns.so:dns/getdns \
		libyaml.so:textproc/libyaml

USES=		autoreconf gmake

USE_GITHUB=	yes
GH_ACCOUNT=	getdnsapi

GNU_CONFIGURE=	yes
CONFIGURE_ARGS=	CFLAGS="-I${LOCALBASE}/include" \
		LDFLAGS="-L${LOCALBASE}/lib"

USE_RC_SUBR=	stubby

.include <bsd.port.mk>
