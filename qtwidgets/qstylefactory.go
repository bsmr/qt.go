package qtwidgets

// /usr/include/qt/QtWidgets/qstylefactory.h
// #include <qstylefactory.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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

type QStyleFactory struct {
	*qtrt.CObject
}

func (this *QStyleFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStyleFactory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStyleFactoryFromPointer(cthis unsafe.Pointer) *QStyleFactory {
	return &QStyleFactory{&qtrt.CObject{cthis}}
}
func (*QStyleFactory) NewFromPointer(cthis unsafe.Pointer) *QStyleFactory {
	return NewQStyleFactoryFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstylefactory.h:54
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList keys()
func (this *QStyleFactory) Keys() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStyleFactory4keysEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QStyleFactory_Keys() *qtcore.QStringList /*123*/ {
	var nilthis *QStyleFactory
	rv := nilthis.Keys()
	return rv
}

// /usr/include/qt/QtWidgets/qstylefactory.h:55
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStyle * create(const QString &)
func (this *QStyleFactory) Create(arg0 string) *QStyle /*777 QStyle **/ {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStyleFactory6createERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QStyleFactory_Create(arg0 string) *QStyle /*777 QStyle **/ {
	var nilthis *QStyleFactory
	rv := nilthis.Create(arg0)
	return rv
}

func DeleteQStyleFactory(this *QStyleFactory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStyleFactoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end