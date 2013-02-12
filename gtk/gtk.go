package gtk

import "C"
import (
	"fmt"
	"github.com/stnmrshx/mokuton-gtk/gdk"
	"github.com/stnmrshx/mokuton-gtk/gdkpixbuf"
	"github.com/stnmrshx/mokuton-gtk/glib"
	"github.com/stnmrshx/mokuton-gtk/pango"
	"log"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

func gint(v int) C.gint {
	return C.gint(v)
}

func guint(v uint) C.guint {
	return C.guint(v)
}

func guint32(v uint32) C.guint32 {
	return C.guint32(v)
}

func glong(v int32) C.glong {
	return C.glong(v)
}

func gdouble(v float64) C.gdouble {
	return C.gdouble(v)
}

func gsize_t(v C.size_t) C.gint {
	return C.gint(v)
}

func gstring(s *C.char) *C.gchar {
	return C.toGstr(s)
}

func cstring(s *C.gchar) *C.char {
	return C.toCstr(s)
}

func gostring(s *C.gchar) string {
	return C.GoString(cstring(s))
}

func gslist(l *glib.SList) *C.GSList {
	if l == nil {
		return nil
	}
	return C.to_gslist(unsafe.Pointer(l.ToSList()))
}

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

func cfree(s *C.char) { C.freeCstr(s) }

func WINDOW(p *Window) *C.GtkWindow {
	return C.toGWindow(p.GWidget)
}

func DIALOG(p *Dialog) *C.GtkDialog {
	return C.toGDialog(p.GWidget)
}

func ABOUT_DIALOG(p *AboutDialog) *C.GtkAboutDialog {
	return C.toGAboutDialog(p.GWidget)
}

func CONTAINER(p *Container) *C.GtkContainer {
	return C.toGContainer(p.GWidget)
}

func FILE_CHOOSER(p *Widget) *C.GtkFileChooser {
	return C.toGFileChooser(p.GWidget)
}

func FONT_SELECTION_DIALOG(p *FontSelectionDialog) *C.GtkFontSelectionDialog {
	return C.toGFontSelectionDialog(p.GWidget)
}

func LABEL(p *Label) *C.GtkLabel {
	return C.toGLabel(p.GWidget)
}

func BUTTON(p *Button) *C.GtkButton {
	return C.toGButton(p.GWidget)
}

func SPIN_BUTTON(p *SpinButton) *C.GtkSpinButton {
	return C.toGSpinButton(p.GWidget)
}

func RADIO_BUTTON(p *RadioButton) *C.GtkRadioButton {
	return C.toGRadioButton(p.GWidget)
}

func FONT_BUTTON(p *FontButton) *C.GtkFontButton {
	return C.toGFontButton(p.GWidget)
}

func LINK_BUTTON(p *LinkButton) *C.GtkLinkButton {
	return C.toGLinkButton(p.GWidget)
}

func COMBO_BOX(p *ComboBox) *C.GtkComboBox {
	return C.toGComboBox(p.GWidget)
}

func COMBO_BOX_ENTRY(p *ComboBoxEntry) *C.GtkComboBoxEntry {
	return C.toGComboBoxEntry(p.GWidget)
}

func MESSAGE_DIALOG(p *MessageDialog) *C.GtkMessageDialog {
	return C.toGMessageDialog(p.GWidget)
}

func COMBO_BOX_TEXT(p *ComboBoxText) *C.GtkComboBoxText {
	return C.toGComboBoxText(p.GWidget)
}

func ACCESSIBLE(p *Accessible) *C.GtkAccessible {
	return C.toGAccessible(unsafe.Pointer(p.Object))
}

func BIN(p *Bin) *C.GtkBin {
	return C.toGBin(p.GWidget)
}

func STATUSBAR(p *Statusbar) *C.GtkStatusbar {
	return C.toGStatusbar(p.GWidget)
}

func INFO_BAR(p *InfoBar) *C.GtkInfoBar {
	return C.toGInfoBar(p.GWidget)
}

func FRAME(p *Frame) *C.GtkFrame {
	return C.toGFrame(p.GWidget)
}

func BOX(p *Box) *C.GtkBox {
	return C.toGBox(p.GWidget)
}

func PANED(p *Paned) *C.GtkPaned {
	return C.toGPaned(p.GWidget)
}

func TOGGLE_BUTTON(p *ToggleButton) *C.GtkToggleButton {
	return C.toGToggleButton(p.GWidget)
}

func ACCEL_LABEL(p *AccelLabel) *C.GtkAccelLabel {
	return C.toGAccelLabel(p.GWidget)
}

func ENTRY(p *Entry) *C.GtkEntry {
	return C.toGEntry(p.GWidget)
}

func ADJUSTMENT(p *Adjustment) *C.GtkAdjustment {
	return p.GAdjustment
}

func TEXT_VIEW(p *TextView) *C.GtkTextView {
	return C.toGTextView(p.GWidget)
}

func TEXT_BUFFER(p unsafe.Pointer) *C.GtkTextBuffer {
	return C.toGTextBuffer(p)
}

func MENU(p *Menu) *C.GtkMenu {
	return C.toGMenu(p.GWidget)
}

func MENU_BAR(p *MenuBar) *C.GtkMenuBar {
	return C.toGMenuBar(p.GWidget)
}

func MENU_SHELL(p *Menu) *C.GtkMenuShell {
	return C.toGMenuShell(p.GWidget)
}

func MENU_BAR_SHELL(p *MenuBar) *C.GtkMenuShell {
	return C.toGMenuShell(p.GWidget)
}

func MENU_ITEM(p *MenuItem) *C.GtkMenuItem {
	return C.toGMenuItem(p.GWidget)
}

func ITEM(p *Item) *C.GtkItem {
	return C.toGItem(p.GWidget)
}

func TOOLBAR(p *Toolbar) *C.GtkToolbar {
	return C.toGToolbar(p.GWidget)
}

func TOOL_ITEM(p *ToolItem) *C.GtkToolItem {
	return C.toGToolItem(p.GWidget)
}

func SEPARATOR_TOOL_ITEM(p *SeparatorToolItem) *C.GtkSeparatorToolItem {
	return C.toGSeparatorToolItem(p.GWidget)
}

func TOOL_BUTTON(p *ToolButton) *C.GtkToolButton {
	return C.toGToolButton(p.GWidget)
}

func TOOL_PALETTE(p *ToolPalette) *C.GtkToolPalette {
	return C.toGToolPalette(p.GWidget)
}

func TOOL_ITEM_GROUP(p *ToolItemGroup) *C.GtkToolItemGroup {
	return C.toGToolItemGroup(p.GWidget)
}

func MENU_TOOL_BUTTON(p *MenuToolButton) *C.GtkMenuToolButton {
	return C.toGMenuToolButton(p.GWidget)
}

func TOGGLE_TOOL_BUTTON(p *ToggleToolButton) *C.GtkToggleToolButton {
	return C.toGToggleToolButton(p.GWidget)
}

func SCROLLED_WINDOW(p *ScrolledWindow) *C.GtkScrolledWindow {
	return C.toGScrolledWindow(p.GWidget)
}

func VIEWPORT(p *Viewport) *C.GtkViewport {
	return C.toGViewport(p.GWidget)
}

func WIDGET(p *Widget) *C.GtkWidget {
	return C.toGWidget(unsafe.Pointer(p.GWidget))
}

func TREE_VIEW(p *TreeView) *C.GtkTreeView {
	return C.toGTreeView(p.GWidget)
}

func ICON_VIEW(p *IconView) *C.GtkIconView {
	return C.toGIconView(p.GWidget)
}

func CELL_RENDERER_TEXT(p *CellRendererText) *C.GtkCellRendererText {
	return C.toGCellRendererText(p.GCellRenderer)
}

func CELL_RENDERER_TOGGLE(p *CellRendererToggle) *C.GtkCellRendererToggle {
	return C.toGCellRendererToggle(p.GCellRenderer)
}

func SCALE(p *Scale) *C.GtkScale {
	return C.toGScale(p.GWidget)
}

func RANGE(p *Range) *C.GtkRange {
	return C.toGRange(p.GWidget)
}

func IMAGE(p *Image) *C.GtkImage {
	return C.toGImage(p.GWidget)
}

func NOTEBOOK(p *Notebook) *C.GtkNotebook {
	return C.toGNotebook(p.GWidget)
}

func TABLE(p *Table) *C.GtkTable {
	return C.toGTable(p.GWidget)
}

func DRAWING_AREA(p *DrawingArea) *C.GtkDrawingArea {
	return C.toGDrawingArea(p.GWidget)
}

func ASSISTANT(p *Assistant) *C.GtkAssistant {
	return C.toGAssistant(p.GWidget)
}

func EXPANDER(p *Expander) *C.GtkExpander {
	return C.toGExpander(p.GWidget)
}

func ALIGNMENT(p *Alignment) *C.GtkAlignment {
	return C.toGAlignment(p.GWidget)
}

func PROGRESS_BAR(p *ProgressBar) *C.GtkProgressBar {
	return C.toGProgressBar(p.GWidget)
}

func FIXED(p *Fixed) *C.GtkFixed {
	return C.toGFixed(p.GWidget)
}

func CHECK_MENU_ITEM(p *CheckMenuItem) *C.GtkCheckMenuItem {
	return C.toGCheckMenuItem(p.GWidget)
}

func RADIO_MENU_ITEM(p *RadioMenuItem) *C.GtkRadioMenuItem {
	return C.toGRadioMenuItem(p.GWidget)
}

func panic_if_version_older(major int, minor int, micro int, function string) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) == 0 {
		log.Panicf("%s is not provided on your GTK, version %d.%d is required\n", function, major, minor)
	}
}

func panic_if_version_older_auto(major, minor, micro int) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) != 0 {
		return
	}
	formatStr := "%s is not provided on your GTK, version %d.%d is required\n"
	if pc, _, _, ok := runtime.Caller(1); ok {
		log.Panicf(formatStr, runtime.FuncForPC(pc).Name(), major, minor)
	} else {
		log.Panicf("GTK version %d.%d is required (unknown caller, see stack)\n", major, minor)
	}
}

func deprecated_since(major int, minor int, micro int, function string) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) != 0 {
		log.Printf("\nWarning: %s is deprecated since gtk %d.%d\n", function, major, minor)
	}
}

func variadicButtonsToArrays(buttons []interface{}) ([]string, []ResponseType) {
	if len(buttons)%2 != 0 {
		argumentPanic("variadic parameter must be even couples of string-ResponseType (button label - button response)")
	}
	text := make([]string, len(buttons)/2)
	res := make([]ResponseType, len(buttons)/2)
	for i := 0; i < len(text); i++ {
		btext, ok := buttons[2*i].(string)
		if !ok {
			argumentPanic("button text must be a string")
		}
		bresponse, ok := buttons[2*i+1].(ResponseType)
		if !ok {
			argumentPanic("button response must be an ResponseType")
		}
		text[i] = btext
		res[i] = bresponse
	}
	return text, res
}

func argumentPanic(message string) {
	if pc, _, _, ok := runtime.Caller(2); ok {
		log.Panicf("Arguments error: %s : %s\n",
			runtime.FuncForPC(pc).Name(), message)
	} else {
		log.Panicln("Arguments error: (unknown caller, see stack):", message)
	}
}

// Main Loop and Events
func SetLocale() {
	C.gtk_set_locale()
}

func Init(args *[]string) {
	if args != nil {
		var argc C.int = C.int(len(*args))
		cargs := make([]*C.char, argc)
		for i, arg := range *args {
			cargs[i] = C.CString(arg)
		}
		C._gtk_init(unsafe.Pointer(&argc), unsafe.Pointer(&cargs))
		goargs := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			goargs[i] = C.GoString(cargs[i])
		}
		for i := 0; i < int(argc); i++ {
			cfree(cargs[i])
		}
		*args = goargs
	} else {
		C._gtk_init(nil, nil)
	}
}

func Main() {
	C.gtk_main()
}

func MainQuit() {
	C.gtk_main_quit()
}

func MainIteration() bool {
	return gobool(C.gtk_main_iteration())
}

func MainIterationDo(blocking bool) bool {
	return gobool(C.gtk_main_iteration_do(gbool(blocking)))
}

func EventsPending() bool {
	return gobool(C.gtk_events_pending())
}

// gtkAccelGroup : represents a group of keyboard accelerators, typically attached to a toplevel GtkWindow with add_accel_group()
type AccelGroup struct {
	GAccelGroup *C.GtkAccelGroup
}

func NewAccelGroup() *AccelGroup {
	return &AccelGroup{C.gtk_accel_group_new()}
}

func AcceleratorGetDefaultModMask() uint {
	return uint(C.gtk_accelerator_get_default_mod_mask())
}

// gtkClipboard
type Clipboard struct {
	GClipboard *C.GtkClipboard
}

func NewClipboardGetForDisplay(display *gdk.Display, selection gdk.Atom) *Clipboard {
	var cdisplay unsafe.Pointer
	if display != nil {
		cdisplay = display.GDisplay
	}
	return &Clipboard{C._gtk_clipboard_get_for_display(cdisplay, unsafe.Pointer(selection))}
}

func (v *Clipboard) Clear() {
	C.gtk_clipboard_clear(v.GClipboard)
}

func (v *Clipboard) SetText(text string) {
	//ptr := C.toCstrV(unsafe.Pointer(&([]byte(text))[0])) // ARRRGGGHHHH i can't understand unsafe pointer >,< Need Fix
	p := C.CString(text)
	defer cfree(p)
	C.gtk_clipboard_set_text(v.GClipboard, gstring(p), gint(-1))
}

func (v *Clipboard) SetImage(pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_clipboard_set_image(v.GClipboard, pixbuf.GPixbuf)
}

func (v *Clipboard) Store() {
	C.gtk_clipboard_store(v.GClipboard)
}

func (v *Clipboard) WaitForText() string {
	return gostring(C.gtk_clipboard_wait_for_text(v.GClipboard))
}

func (v *Clipboard) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(C.toGObject(unsafe.Pointer(v.GClipboard)))).Connect(s, f, datas...)
}

// Drag and Drop
type DestDefaults int

const (
	DEST_DEFAULT_MOTION    DestDefaults = 1 << 0
	DEST_DEFAULT_HIGHLIGHT DestDefaults = 1 << 1
	DEST_DEFAULT_DROP      DestDefaults = 1 << 2
	DEST_DEFAULT_ALL       DestDefaults = 0x07
)

type TargetEntry struct {
	Target string
	Flags  uint
	Info   uint
}

func (v *Widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions gdk.DragAction) {
	ctargets := make([]C.GtkTargetEntry, len(targets))
	for i, target := range targets {
		ptr := C.CString(target.Target)
		defer cfree(ptr)
		ctargets[i].target = gstring(ptr)
		ctargets[i].flags = guint(target.Flags)
		ctargets[i].info = guint(target.Info)
	}
	C.gtk_drag_dest_set(v.GWidget, C.GtkDestDefaults(flags), &ctargets[0], gint(len(targets)), C.GdkDragAction(actions))
}

func (v *Widget) DragSourceSet(start_button_mask gdk.ModifierType, targets []TargetEntry, actions gdk.DragAction) {
	ctargets := make([]C.GtkTargetEntry, len(targets))
	for i, target := range targets {
		ptr := C.CString(target.Target)
		defer cfree(ptr)
		ctargets[i].target = gstring(ptr)
		ctargets[i].flags = guint(target.Flags)
		ctargets[i].info = guint(target.Info)
	}
	C.gtk_drag_source_set(v.GWidget, C.GdkModifierType(start_button_mask), &ctargets[0], gint(len(targets)), C.GdkDragAction(actions))
}

func (v *Widget) DragFinish(context *gdk.DragContext, success bool, del bool, time uint32) {
	C._gtk_drag_finish(unsafe.Pointer(context.DragContext), gbool(success), gbool(del), guint32(time))
}

func (v *Widget) DragDestAddUriTargets() {
	C.gtk_drag_dest_add_uri_targets(v.GWidget)
}

// gtkStockItem
const (
	STOCK_ABOUT                         = "gtk-about"
	STOCK_ADD                           = "gtk-add"
	STOCK_APPLY                         = "gtk-apply"
	STOCK_BOLD                          = "gtk-bold"
	STOCK_CANCEL                        = "gtk-cancel"
	STOCK_CAPS_LOCK_WARNING             = "gtk-caps-lock-warning"
	STOCK_CDROM                         = "gtk-cdrom"
	STOCK_CLEAR                         = "gtk-clear"
	STOCK_CLOSE                         = "gtk-close"
	STOCK_COLOR_PICKER                  = "gtk-color-picker"
	STOCK_CONVERT                       = "gtk-convert"
	STOCK_CONNECT                       = "gtk-connect"
	STOCK_COPY                          = "gtk-copy"
	STOCK_CUT                           = "gtk-cut"
	STOCK_DELETE                        = "gtk-delete"
	STOCK_DIALOG_AUTHENTICATION         = "gtk-dialog-authentication"
	STOCK_DIALOG_INFO                   = "gtk-dialog-info"
	STOCK_DIALOG_WARNING                = "gtk-dialog-warning"
	STOCK_DIALOG_ERROR                  = "gtk-dialog-error"
	STOCK_DIALOG_QUESTION               = "gtk-dialog-question"
	STOCK_DIRECTORY                     = "gtk-directory"
	STOCK_DISCARD                       = "gtk-discard"
	STOCK_DISCONNECT                    = "gtk-disconnect"
	STOCK_DND                           = "gtk-dnd"
	STOCK_DND_MULTIPLE                  = "gtk-dnd-multiple"
	STOCK_EDIT                          = "gtk-edit"
	STOCK_EXECUTE                       = "gtk-execute"
	STOCK_FILE                          = "gtk-file"
	STOCK_FIND                          = "gtk-find"
	STOCK_FIND_AND_REPLACE              = "gtk-find-and-replace"
	STOCK_FLOPPY                        = "gtk-floppy"
	STOCK_FULLSCREEN                    = "gtk-fullscreen"
	STOCK_GOTO_BOTTOM                   = "gtk-goto-bottom"
	STOCK_GOTO_FIRST                    = "gtk-goto-first"
	STOCK_GOTO_LAST                     = "gtk-goto-last"
	STOCK_GOTO_TOP                      = "gtk-goto-top"
	STOCK_GO_BACK                       = "gtk-go-back"
	STOCK_GO_DOWN                       = "gtk-go-down"
	STOCK_GO_FORWARD                    = "gtk-go-forward"
	STOCK_GO_UP                         = "gtk-go-up"
	STOCK_HARDDISK                      = "gtk-harddisk"
	STOCK_HELP                          = "gtk-help"
	STOCK_HOME                          = "gtk-home"
	STOCK_INDEX                         = "gtk-index"
	STOCK_INDENT                        = "gtk-indent"
	STOCK_INFO                          = "gtk-info"
	STOCK_UNINDENT                      = "gtk-unindent"
	STOCK_ITALIC                        = "gtk-italic"
	STOCK_JUMP_TO                       = "gtk-jump-to"
	STOCK_JUSTIFY_CENTER                = "gtk-justify-center"
	STOCK_JUSTIFY_FILL                  = "gtk-justify-fill"
	STOCK_JUSTIFY_LEFT                  = "gtk-justify-left"
	STOCK_JUSTIFY_RIGHT                 = "gtk-justify-right"
	STOCK_LEAVE_FULLSCREEN              = "gtk-leave-fullscreen"
	STOCK_MISSING_IMAGE                 = "gtk-missing-image"
	STOCK_MEDIA_FORWARD                 = "gtk-media-forward"
	STOCK_MEDIA_NEXT                    = "gtk-media-next"
	STOCK_MEDIA_PAUSE                   = "gtk-media-pause"
	STOCK_MEDIA_PLAY                    = "gtk-media-play"
	STOCK_MEDIA_PREVIOUS                = "gtk-media-previous"
	STOCK_MEDIA_RECORD                  = "gtk-media-record"
	STOCK_MEDIA_REWIND                  = "gtk-media-rewind"
	STOCK_MEDIA_STOP                    = "gtk-media-stop"
	STOCK_NETWORK                       = "gtk-network"
	STOCK_NEW                           = "gtk-new"
	STOCK_NO                            = "gtk-no"
	STOCK_OK                            = "gtk-ok"
	STOCK_OPEN                          = "gtk-open"
	STOCK_ORIENTATION_PORTRAIT          = "gtk-orientation-portrait"
	STOCK_ORIENTATION_LANDSCAPE         = "gtk-orientation-landscape"
	STOCK_ORIENTATION_REVERSE_LANDSCAPE = "gtk-orientation-reverse-landscape"
	STOCK_ORIENTATION_REVERSE_PORTRAIT  = "gtk-orientation-reverse-portrait"
	STOCK_PAGE_SETUP                    = "gtk-page-setup"
	STOCK_PASTE                         = "gtk-paste"
	STOCK_PREFERENCES                   = "gtk-preferences"
	STOCK_PRINT                         = "gtk-print"
	STOCK_PRINT_ERROR                   = "gtk-print-error"
	STOCK_PRINT_PAUSED                  = "gtk-print-paused"
	STOCK_PRINT_PREVIEW                 = "gtk-print-preview"
	STOCK_PRINT_REPORT                  = "gtk-print-report"
	STOCK_PRINT_WARNING                 = "gtk-print-warning"
	STOCK_PROPERTIES                    = "gtk-properties"
	STOCK_QUIT                          = "gtk-quit"
	STOCK_REDO                          = "gtk-redo"
	STOCK_REFRESH                       = "gtk-refresh"
	STOCK_REMOVE                        = "gtk-remove"
	STOCK_REVERT_TO_SAVED               = "gtk-revert-to-saved"
	STOCK_SAVE                          = "gtk-save"
	STOCK_SAVE_AS                       = "gtk-save-as"
	STOCK_SELECT_ALL                    = "gtk-select-all"
	STOCK_SELECT_COLOR                  = "gtk-select-color"
	STOCK_SELECT_FONT                   = "gtk-select-font"
	STOCK_SORT_ASCENDING                = "gtk-sort-ascending"
	STOCK_SORT_DESCENDING               = "gtk-sort-descending"
	STOCK_SPELL_CHECK                   = "gtk-spell-check"
	STOCK_STOP                          = "gtk-stop"
	STOCK_STRIKETHROUGH                 = "gtk-strikethrough"
	STOCK_UNDELETE                      = "gtk-undelete"
	STOCK_UNDERLINE                     = "gtk-underline"
	STOCK_UNDO                          = "gtk-undo"
	STOCK_YES                           = "gtk-yes"
	STOCK_ZOOM_100                      = "gtk-zoom-100"
	STOCK_ZOOM_FIT                      = "gtk-zoom-fit"
	STOCK_ZOOM_IN                       = "gtk-zoom-in"
	STOCK_ZOOM_OUT                      = "gtk-zoom-out"
)

type StockItem struct {
	GStockItem *C.GtkStockItem
}

func (v *StockItem) Add(nitems uint) {
	C.gtk_stock_add(v.GStockItem, guint(nitems))
}
func (v *StockItem) AddStatic(nitems uint) {
	C.gtk_stock_add_static(v.GStockItem, guint(nitems))
}
func StockLookup(stock_id string, item *StockItem) bool {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	return gobool(C.gtk_stock_lookup(gstring(ptr), item.GStockItem))
}

func StockListIDs() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_stock_list_ids()))
}

// gtkSettings
type Settings struct {
	GSettings *C.GtkSettings
}

func (s *Settings) SetStringProperty(name string, v_string string, origin string) {
	ptrn := C.CString(name)
	defer cfree(ptrn)
	ptrv := C.CString(v_string)
	defer cfree(ptrv)
	prts := C.CString(origin)
	defer cfree(prts)
	C.gtk_settings_set_string_property(
		s.GSettings, 
		gstring(ptrn), 
		gstring(ptrv), 
		gstring(prts)
	)
}

