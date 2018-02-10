package qtgui

// /usr/include/qt/QtGui/qtransform.h
// #include <qtransform.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

type QTransform struct {
	*qtrt.CObject
}

func (this *QTransform) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTransform) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTransformFromPointer(cthis unsafe.Pointer) *QTransform {
	return &QTransform{&qtrt.CObject{cthis}}
}
func (*QTransform) NewFromPointer(cthis unsafe.Pointer) *QTransform {
	return NewQTransformFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtransform.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTransform(Qt::Initialization)
func NewQTransform(arg0 int) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTransform()
func NewQTransform_1() *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:71
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTransform(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
func NewQTransform_2(h11 float64, h12 float64, h13 float64, h21 float64, h22 float64, h23 float64, h31 float64, h32 float64, h33 float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2Eddddddddd", qtrt.FFI_TYPE_POINTER, h11, h12, h13, h21, h22, h23, h31, h32, h33)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:74
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTransform(qreal, qreal, qreal, qreal, qreal, qreal)
func NewQTransform_3(h11 float64, h12 float64, h21 float64, h22 float64, dx float64, dy float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2Edddddd", qtrt.FFI_TYPE_POINTER, h11, h12, h21, h22, dx, dy)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:76
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QTransform(const QMatrix &)
func NewQTransform_4(mtx *QMatrix) *QTransform {
	var convArg0 = mtx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformC2ERK7QMatrix", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTransform)
	return gothis
}

// /usr/include/qt/QtGui/qtransform.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAffine()
func (this *QTransform) IsAffine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform8isAffineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isIdentity()
func (this *QTransform) IsIdentity() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform10isIdentityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInvertible()
func (this *QTransform) IsInvertible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform12isInvertibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isScaling()
func (this *QTransform) IsScaling() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform9isScalingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRotating()
func (this *QTransform) IsRotating() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform10isRotatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTranslating()
func (this *QTransform) IsTranslating() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform13isTranslatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtransform.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] QTransform::TransformationType type()
func (this *QTransform) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtransform.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal determinant()
func (this *QTransform) Determinant() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform11determinantEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal det()
func (this *QTransform) Det() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3detEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m11()
func (this *QTransform) M11() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m11Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m12()
func (this *QTransform) M12() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m12Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m13()
func (this *QTransform) M13() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m13Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m21()
func (this *QTransform) M21() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m21Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m22()
func (this *QTransform) M22() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m22Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m23()
func (this *QTransform) M23() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m23Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m31()
func (this *QTransform) M31() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m31Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m32()
func (this *QTransform) M32() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m32Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal m33()
func (this *QTransform) M33() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3m33Ev", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal dx()
func (this *QTransform) Dx() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform2dxEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal dy()
func (this *QTransform) Dy() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform2dyEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtransform.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QTransform) SetMatrix(m11 float64, m12 float64, m13 float64, m21 float64, m22 float64, m23 float64, m31 float64, m32 float64, m33 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform9setMatrixEddddddddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m11, m12, m13, m21, m22, m23, m31, m32, m33)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:119
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform inverted(_Bool *)
func (this *QTransform) Inverted(invertible unsafe.Pointer /*666*/) *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform8invertedEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &invertible)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:120
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform adjoint()
func (this *QTransform) Adjoint() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform7adjointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:121
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform transposed()
func (this *QTransform) Transposed() *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:123
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & translate(qreal, qreal)
func (this *QTransform) Translate(dx float64, dy float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform9translateEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:124
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & scale(qreal, qreal)
func (this *QTransform) Scale(sx float64, sy float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform5scaleEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sx, sy)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:125
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & shear(qreal, qreal)
func (this *QTransform) Shear(sh float64, sv float64) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform5shearEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, sv)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:126
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & rotate(qreal, Qt::Axis)
func (this *QTransform) Rotate(a float64, axis int) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform6rotateEdN2Qt4AxisE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, axis)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:127
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform & rotateRadians(qreal, Qt::Axis)
func (this *QTransform) RotateRadians(a float64, axis int) *QTransform {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform13rotateRadiansEdN2Qt4AxisE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, axis)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:129
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool squareToQuad(const QPolygonF &, QTransform &)
func (this *QTransform) SquareToQuad(square *QPolygonF, result *QTransform) bool {
	var convArg0 = square.GetCthis()
	var convArg1 = result.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform12squareToQuadERK9QPolygonFRS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QTransform_SquareToQuad(square *QPolygonF, result *QTransform) bool {
	var nilthis *QTransform
	rv := nilthis.SquareToQuad(square, result)
	return rv
}

// /usr/include/qt/QtGui/qtransform.h:130
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool quadToSquare(const QPolygonF &, QTransform &)
func (this *QTransform) QuadToSquare(quad *QPolygonF, result *QTransform) bool {
	var convArg0 = quad.GetCthis()
	var convArg1 = result.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform12quadToSquareERK9QPolygonFRS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QTransform_QuadToSquare(quad *QPolygonF, result *QTransform) bool {
	var nilthis *QTransform
	rv := nilthis.QuadToSquare(quad, result)
	return rv
}

// /usr/include/qt/QtGui/qtransform.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool quadToQuad(const QPolygonF &, const QPolygonF &, QTransform &)
func (this *QTransform) QuadToQuad(one *QPolygonF, two *QPolygonF, result *QTransform) bool {
	var convArg0 = one.GetCthis()
	var convArg1 = two.GetCthis()
	var convArg2 = result.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QTransform_QuadToQuad(one *QPolygonF, two *QPolygonF, result *QTransform) bool {
	var nilthis *QTransform
	rv := nilthis.QuadToQuad(one, two, result)
	return rv
}

// /usr/include/qt/QtGui/qtransform.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QTransform) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:144
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint map(const QPoint &)
func (this *QTransform) Map(p *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:145
// index:1
// Public Visibility=Default Availability=Available
// [16] QPointF map(const QPointF &)
func (this *QTransform) Map_1(p *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:146
// index:2
// Public Visibility=Default Availability=Available
// [16] QLine map(const QLine &)
func (this *QTransform) Map_2(l *qtcore.QLine) *qtcore.QLine /*123*/ {
	var convArg0 = l.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK5QLine", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLine)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:147
// index:3
// Public Visibility=Default Availability=Available
// [32] QLineF map(const QLineF &)
func (this *QTransform) Map_3(l *qtcore.QLineF) *qtcore.QLineF /*123*/ {
	var convArg0 = l.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK6QLineF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:148
// index:4
// Public Visibility=Default Availability=Available
// [8] QPolygonF map(const QPolygonF &)
func (this *QTransform) Map_4(a *QPolygonF) *QPolygonF /*123*/ {
	var convArg0 = a.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:149
// index:5
// Public Visibility=Default Availability=Available
// [8] QPolygon map(const QPolygon &)
func (this *QTransform) Map_5(a *QPolygon) *QPolygon /*123*/ {
	var convArg0 = a.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK8QPolygon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:150
// index:6
// Public Visibility=Default Availability=Available
// [8] QRegion map(const QRegion &)
func (this *QTransform) Map_6(r *QRegion) *QRegion /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:151
// index:7
// Public Visibility=Default Availability=Available
// [8] QPainterPath map(const QPainterPath &)
func (this *QTransform) Map_7(p *QPainterPath) *QPainterPath /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:155
// index:8
// Public Visibility=Default Availability=Available
// [-2] void map(int, int, int *, int *)
func (this *QTransform) Map_8(x int, y int, tx unsafe.Pointer /*666*/, ty unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapEiiPiS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, &tx, &ty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:156
// index:9
// Public Visibility=Default Availability=Available
// [-2] void map(qreal, qreal, qreal *, qreal *)
func (this *QTransform) Map_9(x float64, y float64, tx unsafe.Pointer /*666*/, ty unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform3mapEddPdS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, &tx, &ty)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtransform.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon mapToPolygon(const QRect &)
func (this *QTransform) MapToPolygon(r *qtcore.QRect) *QPolygon /*123*/ {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform12mapToPolygonERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:153
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect mapRect(const QRect &)
func (this *QTransform) MapRect(arg0 *qtcore.QRect) *qtcore.QRect /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform7mapRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:154
// index:1
// Public Visibility=Default Availability=Available
// [32] QRectF mapRect(const QRectF &)
func (this *QTransform) MapRect_1(arg0 *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform7mapRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:158
// index:0
// Public Visibility=Default Availability=Available
// [48] const QMatrix & toAffine()
func (this *QTransform) ToAffine() *QMatrix {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTransform8toAffineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtGui/qtransform.h:165
// index:0
// Public static Visibility=Default Availability=Available
// [88] QTransform fromTranslate(qreal, qreal)
func (this *QTransform) FromTranslate(dx float64, dy float64) *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform13fromTranslateEdd", qtrt.FFI_TYPE_POINTER, dx, dy)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}
func QTransform_FromTranslate(dx float64, dy float64) *QTransform /*123*/ {
	var nilthis *QTransform
	rv := nilthis.FromTranslate(dx, dy)
	return rv
}

// /usr/include/qt/QtGui/qtransform.h:166
// index:0
// Public static Visibility=Default Availability=Available
// [88] QTransform fromScale(qreal, qreal)
func (this *QTransform) FromScale(dx float64, dy float64) *QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransform9fromScaleEdd", qtrt.FFI_TYPE_POINTER, dx, dy)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}
func QTransform_FromScale(dx float64, dy float64) *QTransform /*123*/ {
	var nilthis *QTransform
	rv := nilthis.FromScale(dx, dy)
	return rv
}

func DeleteQTransform(this *QTransform) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTransformD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QTransform__TransformationType = int

const QTransform__TxNone QTransform__TransformationType = 0
const QTransform__TxTranslate QTransform__TransformationType = 1
const QTransform__TxScale QTransform__TransformationType = 2
const QTransform__TxRotate QTransform__TransformationType = 4
const QTransform__TxShear QTransform__TransformationType = 8
const QTransform__TxProject QTransform__TransformationType = 16

//  body block end