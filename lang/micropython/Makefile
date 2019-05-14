# $FreeBSD: head/lang/micropython/Makefile 498476 2019-04-09 14:04:49Z sunpoet $

PORTNAME=	micropython
PORTVERSION=	1.10
CATEGORIES=	lang python
MASTER_SITES=	https://github.com/micropython/micropython/releases/download/v${PORTVERSION}/

MAINTAINER=	ports@FreeBSD.org
COMMENT=	Implementation of the Python language for microcontrollers

LICENSE=	MIT
LICENSE_FILE=	${WRKDIR}/${DISTNAME}/LICENSE

BUILD_DEPENDS=	bash:shells/bash
RUN_DEPENDS=	bash:shells/bash
LIB_DEPENDS=	libffi.so:devel/libffi

USES=		gmake pkgconfig python:build
USE_GCC=	yes

PLIST_FILES=	bin/micropython

WRKSRC_SUBDIR=	ports/unix
MAKE_ARGS+=	V=1
ALL_TARGET=	axtls micropython

.include <bsd.port.mk>
