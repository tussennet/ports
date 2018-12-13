# $FreeBSD$

PORTNAME=	smtp-proxy
PORTVERSION=	1.0.1
CATEGORIES=	mail
MASTER_SITES=	https://code.bluelife.at/decke/smtp-proxy/archive/
DISTFILES=	v${PORTVERSION}.tar.gz

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Simple Go SMTP relay/proxy server

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go
USE_RC_SUBR=	smtpproxy

USE_GITHUB=	nodefault
GH_TUPLE=	chrj:smtpd:v0.1.1:smtpd/vendor/github.com/chrj/smtpd \
		vharitonsky:iniflags:a33cd0b5f3de:iniflags/vendor/github.com/vharitonsky/iniflags

WRKSRC=		${WRKDIR}/${PORTNAME}
MAKE_ENV+=	GOFLAGS=-mod=vendor

PLIST_FILES=	bin/${PORTNAME} \
		"@sample ${ETCDIR}/smtp-proxy.ini.sample"

post-install:
		${MKDIR} ${STAGEDIR}${ETCDIR}
		${INSTALL_DATA} ${WRKSRC}/smtp-proxy.ini ${STAGEDIR}${ETCDIR}/smtp-proxy.ini.sample

.include <bsd.port.mk>
