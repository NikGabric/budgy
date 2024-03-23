// @generated automatically by Diesel CLI.

diesel::table! {
    budgy_user (user_id) {
        user_id -> Int4,
        #[max_length = 50]
        username -> Varchar,
        #[max_length = 100]
        password -> Varchar,
        #[max_length = 100]
        email -> Varchar,
    }
}

diesel::table! {
    category (category_id) {
        category_id -> Int4,
        #[max_length = 100]
        category_name -> Varchar,
        user_id -> Int4,
    }
}

diesel::table! {
    transaction (transaction_id) {
        transaction_id -> Int4,
        #[max_length = 255]
        description -> Varchar,
        amount -> Numeric,
        category_id -> Int4,
        user_id -> Int4,
        created_at -> Nullable<Timestamp>,
    }
}

diesel::joinable!(category -> budgy_user (user_id));
diesel::joinable!(transaction -> budgy_user (user_id));
diesel::joinable!(transaction -> category (category_id));

diesel::allow_tables_to_appear_in_same_query!(
    budgy_user,
    category,
    transaction,
);
