# $FreeBSD: ports/multimedia/cinelerra-devel/Makefile,v 1.2 2010/03/28 06:32:21 fidaj Exp $

PORTNAME=	cinelerra
PORTVERSION=	2.3
CATEGORIES=	multimedia
MASTER_SITES=	http://www.cinelerra-cv.org/releases/
DISTNAME=	CinelerraCV-${DISTVERSION}

MAINTAINER=	fidaj@ukr.net
COMMENT=	Cinelerra is a prosumer non-linear video editing system

LICENSE=	GPLv2

BUILD_DEPENDS=	${LOCALBASE}/include/linux/videodev.h:${PORTSDIR}/multimedia/v4l_compat \
		nasm:${PORTSDIR}/devel/nasm

LIB_DEPENDS=	libtiff.so:${PORTSDIR}/graphics/tiff \
		libfaad.so:${PORTSDIR}/audio/faad \
		libfaac.so:${PORTSDIR}/audio/faac \
		libfftw3.so:${PORTSDIR}/math/fftw3 \
		liblavjpeg.so:${PORTSDIR}/multimedia/mjpegtools \
		libjpeg.so:${PORTSDIR}/graphics/jpeg \
		liba52.so:${PORTSDIR}/audio/liba52 \
		libpng.so:${PORTSDIR}/graphics/png \
		libsndfile.so:${PORTSDIR}/audio/libsndfile \
		libmp3lame.so:${PORTSDIR}/audio/lame \
		libx264.so:${PORTSDIR}/multimedia/x264 \
		libtheora.so:${PORTSDIR}/multimedia/libtheora \
		libvorbis.so:${PORTSDIR}/audio/libvorbis \
		libavcodec.so:${PORTSDIR}/multimedia/ffmpeg \
		libdv.so:${PORTSDIR}/multimedia/libdv \
		libuuid.so:${PORTSDIR}/misc/e2fsprogs-libuuid \
		libIlmImf.so:${PORTSDIR}/graphics/OpenEXR

USES=		autoreconf gettext gmake tar:xz libtool
GNU_CONFIGURE=	yes
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

#post-configure:
#	${FIND} ${WRKSRC} -name "Makefile*" | ${XARGS} ${REINPLACE_CMD} -e \
#	's|-lc_r|${PTHREAD_LIBS}|g ; \
#	s|-pthread|${PTHREAD_LIBS}|g ; \
#	s|-lpthread|${PTHREAD_LIBS}|g'

.include <bsd.port.mk>
