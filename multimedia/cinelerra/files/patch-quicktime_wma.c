--- quicktime/wma.c.orig	2015-08-13 16:04:04.000000000 +0200
+++ quicktime/wma.c	2016-03-02 18:38:00.050087000 +0100
@@ -67,7 +67,6 @@
 		if(!ffmpeg_initialized)
 		{
 			ffmpeg_initialized = 1;
-			avcodec_init();
 			avcodec_register_all();
 		}
 
@@ -194,12 +193,16 @@
 			codec->packet_buffer,
 			chunk_size);
 #else
+#define	AVCODEC_MAX_AUDIO_FRAME_SIZE   192000
 		bytes_decoded = AVCODEC_MAX_AUDIO_FRAME_SIZE;
-		result = avcodec_decode_audio2(codec->decoder_context,
+		AVPacket pkt;
+		av_init_packet( &pkt );
+		pkt.data = codec->packet_buffer;
+		pkt.size = chunk_size;
+		result = avcodec_decode_audio3(codec->decoder_context,
 			(int16_t*)(codec->work_buffer + codec->output_size * sample_size),
 			&bytes_decoded,
-			codec->packet_buffer,
-			chunk_size);
+			&pkt);
 #endif
 
 		pthread_mutex_unlock(&ffmpeg_lock);
