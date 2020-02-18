# $FreeBSD$

PORTNAME=	addrwatch
DISTVERSIONPREFIX=v
DISTVERSION=	1.0.1
CATEGORIES=	net

MAINTAINER=	m.muenz@gmail.com
COMMENT=	Supports IP/Ethernet pairing for IPv4 and IPv6

LICENSE=	GPLv3

LIB_DEPENDS=	libevent.so:devel/libevent \
		libpcap.so:net/libpcap \
		libargp.so:devel/argp-standalone

USES=		autoreconf gmake pkgconfig
GNU_CONFIGURE=	yes
USE_GITHUB=	yes
GH_ACCOUNT=	fln

.include <bsd.port.mk>
