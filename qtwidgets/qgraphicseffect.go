package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>

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
// void draw(class QPainter *)
func (this *QGraphicsEffect) InheritDraw(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "draw", f)
}

// void sourceChanged(QGraphicsEffect::ChangeFlags)
func (this *QGraphicsEffect) InheritSourceChanged(f func(flags int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sourceChanged", f)
}

// void updateBoundingRect()
func (this *QGraphicsEffect) InheritUpdateBoundingRect(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateBoundingRect", f)
}

// bool sourceIsPixmap()
func (this *QGraphicsEffect) InheritSourceIsPixmap(f func() bool) {
	qtrt.SetAllInheritCallback(this, "sourceIsPixmap", f)
}

// QRectF sourceBoundingRect(Qt::CoordinateSystem)
func (this *QGraphicsEffect) InheritSourceBoundingRect(f func(system int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sourceBoundingRect", f)
}

// void drawSource(class QPainter *)
func (this *QGraphicsEffect) InheritDrawSource(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawSource", f)
}

// QPixmap sourcePixmap(Qt::CoordinateSystem, class QPoint *, enum QGraphicsEffect::PixmapPadMode)
func (this *QGraphicsEffect) InheritSourcePixmap(f func(system int, offset *qtcore.QPoint /*777 QPoint **/, mode int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sourcePixmap", f)
}

type QGraphicsEffect struct {
	*qtcore.QObject
}

func (this *QGraphicsEffect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGraphicsEffect) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGraphicsEffectFromPointer(cthis unsafe.Pointer) *QGraphicsEffect {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsEffect{bcthis0}
}
func (*QGraphicsEffect) NewFromPointer(cthis unsafe.Pointer) *QGraphicsEffect {
	return NewQGraphicsEffectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsEffect(QObject *)
func NewQGraphicsEffect(parent *qtcore.QObject /*777 QObject **/) *QGraphicsEffect {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsEffect()
func DeleteQGraphicsEffect(this *QGraphicsEffect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRectFor(const QRectF &)
func (this *QGraphicsEffect) BoundingRectFor(sourceRect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = sourceRect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect15boundingRectForERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:86
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGraphicsEffect) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QGraphicsEffect) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QGraphicsEffect) SetEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()
func (this *QGraphicsEffect) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void enabledChanged(_Bool)
func (this *QGraphicsEffect) EnabledChanged(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect14enabledChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:99
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void draw(QPainter *)
func (this *QGraphicsEffect) Draw(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect4drawEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:100
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void sourceChanged(QGraphicsEffect::ChangeFlags)
func (this *QGraphicsEffect) SourceChanged(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:101
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateBoundingRect()
func (this *QGraphicsEffect) UpdateBoundingRect() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect18updateBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:103
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool sourceIsPixmap()
func (this *QGraphicsEffect) SourceIsPixmap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect14sourceIsPixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:104
// index:0
// Protected Visibility=Default Availability=Available
// [32] QRectF sourceBoundingRect(Qt::CoordinateSystem)
func (this *QGraphicsEffect) SourceBoundingRect(system int) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:105
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawSource(QPainter *)
func (this *QGraphicsEffect) DrawSource(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsEffect10drawSourceEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [32] QPixmap sourcePixmap(Qt::CoordinateSystem, QPoint *, enum QGraphicsEffect::PixmapPadMode)
func (this *QGraphicsEffect) SourcePixmap(system int, offset *qtcore.QPoint /*777 QPoint **/, mode int) *qtgui.QPixmap /*123*/ {
	var convArg1 = offset.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), system, convArg1, mode)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

type QGraphicsEffect__ChangeFlag = int

const QGraphicsEffect__SourceAttached QGraphicsEffect__ChangeFlag = 1
const QGraphicsEffect__SourceDetached QGraphicsEffect__ChangeFlag = 2
const QGraphicsEffect__SourceBoundingRectChanged QGraphicsEffect__ChangeFlag = 4
const QGraphicsEffect__SourceInvalidated QGraphicsEffect__ChangeFlag = 8

type QGraphicsEffect__PixmapPadMode = int

const QGraphicsEffect__NoPad QGraphicsEffect__PixmapPadMode = 0
const QGraphicsEffect__PadToTransparentBorder QGraphicsEffect__PixmapPadMode = 1
const QGraphicsEffect__PadToEffectiveBoundingRect QGraphicsEffect__PixmapPadMode = 2

//  body block end