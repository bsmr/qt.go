package qtqml

// /usr/include/qt/QtQml/qqmlnetworkaccessmanagerfactory.h
// #include <qqmlnetworkaccessmanagerfactory.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin

type QQmlNetworkAccessManagerFactory struct {
	*qtrt.CObject
}

func (this *QQmlNetworkAccessManagerFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlNetworkAccessManagerFactory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlNetworkAccessManagerFactoryFromPointer(cthis unsafe.Pointer) *QQmlNetworkAccessManagerFactory {
	return &QQmlNetworkAccessManagerFactory{&qtrt.CObject{cthis}}
}
func (*QQmlNetworkAccessManagerFactory) NewFromPointer(cthis unsafe.Pointer) *QQmlNetworkAccessManagerFactory {
	return NewQQmlNetworkAccessManagerFactoryFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlnetworkaccessmanagerfactory.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlNetworkAccessManagerFactory()
func DeleteQQmlNetworkAccessManagerFactory(this *QQmlNetworkAccessManagerFactory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QQmlNetworkAccessManagerFactoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlnetworkaccessmanagerfactory.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QNetworkAccessManager * create(QObject *)
func (this *QQmlNetworkAccessManagerFactory) Create(parent *qtcore.QObject /*777 QObject **/) *qtnetwork.QNetworkAccessManager /*777 QNetworkAccessManager **/ {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN31QQmlNetworkAccessManagerFactory6createEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtnetwork.NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end