package qtwidgets

// /usr/include/qt/QtWidgets/qtablewidget.h
// #include <qtablewidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QTableWidgetItem struct {
	*qtrt.CObject
}

func (this *QTableWidgetItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTableWidgetItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTableWidgetItemFromPointer(cthis unsafe.Pointer) *QTableWidgetItem {
	return &QTableWidgetItem{&qtrt.CObject{cthis}}
}
func (*QTableWidgetItem) NewFromPointer(cthis unsafe.Pointer) *QTableWidgetItem {
	return NewQTableWidgetItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTableWidgetItem(int)
func NewQTableWidgetItem(type_ int) *QTableWidgetItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItemC2Ei", qtrt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTableWidgetItem(const QString &, int)
func NewQTableWidgetItem_1(text string, type_ int) *QTableWidgetItem {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItemC2ERK7QStringi", qtrt.FFI_TYPE_POINTER, convArg0, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:84
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTableWidgetItem(const QIcon &, const QString &, int)
func NewQTableWidgetItem_2(icon *qtgui.QIcon, text string, type_ int) *QTableWidgetItem {
	var convArg0 = icon.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItemC2ERK5QIconRK7QStringi", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTableWidgetItem()
func DeleteQTableWidgetItem(this *QTableWidgetItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QTableWidgetItem * clone()
func (this *QTableWidgetItem) Clone() *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem5cloneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QTableWidget * tableWidget()
func (this *QTableWidgetItem) TableWidget() *QTableWidget /*777 QTableWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem11tableWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQTableWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtablewidget.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int row()
func (this *QTableWidgetItem) Row() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int column()
func (this *QTableWidgetItem) Column() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem6columnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSelected(_Bool)
func (this *QTableWidgetItem) SetSelected(select_ bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem11setSelectedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSelected()
func (this *QTableWidgetItem) IsSelected() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem10isSelectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags()
func (this *QTableWidgetItem) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::ItemFlags)
func (this *QTableWidgetItem) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem8setFlagsE6QFlagsIN2Qt8ItemFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString text()
func (this *QTableWidgetItem) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtablewidget.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QTableWidgetItem) SetText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QTableWidgetItem) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QTableWidgetItem) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString statusTip()
func (this *QTableWidgetItem) StatusTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem9statusTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtablewidget.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)
func (this *QTableWidgetItem) SetStatusTip(statusTip string) {
	var tmpArg0 = qtcore.NewQString_5(statusTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem12setStatusTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toolTip()
func (this *QTableWidgetItem) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtablewidget.h:116
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)
func (this *QTableWidgetItem) SetToolTip(toolTip string) {
	var tmpArg0 = qtcore.NewQString_5(toolTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString whatsThis()
func (this *QTableWidgetItem) WhatsThis() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem9whatsThisEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtablewidget.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)
func (this *QTableWidgetItem) SetWhatsThis(whatsThis string) {
	var tmpArg0 = qtcore.NewQString_5(whatsThis)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem12setWhatsThisERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QFont font()
func (this *QTableWidgetItem) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QTableWidgetItem) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int textAlignment()
func (this *QTableWidgetItem) TextAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem13textAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:131
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextAlignment(int)
func (this *QTableWidgetItem) SetTextAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem16setTextAlignmentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:134
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor backgroundColor()
func (this *QTableWidgetItem) BackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem15backgroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackgroundColor(const QColor &)
func (this *QTableWidgetItem) SetBackgroundColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem18setBackgroundColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:139
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background()
func (this *QTableWidgetItem) Background() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem10backgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)
func (this *QTableWidgetItem) SetBackground(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem13setBackgroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor textColor()
func (this *QTableWidgetItem) TextColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem9textColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextColor(const QColor &)
func (this *QTableWidgetItem) SetTextColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem12setTextColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:149
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground()
func (this *QTableWidgetItem) Foreground() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem10foregroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:151
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(const QBrush &)
func (this *QTableWidgetItem) SetForeground(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem13setForegroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::CheckState checkState()
func (this *QTableWidgetItem) CheckState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem10checkStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:156
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCheckState(Qt::CheckState)
func (this *QTableWidgetItem) SetCheckState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem13setCheckStateEN2Qt10CheckStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QTableWidgetItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:161
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSizeHint(const QSize &)
func (this *QTableWidgetItem) SetSizeHint(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem11setSizeHintERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:164
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(int)
func (this *QTableWidgetItem) Data(role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:165
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(int, const QVariant &)
func (this *QTableWidgetItem) SetData(role int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem7setDataEiRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:170
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void read(QDataStream &)
func (this *QTableWidgetItem) Read(in *qtcore.QDataStream) {
	var convArg0 = in.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTableWidgetItem4readER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:171
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void write(QDataStream &)
func (this *QTableWidgetItem) Write(out *qtcore.QDataStream) {
	var convArg0 = out.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem5writeER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:175
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int type()
func (this *QTableWidgetItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTableWidgetItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QTableWidgetItem__ItemType = int

const QTableWidgetItem__Type QTableWidgetItem__ItemType = 0
const QTableWidgetItem__UserType QTableWidgetItem__ItemType = 1000

//  body block end