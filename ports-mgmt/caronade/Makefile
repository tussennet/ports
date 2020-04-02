# $FreeBSD$

PORTNAME=	caronade
DISTVERSIONPREFIX=v
DISTVERSION=	0.3.99.20200402
CATEGORIES=	ports-mgmt

MAINTAINER=	decke@FreeBSD.org
COMMENT=	Small and light tool to help with FreeBSD Ports CI

LICENSE=	BSD2CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	sassc:textproc/sassc

USES=		go:modules,no_targets
USE_RC_SUBR=	caronade

USE_GITHUB=	yes
GH_ACCOUNT=	decke
GH_TAGNAME=	e03349f

GH_TUPLE=	NYTimes:gziphandler:v1.1.1:nytimes_gziphandler/vendor/github.com/NYTimes/gziphandler \
		go-yaml:yaml:v2.2.8:go_yaml_yaml/vendor/gopkg.in/yaml.v2

.include <bsd.port.mk>