func (s *Settings) SetLongProperty(name string, v_long int32, origin string) {
	ptrn := C.CString(name)
	defer cfree(ptrn)
	prts := C.CString(origin)
	defer cfree(prts)
	C.gtk_settings_set_long_property(
		s.GSettings, 
		gstring(ptrn), 
		glong(v_long), 
		gstring(prts)
	)
}

func (s *Settings) SetDoubleProperty(name string, v_double float64, origin string) {
	ptrn := C.CString(name)
	defer cfree(ptrn)
	prts := C.CString(origin)
	defer cfree(prts)
	C.gtk_settings_set_double_property(
		s.GSettings, 
		gstring(ptrn), 
		gdouble(v_double), 
		gstring(prts)
	)
}

type SelectionData struct {
	GSelectionData unsafe.Pointer
}

func NewSelectionDataFromNative(l unsafe.Pointer) *SelectionData {
	return &SelectionData{l}
}

func (v *SelectionData) GetLength() int {
	return int(C._gtk_selection_data_get_length(v.GSelectionData))
}

func (v *SelectionData) GetData() unsafe.Pointer {
	return unsafe.Pointer(C._gtk_selection_data_get_data(v.GSelectionData))
}

func (v *SelectionData) GetText() string {
	return C.GoString(C._gtk_selection_data_get_text(v.GSelectionData))
}

// gtkDialog
type DialogFlags int

const (
	DIALOG_MODAL               DialogFlags = 1 << 0
	DIALOG_DESTROY_WITH_PARENT DialogFlags = 1 << 1
	DIALOG_NO_SEPARATOR        DialogFlags = 1 << 2
)

type ResponseType int

const (
	RESPONSE_NONE         ResponseType = -1
	RESPONSE_REJECT       ResponseType = -2
	RESPONSE_ACCEPT       ResponseType = -3
	RESPONSE_DELETE_EVENT ResponseType = -4
	RESPONSE_OK           ResponseType = -5
	RESPONSE_CANCEL       ResponseType = -6
	RESPONSE_CLOSE        ResponseType = -7
	RESPONSE_YES          ResponseType = -8
	RESPONSE_NO           ResponseType = -9
	RESPONSE_APPLY        ResponseType = -10
	RESPONSE_HELP         ResponseType = -11
)

type Dialog struct {
	Window
	wfr *Widget
}

func NewDialog() *Dialog {
	return &Dialog{
		Window{
			Bin{
				Container{
					Widget{
						C.gtk_dialog_new()
					}
				}
			}
		}, nil
	}
}

func (v *Dialog) Run() ResponseType {
	return ResponseType(C.gtk_dialog_run(DIALOG(v)))
}

func (v *Dialog) Response(response interface{}, datas ...interface{}) {
	v.Connect("response", response, datas...)
}

func (v *Dialog) AddButton(button_text string, response_id ResponseType) *Button {
	ptr := C.CString(button_text)
	defer cfree(ptr)
	return &Button{Bin{Container{Widget{
		C.gtk_dialog_add_button(DIALOG(v), gstring(ptr), gint(int(response_id)))}}}}
}

func (v *Dialog) SetHasSeparator(f bool) {
	deprecated_since(2, 22, 0, "gtk_dialog_set_has_separator()")
	C.gtk_dialog_set_has_separator(DIALOG(v), gbool(f))
}

func (v *Dialog) SetDefaultResponse(id int) {
	C.gtk_dialog_set_default_response(DIALOG(v), gint(id))
}

func (v *Dialog) GetResponseForWidget(w *Widget) int {
	return int(C.gtk_dialog_get_response_for_widget(DIALOG(v), w.GWidget))
}

func (v *Dialog) GetWidgetForResponse(id int) *Widget {
	panic_if_version_older(2, 20, 0, "gtk_dialog_get_widget_for_response()")
	w := C._gtk_dialog_get_widget_for_response(DIALOG(v), gint(id))
	if v.wfr == nil {
		v.wfr = &Widget{w}
	} else {
		v.wfr.GWidget = w
	}
	return v.wfr
}

func (v *Dialog) GetVBox() *VBox {
	return &VBox{Box{Container{Widget{C._gtk_dialog_get_vbox(v.GWidget)}}}}
}

// gtkMessageDialog
type MessageType int

const (
	MESSAGE_INFO     = 0
	MESSAGE_WARNING  = 1
	MESSAGE_QUESTION = 2
	MESSAGE_ERROR    = 3
	MESSAGE_OTHER    = 4
)

type ButtonsType int

const (
	BUTTONS_NONE      = 0
	BUTTONS_OK        = 1
	BUTTONS_CLOSE     = 2
	BUTTONS_CANCEL    = 3
	BUTTONS_YES_NO    = 4
	BUTTONS_OK_CANCEL = 5
)

type MessageDialog struct {
	Dialog
}

func NewMessageDialog(parent *Window, flag DialogFlags, t MessageType, buttons ButtonsType, format string, args ...interface{}) *MessageDialog {
	ptr := C.CString(strings.Replace(fmt.Sprintf(format, args...), "%", "%%", -1))
	defer cfree(ptr)
	return &MessageDialog{Dialog{Window{Bin{Container{Widget{
		C._gtk_message_dialog_new(
			ToNative(parent),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(buttons),
			gstring(ptr))}}}}, nil}}
}

func NewMessageDialogWithMarkup(parent *Window, flag DialogFlags, t MessageType, buttons ButtonsType, format string, args ...interface{}) *MessageDialog {
	r := &MessageDialog{Dialog{Window{Bin{Container{Widget{
		C._gtk_message_dialog_new_with_markup(
			ToNative(parent),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(buttons),
			nil)}}}}, nil}}
	r.SetMarkup(fmt.Sprintf(format, args...))
	return r
}

func (v *MessageDialog) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_message_dialog_set_markup(MESSAGE_DIALOG(v), gstring(ptr))
}

func (v *MessageDialog) SetImage(image IWidget) {
	C.gtk_message_dialog_set_image(MESSAGE_DIALOG(v), ToNative(image))
}

func (v *MessageDialog) GetImage() *Image {
	return &Image{Widget{C.gtk_message_dialog_get_image(MESSAGE_DIALOG(v))}}
}

// gtkWindow
type WindowType int

const (
	WINDOW_TOPLEVEL WindowType = 0
	WINDOW_POPUP    WindowType = 1
)

type WindowPosition int

const (
	WIN_POS_NONE             WindowPosition = 0
	WIN_POS_CENTER           WindowPosition = 1
	WIN_POS_MOUSE            WindowPosition = 2
	WIN_POS_CENTER_ALWAYS    WindowPosition = 3
	WIN_POS_CENTER_ON_PARENT WindowPosition = 4
)

type Window struct {
	Bin
}

func NewWindow(t WindowType) *Window {
	return &Window{Bin{Container{Widget{
		C.gtk_window_new(C.GtkWindowType(t))}}}}
}

func (v *Window) SetTitle(title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_window_set_title(
		WINDOW(v), 
		gstring(ptr)
	)
}

func (v *Window) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(
		WINDOW(v), 
		gbool(resizable)
	)
}

func (v *Window) GetResizable() bool {
	return gobool(C.gtk_window_get_resizable(WINDOW(v)))
}

func (v *Window) AddAccelGroup(group *AccelGroup) {
	C.gtk_window_add_accel_group(
		WINDOW(v), 
		group.GAccelGroup
	)
}

func (v *Window) SetModal(modal bool) {
	C.gtk_window_set_modal(
		WINDOW(v), 
		gbool(modal)
	)
}

func (v *Window) SetDefaultSize(width int, height int) {
	C.gtk_window_set_default_size(
		WINDOW(v), 
		gint(width), 
		gint(height)
	)
}

func (v *Window) SetPosition(position WindowPosition) {
	C.gtk_window_set_position(
		WINDOW(v), 
		C.GtkWindowPosition(position)
	)
}

func (v *Window) SetTransientFor(parent *Window) {
	C.gtk_window_set_transient_for(
		WINDOW(v), 
		WINDOW(parent)
	)
}

func (v *Window) SetDestroyWithParent(setting bool) {
	C.gtk_window_set_destroy_with_parent(
		WINDOW(v), 
		gbool(setting)
	)
}

func (v *Window) SetDefault(w *Widget) {
	C.gtk_window_set_default(
		WINDOW(v), 
		w.GWidget
	)
}

func (v *Window) Present() {
	C.gtk_window_present(WINDOW(v))
}

func (v *Window) Stick() {
	C.gtk_window_stick(WINDOW(v))
}

func (v *Window) Unstick() {
	C.gtk_window_unstick(WINDOW(v))
}

func (v *Window) Iconify() {
	C.gtk_window_iconify(WINDOW(v))
}

func (v *Window) Deiconify() {
	C.gtk_window_deiconify(WINDOW(v))
}

func (v *Window) Maximize() {
	C.gtk_window_maximize(WINDOW(v))
}
func (v *Window) Unmaximize() {
	C.gtk_window_unmaximize(WINDOW(v))
}

func (v *Window) Fullscreen() {
	C.gtk_window_fullscreen(WINDOW(v))
}

func (v *Window) Unfullscreen() {
	C.gtk_window_unfullscreen(WINDOW(v))
}

func (v *Window) SetKeepAbove(setting bool) {
	C.gtk_window_set_keep_above(
		WINDOW(v), 
		gbool(setting)
	)
}

func (v *Window) SetKeepBelow(setting bool) {
	C.gtk_window_set_keep_below(
		WINDOW(v), 
		gbool(setting)
	)
}

func (v *Window) SetDecorated(setting bool) {
	C.gtk_window_set_decorated(
		WINDOW(v), 
		gbool(setting)
	)
}

func (v *Window) SetDeletable(setting bool) {
	C.gtk_window_set_deletable(
		WINDOW(v), 
		gbool(setting)
	)
}

func (v *Window) SetTypeHint(hint gdk.WindowTypeHint) {
	C.gtk_window_set_type_hint(
		WINDOW(v), 
		C.GdkWindowTypeHint(hint)
	)
}

func (v *Window) SetAcceptFocus(setting bool) {
	C.gtk_window_set_accept_focus(
		WINDOW(v), 
		gbool(setting)
	)
}

func (v *Window) GetDefaultSize(width *int, height *int) {
	var cwidth, cheight C.gint
	C.gtk_window_get_default_size(
		WINDOW(v), 
		&cwidth, 
		&cheight
	)
	*width = int(cwidth)
	*height = int(cheight)
}

func (v *Window) GetDestroyWithParent() bool {
	return gobool(C.gtk_window_get_destroy_with_parent(WINDOW(v)))
}

func (v *Window) GetIconName() string {
	return gostring(C.gtk_window_get_icon_name(WINDOW(v)))
}

func (v *Window) GetModal() bool {
	return gobool(C.gtk_window_get_modal(WINDOW(v)))
}

func (v *Window) GetPosition(root_x *int, root_y *int) {
	var croot_x, croot_y C.gint
	C.gtk_window_get_position(
		WINDOW(v), 
		&croot_x, 
		&croot_y
	)
	*root_x = int(croot_x)
	*root_y = int(croot_y)
}

func (v *Window) GetSize(width *int, height *int) {
	var cwidth, cheight C.gint
	C.gtk_window_get_size(
		WINDOW(v), 
		&cwidth, 
		&cheight
	)
	*width = int(cwidth)
	*height = int(cheight)
}

func (v *Window) GetTitle() string {
	return gostring(C.gtk_window_get_title(WINDOW(v)))
}

func (v *Window) GetTypeHint() gdk.WindowTypeHint {
	return gdk.WindowTypeHint(C.gtk_window_get_type_hint(WINDOW(v)))
}

func (v *Window) GetAcceptFocus() bool {
	return gobool(C.gtk_window_get_accept_focus(WINDOW(v)))
}

func (v *Window) Move(x int, y int) {
	C.gtk_window_move(
		WINDOW(v), 
		gint(x), 
		gint(y)
	)
}

func (v *Window) Resize(width int, height int) {
	C.gtk_window_resize(
		WINDOW(v), 
		gint(width), 
		gint(height)
	)
}

func (v *Window) XID() int32 {
	return gdk.WindowFromUnsafe(unsafe.Pointer(v.GWidget.window)).GetNativeWindowID()
}

func (v *Window) SetIconFromFile(file string) {
	ptr := C.CString(file)
	defer cfree(ptr)
	C.gtk_window_set_icon_from_file(
		WINDOW(v), 
		gstring(ptr), 
		nil
	)
}

func (v *Window) SetIconName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_window_set_icon_name(
		WINDOW(v), 
		gstring(ptr)
	)
}

// gtkAboutDialog
type AboutDialog struct {
	Dialog
}

func NewAboutDialog() *AboutDialog {
	return &AboutDialog{Dialog{Window{Bin{Container{Widget{C.gtk_about_dialog_new()}}}}, nil}}
}

