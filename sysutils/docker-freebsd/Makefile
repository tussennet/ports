# Created by: kmoore@FreeBSD.org
# $FreeBSD: head/sysutils/docker-freebsd/Makefile 490308 2019-01-14 18:09:20Z joneum $

PORTNAME=	docker-freebsd
PORTVERSION=	20170907
DISTVERSIONPREFIX=	v
DISTVERSIONSUFFIX=	-ce
CATEGORIES=	sysutils

MAINTAINER=	joneum@FreeBSD.org
COMMENT=	Docker containment system

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	go>=1.4:lang/go
LIB_DEPENDS=	libepoll-shim.so:devel/libepoll-shim

USES=		go gmake

USE_GITHUB=	yes
GH_ACCOUNT=	freebsd-docker
GH_PROJECT=	docker
GH_TAGNAME=	969c6fd
GH_SUBDIR=	src/github.com/docker/docker

GH_TUPLE=	freebsd-docker:containerd:00a99ef:containerd/src/github.com/containerd/containerd \
		freebsd-docker:libnetwork:acb2447:libnetwork/src/github.com/docker/libnetwork \
		freebsd-docker:runc:9c2d8d1:runc/src/github.com/runc/runc \
		krallin:tini:v0.16.1:tini/src/github.com/krallin/tini

PLIST_FILES=	bin/docker \
		bin/dockerd \
		bin/docker-containerd \
		bin/docker-containerd-ctr \
		bin/docker-containerd-shim \
		bin/docker-proxy

USE_RC_SUBR=	docker

GO_PKGNAME=	github.com/docker/docker
GO_TARGET=	${GO_PKGNAME}/cmd/dockerd
SSP_UNSAFE=	yes

STRIP=	#

#post-extract:
#	${LN} -sf ${WRKSRC}/src ${WRKDIR}
#	${CP} ${FILESDIR}/Makefile.freebsd ${WRKSRC}

#do-build:
#	@cd ${GO_WRKSRC} && export DOCKER_GITCOMMIT=${GH_TAGNAME} && ${SETENV} ${MAKE_ENV} ${GO_ENV} ${GMAKE} -f Makefile.freebsd #binary #daemon_freebsd.go binary

#do-install:
#.for f in ${PLIST_FILES}
#	${INSTALL_PROGRAM} ${GO_WRKSRC}/bundles/${f} ${STAGEDIR}${PREFIX}/bin
#.endfor

.include <bsd.port.mk>
