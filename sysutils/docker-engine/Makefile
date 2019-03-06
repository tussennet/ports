# $FreeBSD$

PORTNAME=	docker-engine
PORTVERSION=	18.09.3
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Docker Engine based on moby

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	docker
GH_PROJECT=	engine
GH_SUBDIR=	src/github.com/docker/docker

GO_PKGNAME=	github.com/docker/docker
GO_TARGET=	${GO_PKGNAME}/cmd/dockerd

do-build:
	@cd ${GO_WRKSRC} && export DOCKER_GITCOMMIT=${GH_TAGNAME} && ${SETENV} ${GO_ENV} ./hack/make.sh binary

.include <bsd.port.mk>
