toflv
=====

toflv converts videos into FLV files using ffmpeg.
**Therefore, ffmpeg needs to be installed on the machine**.

How to use it
-------------

    toflv <files...>
    
As simple as that. Remember that you need to have ffmpeg installed.

Why I wrote that
----------------

If you record your screen using e.g. OBS, you create extremely large files.
If you want to send them via e.g. Telegram, you need to convert those files into a smaller file.
You can use ffmpeg to do this, but I have to google the command every time, which is annoying.
So I wrote this simple go file, which does the job.

Just use a bash script...
-------------------------

A bash script would be an alternative, but difficult to use with Windows.