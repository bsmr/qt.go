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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type VariantData struct {
	*qtrt.CObject
}

func (this *VariantData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *VariantData) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewVariantDataFromPointer(cthis unsafe.Pointer) *VariantData {
	return &VariantData{&qtrt.CObject{cthis}}
}
func (*VariantData) NewFromPointer(cthis unsafe.Pointer) *VariantData {
	return NewVariantDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetatype.h:798
// index:0
// Public inline
// void VariantData(const int, const void *, const uint)
func NewVariantData(metaTypeId_ int, data_ unsafe.Pointer /*666*/, flags_ uint) *VariantData {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QtMetaTypePrivate11VariantDataC2EiPKvj", ffiqt.FFI_TYPE_VOID, cthis, metaTypeId_, data_, flags_)
	gopp.ErrPrint(err, rv)
	gothis := NewVariantDataFromPointer(cthis)
	return gothis
}

//  body block end
