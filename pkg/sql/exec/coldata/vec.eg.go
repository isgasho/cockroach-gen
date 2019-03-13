// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package coldata

import (
	"fmt"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
)

func (m *memColumn) Append(vec Vec, colType types.T, toLength uint64, fromLength uint16) {
	switch colType {
	case types.Bool:
		m.col = append(m.Bool()[:toLength], vec.Bool()[:fromLength]...)
	case types.Bytes:
		m.col = append(m.Bytes()[:toLength], vec.Bytes()[:fromLength]...)
	case types.Decimal:
		m.col = append(m.Decimal()[:toLength], vec.Decimal()[:fromLength]...)
	case types.Int8:
		m.col = append(m.Int8()[:toLength], vec.Int8()[:fromLength]...)
	case types.Int16:
		m.col = append(m.Int16()[:toLength], vec.Int16()[:fromLength]...)
	case types.Int32:
		m.col = append(m.Int32()[:toLength], vec.Int32()[:fromLength]...)
	case types.Int64:
		m.col = append(m.Int64()[:toLength], vec.Int64()[:fromLength]...)
	case types.Float32:
		m.col = append(m.Float32()[:toLength], vec.Float32()[:fromLength]...)
	case types.Float64:
		m.col = append(m.Float64()[:toLength], vec.Float64()[:fromLength]...)
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}

	if fromLength > 0 {
		m.nulls = append(m.nulls, make([]int64, (fromLength-1)>>6+1)...)

		if vec.HasNulls() {
			for i := uint16(0); i < fromLength; i++ {
				if vec.NullAt(i) {
					m.SetNull64(toLength + uint64(i))
				}
			}
		}
	}
}

