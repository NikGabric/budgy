use std::sync::Arc;

use axum::extract::{Request, State};
use axum::http::StatusCode;
use axum::middleware::Next;
use axum::response::IntoResponse;
use diesel::{ExpressionMethods, QueryDsl, RunQueryDsl, SelectableHelper};
use jsonwebtoken::{decode, DecodingKey, Validation};

use crate::AppState;
use crate::helpers::internal_err::internal_error;
use crate::helpers::jwt::TokenClaims;
use crate::models::BudgyUser;
use crate::schema::budgy_user::budgy_user_id;
use crate::schema::budgy_user::dsl::budgy_user;

pub async fn auth(
    State(state): State<Arc<AppState>>,
    mut req: Request,
    next: Next,
) -> Result<impl IntoResponse, (StatusCode, String)> {
    let auth_header = req.headers().get("Authorization").ok_or_else(|| {
        (StatusCode::UNAUTHORIZED, "Authorization token missing");
    })
        .unwrap()
        .to_str()
        .unwrap();
    let token = if auth_header.starts_with("Bearer ") {
        &auth_header[7..]
    } else {
        return Err((StatusCode::UNAUTHORIZED, String::from("Authorization token invalid")));
    };

    let claims = decode::<TokenClaims>(
        &token,
        &DecodingKey::from_secret(state.config.jwt_secret.as_ref()),
        &Validation::default(),
    )
        .map_err(|err| {
            println!("{:?}", err);
            (StatusCode::UNAUTHORIZED, "Authorization token missing 1")
        })
        .unwrap()
        .claims;

    let conn = state.pool.get().await.map_err(internal_error)?;
    let user = conn.interact(move |conn| {
        budgy_user
            .filter(budgy_user_id.eq(claims.sub.parse::<i32>().unwrap()))
            .select(BudgyUser::as_select())
            .first(conn)
    })
        .await
        .map_err(internal_error)?
        .map_err(internal_error)?;

    req.extensions_mut().insert(user);
    Ok(next.run(req).await)
}