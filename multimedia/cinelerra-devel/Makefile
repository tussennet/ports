# New ports collection makefile for:	cinelerra-devel
# Date created:		2010-04-16
# Whom:			Ivan Klymenko <fidaj@ukr.net>
#
# $FreeBSD: ports/multimedia/cinelerra-devel/Makefile,v 1.2 2010/03/28 06:32:21 fidaj Exp $
#

PORTNAME=	cinelerra-devel
PORTVERSION=	2.1
CATEGORIES=	multimedia
MASTER_SITES=	http://pazzle.otdux.com.ua/

MAINTAINER=	fidaj@ukr.net
COMMENT=	Cinelerra is a prosumer non-linear video editing system

BUILD_DEPENDS=	${LOCALBASE}/include/linux/videodev.h:${PORTSDIR}/multimedia/v4l_compat
LIB_DEPENDS=	tiff.4:${PORTSDIR}/graphics/tiff \
		faad.2:${PORTSDIR}/audio/faad \
		faac.0:${PORTSDIR}/audio/faac \
		fftw3:${PORTSDIR}/math/fftw3 \
		lavjpeg-2.0:${PORTSDIR}/multimedia/mjpegtools \
		jpeg:${PORTSDIR}/graphics/jpeg \
		a52.0:${PORTSDIR}/audio/liba52 \
		png.6:${PORTSDIR}/graphics/png \
		sndfile.1:${PORTSDIR}/audio/libsndfile \
		mp3lame.0:${PORTSDIR}/audio/lame \
		x264.123:${PORTSDIR}/multimedia/x264 \
		theora.0:${PORTSDIR}/multimedia/libtheora \
		vorbis.4:${PORTSDIR}/audio/libvorbis \
		avcodec.1:${PORTSDIR}/multimedia/ffmpeg \
		dv.4:${PORTSDIR}/multimedia/libdv \
		uuid.1:${PORTSDIR}/misc/e2fsprogs-libuuid \
		IlmImf.6:${PORTSDIR}/graphics/OpenEXR

USE_AUTOTOOLS=	aclocal:env automake:env autoconf:env libtool:env
GNU_CONFIGURE=	yes
USE_BZIP2=	yes
USE_GETTEXT=	yes
USE_GMAKE=	yes
USE_GNOME=	esound
USE_LDCONFIG=	yes
USE_XORG=	x11 xxf86vm xv
USE_GL=		gl glu
CONFIGURE_ARGS=	--disable-alsa \
		--disable-alsatest \
		--enable-opengl \
		--with-external-ffmpeg
CONFIGURE_ENV+=	CFLAGS="-g -D_THREAD_SAFE ${CFLAGS} -I${LOCALBASE}/include" \
		CPPFLAGS="-g -D_THREAD_SAFE ${CPPFLAGS} -I${LOCALBASE}/include" \
		CXXFLAGS="-g -D_THREAD_SAFE ${CXXFLAGS} -I${LOCALBASE}/include" \
		LDFLAGS="${LDFLAGS} -L${LOCALBASE}/lib" \
		LIBS="${LIBS} -lintl -lGL -pthread"
MAKE_JOBS_SAFE=	yes

pre-configure:
	${FIND} ${WRKSRC} -name "Makefile*" | ${XARGS} ${REINPLACE_CMD} -e \
	's|-ldl||g'
	${FIND} ${WRKSRC} -name "*.m4" | ${XARGS} ${REINPLACE_CMD} -e \
	's|-ldl||g'
	cd ${WRKSRC} && ${SH} ./autogen.sh
post-configure:
	${FIND} ${WRKSRC} -name "Makefile*" | ${XARGS} ${REINPLACE_CMD} -e \
	's|-lc_r|${PTHREAD_LIBS}|g ; \
	s|-pthread|${PTHREAD_LIBS}|g ; \
	s|-lpthread|${PTHREAD_LIBS}|g'

.include <bsd.port.mk>
