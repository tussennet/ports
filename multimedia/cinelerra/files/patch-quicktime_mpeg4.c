--- quicktime/mpeg4.c.orig	2015-08-13 16:04:04.000000000 +0200
+++ quicktime/mpeg4.c	2016-03-02 18:30:14.006312000 +0100
@@ -629,7 +629,6 @@
 			if(!ffmpeg_initialized)
 			{
 				ffmpeg_initialized = 1;
-  				avcodec_init();
 				avcodec_register_all();
 			}
 
@@ -674,7 +673,7 @@
 #if LIBAVCODEC_VERSION_INT < ((52<<16)+(0<<8)+0)
 			context->error_resilience = FF_ER_CAREFUL;
 #else
-			context->error_recognition = FF_ER_CAREFUL;
+			context->err_recognition = AV_EF_CRCCHECK;
 #endif
 			context->error_concealment = 3;
 			context->frame_skip_cmp = FF_CMP_DCTMAX;
@@ -699,7 +698,6 @@
         	context->profile= FF_PROFILE_UNKNOWN;
 			context->rc_buffer_aggressivity = 1.0;
         	context->level= FF_LEVEL_UNKNOWN;
-			context->flags |= CODEC_FLAG_H263P_UMV;
 			context->flags |= CODEC_FLAG_AC_PRED;
 
 // All the forbidden settings can be extracted from libavcodec/mpegvideo.c of ffmpeg...
@@ -717,10 +715,8 @@
 				(codec->ffmpeg_id == CODEC_ID_MPEG4 ||
 			         codec->ffmpeg_id == CODEC_ID_MPEG1VIDEO ||
 			         codec->ffmpeg_id == CODEC_ID_MPEG2VIDEO ||
-			         codec->ffmpeg_id == CODEC_ID_H263P || 
-			         codec->ffmpeg_id == CODEC_FLAG_H263P_SLICE_STRUCT))
+			         codec->ffmpeg_id == CODEC_ID_H263P))
 			{
-				avcodec_thread_init(context, file->cpus);
 				context->thread_count = file->cpus;
 			}
 