func (v *AboutDialog) GetProgramName() string {
	return gostring(C.gtk_about_dialog_get_program_name(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetProgramName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_about_dialog_set_program_name(
		ABOUT_DIALOG(v), 
		gstring(ptr)
	)
}

func (v *AboutDialog) GetVersion() string {
	return gostring(C.gtk_about_dialog_get_version(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetVersion(version string) {
	ptr := C.CString(version)
	defer cfree(ptr)
	C.gtk_about_dialog_set_version(
		ABOUT_DIALOG(v), 
		gstring(ptr)
	)
}

func (v *AboutDialog) GetCopyright() string {
	return gostring(C.gtk_about_dialog_get_copyright(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetCopyright(copyright string) {
	ptr := C.CString(copyright)
	defer cfree(ptr)
	C.gtk_about_dialog_set_copyright(
		ABOUT_DIALOG(v), 
		gstring(ptr)
	)
}

func (v *AboutDialog) GetComments() string {
	return gostring(C.gtk_about_dialog_get_comments(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetComments(comments string) {
	ptr := C.CString(comments)
	defer cfree(ptr)
	C.gtk_about_dialog_set_comments(
		ABOUT_DIALOG(v), 
		gstring(ptr)
	)
}

func (v *AboutDialog) GetLicense() string {
	return gostring(C.gtk_about_dialog_get_license(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetLicense(license string) {
	ptr := C.CString(license)
	defer cfree(ptr)
	C.gtk_about_dialog_set_license(
		ABOUT_DIALOG(v), 
		gstring(ptr)
	)
}

func (v *AboutDialog) GetWrapLicense() bool {
	return gobool(C.gtk_about_dialog_get_wrap_license(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetWrapLicense(wrap_license bool) {
	C.gtk_about_dialog_set_wrap_license(
		ABOUT_DIALOG(v), 
		gbool(wrap_license)
	)
}

func (v *AboutDialog) GetWebsite() string {
	return gostring(C.gtk_about_dialog_get_website(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetWebsite(website string) {
	ptr := C.CString(website)
	defer cfree(ptr)
	C.gtk_about_dialog_set_website(
		ABOUT_DIALOG(v), 
		gstring(ptr)
	)
}

func (v *AboutDialog) GetWebsiteLabel() string {
	return gostring(C.gtk_about_dialog_get_website_label(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetWebsiteLabel(website_label string) {
	ptr := C.CString(website_label)
	defer cfree(ptr)
	C.gtk_about_dialog_set_website_label(
		ABOUT_DIALOG(v), 
		gstring(ptr)
	)
}

func (v *AboutDialog) GetAuthors() []string {
	var authors []string
	cauthors := C.gtk_about_dialog_get_authors(ABOUT_DIALOG(v))
	for {
		authors = append(authors, gostring(*cauthors))
		cauthors = C.nextGstr(cauthors)
		if *cauthors == nil {
			break
		}
	}
	return authors
}

func (v *AboutDialog) SetAuthors(authors []string) {
	cauthors := C.make_strings(C.int(len(authors) + 1))
	for i, author := range authors {
		ptr := C.CString(author)
		defer cfree(ptr)
		C.set_string(cauthors, C.int(i), gstring(ptr))
	}
	C.set_string(cauthors, C.int(len(authors)), nil)
	C.gtk_about_dialog_set_authors(
		ABOUT_DIALOG(v), 
		cauthors
	)
	C.destroy_strings(cauthors)
}

func (v *AboutDialog) GetArtists() []string {
	var artists []string
	cartists := C.gtk_about_dialog_get_artists(ABOUT_DIALOG(v))
	for {
		artists = append(artists, gostring(*cartists))
		cartists = C.nextGstr(cartists)
		if *cartists == nil {
			break
		}
	}
	return artists
}

func (v *AboutDialog) SetArtists(artists []string) {
	cartists := C.make_strings(C.int(len(artists)))
	for i, author := range artists {
		ptr := C.CString(author)
		defer cfree(ptr)
		C.set_string(cartists, C.int(i), gstring(ptr))
	}
	C.gtk_about_dialog_set_artists(
		ABOUT_DIALOG(v), 
		cartists
	)
	C.destroy_strings(cartists)
}

func (v *AboutDialog) GetDocumenters() []string {
	var documenters []string
	cdocumenters := C.gtk_about_dialog_get_documenters(ABOUT_DIALOG(v))
	for {
		documenters = append(documenters, gostring(*cdocumenters))
		cdocumenters = C.nextGstr(cdocumenters)
		if *cdocumenters == nil {
			break
		}
	}
	return documenters
}

func (v *AboutDialog) SetDocumenters(documenters []string) {
	cdocumenters := C.make_strings(C.int(len(documenters)))
	for i, author := range documenters {
		ptr := C.CString(author)
		defer cfree(ptr)
		C.set_string(cdocumenters, C.int(i), gstring(ptr))
	}
	C.gtk_about_dialog_set_documenters(
		ABOUT_DIALOG(v), 
		cdocumenters
	)
	C.destroy_strings(cdocumenters)
}

func (v *AboutDialog) GetTranslatorCredits() string {
	return gostring(C.gtk_about_dialog_get_translator_credits(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetTranslatorCredits(translator_credits string) {
	ptr := C.CString(translator_credits)
	defer cfree(ptr)
	C.gtk_about_dialog_set_translator_credits(
		ABOUT_DIALOG(v), 
		gstring(ptr)
	)
}

func (v *AboutDialog) GetLogo() *gdkpixbuf.Pixbuf {
	return &gdkpixbuf.Pixbuf{C.gtk_about_dialog_get_logo(ABOUT_DIALOG(v))}
}

func (v *AboutDialog) SetLogo(logo *gdkpixbuf.Pixbuf) {
	C.gtk_about_dialog_set_logo(
		ABOUT_DIALOG(v), 
		logo.GPixbuf
	)
}

func (v *AboutDialog) GetLogoIconName() string {
	return gostring(C.gtk_about_dialog_get_logo_icon_name(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetLogoIconName(icon_name string) {
	ptr := C.CString(icon_name)
	defer cfree(ptr)
	C.gtk_about_dialog_set_logo_icon_name(
		ABOUT_DIALOG(v), 
		gstring(ptr)
	)
}

// gtkAssistant
type AssistantPageType int

const (
	ASSISTANT_PAGE_CONTENT  AssistantPageType = 0
	ASSISTANT_PAGE_INTRO    AssistantPageType = 1
	ASSISTANT_PAGE_CONFIRM  AssistantPageType = 2
	ASSISTANT_PAGE_SUMMARY  AssistantPageType = 3
	ASSISTANT_PAGE_PROGRESS AssistantPageType = 4
)

type Assistant struct {
	Widget
}

func NewAssistant() *Assistant {
	return &Assistant{Widget{C.gtk_assistant_new()}}
}

func (v *Assistant) GetCurrentPage() int {
	return int(C.gtk_assistant_get_current_page(ASSISTANT(v)))
}

func (v *Assistant) SetCurrentPage(page_num int) {
	C.gtk_assistant_set_current_page(
		ASSISTANT(v), 
		gint(page_num)
	)
}

func (v *Assistant) GetNPages() int {
	return int(C.gtk_assistant_get_n_pages(ASSISTANT(v)))
}

func (v *Assistant) GetNthPage(page_num int) *Widget {
	return &Widget{
		C.gtk_assistant_get_nth_page(
			ASSISTANT(v), 
			gint(page_num)
		)
	}
}

func (v *Assistant) PrependPage(page IWidget) int {
	return int(C.gtk_assistant_prepend_page(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) AppendPage(page IWidget) int {
	return int(C.gtk_assistant_prepend_page(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) InsertPage(page IWidget, position int) int {
	return int(C.gtk_assistant_insert_page(ASSISTANT(v), ToNative(page), gint(position)))
}

func (v *Assistant) SetPageType(page IWidget, t AssistantPageType) {
	C.gtk_assistant_set_page_type(
		ASSISTANT(v), 
		ToNative(page), 
		C.GtkAssistantPageType(t)
	)
}

func (v *Assistant) GetPageType(page IWidget) AssistantPageType {
	return AssistantPageType(C.gtk_assistant_get_page_type(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) SetPageTitle(page IWidget, title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_assistant_set_page_title(
		ASSISTANT(v), 
		ToNative(page), 
		gstring(ptr)
	)
}

func (v *Assistant) GetPageTitle(page IWidget) string {
	return gostring(C.gtk_assistant_get_page_title(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) SetPageHeaderImage(page IWidget, pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_assistant_set_page_header_image(
		ASSISTANT(v), 
		ToNative(page), 
		pixbuf.GPixbuf
	)
}

func (v *Assistant) GetPageHeaderImage(page IWidget) *gdkpixbuf.Pixbuf {
	return &gdkpixbuf.Pixbuf{
		C.gtk_assistant_get_page_header_image(
			ASSISTANT(v), 
			ToNative(page)
		)
	}
}

func (v *Assistant) SetPageSideImage(page IWidget, pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_assistant_set_page_side_image(
		ASSISTANT(v), 
		ToNative(page), 
		pixbuf.GPixbuf
	)
}

func (v *Assistant) GetPageSideImage(page IWidget) *gdkpixbuf.Pixbuf {
	return &gdkpixbuf.Pixbuf{
		C.gtk_assistant_get_page_side_image(ASSISTANT(v), ToNative(page))
	}
}

func (v *Assistant) SetPageComplete(page IWidget, complete bool) {
	C.gtk_assistant_set_page_complete(
		ASSISTANT(v), 
		ToNative(page), 
		gbool(complete)
	)
}

func (v *Assistant) GetPageComplete(page IWidget) bool {
	return gobool(C.gtk_assistant_get_page_complete(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) AddActionWidget(child IWidget) {
	C.gtk_assistant_add_action_widget(
		ASSISTANT(v), 
		ToNative(child)
	)
}

func (v *Assistant) RemoveActionWidget(child IWidget) {
	C.gtk_assistant_remove_action_widget(
		ASSISTANT(v), 
		ToNative(child)
	)
}

func (v *Assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(ASSISTANT(v))
}

// gtkAccelLabel
type AccelLabel struct {
	Widget
}

func NewAccelLabel(label string) *AccelLabel {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &AccelLabel{Widget{C.gtk_accel_label_new(gstring(ptr))}}
}

func (v *AccelLabel) GetAccelWidget() Widget {
	return Widget{C.gtk_accel_label_get_accel_widget(ACCEL_LABEL(v))}
}

func (v *AccelLabel) SetAccelWidget(w IWidget) {
	C.gtk_accel_label_set_accel_widget(
		ACCEL_LABEL(v), 
		ToNative(w)
	)
}

func (v *AccelLabel) GetAccelWidth() uint {
	return uint(C.gtk_accel_label_get_accel_width(ACCEL_LABEL(v)))
}

func (v *AccelLabel) Refetch() bool {
	return gobool(C.gtk_accel_label_refetch(ACCEL_LABEL(v)))
}

// gtkImage
type IconSize int

const (
	ICON_SIZE_INVALID       IconSize = 0
	ICON_SIZE_MENU          IconSize = 1
	ICON_SIZE_SMALL_TOOLBAR IconSize = 2
	ICON_SIZE_LARGE_TOOLBAR IconSize = 3
	ICON_SIZE_BUTTON        IconSize = 4
	ICON_SIZE_DND           IconSize = 5
	ICON_SIZE_DIALOG        IconSize = 6
)

type Image struct {
	Widget
}

func NewImage() *Image {
	return &Image{Widget{C.gtk_image_new()}}
}

func NewImageFromFile(filename string) *Image {
	ptr := C.CString(filename)
	defer cfree(ptr)
	return &Image{Widget{C.gtk_image_new_from_file(gstring(ptr))}}
}

func NewImageFromPixbuf(pixbuf gdkpixbuf.Pixbuf) *Image {
	return &Image{Widget{C.gtk_image_new_from_pixbuf(pixbuf.GPixbuf)}}
}

func NewImageFromStock(stock_id string, size IconSize) *Image {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	return &Image{Widget{C.gtk_image_new_from_stock(gstring(ptr), C.GtkIconSize(size))}}
}

func (v *Image) GetPixbuf() *gdkpixbuf.Pixbuf {
	return &gdkpixbuf.Pixbuf{
		C.gtk_image_get_pixbuf(IMAGE(v))}
}

func (v *Image) SetFromFile(filename string) {
	ptr := C.CString(filename)
	defer cfree(ptr)
	C.gtk_image_set_from_file(
		IMAGE(v), 
		gstring(ptr)
	)
}

func (v *Image) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_image_set_from_pixbuf(
		IMAGE(v), 
		pixbuf.GPixbuf
	)
}

func (v *Image) SetFromStock(stock_id string, size IconSize) {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	C.gtk_image_set_from_stock(
		IMAGE(v), 
		gstring(ptr), 
		C.GtkIconSize(size)
	)
}

func (v *Image) Clear() {
	C.gtk_image_clear(IMAGE(v))
}

// gtkLabel
type Justification int

const (
	JUSTIFY_LEFT   Justification = 0
	JUSTIFY_RIGHT  Justification = 1
	JUSTIFY_CENTER Justification = 2
	JUSTIFY_FILL   Justification = 3
)

type ILabel interface {
	IWidget
	isILabel()
	GetLabel() string
	SetLabel(label string)
}

type Label struct {
	Widget
}

func (Label) isILabel() {
	//EMPTY WILL RETURN
}

func NewLabel(label string) *Label {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &Label{Widget{C.gtk_label_new(gstring(ptr))}}
}

func (v *Label) SetText(label string) {
	ptr := C.CString(label)
	defer cfree(ptr)
	C.gtk_label_set_text(
		LABEL(v), 
		gstring(ptr)
	)
}

func (v *Label) SetMnemonicWidget(widget IWidget) {
	C.gtk_label_set_mnemonic_widget(
		LABEL(v), 
		ToNative(widget)
	)
}

func (v *Label) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_label_set_markup(
		LABEL(v), 
		gstring(ptr)
	)
}

func (v *Label) GetMnemonicWidget() *Widget {
	return &Widget{C.gtk_label_get_mnemonic_widget(LABEL(v))}
}

func (v *Label) SetPattern(pattern string) {
	ptr := C.CString(pattern)
	defer cfree(ptr)
	C.gtk_label_set_pattern(
		LABEL(v), 
		gstring(ptr)
	)
}

func (v *Label) SetJustify(jtype Justification) {
	C.gtk_label_set_justify(
		LABEL(v), 
		C.GtkJustification(jtype)
	)
}

func (v *Label) SetEllipsize(ellipsize pango.EllipsizeMode) {
	C.gtk_label_set_ellipsize(
		LABEL(v), 
		C.PangoEllipsizeMode(ellipsize)
	)
}

func (v *Label) SetWidthChars(n_chars int) {
	C.gtk_label_set_width_chars(
		LABEL(v), 
		gint(n_chars)
	)
}

func (v *Label) SetMaxWidthChars(n_chars int) {
	C.gtk_label_set_max_width_chars(
		LABEL(v), 
		gint(n_chars)
	)
}

func (v *Label) SetLineWrap(setting bool) {
	C.gtk_label_set_line_wrap(
		LABEL(v), 
		gbool(setting)
	)
}

func (v *Label) SetUseLineWrapMode(wrap_mode pango.WrapMode) {
	C.gtk_label_set_line_wrap_mode(
		LABEL(v), 
		C.PangoWrapMode(wrap_mode)
	)
}

func (v *Label) GetSelectable() bool {
	return gobool(C.gtk_label_get_selectable(LABEL(v)))
}

func (v *Label) GetText() string {
	return gostring(C.gtk_label_get_text(LABEL(v)))
}

func LabelWithMnemonic(label string) *Label {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &Label{Widget{C.gtk_label_new_with_mnemonic(gstring(ptr))}}
}

func (v *Label) SelectRegion(start_offset int, end_offset int) {
	C.gtk_label_select_region(
		LABEL(v), 
		gint(start_offset), 
		gint(end_offset)
	)
}

func (v *Label) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(
		LABEL(v), 
		gbool(setting)
	)
}

func (v *Label) SetTextWithMnemonic(str string) {
	ptr := C.CString(str)
	defer cfree(ptr)
	C.gtk_label_set_text_with_mnemonic(
		LABEL(v), 
		gstring(ptr)
	)
}

func (v *Label) GetJustify() Justification {
	return Justification(C.gtk_label_get_justify(LABEL(v)))
}

func (v *Label) GetEllipsize() pango.EllipsizeMode {
	return pango.EllipsizeMode(C.gtk_label_get_ellipsize(LABEL(v)))
}

func (v *Label) GetWidthChars() int {
	return int(C.gtk_label_get_width_chars(LABEL(v)))
}

func (v *Label) GetMaxWidthChars() int {
	return int(C.gtk_label_get_max_width_chars(LABEL(v)))
}

func (v *Label) GetLabel() string {
	return gostring(C.gtk_label_get_label(LABEL(v)))
}

func (v *Label) GetLineWrap() bool {
	return gobool(C.gtk_label_get_line_wrap(LABEL(v)))
}

func (v *Label) GetLineWrapMode() pango.WrapMode {
	return pango.WrapMode(C.gtk_label_get_line_wrap_mode(LABEL(v)))
}

func (v *Label) GetSelectionBounds(start *int, end *int) {
	var cstart, cend C.gint
	C.gtk_label_get_selection_bounds(
		LABEL(v), 
		&cstart, 
		&cend
	)
	*start = int(cstart)
	*end = int(cend)
}

func (v *Label) GetUseMarkup() bool {
	return gobool(C.gtk_label_get_use_markup(LABEL(v)))
}

func (v *Label) GetUseUnderline() bool {
	return gobool(C.gtk_label_get_use_underline(LABEL(v)))
}

func (v *Label) GetSingleLineMode() bool {
	return gobool(C.gtk_label_get_single_line_mode(LABEL(v)))
}

func (v *Label) GetAngle() float64 {
	return float64(C.gtk_label_get_angle(LABEL(v)))
}

func (v *Label) SetLabel(label string) {
	ptr := C.CString(label)
	defer cfree(ptr)
	C.gtk_label_set_label(
		LABEL(v), 
		gstring(ptr)
	)
}

func (v *Label) SetUseMarkup(setting bool) {
	C.gtk_label_set_use_markup(
		LABEL(v), 
		gbool(setting)
	)
}

func (v *Label) SetUseUnderline(setting bool) {
	C.gtk_label_set_use_underline(
		LABEL(v), 
		gbool(setting)
	)
}

func (v *Label) SetSingleLineMode(single_line bool) {
	C.gtk_label_set_single_line_mode(
		LABEL(v), 
		gbool(single_line)
	)
}

func (v *Label) SetAngle(angle float64) {
	C.gtk_label_set_angle(
		LABEL(v), 
		gdouble(angle)
	)
}

func (v *Label) GetCurrentUri() string {
	panic_if_version_older(2, 18, 0, "gtk_label_get_current_uri()")
	return gostring(C.gtk_label_get_current_uri(LABEL(v)))
}

func (v *Label) SetTrackVisitedLinks(track_links bool) {
	panic_if_version_older(2, 18, 0, "gtk_label_set_track_visited_links()")
	C.gtk_label_set_track_visited_links(
		LABEL(v), 
		gbool(track_links)
	)
}

func (v *Label) GetTrackVisitedLinks() bool {
	panic_if_version_older(2, 18, 0, "gtk_label_get_track_visited_links()")
	return gobool(C.gtk_label_get_track_visited_links(LABEL(v)))
}

// gtkProgressBar
type ProgressBarOrientation int

const (
	PROGRESS_LEFT_TO_RIGHT ProgressBarOrientation = 0
	PROGRESS_RIGHT_TO_LEFT ProgressBarOrientation = 1
	PROGRESS_BOTTOM_TO_TOP ProgressBarOrientation = 2
	PROGRESS_TOP_TO_BOTTOM ProgressBarOrientation = 3
)

type ProgressBar struct {
	Widget
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{Widget{C.gtk_progress_bar_new()}}
}
func (v *ProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(PROGRESS_BAR(v))
}

func (v *ProgressBar) SetText(show_text string) {
	ptr := C.CString(show_text)
	defer cfree(ptr)
	C.gtk_progress_bar_set_text(
		PROGRESS_BAR(v), 
		gstring(ptr)
	)
}

func (v *ProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(
		PROGRESS_BAR(v), 
		gdouble(fraction)
	)
}

func (v *ProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(
		PROGRESS_BAR(v), 
		gdouble(fraction)
	)
}

func (v *ProgressBar) SetOrientation(i ProgressBarOrientation) {
	C.gtk_progress_bar_set_orientation(
		PROGRESS_BAR(v), 
		C.GtkProgressBarOrientation(i)
	)
}

func (v *ProgressBar) GetText() string {
	return gostring(C.gtk_progress_bar_get_text(PROGRESS_BAR(v)))
}

func (v *ProgressBar) GetFraction() float64 {
	r := C.gtk_progress_bar_get_fraction(PROGRESS_BAR(v))
	return float64(r)
}

func (v *ProgressBar) GetPulseStep() float64 {
	r := C.gtk_progress_bar_get_pulse_step(PROGRESS_BAR(v))
	return float64(r)
}

func (v *ProgressBar) GetOrientation() ProgressBarOrientation {
	return ProgressBarOrientation(C.gtk_progress_bar_get_orientation(PROGRESS_BAR(v)))
}

// gtkStatusbar
type Statusbar struct {
	HBox
}

func NewStatusbar() *Statusbar {
	return &Statusbar{HBox{Box{Container{Widget{C.gtk_statusbar_new()}}}}}
}

func (v *Statusbar) GetContextId(content_description string) uint {
	ptr := C.CString(content_description)
	defer cfree(ptr)
	return uint(C.gtk_statusbar_get_context_id(STATUSBAR(v), gstring(ptr)))
}

func (v *Statusbar) Push(context_id uint, text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_statusbar_push(
		STATUSBAR(v), 
		guint(context_id), 
		gstring(ptr)
	)
}

func (v *Statusbar) Pop(context_id uint) {
	C.gtk_statusbar_pop(
		STATUSBAR(v), 
		guint(context_id)
	)
}

func (v *Statusbar) Remove(context_id uint, message_id uint) {
	C.gtk_statusbar_remove(
		STATUSBAR(v), 
		guint(context_id), 
		guint(message_id)
	)
}

func (v *Statusbar) SetHasResizeGrip(add_tearoffs bool) {
	C.gtk_statusbar_set_has_resize_grip(
		STATUSBAR(v), 
		gbool(add_tearoffs)
	)
}

func (v *Statusbar) GetHasResizeGrip() bool {
	return gobool(C.gtk_statusbar_get_has_resize_grip(STATUSBAR(v)))
}

// gtkInfoBar
type InfoBar struct {
	HBox
}

func NewInfoBar() *InfoBar {
	panic_if_version_older_auto(2, 18, 0)
	return &InfoBar{HBox{Box{Container{Widget{C._gtk_info_bar_new()}}}}}
}

func NewInfoBarWithButtons(buttons ...interface{}) *InfoBar {
	panic_if_version_older_auto(2, 18, 0)
	infobar := NewInfoBar()
	text, res := variadicButtonsToArrays(buttons)
	for i := range text {
		infobar.AddButton(text[i], res[i])
	}
	return infobar
}

func (v *InfoBar) AddActionWidget(child IWidget, responseId ResponseType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_add_action_widget(
		INFO_BAR(v), 
		ToNative(child), 
		gint(int(responseId))
	)
}

func (v *InfoBar) AddButton(buttonText string, responseId ResponseType) *Widget {
	panic_if_version_older_auto(2, 18, 0)
	ptr := C.CString(buttonText)
	defer cfree(ptr)
	return &Widget{C._gtk_info_bar_add_button(INFO_BAR(v), gstring(ptr), gint(int(responseId)))}
}

func (v *InfoBar) AddButtons(buttons ...interface{}) {
	panic_if_version_older_auto(2, 18, 0)
	text, res := variadicButtonsToArrays(buttons)
	for i := range text {
		v.AddButton(text[i], res[i])
	}
}

func (v *InfoBar) SetResponseSensitive(responseId ResponseType, setting bool) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_response_sensitive(
		INFO_BAR(v), 
		gint(int(responseId)), 
		gbool(setting)
	)
}

func (v *InfoBar) SetDefaultResponse(responseId ResponseType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_default_response(
		INFO_BAR(v), 
		gint(int(responseId))
	)
}

func (v *InfoBar) Response(responseId ResponseType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_response(
		INFO_BAR(v), 
		gint(int(responseId))
	)
}

func (v *InfoBar) SetMessageType(messageType MessageType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_message_type(
		INFO_BAR(v), 
		C.GtkMessageType(messageType)
	)
}

func (v *InfoBar) GetMessageType() MessageType {
	panic_if_version_older_auto(2, 18, 0)
	return MessageType(C._gtk_info_bar_get_message_type(INFO_BAR(v)))
}

func (v *InfoBar) GetActionArea() *Widget {
	panic_if_version_older_auto(2, 18, 0)
	return &Widget{C._gtk_info_bar_get_action_area(INFO_BAR(v))}
}

func (v *InfoBar) GetContentArea() *Widget {
	panic_if_version_older_auto(2, 18, 0)
	return &Widget{C._gtk_info_bar_get_content_area(INFO_BAR(v))}
}

// gtkStatusIcon
type StatusIcon struct {
	GStatusIcon *C.GtkStatusIcon
}

func NewStatusIcon() *StatusIcon {
	return &StatusIcon{C.gtk_status_icon_new()}
}

func NewStatusIconFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *StatusIcon {
	return &StatusIcon{C.gtk_status_icon_new_from_pixbuf(pixbuf.GPixbuf)}
}

func NewStatusIconFromFile(filename string) *StatusIcon {
	ptr := C.CString(filename)
	defer cfree(ptr)
	return &StatusIcon{C.gtk_status_icon_new_from_file(gstring(ptr))}
}

func NewStatusIconFromStock(stock_id string) *StatusIcon {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	return &StatusIcon{C.gtk_status_icon_new_from_stock(gstring(ptr))}
}

func NewStatusIconFromIconName(icon_name string) *StatusIcon {
	ptr := C.CString(icon_name)
	defer cfree(ptr)
	return &StatusIcon{C.gtk_status_icon_new_from_icon_name(gstring(ptr))}
}

func (v *StatusIcon) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_status_icon_set_from_pixbuf(
		v.GStatusIcon, 
		pixbuf.GPixbuf
	)
}

func (v *StatusIcon) SetFromFile(filename string) {
	ptr := C.CString(filename)
	defer cfree(ptr)
	C.gtk_status_icon_set_from_file(
		v.GStatusIcon, 
		gstring(ptr)
	)
}

func (v *StatusIcon) SetFromStock(stock_id string) {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	C.gtk_status_icon_set_from_stock(
		v.GStatusIcon, 
		gstring(ptr)
	)
}

func (v *StatusIcon) SetFromIconName(icon_name string) {
	ptr := C.CString(icon_name)
	defer cfree(ptr)
	C.gtk_status_icon_set_from_icon_name(
		v.GStatusIcon, 
		gstring(ptr)
	)
}

func (v *StatusIcon) GetPixbuf() *gdkpixbuf.Pixbuf {
	return &gdkpixbuf.Pixbuf{C.gtk_status_icon_get_pixbuf(v.GStatusIcon)}
}

func (v *StatusIcon) GetStock() string {
	return gostring(C.gtk_status_icon_get_stock(v.GStatusIcon))
}

func (v *StatusIcon) GetIconName() string {
	return gostring(C.gtk_status_icon_get_icon_name(v.GStatusIcon))
}

func (v *StatusIcon) SetName(name string) {
	panic_if_version_older(2, 20, 0, "gtk_status_icon_set_name()")
	ptr := C.CString(name)
	defer cfree(ptr)
	C._gtk_status_icon_set_name(
		v.GStatusIcon, 
		gstring(ptr)
	)
}

func (v *StatusIcon) SetTitle(title string) {
	panic_if_version_older(2, 18, 0, "gtk_status_icon_set_title()")
	ptr := C.CString(title)
	defer cfree(ptr)
	C._gtk_status_icon_set_title(
		v.GStatusIcon, 
		gstring(ptr)
	)
}

func (v *StatusIcon) GetTitle() string {
	panic_if_version_older(2, 18, 0, "gtk_status_icon_get_title()")
	return gostring(C._gtk_status_icon_get_title(v.GStatusIcon))
}

func (v *StatusIcon) SetTooltipText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_status_icon_set_tooltip_text(
		v.GStatusIcon, 
		gstring(ptr)
	)
}

func (v *StatusIcon) GetTooltipText() string {
	return gostring(C.gtk_status_icon_get_tooltip_text(v.GStatusIcon))
}

func (v *StatusIcon) SetTooltipMarkup(markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_status_icon_set_tooltip_markup(
		v.GStatusIcon, 
		gstring(ptr)
	)
}

func (v *StatusIcon) GetTooltipMarkup() string {
	return gostring(C.gtk_status_icon_get_tooltip_markup(v.GStatusIcon))
}

func (v *StatusIcon) GetHasTooltip() bool {
	return gobool(C.gtk_status_icon_get_has_tooltip(v.GStatusIcon))
}

func (v *StatusIcon) SetHasTooltip(setting bool) {
	C.gtk_status_icon_set_has_tooltip(
		v.GStatusIcon, 
		gbool(setting)
	)
}

func (v *StatusIcon) GetVisible() bool {
	return gobool(C.gtk_status_icon_get_visible(v.GStatusIcon))
}

func (v *StatusIcon) SetVisible(setting bool) {
	C.gtk_status_icon_set_visible(
		v.GStatusIcon, 
		gbool(setting)
	)
}

func StatusIconPositionMenu(menu *Menu, px, py *int, push_in *bool, data interface{}) {
	x := gint(*px)
	y := gint(*py)
	pi := gbool(*push_in)
	var pdata C.gpointer
	if sm, ok := data.(*StatusIcon); ok {
		pdata = C.gpointer(unsafe.Pointer(sm.GStatusIcon))
	}
	C.gtk_status_icon_position_menu(
		MENU(menu), 
		&x, 
		&y, 
		&pi, 
		pdata
	)
	*px = int(x)
	*py = int(y)
	*push_in = gobool(pi)
}

func (v *StatusIcon) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(C.toGObject(unsafe.Pointer(v.GStatusIcon)))).Connect(s, f, datas...)
}

func PrintContextFromNative(l unsafe.Pointer) *PrintContext {
	return &PrintContext{(*C.GtkPrintContext)(l)}
}

// gtkButton

type ReliefStyle int

const (
	RELIEF_NORMAL ReliefStyle = 0
	RELIEF_HALF   ReliefStyle = 1
	RELIEF_NONE   ReliefStyle = 2
)

type Button struct {
	Bin
}

func (Button) isILabel() {
	//LEFT 4 DEAD
} 

func NewButton() *Button {
	return &Button{Bin{Container{Widget{C.gtk_button_new()}}}}
}

func NewButtonWithLabel(label string) *Button {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &Button{Bin{Container{Widget{C.gtk_button_new_with_label(gstring(ptr))}}}}
}

func NewButtonWithMnemonic(label string) *Button {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &Button{Bin{Container{Widget{C.gtk_button_new_with_mnemonic(gstring(ptr))}}}}
}

func NewButtonFromStock(stock_id string) *Button {
	p_stock_id := C.CString(stock_id)
	defer cfree(p_stock_id)
	return &Button{Bin{Container{Widget{C.gtk_button_new_from_stock(gstring(p_stock_id))}}}}
}

func (v *Button) Pressed() {
	deprecated_since(2, 20, 0, "gtk_button_pressed()")
	C.gtk_button_pressed(BUTTON(v))
}

func (v *Button) Released() {
	deprecated_since(2, 20, 0, "gtk_button_released()")
	C.gtk_button_released(BUTTON(v))
}

func (v *Button) Clicked(onclick interface{}, datas ...interface{}) int {
	return v.Connect("clicked", onclick, datas...)
}

func (v *Button) Enter() {
	deprecated_since(2, 20, 0, "gtk_button_enter()")
	C.gtk_button_enter(BUTTON(v))
}

func (v *Button) Leave() {
	deprecated_since(2, 20, 0, "gtk_button_leave()")
	C.gtk_button_leave(BUTTON(v))
}

func (v *Button) GetRelief() ReliefStyle {
	return ReliefStyle(C.gtk_button_get_relief(BUTTON(v)))
}

func (v *Button) SetRelief(relief ReliefStyle) {
	C.gtk_button_set_relief(
		BUTTON(v), 
		C.GtkReliefStyle(relief)
	)
}

func (v *Button) GetLabel() string {
	return gostring(C.gtk_button_get_label(BUTTON(v)))
}

func (v *Button) SetLabel(label string) {
	ptr := C.CString(label)
	defer cfree(ptr)
	C.gtk_button_set_label(
		BUTTON(v), 
		gstring(ptr)
	)
}

func (v *Button) GetUseStock() bool {
	return gobool(C.gtk_button_get_use_stock(BUTTON(v)))
}

func (v *Button) SetUseStock(use bool) {
	C.gtk_button_set_use_stock(
		BUTTON(v), 
		gbool(use)
	)
}

func (v *Button) GetUseUnderline() bool {
	return gobool(C.gtk_button_get_use_underline(BUTTON(v)))
}

func (v *Button) SetUseUnderline(setting bool) {
	C.gtk_button_set_use_underline(
		BUTTON(v), 
		gbool(setting)
	)
}

func (v *Button) GetFocusOnClick() bool {
	return gobool(C.gtk_button_get_focus_on_click(BUTTON(v)))
}

func (v *Button) SetFocusOnClick(setting bool) {
	C.gtk_button_set_focus_on_click(
		BUTTON(v), 
		gbool(setting)
	)
}

func (v *Button) SetAlignment(xalign, yalign float64) {
	C.gtk_button_set_alignment(
		BUTTON(v), 
		C.gfloat(xalign), 
		C.gfloat(yalign)
	)
}

func (v *Button) GetAlignment() (xalign, yalign float64) {
	var xalign_, yalign_ C.gfloat
	C.gtk_button_get_alignment(
		BUTTON(v), 
		&xalign_, 
		&yalign_
	)
	return float64(xalign_), float64(yalign_)
}

func (v *Button) SetImage(image IWidget) {
	C.gtk_button_set_image(
		BUTTON(v), 
		ToNative(image)
	)
}

func (v *Button) GetImage() *Image {
	return &Image{Widget{C.gtk_button_get_image(BUTTON(v))}}
}

// gtkCheckButton
type CheckButton struct {
	ToggleButton
}

func NewCheckButton() *CheckButton {
	return &CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_check_button_new()}}}}}}
}

func NewCheckButtonWithLabel(label string) *CheckButton {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_check_button_new_with_label(gstring(ptr))}}}}}}
}

func NewCheckButtonWithMnemonic(label string) *CheckButton {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_check_button_new_with_mnemonic(gstring(ptr))}}}}}}
}

// gtkRadioButton
type RadioButton struct {
	CheckButton
}

func NewRadioButton(group *glib.SList) *RadioButton {
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new(gslist(group))}}}}}}}
}

func NewRadioButtonFromWidget(w *RadioButton) *RadioButton {
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_from_widget(RADIO_BUTTON(w))}}}}}}}
}

