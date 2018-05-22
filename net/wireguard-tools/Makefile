# $FreeBSD$

PORTNAME=	wireguard-tools
PORTVERSION=	master
CATEGORIES=	net
MASTER_SITES=	https://git.zx2c4.com/WireGuard/snapshot/
DISTNAME=	WireGuard-${PORTVERSION}

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Tools for wireguard

LICENSE=	GPLv2
LICENSE_FILE=	${WRKDIR}/${DISTNAME}/COPYING

BUILD_DEPENDS=	wireguard-go:net/wireguard-go \
		bash:shells/bash
RUN_DEPENDS=	bash:shells/bash

USES=		gmake tar:xz

WRKSRC=		${WRKDIR}/${DISTNAME}/src/tools

ALL_TARGET=
MAKE_ARGS=	WITH_WGQUICK=yes
MAKE_ENV+=	MANDIR="${MANPREFIX}/man"

.include <bsd.port.mk>
