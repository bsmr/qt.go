// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qsplashscreen.h
// #include <qsplashscreen.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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

// bool event(QEvent *)
func (this *QSplashScreen) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void drawContents(QPainter *)
func (this *QSplashScreen) InheritDrawContents(f func(painter *qtgui.QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawContents", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QSplashScreen) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

/*

 */
type QSplashScreen struct {
	*QWidget
}
type QSplashScreen_ITF interface {
	QWidget_ITF
	QSplashScreen_PTR() *QSplashScreen
}

func (ptr *QSplashScreen) QSplashScreen_PTR() *QSplashScreen { return ptr }

func (this *QSplashScreen) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QSplashScreen) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQSplashScreenFromPointer(cthis unsafe.Pointer) *QSplashScreen {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSplashScreen{bcthis0}
}
func (*QSplashScreen) NewFromPointer(cthis unsafe.Pointer) *QSplashScreen {
	return NewQSplashScreenFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSplashScreen) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSplashScreen10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(const QPixmap &, Qt::WindowFlags)

/*
Construct a splash screen that will display the pixmap.

There should be no need to set the widget flags, f, except perhaps Qt::WindowStaysOnTopHint.
*/
func (*QSplashScreen) NewForInherit(pixmap qtgui.QPixmap_ITF, f int) *QSplashScreen {
	return NewQSplashScreen(pixmap, f)
}
func NewQSplashScreen(pixmap qtgui.QPixmap_ITF, f int) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2ERK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(const QPixmap &, Qt::WindowFlags)

/*
Construct a splash screen that will display the pixmap.

There should be no need to set the widget flags, f, except perhaps Qt::WindowStaysOnTopHint.
*/
func (*QSplashScreen) NewForInheritp() *QSplashScreen {
	return NewQSplashScreenp()
}
func NewQSplashScreenp() *QSplashScreen {
	// arg: 0, const QPixmap &=LValueReference, QPixmap=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2ERK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(const QPixmap &, Qt::WindowFlags)

/*
Construct a splash screen that will display the pixmap.

There should be no need to set the widget flags, f, except perhaps Qt::WindowStaysOnTopHint.
*/
func (*QSplashScreen) NewForInheritp1(pixmap qtgui.QPixmap_ITF) *QSplashScreen {
	return NewQSplashScreenp1(pixmap)
}
func NewQSplashScreenp1(pixmap qtgui.QPixmap_ITF) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2ERK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(QWidget *, const QPixmap &, Qt::WindowFlags)

/*
Construct a splash screen that will display the pixmap.

There should be no need to set the widget flags, f, except perhaps Qt::WindowStaysOnTopHint.
*/
func (*QSplashScreen) NewForInherit1(parent QWidget_ITF /*777 QWidget **/, pixmap qtgui.QPixmap_ITF, f int) *QSplashScreen {
	return NewQSplashScreen1(parent, pixmap, f)
}
func NewQSplashScreen1(parent QWidget_ITF /*777 QWidget **/, pixmap qtgui.QPixmap_ITF, f int) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2EP7QWidgetRK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(QWidget *, const QPixmap &, Qt::WindowFlags)

/*
Construct a splash screen that will display the pixmap.

There should be no need to set the widget flags, f, except perhaps Qt::WindowStaysOnTopHint.
*/
func (*QSplashScreen) NewForInherit1p(parent QWidget_ITF /*777 QWidget **/) *QSplashScreen {
	return NewQSplashScreen1p(parent)
}
func NewQSplashScreen1p(parent QWidget_ITF /*777 QWidget **/) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QPixmap &=LValueReference, QPixmap=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2EP7QWidgetRK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSplashScreen(QWidget *, const QPixmap &, Qt::WindowFlags)

/*
Construct a splash screen that will display the pixmap.

There should be no need to set the widget flags, f, except perhaps Qt::WindowStaysOnTopHint.
*/
func (*QSplashScreen) NewForInherit1p1(parent QWidget_ITF /*777 QWidget **/, pixmap qtgui.QPixmap_ITF) *QSplashScreen {
	return NewQSplashScreen1p1(parent, pixmap)
}
func NewQSplashScreen1p1(parent QWidget_ITF /*777 QWidget **/, pixmap qtgui.QPixmap_ITF) *QSplashScreen {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenC2EP7QWidgetRK7QPixmap6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplashScreenFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplashScreen")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSplashScreen()

/*

 */
func DeleteQSplashScreen(this *QSplashScreen) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreenD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(const QPixmap &)

/*
Sets the pixmap that will be used as the splash screen's image to pixmap.

See also pixmap().
*/
func (this *QSplashScreen) SetPixmap(pixmap qtgui.QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen9setPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:62
// index:0
// Public Visibility=Default Availability=Available
// [32] const QPixmap pixmap() const

/*
Returns the pixmap that is used in the splash screen. The image does not have any of the text drawn by showMessage() calls.

See also setPixmap().
*/
func (this *QSplashScreen) Pixmap() *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSplashScreen6pixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finish(QWidget *)

/*
Makes the splash screen wait until the widget mainWin is displayed before calling close() on itself.
*/
func (this *QSplashScreen) Finish(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen6finishEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void repaint()

/*
This overrides QWidget::repaint(). It differs from the standard repaint function in that it also calls QApplication::processEvents() to ensure the updates are displayed, even when there is no event loop present.
*/
func (this *QSplashScreen) Repaint() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen7repaintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QString message() const

/*
Returns the message that is currently displayed on the splash screen.

This function was introduced in  Qt 5.2.

See also showMessage() and clearMessage().
*/
func (this *QSplashScreen) Message() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSplashScreen7messageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int, const QColor &)

/*
Draws the message text onto the splash screen with color color and aligns the text according to the flags in alignment. This function calls repaint() to make sure the splash screen is repainted immediately. As a result the message is kept up to date with what your application is doing (e.g. loading files).

See also Qt::Alignment, clearMessage(), and message().
*/
func (this *QSplashScreen) ShowMessage(message string, alignment int, color qtgui.QColor_ITF) {
	var tmpArg0 = qtcore.NewQString5(message)
	var convArg0 = tmpArg0.GetCthis()
	var convArg2 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg2 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int, const QColor &)

/*
Draws the message text onto the splash screen with color color and aligns the text according to the flags in alignment. This function calls repaint() to make sure the splash screen is repainted immediately. As a result the message is kept up to date with what your application is doing (e.g. loading files).

See also Qt::Alignment, clearMessage(), and message().
*/
func (this *QSplashScreen) ShowMessagep(message string) {
	var tmpArg0 = qtcore.NewQString5(message)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	alignment := 0 /*Qt::AlignLeft*/
	// arg: 2, const QColor &=LValueReference, QColor=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int, const QColor &)

/*
Draws the message text onto the splash screen with color color and aligns the text according to the flags in alignment. This function calls repaint() to make sure the splash screen is repainted immediately. As a result the message is kept up to date with what your application is doing (e.g. loading files).

See also Qt::Alignment, clearMessage(), and message().
*/
func (this *QSplashScreen) ShowMessagep1(message string, alignment int) {
	var tmpArg0 = qtcore.NewQString5(message)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, const QColor &=LValueReference, QColor=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMessage()

/*
Removes the message being displayed on the splash screen

See also showMessage().
*/
func (this *QSplashScreen) ClearMessage() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen12clearMessageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void messageChanged(const QString &)

/*
This signal is emitted when the message on the splash screen changes. message is the new message and is a null-string when the message has been removed.

See also showMessage() and clearMessage().
*/
func (this *QSplashScreen) MessageChanged(message string) {
	var tmpArg0 = qtcore.NewQString5(message)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen14messageChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QSplashScreen) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:77
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawContents(QPainter *)

/*
Draw the contents of the splash screen using painter painter. The default implementation draws the message passed by showMessage(). Reimplement this function if you want to do your own drawing on the splash screen.
*/
func (this *QSplashScreen) DrawContents(painter qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen12drawContentsEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:78
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QSplashScreen) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSplashScreen15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
