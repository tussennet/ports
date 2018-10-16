# $FreeBSD$

PORTNAME=	wireguard-go
PORTVERSION=	85b2378a07126fb64460f0a23e7a150d0e63d909
CATEGORIES=	net
MASTER_SITES=	https://git.zx2c4.com/wireguard-go/snapshot/ \
		LOCAL/decke
DISTFILES=	${PORTNAME}-${PORTVERSION}${EXTRACT_SUFX} \
		${PORTNAME}-${PORTVERSION}-vendor${EXTRACT_SUFX}

MAINTAINER=	decke@FreeBSD.org
COMMENT=	WireGuard implementation in Go

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/COPYING

BUILD_DEPENDS=	go:lang/go

USES=		gmake tar:xz

ALL_TARGET=	default
MAKE_ENV+=	GOFLAGS=-mod=vendor

PLIST_FILES=	bin/${PORTNAME}

# The vendor tar file is generated with the script from
# https://people.freebsd.org/~decke/wireguard-go-genvendor.sh

.include <bsd.port.mk>
