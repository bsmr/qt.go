package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtGui/qwindow.h
// dst-file: /src/gui/qwindow.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QWindow::unsetCursor();
extern void _ZN7QWindow11unsetCursorEv(void* qthis);
  // proto:  bool QWindow::isVisible();
extern void _ZNK7QWindow9isVisibleEv(void* qthis);
  // proto:  void QWindow::setScreen(QScreen * screen);
extern void _ZN7QWindow9setScreenEP7QScreen(void* qthis, void* arg0);
  // proto:  QSize QWindow::maximumSize();
extern void _ZNK7QWindow11maximumSizeEv(void* qthis);
  // proto:  void QWindow::setTransientParent(QWindow * parent);
extern void _ZN7QWindow18setTransientParentEPS_(void* qthis, void* arg0);
  // proto:  QSurfaceFormat QWindow::format();
extern void _ZNK7QWindow6formatEv(void* qthis);
  // proto:  bool QWindow::isTopLevel();
extern void _ZNK7QWindow10isTopLevelEv(void* qthis);
  // proto:  void QWindow::QWindow(const QWindow & );
extern void* dector_ZN7QWindowC1ERKS_(void* arg0);
extern void _ZN7QWindowC1ERKS_(void* qthis, void* arg0);
  // proto:  void QWindow::setIcon(const QIcon & icon);
extern void _ZN7QWindow7setIconERK5QIcon(void* qthis, void* arg0);
  // proto:  qreal QWindow::opacity();
extern void _ZNK7QWindow7opacityEv(void* qthis);
  // proto:  void QWindow::setMinimumSize(const QSize & size);
extern void _ZN7QWindow14setMinimumSizeERK5QSize(void* qthis, void* arg0);
  // proto:  int QWindow::minimumHeight();
extern void demth_ZNK7QWindow13minimumHeightEv(void* qthis);
  // proto:  QSize QWindow::sizeIncrement();
extern void _ZNK7QWindow13sizeIncrementEv(void* qthis);
  // proto:  void QWindow::resize(const QSize & newSize);
extern void _ZN7QWindow6resizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QWindow::setTitle(const QString & );
extern void _ZN7QWindow8setTitleERK7QString(void* qthis, void* arg0);
  // proto:  void QWindow::raise();
extern void _ZN7QWindow5raiseEv(void* qthis);
  // proto:  QSize QWindow::minimumSize();
extern void _ZNK7QWindow11minimumSizeEv(void* qthis);
  // proto:  QPoint QWindow::mapToGlobal(const QPoint & pos);
extern void _ZNK7QWindow11mapToGlobalERK6QPoint(void* qthis, void* arg0);
  // proto: static QWindow * QWindow::fromWinId(WId id);
extern void _ZN7QWindow9fromWinIdEi(int32_t* arg0);
  // proto:  QMargins QWindow::frameMargins();
extern void _ZNK7QWindow12frameMarginsEv(void* qthis);
  // proto:  void QWindow::setMaximumWidth(int w);
extern void _ZN7QWindow15setMaximumWidthEi(void* qthis, int32_t arg0);
  // proto:  int QWindow::maximumHeight();
extern void demth_ZNK7QWindow13maximumHeightEv(void* qthis);
  // proto:  bool QWindow::isModal();
extern void _ZNK7QWindow7isModalEv(void* qthis);
  // proto:  QRect QWindow::geometry();
extern void _ZNK7QWindow8geometryEv(void* qthis);
  // proto:  void QWindow::setParent(QWindow * parent);
extern void _ZN7QWindow9setParentEPS_(void* qthis, void* arg0);
  // proto:  QRect QWindow::frameGeometry();
extern void _ZNK7QWindow13frameGeometryEv(void* qthis);
  // proto:  QSurfaceFormat QWindow::requestedFormat();
extern void _ZNK7QWindow15requestedFormatEv(void* qthis);
  // proto:  void QWindow::setHeight(int arg);
extern void _ZN7QWindow9setHeightEi(void* qthis, int32_t arg0);
  // proto:  void QWindow::requestActivate();
extern void _ZN7QWindow15requestActivateEv(void* qthis);
  // proto:  QPoint QWindow::mapFromGlobal(const QPoint & pos);
extern void _ZNK7QWindow13mapFromGlobalERK6QPoint(void* qthis, void* arg0);
  // proto:  int QWindow::y();
extern void demth_ZNK7QWindow1yEv(void* qthis);
  // proto:  int QWindow::width();
extern void demth_ZNK7QWindow5widthEv(void* qthis);
  // proto:  void QWindow::setFilePath(const QString & filePath);
