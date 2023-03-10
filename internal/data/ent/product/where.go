// Code generated by ent, DO NOT EDIT.

package product

import (
	"gy_home/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Recommend applies equality check predicate on the "recommend" field. It's identical to RecommendEQ.
func Recommend(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecommend), v))
	})
}

// History applies equality check predicate on the "history" field. It's identical to HistoryEQ.
func History(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHistory), v))
	})
}

// CategoryID applies equality check predicate on the "category_id" field. It's identical to CategoryIDEQ.
func CategoryID(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryID), v))
	})
}

// CategoryName applies equality check predicate on the "category_name" field. It's identical to CategoryNameEQ.
func CategoryName(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryName), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Utime applies equality check predicate on the "utime" field. It's identical to UtimeEQ.
func Utime(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUtime), v))
	})
}

// ServiceLink applies equality check predicate on the "service_link" field. It's identical to ServiceLinkEQ.
func ServiceLink(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceLink), v))
	})
}

// ApplyLink applies equality check predicate on the "apply_link" field. It's identical to ApplyLinkEQ.
func ApplyLink(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplyLink), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// RecommendEQ applies the EQ predicate on the "recommend" field.
func RecommendEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecommend), v))
	})
}

// RecommendNEQ applies the NEQ predicate on the "recommend" field.
func RecommendNEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecommend), v))
	})
}

// RecommendIn applies the In predicate on the "recommend" field.
func RecommendIn(vs ...int) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRecommend), v...))
	})
}

// RecommendNotIn applies the NotIn predicate on the "recommend" field.
func RecommendNotIn(vs ...int) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRecommend), v...))
	})
}

// RecommendGT applies the GT predicate on the "recommend" field.
func RecommendGT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecommend), v))
	})
}

// RecommendGTE applies the GTE predicate on the "recommend" field.
func RecommendGTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecommend), v))
	})
}

// RecommendLT applies the LT predicate on the "recommend" field.
func RecommendLT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecommend), v))
	})
}

// RecommendLTE applies the LTE predicate on the "recommend" field.
func RecommendLTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecommend), v))
	})
}

// HistoryEQ applies the EQ predicate on the "history" field.
func HistoryEQ(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHistory), v))
	})
}

// HistoryNEQ applies the NEQ predicate on the "history" field.
func HistoryNEQ(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHistory), v))
	})
}

// HistoryIn applies the In predicate on the "history" field.
func HistoryIn(vs ...int32) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHistory), v...))
	})
}

// HistoryNotIn applies the NotIn predicate on the "history" field.
func HistoryNotIn(vs ...int32) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHistory), v...))
	})
}

// HistoryGT applies the GT predicate on the "history" field.
func HistoryGT(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHistory), v))
	})
}

// HistoryGTE applies the GTE predicate on the "history" field.
func HistoryGTE(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHistory), v))
	})
}

// HistoryLT applies the LT predicate on the "history" field.
func HistoryLT(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHistory), v))
	})
}

// HistoryLTE applies the LTE predicate on the "history" field.
func HistoryLTE(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHistory), v))
	})
}

// CategoryIDEQ applies the EQ predicate on the "category_id" field.
func CategoryIDEQ(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryID), v))
	})
}

// CategoryIDNEQ applies the NEQ predicate on the "category_id" field.
func CategoryIDNEQ(v int32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategoryID), v))
	})
}

// CategoryIDIn applies the In predicate on the "category_id" field.
func CategoryIDIn(vs ...int32) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCategoryID), v...))
	})
}

// CategoryIDNotIn applies the NotIn predicate on the "category_id" field.
func CategoryIDNotIn(vs ...int32) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCategoryID), v...))
	})
}

// CategoryIDIsNil applies the IsNil predicate on the "category_id" field.
func CategoryIDIsNil() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCategoryID)))
	})
}

// CategoryIDNotNil applies the NotNil predicate on the "category_id" field.
func CategoryIDNotNil() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCategoryID)))
	})
}

// CategoryNameEQ applies the EQ predicate on the "category_name" field.
func CategoryNameEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryName), v))
	})
}

// CategoryNameNEQ applies the NEQ predicate on the "category_name" field.
func CategoryNameNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategoryName), v))
	})
}

// CategoryNameIn applies the In predicate on the "category_name" field.
func CategoryNameIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCategoryName), v...))
	})
}

// CategoryNameNotIn applies the NotIn predicate on the "category_name" field.
func CategoryNameNotIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCategoryName), v...))
	})
}

// CategoryNameGT applies the GT predicate on the "category_name" field.
func CategoryNameGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCategoryName), v))
	})
}

// CategoryNameGTE applies the GTE predicate on the "category_name" field.
func CategoryNameGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCategoryName), v))
	})
}

// CategoryNameLT applies the LT predicate on the "category_name" field.
func CategoryNameLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCategoryName), v))
	})
}

