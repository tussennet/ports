--- quicktime/qtffmpeg.c.orig	2015-08-13 16:04:04.000000000 +0200
+++ quicktime/qtffmpeg.c	2016-03-02 18:35:13.672424000 +0100
@@ -54,7 +54,6 @@
 	if(!ffmpeg_initialized)
 	{
 		ffmpeg_initialized = 1;
-  		avcodec_init();
 		avcodec_register_all();
 	}
 
@@ -90,10 +89,8 @@
 				(ffmpeg_id == CODEC_ID_MPEG4 ||
 			         ffmpeg_id == CODEC_ID_MPEG1VIDEO ||
 			         ffmpeg_id == CODEC_ID_MPEG2VIDEO ||
-			         ffmpeg_id == CODEC_ID_H263P || 
-			         ffmpeg_id == CODEC_FLAG_H263P_SLICE_STRUCT))
+			         ffmpeg_id == CODEC_ID_H263P))
 		{
-			avcodec_thread_init(context, cpus);
 			context->thread_count = cpus;
 		}
 		if(avcodec_open(context, 
@@ -183,7 +180,7 @@
  
 	if(!result) 
 	{ 
-
+		AVPacket pkt;
 
 // No way to determine if there was an error based on nonzero status.
 // Need to test row pointers to determine if an error occurred.
@@ -191,12 +188,13 @@
 			ffmpeg->decoder_context[current_field]->skip_frame = AVDISCARD_NONREF /* AVDISCARD_BIDIR */;
 		else
 			ffmpeg->decoder_context[current_field]->skip_frame = AVDISCARD_DEFAULT;
-		result = avcodec_decode_video(ffmpeg->decoder_context[current_field], 
+		av_init_packet( &pkt );
+		pkt.data = ffmpeg->work_buffer;
+		pkt.size = bytes + header_bytes;
+		result = avcodec_decode_video2(ffmpeg->decoder_context[current_field],
 			&ffmpeg->picture[current_field], 
 			&got_picture, 
-			ffmpeg->work_buffer, 
-			bytes + header_bytes);
-
+			&pkt);
 
 
 		if(ffmpeg->picture[current_field].data[0])
