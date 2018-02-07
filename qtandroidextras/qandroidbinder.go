package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h
// #include <qandroidbinder.h>
// #include <QtAndroidExtras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QAndroidBinder struct {
	*qtrt.CObject
}

func (this *QAndroidBinder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidBinder) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidBinderFromPointer(cthis unsafe.Pointer) *QAndroidBinder {
	return &QAndroidBinder{&qtrt.CObject{cthis}}
}
func (*QAndroidBinder) NewFromPointer(cthis unsafe.Pointer) *QAndroidBinder {
	return NewQAndroidBinderFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidBinder()
func NewQAndroidBinder() *QAndroidBinder {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinderC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidBinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidBinder)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAndroidBinder(const QAndroidJniObject &)
func NewQAndroidBinder_1(binder *QAndroidJniObject) *QAndroidBinder {
	var convArg0 = binder.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinderC2ERK17QAndroidJniObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidBinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidBinder)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidBinder()
func DeleteQAndroidBinder(this *QAndroidBinder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 1)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool onTransact(int, const QAndroidParcel &, const QAndroidParcel &, enum QAndroidBinder::CallType)
func (this *QAndroidBinder) OnTransact(code int, data *QAndroidParcel, reply *QAndroidParcel, flags int) bool {
	var convArg1 = data.GetCthis()
	var convArg2 = reply.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinder10onTransactEiRK14QAndroidParcelS2_NS_8CallTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1, convArg2, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool transact(int, const QAndroidParcel &, QAndroidParcel *, enum QAndroidBinder::CallType)
func (this *QAndroidBinder) Transact(code int, data *QAndroidParcel, reply *QAndroidParcel /*777 QAndroidParcel **/, flags int) bool {
	var convArg1 = data.GetCthis()
	var convArg2 = reply.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidBinder8transactEiRK14QAndroidParcelPS0_NS_8CallTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1, convArg2, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] QAndroidJniObject handle()
func (this *QAndroidBinder) Handle() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidBinder6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

type QAndroidBinder__CallType = int

const QAndroidBinder__Normal QAndroidBinder__CallType = 0
const QAndroidBinder__OneWay QAndroidBinder__CallType = 1

//  body block end
