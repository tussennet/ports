# $FreeBSD$

PORTNAME=	docker-engine
DISTVERSIONPREFIX=	v
DISTVERSION=	19.03.3
CATEGORIES=	sysutils

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Docker Engine based on moby

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	bash:shells/bash

USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	docker
GH_PROJECT=	engine
GH_SUBDIR=	src/github.com/docker/docker

GO_TARGET=	./cmd/dockerd

do-build:
	@cd ${GO_WRKSRC} && export DOCKER_GITCOMMIT=${GH_TAGNAME} && ${SETENV} ${GO_ENV} ./hack/make.sh binary

.include <bsd.port.mk>