extern void _ZN7QWindow11setFilePathERK7QString(void* qthis, void* arg0);
  // proto:  void QWindow::setCursor(const QCursor & );
extern void _ZN7QWindow9setCursorERK7QCursor(void* qthis, void* arg0);
  // proto:  void QWindow::setVisible(bool visible);
extern void _ZN7QWindow10setVisibleEb(void* qthis, bool arg0);
  // proto:  void QWindow::~QWindow();
extern void _ZN7QWindowD0Ev(void* qthis);
  // proto:  bool QWindow::setMouseGrabEnabled(bool grab);
extern void _ZN7QWindow19setMouseGrabEnabledEb(void* qthis, bool arg0);
  // proto:  bool QWindow::isExposed();
extern void _ZNK7QWindow9isExposedEv(void* qthis);
  // proto:  int QWindow::minimumWidth();
extern void demth_ZNK7QWindow12minimumWidthEv(void* qthis);
  // proto:  void QWindow::setPosition(const QPoint & pt);
extern void _ZN7QWindow11setPositionERK6QPoint(void* qthis, void* arg0);
  // proto:  bool QWindow::close();
extern void _ZN7QWindow5closeEv(void* qthis);
  // proto:  int QWindow::x();
extern void demth_ZNK7QWindow1xEv(void* qthis);
  // proto:  void QWindow::setMinimumWidth(int w);
extern void _ZN7QWindow15setMinimumWidthEi(void* qthis, int32_t arg0);
  // proto:  QRegion QWindow::mask();
extern void _ZNK7QWindow4maskEv(void* qthis);
  // proto:  QWindow * QWindow::parent();
extern void _ZNK7QWindow6parentEv(void* qthis);
  // proto:  void QWindow::setFramePosition(const QPoint & point);
extern void _ZN7QWindow16setFramePositionERK6QPoint(void* qthis, void* arg0);
  // proto:  void QWindow::QWindow(QScreen * screen);
extern void* dector_ZN7QWindowC1EP7QScreen(void* arg0);
extern void _ZN7QWindowC1EP7QScreen(void* qthis, void* arg0);
  // proto:  void QWindow::setGeometry(int posx, int posy, int w, int h);
extern void _ZN7QWindow11setGeometryEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  bool QWindow::setKeyboardGrabEnabled(bool grab);
extern void _ZN7QWindow22setKeyboardGrabEnabledEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QWindow::metaObject();
extern void _ZNK7QWindow10metaObjectEv(void* qthis);
  // proto:  void QWindow::QWindow(QWindow * parent);
extern void* dector_ZN7QWindowC1EPS_(void* arg0);
extern void _ZN7QWindowC1EPS_(void* qthis, void* arg0);
  // proto:  void QWindow::setWidth(int arg);
extern void _ZN7QWindow8setWidthEi(void* qthis, int32_t arg0);
  // proto:  void QWindow::setY(int arg);
extern void _ZN7QWindow4setYEi(void* qthis, int32_t arg0);
  // proto:  qreal QWindow::devicePixelRatio();
extern void _ZNK7QWindow16devicePixelRatioEv(void* qthis);
  // proto:  void QWindow::setBaseSize(const QSize & size);
extern void _ZN7QWindow11setBaseSizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QWindow::alert(int msec);
extern void _ZN7QWindow5alertEi(void* qthis, int32_t arg0);
  // proto:  QPlatformWindow * QWindow::handle();
extern void _ZNK7QWindow6handleEv(void* qthis);
  // proto:  void QWindow::destroy();
extern void _ZN7QWindow7destroyEv(void* qthis);
  // proto:  QWindow * QWindow::transientParent();
extern void _ZNK7QWindow15transientParentEv(void* qthis);
  // proto:  void QWindow::setMinimumHeight(int h);
extern void _ZN7QWindow16setMinimumHeightEi(void* qthis, int32_t arg0);
  // proto:  void QWindow::show();
extern void _ZN7QWindow4showEv(void* qthis);
  // proto:  QSize QWindow::baseSize();
extern void _ZNK7QWindow8baseSizeEv(void* qthis);
  // proto:  QString QWindow::title();
extern void _ZNK7QWindow5titleEv(void* qthis);
  // proto:  void QWindow::showMaximized();
extern void _ZN7QWindow13showMaximizedEv(void* qthis);
  // proto:  void QWindow::create();
extern void _ZN7QWindow6createEv(void* qthis);
  // proto:  void QWindow::resize(int w, int h);
