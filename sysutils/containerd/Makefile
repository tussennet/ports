# $FreeBSD$

PORTNAME=	containerd
DISTVERSIONPREFIX=	
DISTVERSION=	481103c8793316c118d9f795cde18060847c370e # somewhere between v1.2.6 and v1.2.7 but this is dependency commit in moby v19.03.12...
CATEGORIES=	sysutils

MAINTAINER=	decke@FreeBSD.org
COMMENT=	containerd part of docker

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	bash:shells/bash
USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	containerd
GH_PROJECT=	containerd
GH_SUBDIR=	src/github.com/containerd/containerd

# GO_TARGET=	./

do-build:
	@cd ${GO_WRKSRC} && export DOCKER_GITCOMMIT=${GH_TAGNAME} && ${SETENV} ${GO_ENV} gmake 

.include <bsd.port.mk>
