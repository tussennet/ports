# Created by: Bernhard Froehlich <decke@bluelife.at>
# $FreeBSD: head/www/helma/Makefile 412351 2016-04-01 14:33:55Z mat $

PORTNAME=	helma
PORTVERSION=	1.7.4
CATEGORIES=	www java

MAINTAINER=	ports@FreeBSD.org
COMMENT=	Web application framework for fast and efficient scripting

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/license.txt

DATADIR=	${JAVASHAREDIR}/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	helma-org
GH_TAGNAME=	v${PORTVERSION}

USE_RC_SUBR=	helma
USE_JAVA=	yes
JAVA_VERSION=	1.8+
USE_ANT=	yes
ALL_TARGET=	clean jar
NO_ARCH=	yes

DBDIR?=		/var/db/${PORTNAME}
LOGDIR?=	/var/log/${PORTNAME}
HELMA_USER=	${WWWOWN}
HELMA_GROUP=	${WWWGRP}
HELMA_CONFIG=	${WRKSRC}/scripts/helma.conf

PLIST_SUB=	DBDIR="${DBDIR}" \
		LOGDIR="${LOGDIR}" \
		HELMA_USER="${HELMA_USER}" \
		HELMA_GROUP="${HELMA_GROUP}"

OPTIONS_DEFINE=	MYSQL
OPTIONS_SUB=	yes
OPTIONS_DESC=	Install MySQL Connector for Java

MYSQL_RUN_DEPENDS=	${JAVALIBDIR}/mysql-connector-java.jar:databases/mysql-connector-java

post-patch:
	@${REINPLACE_CMD} -e 's|JAVA_HOME=/usr/lib/j2sdk1.5-sun|JAVA_HOME=${JAVA_HOME}|' \
		${HELMA_CONFIG}
	@${REINPLACE_CMD} -e 's|HELMA_INSTALL=/usr/local/helma/helma-1.6.x|HELMA_INSTALL=${DATADIR}|' \
		${HELMA_CONFIG}
	@${REINPLACE_CMD} -e 's|HELMA_HOME=\$$HELMA_INSTALL|HELMA_HOME=${WWWDIR}|' \
		${HELMA_CONFIG}
	@${REINPLACE_CMD} -e 's|HELMA_USER=helma|HELMA_USER=${HELMA_USER}|' \
		${HELMA_CONFIG}
	@${REINPLACE_CMD} -e 's|HELMA_LOG=\$$HELMA_HOME/log/helma-out.log|HELMA_LOG=${LOGDIR}/helma-out.log|' \
		${HELMA_CONFIG}
	@${ECHO_CMD} "HELMA_CONFIG=${ETCDIR}" >> \
		${HELMA_CONFIG}

	@${ECHO_CMD} "" >> \
		${WRKSRC}/build/main/server.properties
	@${ECHO_CMD} "# custom directories" >> \
		${WRKSRC}/build/main/server.properties
	@${ECHO_CMD} "hopHome = ${WWWDIR}" >> \
		${WRKSRC}/build/main/server.properties
	@${ECHO_CMD} "appHome = ${WWWDIR}/apps" >> \
		${WRKSRC}/build/main/server.properties
	@${ECHO_CMD} "appsPropFile = ${ETCDIR}/apps.properties" >> \
		${WRKSRC}/build/main/server.properties
	@${ECHO_CMD} "dbPropFile = ${ETCDIR}/db.properties" >> \
		${WRKSRC}/build/main/server.properties
	@${ECHO_CMD} "dbHome = ${DBDIR}" >> \
		${WRKSRC}/build/main/server.properties
	@${REINPLACE_CMD} -e 's|# logDir = console|logDir = ${LOGDIR}|' \
		${WRKSRC}/build/main/server.properties

do-install:
	${INSTALL_SCRIPT} ${FILESDIR}/helma.sh.in ${STAGEDIR}${PREFIX}/bin/helma

	${MKDIR} ${STAGEDIR}${WWWDIR} && \
	(cd ${WRKSRC} && ${COPYTREE_SHARE} "apps modules" ${STAGEDIR}${WWWDIR} \
		"! -name *.html ! -name helma-*.jar ! -name .jsdoc")

	${MKDIR} ${STAGEDIR}${DATADIR} && \
	(cd ${WRKSRC} && ${COPYTREE_SHARE} "lib launcher.jar" ${STAGEDIR}${DATADIR} \
		"! -name *.html ! -name helma-*.jar ! -name .jsdoc")

	${MKDIR} ${STAGEDIR}${LOGDIR}

	${MKDIR} ${STAGEDIR}${DBDIR}

	${MKDIR} ${STAGEDIR}${ETCDIR}; \
	for f in build/main/apps.properties build/main/db.properties build/main/server.properties scripts/helma.conf; do \
		${INSTALL_DATA} ${WRKSRC}/$$f ${STAGEDIR}${ETCDIR}/$${f##*/}.sample; \
	done

do-install-MYSQL-on:
	${LN} -sf ${JAVALIBDIR}/mysql-connector-java.jar ${STAGEDIR}${DATADIR}/lib/ext/mysql-connector-java.jar

.include <bsd.port.mk>
