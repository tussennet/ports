# Created by: Brad Davis <brd@FreeBSD.org>
# $FreeBSD: head/benchmarks/hey/Makefile 463776 2018-03-07 08:06:24Z yuri $

PORTNAME=	hey
PORTVERSION=	0.1.1
DISTVERSIONPREFIX=	v
CATEGORIES=	benchmarks

MAINTAINER=	brd@FreeBSD.org
COMMENT=	Tiny program that sends some load to a web application

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	rakyll

GH_TUPLE=	golang:net:04a2e54:net/vendor/golang.org/x/net \
		golang:text:v0.3.0:text/vendor/golang.org/x/text

PLIST_FILES=	bin/hey

post-patch:
	${MKDIR} ${WRKSRC}/vendor/github.com/rakyll/hey
	${LN} -s ${WRKSRC}/requester ${WRKSRC}/vendor/github.com/rakyll/hey/requester

.include <bsd.port.mk>
