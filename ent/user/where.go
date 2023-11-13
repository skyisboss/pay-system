// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/skyisboss/pay-system/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeletedAt, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// AuthGoogle applies equality check predicate on the "auth_google" field. It's identical to AuthGoogleEQ.
func AuthGoogle(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAuthGoogle, v))
}

// AuthToken applies equality check predicate on the "auth_token" field. It's identical to AuthTokenEQ.
func AuthToken(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAuthToken, v))
}

// Setting applies equality check predicate on the "setting" field. It's identical to SettingEQ.
func Setting(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSetting, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDeletedAt))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRole, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// AuthGoogleEQ applies the EQ predicate on the "auth_google" field.
func AuthGoogleEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAuthGoogle, v))
}

// AuthGoogleNEQ applies the NEQ predicate on the "auth_google" field.
func AuthGoogleNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAuthGoogle, v))
}

// AuthGoogleIn applies the In predicate on the "auth_google" field.
func AuthGoogleIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAuthGoogle, vs...))
}

// AuthGoogleNotIn applies the NotIn predicate on the "auth_google" field.
func AuthGoogleNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAuthGoogle, vs...))
}

// AuthGoogleGT applies the GT predicate on the "auth_google" field.
func AuthGoogleGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAuthGoogle, v))
}

// AuthGoogleGTE applies the GTE predicate on the "auth_google" field.
func AuthGoogleGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAuthGoogle, v))
}

// AuthGoogleLT applies the LT predicate on the "auth_google" field.
func AuthGoogleLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAuthGoogle, v))
}

// AuthGoogleLTE applies the LTE predicate on the "auth_google" field.
func AuthGoogleLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAuthGoogle, v))
}

// AuthGoogleContains applies the Contains predicate on the "auth_google" field.
func AuthGoogleContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAuthGoogle, v))
}

// AuthGoogleHasPrefix applies the HasPrefix predicate on the "auth_google" field.
func AuthGoogleHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAuthGoogle, v))
}

// AuthGoogleHasSuffix applies the HasSuffix predicate on the "auth_google" field.
func AuthGoogleHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAuthGoogle, v))
}

// AuthGoogleEqualFold applies the EqualFold predicate on the "auth_google" field.
func AuthGoogleEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAuthGoogle, v))
}

// AuthGoogleContainsFold applies the ContainsFold predicate on the "auth_google" field.
func AuthGoogleContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAuthGoogle, v))
}

// AuthTokenEQ applies the EQ predicate on the "auth_token" field.
func AuthTokenEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAuthToken, v))
}

// AuthTokenNEQ applies the NEQ predicate on the "auth_token" field.
func AuthTokenNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAuthToken, v))
}

// AuthTokenIn applies the In predicate on the "auth_token" field.
func AuthTokenIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAuthToken, vs...))
}

// AuthTokenNotIn applies the NotIn predicate on the "auth_token" field.
func AuthTokenNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAuthToken, vs...))
}

// AuthTokenGT applies the GT predicate on the "auth_token" field.
func AuthTokenGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAuthToken, v))
}

// AuthTokenGTE applies the GTE predicate on the "auth_token" field.
func AuthTokenGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAuthToken, v))
}

// AuthTokenLT applies the LT predicate on the "auth_token" field.
func AuthTokenLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAuthToken, v))
}

// AuthTokenLTE applies the LTE predicate on the "auth_token" field.
func AuthTokenLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAuthToken, v))
}

// AuthTokenContains applies the Contains predicate on the "auth_token" field.
func AuthTokenContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAuthToken, v))
}

// AuthTokenHasPrefix applies the HasPrefix predicate on the "auth_token" field.
func AuthTokenHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAuthToken, v))
}

// AuthTokenHasSuffix applies the HasSuffix predicate on the "auth_token" field.
func AuthTokenHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAuthToken, v))
}

// AuthTokenEqualFold applies the EqualFold predicate on the "auth_token" field.
func AuthTokenEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAuthToken, v))
}

// AuthTokenContainsFold applies the ContainsFold predicate on the "auth_token" field.
func AuthTokenContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAuthToken, v))
}

// SettingEQ applies the EQ predicate on the "setting" field.
func SettingEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSetting, v))
}

// SettingNEQ applies the NEQ predicate on the "setting" field.
func SettingNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSetting, v))
}

// SettingIn applies the In predicate on the "setting" field.
func SettingIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldSetting, vs...))
}

// SettingNotIn applies the NotIn predicate on the "setting" field.
func SettingNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSetting, vs...))
}

// SettingGT applies the GT predicate on the "setting" field.
func SettingGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldSetting, v))
}

// SettingGTE applies the GTE predicate on the "setting" field.
func SettingGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldSetting, v))
}

// SettingLT applies the LT predicate on the "setting" field.
func SettingLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldSetting, v))
}

// SettingLTE applies the LTE predicate on the "setting" field.
func SettingLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldSetting, v))
}

// SettingContains applies the Contains predicate on the "setting" field.
func SettingContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldSetting, v))
}

// SettingHasPrefix applies the HasPrefix predicate on the "setting" field.
func SettingHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldSetting, v))
}

// SettingHasSuffix applies the HasSuffix predicate on the "setting" field.
func SettingHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldSetting, v))
}

// SettingEqualFold applies the EqualFold predicate on the "setting" field.
func SettingEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldSetting, v))
}

// SettingContainsFold applies the ContainsFold predicate on the "setting" field.
func SettingContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldSetting, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
