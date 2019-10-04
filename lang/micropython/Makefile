# $FreeBSD: head/lang/micropython/Makefile 498476 2019-04-09 14:04:49Z sunpoet $

PORTNAME=	micropython
PORTVERSION=	1.11
CATEGORIES=	lang python
MASTER_SITES=	http://micropython.org/resources/source/

MAINTAINER=	ports@FreeBSD.org
COMMENT=	Implementation of the Python language for microcontrollers

LICENSE=	MIT
LICENSE_FILE=	${WRKDIR}/${DISTNAME}/LICENSE

BUILD_DEPENDS=	bash:shells/bash \
		gcc9:lang/gcc9
RUN_DEPENDS=	bash:shells/bash \
		gcc9:lang/gcc9
LIB_DEPENDS=	libffi.so:devel/libffi

USES=		gmake pkgconfig python:build tar:xz

PLIST_FILES=	bin/micropython

WRKSRC_SUBDIR=	ports/unix
MAKE_ARGS=	V=1 PYTHON=${PYTHON_CMD} CC=gcc9
ALL_TARGET=	axtls micropython

.include <bsd.port.mk>
