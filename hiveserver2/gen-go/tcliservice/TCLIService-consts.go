// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package tcliservice

import (
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var PRIMITIVE_TYPES []TTypeId
var COMPLEX_TYPES []TTypeId
var COLLECTION_TYPES []TTypeId
var TYPE_NAMES map[TTypeId]string
const CHARACTER_MAXIMUM_LENGTH = "characterMaximumLength"
const PRECISION = "precision"
const SCALE = "scale"

func init() {
PRIMITIVE_TYPES = []TTypeId{
    0,     1,     2,     3,     4,     5,     6,     7,     8,     9,     15,     16,     17,     18,     19,     20,     21, }

COMPLEX_TYPES = []TTypeId{
    10,     11,     12,     13,     14, }

COLLECTION_TYPES = []TTypeId{
    10,     11, }

TYPE_NAMES = map[TTypeId]string{
    10: "ARRAY",
    4: "BIGINT",
    9: "BINARY",
    0: "BOOLEAN",
    19: "CHAR",
    17: "DATE",
    15: "DECIMAL",
    6: "DOUBLE",
    5: "FLOAT",
    21: "INTERVAL_DAY_TIME",
    20: "INTERVAL_YEAR_MONTH",
    3: "INT",
    11: "MAP",
    16: "NULL",
    2: "SMALLINT",
    7: "STRING",
    12: "STRUCT",
    8: "TIMESTAMP",
    1: "TINYINT",
    13: "UNIONTYPE",
    18: "VARCHAR",
}

}
