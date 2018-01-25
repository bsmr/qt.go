package qtgui

// /usr/include/qt/QtGui/qfontdatabase.h
// #include <qfontdatabase.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QFontDatabase struct {
	*qtrt.CObject
}

func (this *QFontDatabase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFontDatabase) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQFontDatabaseFromPointer(cthis unsafe.Pointer) *QFontDatabase {
	return &QFontDatabase{&qtrt.CObject{cthis}}
}
func (*QFontDatabase) NewFromPointer(cthis unsafe.Pointer) *QFontDatabase {
	return NewQFontDatabaseFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfontdatabase.h:118
// index:0
// Public
// void QFontDatabase()
func NewQFontDatabase() *QFontDatabase {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabaseC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontDatabaseFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfontdatabase.h:127
// index:0
// Public
// QString styleString(const class QFont &)
func (this *QFontDatabase) StyleString(font *QFont) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase11styleStringERK5QFont", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qfontdatabase.h:128
// index:1
// Public
// QString styleString(const class QFontInfo &)
func (this *QFontDatabase) StyleString_1(fontInfo *QFontInfo) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = fontInfo.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase11styleStringERK9QFontInfo", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qfontdatabase.h:130
// index:0
// Public
// QFont font(const class QString &, const class QString &, int)
func (this *QFontDatabase) Font(family *qtcore.QString, style *qtcore.QString, pointSize int) *QFont /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase4fontERK7QStringS2_i", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1, pointSize)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qfontdatabase.h:132
// index:0
// Public
// bool isBitmapScalable(const class QString &, const class QString &)
func (this *QFontDatabase) IsBitmapScalable(family *qtcore.QString, style *qtcore.QString) bool {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:133
// index:0
// Public
// bool isSmoothlyScalable(const class QString &, const class QString &)
func (this *QFontDatabase) IsSmoothlyScalable(family *qtcore.QString, style *qtcore.QString) bool {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:134
// index:0
// Public
// bool isScalable(const class QString &, const class QString &)
func (this *QFontDatabase) IsScalable(family *qtcore.QString, style *qtcore.QString) bool {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase10isScalableERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:135
// index:0
// Public
// bool isFixedPitch(const class QString &, const class QString &)
func (this *QFontDatabase) IsFixedPitch(family *qtcore.QString, style *qtcore.QString) bool {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase12isFixedPitchERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:137
// index:0
// Public
// bool italic(const class QString &, const class QString &)
func (this *QFontDatabase) Italic(family *qtcore.QString, style *qtcore.QString) bool {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase6italicERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:138
// index:0
// Public
// bool bold(const class QString &, const class QString &)
func (this *QFontDatabase) Bold(family *qtcore.QString, style *qtcore.QString) bool {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase4boldERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:139
// index:0
// Public
// int weight(const class QString &, const class QString &)
func (this *QFontDatabase) Weight(family *qtcore.QString, style *qtcore.QString) int {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase6weightERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qfontdatabase.h:141
// index:0
// Public
// bool hasFamily(const class QString &)
func (this *QFontDatabase) HasFamily(family *qtcore.QString) bool {
	var convArg0 = family.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase9hasFamilyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:142
// index:0
// Public
// bool isPrivateFamily(const class QString &)
func (this *QFontDatabase) IsPrivateFamily(family *qtcore.QString) bool {
	var convArg0 = family.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase15isPrivateFamilyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:144
// index:0
// Public static
// QString writingSystemName(enum QFontDatabase::WritingSystem)
func (this *QFontDatabase) WritingSystemName(writingSystem int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase17writingSystemNameENS_13WritingSystemE", ffiqt.FFI_TYPE_POINTER, writingSystem)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QFontDatabase_WritingSystemName(writingSystem int) *qtcore.QString /*123*/ {
	var nilthis *QFontDatabase
	rv := nilthis.WritingSystemName(writingSystem)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:145
// index:0
// Public static
// QString writingSystemSample(enum QFontDatabase::WritingSystem)
func (this *QFontDatabase) WritingSystemSample(writingSystem int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase19writingSystemSampleENS_13WritingSystemE", ffiqt.FFI_TYPE_POINTER, writingSystem)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QFontDatabase_WritingSystemSample(writingSystem int) *qtcore.QString /*123*/ {
	var nilthis *QFontDatabase
	rv := nilthis.WritingSystemSample(writingSystem)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:147
// index:0
// Public static
// int addApplicationFont(const class QString &)
func (this *QFontDatabase) AddApplicationFont(fileName *qtcore.QString) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase18addApplicationFontERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QFontDatabase_AddApplicationFont(fileName *qtcore.QString) int {
	var nilthis *QFontDatabase
	rv := nilthis.AddApplicationFont(fileName)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:148
// index:0
// Public static
// int addApplicationFontFromData(const class QByteArray &)
func (this *QFontDatabase) AddApplicationFontFromData(fontData *qtcore.QByteArray) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, fontData)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QFontDatabase_AddApplicationFontFromData(fontData *qtcore.QByteArray) int {
	var nilthis *QFontDatabase
	rv := nilthis.AddApplicationFontFromData(fontData)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:150
// index:0
// Public static
// bool removeApplicationFont(int)
func (this *QFontDatabase) RemoveApplicationFont(id int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase21removeApplicationFontEi", ffiqt.FFI_TYPE_POINTER, id)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QFontDatabase_RemoveApplicationFont(id int) bool {
	var nilthis *QFontDatabase
	rv := nilthis.RemoveApplicationFont(id)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:151
// index:0
// Public static
// bool removeAllApplicationFonts()
func (this *QFontDatabase) RemoveAllApplicationFonts() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase25removeAllApplicationFontsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QFontDatabase_RemoveAllApplicationFonts() bool {
	var nilthis *QFontDatabase
	rv := nilthis.RemoveAllApplicationFonts()
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:154
// index:0
// Public static
// bool supportsThreadedFontRendering()
func (this *QFontDatabase) SupportsThreadedFontRendering() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase29supportsThreadedFontRenderingEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QFontDatabase_SupportsThreadedFontRendering() bool {
	var nilthis *QFontDatabase
	rv := nilthis.SupportsThreadedFontRendering()
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:157
// index:0
// Public static
// QFont systemFont(enum QFontDatabase::SystemFont)
func (this *QFontDatabase) SystemFont(type_ int) *QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase10systemFontENS_10SystemFontE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QFontDatabase_SystemFont(type_ int) *QFont /*123*/ {
	var nilthis *QFontDatabase
	rv := nilthis.SystemFont(type_)
	return rv
}

type QFontDatabase__WritingSystem = int

const QFontDatabase__Any QFontDatabase__WritingSystem = 0
const QFontDatabase__Latin QFontDatabase__WritingSystem = 1
const QFontDatabase__Greek QFontDatabase__WritingSystem = 2
const QFontDatabase__Cyrillic QFontDatabase__WritingSystem = 3
const QFontDatabase__Armenian QFontDatabase__WritingSystem = 4
const QFontDatabase__Hebrew QFontDatabase__WritingSystem = 5
const QFontDatabase__Arabic QFontDatabase__WritingSystem = 6
const QFontDatabase__Syriac QFontDatabase__WritingSystem = 7
const QFontDatabase__Thaana QFontDatabase__WritingSystem = 8
const QFontDatabase__Devanagari QFontDatabase__WritingSystem = 9
const QFontDatabase__Bengali QFontDatabase__WritingSystem = 10
const QFontDatabase__Gurmukhi QFontDatabase__WritingSystem = 11
const QFontDatabase__Gujarati QFontDatabase__WritingSystem = 12
const QFontDatabase__Oriya QFontDatabase__WritingSystem = 13
const QFontDatabase__Tamil QFontDatabase__WritingSystem = 14
const QFontDatabase__Telugu QFontDatabase__WritingSystem = 15
const QFontDatabase__Kannada QFontDatabase__WritingSystem = 16
const QFontDatabase__Malayalam QFontDatabase__WritingSystem = 17
const QFontDatabase__Sinhala QFontDatabase__WritingSystem = 18
const QFontDatabase__Thai QFontDatabase__WritingSystem = 19
const QFontDatabase__Lao QFontDatabase__WritingSystem = 20
const QFontDatabase__Tibetan QFontDatabase__WritingSystem = 21
const QFontDatabase__Myanmar QFontDatabase__WritingSystem = 22
const QFontDatabase__Georgian QFontDatabase__WritingSystem = 23
const QFontDatabase__Khmer QFontDatabase__WritingSystem = 24
const QFontDatabase__SimplifiedChinese QFontDatabase__WritingSystem = 25
const QFontDatabase__TraditionalChinese QFontDatabase__WritingSystem = 26
const QFontDatabase__Japanese QFontDatabase__WritingSystem = 27
const QFontDatabase__Korean QFontDatabase__WritingSystem = 28
const QFontDatabase__Vietnamese QFontDatabase__WritingSystem = 29
const QFontDatabase__Symbol QFontDatabase__WritingSystem = 30
const QFontDatabase__Other QFontDatabase__WritingSystem = 30
const QFontDatabase__Ogham QFontDatabase__WritingSystem = 31
const QFontDatabase__Runic QFontDatabase__WritingSystem = 32
const QFontDatabase__Nko QFontDatabase__WritingSystem = 33
const QFontDatabase__WritingSystemsCount QFontDatabase__WritingSystem = 34

type QFontDatabase__SystemFont = int

const QFontDatabase__GeneralFont QFontDatabase__SystemFont = 0
const QFontDatabase__FixedFont QFontDatabase__SystemFont = 1
const QFontDatabase__TitleFont QFontDatabase__SystemFont = 2
const QFontDatabase__SmallestReadableFont QFontDatabase__SystemFont = 3

//  body block end
