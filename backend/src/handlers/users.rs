use axum::{debug_handler, response::Json};
use axum::extract::State;
use axum::http::StatusCode;
use deadpool_diesel::Pool;
use deadpool_diesel::postgres::Manager;
use diesel::prelude::*;

use crate::helpers::internal_err::internal_error;
use crate::models::{BudgyUser, CreateBudgyUser};
use crate::schema::budgy_user;

#[debug_handler]
pub async fn get_users(
    State(state): State<Pool<Manager>>
) -> Result<Json<Vec<BudgyUser>>, (StatusCode, String)> {
    let conn = state.get().await.map_err(internal_error)?;
    let res = conn
        .interact(|conn| budgy_user::table.select(BudgyUser::as_select()).load(conn))
        .await
        .map_err(internal_error)?
        .map_err(internal_error)?;

    Ok(Json(res))
}

#[debug_handler]
pub async fn create_user(
    State(state): State<Pool<Manager>>,
    Json(new_user): Json<CreateBudgyUser>,
) -> Result<Json<BudgyUser>, (StatusCode, String)> {
    let conn = state.get().await.map_err(internal_error)?;
    let res = conn
        .interact(|conn| {
            diesel::insert_into(budgy_user::table)
                .values(new_user)
                .returning(BudgyUser::as_returning())
                .get_result(conn)
        })
        .await
        .map_err(internal_error)?
        .map_err(internal_error)?;

    Ok(Json(res))
}
