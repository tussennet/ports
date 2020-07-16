# $FreeBSD$

PORTNAME=	wireguard-kmod
DISTVERSION=	0.0.0.20200630
CATEGORIES=	net
MASTER_SITES=	https://git.zx2c4.com/wireguard-kmod/snapshot/ \
		LOCAL/decke
DISTNAME=	wireguard-freebsd-${PORTVERSION}

MAINTAINER=	decke@FreeBSD.org
COMMENT=	WireGuard implementation for the FreeBSD kernel

LICENSE=	BSD2CLAUSE

USES=		kmod tar:xz

PLIST_FILES=	${KMODDIR}/if_wg.ko

.include <bsd.port.mk>
