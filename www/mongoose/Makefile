# Created by: Bernhard Froehlich <decke@FreeBSD.org>
# $FreeBSD$

PORTNAME=	mongoose
PORTVERSION=	5.2
CATEGORIES=	www

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Easy to use, powerful, embeddable web server

LICENSE=	GPLv2

USE_GITHUB=	yes
GH_ACCOUNT=	cesanta
GH_COMMIT=	7010add


USES=		gmake perl5
USE_OPENSSL=	yes
USE_PERL5=	build

CFLAGS+=	-DMONGOOSE_USE_IPV6 -DMONGOOSE_USE_SSL -DPF_INET=AF_INET
LDFLAGS+=	-pthread -lssl

MAKE_ENV+=	CFLAGS_EXTRA="${CFLAGS} ${LDFLAGS}"

WRKSRC_SUBDIR=	examples

PLIST_FILES=	bin/mongoose

do-install:
	${MKDIR} ${STAGEDIR}${PREFIX}/bin
	${INSTALL_PROGRAM} ${WRKSRC}/server ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

.include <bsd.port.mk>