// CategoryNameLTE applies the LTE predicate on the "category_name" field.
func CategoryNameLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCategoryName), v))
	})
}

// CategoryNameContains applies the Contains predicate on the "category_name" field.
func CategoryNameContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCategoryName), v))
	})
}

// CategoryNameHasPrefix applies the HasPrefix predicate on the "category_name" field.
func CategoryNameHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCategoryName), v))
	})
}

// CategoryNameHasSuffix applies the HasSuffix predicate on the "category_name" field.
func CategoryNameHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCategoryName), v))
	})
}

// CategoryNameEqualFold applies the EqualFold predicate on the "category_name" field.
func CategoryNameEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCategoryName), v))
	})
}

// CategoryNameContainsFold applies the ContainsFold predicate on the "category_name" field.
func CategoryNameContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCategoryName), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// UtimeEQ applies the EQ predicate on the "utime" field.
func UtimeEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUtime), v))
	})
}

// UtimeNEQ applies the NEQ predicate on the "utime" field.
func UtimeNEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUtime), v))
	})
}

// UtimeIn applies the In predicate on the "utime" field.
func UtimeIn(vs ...time.Time) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUtime), v...))
	})
}

// UtimeNotIn applies the NotIn predicate on the "utime" field.
func UtimeNotIn(vs ...time.Time) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUtime), v...))
	})
}

// UtimeGT applies the GT predicate on the "utime" field.
func UtimeGT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUtime), v))
	})
}

// UtimeGTE applies the GTE predicate on the "utime" field.
func UtimeGTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUtime), v))
	})
}

// UtimeLT applies the LT predicate on the "utime" field.
func UtimeLT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUtime), v))
	})
}

// UtimeLTE applies the LTE predicate on the "utime" field.
func UtimeLTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUtime), v))
	})
}

// ServiceLinkEQ applies the EQ predicate on the "service_link" field.
func ServiceLinkEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkNEQ applies the NEQ predicate on the "service_link" field.
func ServiceLinkNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkIn applies the In predicate on the "service_link" field.
func ServiceLinkIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldServiceLink), v...))
	})
}

// ServiceLinkNotIn applies the NotIn predicate on the "service_link" field.
func ServiceLinkNotIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldServiceLink), v...))
	})
}

// ServiceLinkGT applies the GT predicate on the "service_link" field.
func ServiceLinkGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkGTE applies the GTE predicate on the "service_link" field.
func ServiceLinkGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkLT applies the LT predicate on the "service_link" field.
func ServiceLinkLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkLTE applies the LTE predicate on the "service_link" field.
func ServiceLinkLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkContains applies the Contains predicate on the "service_link" field.
func ServiceLinkContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkHasPrefix applies the HasPrefix predicate on the "service_link" field.
func ServiceLinkHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkHasSuffix applies the HasSuffix predicate on the "service_link" field.
func ServiceLinkHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkEqualFold applies the EqualFold predicate on the "service_link" field.
func ServiceLinkEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldServiceLink), v))
	})
}

// ServiceLinkContainsFold applies the ContainsFold predicate on the "service_link" field.
func ServiceLinkContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldServiceLink), v))
	})
}

// ApplyLinkEQ applies the EQ predicate on the "apply_link" field.
func ApplyLinkEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkNEQ applies the NEQ predicate on the "apply_link" field.
func ApplyLinkNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkIn applies the In predicate on the "apply_link" field.
func ApplyLinkIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldApplyLink), v...))
	})
}

// ApplyLinkNotIn applies the NotIn predicate on the "apply_link" field.
func ApplyLinkNotIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldApplyLink), v...))
	})
}

// ApplyLinkGT applies the GT predicate on the "apply_link" field.
func ApplyLinkGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkGTE applies the GTE predicate on the "apply_link" field.
func ApplyLinkGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkLT applies the LT predicate on the "apply_link" field.
func ApplyLinkLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkLTE applies the LTE predicate on the "apply_link" field.
func ApplyLinkLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkContains applies the Contains predicate on the "apply_link" field.
func ApplyLinkContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkHasPrefix applies the HasPrefix predicate on the "apply_link" field.
func ApplyLinkHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkHasSuffix applies the HasSuffix predicate on the "apply_link" field.
func ApplyLinkHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkEqualFold applies the EqualFold predicate on the "apply_link" field.
func ApplyLinkEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldApplyLink), v))
	})
}

// ApplyLinkContainsFold applies the ContainsFold predicate on the "apply_link" field.
func ApplyLinkContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldApplyLink), v))
	})
}

// HasParentID applies the HasEdge predicate on the "parent_id" edge.
func HasParentID() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentIDTable, ParentIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentIDWith applies the HasEdge predicate on the "parent_id" edge with a given conditions (other predicates).
func HasParentIDWith(preds ...predicate.ProductCategory) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentIDTable, ParentIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		p(s.Not())
	})
}
