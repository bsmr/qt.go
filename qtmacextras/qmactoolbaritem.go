package qtmacextras

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h
// #include <qmactoolbaritem.h>
// #include <Qtmacextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

type QMacToolBarItem struct {
	*qtcore.QObject
}

func (this *QMacToolBarItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMacToolBarItem) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQMacToolBarItemFromPointer(cthis unsafe.Pointer) *QMacToolBarItem {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QMacToolBarItem{bcthis0}
}
func (*QMacToolBarItem) NewFromPointer(cthis unsafe.Pointer) *QMacToolBarItem {
	return NewQMacToolBarItemFromPointer(cthis)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QMacToolBarItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMacToolBarItem(QObject *)
func NewQMacToolBarItem(parent *qtcore.QObject /*777 QObject **/) *QMacToolBarItem {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItemC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMacToolBarItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMacToolBarItem()
func DeleteQMacToolBarItem(this *QMacToolBarItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool selectable()
func (this *QMacToolBarItem) Selectable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem10selectableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectable(_Bool)
func (this *QMacToolBarItem) SetSelectable(selectable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem13setSelectableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selectable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] QMacToolBarItem::StandardItem standardItem()
func (this *QMacToolBarItem) StandardItem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem12standardItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardItem(enum QMacToolBarItem::StandardItem)
func (this *QMacToolBarItem) SetStandardItem(standardItem int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem15setStandardItemENS_12StandardItemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardItem)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QMacToolBarItem) Text() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QMacToolBarItem) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QMacToolBarItem) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QMacToolBarItem) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] NSToolbarItem * nativeToolBarItem()
func (this *QMacToolBarItem) NativeToolBarItem() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMacToolBarItem17nativeToolBarItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbaritem.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated()
func (this *QMacToolBarItem) Activated() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMacToolBarItem9activatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QMacToolBarItem__StandardItem = int

const QMacToolBarItem__NoStandardItem QMacToolBarItem__StandardItem = 0
const QMacToolBarItem__Space QMacToolBarItem__StandardItem = 1
const QMacToolBarItem__FlexibleSpace QMacToolBarItem__StandardItem = 2

//  body block end
