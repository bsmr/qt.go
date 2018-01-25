package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyleOptionMenuItem struct {
	*QStyleOption
}

func (this *QStyleOptionMenuItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionMenuItem) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionMenuItemFromPointer(cthis unsafe.Pointer) *QStyleOptionMenuItem {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionMenuItem{bcthis0}
}
func (*QStyleOptionMenuItem) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionMenuItem {
	return NewQStyleOptionMenuItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:372
// index:0
// Public
// void QStyleOptionMenuItem()
func NewQStyleOptionMenuItem() *QStyleOptionMenuItem {
	cthis := qtrt.Calloc(1, 256) // 136
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionMenuItemC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionMenuItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:376
// index:1
// Protected
// void QStyleOptionMenuItem(int)
func NewQStyleOptionMenuItem_1(version int) *QStyleOptionMenuItem {
	cthis := qtrt.Calloc(1, 256) // 136
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionMenuItemC2Ei", ffiqt.FFI_TYPE_VOID, cthis, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionMenuItemFromPointer(cthis)
	return gothis
}

type QStyleOptionMenuItem__StyleOptionType = int

const QStyleOptionMenuItem__Type QStyleOptionMenuItem__StyleOptionType = 4

type QStyleOptionMenuItem__StyleOptionVersion = int

const QStyleOptionMenuItem__Version QStyleOptionMenuItem__StyleOptionVersion = 1

type QStyleOptionMenuItem__MenuItemType = int

const QStyleOptionMenuItem__Normal QStyleOptionMenuItem__MenuItemType = 0
const QStyleOptionMenuItem__DefaultItem QStyleOptionMenuItem__MenuItemType = 1
const QStyleOptionMenuItem__Separator QStyleOptionMenuItem__MenuItemType = 2
const QStyleOptionMenuItem__SubMenu QStyleOptionMenuItem__MenuItemType = 3
const QStyleOptionMenuItem__Scroller QStyleOptionMenuItem__MenuItemType = 4
const QStyleOptionMenuItem__TearOff QStyleOptionMenuItem__MenuItemType = 5
const QStyleOptionMenuItem__Margin QStyleOptionMenuItem__MenuItemType = 6
const QStyleOptionMenuItem__EmptyArea QStyleOptionMenuItem__MenuItemType = 7

type QStyleOptionMenuItem__CheckType = int

const QStyleOptionMenuItem__NotCheckable QStyleOptionMenuItem__CheckType = 0
const QStyleOptionMenuItem__Exclusive QStyleOptionMenuItem__CheckType = 1
const QStyleOptionMenuItem__NonExclusive QStyleOptionMenuItem__CheckType = 2

//  body block end
