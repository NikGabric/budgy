use chrono::NaiveDateTime;
use diesel::data_types::PgMoney;
use diesel::prelude::*;
use serde::{Deserialize, Serialize};

use crate::schema::budgy_user;

#[derive(Queryable, Selectable, Debug, Serialize)]
#[diesel(table_name = crate::schema::budgy_user)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct BudgyUser {
    pub budgy_user_id: i32,
    pub username: String,
    pub password: String,
    pub email: String,
    pub deleted: bool,
    pub inserted_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

#[derive(Insertable, Deserialize)]
#[diesel(table_name = budgy_user)]
pub struct CreateBudgyUser {
    pub username: String,
    pub password: String,
    pub email: String,
}

#[derive(Queryable, Selectable)]
#[diesel(table_name = crate::schema::transaction)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct Transaction {
    pub transaction_id: i32,
    pub title: String,
    pub description: Option<String>,
    pub amount: PgMoney,
    pub deleted: bool,
    pub inserted_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
    pub budgy_user_id: i32,
    pub transaction_type_id: i32,
}

#[derive(Queryable, Selectable)]
#[diesel(table_name = crate::schema::transaction_type)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct TransactionType {
    pub transaction_type_id: i32,
    pub title: String,
    pub description: Option<String>,
    pub deleted: bool,
    pub inserted_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
    pub budgy_user_id: i32,
}
