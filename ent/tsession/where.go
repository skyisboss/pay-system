// Code generated by ent, DO NOT EDIT.

package tsession

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/skyisboss/pay-system/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.TSession {
	return predicate.TSession(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.TSession {
	return predicate.TSession(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.TSession {
	return predicate.TSession(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.TSession {
	return predicate.TSession(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.TSession {
	return predicate.TSession(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.TSession {
	return predicate.TSession(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.TSession {
	return predicate.TSession(sql.FieldLTE(FieldID, id))
}

// KeyName applies equality check predicate on the "key_name" field. It's identical to KeyNameEQ.
func KeyName(v string) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldKeyName, v))
}

// KeyValue applies equality check predicate on the "key_value" field. It's identical to KeyValueEQ.
func KeyValue(v string) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldKeyValue, v))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldIP, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldUpdatedAt, v))
}

// KeyNameEQ applies the EQ predicate on the "key_name" field.
func KeyNameEQ(v string) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldKeyName, v))
}

// KeyNameNEQ applies the NEQ predicate on the "key_name" field.
func KeyNameNEQ(v string) predicate.TSession {
	return predicate.TSession(sql.FieldNEQ(FieldKeyName, v))
}

// KeyNameIn applies the In predicate on the "key_name" field.
func KeyNameIn(vs ...string) predicate.TSession {
	return predicate.TSession(sql.FieldIn(FieldKeyName, vs...))
}

// KeyNameNotIn applies the NotIn predicate on the "key_name" field.
func KeyNameNotIn(vs ...string) predicate.TSession {
	return predicate.TSession(sql.FieldNotIn(FieldKeyName, vs...))
}

// KeyNameGT applies the GT predicate on the "key_name" field.
func KeyNameGT(v string) predicate.TSession {
	return predicate.TSession(sql.FieldGT(FieldKeyName, v))
}

// KeyNameGTE applies the GTE predicate on the "key_name" field.
func KeyNameGTE(v string) predicate.TSession {
	return predicate.TSession(sql.FieldGTE(FieldKeyName, v))
}

// KeyNameLT applies the LT predicate on the "key_name" field.
func KeyNameLT(v string) predicate.TSession {
	return predicate.TSession(sql.FieldLT(FieldKeyName, v))
}

// KeyNameLTE applies the LTE predicate on the "key_name" field.
func KeyNameLTE(v string) predicate.TSession {
	return predicate.TSession(sql.FieldLTE(FieldKeyName, v))
}

// KeyNameContains applies the Contains predicate on the "key_name" field.
func KeyNameContains(v string) predicate.TSession {
	return predicate.TSession(sql.FieldContains(FieldKeyName, v))
}

// KeyNameHasPrefix applies the HasPrefix predicate on the "key_name" field.
func KeyNameHasPrefix(v string) predicate.TSession {
	return predicate.TSession(sql.FieldHasPrefix(FieldKeyName, v))
}

// KeyNameHasSuffix applies the HasSuffix predicate on the "key_name" field.
func KeyNameHasSuffix(v string) predicate.TSession {
	return predicate.TSession(sql.FieldHasSuffix(FieldKeyName, v))
}

// KeyNameEqualFold applies the EqualFold predicate on the "key_name" field.
func KeyNameEqualFold(v string) predicate.TSession {
	return predicate.TSession(sql.FieldEqualFold(FieldKeyName, v))
}

// KeyNameContainsFold applies the ContainsFold predicate on the "key_name" field.
func KeyNameContainsFold(v string) predicate.TSession {
	return predicate.TSession(sql.FieldContainsFold(FieldKeyName, v))
}

// KeyValueEQ applies the EQ predicate on the "key_value" field.
func KeyValueEQ(v string) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldKeyValue, v))
}

// KeyValueNEQ applies the NEQ predicate on the "key_value" field.
func KeyValueNEQ(v string) predicate.TSession {
	return predicate.TSession(sql.FieldNEQ(FieldKeyValue, v))
}

// KeyValueIn applies the In predicate on the "key_value" field.
func KeyValueIn(vs ...string) predicate.TSession {
	return predicate.TSession(sql.FieldIn(FieldKeyValue, vs...))
}

// KeyValueNotIn applies the NotIn predicate on the "key_value" field.
func KeyValueNotIn(vs ...string) predicate.TSession {
	return predicate.TSession(sql.FieldNotIn(FieldKeyValue, vs...))
}

// KeyValueGT applies the GT predicate on the "key_value" field.
func KeyValueGT(v string) predicate.TSession {
	return predicate.TSession(sql.FieldGT(FieldKeyValue, v))
}

// KeyValueGTE applies the GTE predicate on the "key_value" field.
func KeyValueGTE(v string) predicate.TSession {
	return predicate.TSession(sql.FieldGTE(FieldKeyValue, v))
}

// KeyValueLT applies the LT predicate on the "key_value" field.
func KeyValueLT(v string) predicate.TSession {
	return predicate.TSession(sql.FieldLT(FieldKeyValue, v))
}

// KeyValueLTE applies the LTE predicate on the "key_value" field.
func KeyValueLTE(v string) predicate.TSession {
	return predicate.TSession(sql.FieldLTE(FieldKeyValue, v))
}

// KeyValueContains applies the Contains predicate on the "key_value" field.
func KeyValueContains(v string) predicate.TSession {
	return predicate.TSession(sql.FieldContains(FieldKeyValue, v))
}

// KeyValueHasPrefix applies the HasPrefix predicate on the "key_value" field.
func KeyValueHasPrefix(v string) predicate.TSession {
	return predicate.TSession(sql.FieldHasPrefix(FieldKeyValue, v))
}

// KeyValueHasSuffix applies the HasSuffix predicate on the "key_value" field.
func KeyValueHasSuffix(v string) predicate.TSession {
	return predicate.TSession(sql.FieldHasSuffix(FieldKeyValue, v))
}

// KeyValueEqualFold applies the EqualFold predicate on the "key_value" field.
func KeyValueEqualFold(v string) predicate.TSession {
	return predicate.TSession(sql.FieldEqualFold(FieldKeyValue, v))
}

// KeyValueContainsFold applies the ContainsFold predicate on the "key_value" field.
func KeyValueContainsFold(v string) predicate.TSession {
	return predicate.TSession(sql.FieldContainsFold(FieldKeyValue, v))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.TSession {
	return predicate.TSession(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.TSession {
	return predicate.TSession(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.TSession {
	return predicate.TSession(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.TSession {
	return predicate.TSession(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.TSession {
	return predicate.TSession(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.TSession {
	return predicate.TSession(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.TSession {
	return predicate.TSession(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.TSession {
	return predicate.TSession(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.TSession {
	return predicate.TSession(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.TSession {
	return predicate.TSession(sql.FieldHasSuffix(FieldIP, v))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.TSession {
	return predicate.TSession(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.TSession {
	return predicate.TSession(sql.FieldContainsFold(FieldIP, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.TSession {
	return predicate.TSession(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.TSession {
	return predicate.TSession(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TSession {
	return predicate.TSession(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.TSession {
	return predicate.TSession(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.TSession {
	return predicate.TSession(sql.FieldNotNull(FieldUpdatedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TSession) predicate.TSession {
	return predicate.TSession(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TSession) predicate.TSession {
	return predicate.TSession(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TSession) predicate.TSession {
	return predicate.TSession(sql.NotPredicates(p))
}