func (m *memColumn) AppendSlice(
	vec Vec, colType types.T, destStartIdx uint64, srcStartIdx uint16, srcEndIdx uint16,
) {
	batchSize := srcEndIdx - srcStartIdx
	outputLen := destStartIdx + uint64(batchSize)

	switch colType {
	case types.Bool:
		if outputLen > uint64(len(m.Bool())) {
			m.col = append(m.Bool()[:destStartIdx], vec.Bool()[srcStartIdx:srcEndIdx]...)
		} else {
			copy(m.Bool()[destStartIdx:], vec.Bool()[srcStartIdx:srcEndIdx])
		}
	case types.Bytes:
		if outputLen > uint64(len(m.Bytes())) {
			m.col = append(m.Bytes()[:destStartIdx], vec.Bytes()[srcStartIdx:srcEndIdx]...)
		} else {
			copy(m.Bytes()[destStartIdx:], vec.Bytes()[srcStartIdx:srcEndIdx])
		}
	case types.Decimal:
		if outputLen > uint64(len(m.Decimal())) {
			m.col = append(m.Decimal()[:destStartIdx], vec.Decimal()[srcStartIdx:srcEndIdx]...)
		} else {
			copy(m.Decimal()[destStartIdx:], vec.Decimal()[srcStartIdx:srcEndIdx])
		}
	case types.Int8:
		if outputLen > uint64(len(m.Int8())) {
			m.col = append(m.Int8()[:destStartIdx], vec.Int8()[srcStartIdx:srcEndIdx]...)
		} else {
			copy(m.Int8()[destStartIdx:], vec.Int8()[srcStartIdx:srcEndIdx])
		}
	case types.Int16:
		if outputLen > uint64(len(m.Int16())) {
			m.col = append(m.Int16()[:destStartIdx], vec.Int16()[srcStartIdx:srcEndIdx]...)
		} else {
			copy(m.Int16()[destStartIdx:], vec.Int16()[srcStartIdx:srcEndIdx])
		}
	case types.Int32:
		if outputLen > uint64(len(m.Int32())) {
			m.col = append(m.Int32()[:destStartIdx], vec.Int32()[srcStartIdx:srcEndIdx]...)
		} else {
			copy(m.Int32()[destStartIdx:], vec.Int32()[srcStartIdx:srcEndIdx])
		}
	case types.Int64:
		if outputLen > uint64(len(m.Int64())) {
			m.col = append(m.Int64()[:destStartIdx], vec.Int64()[srcStartIdx:srcEndIdx]...)
		} else {
			copy(m.Int64()[destStartIdx:], vec.Int64()[srcStartIdx:srcEndIdx])
		}
	case types.Float32:
		if outputLen > uint64(len(m.Float32())) {
			m.col = append(m.Float32()[:destStartIdx], vec.Float32()[srcStartIdx:srcEndIdx]...)
		} else {
			copy(m.Float32()[destStartIdx:], vec.Float32()[srcStartIdx:srcEndIdx])
		}
	case types.Float64:
		if outputLen > uint64(len(m.Float64())) {
			m.col = append(m.Float64()[:destStartIdx], vec.Float64()[srcStartIdx:srcEndIdx]...)
		} else {
			copy(m.Float64()[destStartIdx:], vec.Float64()[srcStartIdx:srcEndIdx])
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}

	m.ExtendNulls(vec, destStartIdx, srcStartIdx, batchSize)
}

func (m *memColumn) AppendWithSel(
	vec Vec, sel []uint16, batchSize uint16, colType types.T, toLength uint64,
) {
	switch colType {
	case types.Bool:
		toCol := append(m.Bool()[:toLength], make([]bool, batchSize)...)
		fromCol := vec.Bool()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Bytes:
		toCol := append(m.Bytes()[:toLength], make([][]byte, batchSize)...)
		fromCol := vec.Bytes()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Decimal:
		toCol := append(m.Decimal()[:toLength], make([]apd.Decimal, batchSize)...)
		fromCol := vec.Decimal()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Int8:
		toCol := append(m.Int8()[:toLength], make([]int8, batchSize)...)
		fromCol := vec.Int8()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Int16:
		toCol := append(m.Int16()[:toLength], make([]int16, batchSize)...)
		fromCol := vec.Int16()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Int32:
		toCol := append(m.Int32()[:toLength], make([]int32, batchSize)...)
		fromCol := vec.Int32()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Int64:
		toCol := append(m.Int64()[:toLength], make([]int64, batchSize)...)
		fromCol := vec.Int64()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Float32:
		toCol := append(m.Float32()[:toLength], make([]float32, batchSize)...)
		fromCol := vec.Float32()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Float64:
		toCol := append(m.Float64()[:toLength], make([]float64, batchSize)...)
		fromCol := vec.Float64()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}

	if batchSize > 0 {
		m.nulls = append(m.nulls, make([]int64, (batchSize-1)>>6+1)...)
		for i := uint16(0); i < batchSize; i++ {
			if vec.NullAt(sel[i]) {
				m.SetNull64(toLength + uint64(i))
			}
		}
	}
}

func (m *memColumn) AppendSliceWithSel(
	vec Vec, colType types.T, destStartIdx uint64, srcStartIdx uint16, srcEndIdx uint16, sel []uint16,
) {
	batchSize := srcEndIdx - srcStartIdx
	switch colType {
	case types.Bool:
		toCol := append(m.Bool()[:destStartIdx], make([]bool, batchSize)...)
		fromCol := vec.Bool()

		for i := 0; i < int(batchSize); i++ {
			toCol[uint64(i)+destStartIdx] = fromCol[sel[i+int(srcStartIdx)]]
		}

		m.col = toCol
	case types.Bytes:
		toCol := append(m.Bytes()[:destStartIdx], make([][]byte, batchSize)...)
		fromCol := vec.Bytes()

		for i := 0; i < int(batchSize); i++ {
			toCol[uint64(i)+destStartIdx] = fromCol[sel[i+int(srcStartIdx)]]
		}

		m.col = toCol
	case types.Decimal:
		toCol := append(m.Decimal()[:destStartIdx], make([]apd.Decimal, batchSize)...)
		fromCol := vec.Decimal()

		for i := 0; i < int(batchSize); i++ {
			toCol[uint64(i)+destStartIdx] = fromCol[sel[i+int(srcStartIdx)]]
		}

		m.col = toCol
	case types.Int8:
		toCol := append(m.Int8()[:destStartIdx], make([]int8, batchSize)...)
		fromCol := vec.Int8()

		for i := 0; i < int(batchSize); i++ {
			toCol[uint64(i)+destStartIdx] = fromCol[sel[i+int(srcStartIdx)]]
		}

		m.col = toCol
	case types.Int16:
		toCol := append(m.Int16()[:destStartIdx], make([]int16, batchSize)...)
		fromCol := vec.Int16()

		for i := 0; i < int(batchSize); i++ {
			toCol[uint64(i)+destStartIdx] = fromCol[sel[i+int(srcStartIdx)]]
		}

		m.col = toCol
	case types.Int32:
		toCol := append(m.Int32()[:destStartIdx], make([]int32, batchSize)...)
		fromCol := vec.Int32()

		for i := 0; i < int(batchSize); i++ {
			toCol[uint64(i)+destStartIdx] = fromCol[sel[i+int(srcStartIdx)]]
		}

		m.col = toCol
	case types.Int64:
		toCol := append(m.Int64()[:destStartIdx], make([]int64, batchSize)...)
		fromCol := vec.Int64()

		for i := 0; i < int(batchSize); i++ {
			toCol[uint64(i)+destStartIdx] = fromCol[sel[i+int(srcStartIdx)]]
		}

		m.col = toCol
	case types.Float32:
		toCol := append(m.Float32()[:destStartIdx], make([]float32, batchSize)...)
		fromCol := vec.Float32()

		for i := 0; i < int(batchSize); i++ {
			toCol[uint64(i)+destStartIdx] = fromCol[sel[i+int(srcStartIdx)]]
		}

		m.col = toCol
	case types.Float64:
		toCol := append(m.Float64()[:destStartIdx], make([]float64, batchSize)...)
		fromCol := vec.Float64()

		for i := 0; i < int(batchSize); i++ {
			toCol[uint64(i)+destStartIdx] = fromCol[sel[i+int(srcStartIdx)]]
		}

		m.col = toCol
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}

	m.ExtendNullsWithSel(vec, destStartIdx, srcStartIdx, batchSize, sel)
}

func (m *memColumn) Copy(src Vec, srcStartIdx, srcEndIdx uint64, typ types.T) {
	m.CopyAt(src, 0, srcStartIdx, srcEndIdx, typ)
}

func (m *memColumn) CopyAt(src Vec, destStartIdx, srcStartIdx, srcEndIdx uint64, typ types.T) {
	switch typ {
	case types.Bool:
		copy(m.Bool()[destStartIdx:], src.Bool()[srcStartIdx:srcEndIdx])
	case types.Bytes:
		copy(m.Bytes()[destStartIdx:], src.Bytes()[srcStartIdx:srcEndIdx])
	case types.Decimal:
		copy(m.Decimal()[destStartIdx:], src.Decimal()[srcStartIdx:srcEndIdx])
	case types.Int8:
		copy(m.Int8()[destStartIdx:], src.Int8()[srcStartIdx:srcEndIdx])
	case types.Int16:
		copy(m.Int16()[destStartIdx:], src.Int16()[srcStartIdx:srcEndIdx])
	case types.Int32:
		copy(m.Int32()[destStartIdx:], src.Int32()[srcStartIdx:srcEndIdx])
	case types.Int64:
		copy(m.Int64()[destStartIdx:], src.Int64()[srcStartIdx:srcEndIdx])
	case types.Float32:
		copy(m.Float32()[destStartIdx:], src.Float32()[srcStartIdx:srcEndIdx])
	case types.Float64:
		copy(m.Float64()[destStartIdx:], src.Float64()[srcStartIdx:srcEndIdx])
	default:
		panic(fmt.Sprintf("unhandled type %d", typ))
	}
}

func (m *memColumn) CopyWithSelInt64(vec Vec, sel []uint64, nSel uint16, colType types.T) {
	m.UnsetNulls()

	// todo (changangela): handle the case when nSel > BatchSize
	switch colType {
	case types.Bool:
		toCol := m.Bool()
		fromCol := vec.Bool()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Bytes:
		toCol := m.Bytes()
		fromCol := vec.Bytes()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Decimal:
		toCol := m.Decimal()
		fromCol := vec.Decimal()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int8:
		toCol := m.Int8()
		fromCol := vec.Int8()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int16:
		toCol := m.Int16()
		fromCol := vec.Int16()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int32:
		toCol := m.Int32()
		fromCol := vec.Int32()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int64:
		toCol := m.Int64()
		fromCol := vec.Int64()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Float32:
		toCol := m.Float32()
		fromCol := vec.Float32()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Float64:
		toCol := m.Float64()
		fromCol := vec.Float64()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) CopyWithSelInt16(vec Vec, sel []uint16, nSel uint16, colType types.T) {
	m.UnsetNulls()

	switch colType {
	case types.Bool:
		toCol := m.Bool()
		fromCol := vec.Bool()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Bytes:
		toCol := m.Bytes()
		fromCol := vec.Bytes()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Decimal:
		toCol := m.Decimal()
		fromCol := vec.Decimal()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int8:
		toCol := m.Int8()
		fromCol := vec.Int8()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int16:
		toCol := m.Int16()
		fromCol := vec.Int16()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int32:
		toCol := m.Int32()
		fromCol := vec.Int32()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int64:
		toCol := m.Int64()
		fromCol := vec.Int64()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Float32:
		toCol := m.Float32()
		fromCol := vec.Float32()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Float64:
		toCol := m.Float64()
		fromCol := vec.Float64()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) CopyWithSelAndNilsInt64(
	vec Vec, sel []uint64, nSel uint16, nils []bool, colType types.T,
) {
	m.UnsetNulls()

	switch colType {
	case types.Bool:
		toCol := m.Bool()
		fromCol := vec.Bool()

		if vec.HasNulls() {
			// TODO(jordan): copy the null arrays in batch.
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					if vec.NullAt64(sel[i]) {
						m.SetNull(i)
					} else {
						toCol[i] = fromCol[sel[i]]
					}
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		}
	case types.Bytes:
		toCol := m.Bytes()
		fromCol := vec.Bytes()

		if vec.HasNulls() {
			// TODO(jordan): copy the null arrays in batch.
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					if vec.NullAt64(sel[i]) {
						m.SetNull(i)
					} else {
						toCol[i] = fromCol[sel[i]]
					}
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		}
	case types.Decimal:
		toCol := m.Decimal()
		fromCol := vec.Decimal()

		if vec.HasNulls() {
			// TODO(jordan): copy the null arrays in batch.
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					if vec.NullAt64(sel[i]) {
						m.SetNull(i)
					} else {
						toCol[i] = fromCol[sel[i]]
					}
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		}
	case types.Int8:
		toCol := m.Int8()
		fromCol := vec.Int8()

		if vec.HasNulls() {
			// TODO(jordan): copy the null arrays in batch.
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					if vec.NullAt64(sel[i]) {
						m.SetNull(i)
					} else {
						toCol[i] = fromCol[sel[i]]
					}
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		}
	case types.Int16:
		toCol := m.Int16()
		fromCol := vec.Int16()

		if vec.HasNulls() {
			// TODO(jordan): copy the null arrays in batch.
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					if vec.NullAt64(sel[i]) {
						m.SetNull(i)
					} else {
						toCol[i] = fromCol[sel[i]]
					}
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		}
	case types.Int32:
		toCol := m.Int32()
		fromCol := vec.Int32()

		if vec.HasNulls() {
			// TODO(jordan): copy the null arrays in batch.
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					if vec.NullAt64(sel[i]) {
						m.SetNull(i)
					} else {
						toCol[i] = fromCol[sel[i]]
					}
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		}
	case types.Int64:
		toCol := m.Int64()
		fromCol := vec.Int64()

		if vec.HasNulls() {
			// TODO(jordan): copy the null arrays in batch.
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					if vec.NullAt64(sel[i]) {
						m.SetNull(i)
					} else {
						toCol[i] = fromCol[sel[i]]
					}
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		}
	case types.Float32:
		toCol := m.Float32()
		fromCol := vec.Float32()

		if vec.HasNulls() {
			// TODO(jordan): copy the null arrays in batch.
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					if vec.NullAt64(sel[i]) {
						m.SetNull(i)
					} else {
						toCol[i] = fromCol[sel[i]]
					}
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		}
	case types.Float64:
		toCol := m.Float64()
		fromCol := vec.Float64()

		if vec.HasNulls() {
			// TODO(jordan): copy the null arrays in batch.
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					if vec.NullAt64(sel[i]) {
						m.SetNull(i)
					} else {
						toCol[i] = fromCol[sel[i]]
					}
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				if nils[i] {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) Slice(colType types.T, start uint64, end uint64) Vec {
	switch colType {
	case types.Bool:
		col := m.Bool()
		return &memColumn{
			col: col[start:end],
		}
	case types.Bytes:
		col := m.Bytes()
		return &memColumn{
			col: col[start:end],
		}
	case types.Decimal:
		col := m.Decimal()
		return &memColumn{
			col: col[start:end],
		}
	case types.Int8:
		col := m.Int8()
		return &memColumn{
			col: col[start:end],
		}
	case types.Int16:
		col := m.Int16()
		return &memColumn{
			col: col[start:end],
		}
	case types.Int32:
		col := m.Int32()
		return &memColumn{
			col: col[start:end],
		}
	case types.Int64:
		col := m.Int64()
		return &memColumn{
			col: col[start:end],
		}
	case types.Float32:
		col := m.Float32()
		return &memColumn{
			col: col[start:end],
		}
	case types.Float64:
		col := m.Float64()
		return &memColumn{
			col: col[start:end],
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) PrettyValueAt(colIdx uint16, colType types.T) string {
	if m.NullAt(colIdx) {
		return "NULL"
	}
	switch colType {
	case types.Bool:
		return fmt.Sprintf("%v", m.Bool()[colIdx])
	case types.Bytes:
		return fmt.Sprintf("%v", m.Bytes()[colIdx])
	case types.Decimal:
		return fmt.Sprintf("%v", m.Decimal()[colIdx])
	case types.Int8:
		return fmt.Sprintf("%v", m.Int8()[colIdx])
	case types.Int16:
		return fmt.Sprintf("%v", m.Int16()[colIdx])
	case types.Int32:
		return fmt.Sprintf("%v", m.Int32()[colIdx])
	case types.Int64:
		return fmt.Sprintf("%v", m.Int64()[colIdx])
	case types.Float32:
		return fmt.Sprintf("%v", m.Float32()[colIdx])
	case types.Float64:
		return fmt.Sprintf("%v", m.Float64()[colIdx])
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) ExtendNulls(vec Vec, destStartIdx uint64, srcStartIdx uint16, toAppend uint16) {
	outputLen := destStartIdx + uint64(toAppend)
	if uint64(cap(m.nulls)) < outputLen/64 {
		// (batchSize-1)>>6+1 is the number of Int64s needed to encode the additional elements/nulls in the Vec.
		// This is equivalent to ceil(batchSize/64).
		m.nulls = append(m.nulls, make([]int64, (toAppend-1)>>6+1)...)
	}
	if vec.HasNulls() {
		for i := uint16(0); i < toAppend; i++ {
			// TODO(yuzefovich): this can be done more efficiently with a bitwise OR:
			// like m.nulls[i] |= vec.nulls[i].
			if vec.NullAt(srcStartIdx + i) {
				m.SetNull64(destStartIdx + uint64(i))
			}
		}
	}
}

func (m *memColumn) ExtendNullsWithSel(
	vec Vec, destStartIdx uint64, srcStartIdx uint16, toAppend uint16, sel []uint16,
) {
	outputLen := destStartIdx + uint64(toAppend)
	if uint64(cap(m.nulls)) < outputLen/64 {
		// (batchSize-1)>>6+1 is the number of Int64s needed to encode the additional elements/nulls in the Vec.
		// This is equivalent to ceil(batchSize/64).
		m.nulls = append(m.nulls, make([]int64, (toAppend-1)>>6+1)...)
	}
	for i := uint16(0); i < toAppend; i++ {
		// TODO(yuzefovich): this can be done more efficiently with a bitwise OR:
		// like m.nulls[i] |= vec.nulls[i].
		if vec.NullAt(sel[srcStartIdx+i]) {
			m.SetNull64(destStartIdx + uint64(i))
		}
	}
}