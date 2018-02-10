package qtwidgets

// /usr/include/qt/QtWidgets/qwidgetaction.h
// #include <qwidgetaction.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
// bool event(class QEvent *)
func (this *QWidgetAction) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QWidgetAction) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// QWidget * createWidget(class QWidget *)
func (this *QWidgetAction) InheritCreateWidget(f func(parent *QWidget /*777 QWidget **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createWidget", f)
}

// void deleteWidget(class QWidget *)
func (this *QWidgetAction) InheritDeleteWidget(f func(widget *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "deleteWidget", f)
}

type QWidgetAction struct {
	*QAction
}

func (this *QWidgetAction) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAction.GetCthis()
	}
}
func (this *QWidgetAction) SetCthis(cthis unsafe.Pointer) {
	this.QAction = NewQActionFromPointer(cthis)
}
func NewQWidgetActionFromPointer(cthis unsafe.Pointer) *QWidgetAction {
	bcthis0 := NewQActionFromPointer(cthis)
	return &QWidgetAction{bcthis0}
}
func (*QWidgetAction) NewFromPointer(cthis unsafe.Pointer) *QWidgetAction {
	return NewQWidgetActionFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QWidgetAction) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QWidgetAction10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidgetAction(QObject *)
func NewQWidgetAction(parent *qtcore.QObject /*777 QObject **/) *QWidgetAction {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetActionC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWidgetActionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWidgetAction()
func DeleteQWidgetAction(this *QWidgetAction) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetActionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultWidget(QWidget *)
func (this *QWidgetAction) SetDefaultWidget(w *QWidget /*777 QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction16setDefaultWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * defaultWidget()
func (this *QWidgetAction) DefaultWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QWidgetAction13defaultWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * requestWidget(QWidget *)
func (this *QWidgetAction) RequestWidget(parent *QWidget /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction13requestWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseWidget(QWidget *)
func (this *QWidgetAction) ReleaseWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction13releaseWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QWidgetAction) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QWidgetAction) EventFilter(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:71
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QWidget * createWidget(QWidget *)
func (this *QWidgetAction) CreateWidget(parent *QWidget /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction12createWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:72
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void deleteWidget(QWidget *)
func (this *QWidgetAction) DeleteWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction12deleteWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end