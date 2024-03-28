use sea_orm_migration::prelude::*;

use crate::m20240328_000002_create_transaction_type_table::TransactionType;

use super::m20240328_000001_create_budgy_user_table::BudgyUser;

pub struct Migration;

impl MigrationName for Migration {
    fn name(&self) -> &str {
        "m20240328_000003_create_transaction_table"
    }
}

#[async_trait::async_trait]
impl MigrationTrait for Migration {
    async fn up(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .create_table(
                Table::create()
                    .table(Transaction::Table)
                    .if_not_exists()
                    .col(
                        ColumnDef::new(Transaction::TransactionId)
                            .integer()
                            .not_null()
                            .auto_increment()
                            .primary_key(),
                    )
                    .col(ColumnDef::new(Transaction::Title).string().not_null())
                    .col(ColumnDef::new(Transaction::Description).string().not_null())
                    .col(ColumnDef::new(Transaction::Amount).float().not_null())
                    .col(ColumnDef::new(Transaction::BudgyUserId).integer().not_null())
                    .col(ColumnDef::new(Transaction::TransactionTypeId).integer().not_null())
                    .foreign_key(
                        ForeignKey::create()
                            .name("fk-transaction-user_id")
                            .from(Transaction::Table, Transaction::BudgyUserId)
                            .to(BudgyUser::Table, BudgyUser::BudgyUserId),
                    )
                    .foreign_key(
                        ForeignKey::create()
                            .name("fk-transaction-transaction_type_id")
                            .from(Transaction::Table, Transaction::TransactionTypeId)
                            .to(TransactionType::Table, TransactionType::TransactionTypeId),
                    )
                    .to_owned(),
            )
            .await
    }

    async fn down(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .drop_table(Table::drop().table(TransactionType::Table).to_owned())
            .await
    }
}

#[derive(DeriveIden)]
pub enum Transaction {
    Table,
    TransactionId,
    Title,
    Description,
    Amount,
    BudgyUserId,
    TransactionTypeId,
}