func NewRadioButtonWithLabel(group *glib.SList, label string) *RadioButton {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_with_label(gslist(group), gstring(ptr))}}}}}}}
}

func NewRadioButtonWithLabelFromWidget(w *RadioButton, label string) *RadioButton {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_with_label_from_widget(RADIO_BUTTON(w), gstring(ptr))}}}}}}}
}

func NewRadioButtonWithMnemonic(group *glib.SList, label string) *RadioButton {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_with_mnemonic(gslist(group), gstring(ptr))}}}}}}}
}

func NewRadioButtonWithMnemonicFromWidget(w *RadioButton, label string) *RadioButton {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_with_mnemonic_from_widget(RADIO_BUTTON(w), gstring(ptr))}}}}}}}
}

func (v *RadioButton) GetGroup() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_radio_button_get_group(RADIO_BUTTON(v))))
}

func (v *RadioButton) SetGroup(group *glib.SList) {
	C.gtk_radio_button_set_group(RADIO_BUTTON(v), gslist(group))
}

// gtkToggleButton
type ToggleButton struct {
	Button
}

func NewToggleButton() *ToggleButton {
	return &ToggleButton{Button{Bin{Container{Widget{C.gtk_toggle_button_new()}}}}}
}

func NewToggleButtonWithLabel(label string) *ToggleButton {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &ToggleButton{Button{Bin{Container{Widget{C.gtk_toggle_button_new_with_label(gstring(ptr))}}}}}
}

func NewToggleButtonWithMnemonic(label string) *ToggleButton {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &ToggleButton{Button{Bin{Container{Widget{C.gtk_toggle_button_new_with_mnemonic(gstring(ptr))}}}}}
}

func (v *ToggleButton) SetMode(draw_indicator bool) {
	C.gtk_toggle_button_set_mode(
		TOGGLE_BUTTON(v), 
		gbool(draw_indicator)
	)
}

func (v *ToggleButton) GetMode() bool {
	return gobool(C.gtk_toggle_button_get_mode(TOGGLE_BUTTON(v)))
}

func (v *ToggleButton) GetActive() bool {
	return gobool(C.gtk_toggle_button_get_active(TOGGLE_BUTTON(v)))
}

func (v *ToggleButton) SetActive(is_active bool) {
	C.gtk_toggle_button_set_active(
		TOGGLE_BUTTON(v), 
		gbool(is_active)
	)
}

func (v *ToggleButton) GetInconsistent() bool {
	return gobool(C.gtk_toggle_button_get_inconsistent(TOGGLE_BUTTON(v)))
}

func (v *ToggleButton) SetInconsistent(setting bool) {
	C.gtk_toggle_button_set_inconsistent(
		TOGGLE_BUTTON(v), 
		gbool(setting)
	)
}

// gtkLinkButton
type LinkButton struct {
	Button
}

func NewLinkButton(uri string) *LinkButton {
	ptr := C.CString(uri)
	defer cfree(ptr)
	return &LinkButton{Button{Bin{Container{Widget{
		C.gtk_link_button_new(gstring(ptr))}}}}}
}

func NewLinkButtonWithLabel(uri string, label string) *LinkButton {
	puri := C.CString(uri)
	defer cfree(puri)
	plabel := C.CString(label)
	defer cfree(plabel)
	return &LinkButton{Button{Bin{Container{Widget{C.gtk_link_button_new_with_label(gstring(puri), gstring(plabel))}}}}}
}

func (v *LinkButton) GetUri() string {
	return gostring(C.gtk_link_button_get_uri(LINK_BUTTON(v)))
}

func (v *LinkButton) SetUri(uri string) {
	ptr := C.CString(uri)
	defer cfree(ptr)
	C.gtk_link_button_set_uri(
		LINK_BUTTON(v), 
		gstring(ptr)
	)
}

func (v *LinkButton) GetVisited() bool {
	return gobool(C.gtk_link_button_get_visited(LINK_BUTTON(v)))
}

func (v *LinkButton) SetVisited(visited bool) {
	C.gtk_link_button_set_visited(
		LINK_BUTTON(v), 
		gbool(visited)
	)
}

// gtkEntry
type Entry struct {
	Widget
	Editable
}

func NewEntry() *Entry {
	w := Widget{C.gtk_entry_new()}
	return &Entry{w, Editable{C.toGEditable(w.GWidget)}}
}

func NewEntryWithBuffer(buffer *EntryBuffer) *Entry {
	panic_if_version_older_auto(2, 18, 0)
	w := Widget{C._gtk_entry_new_with_buffer(buffer.GEntryBuffer)}
	return &Entry{w, Editable{C.toGEditable(w.GWidget)}}
}

func (v *Entry) GetBuffer() *EntryBuffer {
	panic_if_version_older_auto(2, 18, 0)
	return &EntryBuffer{C._gtk_entry_get_buffer(ENTRY(v))}
}

func (v *Entry) SetBuffer(buffer *EntryBuffer) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_entry_set_buffer(
		ENTRY(v), 
		buffer.GEntryBuffer
	)
}

func (v *Entry) SetText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_entry_set_text(
		ENTRY(v), 
		gstring(ptr)
	)
}

func (v *Entry) GetText() string {
	return gostring(C.gtk_entry_get_text(ENTRY(v)))
}

func (v *Entry) GetTextLength() int {
	return int(C.gtk_entry_get_text_length(ENTRY(v)))
}

func (v *Entry) SetVisibility(setting bool) {
	C.gtk_entry_set_visibility(
		ENTRY(v), 
		gbool(setting)
	)
}

func (v *Entry) SetInvisibleChar(ch uint8) {
	C.gtk_entry_set_invisible_char(
		ENTRY(v), 
		C.gunichar(ch)
	)
}

func (v *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(ENTRY(v))
}

func (v *Entry) SetMaxLength(i int) {
	C.gtk_entry_set_max_length(
		ENTRY(v), 
		gint(i)
	)
}

func (v *Entry) GetActivatesDefault() bool {
	return gobool(C.gtk_entry_get_activates_default(ENTRY(v)))
}

func (v *Entry) GetHasFrame() bool {
	return gobool(C.gtk_entry_get_has_frame(ENTRY(v)))
}

func (v *Entry) GetWidthChars() int {
	return int(C.gtk_entry_get_width_chars(ENTRY(v)))
}

func (v *Entry) SetActivatesDefault(setting bool) {
	C.gtk_entry_set_activates_default(
		ENTRY(v), 
		gbool(setting)
	)
}

func (v *Entry) SetHasFrame(setting bool) {
	C.gtk_entry_set_has_frame(
		ENTRY(v), 
		gbool(setting)
	)
}

func (v *Entry) SetWidthChars(i int) {
	C.gtk_entry_set_width_chars(
		ENTRY(v), 
		gint(i)
	)
}

func (v *Entry) GetInvisibleChar() uint8 {
	return uint8(C.gtk_entry_get_invisible_char(ENTRY(v)))
}

func (v *Entry) SetAlignment(xalign float64) {
	C.gtk_entry_set_alignment(
		ENTRY(v), 
		C.gfloat(xalign)
	)
}

func (v *Entry) GetAlignment() float64 {
	return float64(C.gtk_entry_get_alignment(ENTRY(v)))
}

func (v *Entry) SetOverwriteMode(mode bool) {
	C.gtk_entry_set_overwrite_mode(
		ENTRY(v), 
		gbool(mode)
	)
}

func (v *Entry) GetOverwriteMode() bool {
	return gobool(C.gtk_entry_get_overwrite_mode(ENTRY(v)))
}

func (v *Entry) GetMaxLength() int {
	return int(C.gtk_entry_get_max_length(ENTRY(v)))
}

func (v *Entry) GetVisibility() bool {
	return gobool(C.gtk_entry_get_visibility(ENTRY(v)))
}

func (v *Entry) SetCompletion(completion *EntryCompletion) {
	C.gtk_entry_set_completion(ENTRY(v), completion.GEntryCompletion)
}

var g_Entry_EntryCompletion *EntryCompletion

func (v *Entry) GetCompletion() *EntryCompletion {
	if g_Entry_EntryCompletion == nil {
		g_Entry_EntryCompletion = &EntryCompletion{C.gtk_entry_get_completion(ENTRY(v))}
	} else {
		g_Entry_EntryCompletion.GEntryCompletion = C.gtk_entry_get_completion(ENTRY(v))
	}
	return g_Entry_EntryCompletion
}

// gtkEntryBuffer
type EntryBuffer struct {
	GEntryBuffer *C.GtkEntryBuffer
}

func NewEntryBuffer(initialText string) *EntryBuffer {
	panic_if_version_older_auto(2, 18, 0)
	if len(initialText) == 0 {
		return &EntryBuffer{C._gtk_entry_buffer_new(nil, gint(-1))}
	}
	ptr := C.CString(initialText)
	defer cfree(ptr)
	return &EntryBuffer{C._gtk_entry_buffer_new(gstring(ptr), gint(len(initialText)))}
}

func (v *EntryBuffer) GetText() string {
	panic_if_version_older_auto(2, 18, 0)
	return gostring(C._gtk_entry_buffer_get_text(v.GEntryBuffer))
}

func (v *EntryBuffer) SetText(text string) {
	panic_if_version_older_auto(2, 18, 0)
	if len(text) == 0 {
		C._gtk_entry_buffer_set_text(v.GEntryBuffer, nil, gint(-1))
	}
	ptr := C.CString(text)
	defer cfree(ptr)
	C._gtk_entry_buffer_set_text(
		v.GEntryBuffer, 
		gstring(ptr), 
		gint(len(text))
	)
}

func (v *EntryBuffer) GetLength() uint {
	panic_if_version_older_auto(2, 18, 0)
	return uint(C._gtk_entry_buffer_get_length(v.GEntryBuffer))
}

func (v *EntryBuffer) GetMaxLength() int {
	panic_if_version_older_auto(2, 18, 0)
	return int(C._gtk_entry_buffer_get_max_length(v.GEntryBuffer))
}

func (v *EntryBuffer) SetMaxLength(maxLength int) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_entry_buffer_set_max_length(v.GEntryBuffer, gint(maxLength))
}

func (v *EntryBuffer) InsertText(position uint, text string) uint {
	panic_if_version_older_auto(2, 18, 0)
	ptr := C.CString(text)
	defer cfree(ptr)
	return uint(C._gtk_entry_buffer_insert_text(v.GEntryBuffer, guint(position), gstring(ptr), gint(len(text))))
}

func (v *EntryBuffer) DeleteText(position uint, nChars int) uint {
	panic_if_version_older_auto(2, 18, 0)
	return uint(C._gtk_entry_buffer_delete_text(v.GEntryBuffer, guint(position), gint(nChars)))
}

// gtkEntryCompletion
type EntryCompletion struct {
	GEntryCompletion *C.GtkEntryCompletion
}

func NewEntryCompletion() *EntryCompletion {
	return &EntryCompletion{C.gtk_entry_completion_new()}
}

func (v *EntryCompletion) GetEntry() *Widget {
	return &Widget{C.gtk_entry_completion_get_entry(v.GEntryCompletion)}
}

func (v *EntryCompletion) SetModel(model *TreeModel) {
	C.gtk_entry_completion_set_model(
		v.GEntryCompletion, 
		model.GTreeModel
	)
}

func (v *EntryCompletion) GetModel() *TreeModel {
	return &TreeModel{C.gtk_entry_completion_get_model(v.GEntryCompletion)}
}

type EntryCompletionMatchFunc func(completion *EntryCompletion, key string, iter *TreeIter, data interface{})

func (v *EntryCompletion) SetMinimumKeyLength(length int) {
	C.gtk_entry_completion_set_minimum_key_length(
		v.GEntryCompletion, 
		gint(length)
	)
}

func (v *EntryCompletion) GetMinimumKeyLength() int {
	return int(C.gtk_entry_completion_get_minimum_key_length(v.GEntryCompletion))
}

func (v *EntryCompletion) Complete() {
	C.gtk_entry_completion_complete(v.GEntryCompletion)
}

func (v *EntryCompletion) GetCompletionPrefix() string {
	return gostring(C.gtk_entry_completion_get_completion_prefix(v.GEntryCompletion))
}

func (v *EntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix(v.GEntryCompletion)
}

func (v *EntryCompletion) InsertActionText(index int, text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_entry_completion_insert_action_text(
		v.GEntryCompletion, 
		gint(index), 
		gstring(ptr)
	)
}

func (v *EntryCompletion) InsertActionMarkup(index int, markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_entry_completion_insert_action_markup(
		v.GEntryCompletion, 
		gint(index), 
		gstring(ptr)
	)
}

func (v *EntryCompletion) DeleteAction(index int) {
	C.gtk_entry_completion_delete_action(
		v.GEntryCompletion, 
		gint(index)
	)
}

func (v *EntryCompletion) SetTextColumn(column int) {
	C.gtk_entry_completion_set_text_column(
		v.GEntryCompletion, 
		gint(column)
	)
}

func (v *EntryCompletion) GetTextColumn() int {
	return int(C.gtk_entry_completion_get_text_column(v.GEntryCompletion))
}

func (v *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	C.gtk_entry_completion_set_inline_completion(
		v.GEntryCompletion,
		gbool(inlineCompletion)
	)
}

func (v *EntryCompletion) GetInlineCompletion() bool {
	return gobool(C.gtk_entry_completion_get_inline_completion(v.GEntryCompletion))
}

func (v *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	C.gtk_entry_completion_set_inline_selection(
		v.GEntryCompletion,
		gbool(inlineSelection)
	)
}

func (v *EntryCompletion) GetInlineSelection() bool {
	return gobool(C.gtk_entry_completion_get_inline_selection(v.GEntryCompletion))
}

func (v *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	C.gtk_entry_completion_set_popup_completion(
		v.GEntryCompletion,
		gbool(popupCompletion)
	)
}

func (v *EntryCompletion) GetPopupCompletion() bool {
	return gobool(C.gtk_entry_completion_get_popup_completion(v.GEntryCompletion))
}

func (v *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	C.gtk_entry_completion_set_popup_set_width(
		v.GEntryCompletion,
		gbool(popupSetWidth)
	)
}

func (v *EntryCompletion) GetPopupSetWidth() bool {
	return gobool(C.gtk_entry_completion_get_popup_set_width(v.GEntryCompletion))
}

func (v *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	C.gtk_entry_completion_set_popup_single_match(
		v.GEntryCompletion,
		gbool(popupSingleMatch)
	)
}

func (v *EntryCompletion) GetPopupSingleMatch() bool {
	return gobool(C.gtk_entry_completion_get_popup_single_match(v.GEntryCompletion))
}

// gtkHScale
func NewHScale(adjustment *Adjustment) *Scale {
	return &Scale{Range{Widget{C.gtk_hscale_new(adjustment.GAdjustment)}}}
}

func NewHScaleWithRange(min, max, step float64) *Scale {
	return &Scale{Range{Widget{C.gtk_hscale_new_with_range(gdouble(min), gdouble(max), gdouble(step))}}}
}

// gtkVScale
func NewVScale(a *Adjustment) *Scale {
	return &Scale{Range{Widget{C.gtk_vscale_new(a.GAdjustment)}}}
}

func NewVScaleWithRange(min, max, step float64) *Scale {
	return &Scale{Range{Widget{C.gtk_vscale_new_with_range(gdouble(min), gdouble(max), gdouble(step))}}}
}

// gtkSpinButton
type SpinButtonUpdatePolicy int

const (
	UPDATE_ALWAYS   = 0
	UPDATE_IF_VALID = 1
)

type SpinType int

const (
	SPIN_STEP_FORWARD  = 0
	SPIN_STEP_BACKWARD = 1
	SPIN_PAGE_FORWARD  = 2
	SPIN_PAGE_BACKWARD = 3
	SPIN_HOME          = 4
	SPIN_END           = 5
	SPIN_USER_DEFINED  = 6
)

type SpinButton struct {
	Entry
}

func NewSpinButton(a *Adjustment, climb float64, digits uint) *SpinButton {
	w := Widget{C.gtk_spin_button_new(a.GAdjustment, gdouble(climb), guint(digits))}
	return &SpinButton{Entry{w, Editable{C.toGEditable(w.GWidget)}}}
}

func NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	w := Widget{C.gtk_spin_button_new_with_range(gdouble(min), gdouble(max), gdouble(step))}
	return &SpinButton{Entry{w, Editable{C.toGEditable(w.GWidget)}}}
}

func (v *SpinButton) OnChangeValue(onclick interface{}, datas ...interface{}) int {
	return v.Connect("change-value", onclick, datas...)
}

func (v *SpinButton) OnInput(onclick interface{}, datas ...interface{}) int {
	return v.Connect("input", onclick, datas...)
}

func (v *SpinButton) OnOutput(onclick interface{}, datas ...interface{}) int {
	return v.Connect("output", onclick, datas...)
}

func (v *SpinButton) OnValueChanged(onclick interface{}, datas ...interface{}) int {
	return v.Connect("value-changed", onclick, datas...)
}

func (v *SpinButton) OnWrapped(onclick interface{}, datas ...interface{}) int {
	return v.Connect("wrapped", onclick, datas...)
}

func (v *SpinButton) Configure(a *Adjustment, climb_rate float64, digits uint) {
	C.gtk_spin_button_configure(
		SPIN_BUTTON(v), 
		a.GAdjustment, 
		gdouble(climb_rate), 
		guint(digits)
	)
}

func (v *SpinButton) SetAdjustment(a *Adjustment) {
	C.gtk_spin_button_set_adjustment(
		SPIN_BUTTON(v), 
		a.GAdjustment
	)
}

func (v *SpinButton) GetAdjustment() *Adjustment {
	return &Adjustment{C.gtk_spin_button_get_adjustment(SPIN_BUTTON(v))}
}

func (v *SpinButton) SetDigits(digits uint) {
	C.gtk_spin_button_set_digits(
		SPIN_BUTTON(v), 
		guint(digits)
	)
}

func (v *SpinButton) SetIncrements(step, page float64) {
	C.gtk_spin_button_set_increments(
		SPIN_BUTTON(v), 
		gdouble(step), 
		gdouble(page)
	)
}

func (v *SpinButton) SetRange(min, max float64) {
	C.gtk_spin_button_set_range(
		SPIN_BUTTON(v), 
		gdouble(min), 
		gdouble(max)
	)
}

