// Code generated by ent, DO NOT EDIT.

package addres

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/skyisboss/pay-system/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Addres {
	return predicate.Addres(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Addres {
	return predicate.Addres(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Addres {
	return predicate.Addres(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Addres {
	return predicate.Addres(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Addres {
	return predicate.Addres(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Addres {
	return predicate.Addres(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Addres {
	return predicate.Addres(sql.FieldLTE(FieldID, id))
}

// ChainID applies equality check predicate on the "chain_id" field. It's identical to ChainIDEQ.
func ChainID(v uint64) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldChainID, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldAddress, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldPassword, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldUUID, v))
}

// UseTo applies equality check predicate on the "use_to" field. It's identical to UseToEQ.
func UseTo(v int64) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldUseTo, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldDeletedAt, v))
}

// ChainIDEQ applies the EQ predicate on the "chain_id" field.
func ChainIDEQ(v uint64) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldChainID, v))
}

// ChainIDNEQ applies the NEQ predicate on the "chain_id" field.
func ChainIDNEQ(v uint64) predicate.Addres {
	return predicate.Addres(sql.FieldNEQ(FieldChainID, v))
}

// ChainIDIn applies the In predicate on the "chain_id" field.
func ChainIDIn(vs ...uint64) predicate.Addres {
	return predicate.Addres(sql.FieldIn(FieldChainID, vs...))
}

// ChainIDNotIn applies the NotIn predicate on the "chain_id" field.
func ChainIDNotIn(vs ...uint64) predicate.Addres {
	return predicate.Addres(sql.FieldNotIn(FieldChainID, vs...))
}

// ChainIDGT applies the GT predicate on the "chain_id" field.
func ChainIDGT(v uint64) predicate.Addres {
	return predicate.Addres(sql.FieldGT(FieldChainID, v))
}

// ChainIDGTE applies the GTE predicate on the "chain_id" field.
func ChainIDGTE(v uint64) predicate.Addres {
	return predicate.Addres(sql.FieldGTE(FieldChainID, v))
}

// ChainIDLT applies the LT predicate on the "chain_id" field.
func ChainIDLT(v uint64) predicate.Addres {
	return predicate.Addres(sql.FieldLT(FieldChainID, v))
}

// ChainIDLTE applies the LTE predicate on the "chain_id" field.
func ChainIDLTE(v uint64) predicate.Addres {
	return predicate.Addres(sql.FieldLTE(FieldChainID, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Addres {
	return predicate.Addres(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Addres {
	return predicate.Addres(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Addres {
	return predicate.Addres(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Addres {
	return predicate.Addres(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Addres {
	return predicate.Addres(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Addres {
	return predicate.Addres(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Addres {
	return predicate.Addres(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Addres {
	return predicate.Addres(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Addres {
	return predicate.Addres(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Addres {
	return predicate.Addres(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Addres {
	return predicate.Addres(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Addres {
	return predicate.Addres(sql.FieldContainsFold(FieldAddress, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Addres {
	return predicate.Addres(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Addres {
	return predicate.Addres(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Addres {
	return predicate.Addres(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Addres {
	return predicate.Addres(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Addres {
	return predicate.Addres(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Addres {
	return predicate.Addres(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Addres {
	return predicate.Addres(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Addres {
	return predicate.Addres(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Addres {
	return predicate.Addres(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Addres {
	return predicate.Addres(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Addres {
	return predicate.Addres(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Addres {
	return predicate.Addres(sql.FieldContainsFold(FieldPassword, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Addres {
	return predicate.Addres(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Addres {
	return predicate.Addres(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Addres {
	return predicate.Addres(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Addres {
	return predicate.Addres(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Addres {
	return predicate.Addres(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Addres {
	return predicate.Addres(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Addres {
	return predicate.Addres(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Addres {
	return predicate.Addres(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Addres {
	return predicate.Addres(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Addres {
	return predicate.Addres(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Addres {
	return predicate.Addres(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Addres {
	return predicate.Addres(sql.FieldContainsFold(FieldUUID, v))
}

// UseToEQ applies the EQ predicate on the "use_to" field.
func UseToEQ(v int64) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldUseTo, v))
}

// UseToNEQ applies the NEQ predicate on the "use_to" field.
func UseToNEQ(v int64) predicate.Addres {
	return predicate.Addres(sql.FieldNEQ(FieldUseTo, v))
}

// UseToIn applies the In predicate on the "use_to" field.
func UseToIn(vs ...int64) predicate.Addres {
	return predicate.Addres(sql.FieldIn(FieldUseTo, vs...))
}

// UseToNotIn applies the NotIn predicate on the "use_to" field.
func UseToNotIn(vs ...int64) predicate.Addres {
	return predicate.Addres(sql.FieldNotIn(FieldUseTo, vs...))
}

// UseToGT applies the GT predicate on the "use_to" field.
func UseToGT(v int64) predicate.Addres {
	return predicate.Addres(sql.FieldGT(FieldUseTo, v))
}

// UseToGTE applies the GTE predicate on the "use_to" field.
func UseToGTE(v int64) predicate.Addres {
	return predicate.Addres(sql.FieldGTE(FieldUseTo, v))
}

// UseToLT applies the LT predicate on the "use_to" field.
func UseToLT(v int64) predicate.Addres {
	return predicate.Addres(sql.FieldLT(FieldUseTo, v))
}

// UseToLTE applies the LTE predicate on the "use_to" field.
func UseToLTE(v int64) predicate.Addres {
	return predicate.Addres(sql.FieldLTE(FieldUseTo, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Addres {
	return predicate.Addres(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Addres {
	return predicate.Addres(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Addres {
	return predicate.Addres(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Addres {
	return predicate.Addres(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Addres {
	return predicate.Addres(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Addres {
	return predicate.Addres(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Addres {
	return predicate.Addres(sql.FieldNotNull(FieldDeletedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Addres) predicate.Addres {
	return predicate.Addres(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Addres) predicate.Addres {
	return predicate.Addres(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Addres) predicate.Addres {
	return predicate.Addres(sql.NotPredicates(p))
}
