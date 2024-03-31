use std::sync::Arc;

use axum::extract::State;
use axum::http::StatusCode;
use axum::Json;
use diesel::{RunQueryDsl, SelectableHelper};

use crate::AppState;
use crate::helpers::internal_err::internal_error;
use crate::models::{CreateTransaction, Transaction};
use crate::schema::transaction;

pub async fn create_transaction(
    State(state): State<Arc<AppState>>,
    Json(new_transaction): Json<CreateTransaction>,
) -> Result<Json<Transaction>, (StatusCode, String)> {
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