package gdk

import "C"
import "unsafe"

func (v *Window) GetNativeWindowID() int32 {
	return int32(C.gdk_x11_drawable_get_xid((*C.GdkDrawable)(unsafe.Pointer(v.GWindow))))
}
