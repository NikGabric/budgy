// @generated automatically by Diesel CLI.

diesel::table! {
    budgy_user (budgy_user_id) {
        budgy_user_id -> Int4,
        username -> Varchar,
        password -> Varchar,
        email -> Varchar,
        deleted -> Bool,
        inserted_at -> Timestamptz,
        updated_at -> Timestamptz,
    }
}

diesel::table! {
    transaction (transaction_id) {
        transaction_id -> Int4,
        title -> Varchar,
        description -> Nullable<Varchar>,
        amount -> Int4,
        deleted -> Bool,
        inserted_at -> Timestamptz,
        updated_at -> Timestamptz,
        budgy_user_id -> Int4,
        transaction_type_id -> Int4,
    }
}

diesel::table! {
    transaction_type (transaction_type_id) {
        transaction_type_id -> Int4,
        title -> Varchar,
        description -> Nullable<Varchar>,
        deleted -> Bool,
        inserted_at -> Timestamptz,
        updated_at -> Timestamptz,
        budgy_user_id -> Int4,
    }
}

diesel::joinable!(transaction -> budgy_user (budgy_user_id));
diesel::joinable!(transaction -> transaction_type (transaction_type_id));
diesel::joinable!(transaction_type -> budgy_user (budgy_user_id));

diesel::allow_tables_to_appear_in_same_query!(
    budgy_user,
    transaction,
    transaction_type,
);
