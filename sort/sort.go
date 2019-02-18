package sort

import (
	"sort"
)

// UintSlice attaches the methods of Interface to []uint, sorting in increasing order.
type UintSlice []uint

// Len implements sort.Interface.
func (u UintSlice) Len() int {
	return len(u)
}

// Less implements sort.Interface.
func (u UintSlice) Less(i, j int) bool {
	return u[i] < u[j]
}

// Swap implements sort.Interface.
func (u UintSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Sort is a convenience method.
// Same as sort.Sort(u).
func (u UintSlice) Sort() {
	sort.Sort(u)
}

// Uints sorts a slice of uints in increasing order.
func Uints(v []uint) {
	sort.Sort(UintSlice(v))
}

// Uint64Slice attaches the methods of Interface to []uint64, sorting in increasing order.
type Uint64Slice []uint64

// Len implements sort.Interface.
func (u Uint64Slice) Len() int {
	return len(u)
}

// Less implements sort.Interface.
func (u Uint64Slice) Less(i, j int) bool {
	return u[i] < u[j]
}

// Swap implements sort.Interface.
func (u Uint64Slice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Sort is a convenience method.
// Same as sort.Sort(u).
func (u Uint64Slice) Sort() {
	sort.Sort(u)
}

// Uint64s sorts a slice of uint64s in increasing order.
func Uint64s(v []uint64) {
	sort.Sort(Uint64Slice(v))
}

// Uint32Slice attaches the methods of Interface to []uint32, sorting in increasing order.
type Uint32Slice []uint32

// Len implements sort.Interface.
func (u Uint32Slice) Len() int {
	return len(u)
}

// Less implements sort.Interface.
func (u Uint32Slice) Less(i, j int) bool {
	return u[i] < u[j]
}

// Swap implements sort.Interface.
func (u Uint32Slice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Sort is a convenience method.
// Same as sort.Sort(u).
func (u Uint32Slice) Sort() {
	sort.Sort(u)
}

// Uint32s sorts a slice of uint32s in increasing order.
func Uint32s(u []uint32) {
	sort.Sort(Uint32Slice(u))
}

// Uint16Slice attaches the methods of Interface to []uint16, sorting in increasing order.
type Uint16Slice []uint16

// Len implements sort.Interface.
func (u Uint16Slice) Len() int {
	return len(u)
}

// Less implements sort.Interface.
func (u Uint16Slice) Less(i, j int) bool {
	return u[i] < u[j]
}

// Swap implements sort.Interface.
func (u Uint16Slice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Sort is a convenience method.
// Same as sort.Sort(u).
func (u Uint16Slice) Sort() {
	sort.Sort(u)
}

// Uint16s sorts a slice of uint16s in increasing order.
func Uint16s(u []uint16) {
	sort.Sort(Uint16Slice(u))
}

// Uint8Slice attaches the methods of Interface to []uint8, sorting in increasing order.
type Uint8Slice []uint8

// Len implements sort.Interface.
func (u Uint8Slice) Len() int {
	return len(u)
}

// Less implements sort.Interface.
func (u Uint8Slice) Less(i, j int) bool {
	return u[i] < u[j]
}

// Swap implements sort.Interface.
func (u Uint8Slice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Sort is a convenience method.
// Same as sort.Sort(u).
func (u Uint8Slice) Sort() {
	sort.Sort(u)
}

// Uint8s sorts a slice of uint8s in increasing order.
func Uint8s(v []uint8) {
	sort.Sort(Uint8Slice(v))
}

// Int64Slice attaches the methods of Interface to []int64, sorting in increasing order.
type Int64Slice []int64

// Len implements sort.Interface.
func (i Int64Slice) Len() int {
	return len(i)
}

// Less implements sort.Interface.
func (i Int64Slice) Less(x, y int) bool {
	return i[x] < i[y]
}

// Swap implements sort.Interface.
func (i Int64Slice) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

// Sort is a convenience method.
func (i Int64Slice) Sort() {
	sort.Sort(i)
}

// Int64s sorts a slice of int64s in increasing order.
func Int64s(i []int64) {
	sort.Sort(Int64Slice(i))
}

// Int32Slice attaches the methods of Interface to []int32, sorting in increasing order.
type Int32Slice []int32

// Len implements sort.Interface.
func (i Int32Slice) Len() int {
	return len(i)
}

// Less implements sort.Interface.
func (i Int32Slice) Less(x, y int) bool {
	return i[x] < i[y]
}

// Swap implements sort.Interface.
func (i Int32Slice) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

// Sort is a convenience method.
// Same as sort.Sort(i).
func (i Int32Slice) Sort() {
	sort.Sort(i)
}

// Int32s sorts a slice of int32s in increasing order.
func Int32s(i []int32) {
	sort.Sort(Int32Slice(i))
}

// Int16Slice attaches the methods of Interface to []int16, sorting in increasing order.
type Int16Slice []int16

// Len implements sort.Interface.
func (i Int16Slice) Len() int {
	return len(i)
}

// Less implements sort.Interface.
func (i Int16Slice) Less(x, y int) bool {
	return i[x] < i[y]
}

// Swap implements sort.Interface.
func (i Int16Slice) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

// Sort is a convenience method.
// Same as sort.Sort(i).
func (i Int16Slice) Sort() {
	sort.Sort(i)
}

// Int16s sorts a slice of int16s in increasing order.
func Int16s(i []int16) {
	sort.Sort(Int16Slice(i))
}

// Int8Slice attaches the methods of Interface to []int8, sorting in increasing order.
type Int8Slice []int8

// Len implements sort.Interface.
func (i Int8Slice) Len() int {
	return len(i)
}

// Less implements sort.Interface.
func (i Int8Slice) Less(x, y int) bool {
	return i[x] < i[y]
}

// Swap implements sort.Interface.
func (i Int8Slice) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

// Sort is a convenience method.
// Same as sort.Sort(i).
func (i Int8Slice) Sort() {
	sort.Sort(i)
}

// Int8s sorts a slice of int8s in increasing order.
func Int8s(v []int8) {
	sort.Sort(Int8Slice(v))
}

// Float32Slice attaches the methods of Interface to []float32, sorting in increasing order.
type Float32Slice []float32

// Len implements sort.Interface.
func (f Float32Slice) Len() int {
	return len(f)
}

// Less implements sort.Interface.
func (f Float32Slice) Less(i, j int) bool {
	return f[i] < f[j]
}

// Swap implements sort.Interface.
func (f Float32Slice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Sort is a convenience method.
// Same as sort.Sort(f).
func (f Float32Slice) Sort() {
	sort.Sort(f)
}

// Float32s sorts a slice of float32s in increasing order.
func Float32s(v []float32) {
	sort.Sort(Float32Slice(v))
}