extern void _ZN7QWindow6resizeEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QScreen * QWindow::screen();
extern void _ZNK7QWindow6screenEv(void* qthis);
  // proto:  void QWindow::setPosition(int posx, int posy);
extern void _ZN7QWindow11setPositionEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QWindow::setOpacity(qreal level);
extern void _ZN7QWindow10setOpacityEd(void* qthis, double arg0);
  // proto:  void QWindow::setGeometry(const QRect & rect);
extern void _ZN7QWindow11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  void QWindow::setSizeIncrement(const QSize & size);
extern void _ZN7QWindow16setSizeIncrementERK5QSize(void* qthis, void* arg0);
  // proto:  void QWindow::showMinimized();
extern void _ZN7QWindow13showMinimizedEv(void* qthis);
  // proto:  QObject * QWindow::focusObject();
extern void _ZNK7QWindow11focusObjectEv(void* qthis);
  // proto:  bool QWindow::isActive();
extern void _ZNK7QWindow8isActiveEv(void* qthis);
  // proto:  QAccessibleInterface * QWindow::accessibleRoot();
extern void _ZNK7QWindow14accessibleRootEv(void* qthis);
  // proto:  QCursor QWindow::cursor();
extern void _ZNK7QWindow6cursorEv(void* qthis);
  // proto:  void QWindow::setFormat(const QSurfaceFormat & format);
extern void _ZN7QWindow9setFormatERK14QSurfaceFormat(void* qthis, void* arg0);
  // proto:  void QWindow::showFullScreen();
extern void _ZN7QWindow14showFullScreenEv(void* qthis);
  // proto:  void QWindow::setX(int arg);
extern void _ZN7QWindow4setXEi(void* qthis, int32_t arg0);
  // proto:  void QWindow::lower();
extern void _ZN7QWindow5lowerEv(void* qthis);
  // proto:  void QWindow::requestUpdate();
extern void _ZN7QWindow13requestUpdateEv(void* qthis);
  // proto:  void QWindow::hide();
extern void _ZN7QWindow4hideEv(void* qthis);
  // proto:  void QWindow::setMask(const QRegion & region);
extern void _ZN7QWindow7setMaskERK7QRegion(void* qthis, void* arg0);
  // proto:  void QWindow::setMaximumSize(const QSize & size);
extern void _ZN7QWindow14setMaximumSizeERK5QSize(void* qthis, void* arg0);
  // proto:  int QWindow::height();
extern void demth_ZNK7QWindow6heightEv(void* qthis);
  // proto:  QSize QWindow::size();
extern void demth_ZNK7QWindow4sizeEv(void* qthis);
  // proto:  int QWindow::maximumWidth();
extern void demth_ZNK7QWindow12maximumWidthEv(void* qthis);
  // proto:  QPoint QWindow::position();
extern void demth_ZNK7QWindow8positionEv(void* qthis);
  // proto:  void QWindow::setMaximumHeight(int h);
extern void _ZN7QWindow16setMaximumHeightEi(void* qthis, int32_t arg0);
  // proto:  QString QWindow::filePath();
extern void _ZNK7QWindow8filePathEv(void* qthis);
  // proto:  void QWindow::showNormal();
extern void _ZN7QWindow10showNormalEv(void* qthis);
  // proto:  QPoint QWindow::framePosition();
extern void _ZNK7QWindow13framePositionEv(void* qthis);
  // proto:  QIcon QWindow::icon();
extern void _ZNK7QWindow4iconEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QWindow)=1
type QWindow struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _modalityChanged QWindow_modalityChanged_signal;
//  _activeChanged QWindow_activeChanged_signal;
//  _heightChanged QWindow_heightChanged_signal;
//  _contentOrientationChanged QWindow_contentOrientationChanged_signal;
//  _minimumWidthChanged QWindow_minimumWidthChanged_signal;
//  _opacityChanged QWindow_opacityChanged_signal;
//  _visibleChanged QWindow_visibleChanged_signal;
//  _screenChanged QWindow_screenChanged_signal;
//  _maximumHeightChanged QWindow_maximumHeightChanged_signal;
//  _yChanged QWindow_yChanged_signal;
//  _widthChanged QWindow_widthChanged_signal;
//  _windowStateChanged QWindow_windowStateChanged_signal;
//  _windowTitleChanged QWindow_windowTitleChanged_signal;
//  _visibilityChanged QWindow_visibilityChanged_signal;
//  _minimumHeightChanged QWindow_minimumHeightChanged_signal;
//  _xChanged QWindow_xChanged_signal;
//  _focusObjectChanged QWindow_focusObjectChanged_signal;
//  _maximumWidthChanged QWindow_maximumWidthChanged_signal;
}

  // proto:  void QWindow::unsetCursor();
