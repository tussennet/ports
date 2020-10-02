# $FreeBSD$

PORTNAME=	runc
DISTVERSIONPREFIX=	v
DISTVERSION=	1.0.0-rc10
CATEGORIES=	sysutils

MAINTAINER=	decke@FreeBSD.org
COMMENT=	runc part of docker

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	bash:shells/bash

USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	opencontainers
GH_PROJECT=	runc
GH_SUBDIR=	src/github.com/opencontainers/runc

# GO_TARGET=	./

.include <bsd.port.mk>
