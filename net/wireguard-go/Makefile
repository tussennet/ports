# $FreeBSD$

PORTNAME=	wireguard-go
PORTVERSION=	master
CATEGORIES=	net
MASTER_SITES=	https://git.zx2c4.com/wireguard-go/snapshot/

MAINTAINER=	decke@FreeBSD.org
COMMENT=	WireGuard implementation in Go

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/COPYING

BUILD_DEPENDS=	${LOCALBASE}/bin/go:lang/go

USES=		gmake tar:xz

GO_CMD=		${LOCALBASE}/bin/go
GO_ENV+=	GOPATH="${WRKDIR}:${LOCALBASE}/bin/go" \
		CGO_FLAGS="-I${LOCALBASE}/include" \
		CGO_LDFLAGS="-L${LOCALBASE}/lib"

PLIST_FILES=	bin/${PORTNAME}

post-extract:
	@(cd ${WRKSRC}; ${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} get -v -d)

.include <bsd.port.mk>