func (this *QWindow) unsetCursor(args ...interface{}) () {
  // unsetCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11unsetCursorEv
    // invoke: void unsetCursor()
    C._ZN7QWindow11unsetCursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "unsetCursor", args)
  }

}

  // proto:  bool QWindow::isVisible();
func (this *QWindow) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow9isVisibleEv
    // invoke: bool isVisible()
    C._ZNK7QWindow9isVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isVisible", args)
  }

}

  // proto:  void QWindow::setScreen(QScreen * screen);
func (this *QWindow) setScreen(args ...interface{}) () {
  // setScreen(class QScreen *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScreen{}) // "QScreen *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setScreenEP7QScreen
    // invoke: void setScreen(class QScreen *)
    var arg0 = args[0].(QScreen).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow9setScreenEP7QScreen(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setScreen", args)
  }

}

  // proto:  QSize QWindow::maximumSize();
func (this *QWindow) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11maximumSizeEv
    // invoke: QSize maximumSize()
    C._ZNK7QWindow11maximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "maximumSize", args)
  }

}

  // proto:  void QWindow::setTransientParent(QWindow * parent);
func (this *QWindow) setTransientParent(args ...interface{}) () {
  // setTransientParent(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow18setTransientParentEPS_
    // invoke: void setTransientParent(class QWindow *)
    var arg0 = args[0].(QWindow).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow18setTransientParentEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setTransientParent", args)
  }

}

  // proto:  QSurfaceFormat QWindow::format();
func (this *QWindow) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6formatEv
    // invoke: QSurfaceFormat format()
    C._ZNK7QWindow6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "format", args)
  }

}

  // proto:  bool QWindow::isTopLevel();
func (this *QWindow) isTopLevel(args ...interface{}) () {
  // isTopLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow10isTopLevelEv
    // invoke: bool isTopLevel()
    C._ZNK7QWindow10isTopLevelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isTopLevel", args)
  }

}

  // proto:  void QWindow::QWindow(const QWindow & );
func NewQWindow(args ...interface{}) QWindow {
  return QWindow{}
}

  // proto:  void QWindow::setIcon(const QIcon & icon);
func (this *QWindow) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setIcon", args)
  }

}

  // proto:  qreal QWindow::opacity();
func (this *QWindow) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow7opacityEv
    // invoke: qreal opacity()
    C._ZNK7QWindow7opacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "opacity", args)
  }

}

  // proto:  void QWindow::setMinimumSize(const QSize & size);
func (this *QWindow) setMinimumSize(args ...interface{}) () {
  // setMinimumSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow14setMinimumSizeERK5QSize
    // invoke: void setMinimumSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow14setMinimumSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMinimumSize", args)
  }

}

  // proto:  int QWindow::minimumHeight();
func (this *QWindow) minimumHeight(args ...interface{}) () {
  // minimumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13minimumHeightEv
    // invoke: int minimumHeight()
    C.demth_ZNK7QWindow13minimumHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "minimumHeight", args)
  }

}

  // proto:  QSize QWindow::sizeIncrement();
func (this *QWindow) sizeIncrement(args ...interface{}) () {
  // sizeIncrement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13sizeIncrementEv
    // invoke: QSize sizeIncrement()
    C._ZNK7QWindow13sizeIncrementEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "sizeIncrement", args)
  }

}

  // proto:  void QWindow::resize(const QSize & newSize);
func (this *QWindow) resize(args ...interface{}) () {
  // resize(const class QSize &)
  // resize(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow6resizeERK5QSize
    // invoke: void resize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow6resizeERK5QSize(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWindow6resizeEii
    // invoke: void resize(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWindow6resizeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWindow", "resize", args)
  }

}

  // proto:  void QWindow::setTitle(const QString & );
func (this *QWindow) setTitle(args ...interface{}) () {
  // setTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow8setTitleERK7QString
    // invoke: void setTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow8setTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setTitle", args)
  }

}

  // proto:  void QWindow::raise();
func (this *QWindow) raise(args ...interface{}) () {
  // raise()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow5raiseEv
    // invoke: void raise()
    C._ZN7QWindow5raiseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "raise", args)
  }

}

  // proto:  QSize QWindow::minimumSize();
func (this *QWindow) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11minimumSizeEv
    // invoke: QSize minimumSize()
    C._ZNK7QWindow11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "minimumSize", args)
  }

}

  // proto:  QPoint QWindow::mapToGlobal(const QPoint & pos);
