# $FreeBSD$

PORTNAME=	wireguard-go
PORTVERSION=	master
CATEGORIES=	net
MASTER_SITES=	https://git.zx2c4.com/wireguard-go/snapshot/:wgo \
		LOCAL/decke:wgodeps
DISTFILES=	${PORTNAME}-${PORTVERSION}${EXTRACT_SUFX}:wgo \
		${PORTNAME}-deps-20180522.tar.xz:wgodeps

MAINTAINER=	decke@FreeBSD.org
COMMENT=	WireGuard implementation in Go

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/COPYING

BUILD_DEPENDS=	go:lang/go

USES=		gmake tar:xz

MAKE_ENV+=	GOPATH="${WRKDIR}:${LOCALBASE}/share/go"

PLIST_FILES=	bin/${PORTNAME}

.include <bsd.port.mk>
