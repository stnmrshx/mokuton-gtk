
install:
	cd pango && go install -x
	cd glib && go install -x
	cd gdk && go install -x
	cd gdkpixbuf && go install -x
	cd gtk && go install -x
