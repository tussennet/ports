# $FreeBSD$

PORTNAME=	gerbera
PORTVERSION=	1.6.1
DISTVERSIONPREFIX=	v
CATEGORIES=	multimedia

MAINTAINER=	decke@FreeBSD.org
COMMENT=	UPnP media server for home network

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/LICENSE.md

LIB_DEPENDS=	libupnp.so:devel/upnp \
		libexpat.so:textproc/expat2 \
		libsqlite3.so:databases/sqlite3 \
		libinotify.so:devel/libinotify \
		libduktape.so:lang/duktape-lib \
		libcurl.so:ftp/curl \
		libtag.so:audio/taglib \
		libexif.so:graphics/libexif \
		libebml.so:textproc/libebml \
		libmatroska.so:multimedia/libmatroska

USES=		cmake:noninja pkgconfig

USE_GITHUB=	yes
CMAKE_ARGS+=	-DWITH_MYSQL=0 -DWITH_LASTFM=0 -DWITH_SYSTEMD=0

.include <bsd.port.mk>
