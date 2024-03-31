use axum::Router;
use axum::routing::{get, post};
use deadpool_diesel::Pool;
use deadpool_diesel::postgres::Manager;

use crate::handlers::transactions::create_transaction;
use crate::handlers::users::{create_user, get_users};

pub fn app_router(state: Pool<Manager>) -> Router {
    Router::new()
        .nest("/api/v1/users", users_router(state.clone()))
        .nest("/api/v1/transactions", transactions_router(state.clone()))
}

fn users_router(state: Pool<Manager>) -> Router {
    Router::new()
        .route("/", get(get_users).post(create_user))
        .with_state(state)
}

fn transactions_router(state: Pool<Manager>) -> Router {
    Router::new()
        .route("/", post(create_transaction))
        .with_state(state)
}