func (v *SpinButton) GetValueAsFloat() float64 {
	return float64(C.gtk_spin_button_get_value_as_float(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetValueAsInt() int {
	return int(C.gtk_spin_button_get_value_as_int(SPIN_BUTTON(v)))
}

func (v *SpinButton) SetValue(val float64) {
	C.gtk_spin_button_set_value(
		SPIN_BUTTON(v), 
		gdouble(val)
	)
}

func (v *SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	C.gtk_spin_button_set_update_policy(
		SPIN_BUTTON(v), 
		C.GtkSpinButtonUpdatePolicy(policy)
	)
}

func (v *SpinButton) SetNumeric(numeric bool) {
	C.gtk_spin_button_set_numeric(
		SPIN_BUTTON(v), 
		gbool(numeric)
	)
}

func (v *SpinButton) Spin(direction SpinType, increment float64) {
	C.gtk_spin_button_spin(
		SPIN_BUTTON(v), 
		C.GtkSpinType(direction), 
		gdouble(increment)
	)
}

func (v *SpinButton) SetWrap(wrap bool) {
	C.gtk_spin_button_set_wrap(
		SPIN_BUTTON(v), 
		gbool(wrap)
	)
}

func (v *SpinButton) SetSnapToTicks(snap_to_ticks bool) {
	C.gtk_spin_button_set_snap_to_ticks(
		SPIN_BUTTON(v), 
		gbool(snap_to_ticks)
	)
}

func (v *SpinButton) Update() {
	C.gtk_spin_button_update(SPIN_BUTTON(v))
}

func (v *SpinButton) GetDigits() uint {
	return uint(C.gtk_spin_button_get_digits(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetIncrements() (float64, float64) {
	var step, page C.gdouble
	C.gtk_spin_button_get_increments(
		SPIN_BUTTON(v), 
		&step, 
		&page
	)
	return float64(step), float64(page)
}

func (v *SpinButton) GetNumeric() bool {
	return gobool(C.gtk_spin_button_get_numeric(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetRange() (float64, float64) {
	var min, max C.gdouble
	C.gtk_spin_button_get_range(
		SPIN_BUTTON(v), 
		&min, 
		&max
	)
	return float64(min), float64(max)
}

func (v *SpinButton) GetSnapToTicks() bool {
	return gobool(C.gtk_spin_button_get_snap_to_ticks(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy {
	return SpinButtonUpdatePolicy(C.gtk_spin_button_get_update_policy(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetValue() float64 {
	return float64(C.gtk_spin_button_get_value(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetWrap() bool {
	return gobool(C.gtk_spin_button_get_wrap(SPIN_BUTTON(v)))
}

// gtkEditable
type Editable struct {
	GEditable *C.GtkEditable
}

func (v *Editable) SelectRegion(startPos int, endPos int) {
	C.gtk_editable_select_region(
		v.GEditable, 
		gint(startPos), 
		gint(endPos)
	)
}

func (v *Editable) GetSelectionBounds() (isSelected bool, startPos int, endPos int) {
	var s, e C.gint
	return gobool(C.gtk_editable_get_selection_bounds(v.GEditable, &s, &e)), int(s), int(e)
}

func (v *Editable) InsertText(newText string, position int) int {
	ptr := C.CString(newText)
	defer cfree(ptr)
	gpos := (C.gint)(position)
	C.gtk_editable_insert_text(
		v.GEditable, 
		gstring(ptr), 
		gint(len(newText)), 
		&gpos
	)
	return int(gpos)
}

func (v *Editable) DeleteText(startPos int, endPos int) {
	C.gtk_editable_delete_text(
		v.GEditable, 
		gint(startPos), 
		gint(endPos)
	)
}

func (v *Editable) GetChars(startPos int, endPos int) string {
	return gostring(C.gtk_editable_get_chars(v.GEditable, gint(startPos), gint(endPos)))
}

func (v *Editable) CutClipboard() {
	C.gtk_editable_cut_clipboard(v.GEditable)
}

func (v *Editable) CopyClipboard() {
	C.gtk_editable_copy_clipboard(v.GEditable)
}

func (v *Editable) PasteClipboard() {
	C.gtk_editable_paste_clipboard(v.GEditable)
}

func (v *Editable) DeleteSelection() {
	C.gtk_editable_delete_selection(v.GEditable)
}

func (v *Editable) SetPosition(position int) {
	C.gtk_editable_set_position(
		v.GEditable, 
		gint(position)
	)
}

func (v *Editable) GetPosition() int {
	return int(C.gtk_editable_get_position(v.GEditable))
}

func (v *Editable) SetEditable(isEditable bool) {
	C.gtk_editable_set_editable(
		v.GEditable, 
		gbool(isEditable)
	)
}

func (v *Editable) GetEditable() bool {
	return gobool(C.gtk_editable_get_editable(v.GEditable))
}

// gtkTextIter
type TextIter struct {
	GTextIter C.GtkTextIter
}

type TextSearchFlags int

const (
	TEXT_SEARCH_VISIBLE_ONLY     TextSearchFlags = 1 << 0
	TEXT_SEARCH_TEXT_ONLY        TextSearchFlags = 1 << 1
	TEXT_SEARCH_CASE_INSENSITIVE TextSearchFlags = 1 << 2
)

func (v *TextIter) GetBuffer() *TextBuffer {
	return newTextBuffer(C.gtk_text_iter_get_buffer(&v.GTextIter))
}

func (v *TextIter) Copy() *TextIter {
	return &TextIter{*C.gtk_text_iter_copy(&v.GTextIter)}
}

func (v *TextIter) Free() {
	C.gtk_text_iter_free(&v.GTextIter)
}

func (v *TextIter) GetOffset() int {
	return int(C.gtk_text_iter_get_offset(&v.GTextIter))
}

func (v *TextIter) GetLine() int {
	return int(C.gtk_text_iter_get_line(&v.GTextIter))
}

func (v *TextIter) GetLineOffset() int {
	return int(C.gtk_text_iter_get_line_offset(&v.GTextIter))
}

func (v *TextIter) GetLineIndex() int {
	return int(C.gtk_text_iter_get_line_index(&v.GTextIter))
}

func (v *TextIter) GetVisibleLineIndex() int {
	return int(C.gtk_text_iter_get_visible_line_index(&v.GTextIter))
}

func (v *TextIter) GetVisibleLineOffset() int {
	return int(C.gtk_text_iter_get_visible_line_offset(&v.GTextIter))
}

func (v *TextIter) GetChar() int {
	return int(C.gtk_text_iter_get_char(&v.GTextIter))
}

func (v *TextIter) GetSlice(end *TextIter) string {
	pchar := cstring(C.gtk_text_iter_get_slice(&v.GTextIter, &end.GTextIter))
	defer cfree(pchar)
	return C.GoString(pchar)
}

func (v *TextIter) GetText(end *TextIter) string {
	pchar := cstring(C.gtk_text_iter_get_text(&v.GTextIter, &end.GTextIter))
	defer cfree(pchar)
	return C.GoString(pchar)
}

func (v *TextIter) GetVisibleSlice(end *TextIter) string {
	return gostring(C.gtk_text_iter_get_visible_slice(&v.GTextIter, &end.GTextIter))
}

func (v *TextIter) GetVisibleText(end *TextIter) string {
	return gostring(C.gtk_text_iter_get_visible_text(&v.GTextIter, &end.GTextIter))
}

func (v *TextIter) GetMarks() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_text_iter_get_marks(&v.GTextIter)))
}

func (v *TextIter) ForwardChar() bool {
	return gobool(C.gtk_text_iter_forward_char(&v.GTextIter))
}

func (v *TextIter) ForwardSearch(str string, flags TextSearchFlags, start *TextIter, end *TextIter, limit *TextIter) bool {
	cstr := C.CString(str)
	defer cfree(cstr)
	return gobool(C.gtk_text_iter_forward_search(&v.GTextIter, gstring(cstr), C.GtkTextSearchFlags(flags), &start.GTextIter, &end.GTextIter, &limit.GTextIter))
}

func (v *TextIter) BackwardSearch(str string, flags TextSearchFlags, start *TextIter, end *TextIter, limit *TextIter) bool {
	cstr := C.CString(str)
	defer cfree(cstr)
	return gobool(C.gtk_text_iter_backward_search(&v.GTextIter, gstring(cstr), C.GtkTextSearchFlags(flags), &start.GTextIter, &end.GTextIter, &limit.GTextIter))
}

func (v *TextIter) Assign(iter *TextIter) {
	C._gtk_text_iter_assign(
		&v.GTextIter, 
		&iter.GTextIter
	)
}

// gtkTextMark
type TextMark struct {
	GTextMark *C.GtkTextMark
}

// gtkTextBuffer
type ITextBuffer interface {
	GetNativeBuffer() unsafe.Pointer
}

type TextBuffer struct {
	GTextBuffer *C.GtkTextBuffer
	*glib.GObject
}

func newTextBuffer(buffer *C.GtkTextBuffer) *TextBuffer { 
	return &TextBuffer{
		GTextBuffer: buffer,
		GObject:     glib.ObjectFromNative(unsafe.Pointer(buffer)),
	}
}

func NewTextBufferFromPointer(v unsafe.Pointer) TextBuffer {
	return *newTextBuffer(TEXT_BUFFER(v))
}

func NewTextBuffer(tagtable *TextTagTable) *TextBuffer {
	return newTextBuffer(C.gtk_text_buffer_new(tagtable.GTextTagTable))
}

func (v *TextBuffer) GetNativeBuffer() unsafe.Pointer {
	return unsafe.Pointer(v.GTextBuffer)
}

func (v *TextBuffer) GetLineCount() int {
	return int(C.gtk_text_buffer_get_line_count(v.GTextBuffer))
}

func (v *TextBuffer) GetCharCount() int {
	return int(C.gtk_text_buffer_get_char_count(v.GTextBuffer))
}

func (v *TextBuffer) GetTagTable() *TextTagTable {
	return &TextTagTable{C.gtk_text_buffer_get_tag_table(v.GTextBuffer)}
}

func (v *TextBuffer) Insert(iter *TextIter, text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_insert(
		v.GTextBuffer, 
		&iter.GTextIter, 
		gstring(ptr), 
		gsize_t(l)
	)
}

func (v *TextBuffer) InsertAtCursor(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_insert_at_cursor(
		v.GTextBuffer, 
		gstring(ptr), 
		gsize_t(l)
	)
}

func (v *TextBuffer) InsertInteractive(iter *TextIter, text string, default_editable bool) bool {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	return gobool(C.gtk_text_buffer_insert_interactive(v.GTextBuffer, &iter.GTextIter, gstring(ptr), gsize_t(l), gbool(default_editable)))
}

func (v *TextBuffer) InsertInteractiveAtCursor(text string, default_editable bool) bool {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	return gobool(C.gtk_text_buffer_insert_interactive_at_cursor(v.GTextBuffer, gstring(ptr), gsize_t(l), gbool(default_editable)))
}

func (v *TextBuffer) InsertRange(iter *TextIter, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_insert_range(
		v.GTextBuffer, 
		&iter.GTextIter, 
		&start.GTextIter, 
		&end.GTextIter
	)
}

func (v *TextBuffer) InsertRangeInteractive(iter *TextIter, start *TextIter, end *TextIter, default_editable bool) bool {
	return gobool(C.gtk_text_buffer_insert_range_interactive(v.GTextBuffer, &iter.GTextIter, &start.GTextIter, &end.GTextIter, gbool(default_editable)))
}

func (v *TextBuffer) InsertWithTag(iter *TextIter, text string, tag *TextTag) {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	C._gtk_text_buffer_insert_with_tag(
		v.GTextBuffer, 
		&iter.GTextIter, 
		gstring(ptr), 
		gsize_t(l), 
		tag.GTextTag
	)
}

func (v *TextBuffer) Delete(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_delete(
		v.GTextBuffer, 
		&start.GTextIter, 
		&end.GTextIter
	)
}

func (v *TextBuffer) DeleteInteractive(start *TextIter, end *TextIter, default_editable bool) bool {
	return gobool(C.gtk_text_buffer_delete_interactive(v.GTextBuffer, &start.GTextIter, &end.GTextIter, gbool(default_editable)))
}

func (v *TextBuffer) Backspace(iter *TextIter, interactive bool, default_editable bool) bool {
	return gobool(C.gtk_text_buffer_backspace(v.GTextBuffer, &iter.GTextIter, gbool(interactive), gbool(default_editable)))
}

func (v *TextBuffer) SetText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_set_text(
		v.GTextBuffer, 
		gstring(ptr), 
		gsize_t(l)
	)
}

func (v *TextBuffer) GetText(start *TextIter, end *TextIter, include_hidden_chars bool) string {
	pchar := cstring(C.gtk_text_buffer_get_text(v.GTextBuffer, &start.GTextIter, &end.GTextIter, gbool(include_hidden_chars)))
	defer cfree(pchar)
	return C.GoString(pchar)
}

func (v *TextBuffer) GetSlice(start *TextIter, end *TextIter, include_hidden_chars bool) string {
	pchar := cstring(C.gtk_text_buffer_get_slice(v.GTextBuffer, &start.GTextIter, &end.GTextIter, gbool(include_hidden_chars)))
	defer cfree(pchar)
	return C.GoString(pchar)
}

func (v *TextBuffer) InsertPixbuf(iter *TextIter, pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_text_buffer_insert_pixbuf(
		v.GTextBuffer, 
		&iter.GTextIter, 
		pixbuf.GPixbuf
	)
}

func (v *TextBuffer) CreateMark(mark_name string, where *TextIter, left_gravity bool) *TextMark {
	ptr := C.CString(mark_name)
	defer cfree(ptr)
	return &TextMark{C.gtk_text_buffer_create_mark(v.GTextBuffer, gstring(ptr), &where.GTextIter, gbool(left_gravity))}
}

func (v *TextBuffer) MoveMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_move_mark(
		v.GTextBuffer, 
		mark.GTextMark, 
		&where.GTextIter
	)
}

func (v *TextBuffer) MoveMarkByName(name string, where *TextIter) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_text_buffer_move_mark_by_name(
		v.GTextBuffer, 
		gstring(ptr), 
		&where.GTextIter
	)
}

func (v *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_add_mark(
		v.GTextBuffer, 
		mark.GTextMark, 
		&where.GTextIter
	)
}

func (v *TextBuffer) DeleteMark(mark *TextMark) {
	C.gtk_text_buffer_delete_mark(
		v.GTextBuffer, 
		mark.GTextMark
	)
}

func (v *TextBuffer) DeleteMarkByName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_text_buffer_delete_mark_by_name(
		v.GTextBuffer, 
		gstring(ptr)
	)
}

func (v *TextBuffer) GetMark(name string) *TextMark {
	ptr := C.CString(name)
	defer cfree(ptr)
	return &TextMark{C.gtk_text_buffer_get_mark(v.GTextBuffer, gstring(ptr))}
}

func (v *TextBuffer) GetInsert() *TextMark {
	return &TextMark{C.gtk_text_buffer_get_insert(v.GTextBuffer)}
}

func (v *TextBuffer) GetSelectionBound() *TextMark {
	return &TextMark{C.gtk_text_buffer_get_selection_bound(v.GTextBuffer)}
}

func (v *TextBuffer) GetHasSelection() bool {
	return gobool(C.gtk_text_buffer_get_has_selection(v.GTextBuffer))
}

func (v *TextBuffer) PlaceCursor(where *TextIter) {
	C.gtk_text_buffer_place_cursor(
		v.GTextBuffer, 
		&where.GTextIter
	)
}

func (v *TextBuffer) SelectRange(ins *TextIter, bound *TextIter) {
	C.gtk_text_buffer_select_range(
		v.GTextBuffer, 
		&ins.GTextIter, 
		&bound.GTextIter
	)
}

func (v *TextBuffer) ApplyTag(tag *TextTag, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_apply_tag(
		v.GTextBuffer, 
		tag.GTextTag, 
		&start.GTextIter, 
		&end.GTextIter
	)
}

func (v *TextBuffer) RemoveTag(tag *TextTag, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_remove_tag(
		v.GTextBuffer, 
		tag.GTextTag, 
		&start.GTextIter, 
		&end.GTextIter
	)
}

func (v *TextBuffer) ApplyTagByName(name string, start *TextIter, end *TextIter) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_text_buffer_apply_tag_by_name(
		v.GTextBuffer, 
		gstring(ptr), 
		&start.GTextIter, 
		&end.GTextIter
	)
}

func (v *TextBuffer) RemoveTagByName(name string, start *TextIter, end *TextIter) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_text_buffer_remove_tag_by_name(
		v.GTextBuffer, 
		gstring(ptr), 
		&start.GTextIter, 
		&end.GTextIter
	)
}

func (v *TextBuffer) RemoveAllTags(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_remove_all_tags(
		v.GTextBuffer, 
		&start.GTextIter, 
		&end.GTextIter
	)
}

func (v *TextBuffer) CreateTag(tag_name string, props map[string]string) *TextTag {
	ptr := C.CString(tag_name)
	defer cfree(ptr)
	tag := C._gtk_text_buffer_create_tag(v.GTextBuffer, gstring(ptr))
	for prop, val := range props {
		pprop := C.CString(prop)
		pval := C.CString(val)
		C._apply_property(
			unsafe.Pointer(tag), 
			gstring(pprop), 
			gstring(pval)
		)
		cfree(pprop)
		cfree(pval)
	}
	return &TextTag{tag}
}

func (v *TextBuffer) GetIterAtLineOffset(iter *TextIter, line_number int, char_offset int) {
	C.gtk_text_buffer_get_iter_at_line_offset(
		v.GTextBuffer, 
		&iter.GTextIter, 
		gint(line_number), 
		gint(char_offset)
	)
}

func (v *TextBuffer) GetIterAtOffset(iter *TextIter, char_offset int) {
	C.gtk_text_buffer_get_iter_at_offset(
		v.GTextBuffer, 
		&iter.GTextIter, 
		gint(char_offset)
	)
}

func (v *TextBuffer) GetIterAtLine(iter *TextIter, line_number int) {
	C.gtk_text_buffer_get_iter_at_line(
		v.GTextBuffer, 
		&iter.GTextIter, 
		gint(line_number)
	)
}

func (v *TextBuffer) GetIterAtLineIndex(iter *TextIter, line_number int, byte_index int) {
	C.gtk_text_buffer_get_iter_at_line_index(
		v.GTextBuffer, 
		&iter.GTextIter, 
		gint(line_number), 
		gint(byte_index)
	)
}

func (v *TextBuffer) GetIterAtMark(iter *TextIter, mark *TextMark) {
	C.gtk_text_buffer_get_iter_at_mark(
		v.GTextBuffer, 
		&iter.GTextIter, 
		mark.GTextMark
	)
}

func (v *TextBuffer) GetIterAtChildAnchor(i *TextIter, a *TextChildAnchor) {
	C.gtk_text_buffer_get_iter_at_child_anchor(
		v.GTextBuffer, 
		&i.GTextIter, 
		a.GTextChildAnchor
	)
}

func (v *TextBuffer) GetStartIter(iter *TextIter) {
	C.gtk_text_buffer_get_start_iter(
		v.GTextBuffer, 
		&iter.GTextIter
	)
}

func (v *TextBuffer) GetEndIter(iter *TextIter) {
	C.gtk_text_buffer_get_end_iter(
		v.GTextBuffer, 
		&iter.GTextIter
	)
}

func (v *TextBuffer) GetBounds(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_get_bounds(
		v.GTextBuffer, 
		&start.GTextIter, 
		&end.GTextIter
	)
}

func (v *TextBuffer) GetModified() bool {
	return gobool(C.gtk_text_buffer_get_modified(v.GTextBuffer))
}

func (v *TextBuffer) SetModified(setting bool) {
	C.gtk_text_buffer_set_modified(
		v.GTextBuffer, 
		gbool(setting)
	)
}

func (v *TextBuffer) DeleteSelection(interactive bool, default_editable bool) {
	C.gtk_text_buffer_delete_selection(
		v.GTextBuffer, 
		gbool(interactive), 
		gbool(default_editable)
	)
}

func (v *TextBuffer) GetSelectionBounds(be, en *TextIter) bool {
	return gobool(C.gtk_text_buffer_get_selection_bounds(v.GTextBuffer, &be.GTextIter, &en.GTextIter))
}

// gtkTextTag
type TextTag struct {
	GTextTag *C.GtkTextTag
}

func NewTextTag(name string) *TextTag {
	ptr := C.CString(name)
	defer cfree(ptr)
	return &TextTag{C.gtk_text_tag_new(gstring(ptr))}
}

func (v *TextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(
		v.GTextTag, 
		gint(priority)
	)
}

func (v *TextTag) GetPriority() int {
	return int(C.gtk_text_tag_get_priority(v.GTextTag))
}

// gtkTextAttributes
type TextAttributes struct {
	GTextAttributes *C.GtkTextAttributes
}

func NewTextAttributes() *TextAttributes {
	return &TextAttributes{C.gtk_text_attributes_new()}
}

func (v *TextAttributes) Copy() *TextAttributes {
	return &TextAttributes{C.gtk_text_attributes_copy(v.GTextAttributes)}
}

func (v *TextAttributes) CopyValues(ta *TextAttributes) {
	C.gtk_text_attributes_copy_values(
		v.GTextAttributes, 
		ta.GTextAttributes
	)
}

func (v *TextAttributes) Unref() {
	C.gtk_text_attributes_unref(v.GTextAttributes)
}

func (v *TextAttributes) Ref() *TextAttributes {
	return &TextAttributes{C.gtk_text_attributes_ref(v.GTextAttributes)}
}

// gtkTextTagTable
type TextTagTable struct {
	GTextTagTable *C.GtkTextTagTable
}

func NewTextTagTable() *TextTagTable {
	return &TextTagTable{C.gtk_text_tag_table_new()}
}

func (v *TextTagTable) Add(tag *TextTag) {
	C.gtk_text_tag_table_add(
		v.GTextTagTable, 
		tag.GTextTag
	)
}

func (v *TextTagTable) Remove(tag *TextTag) {
	C.gtk_text_tag_table_remove(
		v.GTextTagTable, 
		tag.GTextTag
	)
}

func (v *TextTagTable) Lookup(name string) *TextTag {
	ptr := C.CString(name)
	defer cfree(ptr)
	return &TextTag{C.gtk_text_tag_table_lookup(v.GTextTagTable, gstring(ptr))}
}

func (v *TextTagTable) GetSize() int {
	return int(C.gtk_text_tag_table_get_size(v.GTextTagTable))
}

// gtkTextView
type WrapMode int

const (
	WRAP_NONE      WrapMode = 0
	WRAP_CHAR      WrapMode = 1
	WRAP_WORD      WrapMode = 2
	WRAP_WORD_CHAR WrapMode = 3
)

type TextChildAnchor struct {
	GTextChildAnchor *C.GtkTextChildAnchor
}

type TextView struct {
	Container
}

func NewTextView() *TextView {
	return &TextView{Container{Widget{C.gtk_text_view_new()}}}
}

func NewTextViewWithBuffer(b TextBuffer) *TextView {
	return &TextView{Container{Widget{C.gtk_text_view_new_with_buffer(b.GTextBuffer)}}}
}

func (v *TextView) SetBuffer(b ITextBuffer) {
	C.gtk_text_view_set_buffer(
		TEXT_VIEW(v), 
		TEXT_BUFFER(b.GetNativeBuffer())
	)
}

func (v *TextView) GetBuffer() *TextBuffer {
	return newTextBuffer(C.gtk_text_view_get_buffer(TEXT_VIEW(v)))
}

func (v *TextView) ScrollToMark(mark *TextMark, wm float64, ua bool, xa float64, ya float64) {
	C.gtk_text_view_scroll_to_mark(
		TEXT_VIEW(v),
		mark.GTextMark, 
		gdouble(wm), 
		gbool(ua), 
		gdouble(xa), 
		gdouble(ya)
	)
}

func (v *TextView) ScrollToIter(iter *TextIter, wm float64, ua bool, xa float64, ya float64) bool {
	return gobool(C.gtk_text_view_scroll_to_iter(TEXT_VIEW(v), &iter.GTextIter, gdouble(wm), gbool(ua), gdouble(xa), gdouble(ya)))
}

func (v *TextView) GetLineYrange(iter *TextIter, y *int, h *int) {
	var yy, hh C.gint
	C.gtk_text_view_get_line_yrange(
		TEXT_VIEW(v), 
		&iter.GTextIter, 
		&yy, 
		&hh
	)
	*y = int(yy)
	*h = int(hh)
}

func (v *TextView) GetIterAtPosition(iter *TextIter, trailing *int, x int, y int) {
	if nil != trailing {
		var tt C.gint
		C.gtk_text_view_get_iter_at_position(
			TEXT_VIEW(v), 
			&iter.GTextIter, 
			&tt, 
			gint(x), 
			gint(y)
		)
		*trailing = int(tt)
	} else {
		C.gtk_text_view_get_iter_at_position(
			TEXT_VIEW(v), 
			&iter.GTextIter, 
			nil, 
			gint(x), 
			gint(y)
		)
	}
}

func (v *TextView) SetWrapMode(mode WrapMode) {
	C.gtk_text_view_set_wrap_mode(
		TEXT_VIEW(v), 
		C.GtkWrapMode(mode)
	)
}

func (v *TextView) GetWrapMode() WrapMode {
	return WrapMode(C.gtk_text_view_get_wrap_mode(TEXT_VIEW(v)))
}

func (v *TextView) SetEditable(setting bool) {
	C.gtk_text_view_set_editable(
		TEXT_VIEW(v), 
		gbool(setting)
	)
}

func (v *TextView) GetEditable() bool {
	return gobool(C.gtk_text_view_get_editable(TEXT_VIEW(v)))
}

func (v *TextView) SetCursorVisible(setting bool) {
	C.gtk_text_view_set_cursor_visible(
		TEXT_VIEW(v), 
		gbool(setting)
	)
}

func (v *TextView) GetCursorVisible() bool {
	return gobool(C.gtk_text_view_get_cursor_visible(TEXT_VIEW(v)))
}

func (v *TextView) SetOverwrite(overwrite bool) {
	C.gtk_text_view_set_overwrite(
		TEXT_VIEW(v), 
		gbool(overwrite)
	)
}

func (v *TextView) GetOverwrite() bool {
	return gobool(C.gtk_text_view_get_overwrite(TEXT_VIEW(v)))
}

func (v *TextView) SetAcceptsTab(accepts_tab bool) {
	C.gtk_text_view_set_accepts_tab(
		TEXT_VIEW(v), 
		gbool(accepts_tab)
	)
}

func (v *TextView) GetAcceptsTab() bool {
	return gobool(C.gtk_text_view_get_accepts_tab(TEXT_VIEW(v)))
}

// gtkTreePath
type TreePath struct {
	GTreePath *C.GtkTreePath
}

func NewTreePath() *TreePath {
	return &TreePath{C.gtk_tree_path_new()}
}

func NewTreePathFromString(path string) *TreePath {
	ptr := C.CString(path)
	defer cfree(ptr)
	return &TreePath{C.gtk_tree_path_new_from_string(gstring(ptr))}
}

func NewTreePathNewFirst() *TreePath {
	return &TreePath{C.gtk_tree_path_new_first()}
}

func (v *TreePath) String() string {
	return gostring(C.gtk_tree_path_to_string(v.GTreePath))
}

func (v *TreePath) AppendIndex(index int) {
	C.gtk_tree_path_append_index(
		v.GTreePath, 
		gint(index)
	)
}

func (v *TreePath) PrependIndex(index int) {
	C.gtk_tree_path_prepend_index(
		v.GTreePath, 
		gint(index)
	)
}

func (v *TreePath) GetDepth() int {
	return int(C.gtk_tree_path_get_depth(v.GTreePath))
}

func (v *TreePath) Free() {
	C.gtk_tree_path_free(v.GTreePath)
}

func (v *TreePath) Copy() *TreePath {
	return &TreePath{C.gtk_tree_path_copy(v.GTreePath)}
}

func (v *TreePath) Compare(w TreePath) int {
	return int(C.gtk_tree_path_compare(v.GTreePath, w.GTreePath))
}

func (v *TreePath) Next() {
	C.gtk_tree_path_next(v.GTreePath)
}

func (v *TreePath) Prev() bool {
	return gobool(C.gtk_tree_path_prev(v.GTreePath))
}

func (v *TreePath) Up() bool {
	return gobool(C.gtk_tree_path_up(v.GTreePath))
}

func (v *TreePath) Down() {
	C.gtk_tree_path_down(v.GTreePath)
}

func (v *TreePath) IsAncestor(descendant TreePath) bool {
	return gobool(C.gtk_tree_path_is_ancestor(v.GTreePath, descendant.GTreePath))
}

func (v *TreePath) IsDescendant(ancestor TreePath) bool {
	return gobool(C.gtk_tree_path_is_descendant(v.GTreePath, ancestor.GTreePath))
}

// gtkTreeIter
type TreeIter struct {
	GTreeIter C.GtkTreeIter
}

func (v *TreeIter) Assign(to *TreeIter) {
	C._gtk_tree_iter_assign(
		unsafe.Pointer(&v.GTreeIter), 
		unsafe.Pointer(&to.GTreeIter)
	)
}

// gtkTreeModel
type TreeModelFlags int

const (
	TREE_MODEL_ITERS_PERSIST TreeModelFlags = 1 << 0
	TREE_MODEL_LIST_ONLY     TreeModelFlags = 1 << 1
)

type ITreeModel interface {
	cTreeModel() *C.GtkTreeModel
}

type TreeModel struct {
	GTreeModel *C.GtkTreeModel
}

func (v TreeModel) cTreeModel() *C.GtkTreeModel {
	return v.GTreeModel
}

func (v *TreeModel) GetFlags() TreeModelFlags {
	return TreeModelFlags(C.gtk_tree_model_get_flags(v.GTreeModel))
}

func (v *TreeModel) GetNColumns() int {
	return int(C.gtk_tree_model_get_n_columns(v.GTreeModel))
}

func (v *TreeModel) GetIter(iter *TreeIter, path *TreePath) bool {
	return gobool(C.gtk_tree_model_get_iter(v.GTreeModel, &iter.GTreeIter, path.GTreePath))
}

func (v *TreeModel) GetIterFromString(iter *TreeIter, path_string string) bool {
	ptr := C.CString(path_string)
	defer cfree(ptr)
	return gobool(C.gtk_tree_model_get_iter_from_string(v.GTreeModel, &iter.GTreeIter, gstring(ptr)))
}

func (v *TreeModel) GetIterFirst(iter *TreeIter) bool {
	return gobool(C.gtk_tree_model_get_iter_first(v.GTreeModel, &iter.GTreeIter))
}

func (v *TreeModel) GetPath(iter *TreeIter) *TreePath {
	return &TreePath{C._gtk_tree_model_get_path(v.GTreeModel, &iter.GTreeIter)}
}

func (v *TreeModel) GetValue(iter *TreeIter, col int, val *glib.GValue) {
	C.gtk_tree_model_get_value(
		v.GTreeModel, 
		&iter.GTreeIter, 
		gint(col), 
		C.toGValue(unsafe.Pointer(&val.Value))
	)
}

func (v *TreeModel) IterNext(iter *TreeIter) bool {
	return gobool(C.gtk_tree_model_iter_next(v.GTreeModel, &iter.GTreeIter))
}

func (v *TreeModel) IterChildren(iter *TreeIter, parent *TreeIter) bool {
	return gobool(C.gtk_tree_model_iter_children(v.GTreeModel, &iter.GTreeIter, &parent.GTreeIter))
}

func (v *TreeModel) IterHasChild(iter *TreeIter) bool {
	return gobool(C.gtk_tree_model_iter_has_child(v.GTreeModel, &iter.GTreeIter))
}

func (v *TreeModel) IterNChildren(iter *TreeIter) int {
	return int(C.gtk_tree_model_iter_n_children(v.GTreeModel, &iter.GTreeIter))
}

func (v *TreeModel) IterNthChild(iter *TreeIter, parent *TreeIter, n int) bool {
	return gobool(C.gtk_tree_model_iter_nth_child(v.GTreeModel, &iter.GTreeIter, &parent.GTreeIter, gint(n)))
}

func (v *TreeModel) IterParent(iter *TreeIter, child *TreeIter) bool {
	return gobool(C.gtk_tree_model_iter_parent(v.GTreeModel, &iter.GTreeIter, &child.GTreeIter))
}

func (v *TreeModel) GetStringFromIter(i *TreeIter) string {
	return gostring(C.gtk_tree_model_get_string_from_iter(v.GTreeModel, &i.GTreeIter))
}

// gtkTreeSelection
type TreeSelection struct {
	GTreeSelection *C.GtkTreeSelection
}

type SelectionMode int

const (
	SELECTION_NONE     SelectionMode = 0
	SELECTION_SINGLE   SelectionMode = 1
	SELECTION_BROWSE   SelectionMode = 2
	SELECTION_MULTIPLE SelectionMode = 3
	SELECTION_EXTENDED               = SELECTION_MULTIPLE
)

func (v *TreeSelection) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.GTreeSelection)).Connect(s, f, datas...)
}

func (v *TreeSelection) SetMode(m SelectionMode) {
	C.gtk_tree_selection_set_mode(
		v.GTreeSelection, 
		C.GtkSelectionMode(m)
	)
}

func (v *TreeSelection) GetMode() SelectionMode {
	return SelectionMode(C.gtk_tree_selection_get_mode(v.GTreeSelection))
}

func (v *TreeSelection) GetSelected(i *TreeIter) bool {
	return gobool(C.gtk_tree_selection_get_selected(v.GTreeSelection, nil, &i.GTreeIter))
}

func (v *TreeSelection) CountSelectedRows() int {
	return int(C.gtk_tree_selection_count_selected_rows(v.GTreeSelection))
}

func (v *TreeSelection) SelectPath(path *TreePath) {
	C.gtk_tree_selection_select_path(
		v.GTreeSelection, 
		path.GTreePath
	)
}

func (v *TreeSelection) UnselectPath(path *TreePath) {
	C.gtk_tree_selection_unselect_path(
		v.GTreeSelection, 
		path.GTreePath
	)
}

func (v *TreeSelection) PathIsSelected(path *TreePath) bool {
	return gobool(C.gtk_tree_selection_path_is_selected(v.GTreeSelection, path.GTreePath))
}

func (v *TreeSelection) SelectIter(iter *TreeIter) {
	C.gtk_tree_selection_select_iter(v.GTreeSelection, &iter.GTreeIter)
}

func (v *TreeSelection) UnselectIter(iter *TreeIter) {
	C.gtk_tree_selection_unselect_iter(v.GTreeSelection, &iter.GTreeIter)
}

func (v *TreeSelection) IterIsSelected(iter *TreeIter) bool {
	return gobool(C.gtk_tree_selection_iter_is_selected(v.GTreeSelection, &iter.GTreeIter))
}

func (v *TreeSelection) SelectAll() {
	C.gtk_tree_selection_select_all(v.GTreeSelection)
}

func (v *TreeSelection) UnselectAll() {
	C.gtk_tree_selection_unselect_all(v.GTreeSelection)
}

func (v *TreeSelection) SelectRange(start_path *TreePath, end_path *TreePath) {
	C.gtk_tree_selection_select_range(
		v.GTreeSelection, 
		start_path.GTreePath, 
		end_path.GTreePath
	)
}

func (v *TreeSelection) UnselectRange(start_path *TreePath, end_path *TreePath) {
	C.gtk_tree_selection_unselect_range(
		v.GTreeSelection, 
		start_path.GTreePath, 
		end_path.GTreePath
	)
}

// gtkTreeViewColumn
type TreeViewColumnSizing int

const (
	TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = 0
	TREE_VIEW_COLUMN_AUTOSIZE  TreeViewColumnSizing = 1
	TREE_VIEW_COLUMN_FIXED     TreeViewColumnSizing = 2
)

type TreeViewColumn struct {
	GTreeViewColumn *C.GtkTreeViewColumn
	*glib.GObject
}

func newTreeViewColumn(column *C.GtkTreeViewColumn) *TreeViewColumn {
	return &TreeViewColumn{
		GTreeViewColumn: column,
		GObject:         glib.ObjectFromNative(unsafe.Pointer(column)),
	}
}

func NewTreeViewColumn() *TreeViewColumn {
	return newTreeViewColumn(C.gtk_tree_view_column_new())
}

func NewTreeViewColumnWithAttributes2(title string, cell ICellRenderer, attributes ...interface{}) *TreeViewColumn {
	if len(attributes)%2 != 0 {
		log.Panic("Error in gtk.TreeViewColumnWithAttributes: last attribute isn't associated to a value, len(attributes) must be even")
	}
	ptrTitle := C.CString(title)
	defer cfree(ptrTitle)
	ret := newTreeViewColumn(C._gtk_tree_view_column_new_with_attribute(
		gstring(ptrTitle), cell.ToCellRenderer()))
	for i := 0; i < len(attributes)/2; i++ {
		attribute, ok := attributes[2*i].(string)
		if !ok {
			log.Panic("Error calling gtk.TreeViewColumnWithAttributes: property name must be a string")
		}
		ptrAttribute := C.CString(attribute)
		column, ok := attributes[2*i].(int)
		if !ok {
			log.Panic("Error calling gtk.TreeViewColumnWithAttributes: attributes column must be an int")
		}
		C.gtk_tree_view_column_add_attribute(
			ret.GTreeViewColumn,
			cell.ToCellRenderer(), 
			gstring(ptrAttribute), 
			gint(column)
		)
	}
	return ret
}

func NewTreeViewColumnWithAttribute(title string, cell ICellRenderer) *TreeViewColumn {
	ptitle := C.CString(title)
	defer cfree(ptitle)
	return newTreeViewColumn(
		C._gtk_tree_view_column_new_with_attribute(
			gstring(ptitle), 
			cell.ToCellRenderer()
		)
	)
}

func NewTreeViewColumnWithAttributes(title string, cell ICellRenderer, attribute string, column int) *TreeViewColumn {
	ptitle := C.CString(title)
	defer cfree(ptitle)
	pattribute := C.CString(attribute)
	defer cfree(pattribute)
	return newTreeViewColumn(
		C._gtk_tree_view_column_new_with_attributes(
			gstring(ptitle), 
			cell.ToCellRenderer(), 
			gstring(pattribute), 
			gint(column)
		)
	)
}

func (v *TreeViewColumn) PackStart(cell ICellRenderer, expand bool) {
	C.gtk_tree_view_column_pack_start(
		v.GTreeViewColumn, 
		cell.ToCellRenderer(), 
		gbool(expand)
	)
}

func (v *TreeViewColumn) PackEnd(cell ICellRenderer, expand bool) {
	C.gtk_tree_view_column_pack_end(
		v.GTreeViewColumn, 
		cell.ToCellRenderer(), 
		gbool(expand)
	)
}

func (v *TreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear(v.GTreeViewColumn)
}

func (v *TreeViewColumn) AddAttribute(cell ICellRenderer, attribute string, column int) {
	ptr := C.CString(attribute)
	defer cfree(ptr)
	C.gtk_tree_view_column_add_attribute(
		v.GTreeViewColumn, 
		cell.ToCellRenderer(), 
		gstring(ptr), 
		gint(column)
	)
}

func (v *TreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(
		v.GTreeViewColumn, 
		gint(spacing)
	)
}

func (v *TreeViewColumn) GetSpacing() int {
	return int(C.gtk_tree_view_column_get_spacing(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetVisible(resizable bool) {
	C.gtk_tree_view_column_set_visible(
		v.GTreeViewColumn, 
		gbool(resizable)
	)
}

func (v *TreeViewColumn) GetVisible() bool {
	return gobool(C.gtk_tree_view_column_get_visible(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetResizable(resizable bool) {
	C.gtk_tree_view_column_set_resizable(
		v.GTreeViewColumn, 
		gbool(resizable)
	)
}

func (v *TreeViewColumn) GetResizable() bool {
	return gobool(C.gtk_tree_view_column_get_resizable(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetSizing(s TreeViewColumnSizing) {
	C.gtk_tree_view_column_set_sizing(
		v.GTreeViewColumn, 
		C.GtkTreeViewColumnSizing(s)
	)
}

func (v *TreeViewColumn) GetSizing() TreeViewColumnSizing {
	return TreeViewColumnSizing(C.gtk_tree_view_column_get_sizing(v.GTreeViewColumn))
}

func (v *TreeViewColumn) GetWidth() int {
	return int(C.gtk_tree_view_column_get_width(v.GTreeViewColumn))
}

func (v *TreeViewColumn) GetFixedWidth() int {
	return int(C.gtk_tree_view_column_get_fixed_width(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetFixedWidth(w int) {
	C.gtk_tree_view_column_set_fixed_width(
		v.GTreeViewColumn, 
		gint(w)
	)
}

func (v *TreeViewColumn) SetMinWidth(w int) {
	C.gtk_tree_view_column_set_min_width(
		v.GTreeViewColumn, 
		gint(w)
	)
}

func (v *TreeViewColumn) GetMinWidth() int {
	return int(C.gtk_tree_view_column_get_min_width(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetMaxWidth(w int) {
	C.gtk_tree_view_column_set_max_width(
		v.GTreeViewColumn, 
		gint(w)
	)
}

func (v *TreeViewColumn) GetMaxWidth() int {
	return int(C.gtk_tree_view_column_get_max_width(v.GTreeViewColumn))
}

func (v *TreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked(v.GTreeViewColumn)
}

func (v *TreeViewColumn) SetTitle(title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_tree_view_column_set_title(
		v.GTreeViewColumn, 
		gstring(ptr)
	)
}

func (v *TreeViewColumn) GetTitle() string {
	return gostring(C.gtk_tree_view_column_get_title(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetExpand(expand bool) {
	C.gtk_tree_view_column_set_expand(
		v.GTreeViewColumn, 
		gbool(expand)
	)
}

func (v *TreeViewColumn) GetExpand() bool {
	return gobool(C.gtk_tree_view_column_get_expand(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetClickable(clickable bool) {
	C.gtk_tree_view_column_set_clickable(
		v.GTreeViewColumn, 
		gbool(clickable)
	)
}

func (v *TreeViewColumn) GetClickable() bool {
	return gobool(C.gtk_tree_view_column_get_clickable(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetReorderable(reorderable bool) {
	C.gtk_tree_view_column_set_reorderable(
		v.GTreeViewColumn, 
		gbool(reorderable)
	)
}

func (v *TreeViewColumn) GetReorderable() bool {
	return gobool(C.gtk_tree_view_column_get_reorderable(v.GTreeViewColumn))
}

// gtkTreeView
type TreeView struct {
	Container
}

func NewTreeView() *TreeView {
	return &TreeView{Container{Widget{C.gtk_tree_view_new()}}}
}

func (v *TreeView) SetModel(model ITreeModel) {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	C.gtk_tree_view_set_model(
		TREE_VIEW(v), 
		tm
	)
}

func (v *TreeView) GetSelection() *TreeSelection {
	return &TreeSelection{C.gtk_tree_view_get_selection(TREE_VIEW(v))}
}

func (v *TreeView) SetHeadersVisible(flag bool) {
	C.gtk_tree_view_set_headers_visible(
		TREE_VIEW(v), 
		gbool(flag)
	)
}

func (v *TreeView) AppendColumn(c *TreeViewColumn) int {
	return int(C.gtk_tree_view_append_column(TREE_VIEW(v), c.GTreeViewColumn))
}

func (v *TreeView) GetColumn(n int) *TreeViewColumn {
	return newTreeViewColumn(C.gtk_tree_view_get_column(TREE_VIEW(v), gint(n)))
}

func (v *TreeView) ScrollToCell(path *TreePath, col *TreeViewColumn, use bool, ra, ca float64) {
	var pcol *C.GtkTreeViewColumn
	if nil == col {
		pcol = nil
	} else {
		pcol = col.GTreeViewColumn
	}
	C.gtk_tree_view_scroll_to_cell(
		TREE_VIEW(v), 
		path.GTreePath, 
		pcol, 
		gbool(use), 
		C.gfloat(ra), 
		C.gfloat(ca)
	)
}

func (v *TreeView) SetCursor(path *TreePath, col *TreeViewColumn, se bool) {
	var pcol *C.GtkTreeViewColumn
	if nil == col {
		pcol = nil
	} else {
		pcol = col.GTreeViewColumn
	}
	C.gtk_tree_view_set_cursor(
		TREE_VIEW(v), 
		path.GTreePath, 
		pcol, 
		gbool(se)
	)
}

func (v *TreeView) GetCursor(path **TreePath, focus_column **TreeViewColumn) {
	*path = &TreePath{nil}
	if nil != focus_column {
		*focus_column = &TreeViewColumn{nil, nil}
		C.gtk_tree_view_get_cursor(
			TREE_VIEW(v), 
			&(*path).GTreePath, 
			&(*focus_column).GTreeViewColumn
		)
	} else {
		C.gtk_tree_view_get_cursor(
			TREE_VIEW(v), 
			&(*path).GTreePath, 
			nil
		)
	}
}

func (v *TreeView) ExpandAll() {
	C.gtk_tree_view_expand_all(TREE_VIEW(v))
}

func (v *TreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(TREE_VIEW(v))
}

func (v *TreeView) ExpandRow(path *TreePath, openall bool) bool {
	return gobool(C.gtk_tree_view_expand_row(TREE_VIEW(v), path.GTreePath, gbool(openall)))
}

func (v *TreeView) CollapseRow(path *TreePath) bool {
	return gobool(C.gtk_tree_view_collapse_row(TREE_VIEW(v), path.GTreePath))
}

func (v *TreeView) RowExpanded(path *TreePath) bool {
	return gobool(C.gtk_tree_view_row_expanded(TREE_VIEW(v), path.GTreePath))
}

// gtkIconView
type IconView struct {
	Container
}

func NewIconView() *IconView {
	return &IconView{Container{Widget{
		C.gtk_icon_view_new()}}}
}

func NewIconViewWithModel(model ITreeModel) *IconView {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	return &IconView{Container{Widget{C.gtk_icon_view_new_with_model(tm)}}}
}

func (v *IconView) GetModel() *TreeModel {
	return &TreeModel{C.gtk_icon_view_get_model(ICON_VIEW(v))}
}

func (v *IconView) SetModel(model ITreeModel) {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	C.gtk_icon_view_set_model(
		ICON_VIEW(v), 
		tm
	)
}

func (v *IconView) GetTextColumn() int {
	return int(C.gtk_icon_view_get_text_column(ICON_VIEW(v)))
}

func (v *IconView) SetTextColumn(text_column int) {
	C.gtk_icon_view_set_text_column(
		ICON_VIEW(v), 
		gint(text_column)
	)
}

func (v *IconView) GetMarkupColumn() int {
	return int(C.gtk_icon_view_get_markup_column(ICON_VIEW(v)))
}

func (v *IconView) SetMarkupColumn(markup_column int) {
	C.gtk_icon_view_set_markup_column(
		ICON_VIEW(v), 
		gint(markup_column)
	)
}

func (v *IconView) GetPixbufColumn() int {
	return int(C.gtk_icon_view_get_pixbuf_column(ICON_VIEW(v)))
}

func (v *IconView) SetPixbufColumn(pixbuf_column int) {
	C.gtk_icon_view_set_pixbuf_column(
		ICON_VIEW(v), 
		gint(pixbuf_column)
	)
}

func (v *IconView) ScrollToPath(path *TreePath, use bool, ra float64, ca float64) {
	C.gtk_icon_view_scroll_to_path(
		ICON_VIEW(v), 
		path.GTreePath,
		gbool(use), 
		C.gfloat(ra), 
		C.gfloat(ca)
	)
}

// gtkCellRenderer
type ICellRenderer interface {
	ToCellRenderer() *C.GtkCellRenderer
}

type CellRenderer struct {
	GCellRenderer *C.GtkCellRenderer
	ICellRenderer
}

func (v *CellRenderer) ToCellRenderer() *C.GtkCellRenderer {
	return v.GCellRenderer
}

func (v *CellRenderer) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.GCellRenderer)).Connect(s, f, datas...)
}

// gtkCellRendererAccel
type CellRendererAccel struct {
	CellRenderer
}

func NewCellRendererAccel() *CellRendererAccel {
	return &CellRendererAccel{CellRenderer{C.gtk_cell_renderer_accel_new(), nil}}
}

// gtkCellRendererCombo
type CellRendererCombo struct {
	CellRenderer
}

func NewCellRendererCombo() *CellRendererCombo {
	return &CellRendererCombo{CellRenderer{C.gtk_cell_renderer_combo_new(), nil}}
}

// gtkCellRendererPixbuf
type CellRendererPixbuf struct {
	CellRenderer
}

func NewCellRendererPixbuf() *CellRendererPixbuf {
	return &CellRendererPixbuf{CellRenderer{C.gtk_cell_renderer_pixbuf_new(), nil}}
}

// gtkCellRendererProgress
type CellRendererProgress struct {
	CellRenderer
}

func NewCellRendererProgress() *CellRendererProgress {
	return &CellRendererProgress{CellRenderer{C.gtk_cell_renderer_progress_new(), nil}}
}

// gtkCellRendererSpin
type CellRendererSpin struct {
	CellRenderer
}

func NewCellRendererSpin() *CellRendererSpin {
	return &CellRendererSpin{CellRenderer{C.gtk_cell_renderer_spin_new(), nil}}
}

// gtkCellRendererText
type CellRendererText struct {
	CellRenderer
}

func NewCellRendererText() *CellRendererText {
	return &CellRendererText{CellRenderer{C.gtk_cell_renderer_text_new(), nil}}
}
func (v *CellRendererText) SetFixedHeightFromFont(number_of_rows int) {
	C.gtk_cell_renderer_text_set_fixed_height_from_font(
		CELL_RENDERER_TEXT(v), 
		gint(number_of_rows)
	)
}

// gtkCellRendererToggle
type CellRendererToggle struct {
	CellRenderer
}

func NewCellRendererToggle() *CellRendererToggle {
	return &CellRendererToggle{CellRenderer{C.gtk_cell_renderer_toggle_new(), nil}}
}

func (v *CellRendererToggle) GetRadio() bool {
	return gobool(C.gtk_cell_renderer_toggle_get_radio(CELL_RENDERER_TOGGLE(v)))
}

func (v *CellRendererToggle) SetRadio(radio bool) {
	C.gtk_cell_renderer_toggle_set_radio(
		CELL_RENDERER_TOGGLE(v), 
		gbool(radio)
	)
}

func (v *CellRendererToggle) GetActive() bool {
	return gobool(C.gtk_cell_renderer_toggle_get_active(CELL_RENDERER_TOGGLE(v)))
}

func (v *CellRendererToggle) SetActive(active bool) {
	C.gtk_cell_renderer_toggle_set_active(
		CELL_RENDERER_TOGGLE(v), 
		gbool(active)
	)
}

func (v *CellRendererToggle) GetActivatable() bool {
	panic_if_version_older(2, 18, 0, "gtk_cell_renderer_toggle_get_activatable()")
	return gobool(C._gtk_cell_renderer_toggle_get_activatable(CELL_RENDERER_TOGGLE(v)))
}

func (v *CellRendererToggle) SetActivatable(activatable bool) {
	panic_if_version_older(2, 18, 0, "gtk_cell_renderer_toggle_set_activatable()")
	C._gtk_cell_renderer_toggle_set_activatable(
		CELL_RENDERER_TOGGLE(v), 
		gbool(activatable)
	)
}

// gtkCellRendererSpinner
type CellRendererSpinner struct {
	CellRenderer
}

func NewCellRendererSpinner() *CellRendererSpinner {
	panic_if_version_older(2, 20, 0, "gtk_cell_renderer_spinner_new()")
	return &CellRendererSpinner{CellRenderer{C._gtk_cell_renderer_spinner_new(), nil}}
}

// gtkListStore
const (
	TYPE_CHAR    = glib.G_TYPE_CHAR
	TYPE_UCHAR   = glib.G_TYPE_UCHAR
	TYPE_BOOL    = glib.G_TYPE_BOOL
	TYPE_INT     = glib.G_TYPE_INT
	TYPE_UINT    = glib.G_TYPE_UINT
	TYPE_LONG    = glib.G_TYPE_LONG
	TYPE_ULONG   = glib.G_TYPE_ULONG
	TYPE_FLOAT   = glib.G_TYPE_FLOAT
	TYPE_DOUBLE  = glib.G_TYPE_DOUBLE
	TYPE_STRING  = glib.G_TYPE_STRING
	TYPE_BOXED   = glib.G_TYPE_BOXED
	TYPE_POINTER = glib.G_TYPE_POINTER
	TYPE_PIXBUF  = TYPE_POINTER
)

type ListStore struct {
	TreeModel
	GListStore *C.GtkListStore
}

func NewListStore(v ...interface{}) *ListStore {
	types := C.make_gtypes(C.int(len(v)))
	for n := range v {
		C.set_gtype(
			types, 
			C.int(n), 
			C.int(v[n].(int))
		)
	}
	defer C.destroy_gtypes(types)
	cliststore := C.gtk_list_store_newv(gint(len(v)), types)
	return &ListStore{TreeModel{C.toGTreeModelFromListStore(cliststore)}, cliststore}
}

func (v *ListStore) Set(iter *TreeIter, a ...interface{}) {
	for r := range a {
		v.SetValue(iter, r, a[r])
	}
}

func (v *ListStore) SetValue(iter *TreeIter, column int, a interface{}) {
	gv := glib.GValueFromNative(a)
	if gv != nil {
		C.gtk_list_store_set_value(
			v.GListStore, 
			&iter.GTreeIter, 
			gint(column), 
			C.toGValue(unsafe.Pointer(gv))
		)
	} else {
		if pv, ok := a.(*[0]uint8); ok {
			C._gtk_list_store_set_ptr(
				v.GListStore, 
				&iter.GTreeIter, 
				gint(column), 
				unsafe.Pointer(pv)
			)
		} else {
			av := reflect.ValueOf(a)
			if av.CanAddr() {
				C._gtk_list_store_set_addr(
					v.GListStore, 
					&iter.GTreeIter, 
					gint(column), 
					unsafe.Pointer(av.UnsafeAddr())
				)
			} else {
				C._gtk_list_store_set_addr(
					v.GListStore, 
					&iter.GTreeIter, 
					gint(column), 
					unsafe.Pointer(&a)
				)
			}
		}
	}
}

func (v *ListStore) Remove(iter *TreeIter) bool {
	return gobool(C.gtk_list_store_remove(v.GListStore, &iter.GTreeIter))
}

func (v *ListStore) Insert(iter *TreeIter, position int) {
	C.gtk_list_store_insert(
		v.GListStore, 
		&iter.GTreeIter, 
		gint(position)
	)
}

func (v *ListStore) InsertBefore(iter *TreeIter, sibling *TreeIter) {
	C.gtk_list_store_insert_before(
		v.GListStore, 
		&iter.GTreeIter, 
		&sibling.GTreeIter
	)
}

func (v *ListStore) InsertAfter(iter *TreeIter, sibling *TreeIter) {
	C.gtk_list_store_insert_after(
		v.GListStore, 
		&iter.GTreeIter, 
		&sibling.GTreeIter
	)
}

func (v *ListStore) Prepend(iter *TreeIter) {
	C.gtk_list_store_prepend(
		v.GListStore, 
		&iter.GTreeIter
	)
}

func (v *ListStore) Append(iter *TreeIter) {
	C.gtk_list_store_append(
		v.GListStore, 
		&iter.GTreeIter
	)
}

func (v *ListStore) Clear() {
	C.gtk_list_store_clear(v.GListStore)
}

func (v *ListStore) IterIsValid(iter *TreeIter) bool {
	log.Println("Attention: ListStore.IterIsValid: You should only use this function for debugging or testing stages.")
	return gobool(C.gtk_list_store_iter_is_valid(v.GListStore, &iter.GTreeIter))
}

func (v *ListStore) Reorder(i *int) {
	gi := gint(*i)
	C.gtk_list_store_reorder(v.GListStore, &gi)
	*i = int(gi)
}

func (v *ListStore) Swap(a *TreeIter, b *TreeIter) {
	C.gtk_list_store_swap(v.GListStore, &a.GTreeIter, &b.GTreeIter)
}

func (v *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	C.gtk_list_store_move_before(
		v.GListStore, 
		&iter.GTreeIter, 
		&position.GTreeIter
	)
}

func (v *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	C.gtk_list_store_move_after(
		v.GListStore, 
		&iter.GTreeIter, 
		&position.GTreeIter
	)
}

// gtkTreeStore
type TreeStore struct {
	TreeModel
	GTreeStore *C.GtkTreeStore
}

func NewTreeStore(v ...interface{}) *TreeStore {
	types := C.make_gtypes(C.int(len(v)))
	for n := range v {
		C.set_gtype(types, C.int(n), C.int(v[n].(int)))
	}
	defer C.destroy_gtypes(types)
	ctreestore := C.gtk_tree_store_newv(gint(len(v)), types)
	return &TreeStore{TreeModel{C.toGTreeModelFromTreeStore(ctreestore)}, ctreestore}
}

func (v *TreeStore) Set(iter *TreeIter, a ...interface{}) {
	for r := range a {
		v.SetValue(iter, r, a[r])
	}
}

func (v *TreeStore) SetValue(iter *TreeIter, column int, a interface{}) {
	gv := glib.GValueFromNative(a)
	if gv != nil {
		C.gtk_tree_store_set_value(
			v.GTreeStore, 
			&iter.GTreeIter, 
			gint(column), 
			C.toGValue(unsafe.Pointer(gv))
		)
	} else {
		if pv, ok := a.(*[0]uint8); ok {
			C._gtk_tree_store_set_ptr(
				v.GTreeStore, 
				&iter.GTreeIter, 
				gint(column), 
				unsafe.Pointer(pv)
			)
		} else {
			av := reflect.ValueOf(a)
			if av.CanAddr() {
				C._gtk_tree_store_set_addr(
					v.GTreeStore, 
					&iter.GTreeIter, 
					gint(column), 
					unsafe.Pointer(av.UnsafeAddr())
				)
			} else {
				C._gtk_tree_store_set_addr(
					v.GTreeStore, 
					&iter.GTreeIter, 
					gint(column), 
					unsafe.Pointer(&a)
				)
			}
		}
	}
}

func (v *TreeStore) Remove(iter *TreeIter) bool {
	return gobool(C.gtk_tree_store_remove(v.GTreeStore, &iter.GTreeIter))
}

func (v *TreeStore) Insert(iter *TreeIter, parent *TreeIter, position int) {
	C.gtk_tree_store_insert(
		v.GTreeStore, 
		&iter.GTreeIter, 
		&parent.GTreeIter, 
		gint(position)
	)
}

func (v *TreeStore) InsertBefore(iter *TreeIter, parent *TreeIter, sibling *TreeIter) {
	C.gtk_tree_store_insert_before(
		v.GTreeStore, 
		&iter.GTreeIter, 
		&parent.GTreeIter, 
		&sibling.GTreeIter
	)
}

func (v *TreeStore) InsertAfter(iter *TreeIter, parent *TreeIter, sibling *TreeIter) {
	C.gtk_tree_store_insert_after(
		v.GTreeStore, 
		&iter.GTreeIter, 
		&parent.GTreeIter, 
		&sibling.GTreeIter
	)
}

func (v *TreeStore) Prepend(iter *TreeIter, parent *TreeIter) {
	if parent == nil {
		C.gtk_tree_store_prepend(
			v.GTreeStore, 
			&iter.GTreeIter, 
			nil
		)
	} else {
		C.gtk_tree_store_prepend(
			v.GTreeStore, 
			&iter.GTreeIter, 
			&parent.GTreeIter
		)
	}
}

func (v *TreeStore) Append(iter *TreeIter, parent *TreeIter) {
	if parent == nil {
		C.gtk_tree_store_append(
			v.GTreeStore, 
			&iter.GTreeIter, 
			nil
		)
	} else {
		C.gtk_tree_store_append(
			v.GTreeStore, 
			&iter.GTreeIter, 
			&parent.GTreeIter
		)
	}
}

func (v *TreeStore) ToTreeModel() *TreeModel {
	return &TreeModel{C.toGTreeModelFromTreeStore(v.GTreeStore)}
}

func (v *TreeStore) IterDepth(iter *TreeIter) int {
	return int(C.gtk_tree_store_iter_depth(v.GTreeStore, &iter.GTreeIter))
}

func (v *TreeStore) Clear() {
	C.gtk_tree_store_clear(v.GTreeStore)
}

func (v *TreeStore) IterIsValid(iter *TreeIter) bool {
	log.Println("Attention: TreeStore.IterIsValid: You should only use this function for debugging or testing stages.")
	return gobool(C.gtk_tree_store_iter_is_valid(v.GTreeStore, &iter.GTreeIter))
}

func (v *TreeStore) Reorder(iter *TreeIter, i *int) {
	gi := gint(*i)
	C.gtk_tree_store_reorder(
		v.GTreeStore, 
		&iter.GTreeIter, 
		&gi
	)
	*i = int(gi)
}

func (v *TreeStore) Swap(a *TreeIter, b *TreeIter) {
	C.gtk_tree_store_swap(
		v.GTreeStore, 
		&a.GTreeIter, 
		&b.GTreeIter
	)
}

func (v *TreeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	C.gtk_tree_store_move_before(
		v.GTreeStore, 
		&iter.GTreeIter, 
		&position.GTreeIter
	)
}

func (v *TreeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	C.gtk_tree_store_move_after(
		v.GTreeStore, 
		&iter.GTreeIter, 
		&position.GTreeIter
	)
}

// gtkComboBox
type ComboBox struct {
	Bin
}

func NewComboBox() *ComboBox {
	return &ComboBox{Bin{Container{Widget{C.gtk_combo_box_new()}}}}
}

func NewComboBoxWithEntry() *ComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_with_entry()")
	return &ComboBox{Bin{Container{Widget{C._gtk_combo_box_new_with_entry()}}}}
}

func NewComboBoxWithModel(model *TreeModel) *ComboBox {
	return &ComboBox{Bin{Container{Widget{C.gtk_combo_box_new_with_model(model.GTreeModel)}}}}
}

func NewComboBoxWithModelAndEntry(model *TreeModel) *ComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_with_model_and_entry()")
	return &ComboBox{Bin{Container{Widget{C._gtk_combo_box_new_with_model_and_entry(model.GTreeModel)}}}}
}

func (v *ComboBox) GetWrapWidth() int {
	return int(C.gtk_combo_box_get_wrap_width(COMBO_BOX(v)))
}

func (v *ComboBox) SetWrapWidth(width int) {
	C.gtk_combo_box_set_wrap_width(
		COMBO_BOX(v), 
		gint(width)
	)
}

func (v *ComboBox) GetRowSpanColumn() int {
	return int(C.gtk_combo_box_get_row_span_column(COMBO_BOX(v)))
}

func (v *ComboBox) SetRowSpanColumn(row_span int) {
	C.gtk_combo_box_set_row_span_column(
		COMBO_BOX(v), 
		gint(row_span)
	)
}

func (v *ComboBox) GetColumnSpanColumn() int {
	return int(C.gtk_combo_box_get_column_span_column(COMBO_BOX(v)))
}

func (v *ComboBox) SetColumnSpanColumn(column_span int) {
	C.gtk_combo_box_set_column_span_column(
		COMBO_BOX(v), 
		gint(column_span)
	)
}

func (v *ComboBox) GetActive() int {
	return int(C.gtk_combo_box_get_active(COMBO_BOX(v)))
}

func (v *ComboBox) SetActive(width int) {
	C.gtk_combo_box_set_active(
		COMBO_BOX(v), 
		gint(width)
	)
}

func (v *ComboBox) GetActiveIter(iter *TreeIter) bool {
	return gobool(C.gtk_combo_box_get_active_iter(COMBO_BOX(v), &iter.GTreeIter))
}

func (v *ComboBox) SetActiveIter(iter *TreeIter) {
	C.gtk_combo_box_set_active_iter(
		COMBO_BOX(v), 
		&iter.GTreeIter
	)
}

func (v *ComboBox) GetModel() *TreeModel {
	return &TreeModel{
		C.gtk_combo_box_get_model(COMBO_BOX(v))}
}

func (v *ComboBox) SetModel(model *TreeModel) {
	C.gtk_combo_box_set_model(
		COMBO_BOX(v), 
		model.GTreeModel
	)
}

func NewComboBoxNewText() *ComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_text()")
	return &ComboBox{Bin{Container{Widget{C.gtk_combo_box_new_text()}}}}
}

func (v *ComboBox) AppendText(text string) {
	deprecated_since(2, 24, 0, "gtk_combo_box_append_text()")
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_combo_box_append_text(
		COMBO_BOX(v), 
		gstring(ptr)
	)
}

func (v *ComboBox) InsertText(text string, position int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_insert_text()")
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_combo_box_insert_text(
		COMBO_BOX(v), 
		gint(position), 
		gstring(ptr)
	)
}

func (v *ComboBox) PrependText(text string) {
	deprecated_since(2, 24, 0, "gtk_combo_box_prepend_text()")
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_combo_box_prepend_text(
		COMBO_BOX(v), 
		gstring(ptr)
	)
}

func (v *ComboBox) RemoveText(position int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_remove_text()")
	C.gtk_combo_box_remove_text(
		COMBO_BOX(v), 
		gint(position)
	)
}

func (v *ComboBox) GetActiveText() string {
	deprecated_since(2, 24, 0, "gtk_combo_box_get_active_text()")
	return gostring(C.gtk_combo_box_get_active_text(COMBO_BOX(v)))
}

func (v *ComboBox) Popup() {
	C.gtk_combo_box_popup(COMBO_BOX(v))
}

func (v *ComboBox) Popdown() {
	C.gtk_combo_box_popdown(COMBO_BOX(v))
}

func (v *ComboBox) SetAddTearoffs(add_tearoffs bool) {
	C.gtk_combo_box_set_add_tearoffs(
		COMBO_BOX(v), 
		gbool(add_tearoffs)
	)
}

func (v *ComboBox) GetAddTearoffs() bool {
	return gobool(C.gtk_combo_box_get_add_tearoffs(COMBO_BOX(v)))
}

func (v *ComboBox) SetTitle(title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_combo_box_set_title(
		COMBO_BOX(v), 
		gstring(ptr)
	)
}

func (v *ComboBox) GetTitle() string {
	return gostring(C.gtk_combo_box_get_title(COMBO_BOX(v)))
}

func (v *ComboBox) SetFocusOnClick(focus_on_click bool) {
	C.gtk_combo_box_set_focus_on_click(
		COMBO_BOX(v), 
		gbool(focus_on_click)
	)
}

func (v *ComboBox) GetFocusOnClick() bool {
	return gobool(C.gtk_combo_box_get_focus_on_click(COMBO_BOX(v)))
}

// gtkComboBoxText
type ComboBoxText struct {
	ComboBox
}

func NewComboBoxText() *ComboBoxText {
	return &ComboBoxText{ComboBox{Bin{Container{Widget{C._gtk_combo_box_text_new()}}}}}
}

func NewComboBoxTextWithEntry() *ComboBoxText {
	return &ComboBoxText{ComboBox{Bin{Container{Widget{C._gtk_combo_box_text_new_with_entry()}}}}}
}

func (v *ComboBoxText) AppendText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C._gtk_combo_box_text_append_text(
		COMBO_BOX_TEXT(v), 
		gstring(ptr)
	)
}

func (v *ComboBoxText) InsertText(position int, text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C._gtk_combo_box_text_insert_text(
		COMBO_BOX_TEXT(v), 
		gint(position), 
		gstring(ptr)
	)
}

func (v *ComboBoxText) PrependText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C._gtk_combo_box_text_prepend_text(
		COMBO_BOX_TEXT(v), 
		gstring(ptr)
	)
}

func (v *ComboBoxText) Remove(position int) {
	C._gtk_combo_box_text_remove(
		COMBO_BOX_TEXT(v), 
		gint(position)
	)
}

func (v *ComboBoxText) GetActiveText() string {
	return gostring(C._gtk_combo_box_text_get_active_text(COMBO_BOX_TEXT(v)))
}

// gtkComboBoxEntry
type ComboBoxEntry struct {
	ComboBox
}

func NewComboBoxEntry() *ComboBoxEntry {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_new()")
	return &ComboBoxEntry{ComboBox{Bin{Container{Widget{C.gtk_combo_box_entry_new()}}}}}
}

func NewComboBoxEntryNewText() *ComboBoxEntry {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_new_text()")
	return &ComboBoxEntry{ComboBox{Bin{Container{Widget{C.gtk_combo_box_entry_new_text()}}}}}
}

func (v *ComboBoxEntry) GetTextColumn() int {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_get_text_column()")
	return int(C.gtk_combo_box_entry_get_text_column(COMBO_BOX_ENTRY(v)))
}

func (v *ComboBoxEntry) SetTextColumn(text_column int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_set_text_column()")
	C.gtk_combo_box_entry_set_text_column(
		COMBO_BOX_ENTRY(v), 
		gint(text_column)
	)
}

// gtkMenu
type Menu struct {
	Container
}

func NewMenu() *Menu {
	return &Menu{Container{Widget{C.gtk_menu_new()}}}
}

func (v *Menu) Append(child IWidget) {
	C.gtk_menu_shell_append(
		MENU_SHELL(v), 
		ToNative(child)
	)
}

func (v *Menu) Prepend(child IWidget) {
	C.gtk_menu_shell_prepend(
		MENU_SHELL(v), 
		ToNative(child)
	)
}

func (v *Menu) Insert(child IWidget, position int) {
	C.gtk_menu_shell_insert(
		MENU_SHELL(v), 
		ToNative(child), 
		gint(position)
	)
}

func (v *Menu) Popup(parent_menu_shell, parent_menu_item IWidget, f MenuPositionFunc, data interface{}, button uint, active_item uint32) {
	var pms, pmi *C.GtkWidget
	if parent_menu_shell != nil {
		pms = ToNative(parent_menu_shell)
	}
	if parent_menu_item != nil {
		pmi = ToNative(parent_menu_item)
	}
	C._gtk_menu_popup(
		v.GWidget, 
		pms, 
		pmi, 
		unsafe.Pointer(&MenuPositionFuncInfo{v, f, data}), 
		guint(button), 
		guint32(active_item)
	)
}

func (v *Menu) GetTearoffState() bool {
	return gobool(C.gtk_menu_get_tearoff_state(MENU(v)))
}

func (v *Menu) SetReserveToggleSize(b bool) {
	panic_if_version_older(2, 18, 0, "gtk_menu_set_reserve_toggle_size()")
	C._gtk_menu_set_reserve_toggle_size(
		MENU(v), 
		gbool(b)
	)
}

func (v *Menu) GetReserveToggleSize() bool {
	panic_if_version_older(2, 18, 0, "gtk_menu_get_reserve_toggle_size()")
	return gobool(C._gtk_menu_get_reserve_toggle_size(MENU(v)))
}

func (v *Menu) Popdown() {
	C.gtk_menu_popdown(MENU(v))
}

func (v *Menu) Reposition() {
	C.gtk_menu_reposition(MENU(v))
}

func (v *Menu) GetActive() *Widget {
	return &Widget{C.gtk_menu_get_active(MENU(v))}
}

func (v *Menu) SetTearoffState(b bool) {
	C.gtk_menu_set_tearoff_state(
		MENU(v), 
		gbool(b)
	)
}

func (v *Menu) Detach() {
	C.gtk_menu_detach(MENU(v))
}

func (v *Menu) GetAttachWidget() *Widget {
	return &Widget{C.gtk_menu_get_attach_widget(MENU(v))}
}

type MenuPositionFunc func(menu *Menu, px, py *int, push_in *bool, data interface{})

type MenuPositionFuncInfo struct {
	menu *Menu
	f    MenuPositionFunc
	data interface{}
}

func _go_gtk_menu_position_func(gmpfi *C._gtk_menu_position_func_info) {
	if gmpfi == nil {
		return
	}
	gmpfigo := (*MenuPositionFuncInfo)(gmpfi.data)
	if gmpfigo.f == nil {
		return
	}
	x := int(gmpfi.x)
	y := int(gmpfi.y)
	push_in := gobool(gmpfi.push_in)
	gmpfigo.f(gmpfigo.menu, &x, &y, &push_in, gmpfigo.data)
	gmpfi.x = gint(x)
	gmpfi.y = gint(y)
	gmpfi.push_in = gbool(push_in)
}

// gtkMenuBar
type PackDirection int

const (
	PACK_DIRECTION_LTR PackDirection = 0
	PACK_DIRECTION_RTL PackDirection = 1
	PACK_DIRECTION_TTB PackDirection = 2
	PACK_DIRECTION_BTT PackDirection = 3
)

type MenuBar struct {
	Widget
}

func NewMenuBar() *MenuBar {
	return &MenuBar{Widget{C.gtk_menu_bar_new()}}
}

func (v *MenuBar) SetPackDirection(pack_dir PackDirection) {
	C.gtk_menu_bar_set_pack_direction(
		MENU_BAR(v), 
		C.GtkPackDirection(pack_dir)
	)
}

func (v *MenuBar) GetPackDirection() PackDirection {
	return PackDirection(C.gtk_menu_bar_get_pack_direction(MENU_BAR(v)))
}

func (v *MenuBar) SetChildPackDirection(pack_dir PackDirection) {
	C.gtk_menu_bar_set_child_pack_direction(
		MENU_BAR(v), 
		C.GtkPackDirection(pack_dir)
	)
}

func (v *MenuBar) GetChildPackDirection() PackDirection {
	return PackDirection(C.gtk_menu_bar_get_child_pack_direction(MENU_BAR(v)))
}

func (v *MenuBar) Append(child IWidget) {
	C.gtk_menu_shell_append(
		MENU_BAR_SHELL(v), 
		ToNative(child)
	)
}

func (v *MenuBar) Prepend(child IWidget) {
	C.gtk_menu_shell_prepend(
		MENU_BAR_SHELL(v), 
		ToNative(child)
	)
}

func (v *MenuBar) Insert(child IWidget, position int) {
	C.gtk_menu_shell_insert(
		MENU_BAR_SHELL(v), 
		ToNative(child), 
		gint(position)
	)
}

// gtkMenuItem
type MenuItem struct {
	Item
}

func NewMenuItem() *MenuItem {
	return &MenuItem{Item{Bin{Container{Widget{C.gtk_menu_item_new()}}}}}
}

func NewMenuItemWithLabel(label string) *MenuItem {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &MenuItem{Item{Bin{Container{Widget{
		C.gtk_menu_item_new_with_label(gstring(ptr))}}}}}
}

func NewMenuItemWithMnemonic(label string) *MenuItem {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &MenuItem{Item{Bin{Container{Widget{
		C.gtk_menu_item_new_with_mnemonic(gstring(ptr))}}}}}
}

func (v *MenuItem) SetRightJustified(b bool) {
	C.gtk_menu_item_set_right_justified(
		MENU_ITEM(v), 
		gbool(b)
	)
}

func (v *MenuItem) GetRightJustified() bool {
	return gobool(C.gtk_menu_item_get_right_justified(MENU_ITEM(v)))
}

func (v *MenuItem) GetUseUnderline() bool {
	return gobool(C.gtk_menu_item_get_use_underline(MENU_ITEM(v)))
}

func (v *MenuItem) SetUseUnderline(setting bool) {
	C.gtk_menu_item_set_use_underline(
		MENU_ITEM(v), 
		gbool(setting)
	)
}

func (v *MenuItem) SetSubmenu(w IWidget) {
	C.gtk_menu_item_set_submenu(
		MENU_ITEM(v), 
		ToNative(w)
	)
}

func (v *MenuItem) GetSubmenu() *Widget {
	return &Widget{C.gtk_menu_item_get_submenu(MENU_ITEM(v))}
}

func (v *MenuItem) RemoveSubmenu() {
	deprecated_since(2, 12, 0, "gtk_menu_item_remove_submenu()")
	C.gtk_menu_item_remove_submenu(MENU_ITEM(v))
}

func (v *MenuItem) Select() {
	C.gtk_menu_item_select(MENU_ITEM(v))
}

func (v *MenuItem) Deselect() {
	C.gtk_menu_item_deselect(MENU_ITEM(v))
}

func (v *MenuItem) Activate() {
	C.gtk_menu_item_activate(MENU_ITEM(v))
}

func (v *MenuItem) ToggleSizeRequest(i *int) {
	gi := gint(*i)
	C.gtk_menu_item_toggle_size_request(
		MENU_ITEM(v), 
		&gi
	)
	*i = int(gi)
}

func (v *MenuItem) ToggleSizeAllocate(i int) {
	C.gtk_menu_item_toggle_size_allocate(
		MENU_ITEM(v), 
		gint(i)
	)
}

// gtkRadioMenuItem
type RadioMenuItem struct {
	CheckMenuItem
}

func NewRadioMenuItem(group *glib.SList) *RadioMenuItem {
	return &RadioMenuItem{CheckMenuItem{MenuItem{Item{Bin{Container{Widget{C.gtk_radio_menu_item_new(gslist(group))}}}}}}}
}

func NewRadioMenuItemWithLabel(group *glib.SList, label string) *RadioMenuItem {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &RadioMenuItem{CheckMenuItem{MenuItem{Item{Bin{Container{Widget{C.gtk_radio_menu_item_new_with_label(gslist(group), gstring(ptr))}}}}}}}
}

func (v *RadioMenuItem) SetGroup(group *glib.SList) {
	C.gtk_radio_menu_item_set_group(
		RADIO_MENU_ITEM(v), 
		gslist(group)
	)
}

func (v *RadioMenuItem) GetGroup() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(
		C.gtk_radio_menu_item_get_group(RADIO_MENU_ITEM(v)))
	)
}

// gtkCheckMenuItem
type CheckMenuItem struct {
	MenuItem
}

func NewCheckMenuItem() *CheckMenuItem {
	return &CheckMenuItem{MenuItem{Item{Bin{Container{Widget{C.gtk_check_menu_item_new()}}}}}}
}

func NewCheckMenuItemWithLabel(label string) *CheckMenuItem {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &CheckMenuItem{MenuItem{Item{Bin{Container{Widget{C.gtk_check_menu_item_new_with_label(gstring(ptr))}}}}}}
}

func NewCheckMenuItemWithMnemonic(label string) *CheckMenuItem {
	ptr := C.CString(label)
	defer cfree(ptr)
	return &CheckMenuItem{MenuItem{Item{Bin{Container{Widget{C.gtk_check_menu_item_new_with_mnemonic(gstring(ptr))}}}}}}
}

func (v *CheckMenuItem) GetActive() bool {
	return gobool(C.gtk_check_menu_item_get_active(CHECK_MENU_ITEM(v)))
}

func (v *CheckMenuItem) SetActive(setting bool) {
	C.gtk_check_menu_item_set_active(
		CHECK_MENU_ITEM(v), 
		gbool(setting)
	)
}

func (v *CheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(CHECK_MENU_ITEM(v))
}

func (v *CheckMenuItem) GetInconsistent() bool {
	return gobool(C.gtk_check_menu_item_get_inconsistent(CHECK_MENU_ITEM(v)))
}

func (v *CheckMenuItem) SetInconsistent(setting bool) {
	C.gtk_check_menu_item_set_inconsistent(
		CHECK_MENU_ITEM(v), 
		gbool(setting)
	)
}

func (v *CheckMenuItem) SetDrawAsRadio(setting bool) {
	C.gtk_check_menu_item_set_draw_as_radio(
		CHECK_MENU_ITEM(v), 
		gbool(setting)
	)
}

func (v *CheckMenuItem) GetDrawAsRadio() bool {
	return gobool(C.gtk_check_menu_item_get_draw_as_radio(CHECK_MENU_ITEM(v)))
}

// gtkSeparatorMenuItem
type SeparatorMenuItem struct {
	MenuItem
}

func NewSeparatorMenuItem() *SeparatorMenuItem {
	return &SeparatorMenuItem{MenuItem{Item{Bin{Container{Widget{C.gtk_separator_menu_item_new()}}}}}}
}

// gtkTearoffMenuItem
type TearoffMenuItem struct {
	MenuItem
}

func NewTearoffMenuItem() *TearoffMenuItem {
	return &TearoffMenuItem{MenuItem{Item{Bin{Container{Widget{C.gtk_tearoff_menu_item_new()}}}}}}
}

// gtkToolbar

type Orientation int

const (
	ORIENTATION_HORIZONTAL = 0
	ORIENTATION_VERTICAL   = 1
)

type ToolbarStyle int

const (
	TOOLBAR_ICONS      = 0
	TOOLBAR_TEXT       = 1
	TOOLBAR_BOTH       = 2
	TOOLBAR_BOTH_HORIZ = 3
)

type Toolbar struct {
	Container
	items map[*C.GtkToolItem]IWidget
}

func NewToolbar() *Toolbar {
	return &Toolbar{Container{Widget{C.gtk_toolbar_new()}}, make(map[*C.GtkToolItem]IWidget)}
}

func (v *Toolbar) OnFocusHomeOrEnd(onclick interface{}, datas ...interface{}) int {
	return v.Connect("focus-home-or-end", onclick, datas...)
}

func (v *Toolbar) OnOrientationChanged(onclick interface{}, datas ...interface{}) int {
	return v.Connect("orientation-changed", onclick, datas...)
}

func (v *Toolbar) OnPopupContextMenu(onclick interface{}, datas ...interface{}) int {
	return v.Connect("popup-context-menu", onclick, datas...)
}

func (v *Toolbar) OnStyleChanged(onclick interface{}, datas ...interface{}) int {
	return v.Connect("style-changed", onclick, datas...)
}

func (v *Toolbar) Insert(item IWidget, pos int) {
	p_tool_item := C.toGToolItem(ToNative(item))
	v.items[p_tool_item] = item
	C.gtk_toolbar_insert(TOOLBAR(v), p_tool_item, gint(pos))
}

func (v *Toolbar) GetItemIndex(item IWidget) int {
	return int(C.gtk_toolbar_get_item_index(TOOLBAR(v), C.toGToolItem(ToNative(item))))
}

func (v *Toolbar) GetNItems() int {
	return int(C.gtk_toolbar_get_n_items(TOOLBAR(v)))
}

func (v *Toolbar) GetNthItem(n int) IWidget {
	p_tool_item := C.gtk_toolbar_get_nth_item(TOOLBAR(v), gint(n))
	if p_tool_item == nil {
		panic("Toolbar.GetNthItem: index out of bounds")
	}
	if _, ok := v.items[p_tool_item]; !ok {
		panic("Toolbar.GetNthItem: interface not found in map")
	}
	return v.items[p_tool_item]
}

func (v *Toolbar) GetDropIndex(x, y int) int {
	return int(C.gtk_toolbar_get_drop_index(TOOLBAR(v), gint(x), gint(y)))
}

func (v *Toolbar) SetDropHighlightItem(item IWidget, index int) {
	C.gtk_toolbar_set_drop_highlight_item(
		TOOLBAR(v), 
		C.toGToolItem(ToNative(item)), 
		gint(index)
	)
}

func (v *Toolbar) SetShowArrow(show_arrow bool) {
	C.gtk_toolbar_set_show_arrow(
		TOOLBAR(v), 
		gbool(show_arrow)
	)
}

func (v *Toolbar) SetOrientation(orientation Orientation) {
	C.gtk_toolbar_set_orientation(
		TOOLBAR(v), 
		C.GtkOrientation(orientation)
	)
}

func (v *Toolbar) SetTooltips(enable bool) {
	C.gtk_toolbar_set_tooltips(
		TOOLBAR(v), 
		gbool(enable)
	)
}

func (v *Toolbar) UnsetIconSize() {
	C.gtk_toolbar_unset_icon_size(TOOLBAR(v))
}

func (v *Toolbar) GetShowArrow() bool {
	return gobool(C.gtk_toolbar_get_show_arrow(TOOLBAR(v)))
}

func (v *Toolbar) GetOrientation() Orientation {
	return Orientation(C.gtk_toolbar_get_orientation(TOOLBAR(v)))
}

func (v *Toolbar) GetStyle() ToolbarStyle {
	return ToolbarStyle(C.gtk_toolbar_get_style(TOOLBAR(v)))
}

func (v *Toolbar) GetIconSize() IconSize {
	return IconSize(C.gtk_toolbar_get_icon_size(TOOLBAR(v)))
}

func (v *Toolbar) GetTooltips() bool {
	return gobool(C.gtk_toolbar_get_tooltips(TOOLBAR(v)))
}

func (v *Toolbar) GetReliefStyle() ReliefStyle {
	return ReliefStyle(C.gtk_toolbar_get_relief_style(TOOLBAR(v)))
}

func (v *Toolbar) SetStyle(style ToolbarStyle) {
	C.gtk_toolbar_set_style(
		TOOLBAR(v), 
		C.GtkToolbarStyle(style)
	)
}

func (v *Toolbar) SetIconSize(icon_size IconSize) {
	C.gtk_toolbar_set_icon_size(
		TOOLBAR(v), 
		C.GtkIconSize(icon_size)
	)
}

func (v *Toolbar) UnsetStyle() {
	C.gtk_toolbar_unset_style(TOOLBAR(v))
}

// gtkToolItem
type ToolItem struct {
	Bin
}

func NewToolItem() *ToolItem {
	return &ToolItem{Bin{Container{Widget{C.toGWidget(unsafe.Pointer(C.gtk_tool_item_new()))}}}}
}

func (v *ToolItem) OnCreateMenuProxy(onclick interface{}, datas ...interface{}) int {
	return v.Connect("create-menu-proxy", onclick, datas...)
}

func (v *ToolItem) OnSetTooltip(onclick interface{}, datas ...interface{}) int {
	return v.Connect("set-tooltip", onclick, datas...)
}

func (v *ToolItem) OnToolbarReconfigured(onclick interface{}, datas ...interface{}) int {
	return v.Connect("toolbar-reconfigured", onclick, datas...)
}

func (v *ToolItem) SetHomogeneous(homogeneous bool) {
	C.gtk_tool_item_set_homogeneous(
		TOOL_ITEM(v), 
		gbool(homogeneous)
	)
}

func (v *ToolItem) GetHomogeneous() bool {
	return gobool(C.gtk_tool_item_get_homogeneous(TOOL_ITEM(v)))
}

func (v *ToolItem) SetExpand(expand bool) {
	C.gtk_tool_item_set_expand(
		TOOL_ITEM(v), 
		gbool(expand)
	)
}

func (v *ToolItem) GetExpand() bool {
	return gobool(C.gtk_tool_item_get_expand(TOOL_ITEM(v)))
}

func (v *ToolItem) SetArrowTooltipText(text string) {
	pt := C.CString(text)
	defer cfree(pt)
	C.gtk_tool_item_set_tooltip_text(
		TOOL_ITEM(v), 
		gstring(pt)
	)
}

func (v *ToolItem) SetArrowTooltipMarkup(markup string) {
	pm := C.CString(markup)
	defer cfree(pm)
	C.gtk_tool_item_set_tooltip_markup(
		TOOL_ITEM(v), 
		gstring(pm)
	)
}

func (v *ToolItem) SetTooltipMarkup(markup string) {
	p_markup := C.CString(markup)
	defer cfree(p_markup)
	C.gtk_tool_item_set_tooltip_markup(
		TOOL_ITEM(v), 
		gstring(p_markup)
	)
}

func (v *ToolItem) GetToolbarStyle() ToolbarStyle {
	return ToolbarStyle(C.gtk_tool_item_get_toolbar_style(TOOL_ITEM(v)))
}

func (v *ToolItem) GetReliefStyle() ReliefStyle {
	return ReliefStyle(C.gtk_tool_item_get_relief_style(TOOL_ITEM(v)))
}

func (v *ToolItem) GetTextAlignment() float64 {
	return float64(C.gtk_tool_item_get_text_alignment(TOOL_ITEM(v)))
}

func (v *ToolItem) GetTextOrientation() Orientation {
	return Orientation(C.gtk_tool_item_get_text_orientation(TOOL_ITEM(v)))
}

func (v *ToolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu(TOOL_ITEM(v))
}

// gtkToolPalette
type ToolPalette struct {
	Container
}

func NewToolPalette() *ToolPalette {
	return &ToolPalette{Container{Widget{C.gtk_tool_palette_new()}}}
}

func (v *ToolPalette) GetExclusive(group *ToolItemGroup) bool {
	return gobool(C.gtk_tool_palette_get_exclusive(TOOL_PALETTE(v), TOOL_ITEM_GROUP(group)))
}

func (v *ToolPalette) SetExclusive(group *ToolItemGroup, exclusive bool) {
	C.gtk_tool_palette_set_exclusive(
		TOOL_PALETTE(v), 
		TOOL_ITEM_GROUP(group), 
		gbool(exclusive)
	)
}

func (v *ToolPalette) GetExpand(group *ToolItemGroup) bool {
	return gobool(C.gtk_tool_palette_get_expand(TOOL_PALETTE(v), TOOL_ITEM_GROUP(group)))
}

func (v *ToolPalette) SetExpand(group *ToolItemGroup, expand bool) {
	C.gtk_tool_palette_set_expand(
		TOOL_PALETTE(v), 
		TOOL_ITEM_GROUP(group), 
		gbool(expand)
	)
}

func (v *ToolPalette) GetGroupPosition(group *ToolItemGroup) int {
	return int(C.gtk_tool_palette_get_group_position(TOOL_PALETTE(v), TOOL_ITEM_GROUP(group)))
}

func (v *ToolPalette) SetGroupPosition(group *ToolItemGroup, pos int) {
	C.gtk_tool_palette_set_group_position(
		TOOL_PALETTE(v), 
		TOOL_ITEM_GROUP(group), 
		gint(pos)
	)
}

func (v *ToolPalette) GetIconSize() IconSize {
	return IconSize(C.gtk_tool_palette_get_icon_size(TOOL_PALETTE(v)))
}

func (v *ToolPalette) SetIconSize(size IconSize) {
	C.gtk_tool_palette_set_icon_size(
		TOOL_PALETTE(v), 
		C.GtkIconSize(size)
	)
}

func (v *ToolPalette) UnsetIconSize() {
	C.gtk_tool_palette_unset_icon_size(TOOL_PALETTE(v))
}

func (v *ToolPalette) GetStyle() ToolbarStyle {
	return ToolbarStyle(C.gtk_tool_palette_get_style(TOOL_PALETTE(v)))
}

func (v *ToolPalette) SetStyle(style ToolbarStyle) {
	C.gtk_tool_palette_set_style(
		TOOL_PALETTE(v), 
		C.GtkToolbarStyle(style)
	)
}

func (v *ToolPalette) UnsetStyle() {
	C.gtk_tool_palette_unset_style(TOOL_PALETTE(v))
}

// gtkToolItemGroup
type ToolItemGroup struct {
	Container
	items map[*C.GtkToolItem]IWidget
}

func NewToolItemGroup(label string) *ToolItemGroup {
	l := C.CString(label)
	defer cfree(l)
	return &ToolItemGroup{Container{Widget{C.gtk_tool_item_group_new(gstring(l))}}, make(map[*C.GtkToolItem]IWidget)}
}

func (v *ToolItemGroup) Insert(item IWidget, pos int) {
	pitem := C.toGToolItem(ToNative(item))
	C.gtk_tool_item_group_insert(
		TOOL_ITEM_GROUP(v), 
		pitem, 
		gint(pos)
	)
	v.items[pitem] = item
}

func (v *ToolItemGroup) GetCollapsed() bool {
	return gobool(C.gtk_tool_item_group_get_collapsed(TOOL_ITEM_GROUP(v)))
}

func (v *ToolItemGroup) GetLabel() string {
	return gostring(C.gtk_tool_item_group_get_label(TOOL_ITEM_GROUP(v)))
}

func (v *ToolItemGroup) GetHeaderRelief() ReliefStyle {
	return ReliefStyle(C.gtk_tool_item_group_get_header_relief(TOOL_ITEM_GROUP(v)))
}

func (v *ToolItemGroup) SetCollapsed(collapsed bool) {
	C.gtk_tool_item_group_set_collapsed(
		TOOL_ITEM_GROUP(v), 
		gbool(collapsed)
	)
}

func (v *ToolItemGroup) SetLabel(label string) {
	l := C.CString(label)
	defer cfree(l)
	C.gtk_tool_item_group_set_label(
		TOOL_ITEM_GROUP(v), 
		gstring(l)
	)
}

func (v *ToolItemGroup) SetHeaderRelief(style ReliefStyle) {
	C.gtk_tool_item_group_set_header_relief(
		TOOL_ITEM_GROUP(v), 
		C.GtkReliefStyle(style)
	)
}

// gtkSeparatorToolItem
type SeparatorToolItem struct {
	ToolItem
}

func NewSeparatorToolItem() *SeparatorToolItem {
	return &SeparatorToolItem{ToolItem{Bin{Container{Widget{C.toGWidget(unsafe.Pointer(C.gtk_separator_tool_item_new()))}}}}}
}

func (v *SeparatorToolItem) SetDraw(draw bool) {
	C.gtk_separator_tool_item_set_draw(
		SEPARATOR_TOOL_ITEM(v), 
		gbool(draw)
	)
}

func (v *SeparatorToolItem) GetDraw() bool {
	return gobool(C.gtk_separator_tool_item_get_draw(SEPARATOR_TOOL_ITEM(v)))
}

// gtkToolButton
type ToolButton struct {
	ToolItem
	iw, lw *Widget
}

func NewToolButton(icon IWidget, text string) *ToolButton {
	ctext := C.CString(text)
	defer cfree(ctext)
	tb := C.toGWidget(unsafe.Pointer(C.gtk_tool_button_new(ToNative(icon), gstring(ctext))))
	return &ToolButton{ToolItem{Bin{Container{Widget{tb}}}}, nil, nil}
}

func NewToolButtonFromStock(stock_id string) *ToolButton {
	si := C.CString(stock_id)
	defer cfree(si)
	tb := C.toGWidget(unsafe.Pointer(C.gtk_tool_button_new_from_stock(gstring(si))))
	return &ToolButton{ToolItem{Bin{Container{Widget{tb}}}}, nil, nil}
}

func (v *ToolButton) OnClicked(onclick interface{}, datas ...interface{}) int {
	return v.Connect("clicked", onclick, datas...)
}

func (v *ToolButton) SetLabel(label string) {
	p_label := C.CString(label)
	defer cfree(p_label)
	C.gtk_tool_button_set_label(
		TOOL_BUTTON(v), 
		gstring(p_label)
	)
}

func (v *ToolButton) GetLabel() string {
	return gostring(C.gtk_tool_button_get_label(TOOL_BUTTON(v)))
}

func (v *ToolButton) SetUseUnderline(use_underline bool) {
	C.gtk_tool_button_set_use_underline(
		TOOL_BUTTON(v), 
		gbool(use_underline)
	)
}

func (v *ToolButton) GetUseUnderline() bool {
	return gobool(C.gtk_tool_button_get_use_underline(TOOL_BUTTON(v)))
}

func (v *ToolButton) SetStockId(stock_id string) {
	p_stock_id := C.CString(stock_id)
	defer cfree(p_stock_id)
	C.gtk_tool_button_set_stock_id(
		TOOL_BUTTON(v), 
		gstring(p_stock_id)
	)
}

func (v *ToolButton) GetStockId() string {
	return gostring(C.gtk_tool_button_get_stock_id(TOOL_BUTTON(v)))
}

func (v *ToolButton) SetIconName(icon_name string) {
	p_icon_name := C.CString(icon_name)
	defer cfree(p_icon_name)
	C.gtk_tool_button_set_icon_name(
		TOOL_BUTTON(v), 
		gstring(p_icon_name)
	)
}

func (v *ToolButton) GetIconName() string {
	return gostring(C.gtk_tool_button_get_icon_name(TOOL_BUTTON(v)))
}

func (v *ToolButton) SetIconWidget(icon_widget *Widget) {
	v.iw = icon_widget
	C.gtk_tool_button_set_icon_widget(
		TOOL_BUTTON(v), 
		icon_widget.GWidget
	)
}

func (v *ToolButton) GetIconWidget() *Widget {
	if v.iw == nil {
		v.iw = &Widget{C.gtk_tool_button_get_icon_widget(TOOL_BUTTON(v))}
	}
	return v.iw
}

func (v *ToolButton) SetLabelWidget(label_widget *Widget) {
	v.lw = label_widget
	C.gtk_tool_button_set_label_widget(
		TOOL_BUTTON(v), 
		label_widget.GWidget
	)
}

func (v *ToolButton) GetLabelWidget() *Widget {
	if v.lw == nil {
		v.lw = &Widget{C.gtk_tool_button_get_label_widget(TOOL_BUTTON(v))}
	}
	return v.lw
}

// gtkMenuToolButton
type MenuToolButton struct {
	ToolButton
	mw *Menu
}

func NewMenuToolButton(icon IWidget, text string) *MenuToolButton {
	ctext := C.CString(text)
	defer cfree(ctext)
	mtb := C.toGWidget(unsafe.Pointer(C.gtk_menu_tool_button_new(ToNative(icon), gstring(ctext))))
	return &MenuToolButton{ToolButton{ToolItem{Bin{Container{Widget{mtb}}}}, nil, nil}, nil}
}

func NewMenuToolButtonFromStock(stock_id string) *MenuToolButton {
	si := C.CString(stock_id)
	defer cfree(si)
	mtb := C.toGWidget(unsafe.Pointer(C.gtk_menu_tool_button_new_from_stock(gstring(si))))
	return &MenuToolButton{ToolButton{ToolItem{Bin{Container{Widget{mtb}}}}, nil, nil}, nil}
}

func (v *MenuToolButton) SetMenu(menu *Menu) {
	v.mw = menu
	C.gtk_menu_tool_button_set_menu(
		MENU_TOOL_BUTTON(v), 
		menu.GWidget
	)
}

func (v *MenuToolButton) GetMenu() *Menu {
	if v.mw == nil {
		v.mw = &Menu{Container{Widget{C.gtk_menu_tool_button_get_menu(MENU_TOOL_BUTTON(v))}}}
	}
	return v.mw
}

func (v *MenuToolButton) SetArrowTooltipText(text string) {
	pt := C.CString(text)
	defer cfree(pt)
	C.gtk_menu_tool_button_set_arrow_tooltip_text(
		MENU_TOOL_BUTTON(v), 
		gstring(pt)
	)
}

func (v *MenuToolButton) SetArrowTooltipMarkup(markup string) {
	pm := C.CString(markup)
	defer cfree(pm)
	C.gtk_menu_tool_button_set_arrow_tooltip_text(
		MENU_TOOL_BUTTON(v), 
		gstring(pm)
	)
}

