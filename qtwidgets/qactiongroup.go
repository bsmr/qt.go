package qtwidgets

// /usr/include/qt/QtWidgets/qactiongroup.h
// #include <qactiongroup.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 62
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

type QActionGroup struct {
	*qtcore.QObject
}

func (this *QActionGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QActionGroup) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQActionGroupFromPointer(cthis unsafe.Pointer) *QActionGroup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QActionGroup{bcthis0}
}
func (*QActionGroup) NewFromPointer(cthis unsafe.Pointer) *QActionGroup {
	return NewQActionGroupFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QActionGroup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QActionGroup(QObject *)
func NewQActionGroup(parent *qtcore.QObject /*777 QObject **/) *QActionGroup {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qactiongroup.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QActionGroup()
func DeleteQActionGroup(this *QActionGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(QAction *)
func (this *QActionGroup) AddAction(a *QAction /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 = a.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup9addActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:67
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &)
func (this *QActionGroup) AddAction_1(text string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup9addActionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:68
// index:2
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &)
func (this *QActionGroup) AddAction_2(icon *qtgui.QIcon, text string) *QAction /*777 QAction **/ {
	var convArg0 = icon.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup9addActionERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAction(QAction *)
func (this *QActionGroup) RemoveAction(a *QAction /*777 QAction **/) {
	var convArg0 = a.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup12removeActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * checkedAction()
func (this *QActionGroup) CheckedAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup13checkedActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExclusive()
func (this *QActionGroup) IsExclusive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup11isExclusiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qactiongroup.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QActionGroup) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qactiongroup.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QActionGroup) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qactiongroup.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QActionGroup) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setDisabled(_Bool)
func (this *QActionGroup) SetDisabled(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup11setDisabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QActionGroup) SetVisible(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExclusive(_Bool)
func (this *QActionGroup) SetExclusive(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup12setExclusiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggered(QAction *)
func (this *QActionGroup) Triggered(arg0 *QAction /*777 QAction **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup9triggeredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hovered(QAction *)
func (this *QActionGroup) Hovered(arg0 *QAction /*777 QAction **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup7hoveredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end