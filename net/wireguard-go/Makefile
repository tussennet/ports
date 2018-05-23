# $FreeBSD$

PORTNAME=	wireguard-go
PORTVERSION=	master
CATEGORIES=	net
MASTER_SITES=	https://git.zx2c4.com/wireguard-go/snapshot/ \
		LOCAL/decke \
		http://web.bluelife.at/
DISTFILES=	${PORTNAME}-${PORTVERSION}${EXTRACT_SUFX} \
		${PORTNAME}-0.0.2018052317-vendor${EXTRACT_SUFX}

MAINTAINER=	decke@FreeBSD.org
COMMENT=	WireGuard implementation in Go

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/COPYING

BUILD_DEPENDS=	go:lang/go \
		dep:devel/dep

USES=		gmake tar:xz

PLIST_FILES=	bin/${PORTNAME}

# To generate the vendor source tarball with all dependencies
genvendor:
	(cd ${WRKDIR} ; ${TAR} cvfJ ${PORTNAME}-${PORTVERSION}-vendor.tar.xz ${PORTNAME}-${PORTVERSION}/vendor)

.include <bsd.port.mk>
