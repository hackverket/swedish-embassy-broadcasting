<img src="http://gfx.bloggar.aftonbladet-cdn.se/wp-content/blogs.dir/428/files/2014/09/48211100457346.png">

# THE SWEDISH EMBASSY BROADCASTING SYSTEM

This was written by highly intoxicated swedes at SHA2017 in Holland in order to drive the radio broadcast from the Swedish Embassy camp. For maximal headache, every path is hard coded. The system supports youtube, soundcloud (and every source youtube-dl supports), file upload, instant wav playback for sound effects and a text to speech readback.

## Running it

The project relies on
 - MOTU ULTRALITE AVB
 - ffmpeg
 - mpd
 - ffmpeg-normalize
 - youtube-dl
 - Amazon AWS Polly
 - Video.JS with hls-contrib addition for web playback

Since proper devops was conducted, the file to run the software is named "test.go", to start the webstream & point to point MPEG-TS stream you have to also run "webstream.sh"

Last but not least:
NÄR MAN FESTAR FESTAR MAN OCH DÅ FESTAR MAN REJÄLT

## License

Lol, do whatever you want with this crap
