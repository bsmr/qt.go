package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qregexp.h
// dst-file: /src/core/qregexp.go
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
  // proto:  void QRegExp::QRegExp(const QRegExp & rx);
extern void* dector_ZN7QRegExpC1ERKS_(void* arg0);
extern void _ZN7QRegExpC1ERKS_(void* qthis, void* arg0);
  // proto:  QStringList QRegExp::capturedTexts();
extern void _ZN7QRegExp13capturedTextsEv(void* qthis);
  // proto:  int QRegExp::captureCount();
extern void _ZNK7QRegExp12captureCountEv(void* qthis);
  // proto: static QString QRegExp::escape(const QString & str);
extern void _ZN7QRegExp6escapeERK7QString(void* arg0);
  // proto:  bool QRegExp::isEmpty();
extern void _ZNK7QRegExp7isEmptyEv(void* qthis);
  // proto:  bool QRegExp::isMinimal();
extern void _ZNK7QRegExp9isMinimalEv(void* qthis);
  // proto:  int QRegExp::matchedLength();
extern void _ZNK7QRegExp13matchedLengthEv(void* qthis);
  // proto:  QString QRegExp::pattern();
extern void _ZNK7QRegExp7patternEv(void* qthis);
  // proto:  void QRegExp::setPattern(const QString & pattern);
extern void _ZN7QRegExp10setPatternERK7QString(void* qthis, void* arg0);
  // proto:  bool QRegExp::isValid();
extern void _ZNK7QRegExp7isValidEv(void* qthis);
  // proto:  void QRegExp::~QRegExp();
extern void _ZN7QRegExpD0Ev(void* qthis);
  // proto:  bool QRegExp::exactMatch(const QString & str);
extern void _ZNK7QRegExp10exactMatchERK7QString(void* qthis, void* arg0);
  // proto:  void QRegExp::swap(QRegExp & other);
extern void demth_ZN7QRegExp4swapERS_(void* qthis, void* arg0);
  // proto:  int QRegExp::pos(int nth);
extern void _ZN7QRegExp3posEi(void* qthis, int32_t arg0);
  // proto:  void QRegExp::QRegExp();
extern void* dector_ZN7QRegExpC1Ev();
extern void _ZN7QRegExpC1Ev(void* qthis);
  // proto:  QString QRegExp::cap(int nth);
extern void _ZN7QRegExp3capEi(void* qthis, int32_t arg0);
  // proto:  QString QRegExp::errorString();
extern void _ZN7QRegExp11errorStringEv(void* qthis);
  // proto:  void QRegExp::setMinimal(bool minimal);
extern void _ZN7QRegExp10setMinimalEb(void* qthis, bool arg0);
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

// class sizeof(QRegExp)=8
type QRegExp struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QRegExp::QRegExp(const QRegExp & rx);
func NewQRegExp(args ...interface{}) QRegExp {
  return QRegExp{}
}

  // proto:  QStringList QRegExp::capturedTexts();
func (this *QRegExp) capturedTexts(args ...interface{}) () {
  // capturedTexts()
  // capturedTexts()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp13capturedTextsEv
    // invoke: QStringList capturedTexts()
    C._ZN7QRegExp13capturedTextsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "capturedTexts", args)
  }

}

  // proto:  int QRegExp::captureCount();
func (this *QRegExp) captureCount(args ...interface{}) () {
  // captureCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp12captureCountEv
    // invoke: int captureCount()
    C._ZNK7QRegExp12captureCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "captureCount", args)
  }

}

  // proto: static QString QRegExp::escape(const QString & str);
func (this *QRegExp) escape_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegExp", "escape", args)
  }

}

  // proto:  bool QRegExp::isEmpty();
func (this *QRegExp) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK7QRegExp7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "isEmpty", args)
  }

}

  // proto:  bool QRegExp::isMinimal();
func (this *QRegExp) isMinimal(args ...interface{}) () {
  // isMinimal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp9isMinimalEv
    // invoke: bool isMinimal()
    C._ZNK7QRegExp9isMinimalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "isMinimal", args)
  }

}

  // proto:  int QRegExp::matchedLength();
func (this *QRegExp) matchedLength(args ...interface{}) () {
  // matchedLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp13matchedLengthEv
    // invoke: int matchedLength()
    C._ZNK7QRegExp13matchedLengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "matchedLength", args)
  }

}

  // proto:  QString QRegExp::pattern();
func (this *QRegExp) pattern(args ...interface{}) () {
  // pattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp7patternEv
    // invoke: QString pattern()
    C._ZNK7QRegExp7patternEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "pattern", args)
  }

}

  // proto:  void QRegExp::setPattern(const QString & pattern);
func (this *QRegExp) setPattern(args ...interface{}) () {
  // setPattern(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp10setPatternERK7QString
    // invoke: void setPattern(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QRegExp10setPatternERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExp", "setPattern", args)
  }

}

  // proto:  bool QRegExp::isValid();
func (this *QRegExp) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp7isValidEv
    // invoke: bool isValid()
    C._ZNK7QRegExp7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "isValid", args)
  }

}

  // proto:  void QRegExp::~QRegExp();
func (this *QRegExp) FreeQRegExp(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegExp", "~QRegExp", args)
  }

}

  // proto:  bool QRegExp::exactMatch(const QString & str);
func (this *QRegExp) exactMatch(args ...interface{}) () {
  // exactMatch(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp10exactMatchERK7QString
    // invoke: bool exactMatch(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QRegExp10exactMatchERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExp", "exactMatch", args)
  }

}

  // proto:  void QRegExp::swap(QRegExp & other);
func (this *QRegExp) swap(args ...interface{}) () {
  // swap(class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp4swapERS_
    // invoke: void swap(class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QRegExp4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExp", "swap", args)
  }

}

  // proto:  int QRegExp::pos(int nth);
func (this *QRegExp) pos(args ...interface{}) () {
  // pos(int)
  // pos(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp3posEi
    // invoke: int pos(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QRegExp3posEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExp", "pos", args)
  }

}

  // proto:  QString QRegExp::cap(int nth);
func (this *QRegExp) cap(args ...interface{}) () {
  // cap(int)
  // cap(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp3capEi
    // invoke: QString cap(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QRegExp3capEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExp", "cap", args)
  }

}

  // proto:  QString QRegExp::errorString();
func (this *QRegExp) errorString(args ...interface{}) () {
  // errorString()
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp11errorStringEv
    // invoke: QString errorString()
    C._ZN7QRegExp11errorStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRegExp", "errorString", args)
  }

}

  // proto:  void QRegExp::setMinimal(bool minimal);
func (this *QRegExp) setMinimal(args ...interface{}) () {
  // setMinimal(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp10setMinimalEb
    // invoke: void setMinimal(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN7QRegExp10setMinimalEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRegExp", "setMinimal", args)
  }

}

// <= body block end