func (this *QWindow) mapToGlobal(args ...interface{}) () {
  // mapToGlobal(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11mapToGlobalERK6QPoint
    // invoke: QPoint mapToGlobal(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWindow11mapToGlobalERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "mapToGlobal", args)
  }

}

  // proto: static QWindow * QWindow::fromWinId(WId id);
func (this *QWindow) fromWinId_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWindow", "fromWinId", args)
  }

}

  // proto:  QMargins QWindow::frameMargins();
func (this *QWindow) frameMargins(args ...interface{}) () {
  // frameMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow12frameMarginsEv
    // invoke: QMargins frameMargins()
    C._ZNK7QWindow12frameMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "frameMargins", args)
  }

}

  // proto:  void QWindow::setMaximumWidth(int w);
func (this *QWindow) setMaximumWidth(args ...interface{}) () {
  // setMaximumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow15setMaximumWidthEi
    // invoke: void setMaximumWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow15setMaximumWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMaximumWidth", args)
  }

}

  // proto:  int QWindow::maximumHeight();
func (this *QWindow) maximumHeight(args ...interface{}) () {
  // maximumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13maximumHeightEv
    // invoke: int maximumHeight()
    C.demth_ZNK7QWindow13maximumHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "maximumHeight", args)
  }

}

  // proto:  bool QWindow::isModal();
func (this *QWindow) isModal(args ...interface{}) () {
  // isModal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow7isModalEv
    // invoke: bool isModal()
    C._ZNK7QWindow7isModalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isModal", args)
  }

}

  // proto:  QRect QWindow::geometry();
func (this *QWindow) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8geometryEv
    // invoke: QRect geometry()
    C._ZNK7QWindow8geometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "geometry", args)
  }

}

  // proto:  void QWindow::setParent(QWindow * parent);
func (this *QWindow) setParent(args ...interface{}) () {
  // setParent(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setParentEPS_
    // invoke: void setParent(class QWindow *)
    var arg0 = args[0].(QWindow).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow9setParentEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setParent", args)
  }

}

  // proto:  QRect QWindow::frameGeometry();
func (this *QWindow) frameGeometry(args ...interface{}) () {
  // frameGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13frameGeometryEv
    // invoke: QRect frameGeometry()
    C._ZNK7QWindow13frameGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "frameGeometry", args)
  }

}

  // proto:  QSurfaceFormat QWindow::requestedFormat();
func (this *QWindow) requestedFormat(args ...interface{}) () {
  // requestedFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow15requestedFormatEv
    // invoke: QSurfaceFormat requestedFormat()
    C._ZNK7QWindow15requestedFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "requestedFormat", args)
  }

}

  // proto:  void QWindow::setHeight(int arg);
func (this *QWindow) setHeight(args ...interface{}) () {
  // setHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setHeightEi
    // invoke: void setHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow9setHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setHeight", args)
  }

}

  // proto:  void QWindow::requestActivate();
func (this *QWindow) requestActivate(args ...interface{}) () {
  // requestActivate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow15requestActivateEv
    // invoke: void requestActivate()
    C._ZN7QWindow15requestActivateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "requestActivate", args)
  }

}

  // proto:  QPoint QWindow::mapFromGlobal(const QPoint & pos);
func (this *QWindow) mapFromGlobal(args ...interface{}) () {
  // mapFromGlobal(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13mapFromGlobalERK6QPoint
    // invoke: QPoint mapFromGlobal(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWindow13mapFromGlobalERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "mapFromGlobal", args)
  }

}

  // proto:  int QWindow::y();
func (this *QWindow) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow1yEv
    // invoke: int y()
    C.demth_ZNK7QWindow1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "y", args)
  }

}

  // proto:  int QWindow::width();
func (this *QWindow) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow5widthEv
    // invoke: int width()
    C.demth_ZNK7QWindow5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "width", args)
  }

}

  // proto:  void QWindow::setFilePath(const QString & filePath);
func (this *QWindow) setFilePath(args ...interface{}) () {
  // setFilePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11setFilePathERK7QString
    // invoke: void setFilePath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow11setFilePathERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setFilePath", args)
  }

}

  // proto:  void QWindow::setCursor(const QCursor & );
func (this *QWindow) setCursor(args ...interface{}) () {
  // setCursor(const class QCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setCursorERK7QCursor
    // invoke: void setCursor(const class QCursor &)
    var arg0 = args[0].(QCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow9setCursorERK7QCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setCursor", args)
  }

}

  // proto:  void QWindow::setVisible(bool visible);
