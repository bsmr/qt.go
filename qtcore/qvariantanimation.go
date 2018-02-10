package qtcore

// /usr/include/qt/QtCore/qvariantanimation.h
// #include <qvariantanimation.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
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
// bool event(class QEvent *)
func (this *QVariantAnimation) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentTime(int)
func (this *QVariantAnimation) InheritUpdateCurrentTime(f func(arg0 int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentTime", f)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QVariantAnimation) InheritUpdateState(f func(newState int, oldState int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateState", f)
}

// void updateCurrentValue(const class QVariant &)
func (this *QVariantAnimation) InheritUpdateCurrentValue(f func(value *QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentValue", f)
}

// QVariant interpolated(const class QVariant &, const class QVariant &, qreal)
func (this *QVariantAnimation) InheritInterpolated(f func(from *QVariant, to *QVariant, progress float64) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "interpolated", f)
}

type QVariantAnimation struct {
	*QAbstractAnimation
}

func (this *QVariantAnimation) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractAnimation.GetCthis()
	}
}
func (this *QVariantAnimation) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractAnimation = NewQAbstractAnimationFromPointer(cthis)
}
func NewQVariantAnimationFromPointer(cthis unsafe.Pointer) *QVariantAnimation {
	bcthis0 := NewQAbstractAnimationFromPointer(cthis)
	return &QVariantAnimation{bcthis0}
}
func (*QVariantAnimation) NewFromPointer(cthis unsafe.Pointer) *QVariantAnimation {
	return NewQVariantAnimationFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariantanimation.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QVariantAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qvariantanimation.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVariantAnimation(QObject *)
func NewQVariantAnimation(parent *QObject /*777 QObject **/) *QVariantAnimation {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qvariantanimation.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QVariantAnimation()
func DeleteQVariantAnimation(this *QVariantAnimation) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qvariantanimation.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant startValue()
func (this *QVariantAnimation) StartValue() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation10startValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartValue(const QVariant &)
func (this *QVariantAnimation) SetStartValue(value *QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation13setStartValueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:74
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant endValue()
func (this *QVariantAnimation) EndValue() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation8endValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEndValue(const QVariant &)
func (this *QVariantAnimation) SetEndValue(value *QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation11setEndValueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:77
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant keyValueAt(qreal)
func (this *QVariantAnimation) KeyValueAt(step float64) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation10keyValueAtEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyValueAt(qreal, const QVariant &)
func (this *QVariantAnimation) SetKeyValueAt(step float64, value *QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation13setKeyValueAtEdRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), step, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:83
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant currentValue()
func (this *QVariantAnimation) CurrentValue() *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation12currentValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int duration()
func (this *QVariantAnimation) Duration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qvariantanimation.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDuration(int)
func (this *QVariantAnimation) SetDuration(msecs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation11setDurationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QEasingCurve easingCurve()
func (this *QVariantAnimation) EasingCurve() *QEasingCurve /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation11easingCurveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQEasingCurve)
	return rv2
}

// /usr/include/qt/QtCore/qvariantanimation.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEasingCurve(const QEasingCurve &)
func (this *QVariantAnimation) SetEasingCurve(easing *QEasingCurve) {
	var convArg0 = easing.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation14setEasingCurveERK12QEasingCurve", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(const QVariant &)
func (this *QVariantAnimation) ValueChanged(value *QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation12valueChangedERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QVariantAnimation) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariantanimation.h:100
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentTime(int)
func (this *QVariantAnimation) UpdateCurrentTime(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation17updateCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:101
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateState(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QVariantAnimation) UpdateState(newState int, oldState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation11updateStateEN18QAbstractAnimation5StateES1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:103
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentValue(const QVariant &)
func (this *QVariantAnimation) UpdateCurrentValue(value *QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QVariantAnimation18updateCurrentValueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariantanimation.h:104
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant interpolated(const QVariant &, const QVariant &, qreal)
func (this *QVariantAnimation) Interpolated(from *QVariant, to *QVariant, progress float64) *QVariant /*123*/ {
	var convArg0 = from.GetCthis()
	var convArg1 = to.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QVariantAnimation12interpolatedERK8QVariantS2_d", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, progress)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

//  body block end