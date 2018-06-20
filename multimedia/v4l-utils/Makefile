# Created by: Juergen Lock <nox@FreeBSD.org>
# $FreeBSD: head/multimedia/v4l-utils/Makefile 424411 2016-10-21 12:51:40Z mat $

PORTNAME?=	v4l-utils
PORTREVISION=	0

COMMENT=	Video4Linux utilities

MASTERDIR=	${.CURDIR}/../../multimedia/libv4l
DESCR=		${.CURDIR}/pkg-descr
PLIST=		${.CURDIR}/pkg-plist
LIBV4L_SLAVE=   utils

INSTALL_WRKSRC=	${WRKSRC}/utils

# no need for udev rules
post-install:
	@${RM} -r ${STAGEDIR}/lib/udev

.include "${MASTERDIR}/Makefile"