func (this *QWindow) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setVisible", args)
  }

}

  // proto:  void QWindow::~QWindow();
func (this *QWindow) FreeQWindow(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWindow", "~QWindow", args)
  }

}

  // proto:  bool QWindow::setMouseGrabEnabled(bool grab);
func (this *QWindow) setMouseGrabEnabled(args ...interface{}) () {
  // setMouseGrabEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow19setMouseGrabEnabledEb
    // invoke: bool setMouseGrabEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow19setMouseGrabEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMouseGrabEnabled", args)
  }

}

  // proto:  bool QWindow::isExposed();
func (this *QWindow) isExposed(args ...interface{}) () {
  // isExposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow9isExposedEv
    // invoke: bool isExposed()
    C._ZNK7QWindow9isExposedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isExposed", args)
  }

}

  // proto:  int QWindow::minimumWidth();
func (this *QWindow) minimumWidth(args ...interface{}) () {
  // minimumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow12minimumWidthEv
    // invoke: int minimumWidth()
    C.demth_ZNK7QWindow12minimumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "minimumWidth", args)
  }

}

  // proto:  void QWindow::setPosition(const QPoint & pt);
func (this *QWindow) setPosition(args ...interface{}) () {
  // setPosition(const class QPoint &)
  // setPosition(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11setPositionERK6QPoint
    // invoke: void setPosition(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow11setPositionERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWindow11setPositionEii
    // invoke: void setPosition(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWindow11setPositionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWindow", "setPosition", args)
  }

}

  // proto:  bool QWindow::close();
func (this *QWindow) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow5closeEv
    // invoke: bool close()
    C._ZN7QWindow5closeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "close", args)
  }

}

  // proto:  int QWindow::x();
func (this *QWindow) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow1xEv
    // invoke: int x()
    C.demth_ZNK7QWindow1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "x", args)
  }

}

  // proto:  void QWindow::setMinimumWidth(int w);
func (this *QWindow) setMinimumWidth(args ...interface{}) () {
  // setMinimumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow15setMinimumWidthEi
    // invoke: void setMinimumWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow15setMinimumWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMinimumWidth", args)
  }

}

  // proto:  QRegion QWindow::mask();
func (this *QWindow) mask(args ...interface{}) () {
  // mask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow4maskEv
    // invoke: QRegion mask()
    C._ZNK7QWindow4maskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "mask", args)
  }

}

  // proto:  QWindow * QWindow::parent();
func (this *QWindow) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6parentEv
    // invoke: QWindow * parent()
    C._ZNK7QWindow6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "parent", args)
  }

}

  // proto:  void QWindow::setFramePosition(const QPoint & point);
func (this *QWindow) setFramePosition(args ...interface{}) () {
  // setFramePosition(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow16setFramePositionERK6QPoint
    // invoke: void setFramePosition(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow16setFramePositionERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setFramePosition", args)
  }

}

  // proto:  void QWindow::setGeometry(int posx, int posy, int w, int h);
func (this *QWindow) setGeometry(args ...interface{}) () {
  // setGeometry(int, int, int, int)
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11setGeometryEiiii
    // invoke: void setGeometry(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN7QWindow11setGeometryEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN7QWindow11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setGeometry", args)
  }

}

  // proto:  bool QWindow::setKeyboardGrabEnabled(bool grab);
func (this *QWindow) setKeyboardGrabEnabled(args ...interface{}) () {
  // setKeyboardGrabEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow22setKeyboardGrabEnabledEb
    // invoke: bool setKeyboardGrabEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow22setKeyboardGrabEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setKeyboardGrabEnabled", args)
  }

}

  // proto:  const QMetaObject * QWindow::metaObject();
func (this *QWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK7QWindow10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "metaObject", args)
  }

}

  // proto:  void QWindow::setWidth(int arg);
func (this *QWindow) setWidth(args ...interface{}) () {
  // setWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow8setWidthEi
    // invoke: void setWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow8setWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setWidth", args)
  }

}

  // proto:  void QWindow::setY(int arg);
func (this *QWindow) setY(args ...interface{}) () {
  // setY(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow4setYEi
    // invoke: void setY(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow4setYEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setY", args)
  }

}

  // proto:  qreal QWindow::devicePixelRatio();
func (this *QWindow) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow16devicePixelRatioEv
    // invoke: qreal devicePixelRatio()
    C._ZNK7QWindow16devicePixelRatioEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "devicePixelRatio", args)
  }

}

  // proto:  void QWindow::setBaseSize(const QSize & size);
