use std::sync::Arc;

use axum::debug_handler;
use axum::extract::State;
use axum::http::StatusCode;
use axum::response::Json;
use diesel::prelude::*;
use jsonwebtoken::{encode, EncodingKey, Header};
use serde::Deserialize;

use crate::AppState;
use crate::helpers::internal_err::internal_error;
use crate::helpers::jwt::{TokenClaims, UserClaims};
use crate::models::{BudgyUser, CreateBudgyUser};
use crate::schema::budgy_user;
use crate::schema::budgy_user::username;

pub async fn get_users(
    State(state): State<Arc<AppState>>,
) -> Result<Json<Vec<BudgyUser>>, (StatusCode, String)> {
    let conn = state.pool.get().await.map_err(internal_error)?;
    let res = conn
        .interact(|conn| budgy_user::table.select(BudgyUser::as_select()).load(conn))
        .await
        .map_err(internal_error)?
        .map_err(internal_error)?;

    Ok(Json(res))
}

#[derive(Deserialize)]
pub struct LoginBudgyUser {
    pub username: String,
    pub password: String,
}

#[debug_handler]
pub async fn login_user(
    State(state): State<Arc<AppState>>,
    Json(login_budgy_user): Json<LoginBudgyUser>,
) -> Result<String, (StatusCode, String)> {
    let conn = state.pool.get().await.map_err(internal_error)?;
    let db_user = conn
        .interact(|conn| {
            budgy_user::table
                .filter(username.eq(login_budgy_user.username))
                .select(BudgyUser::as_select())
                .first(conn)
        })
        .await
        .map_err(internal_error)?
        .map_err(internal_error)?;

    let is_valid = login_budgy_user.password == db_user.password;
    if !is_valid {
        return Err((StatusCode::BAD_REQUEST, String::from("Invalid username or password")));
    }

    let now = chrono::Utc::now();
    let iat = now.timestamp() as usize;
    let exp = (now + chrono::Duration::minutes(60)).timestamp() as usize;
    let claims = TokenClaims {
        sub: db_user.budgy_user_id.to_string(),
        exp,
        iat,
        usr: UserClaims {
            username: db_user.username,
            email: db_user.email,
        },
    };

    let token = encode(
        &Header::default(),
        &claims,
        &EncodingKey::from_secret(state.config.jwt_secret.as_ref()),
    ).unwrap();

    Ok(token)
}

#[debug_handler]
pub async fn create_user(
    State(state): State<Arc<AppState>>,
    Json(new_user): Json<CreateBudgyUser>,
) -> Result<Json<BudgyUser>, (StatusCode, String)> {
    let conn = state.pool.get().await.map_err(internal_error)?;
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
