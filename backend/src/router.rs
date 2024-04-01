use std::sync::Arc;

use axum::middleware::from_fn_with_state;
use axum::Router;
use axum::routing::{get, post};
use tower_http::cors::CorsLayer;

use crate::AppState;
use crate::handlers::transactions::{create_transaction, create_transaction_type, get_transactions};
use crate::handlers::users::{create_user, get_users, login_user};
use crate::middleware::auth::auth;

pub fn app_router(state: Arc<AppState>) -> Router {
    Router::new()
        .nest("/api/v1/users", users_router(state.clone()))
        .nest("/api/v1/transactions", transactions_router(state.clone()))
}

fn users_router(state: Arc<AppState>) -> Router {
    Router::new()
        .route("/", get(get_users).post(create_user))
        .route("/login", post(login_user))
        .with_state(state)
        .layer(CorsLayer::permissive())
}

fn transactions_router(state: Arc<AppState>) -> Router {
    Router::new()
        .route("/", get(get_transactions).post(create_transaction))
        .route("/type", post(create_transaction_type))
        .layer(from_fn_with_state(state.clone(), auth))
        .with_state(state)
        .layer(CorsLayer::permissive())
}