func (this *QWindow) setBaseSize(args ...interface{}) () {
  // setBaseSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow11setBaseSizeERK5QSize
    // invoke: void setBaseSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow11setBaseSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setBaseSize", args)
  }

}

  // proto:  void QWindow::alert(int msec);
func (this *QWindow) alert(args ...interface{}) () {
  // alert(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow5alertEi
    // invoke: void alert(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow5alertEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "alert", args)
  }

}

  // proto:  QPlatformWindow * QWindow::handle();
func (this *QWindow) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6handleEv
    // invoke: QPlatformWindow * handle()
    C._ZNK7QWindow6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "handle", args)
  }

}

  // proto:  void QWindow::destroy();
func (this *QWindow) destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow7destroyEv
    // invoke: void destroy()
    C._ZN7QWindow7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "destroy", args)
  }

}

  // proto:  QWindow * QWindow::transientParent();
func (this *QWindow) transientParent(args ...interface{}) () {
  // transientParent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow15transientParentEv
    // invoke: QWindow * transientParent()
    C._ZNK7QWindow15transientParentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "transientParent", args)
  }

}

  // proto:  void QWindow::setMinimumHeight(int h);
func (this *QWindow) setMinimumHeight(args ...interface{}) () {
  // setMinimumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow16setMinimumHeightEi
    // invoke: void setMinimumHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow16setMinimumHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMinimumHeight", args)
  }

}

  // proto:  void QWindow::show();
func (this *QWindow) show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow4showEv
    // invoke: void show()
    C._ZN7QWindow4showEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "show", args)
  }

}

  // proto:  QSize QWindow::baseSize();
func (this *QWindow) baseSize(args ...interface{}) () {
  // baseSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8baseSizeEv
    // invoke: QSize baseSize()
    C._ZNK7QWindow8baseSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "baseSize", args)
  }

}

  // proto:  QString QWindow::title();
func (this *QWindow) title(args ...interface{}) () {
  // title()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow5titleEv
    // invoke: QString title()
    C._ZNK7QWindow5titleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "title", args)
  }

}

  // proto:  void QWindow::showMaximized();
func (this *QWindow) showMaximized(args ...interface{}) () {
  // showMaximized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow13showMaximizedEv
    // invoke: void showMaximized()
    C._ZN7QWindow13showMaximizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "showMaximized", args)
  }

}

  // proto:  void QWindow::create();
func (this *QWindow) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow6createEv
    // invoke: void create()
    C._ZN7QWindow6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "create", args)
  }

}

  // proto:  QScreen * QWindow::screen();
func (this *QWindow) screen(args ...interface{}) () {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6screenEv
    // invoke: QScreen * screen()
    C._ZNK7QWindow6screenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "screen", args)
  }

}

  // proto:  void QWindow::setOpacity(qreal level);
func (this *QWindow) setOpacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow10setOpacityEd
    // invoke: void setOpacity(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow10setOpacityEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setOpacity", args)
  }

}

  // proto:  void QWindow::setSizeIncrement(const QSize & size);
func (this *QWindow) setSizeIncrement(args ...interface{}) () {
  // setSizeIncrement(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow16setSizeIncrementERK5QSize
    // invoke: void setSizeIncrement(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow16setSizeIncrementERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setSizeIncrement", args)
  }

}

  // proto:  void QWindow::showMinimized();
func (this *QWindow) showMinimized(args ...interface{}) () {
  // showMinimized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow13showMinimizedEv
    // invoke: void showMinimized()
    C._ZN7QWindow13showMinimizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "showMinimized", args)
  }

}

  // proto:  QObject * QWindow::focusObject();
func (this *QWindow) focusObject(args ...interface{}) () {
  // focusObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow11focusObjectEv
    // invoke: QObject * focusObject()
    C._ZNK7QWindow11focusObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "focusObject", args)
  }

}

  // proto:  bool QWindow::isActive();
func (this *QWindow) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8isActiveEv
    // invoke: bool isActive()
    C._ZNK7QWindow8isActiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "isActive", args)
  }

}

  // proto:  QAccessibleInterface * QWindow::accessibleRoot();
func (this *QWindow) accessibleRoot(args ...interface{}) () {
  // accessibleRoot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow14accessibleRootEv
    // invoke: QAccessibleInterface * accessibleRoot()
    C._ZNK7QWindow14accessibleRootEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "accessibleRoot", args)
  }

}

  // proto:  QCursor QWindow::cursor();
