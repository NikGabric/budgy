use std::sync::Arc;

use axum::{Extension, Json};
use axum::extract::State;
use axum::http::StatusCode;
use diesel::{RunQueryDsl, SelectableHelper};

use crate::AppState;
use crate::helpers::internal_err::internal_error;
use crate::models::{BudgyUser, CreateTransaction, CreateTransactionDto, CreateTransactionType, CreateTransactionTypeDto, Transaction, TransactionType};
use crate::schema::{transaction, transaction_type};

pub async fn create_transaction(
    Extension(user): Extension<BudgyUser>,
    State(state): State<Arc<AppState>>,
    Json(new_transaction_dto): Json<CreateTransactionDto>,
) -> Result<Json<Transaction>, (StatusCode, String)> {
    let new_transaction = CreateTransaction {
        title: new_transaction_dto.title,
        description: new_transaction_dto.description,
        amount: new_transaction_dto.amount,
        transaction_type_id: new_transaction_dto.transaction_type_id,
        budgy_user_id: user.budgy_user_id,
    };

    let conn = state.pool.get().await.map_err(internal_error)?;
    let res = conn
        .interact(|conn| {
            diesel::insert_into(transaction::table)
                .values(new_transaction)
                .returning(Transaction::as_returning())
                .get_result(conn)
        })
        .await
        .map_err(internal_error)?
        .map_err(internal_error)?;

    Ok(Json(res))
}

pub async fn create_transaction_type(
    Extension(user): Extension<BudgyUser>,
    State(state): State<Arc<AppState>>,
    Json(new_transaction_dto): Json<CreateTransactionTypeDto>,
) -> Result<Json<TransactionType>, (StatusCode, String)> {
    let new_transaction_type = CreateTransactionType {
        title: new_transaction_dto.title,
        description: new_transaction_dto.description,
        budgy_user_id: user.budgy_user_id,
    };

    let conn = state.pool.get().await.map_err(internal_error)?;
    let res = conn
        .interact(|conn| {
            diesel::insert_into(transaction_type::table)
                .values(new_transaction_type)
                .returning(TransactionType::as_returning())
                .get_result(conn)
        })
        .await
        .map_err(internal_error)?
        .map_err(internal_error)?;

    Ok(Json(res))
}