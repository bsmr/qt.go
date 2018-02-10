package qtcore

// /usr/include/qt/QtCore/qmetatype.h
// #include <qmetatype.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type AbstractConverterFunction struct {
	*qtrt.CObject
}

func (this *AbstractConverterFunction) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *AbstractConverterFunction) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewAbstractConverterFunctionFromPointer(cthis unsafe.Pointer) *AbstractConverterFunction {
	return &AbstractConverterFunction{&qtrt.CObject{cthis}}
}
func (*AbstractConverterFunction) NewFromPointer(cthis unsafe.Pointer) *AbstractConverterFunction {
	return NewAbstractConverterFunctionFromPointer(cthis)
}

func DeleteAbstractConverterFunction(this *AbstractConverterFunction) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25AbstractConverterFunctionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end