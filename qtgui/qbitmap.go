package qtgui

// /usr/include/qt/QtGui/qbitmap.h
// #include <qbitmap.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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

type QBitmap struct {
	*QPixmap
}

func (this *QBitmap) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPixmap.GetCthis()
	}
}
func (this *QBitmap) SetCthis(cthis unsafe.Pointer) {
	this.QPixmap = NewQPixmapFromPointer(cthis)
}
func NewQBitmapFromPointer(cthis unsafe.Pointer) *QBitmap {
	bcthis0 := NewQPixmapFromPointer(cthis)
	return &QBitmap{bcthis0}
}
func (*QBitmap) NewFromPointer(cthis unsafe.Pointer) *QBitmap {
	return NewQBitmapFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbitmap.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBitmap()
func NewQBitmap() *QBitmap {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmapC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBitmap)
	return gothis
}

// /usr/include/qt/QtGui/qbitmap.h:55
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QBitmap(const QPixmap &)
func NewQBitmap_1(arg0 *QPixmap) *QBitmap {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmapC2ERK7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBitmap)
	return gothis
}

// /usr/include/qt/QtGui/qbitmap.h:56
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QBitmap(int, int)
func NewQBitmap_2(w int, h int) *QBitmap {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmapC2Eii", qtrt.FFI_TYPE_POINTER, w, h)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBitmap)
	return gothis
}

// /usr/include/qt/QtGui/qbitmap.h:57
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QBitmap(const QSize &)
func NewQBitmap_3(arg0 *qtcore.QSize) *QBitmap {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmapC2ERK5QSize", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBitmap)
	return gothis
}

// /usr/include/qt/QtGui/qbitmap.h:58
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QBitmap(const QString &, const char *)
func NewQBitmap_4(fileName string, format string) *QBitmap {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmapC2ERK7QStringPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBitmap)
	return gothis
}

// /usr/include/qt/QtGui/qbitmap.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QBitmap()
func DeleteQBitmap(this *QBitmap) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmapD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qbitmap.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QBitmap &)
func (this *QBitmap) Swap(other *QBitmap) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmap4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbitmap.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clear()
func (this *QBitmap) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmap5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbitmap.h:74
// index:0
// Public static Visibility=Default Availability=Available
// [32] QBitmap fromImage(const QImage &, Qt::ImageConversionFlags)
func (this *QBitmap) FromImage(image *QImage, flags int) *QBitmap /*123*/ {
	var convArg0 = image.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmap9fromImageERK6QImage6QFlagsIN2Qt19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}
func QBitmap_FromImage(image *QImage, flags int) *QBitmap /*123*/ {
	var nilthis *QBitmap
	rv := nilthis.FromImage(image, flags)
	return rv
}

// /usr/include/qt/QtGui/qbitmap.h:75
// index:0
// Public static Visibility=Default Availability=Available
// [32] QBitmap fromData(const QSize &, const uchar *, QImage::Format)
func (this *QBitmap) FromData(size *qtcore.QSize, bits unsafe.Pointer /*666*/, monoFormat int) *QBitmap /*123*/ {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitmap8fromDataERK5QSizePKhN6QImage6FormatE", qtrt.FFI_TYPE_POINTER, convArg0, &bits, monoFormat)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}
func QBitmap_FromData(size *qtcore.QSize, bits unsafe.Pointer /*666*/, monoFormat int) *QBitmap /*123*/ {
	var nilthis *QBitmap
	rv := nilthis.FromData(size, bits, monoFormat)
	return rv
}

// /usr/include/qt/QtGui/qbitmap.h:78
// index:0
// Public Visibility=Default Availability=Available
// [32] QBitmap transformed(const QMatrix &)
func (this *QBitmap) Transformed(arg0 *QMatrix) *QBitmap /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBitmap11transformedERK7QMatrix", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

// /usr/include/qt/QtGui/qbitmap.h:79
// index:1
// Public Visibility=Default Availability=Available
// [32] QBitmap transformed(const QTransform &)
func (this *QBitmap) Transformed_1(matrix *QTransform) *QBitmap /*123*/ {
	var convArg0 = matrix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBitmap11transformedERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitmap)
	return rv2
}

//  body block end