func (this *QWindow) cursor(args ...interface{}) () {
  // cursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6cursorEv
    // invoke: QCursor cursor()
    C._ZNK7QWindow6cursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "cursor", args)
  }

}

  // proto:  void QWindow::setFormat(const QSurfaceFormat & format);
func (this *QWindow) setFormat(args ...interface{}) () {
  // setFormat(const class QSurfaceFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow9setFormatERK14QSurfaceFormat
    // invoke: void setFormat(const class QSurfaceFormat &)
    var arg0 = args[0].(QSurfaceFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow9setFormatERK14QSurfaceFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setFormat", args)
  }

}

  // proto:  void QWindow::showFullScreen();
func (this *QWindow) showFullScreen(args ...interface{}) () {
  // showFullScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow14showFullScreenEv
    // invoke: void showFullScreen()
    C._ZN7QWindow14showFullScreenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "showFullScreen", args)
  }

}

  // proto:  void QWindow::setX(int arg);
func (this *QWindow) setX(args ...interface{}) () {
  // setX(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow4setXEi
    // invoke: void setX(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow4setXEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setX", args)
  }

}

  // proto:  void QWindow::lower();
func (this *QWindow) lower(args ...interface{}) () {
  // lower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow5lowerEv
    // invoke: void lower()
    C._ZN7QWindow5lowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "lower", args)
  }

}

  // proto:  void QWindow::requestUpdate();
func (this *QWindow) requestUpdate(args ...interface{}) () {
  // requestUpdate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow13requestUpdateEv
    // invoke: void requestUpdate()
    C._ZN7QWindow13requestUpdateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "requestUpdate", args)
  }

}

  // proto:  void QWindow::hide();
func (this *QWindow) hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow4hideEv
    // invoke: void hide()
    C._ZN7QWindow4hideEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "hide", args)
  }

}

  // proto:  void QWindow::setMask(const QRegion & region);
func (this *QWindow) setMask(args ...interface{}) () {
  // setMask(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow7setMaskERK7QRegion
    // invoke: void setMask(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow7setMaskERK7QRegion(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMask", args)
  }

}

  // proto:  void QWindow::setMaximumSize(const QSize & size);
func (this *QWindow) setMaximumSize(args ...interface{}) () {
  // setMaximumSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow14setMaximumSizeERK5QSize
    // invoke: void setMaximumSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWindow14setMaximumSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMaximumSize", args)
  }

}

  // proto:  int QWindow::height();
func (this *QWindow) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow6heightEv
    // invoke: int height()
    C.demth_ZNK7QWindow6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "height", args)
  }

}

  // proto:  QSize QWindow::size();
func (this *QWindow) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow4sizeEv
    // invoke: QSize size()
    C.demth_ZNK7QWindow4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "size", args)
  }

}

  // proto:  int QWindow::maximumWidth();
func (this *QWindow) maximumWidth(args ...interface{}) () {
  // maximumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow12maximumWidthEv
    // invoke: int maximumWidth()
    C.demth_ZNK7QWindow12maximumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "maximumWidth", args)
  }

}

  // proto:  QPoint QWindow::position();
func (this *QWindow) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8positionEv
    // invoke: QPoint position()
    C.demth_ZNK7QWindow8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "position", args)
  }

}

  // proto:  void QWindow::setMaximumHeight(int h);
func (this *QWindow) setMaximumHeight(args ...interface{}) () {
  // setMaximumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow16setMaximumHeightEi
    // invoke: void setMaximumHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWindow16setMaximumHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWindow", "setMaximumHeight", args)
  }

}

  // proto:  QString QWindow::filePath();
func (this *QWindow) filePath(args ...interface{}) () {
  // filePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow8filePathEv
    // invoke: QString filePath()
    C._ZNK7QWindow8filePathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "filePath", args)
  }

}

  // proto:  void QWindow::showNormal();
func (this *QWindow) showNormal(args ...interface{}) () {
  // showNormal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWindow10showNormalEv
    // invoke: void showNormal()
    C._ZN7QWindow10showNormalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "showNormal", args)
  }

}

  // proto:  QPoint QWindow::framePosition();
func (this *QWindow) framePosition(args ...interface{}) () {
  // framePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow13framePositionEv
    // invoke: QPoint framePosition()
    C._ZNK7QWindow13framePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "framePosition", args)
  }

}

  // proto:  QIcon QWindow::icon();
func (this *QWindow) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWindow4iconEv
    // invoke: QIcon icon()
    C._ZNK7QWindow4iconEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindow", "icon", args)
  }

}

// <= body block end

