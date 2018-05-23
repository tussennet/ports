# $FreeBSD$

PORTNAME=	wireguard
PORTVERSION=	0.0.2018052306
CATEGORIES=	net
MASTER_SITES=	https://git.zx2c4.com/WireGuard/snapshot/ \
		http://web.bluelife.at/
PKGNAMESUFFIX=	-tools
DISTNAME=	WireGuard-${PORTVERSION}

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Fast, modern and secure VPN Tunnel

LICENSE=	GPLv2
LICENSE_FILE=	${WRKDIR}/${DISTNAME}/COPYING

BUILD_DEPENDS=	bash:shells/bash
RUN_DEPENDS=	bash:shells/bash \
		wireguard-go:net/wireguard-go

USES=		gmake tar:xz

WRKSRC=		${WRKDIR}/${DISTNAME}/src/tools

ALL_TARGET=
MAKE_ARGS=	WITH_WGQUICK=yes
MAKE_ENV+=	MANDIR="${MANPREFIX}/man" \
		SYSCONFDIR="${PREFIX}/etc"

.include <bsd.port.mk>
