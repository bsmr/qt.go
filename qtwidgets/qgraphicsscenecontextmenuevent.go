// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QGraphicsSceneContextMenuEvent struct {
	*QGraphicsSceneEvent
}
type QGraphicsSceneContextMenuEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneContextMenuEvent_PTR() *QGraphicsSceneContextMenuEvent
}

func (ptr *QGraphicsSceneContextMenuEvent) QGraphicsSceneContextMenuEvent_PTR() *QGraphicsSceneContextMenuEvent {
	return ptr
}

func (this *QGraphicsSceneContextMenuEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneContextMenuEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneContextMenuEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneContextMenuEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneContextMenuEvent{bcthis0}
}
func (*QGraphicsSceneContextMenuEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneContextMenuEvent {
	return NewQGraphicsSceneContextMenuEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneContextMenuEvent(QEvent::Type)

/*

 */
func (*QGraphicsSceneContextMenuEvent) NewForInherit(type_ int) *QGraphicsSceneContextMenuEvent {
	return NewQGraphicsSceneContextMenuEvent(type_)
}
func NewQGraphicsSceneContextMenuEvent(type_ int) *QGraphicsSceneContextMenuEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneContextMenuEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneContextMenuEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneContextMenuEvent(QEvent::Type)

/*

 */
func (*QGraphicsSceneContextMenuEvent) NewForInheritp() *QGraphicsSceneContextMenuEvent {
	return NewQGraphicsSceneContextMenuEventp()
}
func NewQGraphicsSceneContextMenuEventp() *QGraphicsSceneContextMenuEvent {
	// arg: 0, QEvent::Type=Enum, QEvent::Type=Enum, , Invalid
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEventC2EN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneContextMenuEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneContextMenuEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:175
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneContextMenuEvent()

/*

 */
func DeleteQGraphicsSceneContextMenuEvent(this *QGraphicsSceneContextMenuEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:177
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pos() const

/*

 */
func (this *QGraphicsSceneContextMenuEvent) Pos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPos(const QPointF &)

/*

 */
func (this *QGraphicsSceneContextMenuEvent) SetPos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent6setPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:180
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos() const

/*

 */
func (this *QGraphicsSceneContextMenuEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent8scenePosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScenePos(const QPointF &)

/*

 */
func (this *QGraphicsSceneContextMenuEvent) SetScenePos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent11setScenePosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint screenPos() const

/*

 */
func (this *QGraphicsSceneContextMenuEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent9screenPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreenPos(const QPoint &)

/*

 */
func (this *QGraphicsSceneContextMenuEvent) SetScreenPos(pos qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent12setScreenPosERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:186
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers() const

/*

 */
func (this *QGraphicsSceneContextMenuEvent) Modifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent9modifiersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModifiers(Qt::KeyboardModifiers)

/*

 */
func (this *QGraphicsSceneContextMenuEvent) SetModifiers(modifiers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modifiers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:189
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsSceneContextMenuEvent::Reason reason() const

/*

 */
func (this *QGraphicsSceneContextMenuEvent) Reason() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent6reasonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReason(QGraphicsSceneContextMenuEvent::Reason)

/*

 */
func (this *QGraphicsSceneContextMenuEvent) SetReason(reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent9setReasonENS_6ReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QGraphicsSceneContextMenuEvent__Reason = int

//
const QGraphicsSceneContextMenuEvent__Mouse QGraphicsSceneContextMenuEvent__Reason = 0

//
const QGraphicsSceneContextMenuEvent__Keyboard QGraphicsSceneContextMenuEvent__Reason = 1

//
const QGraphicsSceneContextMenuEvent__Other QGraphicsSceneContextMenuEvent__Reason = 2

func (this *QGraphicsSceneContextMenuEvent) ReasonItemName(val int) string {
	switch val {
	case QGraphicsSceneContextMenuEvent__Mouse: // 0
		return "Mouse"
	case QGraphicsSceneContextMenuEvent__Keyboard: // 1
		return "Keyboard"
	case QGraphicsSceneContextMenuEvent__Other: // 2
		return "Other"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsSceneContextMenuEvent_ReasonItemName(val int) string {
	var nilthis *QGraphicsSceneContextMenuEvent
	return nilthis.ReasonItemName(val)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
